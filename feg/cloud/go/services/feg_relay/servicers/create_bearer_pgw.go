/*
Copyright 2021 The Magma Authors.

This source code is licensed under the BSD-style license found in the
LICENSE file in the root directory of this source tree.

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package servicers

import (
	"context"
	"fmt"

	fegprotos "magma/feg/cloud/go/protos"
	"magma/orc8r/cloud/go/services/dispatcher/gateway_registry"
	"magma/orc8r/lib/go/errors"
	orc8r_protos "magma/orc8r/lib/go/protos"

	"github.com/golang/glog"
)

const GTPCauseNotAvailable = 73 // gtp CauseNoResourcesAvailable

// CreateBearer relays the CreateBearerRequest from S8_proxy to a corresponding
// dispatcher service instance, who will in turn relay the request to the
// corresponding AGW gateway
func (srv *FegToGwRelayServer) CreateBearer(
	ctx context.Context,
	req *fegprotos.CreateBearerRequestPgw,
) (*orc8r_protos.Void, error) {
	if err := validateFegContext(ctx); err != nil {
		return nil, err
	}
	if req == nil {
		err := fmt.Errorf("unable to send CreateBearerPGW, request is nil: ")
		glog.Error(err)
		return nil, err
	}
	teid := fmt.Sprint(req.CAgwTeid)
	hwId, err := getHwIDFromTeid(ctx, teid)
	if err != nil {
		err = fmt.Errorf("unable to get HwID from TEID %s. err: %v", teid, err)
		glog.Error(err)
		if _, ok := err.(errors.ClientInitError); ok {
			// CauseNoResourcesAvailable uint8 = 73
			return nil, err
		}
	}
	conn, ctx, err := gateway_registry.GetGatewayConnection(
		gateway_registry.GwS8Service, hwId)
	if err != nil {
		err = fmt.Errorf("unable to get connection to the gateway ID: %s", hwId)
		return nil, err
	}
	client := fegprotos.NewS8ProxyResponderClient(conn)
	return client.CreateBearer(ctx, req)
}
