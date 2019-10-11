"""
Copyright (c) 2016-present, Facebook, Inc.
All rights reserved.

This source code is licensed under the BSD-style license found in the
LICENSE file in the root directory of this source tree. An additional grant
of patent rights can be found in the PATENTS file in the same directory.
"""

import unittest
import s1ap_types
import time

from integ_tests.s1aptests import s1ap_wrapper
from integ_tests.s1aptests.s1ap_utils import SpgwUtil


class TestAttachDetachDedicatedInvalidEbi(unittest.TestCase):

    def setUp(self):
        self._s1ap_wrapper = s1ap_wrapper.TestWrapper()
        self._spgw_util = SpgwUtil()

    def tearDown(self):
        self._s1ap_wrapper.cleanup()

    def test_attach_detach(self):
        """ attach/detach + dedicated bearer deactivation with invalid EBI
            test with a single UE """
        num_ues = 1
        detach_type = [s1ap_types.ueDetachType_t.UE_NORMAL_DETACH.value,
                       s1ap_types.ueDetachType_t.UE_SWITCHOFF_DETACH.value]
        wait_for_s1 = [True, False]
        self._s1ap_wrapper.configUEDevice(num_ues)

        for i in range(num_ues):
            req = self._s1ap_wrapper.ue_req
            print("********************** Running End to End attach for ",
                  "UE id ", req.ue_id)
            # Now actually complete the attach
            self._s1ap_wrapper._s1_util.attach(
                req.ue_id, s1ap_types.tfwCmd.UE_END_TO_END_ATTACH_REQUEST,
                s1ap_types.tfwCmd.UE_ATTACH_ACCEPT_IND,
                s1ap_types.ueAttachAccept_t)

            # Wait on EMM Information from MME
            self._s1ap_wrapper._s1_util.receive_emm_info()

            time.sleep(5)
            print("********************** Adding dedicated bearer to IMSI",
                  ''.join([str(i) for i in req.imsi]))
            self._spgw_util.create_bearer(
                'IMSI' + ''.join([str(i) for i in req.imsi]), 5)

            response = self._s1ap_wrapper.s1_util.get_response()
            self.assertTrue(
                response, s1ap_types.tfwCmd.UE_ACT_DED_BER_REQ.value)
            act_ded_ber_ctxt_req = response.cast(
                s1ap_types.UeActDedBearCtxtReq_t)
            ded_bearer_acc = s1ap_types.UeActDedBearCtxtAcc_t()
            ded_bearer_acc.ue_Id = 1
            ded_bearer_acc.bearerId = act_ded_ber_ctxt_req.bearerId
            self._s1ap_wrapper._s1_util.issue_cmd(
                s1ap_types.tfwCmd.UE_ACT_DED_BER_ACC, ded_bearer_acc)

            time.sleep(5)
            print("********************** Deleting dedicated bearer for IMSI",
                  ''.join([str(i) for i in req.imsi]))
            # Deleting bearer with invalid EBI - 7
            self._spgw_util.delete_bearer(
                'IMSI' + ''.join([str(i) for i in req.imsi]), 5, 7)

            time.sleep(5)
            print("********************** Running UE detach for UE id ",
                  req.ue_id)
            # Now detach the UE
            self._s1ap_wrapper.s1_util.detach(
                req.ue_id, detach_type[i], wait_for_s1[i])


if __name__ == "__main__":
    unittest.main()
