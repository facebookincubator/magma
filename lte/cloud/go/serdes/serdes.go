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

package serdes

import (
	"magma/lte/cloud/go/lte"
	lte_models "magma/lte/cloud/go/services/lte/obsidian/models"
	subscriberdb_models "magma/lte/cloud/go/services/subscriberdb/obsidian/models"
	"magma/orc8r/cloud/go/serde"
	"magma/orc8r/cloud/go/serdes"
	"magma/orc8r/cloud/go/services/state"
)

var (
	// LTEStateSerdes contains the LTE-specific state serdes
	LTEStateSerdes = serde.NewRegistry(
		state.NewStateSerde(lte.EnodebStateType, &lte_models.EnodebState{}),
		state.NewStateSerde(lte.ICMPStateType, &subscriberdb_models.IcmpStatus{}),

		// AGW state messages which use arbitrary untyped JSON serdes because
		// they're defined/used as protos in the AGW codebase
		state.NewStateSerde(lte.MMEStateType, &state.ArbitraryJSON{}),
		state.NewStateSerde(lte.SPGWStateType, &state.ArbitraryJSON{}),
		state.NewStateSerde(lte.S1APStateType, &state.ArbitraryJSON{}),
		state.NewStateSerde(lte.MobilitydStateType, &state.ArbitraryJSON{}),
	)

	// StateSerdes contains the full set of state serdes used in the LTE module
	StateSerdes = LTEStateSerdes.MustMerge(serdes.StateSerdes)
)
