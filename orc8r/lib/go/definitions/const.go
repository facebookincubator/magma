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

// Package definitions defines consts, vars & types common to gateway & cloud
package definitions

const (
	MconfigStreamName = "configs"

	ControlProxyServiceName = "control_proxy"
	DispatcherServiceName   = "dispatcher"
	MagmadServiceName       = "magmad"
	MetricsdServiceName     = "metricsd"
	StateServiceName        = "state"
	StreamerServiceName     = "streamer"

	SQLiteDriver = "sqlite3"
)
