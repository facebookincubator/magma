package control_plane

import (
	"context"
	"fmt"
	"net"
	"net/http"

	"sync"
	"sync/atomic"
	"time"

	v2 "github.com/envoyproxy/go-control-plane/envoy/api/v2"
	core "github.com/envoyproxy/go-control-plane/envoy/api/v2/core"

	listener "github.com/envoyproxy/go-control-plane/envoy/api/v2/listener"
	v2route "github.com/envoyproxy/go-control-plane/envoy/api/v2/route"

	"github.com/envoyproxy/go-control-plane/pkg/cache"
	xds "github.com/envoyproxy/go-control-plane/pkg/server"

	"github.com/golang/glog"
	"google.golang.org/grpc"

	discovery "github.com/envoyproxy/go-control-plane/envoy/service/discovery/v2"

	hcm "github.com/envoyproxy/go-control-plane/envoy/config/filter/network/http_connection_manager/v2"
	orig_src "github.com/envoyproxy/go-control-plane/envoy/extensions/filters/listener/original_src/v3"

	"github.com/envoyproxy/go-control-plane/pkg/wellknown"
	"github.com/golang/protobuf/ptypes"
	"github.com/golang/protobuf/ptypes/wrappers"
	"magma/feg/cloud/go/protos"
)

type EnvoyController interface {
	UpdateSnapshot([]*protos.AddUEHeaderEnrichmentRequest)
}

type ControllerClient struct {
	version int32
	config  cache.SnapshotCache
}

const (
	XdsCluster  = "xds_cluster"
	Ads         = "ads"
	Xds         = "xds"
	Rest        = "rest"
	debug       = true
	port        = 18000
	gatewayPort = 18001
	mode        = Ads
	any_addr = "0.0.0.0"
	maxConcurrentStreams = 16
	initialStreamWindowSize = 65536  // 64Kib
	initialConnectionWindowSize = 1048576  // 1 MiB
	connectTimeout = 6 * time.Second
	idleTimeout = 3600 * time.Second
)

type logger struct{}

func (logger logger) Infof(format string, args ...interface{}) {
	glog.Infof(format, args...)
}
func (logger logger) Errorf(format string, args ...interface{}) {
	glog.Errorf(format, args...)
}
func (cb *callbacks) Report() {
	cb.mu.Lock()
	defer cb.mu.Unlock()
	glog.Infof("cb.Report() fetches %d,  callbacks %d", cb.fetches, cb.requests)
}
func (cb *callbacks) OnStreamOpen(ctx context.Context, id int64, typ string) error {
	glog.Infof("OnStreamOpen %d open for %s", id, typ)
	return nil
}
func (cb *callbacks) OnStreamClosed(id int64) {
	glog.Infof("OnStreamClosed %d closed", id)
}
func (cb *callbacks) OnStreamRequest(int64, *v2.DiscoveryRequest) error {
	glog.Infof("OnStreamRequest")
	cb.mu.Lock()
	defer cb.mu.Unlock()
	cb.requests++
	if cb.signal != nil {
		close(cb.signal)
		cb.signal = nil
	}
	return nil
}
func (cb *callbacks) OnStreamResponse(int64, *v2.DiscoveryRequest, *v2.DiscoveryResponse) {
	glog.Infof("OnStreamResponse...")
	cb.Report()
}
func (cb *callbacks) OnFetchRequest(ctx context.Context, req *v2.DiscoveryRequest) error {
	glog.Infof("OnFetchRequest...")
	cb.mu.Lock()
	defer cb.mu.Unlock()
	cb.fetches++
	if cb.signal != nil {
		close(cb.signal)
		cb.signal = nil
	}
	return nil
}
func (cb *callbacks) OnFetchResponse(*v2.DiscoveryRequest, *v2.DiscoveryResponse) {}

type callbacks struct {
	signal   chan struct{}
	fetches  int
	requests int
	mu       sync.Mutex
}

// Hasher returns node ID as an ID
type Hasher struct {
}

// ID function
func (h Hasher) ID(node *core.Node) string {
	if node == nil {
		return "unknown"
	}
	return node.Id
}

const grpcMaxConcurrentStreams = 1000000

// RunManagementServer starts an xDS server at the given port.
func RunManagementServer(ctx context.Context, server xds.Server, port uint) {
	var grpcOptions []grpc.ServerOption
	grpcOptions = append(grpcOptions, grpc.MaxConcurrentStreams(grpcMaxConcurrentStreams))
	grpcServer := grpc.NewServer(grpcOptions...)

	lis, err := net.Listen("tcp", fmt.Sprintf(":%d", port))
	if err != nil {
		glog.Fatalf("failed to listen %s", err)
	}

	// register services
	discovery.RegisterAggregatedDiscoveryServiceServer(grpcServer, server)
	v2.RegisterEndpointDiscoveryServiceServer(grpcServer, server)
	v2.RegisterClusterDiscoveryServiceServer(grpcServer, server)
	v2.RegisterRouteDiscoveryServiceServer(grpcServer, server)
	v2.RegisterListenerDiscoveryServiceServer(grpcServer, server)

	glog.Infof("Management server listening on port %d", port)
	go func() {
		if err = grpcServer.Serve(lis); err != nil {
			glog.Error(err)
		}
	}()
	<-ctx.Done()

	grpcServer.GracefulStop()
}

// RunManagementGateway starts an HTTP gateway to an xDS server.
func RunManagementGateway(ctx context.Context, srv xds.Server, port uint) {
	glog.Infof("Gateway listening HTTP/1.1 on port %d", port)
	server := &http.Server{Addr: fmt.Sprintf(":%d", port), Handler: &xds.HTTPGateway{Server: srv}}
	go func() {
		if err := server.ListenAndServe(); err != nil {
			glog.Error(err)
		}
	}()
}

func GetControllerClient() *ControllerClient {
	cli := ControllerClient{}
	ctx := context.Background()

	glog.Infof("Starting Envoy control plane")

	signal := make(chan struct{})
	cb := &callbacks{
		signal:   signal,
		fetches:  0,
		requests: 0,
	}
	cli.config = cache.NewSnapshotCache(mode == Ads, Hasher{}, nil)

	srv := xds.NewServer(ctx, cli.config, cb)

	// start the xDS server
	go RunManagementServer(ctx, srv, port)
	go RunManagementGateway(ctx, srv, gatewayPort)

	<-signal

	cb.Report()

	return &cli
}

func (cli *ControllerClient) UpdateSnapshot(ues []*protos.AddUEHeaderEnrichmentRequest) {
	nodeId := cli.config.GetStatusKeys()[0]

	var clusterName = "cluster1"
	cluster := []cache.Resource{
		&v2.Cluster{
			Name:                 clusterName,
			ClusterDiscoveryType: &v2.Cluster_Type{Type: v2.Cluster_ORIGINAL_DST},
			ConnectTimeout:       ptypes.DurationProto(connectTimeout),
			LbPolicy:             v2.Cluster_CLUSTER_PROVIDED,
		},
	}

	filterChains := []*listener.FilterChain{}
	var listenerName = "default_http"
	var targetPrefix = "/"
	var virtualHostName = "local_service"
	var routeConfigName = "local_route"

	glog.Errorf("Updating snapshot content")

	for _, req := range ues {
		var ue_ip_addr = string(req.UeIp.Address)
		requestHeadersToAdd := []*core.HeaderValueOption{}

		glog.Infof("Adding UE " + ue_ip_addr)

		for _, header := range req.Headers {
			requestHeadersToAdd = append(requestHeadersToAdd, &core.HeaderValueOption{
				Header: &core.HeaderValue{
					Key:   header.Name,
					Value: header.Value,
				},
			})
		}

		virtualHosts := []*v2route.VirtualHost{}
		virtualHosts = append(virtualHosts, &v2route.VirtualHost{
			Name:                virtualHostName,
			Domains:             req.Websites,
			RequestHeadersToAdd: requestHeadersToAdd,
			Routes: []*v2route.Route{{
				Match: &v2route.RouteMatch{
					PathSpecifier: &v2route.RouteMatch_Prefix{
						Prefix: targetPrefix,
					},
				},
				Action: &v2route.Route_Route{
					Route: &v2route.RouteAction{
						ClusterSpecifier: &v2route.RouteAction_Cluster{
							Cluster: clusterName,
						},
					},
				},
			}}})

		httpManager := &hcm.HttpConnectionManager{
			CodecType:        hcm.HttpConnectionManager_AUTO,
			StatPrefix:       "ingress_http",
			UseRemoteAddress: &wrappers.BoolValue{Value: true},
			CommonHttpProtocolOptions: &core.HttpProtocolOptions{
				IdleTimeout: ptypes.DurationProto(idleTimeout),
				//HeadersWithUnderscoresAction: core.HttpProtocolOptions_REJECT_REQUEST,
			},
			Http2ProtocolOptions: &core.Http2ProtocolOptions{
				MaxConcurrentStreams:        &wrappers.UInt32Value{Value: maxConcurrentStreams},
				InitialStreamWindowSize:     &wrappers.UInt32Value{Value: initialStreamWindowSize},
				InitialConnectionWindowSize: &wrappers.UInt32Value{Value: initialConnectionWindowSize},
			},
			RouteSpecifier: &hcm.HttpConnectionManager_RouteConfig{
				RouteConfig: &v2.RouteConfiguration{
					Name:         routeConfigName,
					VirtualHosts: virtualHosts,
				},
			},
			HttpFilters: []*hcm.HttpFilter{{
				Name: wellknown.Router,
			}},
		}

		pbst, err := ptypes.MarshalAny(httpManager)
		if err != nil {
			panic(err)
		}

		filterChains = append(filterChains, &listener.FilterChain{
			FilterChainMatch: &listener.FilterChainMatch{
				SourcePrefixRanges: []*core.CidrRange{{
					AddressPrefix: ue_ip_addr,
					PrefixLen:     &wrappers.UInt32Value{Value: 32},
				}},
			},
			Filters: []*listener.Filter{{
				Name: wellknown.HTTPConnectionManager,
				ConfigType: &listener.Filter_TypedConfig{
					TypedConfig: pbst,
				},
			}},
		})

	}

	o_src := &orig_src.OriginalSrc{}
	mo_src, err := ptypes.MarshalAny(o_src)
	if err != nil {
		panic(err)
	}

	glog.Infof("Creating listener " + listenerName)
	var listener = []cache.Resource{
		&v2.Listener{
			Name:        listenerName,
			Transparent: &wrappers.BoolValue{Value: true},
			Address: &core.Address{
				Address: &core.Address_SocketAddress{
					SocketAddress: &core.SocketAddress{
						Address: any_addr,
						PortSpecifier: &core.SocketAddress_PortValue{
							PortValue: 80,
						},
					},
				},
			},
			FilterChains: filterChains,
			ListenerFilters: []*listener.ListenerFilter{
				{
					Name: "envoy.filters.listener.original_dst",
				},
				{
					Name: "envoy.filters.listener.original_src",
					ConfigType: &listener.ListenerFilter_TypedConfig{
						TypedConfig: mo_src,
					},
				},
			},
		}}

	// Save snapshot
	atomic.AddInt32(&cli.version, 1)
	glog.Infof("Saved snapshot version " + fmt.Sprint(cli.version))
	snap := cache.NewSnapshot(fmt.Sprint(cli.version), nil, cluster, nil, listener, nil)
	cli.config.SetSnapshot(nodeId, snap)
}
