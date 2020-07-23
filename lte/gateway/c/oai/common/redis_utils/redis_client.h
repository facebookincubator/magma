/*
 * Licensed to the OpenAirInterface (OAI) Software Alliance under one or more
 * contributor license agreements.  See the NOTICE file distributed with
 * this work for additional information regarding copyright ownership.
 * The OpenAirInterface Software Alliance licenses this file to You under
 * the terms found in the LICENSE file in the root of this
 * source tree.
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 *------------------------------------------------------------------------------
 * For more information about the OpenAirInterface (OAI) Software Alliance:
 *      contact@openairinterface.org
 */

#pragma once

#include <string>

#include <cpp_redis/cpp_redis>
#include <google/protobuf/message.h>

#include "orc8r/protos/redis.pb.h"

namespace magma {
namespace lte {

class RedisClient {
 public:
  RedisClient();
  ~RedisClient() = default;

  /**
   * Initializes a connection to the redis datastore configured in redis.yml
   * @return response code of success / error with db connection
   */
  void init_db_connection();

  /**
   * Returns the value on redis db mapped to a key
   * @param key
   * @return string repr of value
   */
  std::string read(const std::string& key);

  /**
   * Writes a str value to redis mapped to str key
   * @param key
   * @param value
   * @return response code of operation
   */
  int write(const std::string& key, const std::string& value);

  /**
   * Writes a protobuf object to redis
   * @param key
   * @param proto_msg
   * @return response code of operation
   */
  int write_proto(
      const std::string& key, const google::protobuf::Message& proto_msg);

  /**
   * Reads value from redis mapped to key and returns proto object
   * @param key
   * @return response code of operation
   */
  int read_proto(const std::string& key, google::protobuf::Message& proto_msg);

  int clear_keys(const std::vector<std::string>& keys_to_clear);

  std::vector<std::string> get_keys(const std::string& pattern);

  bool is_connected() { return is_connected_; }

 private:
  std::unique_ptr<cpp_redis::client> db_client_;
  bool is_connected_;

  /**
   * Read the wrapper RedisState value from Redis for a key
   * @param key
   * @param state_out
   * @return
   */
  int read_redis_state(const std::string& key, orc8r::RedisState& state_out);

  /**
   * Check for existence of a key in Redis.
   * @param key
   * @throws std::runtime_error if the Redis call fails
   * @return
   */
  bool key_exists(const std::string& key);

  /**
   * Converts protobuf Message and parses it to string
   * @param proto_msg
   * @param str_to_serialize
   */
  int serialize(
      const google::protobuf::Message& proto_msg,
      std::string& str_to_serialize);
  /**
   * Takes a string and parses it to protobuf Message
   * @param proto_msg
   * @param str_to_deserialize
   */
  int deserialize(
      google::protobuf::Message& proto_msg,
      const std::string& str_to_deserialize);
};

}  // namespace lte
}  // namespace magma
