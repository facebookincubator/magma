/*
Copyright 2020 The Magma Authors.
This source code is licensed under the BSD-style license found in the
LICENSE file in the root directory of this source tree.
Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

#include <string>

#include "lte/protos/session_manager.pb.h"
#include "lte/protos/spgw_service.pb.h"
#include <folly/IPAddress.h>

extern "C" {
#include "spgw_service_handler.h"
#include "log.h"
}
#include "AmfServiceImpl.h"

namespace grpc {
class ServerContext;
}  // namespace grpc

using grpc::ServerContext;
using grpc::Status;
using magma::lte::SetSMSessionContextAccess;
using magma::lte::SmContextVoid;
using magma::lte::SmfPduSessionSmContext;

namespace magma {
using namespace lte;

AmfServiceImpl::AmfServiceImpl() {}
//Set message from SessionD received
Status AmfServiceImpl::SetSmfSessionContext(
    ServerContext* context, const SetSMSessionContextAccess* request,
    SmContextVoid* response) {
  OAILOG_INFO(LOG_UTIL, "Received GRPC SetSMSessionContextAccess request\n"); 
    
//ToDo processing ITTI,ZMQ

   return Status::OK;
 }

}  // namespace magma
