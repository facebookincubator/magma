"""
Copyright (c) 2018-present, Facebook, Inc.
All rights reserved.

This source code is licensed under the BSD-style license found in the
LICENSE file in the root directory of this source tree. An additional grant
of patent rights can be found in the PATENTS file in the same directory.
"""

import unittest
from concurrent.futures import Future

import warnings
from lte.protos.mconfig.mconfigs_pb2 import PipelineD
from lte.protos.policydb_pb2 import FlowDescription, FlowMatch, PolicyRule
from magma.pipelined.app.gy import GYController
from magma.pipelined.bridge_util import BridgeTools
from magma.pipelined.policy_converters import flow_match_to_magma_match
from magma.pipelined.tests.app.flow_query import RyuDirectFlowQuery \
    as FlowQuery
from magma.pipelined.tests.app.packet_builder import IPPacketBuilder, \
    TCPPacketBuilder
from magma.pipelined.tests.app.packet_injector import ScapyPacketInjector
from magma.pipelined.tests.app.start_pipelined import PipelinedController, \
    TestSetup
from magma.pipelined.tests.app.subscriber import RyuDirectSubscriberContext
from magma.pipelined.tests.app.table_isolation import RyuDirectTableIsolator, \
    RyuForwardFlowArgsBuilder
from magma.pipelined.tests.pipelined_test_util import FlowTest, FlowVerifier, \
    PktsToSend, SubTest, create_service_manager, start_ryu_app_thread, \
    stop_ryu_app_thread, wait_after_send, SnapshotVerifier


class GYTableTest(unittest.TestCase):
    BRIDGE = 'testing_br'
    IFACE = 'testing_br'
    MAC_DEST = "5e:cc:cc:b1:49:4b"

    @classmethod
    def setUpClass(cls):
        """
        Starts the thread which launches ryu apps

        Create a testing bridge, add a port, setup the port interfaces. Then
        launch the ryu apps for testing pipelined. Gets the references
        to apps launched by using futures, mocks the redis policy_dictionary
        of gy_controller
        """
        super(GYTableTest, cls).setUpClass()
        warnings.simplefilter('ignore')
        cls._static_rule_dict = {}
        cls.service_manager = create_service_manager([PipelineD.ENFORCEMENT])
        cls._tbl_num = cls.service_manager.get_table_num(
            GYController.APP_NAME)

        gy_controller_reference = Future()
        testing_controller_reference = Future()
        test_setup = TestSetup(
            apps=[PipelinedController.GY,
                  PipelinedController.Testing,
                  PipelinedController.StartupFlows],
            references={
                PipelinedController.GY:
                    gy_controller_reference,
                PipelinedController.Testing:
                    testing_controller_reference,
                PipelinedController.StartupFlows:
                    Future(),
            },
            config={
                'bridge_name': cls.BRIDGE,
                'bridge_ip_address': '192.168.128.1',
                'nat_iface': 'eth2',
                'enodeb_iface': 'eth1',
                'enable_queue_pgm': False,
                'clean_restart': True,
            },
            mconfig=PipelineD(
                relay_enabled=True
            ),
            loop=None,
            service_manager=cls.service_manager,
            integ_test=False
        )

        BridgeTools.create_bridge(cls.BRIDGE, cls.IFACE)

        cls.thread = start_ryu_app_thread(test_setup)

        cls.gy_controller = gy_controller_reference.result()
        cls.testing_controller = testing_controller_reference.result()

        cls.gy_controller._policy_dict = cls._static_rule_dict

    @classmethod
    def tearDownClass(cls):
        stop_ryu_app_thread(cls.thread)
        BridgeTools.destroy_bridge(cls.BRIDGE)

    def test_subscriber_policy(self):
        """
        Add policy to subscriber, send 4096 packets

        Assert:
            Packets are properly matched with the 'simple_match' policy
            Send /20 (4096) packets, match /16 (256) packets
        """
        imsi = 'IMSI010000000088888'
        sub_ip = '192.168.128.74'
        flow_list1 = [FlowDescription(
            match=FlowMatch(
                ipv4_dst='45.10.0.0/24', direction=FlowMatch.UPLINK),
            action=FlowDescription.PERMIT)
        ]
        policies = [
            PolicyRule(id='simple_match', priority=2, flow_list=flow_list1)
        ]
        pkts_matched = 256
        pkts_sent = 4096

        self._static_rule_dict[policies[0].id] = policies[0]

        # ============================ Subscriber ============================
        sub_context = RyuDirectSubscriberContext(
            imsi, sub_ip, self.gy_controller, self._tbl_num
        ).add_static_rule(policies[0].id)
        isolator = RyuDirectTableIsolator(
            RyuForwardFlowArgsBuilder.from_subscriber(sub_context.cfg)
                                     .build_requests(),
            self.testing_controller
        )
        pkt_sender = ScapyPacketInjector(self.IFACE)
        packet = IPPacketBuilder()\
            .set_ip_layer('45.10.0.0/20', sub_ip)\
            .set_ether_layer(self.MAC_DEST, "00:00:00:00:00:00")\
            .build()
        flow_query = FlowQuery(
            self._tbl_num, self.testing_controller,
            match=flow_match_to_magma_match(flow_list1[0].match)
        )

        # =========================== Verification ===========================
        # Verify aggregate table stats, subscriber 1 'simple_match' pkt count
        flow_verifier = FlowVerifier([
            FlowTest(FlowQuery(self._tbl_num, self.testing_controller),
                     pkts_sent),
            FlowTest(flow_query, pkts_matched)
        ], lambda: wait_after_send(self.testing_controller))
        snapshot_verifier = SnapshotVerifier(self, self.BRIDGE,
                                             self.service_manager)

        with isolator, flow_verifier, snapshot_verifier:
        #with isolator, sub_context, flow_verifier, snapshot_verifier:
            pkt_sender.send(packet)

        flow_verifier.verify()

if __name__ == "__main__":
    unittest.main()
