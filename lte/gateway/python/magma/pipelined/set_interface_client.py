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

import grpc
import logging

from magma.common.service_registry import ServiceRegistry
from lte.protos.session_manager_pb2 import UPFNodeState
from lte.protos.session_manager_pb2_grpc import SetInterfaceForUserPlaneStub

SESSION_SERVICE_NAME = "sessiond"

DEFAULT_GRPC_TIMEOUT = 5

def send_node_state_association_request(node_state_info: UPFNodeState,
                                        session_chan=None):
    """
    Make RPC call to send Node Association Setup/Release request to
    sessionD (SMF)
    """
    if session_chan is None:
        session_chan = ServiceRegistry.get_rpc_channel(SESSION_SERVICE_NAME,
                                                       ServiceRegistry.LOCAL)

    upf_stub = SetInterfaceForUserPlaneStub(session_chan)
    try:
        upf_stub.SetUPFNodeState(node_state_info, DEFAULT_GRPC_TIMEOUT)
        return True
    except grpc.RpcError as err:
        logging.error(
            "send_node_state_association_request error[%s] %s",
            err.code(),
            err.details())

        return False
