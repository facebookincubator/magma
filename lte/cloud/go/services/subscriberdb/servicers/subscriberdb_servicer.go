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
	"sort"

	"magma/lte/cloud/go/services/subscriberdb"
	"magma/orc8r/cloud/go/mproto"

	"github.com/golang/glog"
	"github.com/golang/protobuf/proto"
	"github.com/pkg/errors"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	"magma/lte/cloud/go/lte"
	lte_protos "magma/lte/cloud/go/protos"
	"magma/lte/cloud/go/serdes"
	lte_models "magma/lte/cloud/go/services/lte/obsidian/models"
	"magma/lte/cloud/go/services/subscriberdb/obsidian/models"
	"magma/orc8r/cloud/go/services/configurator"
	"magma/orc8r/lib/go/protos"
)

type subscriberdbServicer struct {
	flatDigestEnabled bool
}

const defaultSubProfile = "default"

func NewSubscriberdbServicer(config subscriberdb.Config) lte_protos.SubscriberDBCloudServer {
	return &subscriberdbServicer{flatDigestEnabled: config.FlatDigestEnabled}
}

// ListSubscribers returns a page of subscribers and a token to be used on
// subsequent requests. The page token specified in the request is used to
// determine the first subscriber to include in the page. The page size
// specified in the request determines the maximum number of entities to
// return. If no page size is specified, the maximum size configured in the
// configurator service will be returned.
func (s *subscriberdbServicer) ListSubscribers(ctx context.Context, req *lte_protos.ListSubscribersRequest) (*lte_protos.ListSubscribersResponse, error) {
	gateway := protos.GetClientGateway(ctx)
	if gateway == nil {
		return nil, status.Errorf(codes.PermissionDenied, "missing gateway identity")
	}
	if !gateway.Registered() {
		return nil, status.Errorf(codes.PermissionDenied, "gateway is not registered")
	}
	networkID := gateway.NetworkId
	gatewayID := gateway.LogicalId

	lteGateway, err := configurator.LoadEntity(
		networkID, lte.CellularGatewayEntityType, gatewayID,
		configurator.EntityLoadCriteria{LoadAssocsFromThis: true},
		serdes.Entity,
	)
	if err != nil {
		return nil, errors.Wrapf(err, "load cellular gateway for gateway %s", gatewayID)
	}
	apnsByName, apnResourcesByAPN, err := loadAPNs(lteGateway)
	if err != nil {
		return nil, err
	}

	subProtos, nextToken, err := loadSubProtosPage(req.PageSize, req.PageToken, networkID, apnsByName, apnResourcesByAPN)
	if err != nil {
		return nil, err
	}

	digest, noUpdates := s.getDigestAndNoUpdates(req, networkID, apnsByName, apnResourcesByAPN)
	if noUpdates {
		subProtos = []*lte_protos.SubscriberData{}
	}

	listRes := &lte_protos.ListSubscribersResponse{
		Subscribers:   subProtos,
		NextPageToken: nextToken,
		Digest:        digest,
		NoUpdates:     noUpdates,
	}

	return listRes, nil
}

func (s *subscriberdbServicer) getDigestAndNoUpdates(
	req *lte_protos.ListSubscribersRequest,
	networkID string,
	apnsByName map[string]*lte_models.ApnConfiguration,
	apnResourcesByAPN lte_models.ApnResources,
) (*lte_protos.SubscribersDigest, bool) {

	digest, noUpdates := req.Digest, false
	// This functionality is currently placed behind a feature flag.
	if s.flatDigestEnabled && req.PageToken == "" {
		digestString, err := getDigest(networkID, apnsByName, apnResourcesByAPN)
		// If digest generation fails, the error is swallowed to not affect the main functionality.
		if err != nil {
			glog.Errorf("Generating digest for subscribers in network of gateway %s failed", networkID)
		} else {
			digest = &lte_protos.SubscribersDigest{Md5Base64Digest: digestString}
			noUpdates = req.Digest.Md5Base64Digest == digestString
		}
	}

	return digest, noUpdates
}

// getDigest loads all subscribers registered on the current network, and returns
// a deterministic subscriber flat digest.
func getDigest(
	networkID string,
	apnsByName map[string]*lte_models.ApnConfiguration,
	apnResourcesByAPN lte_models.ApnResources,
) (string, error) {
	subProtosById := map[string]proto.Message{}
	curToken := ""
	for {
		subProtos, nextToken, err := loadSubProtosPage(0, curToken, networkID, apnsByName, apnResourcesByAPN)
		if err != nil {
			return "", err
		}
		for _, subProto := range subProtos {
			index := subProto.Sid.Id
			subProtosById[index] = subProto
		}
		if nextToken == "" {
			break
		}
		curToken = nextToken
	}
	return mproto.HashManyDeterministic(subProtosById)
}

func loadSubProtosPage(
	pageSize uint32, pageToken string, networkID string,
	apnsByName map[string]*lte_models.ApnConfiguration,
	apnResourcesByAPN lte_models.ApnResources,
) ([]*lte_protos.SubscriberData, string, error) {

	lc := configurator.EntityLoadCriteria{
		PageSize:           pageSize,
		PageToken:          pageToken,
		LoadConfig:         true,
		LoadAssocsToThis:   true,
		LoadAssocsFromThis: true,
	}

	subEnts, nextToken, err := configurator.LoadAllEntitiesOfType(
		networkID, lte.SubscriberEntityType, lc, serdes.Entity,
	)
	if err != nil {
		return nil, "", errors.Wrapf(err, "load subscribers in network of gateway %s", networkID)
	}

	subProtos := make([]*lte_protos.SubscriberData, 0, len(subEnts))
	for _, sub := range subEnts {
		subProto, err := convertSubEntsToProtos(sub, apnsByName, apnResourcesByAPN)
		if err != nil {
			return nil, "", err
		}
		subProto.NetworkId = &protos.NetworkID{Id: networkID}
		subProtos = append(subProtos, subProto)
	}

	return subProtos, nextToken, nil
}

func loadAPNs(gateway configurator.NetworkEntity) (map[string]*lte_models.ApnConfiguration, lte_models.ApnResources, error) {
	apns, _, err := configurator.LoadAllEntitiesOfType(
		gateway.NetworkID, lte.APNEntityType,
		configurator.EntityLoadCriteria{LoadConfig: true},
		serdes.Entity,
	)
	if err != nil {
		return nil, nil, err
	}
	apnsByName := map[string]*lte_models.ApnConfiguration{}
	for _, ent := range apns {
		apnsByName[ent.Key] = ent.Config.(*lte_models.ApnConfiguration)
	}

	apnResources, err := lte_models.LoadAPNResources(gateway.NetworkID, gateway.Associations.Filter(lte.APNResourceEntityType).Keys())
	if err != nil {
		return nil, nil, err
	}

	return apnsByName, apnResources, nil
}

func convertSubEntsToProtos(ent configurator.NetworkEntity, apnConfigs map[string]*lte_models.ApnConfiguration, apnResources lte_models.ApnResources) (*lte_protos.SubscriberData, error) {
	subData := &lte_protos.SubscriberData{}
	t, err := lte_protos.SidProto(ent.Key)
	if err != nil {
		return nil, err
	}

	subData.Sid = t
	if ent.Config == nil {
		return subData, nil
	}

	cfg := ent.Config.(*models.SubscriberConfig)
	subData.Lte = &lte_protos.LTESubscription{
		State:    lte_protos.LTESubscription_LTESubscriptionState(lte_protos.LTESubscription_LTESubscriptionState_value[cfg.Lte.State]),
		AuthAlgo: lte_protos.LTESubscription_LTEAuthAlgo(lte_protos.LTESubscription_LTEAuthAlgo_value[cfg.Lte.AuthAlgo]),
		AuthKey:  cfg.Lte.AuthKey,
		AuthOpc:  cfg.Lte.AuthOpc,
	}

	if cfg.Lte.SubProfile != "" {
		subData.SubProfile = string(cfg.Lte.SubProfile)
	} else {
		subData.SubProfile = defaultSubProfile
	}

	for _, assoc := range ent.ParentAssociations {
		if assoc.Type == lte.BaseNameEntityType {
			subData.Lte.AssignedBaseNames = append(subData.Lte.AssignedBaseNames, assoc.Key)
		} else if assoc.Type == lte.PolicyRuleEntityType {
			subData.Lte.AssignedPolicies = append(subData.Lte.AssignedPolicies, assoc.Key)
		}
	}

	// Construct the non-3gpp profile
	non3gpp := &lte_protos.Non3GPPUserProfile{
		ApnConfig: make([]*lte_protos.APNConfiguration, 0, len(ent.Associations)),
	}
	for _, assoc := range ent.Associations {
		apnConfig, apnFound := apnConfigs[assoc.Key]
		if !apnFound {
			continue
		}
		var apnResource *lte_protos.APNConfiguration_APNResource
		if apnResourceModel, ok := apnResources[assoc.Key]; ok {
			apnResource = apnResourceModel.ToProto()
		}
		apnProto := &lte_protos.APNConfiguration{
			ServiceSelection: assoc.Key,
			Ambr: &lte_protos.AggregatedMaximumBitrate{
				MaxBandwidthUl: *(apnConfig.Ambr.MaxBandwidthUl),
				MaxBandwidthDl: *(apnConfig.Ambr.MaxBandwidthDl),
			},
			QosProfile: &lte_protos.APNConfiguration_QoSProfile{
				ClassId:                 *(apnConfig.QosProfile.ClassID),
				PriorityLevel:           *(apnConfig.QosProfile.PriorityLevel),
				PreemptionCapability:    *(apnConfig.QosProfile.PreemptionCapability),
				PreemptionVulnerability: *(apnConfig.QosProfile.PreemptionVulnerability),
			},
			Resource: apnResource,
		}
		if staticIP, found := cfg.StaticIps[assoc.Key]; found {
			apnProto.AssignedStaticIp = string(staticIP)
		}
		non3gpp.ApnConfig = append(non3gpp.ApnConfig, apnProto)
	}
	sort.Slice(non3gpp.ApnConfig, func(i, j int) bool {
		return non3gpp.ApnConfig[i].ServiceSelection < non3gpp.ApnConfig[j].ServiceSelection
	})
	subData.Non_3Gpp = non3gpp

	return subData, nil
}
