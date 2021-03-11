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
	"fmt"
	"net"
	"time"

	"github.com/wmnsk/go-gtp/gtpv2"
	"github.com/wmnsk/go-gtp/gtpv2/ie"
	"github.com/wmnsk/go-gtp/gtpv2/message"

	"magma/feg/cloud/go/protos"
	"magma/feg/gateway/gtp"
)

// buildCreateSessionRequestIE creates a Message with all the IE needed for a Create Session Request
func buildCreateSessionRequestMsg(cPgwUDPAddr *net.UDPAddr, req *protos.CreateSessionRequestPgw) (message.Message, error) {
	// Create session needs two FTEIDs:
	// - S8 control plane FTEID will be built using local address and control TEID
	//	 passed by MME
	// - S8 user plane FTEID, provided by MME in the requested bearer

	// TODO: look for a better way to find the local ip (avoid pinging on each request)
	// (obtain the IP that is going to send the packet first)
	ip, err := gtp.GetLocalOutboundIP(cPgwUDPAddr)
	if err != nil {
		return nil, err
	}

	// Control plane TEID
	cFegFTeid := ie.NewFullyQualifiedTEID(gtpv2.IFTypeS5S8SGWGTPC,
		req.CAgwTeid, ip.String(), "").WithInstance(0)

	// User plane TEID (ip belongs to pipelined GTP-U interface)
	uAgwFTeidReq := req.BearerContext.GetUserPlaneFteid()
	uAgwFTeid := ie.NewFullyQualifiedTEID(gtpv2.IFTypeS5S8SGWGTPU,
		uAgwFTeidReq.Teid, uAgwFTeidReq.Ipv4Address, uAgwFTeidReq.Ipv6Address).WithInstance(2)

	// Qos
	qos := req.BearerContext.GetQos()
	ieQos := ie.NewBearerQoS(uint8(qos.Pci), uint8(qos.PriorityLevel), uint8(qos.PreemptionVulnerability),
		uint8(qos.Qci), qos.Mbr.BrUl, qos.Mbr.BrDl, qos.Gbr.BrUl, qos.Gbr.BrDl)

	// bearer
	bearerId := ie.NewEPSBearerID(uint8(req.BearerContext.Id))
	bearer := ie.NewBearerContext(bearerId, uAgwFTeid, ieQos)

	//timezone
	offset := time.Duration(req.TimeZone.DeltaSeconds) * time.Second
	daylightSavingTime := uint8(req.TimeZone.DaylightSavingTime)

	ies := []*ie.IE{
		ie.NewIMSI(req.GetImsi()),
		bearer,
		cFegFTeid,
		getUserLocationIndication(req.ServingNetwork.Mcc, req.ServingNetwork.Mcc, req.Uli),
		getPdnType(req.PdnType),
		getPDNAddressAllocation(req),
		getRatType(req.RatType),
		getSelectionModeType(req.SelectionMode),
		ie.NewMSISDN(string(req.Msisdn[:])),
		ie.NewMobileEquipmentIdentity(req.Mei),
		ie.NewServingNetwork(req.ServingNetwork.Mcc, req.ServingNetwork.Mnc),
		ie.NewAccessPointName(req.Apn),
		ie.NewAggregateMaximumBitRate(uint32(req.Ambr.BrUl), uint32(req.Ambr.BrDl)),
		ie.NewUETimeZone(offset, daylightSavingTime),

		// TODO: Hardcoded values
		ie.NewIndicationFromOctets(0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00),
		ie.NewAPNRestriction(gtpv2.APNRestrictionNoExistingContextsorRestriction),
		// TODO: set charging characteristics
	}

	msg := message.NewCreateSessionRequest(0, 0, ies...)

	return msg, nil
}

func buildDeleteSessionRequestMsg(req *protos.DeleteSessionRequestPgw) message.Message {
	ies := []*ie.IE{
		ie.NewEPSBearerID(uint8(req.BearerId)),
	}
	return message.NewDeleteSessionRequest(req.CPgwFteid.Teid, 0, ies...)
}

func getPDNAddressAllocation(req *protos.CreateSessionRequestPgw) *ie.IE {
	var res *ie.IE
	if req.PdnType == protos.PDNType_IPV4 {
		res = ie.NewPDNAddressAllocation(req.Paa.Ipv4Address)
	}
	if req.PdnType == protos.PDNType_IPV6 {
		res = ie.NewPDNAddressAllocationIPv6(req.Paa.Ipv6Address, uint8(req.Paa.Ipv6Prefix))
	}
	if req.PdnType == protos.PDNType_IPV4V6 {
		res = ie.NewPDNAddressAllocationDual(req.Paa.Ipv4Address, req.Paa.Ipv6Address, uint8(req.Paa.Ipv6Prefix))
	}
	return res
}

// getPdnType convert proto PDNType into GTP PDN type
func getPdnType(pdnType protos.PDNType) *ie.IE {
	var res = uint8(0)
	switch pdnType {
	case protos.PDNType_IPV4:
		res = gtpv2.PDNTypeIPv4 // v4
	case protos.PDNType_IPV6:
		res = gtpv2.PDNTypeIPv6 // v6
	case protos.PDNType_IPV4V6:
		res = gtpv2.PDNTypeIPv4 // v4v6
	case protos.PDNType_NonIP:
		res = gtpv2.PDNTypeNonIP // nonIP
	default:
		panic(fmt.Sprintf("PdnType %d does not exist", pdnType))
	}
	return ie.NewPDNType(res)
}

func getUserLocationIndication(mcc, mnc string, uli *protos.UserLocationInformation) *ie.IE {
	var (
		cgi    *ie.CGI    = nil
		sai    *ie.SAI    = nil
		rai    *ie.RAI    = nil
		tai    *ie.TAI    = nil
		ecgi   *ie.ECGI   = nil
		lai    *ie.LAI    = nil
		menbi  *ie.MENBI  = nil
		emenbi *ie.EMENBI = nil
	)

	if uli.Lac != 0 && uli.Ci != 0 {
		cgi = ie.NewCGI(mcc, mnc, uint16(uli.Lac), uint16(uli.Ci))
	}
	if uli.Lac != 0 && uli.Sac != 0 {
		sai = ie.NewSAI(mcc, mnc, uint16(uli.Lac), uint16(uli.Sac))
	}
	if uli.Lac != 0 && uli.Rac != 0 {
		rai = ie.NewRAI(mcc, mnc, uint16(uli.Lac), uint16(uli.Rac))
	}
	if uli.Tac != 0 {
		tai = ie.NewTAI(mcc, mnc, uint16(uli.Tac))
	}
	if uli.Eci != 0 {
		ecgi = ie.NewECGI(mcc, mnc, uli.Eci)
	}
	if uli.Lac != 0 {
		lai = ie.NewLAI(mcc, mnc, uint16(uli.Lac))
	}
	if uli.MeNbi != 0 {
		menbi = ie.NewMENBI(mcc, mnc, uli.MeNbi)
	}
	if uli.EMeNbi != 0 {
		emenbi = ie.NewEMENBI(mcc, mnc, uli.EMeNbi)
	}
	return ie.NewUserLocationInformationStruct(cgi, sai, rai, tai, ecgi, lai, menbi, emenbi)
}

func getRatType(ratType protos.RATType) *ie.IE {
	var rType uint8
	switch ratType {
	case protos.RATType_RESERVED:
		rType = 0
	case protos.RATType_UTRAN:
		rType = 1
	case protos.RATType_GERAN:
		rType = 2
	case protos.RATType_WLAN:
		rType = 3
	case protos.RATType_GAN:
		rType = 4
	case protos.RATType_HSPA:
		rType = 5
	case protos.RATType_EUTRAN:
		rType = 6
	case protos.RATType_VIRTUAL:
		rType = 7
	case protos.RATType_EUTRAN_NB_IOT:
		rType = 8
	case protos.RATType_LTE_M:
		rType = 9
	case protos.RATType_NR:
		rType = 10
	default:
		panic(fmt.Sprintf("RatType %d does not exist", ratType))
	}
	return ie.NewRATType(rType)
}

func getSelectionModeType(selMode protos.SelectionModeType) *ie.IE {
	var rType uint8
	switch selMode {
	case protos.SelectionModeType_APN_provided_subscription_verified:
		rType = gtpv2.SelectionModeMSorNetworkProvidedAPNSubscribedVerified
	case protos.SelectionModeType_ms_APN_subscription_not_verified:
		rType = gtpv2.SelectionModeMSProvidedAPNSubscriptionNotVerified
	case protos.SelectionModeType_network_APN_subscription_not_verified:
		rType = gtpv2.SelectionModeNetworkProvidedAPNSubscriptionNotVerified
	default:
		panic(fmt.Sprintf("RatType %d does not exist", selMode))
	}
	return ie.NewSelectionMode(rType)
}
