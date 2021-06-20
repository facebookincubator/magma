---
#
# Copyright 2020 The Magma Authors.

# This source code is licensed under the BSD-style license found in the
# LICENSE file in the root directory of this source tree.

# Unless required by applicable law or agreed to in writing, software
# distributed under the License is distributed on an "AS IS" BASIS,
# WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
# See the License for the specific language governing permissions and
# limitations under the License.

log_level: INFO
print_grpc_payload: false

rule_update_inteval_sec: 1

# Session manager will report the usage when the usage is greater than
# usage_reporting_threshold * available quota since last update
# In this way, session manager will report the usage before the subscriber
# completely uses up the quota.
usage_reporting_threshold: 0.8

# Set to true to terminate service when the quota of a session is exhausted.
# An user can still use up to the extra margin.
# Set to false to allow users to use without any constraint.
terminate_service_when_quota_exhausted: true

# Maximum time to wait for the flow to be deleted by pipelined before forcefully
# terminating the session. This should be at least twice the poll interval of
# pipelined
session_force_termination_timeout_ms: 5000

# Set to true to enable sessiond support of carrier wifi
support_carrier_wifi: true

# CWF only
# Number of ms before sessiond terminates a session when it is started without
# quota.
cwf_quota_exhaustion_termination_on_init_ms: 30000

# If this flag is set, we will send the GW's timezone information to the policy
# component as part of CreateSessionRequest.
send_access_timezone: false

# Set to true to store session state in Redis for persistence.
support_stateless: {{ .Values.cwf.gateway_ha.enabled }}

# Redis table name for session state.
sessions_table: sessiond:sessions

# Sets the amount of octets that will be requested on the CCR-I for credit if
# credit is given. If unset, defaults to 200000
default_requested_units: 200000
