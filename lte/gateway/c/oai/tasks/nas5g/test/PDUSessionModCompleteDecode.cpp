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

/* using this stub code we are going to test Decoding functionality of
 * PDU Session Modification Request Message */

#include <iostream>
#include <M5GPDUSessionModificationComplete.h>

using namespace std;
using namespace magma5g;

namespace magma5g {
// Testing Decoding functionality
int decode(void) {
  int ret = 0;

  // Message to be decoded
  uint8_t buffer[] = {0x2e, 0x01, 0x02, 0xCC};
  int len          = 4;
  PDUSessionModificationCompleteMsg Req;

  // Decoding Message
  ret = Req.DecodePDUSessionModificationCompleteMsg(&Req, buffer, len);
  return ret;
}
}  // namespace magma5g

// Main function to call test decode function
int main(void) {
  int ret;
  ret = magma5g::decode();
  return ret;
}
