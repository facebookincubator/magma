/**
 * Copyright 2020 The Magma Authors.
 *
 * This source code is licensed under the BSD-style license found in the
 * LICENSE file in the root directory of this source tree.
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 */

// WARNING: Do not include this header directly. Use intertask_interface.h
// instead.

MESSAGE_DEF(
    AMF_APP_INITIAL_CONTEXT_SETUP_RSP, itti_amf_app_initial_context_setup_rsp_t,
    amf_app_initial_context_setup_rsp)
MESSAGE_DEF(
    AMF_APP_INITIAL_CONTEXT_SETUP_FAILURE,
    itti_amf_app_initial_context_setup_failure_t,
    amf_app_initial_context_setup_failure)
MESSAGE_DEF(
    AMF_APP_NGAP_AMF_UE_ID_NOTIFICATION,
    itti_amf_app_ngap_amf_ue_id_notification_t,
    amf_app_ngap_amf_ue_id_notification)
MESSAGE_DEF(
    AMF_APP_UPLINK_DATA_IND, itti_amf_app_ul_data_ind_t, amf_app_ul_data_ind)
MESSAGE_DEF(
    AMF_APP_DOWNLINK_DATA_REJ, itti_amf_app_dl_data_rej_t, amf_app_dl_data_rej)
