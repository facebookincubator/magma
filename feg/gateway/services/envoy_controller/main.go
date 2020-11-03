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

// Magma's Envoy Controller Service configures Envoy proxy with specified configuration
package main

import (
	"flag"

	"magma/feg/gateway/registry"
	"magma/feg/gateway/services/envoy_controller/servicers"
	"magma/orc8r/lib/go/service"

	"magma/feg/cloud/go/protos"

	lte_p "magma/lte/cloud/go/protos"

	"github.com/golang/glog"

	"magma/feg/gateway/services/envoy_controller/control_plane"
)

func init() {
    // Temp
    flag.Set("logtostderr", "true")
    flag.Set("stderrthreshold", "INFO")
    flag.Set("v", "2")
	flag.Parse()
}

func main() {
	// Create the service
	glog.Infof("Creating '%s' Service", registry.ENVOYD)
	srv, err := service.NewServiceWithOptions(registry.ModuleName, registry.ENVOYD)
	if err != nil {
		glog.Fatalf("Error creating Envoy Controller service: %s", err)
	}

	cli := control_plane.GetControllerClient()


	// Create servicers
	servicer := servicers.NewEnvoydService(cli)


	// Register services
	protos.RegisterEnvoydServer(srv.GrpcServer, servicer)

	test_proto := []*protos.AddUEHeaderEnrichmentRequest{{
		UeIp: &lte_p.IPAddress{
			Version: lte_p.IPAddress_IPV4,
			Address: []byte("1.2.3.3"),
		},
		Websites: []string{"neverssl.com", "google.com"},
		Headers: []*protos.Header{{
			Name:  "IMSI",
			Value: "024212312312",
		}},
	},
		{
			UeIp: &lte_p.IPAddress{
				Version: lte_p.IPAddress_IPV4,
				Address: []byte("2.2.2.2"),
			},
			Websites: []string{"magma.com", "qqq.com"},
			Headers: []*protos.Header{{
				Name:  "IMSI",
				Value: "111111",
			},
				{
					Name:  "MSISDN",
					Value: "THIS_IS_MSISDN",
				}},
		}}
	cli.UpdateSnapshot(test_proto)

	// Run the service
	err = srv.Run()
	if err != nil {
		glog.Fatalf("Error running service: %s", err)
	}
	glog.Infof("Starting '%s' Service", registry.ENVOYD)
}
