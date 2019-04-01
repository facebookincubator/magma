/*
Copyright (c) Facebook, Inc. and its affiliates.
All rights reserved.

This source code is licensed under the BSD-style license found in the
LICENSE file in the root directory of this source tree.
*/

// This is a http/2 h2c server. It's run within the same process as SyncRPC grpc servicer.
// When a client wants to send a grpc request to some service on gateway with hardwareId hwId,
// it identifies the addr of the SyncRPC grpc servicer to which the gateway has a bidirectional stream to,
// and sends a grpc request to the httpServer that's within the same process of that grpc servicer.
// This httpServer converts httpRequest to GatewayRequest, send it over to grpc servicer using
// GatewayRPCBroker, waits for a response, and converts the GatewayResponse to a HttpResponse and send it back
// to the client.
package httpserver

import (
	"fmt"
	"io"
	"io/ioutil"
	"net"
	"net/http"
	"net/url"
	"strconv"
	"strings"
	"time"

	"magma/orc8r/cloud/go/http2"
	"magma/orc8r/cloud/go/protos"
	"magma/orc8r/cloud/go/services/dispatcher/broker"
	"magma/orc8r/cloud/go/services/dispatcher/gateway_registry"

	"github.com/golang/glog"
	"google.golang.org/grpc/codes"
)

const (
	DEFAULT_HTTP_RESPONSE_STATUS = 200
)

type SyncRPCHttpServer struct {
	*http2.H2CServer
	broker broker.GatewayRPCBroker
}

func NewSyncRPCHttpServer(broker broker.GatewayRPCBroker) *SyncRPCHttpServer {
	return &SyncRPCHttpServer{http2.NewH2CServer(), broker}
}

func (server *SyncRPCHttpServer) Run(addr string) {
	server.H2CServer.Run(addr, server.rootHandler)
}

func (server *SyncRPCHttpServer) Serve(listener net.Listener) {
	server.H2CServer.Serve(listener, server.rootHandler)
}

func (server *SyncRPCHttpServer) rootHandler(responseWriter http.ResponseWriter, req *http.Request) {
	http2.LogRequestWithVerbosity(req, 4)
	respChan, err := server.sendRequest(req)
	if err != nil {
		glog.Errorf(err.Msg)
		// also write to client
		http2.WriteErrResponse(responseWriter, err)
		return
	}
	// wait for response or timeout
	select {
	case gwResponse := <-respChan:
		err := processResponse(responseWriter, gwResponse)
		if err != nil {
			glog.Errorf(err.Msg)
			http2.WriteErrResponse(responseWriter, err)
		}
		return
	case <-time.After(time.Second * 10):
		http2.WriteErrResponse(responseWriter,
			http2.NewHTTPGrpcError("Request timed out",
				int(codes.DeadlineExceeded),
				http.StatusRequestTimeout))
		return
	}

}

func (server *SyncRPCHttpServer) sendRequest(req *http.Request) (chan *protos.GatewayResponse, *http2.HTTPGrpcError) {
	gwReq, err := createRequest(req)
	if err != nil {
		return nil, err
	}
	respChan, sendReqErr := server.broker.SendRequestToGateway(gwReq)
	if sendReqErr != nil {
		errMsg := fmt.Sprintf("err sending request %v to gateway: %v", gwReq, sendReqErr)
		return nil, http2.NewHTTPGrpcError(errMsg,
			int(codes.Internal), http.StatusInternalServerError)
	}
	return respChan, nil
}

func createRequest(req *http.Request) (*protos.GatewayRequest, *http2.HTTPGrpcError) {
	headers := req.Header
	gwIds := headers[gateway_registry.GATEWAYID_HEADER_KEY]
	if len(gwIds) == 0 || len(gwIds[0]) == 0 {
		str := fmt.Sprintf("No Gatewayid provided in metaData")
		return nil, http2.NewHTTPGrpcError(
			str,
			int(codes.InvalidArgument),
			http.StatusBadRequest)
	}
	gwId := gwIds[0]
	delete(headers, gateway_registry.GATEWAYID_HEADER_KEY)
	authority, err := getAuthority(req.Host)
	if err != nil {
		return nil, err
	}
	path, err := getPath(req.URL)
	if err != nil {
		return nil, err
	}
	body, err := getPayload(req.Body)
	if err != nil {
		return nil, err
	}
	gwReq := &protos.GatewayRequest{
		GwId:      gwId,
		Authority: authority,
		Path:      path,
		Headers:   convertHeadersForProto(headers),
		Payload:   body,
	}
	return gwReq, nil
}

func getAuthority(host string) (string, *http2.HTTPGrpcError) {
	if len(host) == 0 {
		errMsg := "No authority provided"
		return "", http2.NewHTTPGrpcError(errMsg, int(codes.InvalidArgument), http.StatusBadRequest)
	} else {
		return host, nil
	}
}

func getPath(url *url.URL) (string, *http2.HTTPGrpcError) {
	if url == nil || len(url.Path) == 0 {
		errMsg := "No url path provided"
		return "", http2.NewHTTPGrpcError(errMsg, int(codes.InvalidArgument),
			http.StatusBadRequest)
	}
	return url.Path, nil
}

func getPayload(body io.ReadCloser) ([]byte, *http2.HTTPGrpcError) {
	payload, err := ioutil.ReadAll(body)
	defer body.Close()
	if err != nil {
		errMsg := fmt.Sprintf("err reading req body: %v", err)
		return nil, http2.NewHTTPGrpcError(
			errMsg,
			int(codes.InvalidArgument),
			http.StatusBadRequest)
	}
	return payload, nil
}

func convertHeadersForProto(headers http.Header) map[string]string {
	ret := make(map[string]string)
	for k, vals := range headers {
		ret[k] = concatenateHeaders(vals)
	}
	return ret
}

func concatenateHeaders(headers []string) string {
	return strings.Join(headers, ",")
}

func processResponse(w http.ResponseWriter, gwResp *protos.GatewayResponse) *http2.HTTPGrpcError {
	// remain for backward compatibility, but it shouldn't get forwarded
	// a nil GatewayResponse in new versions.
	if gwResp == nil {
		errMsg := "nil GatewayResponse"
		return http2.NewHTTPGrpcError(errMsg,
			int(codes.Internal), http.StatusInternalServerError)
	}
	if gwResp.Err != "" {
		return http2.NewHTTPGrpcError(gwResp.Err, int(codes.Internal),
			http.StatusInternalServerError)
	}
	headers := gwResp.GetHeaders()
	writeHeadersToResponse(headers, w)
	httpStatus, err := getHttpStatusFromGatewayResponse(gwResp.Status)
	w.WriteHeader(httpStatus)
	if gwResp.Payload != nil {
		w.Write(gwResp.Payload)
	}
	if err != nil {
		// only log, and do not send to client
		glog.Errorf("%v\n", err)
	}
	return nil
}

func getHttpStatusFromGatewayResponse(gwRespStatus string) (int, error) {
	httpStatus, err := strconv.Atoi(gwRespStatus)
	if err != nil {
		return DEFAULT_HTTP_RESPONSE_STATUS,
			fmt.Errorf("cannot parse status of gatewayResponse: %v\n", err)
	}
	// invalid status code, defaults to 200
	if statusText := http.StatusText(httpStatus); len(statusText) == 0 {
		return DEFAULT_HTTP_RESPONSE_STATUS, fmt.Errorf("Unrecognized httpStatus: %v\n", httpStatus)
	}
	return httpStatus, nil
}

func writeHeadersToResponse(headers map[string]string, w http.ResponseWriter) {
	// see how to write trailers: https://golang.org/pkg/net/http/#example_ResponseWriter_trailers
	w.Header().Set("Trailer", "Grpc-Status")
	w.Header().Add("Trailer", "Grpc-Message")
	for k, v := range headers {
		vals := strings.Split(v, ",")
		for _, val := range vals {
			if len(w.Header().Get(k)) == 0 {
				w.Header().Set(k, val)
			} else {
				w.Header().Add(k, val)
			}
		}
	}
}
