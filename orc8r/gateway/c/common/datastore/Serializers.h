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
#pragma once

#include <functional>
#include <google/protobuf/message.h>

using google::protobuf::Message;
namespace magma {
// This file defines some common serializer methods

/**
 * Serialize a protobuf message into the standard string form
 */
std::function<bool(const Message&, std::string&, uint64_t&)>
get_proto_serializer();

/**
 * Deserialize a string into a protobuf message
 */
std::function<bool(const std::string&, Message&)>
get_proto_deserializer();

}
