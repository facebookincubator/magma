"""
Copyright 2020 The Magma Authors.

This source code is licensed under the BSD-style license found in the
LICENSE file in the root directory of this source tree.

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
"""

import warnings
from typing import List
import subprocess
import unittest
from unittest import TestCase
import unittest.mock
from unittest.mock import MagicMock
import grpc
from concurrent import futures

from magma.pipelined.bridge_util import BridgeTools
from magma.pipelined.tests.app.start_pipelined import (
    TestSetup,
    PipelinedController)

from magma.pipelined.tests.pipelined_test_util import (
    start_ryu_app_thread,
    stop_ryu_app_thread,
    create_service_manager,
    wait_after_send)

import threading

from lte.protos import session_manager_pb2_grpc
from lte.protos.session_manager_pb2_grpc import SetInterfaceForUserPlaneStub
from lte.protos.session_manager_pb2 import UPFNodeState
from magma.pipelined.ng_manager.node_state_manager import NodeStateManager
from magma.pipelined.set_interface_client import send_node_state_association_request
from ryu.lib import hub
from orc8r.protos.common_pb2 import Void

class SMFAssociationServerTest(session_manager_pb2_grpc.SetInterfaceForUserPlaneServicer):

     def __init__ (self, loop):
         self._loop = loop

     def add_to_server(self, server):
        """
        Add the servicer to a gRPC server
        """
        session_manager_pb2_grpc.add_SetInterfaceForUserPlaneServicer_to_server(self, server)

     def SetUPFNodeState(self, request, context):
         return (Void())

class RpcTests(unittest.TestCase):
    """
    Tests NG Node related servicers
    """
    BRIDGE = 'testing_br'
    IFACE = 'testing_br'
    ASSOCIATED = 0

    def setUp(self):
        """
        Starts the thread which launches ryu apps

        Create a testing bridge, add a port, setup the port interfaces. Then
        launch the ryu apps for testing pipelined. Gets the references
        to apps launched by using futures.
        """
        warnings.simplefilter('ignore')

        loop_mock = MagicMock()

        # Bind the rpc server to a free port
        thread_pool = futures.ThreadPoolExecutor(max_workers=10)
        self._rpc_server = grpc.server(thread_pool)
        port = self._rpc_server.add_insecure_port('0.0.0.0:0')

        self._servicer = SMFAssociationServerTest(loop_mock)
        self._servicer.add_to_server(self._rpc_server)
        self._rpc_server.start()

        # Create a rpc stub
        channel = grpc.insecure_channel('0.0.0.0:{}'.format(port))
        self._channel = channel

        config_mock ={
                   'enodeb_iface': 'eth1',
                   'clean_restart': True,
                   '5G_feature_set': {'enable': True},
                   '5G_feature_set': {'node_identifier': '192.168.220.1'},
                   'bridge_name': self.BRIDGE,
               }

        self._ng_node_mgr = NodeStateManager(loop_mock, channel, config_mock)

    def tearDown(self):
        self._rpc_server.stop(0)


    def mock_sessiond_failure_case(self, node_message):
        return False

    def test_assoc_setup_message_request(self):
        node_mgr = self._ng_node_mgr

        if node_mgr._assoc_mon_thread:
            hub.kill(node_mgr._assoc_mon_thread)

        node_mgr._monitor_association()
        TestCase().assertEqual(node_mgr.assoc_message_count, 1)

    def test_assoc_setup_message_request_fail_attempt(self):
        node_mgr = self._ng_node_mgr

        if node_mgr._assoc_mon_thread:
            hub.kill(node_mgr._assoc_mon_thread)

        # Change the mock function to see if the send is failed
        node_mgr._sessiond_chan = None
        node_mgr._send_association_request_message(node_mgr._assoc_message)
        TestCase().assertEqual(node_mgr.assoc_message_count, 0)

        # Change the mock function to see if the send is successfull
        node_mgr._sessiond_chan = self._channel
        node_mgr._send_association_request_message(node_mgr._assoc_message)
        TestCase().assertEqual(node_mgr.assoc_message_count, 1)
