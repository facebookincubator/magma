// Code generated by protoc-gen-go. DO NOT EDIT.
// source: southbound.proto

package protos

import (
	context "context"
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
	protos "magma/orc8r/lib/go/protos"
	math "math"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion3 // please upgrade the proto package

type GetMconfigRequest struct {
	HardwareID           string   `protobuf:"bytes,1,opt,name=hardwareID,proto3" json:"hardwareID,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *GetMconfigRequest) Reset()         { *m = GetMconfigRequest{} }
func (m *GetMconfigRequest) String() string { return proto.CompactTextString(m) }
func (*GetMconfigRequest) ProtoMessage()    {}
func (*GetMconfigRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_480661e00faacec1, []int{0}
}

func (m *GetMconfigRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetMconfigRequest.Unmarshal(m, b)
}
func (m *GetMconfigRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetMconfigRequest.Marshal(b, m, deterministic)
}
func (m *GetMconfigRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetMconfigRequest.Merge(m, src)
}
func (m *GetMconfigRequest) XXX_Size() int {
	return xxx_messageInfo_GetMconfigRequest.Size(m)
}
func (m *GetMconfigRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_GetMconfigRequest.DiscardUnknown(m)
}

var xxx_messageInfo_GetMconfigRequest proto.InternalMessageInfo

func (m *GetMconfigRequest) GetHardwareID() string {
	if m != nil {
		return m.HardwareID
	}
	return ""
}

type GetMconfigResponse struct {
	Configs   *protos.GatewayConfigs `protobuf:"bytes,1,opt,name=configs,proto3" json:"configs,omitempty"`
	LogicalID string                 `protobuf:"bytes,2,opt,name=logicalID,proto3" json:"logicalID,omitempty"`
	// TODO(8/5/20): revert json_configs once we send proto descriptors from mconfig_builders
	JsonConfigs          []byte   `protobuf:"bytes,3,opt,name=json_configs,json=jsonConfigs,proto3" json:"json_configs,omitempty"` // Deprecated: Do not use.
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *GetMconfigResponse) Reset()         { *m = GetMconfigResponse{} }
func (m *GetMconfigResponse) String() string { return proto.CompactTextString(m) }
func (*GetMconfigResponse) ProtoMessage()    {}
func (*GetMconfigResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_480661e00faacec1, []int{1}
}

func (m *GetMconfigResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetMconfigResponse.Unmarshal(m, b)
}
func (m *GetMconfigResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetMconfigResponse.Marshal(b, m, deterministic)
}
func (m *GetMconfigResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetMconfigResponse.Merge(m, src)
}
func (m *GetMconfigResponse) XXX_Size() int {
	return xxx_messageInfo_GetMconfigResponse.Size(m)
}
func (m *GetMconfigResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_GetMconfigResponse.DiscardUnknown(m)
}

var xxx_messageInfo_GetMconfigResponse proto.InternalMessageInfo

func (m *GetMconfigResponse) GetConfigs() *protos.GatewayConfigs {
	if m != nil {
		return m.Configs
	}
	return nil
}

func (m *GetMconfigResponse) GetLogicalID() string {
	if m != nil {
		return m.LogicalID
	}
	return ""
}

// Deprecated: Do not use.
func (m *GetMconfigResponse) GetJsonConfigs() []byte {
	if m != nil {
		return m.JsonConfigs
	}
	return nil
}

func init() {
	proto.RegisterType((*GetMconfigRequest)(nil), "magma.orc8r.configurator.GetMconfigRequest")
	proto.RegisterType((*GetMconfigResponse)(nil), "magma.orc8r.configurator.GetMconfigResponse")
}

func init() { proto.RegisterFile("southbound.proto", fileDescriptor_480661e00faacec1) }

var fileDescriptor_480661e00faacec1 = []byte{
	// 295 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x8c, 0x91, 0x4f, 0x4b, 0xc3, 0x40,
	0x10, 0xc5, 0xbb, 0x15, 0xd4, 0x4e, 0x7b, 0xb0, 0x7b, 0x90, 0x10, 0x45, 0x43, 0x40, 0x28, 0x28,
	0x1b, 0x68, 0x11, 0x3c, 0x79, 0x68, 0x0b, 0x25, 0x07, 0x2f, 0x11, 0x3c, 0x78, 0x91, 0x6d, 0xb2,
	0xa6, 0x91, 0x64, 0xa7, 0xdd, 0xdd, 0x50, 0xfc, 0x16, 0x7e, 0x2f, 0xbf, 0x94, 0x98, 0xb5, 0x76,
	0xc5, 0x3f, 0x78, 0x0a, 0xbc, 0xf9, 0xbd, 0xc9, 0xdb, 0x37, 0x70, 0xa0, 0xb1, 0x36, 0x8b, 0x39,
	0xd6, 0x32, 0x63, 0x4b, 0x85, 0x06, 0xa9, 0x57, 0xf1, 0xbc, 0xe2, 0x0c, 0x55, 0x7a, 0xa5, 0x58,
	0x8a, 0xf2, 0xb1, 0xc8, 0x6b, 0xc5, 0x0d, 0x2a, 0x3f, 0x68, 0x26, 0x51, 0x33, 0x89, 0x1a, 0x58,
	0x47, 0x95, 0x25, 0xac, 0xd7, 0x3f, 0xfd, 0x81, 0x48, 0xb1, 0xaa, 0x50, 0x5a, 0x20, 0x1c, 0x41,
	0x7f, 0x26, 0xcc, 0x8d, 0x35, 0x25, 0x62, 0x55, 0x0b, 0x6d, 0xe8, 0x09, 0xc0, 0x82, 0xab, 0x6c,
	0xcd, 0x95, 0x88, 0xa7, 0x1e, 0x09, 0xc8, 0xa0, 0x93, 0x38, 0x4a, 0xf8, 0x42, 0x80, 0xba, 0x2e,
	0xbd, 0x44, 0xa9, 0x05, 0xbd, 0x84, 0x3d, 0xab, 0xe8, 0xc6, 0xd3, 0x1d, 0x1e, 0x31, 0x37, 0xfa,
	0x8c, 0x1b, 0xb1, 0xe6, 0xcf, 0x13, 0x8b, 0x24, 0x1b, 0x96, 0x1e, 0x43, 0xa7, 0xc4, 0xbc, 0x48,
	0x79, 0x19, 0x4f, 0xbd, 0x76, 0xf3, 0xb3, 0xad, 0x40, 0xcf, 0xa0, 0xf7, 0xa4, 0x51, 0x3e, 0x6c,
	0x36, 0xef, 0x04, 0x64, 0xd0, 0x1b, 0xb7, 0x3d, 0x92, 0x74, 0xdf, 0xf5, 0x8f, 0x6d, 0xc3, 0x57,
	0x02, 0x87, 0xb7, 0x9f, 0xcd, 0x4d, 0x9c, 0x96, 0xe8, 0x35, 0xc0, 0x36, 0x2c, 0xed, 0x7f, 0xc9,
	0x74, 0x87, 0x45, 0xe6, 0xff, 0x15, 0x33, 0x6c, 0xd1, 0x95, 0xfb, 0xd8, 0x58, 0x1a, 0xa1, 0x24,
	0x2f, 0xe9, 0x39, 0xfb, 0xed, 0x2c, 0xec, 0x5b, 0xa1, 0xfe, 0xc5, 0xff, 0x60, 0xdb, 0x63, 0xd8,
	0x1a, 0xef, 0xdf, 0xef, 0xda, 0x63, 0xcd, 0xed, 0x77, 0xf4, 0x16, 0x00, 0x00, 0xff, 0xff, 0x07,
	0x47, 0x66, 0x4e, 0x17, 0x02, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConnInterface

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion6

// SouthboundConfiguratorClient is the client API for SouthboundConfigurator service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type SouthboundConfiguratorClient interface {
	GetMconfig(ctx context.Context, in *protos.Void, opts ...grpc.CallOption) (*protos.GatewayConfigs, error)
	// GetMconfigInternal exists to support the existing streamer mconfig
	// policy. This should be removed when we migrate gateway mconfig updates
	// from streamer to this southbound configurator servicer.
	GetMconfigInternal(ctx context.Context, in *GetMconfigRequest, opts ...grpc.CallOption) (*GetMconfigResponse, error)
}

type southboundConfiguratorClient struct {
	cc grpc.ClientConnInterface
}

func NewSouthboundConfiguratorClient(cc grpc.ClientConnInterface) SouthboundConfiguratorClient {
	return &southboundConfiguratorClient{cc}
}

func (c *southboundConfiguratorClient) GetMconfig(ctx context.Context, in *protos.Void, opts ...grpc.CallOption) (*protos.GatewayConfigs, error) {
	out := new(protos.GatewayConfigs)
	err := c.cc.Invoke(ctx, "/magma.orc8r.configurator.SouthboundConfigurator/GetMconfig", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *southboundConfiguratorClient) GetMconfigInternal(ctx context.Context, in *GetMconfigRequest, opts ...grpc.CallOption) (*GetMconfigResponse, error) {
	out := new(GetMconfigResponse)
	err := c.cc.Invoke(ctx, "/magma.orc8r.configurator.SouthboundConfigurator/GetMconfigInternal", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// SouthboundConfiguratorServer is the server API for SouthboundConfigurator service.
type SouthboundConfiguratorServer interface {
	GetMconfig(context.Context, *protos.Void) (*protos.GatewayConfigs, error)
	// GetMconfigInternal exists to support the existing streamer mconfig
	// policy. This should be removed when we migrate gateway mconfig updates
	// from streamer to this southbound configurator servicer.
	GetMconfigInternal(context.Context, *GetMconfigRequest) (*GetMconfigResponse, error)
}

// UnimplementedSouthboundConfiguratorServer can be embedded to have forward compatible implementations.
type UnimplementedSouthboundConfiguratorServer struct {
}

func (*UnimplementedSouthboundConfiguratorServer) GetMconfig(ctx context.Context, req *protos.Void) (*protos.GatewayConfigs, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetMconfig not implemented")
}
func (*UnimplementedSouthboundConfiguratorServer) GetMconfigInternal(ctx context.Context, req *GetMconfigRequest) (*GetMconfigResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetMconfigInternal not implemented")
}

func RegisterSouthboundConfiguratorServer(s *grpc.Server, srv SouthboundConfiguratorServer) {
	s.RegisterService(&_SouthboundConfigurator_serviceDesc, srv)
}

func _SouthboundConfigurator_GetMconfig_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(protos.Void)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SouthboundConfiguratorServer).GetMconfig(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/magma.orc8r.configurator.SouthboundConfigurator/GetMconfig",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SouthboundConfiguratorServer).GetMconfig(ctx, req.(*protos.Void))
	}
	return interceptor(ctx, in, info, handler)
}

func _SouthboundConfigurator_GetMconfigInternal_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetMconfigRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SouthboundConfiguratorServer).GetMconfigInternal(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/magma.orc8r.configurator.SouthboundConfigurator/GetMconfigInternal",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SouthboundConfiguratorServer).GetMconfigInternal(ctx, req.(*GetMconfigRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _SouthboundConfigurator_serviceDesc = grpc.ServiceDesc{
	ServiceName: "magma.orc8r.configurator.SouthboundConfigurator",
	HandlerType: (*SouthboundConfiguratorServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetMconfig",
			Handler:    _SouthboundConfigurator_GetMconfig_Handler,
		},
		{
			MethodName: "GetMconfigInternal",
			Handler:    _SouthboundConfigurator_GetMconfigInternal_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "southbound.proto",
}
