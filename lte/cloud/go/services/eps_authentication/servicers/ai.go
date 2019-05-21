/*
Copyright (c) Facebook, Inc. and its affiliates.
All rights reserved.

This source code is licensed under the BSD-style license found in the
LICENSE file in the root directory of this source tree.
*/

package servicers

import (
	"errors"
	"fmt"

	fegprotos "magma/feg/cloud/go/protos"
	lteprotos "magma/lte/cloud/go/protos"
	"magma/lte/cloud/go/services/eps_authentication/crypto"
	"magma/lte/cloud/go/services/eps_authentication/metrics"
	"magma/orc8r/cloud/go/identity"

	"github.com/golang/glog"
	"golang.org/x/net/context"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (srv *EPSAuthServer) AuthenticationInformation(ctx context.Context, air *fegprotos.AuthenticationInformationRequest) (*fegprotos.AuthenticationInformationAnswer, error) {
	glog.V(2).Infof("received AIR from: %s", air.GetUserName())
	metrics.AIRequests.Inc()
	if err := validateAIR(air); err != nil {
		glog.V(2).Infof("AIR is invalid: %v", err.Error())
		metrics.InvalidRequests.Inc()
		return nil, status.Errorf(codes.InvalidArgument, err.Error())
	}

	networkID, err := identity.GetClientNetworkID(ctx)
	if err != nil {
		glog.V(2).Infof("could not lookup networkID: %v", err.Error())
		metrics.NetworkIDErrors.Inc()
		return nil, err
	}
	config, err := getConfig(networkID)
	if err != nil {
		glog.V(2).Infof("could not lookup config for networkID '%s': %v", networkID, err.Error())
		metrics.ConfigErrors.Inc()
		return nil, err
	}
	subscriber, errorCode, err := srv.lookupSubscriber(air.UserName, networkID)
	if err != nil {
		glog.V(2).Infof("failed to lookup subscriber '%s': %v", air.UserName, err.Error())
		metrics.UnknownSubscribers.Inc()
		return &fegprotos.AuthenticationInformationAnswer{ErrorCode: errorCode}, err
	}

	if subscriber.State == nil {
		glog.V(2).Infof("subscriber state was nil, setting to a default state of 0")
		subscriber.State = &lteprotos.SubscriberState{}
	}

	lteAuthNextSeq, err := ResyncLteAuthSeq(subscriber, air.ResyncInfo, config.LteAuthOp)
	if err != nil {
		glog.V(2).Infof("resync auth request failed: %v", err.Error())
		metrics.ResyncAuthErrors.Inc()
		return convertAuthErrorToAuthenticationAnswer(err)
	}
	if err = srv.setLteAuthNextSeq(subscriber, lteAuthNextSeq); err != nil {
		glog.V(2).Infof("failed to store sequence number after resync: %v", err.Error())
		metrics.StorageErrors.Inc()
		return &fegprotos.AuthenticationInformationAnswer{ErrorCode: fegprotos.ErrorCode_AUTHENTICATION_DATA_UNAVAILABLE}, err
	}

	milenage, err := crypto.NewMilenageCipher(config.LteAuthAmf)
	if err != nil {
		glog.V(2).Infof("could not create milenage cipher: %v", err.Error())
		metrics.AuthErrors.Inc()
		return &fegprotos.AuthenticationInformationAnswer{ErrorCode: fegprotos.ErrorCode_AUTHORIZATION_REJECTED},
			status.Errorf(codes.FailedPrecondition, "Could not create milenage cipher: %s", err.Error())
	}

	vectors, lteAuthNextSeq, err := GenerateLteAuthVectors(
		air.NumRequestedEutranVectors,
		milenage,
		subscriber,
		air.VisitedPlmn,
		config.LteAuthOp,
		0,
	)
	if err != nil {
		glog.V(2).Infof("could not generate lte auth vectors: %v", err.Error())
		metrics.AuthErrors.Inc()
		return convertAuthErrorToAuthenticationAnswer(err)
	}
	if err = srv.setLteAuthNextSeq(subscriber, lteAuthNextSeq); err != nil {
		glog.V(2).Infof("failed to store sequence number after generating auth vectors: %v", err.Error())
		metrics.StorageErrors.Inc()
		return &fegprotos.AuthenticationInformationAnswer{ErrorCode: fegprotos.ErrorCode_AUTHENTICATION_DATA_UNAVAILABLE}, err
	}

	return &fegprotos.AuthenticationInformationAnswer{
		ErrorCode:     fegprotos.ErrorCode_SUCCESS,
		EutranVectors: convertEutranVectorsToProto(vectors),
	}, nil
}

// validateAIR returns an error iff the AIR is invalid.
func validateAIR(air *fegprotos.AuthenticationInformationRequest) error {
	if air == nil {
		return errors.New("received a nil AuthenticationInformationRequest")
	}
	if len(air.UserName) == 0 {
		return errors.New("user name was empty")
	}
	if len(air.VisitedPlmn) != crypto.ExpectedPlmnBytes {
		return fmt.Errorf("expected Visited PLMN to be %v bytes, but got %v bytes", crypto.ExpectedPlmnBytes, len(air.VisitedPlmn))
	}
	if air.NumRequestedEutranVectors == 0 {
		return errors.New("0 E-UTRAN vectors were requested")
	}
	return nil
}

// convertAuthErrorToAuthenticationAnswer converts an auth error to a result which can be returned by AuthenticationInformation.
func convertAuthErrorToAuthenticationAnswer(err error) (*fegprotos.AuthenticationInformationAnswer, error) {
	var errorCode fegprotos.ErrorCode
	var grpcErr error

	switch err.(type) {
	case AuthRejectedError:
		errorCode = fegprotos.ErrorCode_AUTHORIZATION_REJECTED
		grpcErr = status.Errorf(codes.Unauthenticated, err.Error())
	case AuthDataUnavailableError:
		errorCode = fegprotos.ErrorCode_AUTHENTICATION_DATA_UNAVAILABLE
		grpcErr = status.Errorf(codes.Unavailable, err.Error())
	default:
		errorCode = fegprotos.ErrorCode_UNDEFINED
		grpcErr = status.Errorf(codes.Unknown, err.Error())
	}

	answer := &fegprotos.AuthenticationInformationAnswer{ErrorCode: errorCode}
	return answer, grpcErr
}

// setLteAuthNextSeq sets the subscriber's LteAuthNextSeq field in the database.
func (srv *EPSAuthServer) setLteAuthNextSeq(subscriber *lteprotos.SubscriberData, lteAuthNextSeq uint64) error {
	if subscriber.GetState() == nil {
		return NewAuthDataUnavailableError("subscriber state was nil")
	}
	subscriber.State.LteAuthNextSeq = lteAuthNextSeq
	_, err := srv.Store.UpdateSubscriber(subscriber)
	return err
}

// convertEutranVectorsToProto serialized a list of E-UTRAN vectors to proto.
func convertEutranVectorsToProto(vectors []*crypto.EutranVector) []*fegprotos.AuthenticationInformationAnswer_EUTRANVector {
	result := make([]*fegprotos.AuthenticationInformationAnswer_EUTRANVector, len(vectors))
	for i, vector := range vectors {
		result[i] = &fegprotos.AuthenticationInformationAnswer_EUTRANVector{
			Rand:  vector.Rand[:],
			Xres:  vector.Xres[:],
			Autn:  vector.Autn[:],
			Kasme: vector.Kasme[:],
		}
	}
	return result
}
