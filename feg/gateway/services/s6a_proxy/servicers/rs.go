/*
Copyright (c) Facebook, Inc. and its affiliates.
All rights reserved.

This source code is licensed under the BSD-style license found in the
LICENSE file in the root directory of this source tree.
*/

// Package service implements S6a GRPC proxy service which sends AIR, ULR messages over diameter connection,
// waits (blocks) for diameter's AIAs, ULAs & returns their RPC representation
// It also handles CLR & RSR sends sync rpc request to gateway, then returns a CLA/RSA over diameter connection.
package servicers

import (
	"github.com/fiorix/go-diameter/diam"
	"github.com/fiorix/go-diameter/diam/avp"
	"github.com/fiorix/go-diameter/diam/datatype"
	"github.com/golang/glog"

	fegprotos "magma/feg/cloud/go/protos"
	"magma/feg/gateway/services/s6a_proxy"
	lteprotos "magma/lte/cloud/go/protos"
)

const (
	// MaxDiamRsRetries - number of retries for forwarding RSR to a gateway
	MaxDiamRsRetries = 1
)

// S6a CLR
func handleRSR(s *s6aProxy) diam.HandlerFunc {
	return func(c diam.Conn, m *diam.Message) {
		glog.V(2).Infof("handling RSR\n")
		var code lteprotos.ErrorCode //result-code
		var rsr RSR
		err := m.Unmarshal(&rsr)
		if err != nil {
			glog.Errorf("RSR Unmarshal failed for remote %s & message %s: %s", c.RemoteAddr(), m, err)
			return
		}
		var retries = MaxSyncRPCRetries
		for ; retries >= 0; retries-- {
			code, err = forwardRSRToGateway(&rsr)
			if err != nil {
				glog.Errorf("Failed to forward RSR to gateway. err: %v. Retries left: %v\n", err, retries)
			} else {
				break
			}
		}
		err = s.sendRSA(c, m, code, &rsr, MaxDiamRsRetries)
		if err != nil {
			glog.Errorf("Failed to send RSA: %s", err.Error())
		}
	}
}

func forwardRSRToGateway(rsr *RSR) (lteprotos.ErrorCode, error) {
	if rsr == nil {
		return diam.MissingAVP, nil
	}
	in := new(fegprotos.ResetRequest)

	res, err := s6a_proxy.GWS6AProxyReset(in)
	if err != nil {
		return lteprotos.ErrorCode_UNABLE_TO_DELIVER, err
	}
	return res.ErrorCode, nil
}

func (s *s6aProxy) sendRSA(c diam.Conn, m *diam.Message, code lteprotos.ErrorCode, rsr *RSR, retries uint) error {
	ans := m.Answer(uint32(code))
	// SessionID is required to be the AVP in position 1
	ans.InsertAVP(diam.NewAVP(avp.SessionID, avp.Mbit, 0, datatype.UTF8String(rsr.SessionID)))
	ans.NewAVP(avp.AuthSessionState, avp.Mbit, 0, datatype.Enumerated(rsr.AuthSessionState))
	s.addDiamOriginAVPs(ans)

	_, err := ans.WriteToWithRetry(c, retries)
	return err

}
