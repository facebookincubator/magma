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
#include <stdint.h>                      // for uint32_t
#include <atomic>                        // for atomic
#include <functional>                    // for function
#include <vector>                        // for vector
#include "lte/protos/subscriberdb.pb.h"  // for lte
namespace magma {
using namespace lte;

class PollStatsThread {
 public:
  /**
   * start_loop is the main function to call to initiate a load lv oop. Based
   * on the given loop interval length, this function will poll stats from
   * Pipelined and call the enforcer callback
   */
  void start_loop(
      std::shared_ptr<LocalEnforcer>&, uint32_t loop_interval_seconds);

  /**
   * Stop the config loop on the next loop
   */
  void stop();

 private:
  std::atomic<bool> is_running_;
};
}  // namespace magma