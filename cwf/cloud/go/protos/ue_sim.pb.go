// Code generated by protoc-gen-go. DO NOT EDIT.
// source: cwf/protos/ue_sim.proto

package protos

import (
	context "context"
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	wrappers "github.com/golang/protobuf/ptypes/wrappers"
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

// --------------------------------------------------------------------------
// UE config
// --------------------------------------------------------------------------
type UEConfig struct {
	// Unique identifier for the UE.
	Imsi string `protobuf:"bytes,1,opt,name=imsi,proto3" json:"imsi,omitempty"`
	// Authentication key (k).
	AuthKey []byte `protobuf:"bytes,2,opt,name=auth_key,json=authKey,proto3" json:"auth_key,omitempty"`
	// Operator configuration field (Op) signed with authentication key (k).
	AuthOpc []byte `protobuf:"bytes,3,opt,name=auth_opc,json=authOpc,proto3" json:"auth_opc,omitempty"`
	// Sequence Number (SEQ).
	Seq                  uint64   `protobuf:"varint,4,opt,name=seq,proto3" json:"seq,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *UEConfig) Reset()         { *m = UEConfig{} }
func (m *UEConfig) String() string { return proto.CompactTextString(m) }
func (*UEConfig) ProtoMessage()    {}
func (*UEConfig) Descriptor() ([]byte, []int) {
	return fileDescriptor_01bc05ea16f96cbc, []int{0}
}

func (m *UEConfig) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_UEConfig.Unmarshal(m, b)
}
func (m *UEConfig) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_UEConfig.Marshal(b, m, deterministic)
}
func (m *UEConfig) XXX_Merge(src proto.Message) {
	xxx_messageInfo_UEConfig.Merge(m, src)
}
func (m *UEConfig) XXX_Size() int {
	return xxx_messageInfo_UEConfig.Size(m)
}
func (m *UEConfig) XXX_DiscardUnknown() {
	xxx_messageInfo_UEConfig.DiscardUnknown(m)
}

var xxx_messageInfo_UEConfig proto.InternalMessageInfo

func (m *UEConfig) GetImsi() string {
	if m != nil {
		return m.Imsi
	}
	return ""
}

func (m *UEConfig) GetAuthKey() []byte {
	if m != nil {
		return m.AuthKey
	}
	return nil
}

func (m *UEConfig) GetAuthOpc() []byte {
	if m != nil {
		return m.AuthOpc
	}
	return nil
}

func (m *UEConfig) GetSeq() uint64 {
	if m != nil {
		return m.Seq
	}
	return 0
}

type AuthenticateRequest struct {
	Imsi                 string   `protobuf:"bytes,1,opt,name=imsi,proto3" json:"imsi,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *AuthenticateRequest) Reset()         { *m = AuthenticateRequest{} }
func (m *AuthenticateRequest) String() string { return proto.CompactTextString(m) }
func (*AuthenticateRequest) ProtoMessage()    {}
func (*AuthenticateRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_01bc05ea16f96cbc, []int{1}
}

func (m *AuthenticateRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_AuthenticateRequest.Unmarshal(m, b)
}
func (m *AuthenticateRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_AuthenticateRequest.Marshal(b, m, deterministic)
}
func (m *AuthenticateRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_AuthenticateRequest.Merge(m, src)
}
func (m *AuthenticateRequest) XXX_Size() int {
	return xxx_messageInfo_AuthenticateRequest.Size(m)
}
func (m *AuthenticateRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_AuthenticateRequest.DiscardUnknown(m)
}

var xxx_messageInfo_AuthenticateRequest proto.InternalMessageInfo

func (m *AuthenticateRequest) GetImsi() string {
	if m != nil {
		return m.Imsi
	}
	return ""
}

type AuthenticateResponse struct {
	RadiusPacket         []byte   `protobuf:"bytes,1,opt,name=radiusPacket,proto3" json:"radiusPacket,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *AuthenticateResponse) Reset()         { *m = AuthenticateResponse{} }
func (m *AuthenticateResponse) String() string { return proto.CompactTextString(m) }
func (*AuthenticateResponse) ProtoMessage()    {}
func (*AuthenticateResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_01bc05ea16f96cbc, []int{2}
}

func (m *AuthenticateResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_AuthenticateResponse.Unmarshal(m, b)
}
func (m *AuthenticateResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_AuthenticateResponse.Marshal(b, m, deterministic)
}
func (m *AuthenticateResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_AuthenticateResponse.Merge(m, src)
}
func (m *AuthenticateResponse) XXX_Size() int {
	return xxx_messageInfo_AuthenticateResponse.Size(m)
}
func (m *AuthenticateResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_AuthenticateResponse.DiscardUnknown(m)
}

var xxx_messageInfo_AuthenticateResponse proto.InternalMessageInfo

func (m *AuthenticateResponse) GetRadiusPacket() []byte {
	if m != nil {
		return m.RadiusPacket
	}
	return nil
}

type DisconnectRequest struct {
	Imsi                 string   `protobuf:"bytes,1,opt,name=imsi,proto3" json:"imsi,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *DisconnectRequest) Reset()         { *m = DisconnectRequest{} }
func (m *DisconnectRequest) String() string { return proto.CompactTextString(m) }
func (*DisconnectRequest) ProtoMessage()    {}
func (*DisconnectRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_01bc05ea16f96cbc, []int{3}
}

func (m *DisconnectRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_DisconnectRequest.Unmarshal(m, b)
}
func (m *DisconnectRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_DisconnectRequest.Marshal(b, m, deterministic)
}
func (m *DisconnectRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_DisconnectRequest.Merge(m, src)
}
func (m *DisconnectRequest) XXX_Size() int {
	return xxx_messageInfo_DisconnectRequest.Size(m)
}
func (m *DisconnectRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_DisconnectRequest.DiscardUnknown(m)
}

var xxx_messageInfo_DisconnectRequest proto.InternalMessageInfo

func (m *DisconnectRequest) GetImsi() string {
	if m != nil {
		return m.Imsi
	}
	return ""
}

type DisconnectResponse struct {
	RadiusPacket         []byte   `protobuf:"bytes,1,opt,name=radiusPacket,proto3" json:"radiusPacket,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *DisconnectResponse) Reset()         { *m = DisconnectResponse{} }
func (m *DisconnectResponse) String() string { return proto.CompactTextString(m) }
func (*DisconnectResponse) ProtoMessage()    {}
func (*DisconnectResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_01bc05ea16f96cbc, []int{4}
}

func (m *DisconnectResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_DisconnectResponse.Unmarshal(m, b)
}
func (m *DisconnectResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_DisconnectResponse.Marshal(b, m, deterministic)
}
func (m *DisconnectResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_DisconnectResponse.Merge(m, src)
}
func (m *DisconnectResponse) XXX_Size() int {
	return xxx_messageInfo_DisconnectResponse.Size(m)
}
func (m *DisconnectResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_DisconnectResponse.DiscardUnknown(m)
}

var xxx_messageInfo_DisconnectResponse proto.InternalMessageInfo

func (m *DisconnectResponse) GetRadiusPacket() []byte {
	if m != nil {
		return m.RadiusPacket
	}
	return nil
}

type GenTrafficRequest struct {
	Imsi                 string                `protobuf:"bytes,1,opt,name=imsi,proto3" json:"imsi,omitempty"`
	Volume               *wrappers.StringValue `protobuf:"bytes,2,opt,name=volume,proto3" json:"volume,omitempty"`
	XXX_NoUnkeyedLiteral struct{}              `json:"-"`
	XXX_unrecognized     []byte                `json:"-"`
	XXX_sizecache        int32                 `json:"-"`
}

func (m *GenTrafficRequest) Reset()         { *m = GenTrafficRequest{} }
func (m *GenTrafficRequest) String() string { return proto.CompactTextString(m) }
func (*GenTrafficRequest) ProtoMessage()    {}
func (*GenTrafficRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_01bc05ea16f96cbc, []int{5}
}

func (m *GenTrafficRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GenTrafficRequest.Unmarshal(m, b)
}
func (m *GenTrafficRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GenTrafficRequest.Marshal(b, m, deterministic)
}
func (m *GenTrafficRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GenTrafficRequest.Merge(m, src)
}
func (m *GenTrafficRequest) XXX_Size() int {
	return xxx_messageInfo_GenTrafficRequest.Size(m)
}
func (m *GenTrafficRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_GenTrafficRequest.DiscardUnknown(m)
}

var xxx_messageInfo_GenTrafficRequest proto.InternalMessageInfo

func (m *GenTrafficRequest) GetImsi() string {
	if m != nil {
		return m.Imsi
	}
	return ""
}

func (m *GenTrafficRequest) GetVolume() *wrappers.StringValue {
	if m != nil {
		return m.Volume
	}
	return nil
}

func init() {
	proto.RegisterType((*UEConfig)(nil), "magma.cwf.UEConfig")
	proto.RegisterType((*AuthenticateRequest)(nil), "magma.cwf.AuthenticateRequest")
	proto.RegisterType((*AuthenticateResponse)(nil), "magma.cwf.AuthenticateResponse")
	proto.RegisterType((*DisconnectRequest)(nil), "magma.cwf.DisconnectRequest")
	proto.RegisterType((*DisconnectResponse)(nil), "magma.cwf.DisconnectResponse")
	proto.RegisterType((*GenTrafficRequest)(nil), "magma.cwf.GenTrafficRequest")
}

func init() { proto.RegisterFile("cwf/protos/ue_sim.proto", fileDescriptor_01bc05ea16f96cbc) }

var fileDescriptor_01bc05ea16f96cbc = []byte{
	// 415 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x94, 0x93, 0x4f, 0x6f, 0xd3, 0x40,
	0x10, 0xc5, 0x9b, 0x36, 0x2d, 0xed, 0xe0, 0x03, 0xd9, 0x22, 0xe1, 0x98, 0x52, 0x22, 0x5f, 0x08,
	0x97, 0xb5, 0x28, 0x1c, 0x2a, 0x2e, 0xa8, 0x40, 0xc4, 0xa1, 0x07, 0xc0, 0x25, 0x3d, 0x20, 0xa1,
	0x6a, 0xbb, 0x1e, 0xbb, 0xab, 0x7a, 0x77, 0xdd, 0xfd, 0x43, 0xd4, 0xcf, 0xc2, 0x97, 0x45, 0xb1,
	0xf3, 0xc7, 0x51, 0xd3, 0x4a, 0x3d, 0x79, 0x77, 0xde, 0x9b, 0x19, 0xf9, 0xf7, 0x6c, 0x78, 0xc1,
	0x27, 0x79, 0x52, 0x19, 0xed, 0xb4, 0x4d, 0x3c, 0x5e, 0x58, 0x21, 0x69, 0x7d, 0x23, 0x7b, 0x92,
	0x15, 0x92, 0x51, 0x3e, 0xc9, 0xa3, 0xbe, 0x36, 0xfc, 0xd8, 0xcc, 0x5d, 0x5c, 0x4b, 0xa9, 0x55,
	0xe3, 0x8a, 0x0e, 0x0b, 0xad, 0x8b, 0x12, 0x1b, 0xed, 0xd2, 0xe7, 0xc9, 0xc4, 0xb0, 0xaa, 0x42,
	0x63, 0x1b, 0x3d, 0xce, 0x61, 0x77, 0x3c, 0xfa, 0xa2, 0x55, 0x2e, 0x0a, 0x42, 0xa0, 0x2b, 0xa4,
	0x15, 0x61, 0x67, 0xd0, 0x19, 0xee, 0xa5, 0xf5, 0x99, 0xf4, 0x61, 0x97, 0x79, 0x77, 0x75, 0x71,
	0x8d, 0xb7, 0xe1, 0xe6, 0xa0, 0x33, 0x0c, 0xd2, 0x27, 0xd3, 0xfb, 0x29, 0xde, 0x2e, 0x24, 0x5d,
	0xf1, 0x70, 0x6b, 0x29, 0x7d, 0xaf, 0x38, 0x79, 0x06, 0x5b, 0x16, 0x6f, 0xc2, 0xee, 0xa0, 0x33,
	0xec, 0xa6, 0xd3, 0x63, 0xfc, 0x16, 0xf6, 0x4f, 0xbc, 0xbb, 0x42, 0xe5, 0x04, 0x67, 0x0e, 0x53,
	0xbc, 0xf1, 0x68, 0xdd, 0xba, 0x95, 0xf1, 0x47, 0x78, 0xbe, 0x6a, 0xb5, 0x95, 0x56, 0x16, 0x49,
	0x0c, 0x81, 0x61, 0x99, 0xf0, 0xf6, 0x07, 0xe3, 0xd7, 0xe8, 0xea, 0x9e, 0x20, 0x5d, 0xa9, 0xc5,
	0x6f, 0xa0, 0xf7, 0x55, 0x58, 0xae, 0x95, 0x42, 0xee, 0x1e, 0x5a, 0x72, 0x0c, 0xa4, 0x6d, 0x7c,
	0xc4, 0x8a, 0x3f, 0xd0, 0xfb, 0x86, 0xea, 0x97, 0x61, 0x79, 0x2e, 0xf8, 0x03, 0x2b, 0xc8, 0x07,
	0xd8, 0xf9, 0xab, 0x4b, 0x2f, 0xb1, 0x06, 0xf7, 0xf4, 0xe8, 0x80, 0x36, 0x59, 0xd0, 0x79, 0x16,
	0xf4, 0xcc, 0x19, 0xa1, 0x8a, 0x73, 0x56, 0x7a, 0x4c, 0x67, 0xde, 0xa3, 0x7f, 0x9b, 0xb0, 0x3d,
	0x1e, 0x9d, 0x09, 0x49, 0xde, 0xc1, 0xf6, 0x49, 0x96, 0x8d, 0x47, 0x64, 0x9f, 0x2e, 0xa2, 0xa6,
	0xf3, 0xb0, 0xa2, 0xde, 0xac, 0x58, 0x47, 0x4f, 0xcf, 0xb5, 0xc8, 0xe2, 0x0d, 0xf2, 0x13, 0x82,
	0x36, 0x3a, 0x72, 0xd8, 0xea, 0x5c, 0x83, 0x3f, 0x7a, 0x7d, 0xaf, 0xde, 0x00, 0x89, 0x37, 0xc8,
	0x29, 0xc0, 0x12, 0x14, 0x39, 0x68, 0x35, 0xdc, 0x01, 0x1d, 0xbd, 0xba, 0x47, 0x5d, 0x0c, 0xfb,
	0x04, 0xb0, 0x64, 0xb7, 0x32, 0xec, 0x0e, 0xd2, 0xb5, 0x2f, 0xf8, 0xf9, 0xe5, 0xef, 0x7e, 0x5d,
	0x4d, 0xa6, 0x7f, 0x05, 0x2f, 0xb5, 0xcf, 0x92, 0x42, 0xcf, 0x3e, 0xfc, 0xcb, 0x9d, 0xfa, 0xf9,
	0xfe, 0x7f, 0x00, 0x00, 0x00, 0xff, 0xff, 0xb0, 0x33, 0x09, 0xa2, 0x33, 0x03, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConnInterface

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion6

// UESimClient is the client API for UESim service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type UESimClient interface {
	// Adds a new UE to the store.
	//
	AddUE(ctx context.Context, in *UEConfig, opts ...grpc.CallOption) (*protos.Void, error)
	// Triggers an authentication for the UE with the specified imsi.
	//
	Authenticate(ctx context.Context, in *AuthenticateRequest, opts ...grpc.CallOption) (*AuthenticateResponse, error)
	Disconnect(ctx context.Context, in *DisconnectRequest, opts ...grpc.CallOption) (*DisconnectResponse, error)
	// Triggers iperf traffic towards the CWAG
	GenTraffic(ctx context.Context, in *GenTrafficRequest, opts ...grpc.CallOption) (*protos.Void, error)
}

type uESimClient struct {
	cc grpc.ClientConnInterface
}

func NewUESimClient(cc grpc.ClientConnInterface) UESimClient {
	return &uESimClient{cc}
}

func (c *uESimClient) AddUE(ctx context.Context, in *UEConfig, opts ...grpc.CallOption) (*protos.Void, error) {
	out := new(protos.Void)
	err := c.cc.Invoke(ctx, "/magma.cwf.UESim/AddUE", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *uESimClient) Authenticate(ctx context.Context, in *AuthenticateRequest, opts ...grpc.CallOption) (*AuthenticateResponse, error) {
	out := new(AuthenticateResponse)
	err := c.cc.Invoke(ctx, "/magma.cwf.UESim/Authenticate", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *uESimClient) Disconnect(ctx context.Context, in *DisconnectRequest, opts ...grpc.CallOption) (*DisconnectResponse, error) {
	out := new(DisconnectResponse)
	err := c.cc.Invoke(ctx, "/magma.cwf.UESim/Disconnect", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *uESimClient) GenTraffic(ctx context.Context, in *GenTrafficRequest, opts ...grpc.CallOption) (*protos.Void, error) {
	out := new(protos.Void)
	err := c.cc.Invoke(ctx, "/magma.cwf.UESim/GenTraffic", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// UESimServer is the server API for UESim service.
type UESimServer interface {
	// Adds a new UE to the store.
	//
	AddUE(context.Context, *UEConfig) (*protos.Void, error)
	// Triggers an authentication for the UE with the specified imsi.
	//
	Authenticate(context.Context, *AuthenticateRequest) (*AuthenticateResponse, error)
	Disconnect(context.Context, *DisconnectRequest) (*DisconnectResponse, error)
	// Triggers iperf traffic towards the CWAG
	GenTraffic(context.Context, *GenTrafficRequest) (*protos.Void, error)
}

// UnimplementedUESimServer can be embedded to have forward compatible implementations.
type UnimplementedUESimServer struct {
}

func (*UnimplementedUESimServer) AddUE(ctx context.Context, req *UEConfig) (*protos.Void, error) {
	return nil, status.Errorf(codes.Unimplemented, "method AddUE not implemented")
}
func (*UnimplementedUESimServer) Authenticate(ctx context.Context, req *AuthenticateRequest) (*AuthenticateResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Authenticate not implemented")
}
func (*UnimplementedUESimServer) Disconnect(ctx context.Context, req *DisconnectRequest) (*DisconnectResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Disconnect not implemented")
}
func (*UnimplementedUESimServer) GenTraffic(ctx context.Context, req *GenTrafficRequest) (*protos.Void, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GenTraffic not implemented")
}

func RegisterUESimServer(s *grpc.Server, srv UESimServer) {
	s.RegisterService(&_UESim_serviceDesc, srv)
}

func _UESim_AddUE_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UEConfig)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UESimServer).AddUE(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/magma.cwf.UESim/AddUE",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UESimServer).AddUE(ctx, req.(*UEConfig))
	}
	return interceptor(ctx, in, info, handler)
}

func _UESim_Authenticate_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(AuthenticateRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UESimServer).Authenticate(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/magma.cwf.UESim/Authenticate",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UESimServer).Authenticate(ctx, req.(*AuthenticateRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _UESim_Disconnect_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DisconnectRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UESimServer).Disconnect(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/magma.cwf.UESim/Disconnect",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UESimServer).Disconnect(ctx, req.(*DisconnectRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _UESim_GenTraffic_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GenTrafficRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UESimServer).GenTraffic(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/magma.cwf.UESim/GenTraffic",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UESimServer).GenTraffic(ctx, req.(*GenTrafficRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _UESim_serviceDesc = grpc.ServiceDesc{
	ServiceName: "magma.cwf.UESim",
	HandlerType: (*UESimServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "AddUE",
			Handler:    _UESim_AddUE_Handler,
		},
		{
			MethodName: "Authenticate",
			Handler:    _UESim_Authenticate_Handler,
		},
		{
			MethodName: "Disconnect",
			Handler:    _UESim_Disconnect_Handler,
		},
		{
			MethodName: "GenTraffic",
			Handler:    _UESim_GenTraffic_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "cwf/protos/ue_sim.proto",
}
