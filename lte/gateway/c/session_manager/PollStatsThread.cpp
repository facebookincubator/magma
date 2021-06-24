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
/**
 * PollStats is used to request stats from Pipelined
 */
class PollStatsThread {
 public:
  /**
   * start_loop is the main function to call to initiate a load loop. Based on
   * the given loop interval length, this function will  and call the processor
   * callback.
   */
  void start_loop(
      std::shared_ptr<LocalEnforcer>&, uint32_t loop_interval_seconds) {
    is_running_ = true;
    while (is_running_) {
      // call local enforcer(arguments might need to be passed in)
      // call Pipelined Client Poll Stats(need to receive arguments from
      // some location)
      printf("Calling pipelined service\n");
      int cookie      = 0;
      int cookie_mask = 0;
      enforcer->PollStats();
      std::this_thread::sleep_for(std::chrono::seconds(loop_interval_seconds));
    }
  }

  /**
   * Stop the config loop on the next loop
   */
  void stop() { is_running_ = false; }

 private:
  std::atomic<bool> is_running_;
};
}  // namespace magma