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

package servicers

import (
	"context"

	"magma/feg/cloud/go/protos"
	"magma/feg/cloud/go/services/feg_relay/gw_to_feg_relay"
	"magma/orc8r/cloud/go/services/dispatcher/gateway_registry"
	orcprotos "magma/orc8r/lib/go/protos"
)

const FegS8Proxy gateway_registry.GwServiceType = "s8_proxy"

type S8RelayRouter struct {
	*gw_to_feg_relay.Router
}

// NewRelayRouter creates & returns a new RelayRouter
func NewS8RelayRouter(router *gw_to_feg_relay.Router) S8RelayRouter {
	if router == nil {
		router = &gw_to_feg_relay.Router{}
	}
	return S8RelayRouter{Router: router}
}

// S8ProxyServer implementation
//
func (s S8RelayRouter) CreateSession(
	c context.Context, r *protos.CreateSessionRequestPgw) (*protos.CreateSessionResponsePgw, error) {

	client, ctx, cancel, err := s.getS8Client(c, r.GetImsi())
	if err != nil {
		return nil, err
	}
	defer cancel()
	return client.CreateSession(ctx, r)
}

/*** TODO ***/
func (s S8RelayRouter) ModifyBearer(c context.Context, r *protos.ModifyBearerRequestPgw) (*protos.ModifyBearerResponsePgw, error) {
	return &protos.ModifyBearerResponsePgw{}, nil
}

func (s S8RelayRouter) DeleteSession(c context.Context, req *protos.DeleteSessionRequestPgw) (*protos.DeleteSessionResponsePgw, error) {
	return &protos.DeleteSessionResponsePgw{}, nil
}

func (s S8RelayRouter) SendEcho(c context.Context, req *orcprotos.Void) (*protos.EchoResponse, error) {
	return &protos.EchoResponse{}, nil
}

func (s S8RelayRouter) getS8Client(
	c context.Context, imsi string) (protos.S8ProxyClient, context.Context, context.CancelFunc, error) {

	conn, ctx, cancel, err := s.GetFegServiceConnection(c, imsi, FegS8Proxy)
	if err != nil {
		return nil, nil, nil, err
	}
	return protos.NewS8ProxyClient(conn), ctx, cancel, nil
}
