/*
 * Copyright 2020 The Magma Authors.
 * This source code is licensed under the BSD-style license found in the
 * LICENSE file in the root directory of this source tree.
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 */
/*****************************************************************************
  Source      	SetMessageManagerHandler.h
  Version     	0.1
  Date       	2020/08/08
  Product     	SessionD
  Subsystem   	5G Landing object in SessionD
  Author/Editor Sanjay Kumar Ojha
  Description 	Defines Access and Mobility Management Messages
*****************************************************************************/
#pragma once
#include <functional>

#include <grpc++/grpc++.h>
#include <lte/protos/session_manager.grpc.pb.h>

#include "SessionStateEnforcer.h"
#include "SessionID.h"
//#include "SessionReporter.h"
//#include "SessionStore.h"

using grpc::Server;
using grpc::ServerContext;
using grpc::Status;
namespace magma {
using namespace orc8r;
/* SetMessageManagerHandler processes gRPC requests for the sessionD
 * This composites the earlier LocalSessionManagerHandlerImpl and uses the
 * exiting functionalities.
 */

/* This the landing object of 5G gRPC call by set message*/
class SetMessageManager
{
 public:

 virtual ~SetMessageManager() {}

   /* RPC call from AMF "rpc SetAmfSessionContext (SetSMSessionContext) returns (SmContextVoid);" 
    * as its set-interface API, no need to send response back, response is void and gRPC
    * will take care on acknowledgement 
    */
   //void SetAmfSessionContext(ServerContext* context, const SetSMSessionContext* request);
   virtual void SetAmfSessionContext(ServerContext* context, const SetSMSessionContext* request,
		          std::function<void(Status, SmContextVoid)> response_callback) = 0;
}; //end of abstract class
class SetMessageManagerHandler: public SetMessageManager {
public:
      SetMessageManagerHandler(	//recheck on the attributes of constructor TODO.
      std::shared_ptr<SessionStateEnforcer> m5G_monitor,
      /*SessionReporter* reporter,
      std::shared_ptr<AsyncDirectorydClient> directoryd_client,*/
      SessionStore& session_store);
      ~SetMessageManagerHandler() {}

      virtual void SetAmfSessionContext(ServerContext* context, const SetSMSessionContext* request,
		          std::function<void(Status, SmContextVoid)> response_callback);
    /*
    * Send session creation related request to the CentralSessionController.
    * which is policy/QoS related. On successful, creates and populate, 
    * session_map in memoery and response set message to AMF by gRPC.
    * It uses SessionStateEnforcer object to create new session state.
    */
   //void send_create_session
   /*void send_create_session(SessionMap& session_map,
		   const SetSMSessionContext& request,
		   const std::string& imsi, const std::string& session_ctx_id,
		   const SessionConfig& cfg);*/
   /*initialize the session message from proto message
   SessionConfig m5g_build_session_config(const SetSMSessionContext& request);*/

private:
   SessionStore& session_store_;
   std::shared_ptr<SessionStateEnforcer> m5g_enforcer_;
   SessionIDGenerator id_gen_;

   void send_create_session(SessionMap& session_map,
                   /*const SetSMSessionContext& request,*/
                   const std::string& imsi, const std::string& session_ctx_id,
                   const SessionConfig& cfg);
   /*initialize the session message from proto message*/
   SessionConfig m5g_build_session_config(const SetSMSessionContext& request);


}; // of class SetMessageManagerHandlerImpl
}//end namespace magma



