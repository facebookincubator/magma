"""
Copyright (c) 2016-present, Facebook, Inc.
All rights reserved.

This source code is licensed under the BSD-style license found in the
LICENSE file in the root directory of this source tree. An additional grant
of patent rights can be found in the PATENTS file in the same directory.
"""

import unittest
import ctypes
import s1ap_types
from integ_tests.s1aptests.s1ap_utils import S1ApUtil


class TestS1SetupIncorrectPlmn(unittest.TestCase):
    def setUp(self):
        self._s1_util = S1ApUtil()

    def tearDown(self):
        print("************************* Sending SCTP SHUTDOWN")
        self._s1_util.issue_cmd(s1ap_types.tfwCmd.SCTP_SHUTDOWN_REQ, None)
        self._s1_util.cleanup()

    def test_s1setup_incorrect_plmn(self):
        """ S1 Setup with incorrect plmn ID """

        print("************************* Enb tester configuration")
        req = s1ap_types.FwNbConfigReq_t()
        req.cellId_pr.pres = True
        req.cellId_pr.cell_id = 10
        req.plmnId_pr.pres = True
        # Convert PLMN to ASCII character array of MCC and MNC digits
        # For 5 digit PLMN add \0 in the end, e.g., "00101\0"
        req.plmnId_pr.plmn_id = (ctypes.c_ubyte * 6).from_buffer_copy(
            bytearray(b"333333")
        )

        print("************************* Sending ENB configuration Request")
        assert self._s1_util.issue_cmd(s1ap_types.tfwCmd.ENB_CONFIG, req) == 0
        response = self._s1_util.get_response()
        assert response.msg_type == s1ap_types.tfwCmd.ENB_CONFIG_CONFIRM.value
        res = response.cast(s1ap_types.FwNbConfigCfm_t)
        assert res.status == s1ap_types.CfgStatus.CFG_DONE.value

        print("************************* Sending S1-setup Request")
        req = None
        assert (
            self._s1_util.issue_cmd(s1ap_types.tfwCmd.ENB_S1_SETUP_REQ, req)
            == 0
        )
        response = self._s1_util.get_response()
        assert response.msg_type == s1ap_types.tfwCmd.ENB_S1_SETUP_RESP.value
        res = response.cast(s1ap_types.FwNbS1setupRsp_t)
        assert res.res == s1ap_types.S1_setp_Result.S1_SETUP_FAILED.value


if __name__ == "__main__":
    unittest.main()
