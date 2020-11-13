/*
 * Licensed to the OpenAirInterface (OAI) Software Alliance under one or more
 * contributor license agreements.  See the NOTICE file distributed with
 * this work for additional information regarding copyright ownership.
 * The OpenAirInterface Software Alliance licenses this file to You under
 * the terms found in the LICENSE file in the root of this source tree.
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 *-------------------------------------------------------------------------------
 * For more information about the OpenAirInterface (OAI) Software Alliance:
 *      contact@openairinterface.org
 */

extern "C" {
#include "sctpd_uplink_server.h"

// #include "assertions.h"
#include "bstrlib.h"
#include "log.h"

#include "sctp_defs.h"
#include "sctp_itti_messaging.h"
}

#include <memory>

#include <grpcpp/grpcpp.h>

#include <lte/protos/sctpd.grpc.pb.h>

namespace magma {
namespace mme {

using grpc::ServerContext;
using grpc::Status;

using magma::sctpd::CloseAssocReq;
using magma::sctpd::CloseAssocRes;
using magma::sctpd::NewAssocReq;
using magma::sctpd::NewAssocRes;
using magma::sctpd::SctpdUplink;
using magma::sctpd::SendUlReq;
using magma::sctpd::SendUlRes;

class SctpdUplinkImpl final : public SctpdUplink::Service {
 public:
  SctpdUplinkImpl();

  Status SendUl(
      ServerContext* context, const SendUlReq* req, SendUlRes* res) override;
  Status NewAssoc(
      ServerContext* context, const NewAssocReq* req,
      NewAssocRes* res) override;
  Status CloseAssoc(
      ServerContext* context, const CloseAssocReq* req,
      CloseAssocRes* res) override;
};

SctpdUplinkImpl::SctpdUplinkImpl() {}

Status SctpdUplinkImpl::SendUl(
    ServerContext* context, const SendUlReq* req, SendUlRes* res) {
  bstring payload;
  uint32_t ppid, i ;
  uint32_t assoc_id;
  uint16_t stream;

/*initial_ue_msg */

const char *new_data ="\x00\x0f\x40\x3d\x00\x00\x04\x00\x55\x00\x02\x00\x01\x00\x26\x00\x18\x17\x7e\x00\x41\x71\x00\x0d\x01\x13\x00\x14\xf0\xff\x00\x00\x01\x00\x00\x00\xf1\x2e\x02\x80\x40\x00\x79\x00\x0f\x40\x13\x40\x01\x00\x00\x00\x5c\xd0\x13\x40\x01\x00\xa0\x00\x00\x5a\x40\x01\x18";

/*Uplink Nas transport reponse*/
//const char *new_data = "\x00\x2e\x40\x3c\x00\x00\x04\x00\x0a\x00\x02\x00\x01\x00\x55\x00\x02\x00\x01\x00\x26\x00\x16\x15\x7e\x00\x57\x2d\x10\x25\xe8\x7b\x06\x52\xc3\xc6\x3b\x36\x82\x8b\x54\x51\x7e\xbf\x15\x00\x79\x40\x0f\x40\x13\x40\x01\x00\x00\x00\x5c\xd0\x13\x40\x01\x00\xa0\x00"; 

/*Ngsetup */
//const char *new_data = "\x00\x15\x00\x5e\x00\x00\x04\x00\x1b\x00\x09\x00\x02\xf8\x59\x50\x00\x0f\xff\xff\x00\x52\x40\x07\x02\x00\x67\x6e\x62\x2d\x31\x00\x66\x00\x3a\x01\x00\x00\xa0\x00\x10\x02\xf7\x45\x00\x00\x10\x80\x58\x39\x22\x00\x02\xf8\x59\x00\x01\x11\x38\x38\x53\x82\x16\xf0\x12\x30\x00\x00\x00\xa0\x01\x10\x02\x03\x10\x00\x00\x17\xf8\x12\x34\x83\x00\x02\x12\x00\x00\x00\x11\x18\x32\x12\x36\x00\x15\x40\x01\x00";



 payload = blk2bstr((unsigned char *)new_data, req->payload().size());
  if (payload == NULL) {
    OAILOG_ERROR(LOG_SCTP, "failed to allocate bstr for SendUl\n");
    return Status::OK;
  
}

  ppid=req->ppid();
  assoc_id = req->assoc_id();
  stream   = req->stream();


  if (sctp_itti_send_new_message_ind(&payload,ppid, assoc_id, stream) < 0) {
    OAILOG_ERROR(LOG_SCTP, "failed to send new_message_ind for SendUl\n");
    return Status::OK;
  }

  return Status::OK;
}

#include <assert.h>

Status SctpdUplinkImpl::NewAssoc(
    ServerContext* context, const NewAssocReq* req, NewAssocRes* res) {

  uint32_t ppid;
  uint32_t assoc_id;
  uint16_t instreams;
  uint16_t outstreams;

  ppid	     =req->ppid();
  assoc_id   = req->assoc_id();
  instreams  = req->instreams();
  outstreams = req->outstreams();

 
 if (sctp_itti_send_new_association(ppid, assoc_id, instreams, outstreams) < 0) {
    return Status::OK;
  }

  return Status::OK;
}

Status SctpdUplinkImpl::CloseAssoc(
    ServerContext* context, const CloseAssocReq* req, CloseAssocRes* res) {
  uint32_t ppid;
  uint32_t assoc_id;
  bool reset;

  ppid	   = req->ppid();
  assoc_id = req->assoc_id();
  reset    = req->is_reset();

  if (sctp_itti_send_com_down_ind(ppid, assoc_id, reset) < 0) {
    OAILOG_ERROR(LOG_SCTP, "failed to send com_down_ind for CloseAssoc\n");
    return Status::OK;
  }

  return Status::OK;
}

}  // namespace mme
}  // namespace magma

using grpc::Server;
using grpc::ServerBuilder;

using magma::mme::SctpdUplinkImpl;

std::shared_ptr<SctpdUplinkImpl> _service = nullptr;
std::unique_ptr<Server> _server           = nullptr;

int start_sctpd_uplink_server(void) {
  _service = std::make_shared<SctpdUplinkImpl>();

  ServerBuilder builder;
  builder.AddListeningPort(UPSTREAM_SOCK, grpc::InsecureServerCredentials());
  builder.RegisterService(_service.get());

  _server = builder.BuildAndStart();

  return 0;
}

void stop_sctpd_uplink_server(void) {
  if (_server != nullptr) {
    _server->Shutdown();
    _server->Wait();
    _server = nullptr;
  }
  _service = nullptr;
}
