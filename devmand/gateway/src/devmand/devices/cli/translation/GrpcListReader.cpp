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

#include <devmand/devices/cli/translation/GrpcListReader.h>

namespace devmand {
namespace devices {
namespace cli {

using namespace folly;
using namespace std;
using namespace grpc;
using namespace devmand::channels::cli::plugin;
using namespace devmand::channels::cli;

GrpcListReader::GrpcListReader(
    shared_ptr<grpc::Channel> channel,
    const string _id,
    shared_ptr<Executor> _executor)
    : GrpcCliHandler(_id, _executor),
      stub_(devmand::channels::cli::plugin::ReaderPlugin::NewStub(channel)) {}

Future<vector<dynamic>> GrpcListReader::readKeys(
    const Path& path,
    const DeviceAccess& device) const {
  MLOG(MDEBUG) << "[" << id << "] readKeys " << path;
  ActualReadRequest* actualRequest = new ActualReadRequest();
  actualRequest->set_path(path.str());
  ReadRequest request;
  request.set_allocated_actualreadrequest(actualRequest);

  return finish<ReadRequest, ReadResponse, vector<dynamic>>(
      request,
      device,
      [this](auto context) { return stub_->Read(context); },
      [this](auto response) -> vector<dynamic> {
        dynamic result = parseJson(response.actualreadresponse().json());
        if (not result.isArray()) {
          MLOG(MERROR) << "[" << id << "] Response is not json array:"
                       << response.actualreadresponse().json();
          throw runtime_error("Response is not json array");
        }
        vector<dynamic> values;
        for (auto& k : result) {
          MLOG(MDEBUG) << "pushing " << toJson(k);
          values.push_back(k);
        }
        return values;
      });
}

} // namespace cli
} // namespace devices
} // namespace devmand
