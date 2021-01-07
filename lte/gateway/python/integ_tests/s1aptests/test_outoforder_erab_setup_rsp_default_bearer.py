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

import unittest
import time

import s1ap_types
import s1ap_wrapper
import ipaddress


class TestOutOfOrderErabSetupRspDefaultBearer(unittest.TestCase):
    def setUp(self):
        self._s1ap_wrapper = s1ap_wrapper.TestWrapper()

    def tearDown(self):
        self._s1ap_wrapper.cleanup()

    def test_outoforder_erab_setup_rsp_default_bearer(self):
        """ Attach a single UE and send erab setup rsp message
        out of order for the secondary pdn"""
        num_ue = 1

        self._s1ap_wrapper.configUEDevice(num_ue)
        req = self._s1ap_wrapper.ue_req
        ue_id = req.ue_id

        # APN of the secondary PDN
        ims = {
            "apn_name": "ims",  # APN-name
            "qci": 5,  # qci
            "priority": 15,  # priority
            "pre_cap": 0,  # preemption-capability
            "pre_vul": 0,  # preemption-vulnerability
            "mbr_ul": 200000000,  # MBR UL
            "mbr_dl": 100000000,  # MBR DL
        }

        # APN list to be configured
        apn_list = [ims]

        self._s1ap_wrapper.configAPN(
            "IMSI" + "".join([str(i) for i in req.imsi]), apn_list
        )
        print(
            "************************* Running End to End attach for UE id ",
            ue_id,
        )
        # Attach
        attach = self._s1ap_wrapper.s1_util.attach(
            ue_id,
            s1ap_types.tfwCmd.UE_END_TO_END_ATTACH_REQUEST,
            s1ap_types.tfwCmd.UE_ATTACH_ACCEPT_IND,
            s1ap_types.ueAttachAccept_t,
        )

        # Wait on EMM Information from MME
        self._s1ap_wrapper._s1_util.receive_emm_info()

        addr = attach.esmInfo.pAddr.addrInfo
        default_ip = ipaddress.ip_address(bytes(addr[:4]))

        # Send indication to delay sending of erab setup rsp
        delay_erab_setup_resp = s1ap_types.UeDelayErabSetupRsp()
        delay_erab_setup_resp.ue_Id = ue_id
        delay_erab_setup_resp.flag = 1
        # Timer value in secs to delay erab setup rsp
        delay_erab_setup_resp.tmrVal = 6000
        self._s1ap_wrapper._s1_util.issue_cmd(
            s1ap_types.tfwCmd.UE_SET_DELAY_ERAB_SETUP_RSP,
            delay_erab_setup_resp,
        )
        print(
            "Sent UE_SET_DELAY_ERAB_SETUP_RSP with delay value of %d secs"
            % (delay_erab_setup_resp.tmrVal)
        )

        # Send PDN Connectivity Request
        apn = "ims"
        self._s1ap_wrapper.sendPdnConnectivityReq(ue_id, apn)
        # Receive PDN CONN RSP/Activate default EPS bearer context request
        response = self._s1ap_wrapper.s1_util.get_response()
        self.assertEqual(
            response.msg_type, s1ap_types.tfwCmd.UE_PDN_CONN_RSP_IND.value
        )
        act_def_bearer_req = response.cast(s1ap_types.uePdnConRsp_t)
        addr = act_def_bearer_req.m.pdnInfo.pAddr.addrInfo
        sec_ip = ipaddress.ip_address(bytes(addr[:4]))

        print(
            "************************* Sending Activate default EPS bearer "
            "context accept for UE id ",
            ue_id,
        )

        # Delay to ensure erab setup rsp is sent out of order
        print("Sleeping for 10 seconds")
        time.sleep(10)

        dl_flow_rules = {
            default_ip: [],
            sec_ip: [],
        }
        # default bearer + ims bearer
        num_ul_flows = 2
        # Verify if flow rules are created
        self._s1ap_wrapper.s1_util.verify_flow_rules(
            num_ul_flows, dl_flow_rules
        )

        # Send PDN Disconnect
        pdn_disconnect_req = s1ap_types.uepdnDisconnectReq_t()
        pdn_disconnect_req.ue_Id = ue_id
        pdn_disconnect_req.epsBearerId = (
            act_def_bearer_req.m.pdnInfo.epsBearerId
        )
        self._s1ap_wrapper._s1_util.issue_cmd(
            s1ap_types.tfwCmd.UE_PDN_DISCONNECT_REQ, pdn_disconnect_req
        )

        # Receive UE_DEACTIVATE_BER_REQ
        response = self._s1ap_wrapper.s1_util.get_response()
        self.assertEqual(
            response.msg_type, s1ap_types.tfwCmd.UE_DEACTIVATE_BER_REQ.value
        )

        print(
            "******************* Received deactivate eps bearer context"
            " request"
        )
        # Send DeactDedicatedBearerAccept
        deactv_bearer_req = response.cast(s1ap_types.UeDeActvBearCtxtReq_t)
        self._s1ap_wrapper.sendDeactDedicatedBearerAccept(
            ue_id, deactv_bearer_req.bearerId
        )

        print(
            "************************* Running UE detach (switch-off) for ",
            "UE id ",
            ue_id,
        )
        # Now detach the UE
        self._s1ap_wrapper.s1_util.detach(
            ue_id, s1ap_types.ueDetachType_t.UE_SWITCHOFF_DETACH.value, False
        )


if __name__ == "__main__":
    unittest.main()
