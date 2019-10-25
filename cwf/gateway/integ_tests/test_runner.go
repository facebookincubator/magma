/*
 * Copyright (c) Facebook, Inc. and its affiliates.
 * All rights reserved.
 *
 * This source code is licensed under the BSD-style license found in the
 * LICENSE file in the root directory of this source tree.
 */

package integ_tests

import (
	"fmt"
	"math/rand"
	"strconv"

	"fbc/lib/go/radius"
	cwfprotos "magma/cwf/cloud/go/protos"
	"magma/cwf/gateway/registry"
	"magma/cwf/gateway/services/uesim"
	fegprotos "magma/feg/cloud/go/protos"
	"magma/lte/cloud/go/crypto"
	lteprotos "magma/lte/cloud/go/protos"

	"github.com/golang/protobuf/ptypes/wrappers"
	"github.com/pkg/errors"
)

// todo make Op configurable, or export it in the UESimServer.
const (
	Op              = "\x11\x11\x11\x11\x11\x11\x11\x11\x11\x11\x11\x11\x11\x11\x11\x11"
	Secret          = "123456"
	MockHSSRemote   = "HSS_REMOTE"
	MockPCRFRemote  = "PCRF_REMOTE"
	MockOCSRemote   = "OCS_REMOTE"
	PipelinedRemote = "pipelined.local"
	CwagIP          = "192.168.70.101"
	HSSPort         = 9204
	OCSPort         = 9201
	PCRFPort        = 9202
	PipelinedPort   = 8443

	defaultMSISDN = "5100001234"
)

type TestRunner struct {
	imsis map[string]bool
}

// NewTestRunner initializes a new TestRunner by making a UESim client and
// and setting the next IMSI.
func NewTestRunner() *TestRunner {
	fmt.Println("************************* TestRunner setup")
	testRunner := &TestRunner{}

	testRunner.imsis = make(map[string]bool)
	fmt.Printf("Adding Mock HSS service at %s:%d\n", CwagIP, HSSPort)
	registry.AddService(MockHSSRemote, CwagIP, HSSPort)
	fmt.Printf("Adding Mock PCRF service at %s:%d\n", CwagIP, PCRFPort)
	registry.AddService(MockPCRFRemote, CwagIP, PCRFPort)
	fmt.Printf("Adding Mock OCS service at %s:%d\n", CwagIP, OCSPort)
	registry.AddService(MockOCSRemote, CwagIP, OCSPort)
	fmt.Printf("Adding Pipelined service at %s:%d\n", CwagIP, PipelinedPort)
	registry.AddService(PipelinedRemote, CwagIP, PipelinedPort)

	return testRunner
}

// ConfigUEs creates and adds the specified number of UEs and Subscribers
// to the UE Simulator and the HSS.
func (testRunner *TestRunner) ConfigUEs(numUEs int) ([]*cwfprotos.UEConfig, error) {
	fmt.Printf("************************* Configuring %d UE(s)\n", numUEs)
	ues := make([]*cwfprotos.UEConfig, 0)
	for i := 0; i < numUEs; i++ {
		imsi := ""
		for {
			imsi = getRandomIMSI()
			_, present := testRunner.imsis[imsi]
			if !present {
				break
			}
		}
		key, opc, err := getRandKeyOpcFromOp([]byte(Op))
		if err != nil {
			return nil, err
		}
		seq := getRandSeq()

		ue := makeUE(imsi, key, opc, seq)
		sub := makeSubscriber(imsi, key, opc, seq+1)

		err = uesim.AddUE(ue)
		if err != nil {
			return nil, errors.Wrap(err, "Error adding UE to UESimServer")
		}
		err = addSubscriberToHSS(sub)
		if err != nil {
			return nil, errors.Wrap(err, "Error adding Subscriber to HSS")
		}
		err = addSubscriberToPCRF(sub.GetSid())
		if err != nil {
			return nil, errors.Wrap(err, "Error adding Subscriber to PCRF")
		}
		err = addSubscriberToOCS(sub.GetSid())
		if err != nil {
			return nil, errors.Wrap(err, "Error adding Subscriber to OCS")
		}

		ues = append(ues, ue)
		fmt.Printf("Added UE to Simulator, HSS, PCRF, and OCS:\n"+
			"\tIMSI: %s\tKey: %x\tOpc: %x\tSeq: %d\n", imsi, key, opc, seq)
		testRunner.imsis[imsi] = true
	}
	fmt.Println("Successfully configured UE(s)")
	return ues, nil
}

// Authenticate simulates an authentication between the UE with the specified
// IMSI and the HSS, and returns the resulting Radius packet.
func (testRunner *TestRunner) Authenticate(imsi string) (*radius.Packet, error) {
	fmt.Printf("************************* Authenticating UE with IMSI: %s\n", imsi)
	res, err := uesim.Authenticate(&cwfprotos.AuthenticateRequest{Imsi: imsi})
	if err != nil {
		fmt.Println(err)
		return &radius.Packet{}, err
	}
	encoded := res.GetRadiusPacket()
	radiusP, err := radius.Parse(encoded, []byte(Secret))
	if err != nil {
		err = errors.Wrap(err, "Error while parsing encoded Radius packet")
		fmt.Println(err)
		return &radius.Packet{}, err
	}
	fmt.Printf("Finished Authenticating UE. Resulting RADIUS Packet: %d\n", radiusP)
	return radiusP, nil
}

// GenULTraffic simulates the UE sending traffic through the CWAG to the Internet
// by running an iperf3 client on the UE simulator and an iperf3 server on the
// Magma traffic server. volume, if provided, specifies the volume of data
// generated and it should be in the form of "1024K", "2048M" etc
func (testRunner *TestRunner) GenULTraffic(imsi string, volume *string) error {
	fmt.Printf("************************* Generating Traffic for UE with IMSI: %s\n", imsi)
	req := &cwfprotos.GenTrafficRequest{
		Imsi: imsi,
	}
	if volume != nil {
		req.Volume = &wrappers.StringValue{Value: *volume}
	}
	return uesim.GenTraffic(req)
}

// Remove subscribers, rules, flows, and monitors to clean up the state for
// consecutive test runs
func (testRunner *TestRunner) CleanUp() error {
	for imsi, _ := range testRunner.imsis {
		err := deleteSubscribersFromHSS(imsi)
		if err != nil {
			return err
		}
		err = deactivateSubscriberFlows(imsi)
		if err != nil {
			return err
		}
	}
	err := clearSubscribersFromPCRF()
	if err != nil {
		return err
	}
	err = clearSubscribersFromOCS()
	if err != nil {
		return err
	}

	return nil
}

// Add an enforcement rule that passes the subscriber from any ip to any ip.
func (tr *TestRunner) AddPassThroughPCRFRules(imsi string) error {
	fmt.Printf("************************* Adding Pass-through PCRF Rule for UE with IMSI: %s\n", imsi)
	rules := &fegprotos.AccountRules{
		Imsi:          imsi,
		RuleNames:     []string{},
		RuleBaseNames: []string{},
		RuleDefinitions: []*fegprotos.RuleDefinition{
			{
				ChargineRuleName: makeRuleIDFromIMSI(imsi),
				Precedence:       100,
				FlowDescriptions: []string{"permit out ip from any to any", "permit in ip from any to any"},
			},
		},
	}
	return addPCRFRules(rules)
}

func (tr *TestRunner) AddPCRFRules(rules *fegprotos.AccountRules) error {
	fmt.Printf("************************* Adding PCRF Rule for UE with IMSI: %s\n", rules.Imsi)
	return addPCRFRules(rules)
}

func (tr *TestRunner) AddPCRFUsageMonitors(monitorInfo *fegprotos.UsageMonitorInfo) error {
	fmt.Printf("************************* Adding PCRF Usage Monitor for UE with IMSI: %s\n", monitorInfo.Imsi)
	return addPCRFUsageMonitors(monitorInfo)
}

func (tr *TestRunner) GetPolicyUsage() (map[string]*lteprotos.RuleRecord, error) {
	recordsBySubID := map[string]*lteprotos.RuleRecord{}
	table, err := getPolicyUsage()
	if err != nil {
		return recordsBySubID, err
	}
	for _, record := range table.Records {
		recordsBySubID[record.Sid] = record
	}
	return recordsBySubID, nil
}

func makeRuleIDFromIMSI(imsi string) string {
	return "dynrule1-" + imsi
}

// getRandomIMSI makes a random 15-digit IMSI that is not added to the UESim or HSS.
func getRandomIMSI() string {
	imsi := ""
	for len(imsi) < 15 {
		imsi += strconv.Itoa(rand.Intn(10))
	}
	return imsi
}

// RandKeyOpc makes a random 16-byte key and calculates the Opc based off the Op.
func getRandKeyOpcFromOp(op []byte) (key, opc []byte, err error) {
	key = make([]byte, 16)
	rand.Read(key)

	tempOpc, err := crypto.GenerateOpc(key, op)
	if err != nil {
		return nil, nil, err
	}
	opc = tempOpc[:]
	return
}

// getRandSeq makes a random 43-bit Seq.
func getRandSeq() uint64 {
	return rand.Uint64() >> 21
}

// makeUE creates a new UE using the given values.
func makeUE(imsi string, key []byte, opc []byte, seq uint64) *cwfprotos.UEConfig {
	return &cwfprotos.UEConfig{
		Imsi:    imsi,
		AuthKey: key,
		AuthOpc: opc,
		Seq:     seq,
	}
}

// MakeSubcriber creates a new Subscriber using the given values.
func makeSubscriber(imsi string, key []byte, opc []byte, seq uint64) *lteprotos.SubscriberData {
	return &lteprotos.SubscriberData{
		Sid: &lteprotos.SubscriberID{
			Id:   imsi,
			Type: 1,
		},
		Lte: &lteprotos.LTESubscription{
			State:    1,
			AuthAlgo: 0,
			AuthKey:  key,
			AuthOpc:  opc,
		},
		State: &lteprotos.SubscriberState{
			LteAuthNextSeq: seq,
		},
		Non_3Gpp: &lteprotos.Non3GPPUserProfile{
			Msisdn:              defaultMSISDN,
			Non_3GppIpAccess:    lteprotos.Non3GPPUserProfile_NON_3GPP_SUBSCRIPTION_ALLOWED,
			Non_3GppIpAccessApn: lteprotos.Non3GPPUserProfile_NON_3GPP_APNS_ENABLE,
			ApnConfig:           &lteprotos.APNConfiguration{},
		},
	}
}
