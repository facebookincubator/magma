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

package orchestrator

//Config represents the configuration provided to lte service
import (
	"magma/orc8r/cloud/go/services/analytics/calculations"
)

//Config represents the configuration provided to lte service
type Config struct {
	UseGRPCExporter           bool                         `yaml:"useGRPCExporter"`
	PrometheusGRPCPushAddress string                       `yaml:"prometheusGRPCPushAddress"`
	PrometheusPushAddresses   []string                     `yaml:"prometheusPushAddresses"`
	Analytics                 calculations.AnalyticsConfig `yaml:"analytics"`
}
