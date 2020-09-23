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
#include <iostream>
#include <sstream>
#include <cstdint>
#include "MessageType.h"
#include "CommonDefs.h"
using namespace std;
namespace magma5g {
MessageTypeMsg::MessageTypeMsg(){};

MessageTypeMsg::~MessageTypeMsg(){};

// Decode MessageType IE
int MessageTypeMsg::DecodeMessageTypeMsg(
    MessageTypeMsg* messagetype, uint8_t iei, uint8_t* buffer, uint32_t len) {
  uint8_t decoded = 0;

  MLOG(MDEBUG) << "   DecodeMessageTypeMsg : ";
  messagetype->msgtype = *(buffer + decoded);
  decoded++;
  MLOG(MDEBUG) << " msgtype = 0x" << hex << int(messagetype->msgtype) << "\n";
  return (decoded);
};

// Encode MessageType IE
int MessageTypeMsg::EncodeMessageTypeMsg(
    MessageTypeMsg* messagetype, uint8_t iei, uint8_t* buffer, uint32_t len) {
  uint8_t encoded = 0;

  MLOG(MDEBUG) << " EncodeMessageTypeMsg : ";
  *(buffer + encoded) = messagetype->msgtype;
  MLOG(MDEBUG) << "Message type = 0x" << hex << int(*(buffer + encoded))
               << "\n";
  encoded++;
  return (encoded);
};
}  // namespace magma5g
