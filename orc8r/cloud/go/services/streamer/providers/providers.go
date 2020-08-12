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

package providers

import (
	"encoding/json"
	"time"

	"magma/orc8r/cloud/go/services/configurator"
	"magma/orc8r/lib/go/definitions"
	"magma/orc8r/lib/go/protos"

	"github.com/golang/protobuf/ptypes"
	"github.com/golang/protobuf/ptypes/any"
	"github.com/golang/protobuf/ptypes/wrappers"
	"github.com/pkg/errors"
)

// MconfigProvider provides streamer mconfigs (magma configs).
type MconfigProvider struct{}

func (p *MconfigProvider) GetStreamName() string {
	return definitions.MconfigStreamName
}

func (p *MconfigProvider) GetUpdates(gatewayId string, extraArgs *any.Any) ([]*protos.DataUpdate, error) {
	res, err := configurator.GetMconfigFor(gatewayId)
	if err != nil {
		return nil, errors.Wrap(err, "get mconfig from configurator")
	}

	// TODO(T71525030): add back below support for mconfig digests
	//if extraArgs != nil {
	//	// Currently, only use of extraArgs is mconfig hash
	//	receivedDigest := &protos.GatewayConfigsDigest{}
	//	if err := ptypes.UnmarshalAny(extraArgs, receivedDigest); err == nil {
	//		glog.V(2).Infof(
	//			"Received, generated config digests: %v, %v\n",
	//			receivedDigest,
	//			res.Configs.Metadata.Digest.Md5HexDigest,
	//		)
	//		return mconfigToUpdate(res.Configs, res.LogicalID, receivedDigest.Md5HexDigest)
	//	}
	//}
	//

	return mconfigToUpdate(res.Configs, res.LogicalID, "")
}

func mconfigToUpdate(configs *protos.GatewayConfigs, logicalID string, digest string) ([]*protos.DataUpdate, error) {
	// TODO(T71525030): add back the digest equality check, and generally revert this fn

	// Early/empty return if gateway already has this config
	//if digest == configs.Metadata.Digest.Md5HexDigest {
	//	return []*protos.DataUpdate{}, nil
	//}

	//marshaledConfig, err := protos.MarshalJSON(configs)
	//if err != nil {
	//	return nil, errors.Wrap(err, "marshal gateway mconfig")
	//}

	marshaledConfig, err := marshalJSONConfigs(configs)
	if err != nil {
		return nil, err
	}
	return []*protos.DataUpdate{{Key: logicalID, Value: marshaledConfig}}, nil
}

type mconfigTemplate struct {
	ConfigsByKey map[string]json.RawMessage    `json:"configsByKey"`
	Metadata     protos.GatewayConfigsMetadata `json:"metadata"`
}

// marshalJSONConfigs marshals gateway configs by "manually" constructing the
// JSON-marshaled bytes for the data update.
// This hack is a temporary workaround until we upgrade to Go's protobuf
// library APIv2, which has built-in support for dynamic any.Any resolution,
// which is required to marshal an any.Any proto to JSON.
func marshalJSONConfigs(configs *protos.GatewayConfigs) ([]byte, error) {
	configsByKey := map[string]json.RawMessage{}

	for k, v := range configs.ConfigsByKey {
		bytesVal := &wrappers.BytesValue{}
		err := ptypes.UnmarshalAny(v, bytesVal)
		if err != nil {
			return nil, err
		}
		configsByKey[k] = bytesVal.Value
	}

	mconfig := &mconfigTemplate{
		ConfigsByKey: configsByKey,
		Metadata:     protos.GatewayConfigsMetadata{CreatedAt: uint64(time.Now().Unix())},
	}
	marshaledMconfig, err := json.Marshal(mconfig)
	if err != nil {
		return nil, err
	}

	return marshaledMconfig, nil
}
