// Code generated by protoc-gen-go. DO NOT EDIT.
// source: protos/eap_auth.proto

package protos // import "magma/feg/gateway/services/eap/protos"

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

import (
	context "golang.org/x/net/context"
	grpc "google.golang.org/grpc"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

type EapType int32

const (
	// Mandatory EAP Method types
	EapType_Reserved      EapType = 0
	EapType_Identity      EapType = 1
	EapType_Notification  EapType = 2
	EapType_Legacy_Nak    EapType = 3
	EapType_MD5_Challenge EapType = 4
	EapType_Expanded      EapType = 254
	EapType_Experimental  EapType = 255
	// EAP Method Authenticator types
	EapType_TLS      EapType = 13
	EapType_SIM      EapType = 18
	EapType_AKA      EapType = 23
	EapType_AKAPrime EapType = 50
)

var EapType_name = map[int32]string{
	0:   "Reserved",
	1:   "Identity",
	2:   "Notification",
	3:   "Legacy_Nak",
	4:   "MD5_Challenge",
	254: "Expanded",
	255: "Experimental",
	13:  "TLS",
	18:  "SIM",
	23:  "AKA",
	50:  "AKAPrime",
}
var EapType_value = map[string]int32{
	"Reserved":      0,
	"Identity":      1,
	"Notification":  2,
	"Legacy_Nak":    3,
	"MD5_Challenge": 4,
	"Expanded":      254,
	"Experimental":  255,
	"TLS":           13,
	"SIM":           18,
	"AKA":           23,
	"AKAPrime":      50,
}

func (x EapType) String() string {
	return proto.EnumName(EapType_name, int32(x))
}
func (EapType) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_eap_auth_657ecccee8c48a12, []int{0}
}

type EapCode int32

const (
	EapCode_Undefined EapCode = 0
	EapCode_Request   EapCode = 1
	EapCode_Response  EapCode = 2
	EapCode_Success   EapCode = 3
	EapCode_Failure   EapCode = 4
)

var EapCode_name = map[int32]string{
	0: "Undefined",
	1: "Request",
	2: "Response",
	3: "Success",
	4: "Failure",
}
var EapCode_value = map[string]int32{
	"Undefined": 0,
	"Request":   1,
	"Response":  2,
	"Success":   3,
	"Failure":   4,
}

func (x EapCode) String() string {
	return proto.EnumName(EapCode_name, int32(x))
}
func (EapCode) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_eap_auth_657ecccee8c48a12, []int{1}
}

type EapContext struct {
	SessionId            string   `protobuf:"bytes,1,opt,name=session_id,json=sessionId,proto3" json:"session_id,omitempty"`
	Msk                  []byte   `protobuf:"bytes,2,opt,name=msk,proto3" json:"msk,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *EapContext) Reset()         { *m = EapContext{} }
func (m *EapContext) String() string { return proto.CompactTextString(m) }
func (*EapContext) ProtoMessage()    {}
func (*EapContext) Descriptor() ([]byte, []int) {
	return fileDescriptor_eap_auth_657ecccee8c48a12, []int{0}
}
func (m *EapContext) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_EapContext.Unmarshal(m, b)
}
func (m *EapContext) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_EapContext.Marshal(b, m, deterministic)
}
func (dst *EapContext) XXX_Merge(src proto.Message) {
	xxx_messageInfo_EapContext.Merge(dst, src)
}
func (m *EapContext) XXX_Size() int {
	return xxx_messageInfo_EapContext.Size(m)
}
func (m *EapContext) XXX_DiscardUnknown() {
	xxx_messageInfo_EapContext.DiscardUnknown(m)
}

var xxx_messageInfo_EapContext proto.InternalMessageInfo

func (m *EapContext) GetSessionId() string {
	if m != nil {
		return m.SessionId
	}
	return ""
}

func (m *EapContext) GetMsk() []byte {
	if m != nil {
		return m.Msk
	}
	return nil
}

type EapMessage struct {
	Payload              []byte   `protobuf:"bytes,1,opt,name=payload,proto3" json:"payload,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *EapMessage) Reset()         { *m = EapMessage{} }
func (m *EapMessage) String() string { return proto.CompactTextString(m) }
func (*EapMessage) ProtoMessage()    {}
func (*EapMessage) Descriptor() ([]byte, []int) {
	return fileDescriptor_eap_auth_657ecccee8c48a12, []int{1}
}
func (m *EapMessage) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_EapMessage.Unmarshal(m, b)
}
func (m *EapMessage) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_EapMessage.Marshal(b, m, deterministic)
}
func (dst *EapMessage) XXX_Merge(src proto.Message) {
	xxx_messageInfo_EapMessage.Merge(dst, src)
}
func (m *EapMessage) XXX_Size() int {
	return xxx_messageInfo_EapMessage.Size(m)
}
func (m *EapMessage) XXX_DiscardUnknown() {
	xxx_messageInfo_EapMessage.DiscardUnknown(m)
}

var xxx_messageInfo_EapMessage proto.InternalMessageInfo

func (m *EapMessage) GetPayload() []byte {
	if m != nil {
		return m.Payload
	}
	return nil
}

type Eap struct {
	Payload              []byte      `protobuf:"bytes,1,opt,name=payload,proto3" json:"payload,omitempty"`
	Ctx                  *EapContext `protobuf:"bytes,2,opt,name=ctx,proto3" json:"ctx,omitempty"`
	XXX_NoUnkeyedLiteral struct{}    `json:"-"`
	XXX_unrecognized     []byte      `json:"-"`
	XXX_sizecache        int32       `json:"-"`
}

func (m *Eap) Reset()         { *m = Eap{} }
func (m *Eap) String() string { return proto.CompactTextString(m) }
func (*Eap) ProtoMessage()    {}
func (*Eap) Descriptor() ([]byte, []int) {
	return fileDescriptor_eap_auth_657ecccee8c48a12, []int{2}
}
func (m *Eap) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Eap.Unmarshal(m, b)
}
func (m *Eap) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Eap.Marshal(b, m, deterministic)
}
func (dst *Eap) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Eap.Merge(dst, src)
}
func (m *Eap) XXX_Size() int {
	return xxx_messageInfo_Eap.Size(m)
}
func (m *Eap) XXX_DiscardUnknown() {
	xxx_messageInfo_Eap.DiscardUnknown(m)
}

var xxx_messageInfo_Eap proto.InternalMessageInfo

func (m *Eap) GetPayload() []byte {
	if m != nil {
		return m.Payload
	}
	return nil
}

func (m *Eap) GetCtx() *EapContext {
	if m != nil {
		return m.Ctx
	}
	return nil
}

func init() {
	proto.RegisterType((*EapContext)(nil), "eap.eap_context")
	proto.RegisterType((*EapMessage)(nil), "eap.eap_message")
	proto.RegisterType((*Eap)(nil), "eap.eap")
	proto.RegisterEnum("eap.EapType", EapType_name, EapType_value)
	proto.RegisterEnum("eap.EapCode", EapCode_name, EapCode_value)
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// EapServiceClient is the client API for EapService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type EapServiceClient interface {
	Handle(ctx context.Context, in *Eap, opts ...grpc.CallOption) (*Eap, error)
}

type eapServiceClient struct {
	cc *grpc.ClientConn
}

func NewEapServiceClient(cc *grpc.ClientConn) EapServiceClient {
	return &eapServiceClient{cc}
}

func (c *eapServiceClient) Handle(ctx context.Context, in *Eap, opts ...grpc.CallOption) (*Eap, error) {
	out := new(Eap)
	err := c.cc.Invoke(ctx, "/eap.eap_service/handle", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// EapServiceServer is the server API for EapService service.
type EapServiceServer interface {
	Handle(context.Context, *Eap) (*Eap, error)
}

func RegisterEapServiceServer(s *grpc.Server, srv EapServiceServer) {
	s.RegisterService(&_EapService_serviceDesc, srv)
}

func _EapService_Handle_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Eap)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(EapServiceServer).Handle(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/eap.eap_service/Handle",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(EapServiceServer).Handle(ctx, req.(*Eap))
	}
	return interceptor(ctx, in, info, handler)
}

var _EapService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "eap.eap_service",
	HandlerType: (*EapServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "handle",
			Handler:    _EapService_Handle_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "protos/eap_auth.proto",
}

func init() { proto.RegisterFile("protos/eap_auth.proto", fileDescriptor_eap_auth_657ecccee8c48a12) }

var fileDescriptor_eap_auth_657ecccee8c48a12 = []byte{
	// 402 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x74, 0x52, 0x51, 0x6b, 0x13, 0x41,
	0x10, 0xee, 0xe5, 0x4a, 0x93, 0x4c, 0x12, 0x99, 0x2e, 0x88, 0x41, 0x50, 0xca, 0x81, 0xb4, 0x14,
	0xcc, 0x41, 0xc4, 0x57, 0x21, 0xd6, 0x0a, 0xa1, 0x6d, 0x90, 0x8b, 0xbe, 0xf8, 0x12, 0xc6, 0xdb,
	0xc9, 0x65, 0xe9, 0xdd, 0xee, 0x9a, 0xdd, 0x68, 0xee, 0x0f, 0xf9, 0x33, 0x55, 0xf6, 0x2e, 0x15,
	0x5f, 0xfa, 0xb4, 0xdf, 0x37, 0x33, 0xdf, 0xce, 0x37, 0xc3, 0xc0, 0x53, 0xbb, 0x35, 0xde, 0xb8,
	0x94, 0xc9, 0xae, 0x68, 0xe7, 0x37, 0x93, 0x86, 0x8b, 0x98, 0xc9, 0x26, 0xef, 0x60, 0x10, 0xc2,
	0xb9, 0xd1, 0x9e, 0xf7, 0x5e, 0xbc, 0x00, 0x70, 0xec, 0x9c, 0x32, 0x7a, 0xa5, 0xe4, 0x38, 0x3a,
	0x8b, 0x2e, 0xfa, 0x59, 0xff, 0x10, 0x99, 0x4b, 0x81, 0x10, 0x57, 0xee, 0x7e, 0xdc, 0x39, 0x8b,
	0x2e, 0x86, 0x59, 0x80, 0xc9, 0x79, 0xab, 0xaf, 0xd8, 0x39, 0x2a, 0x58, 0x8c, 0xa1, 0x6b, 0xa9,
	0x2e, 0x0d, 0xb5, 0xe2, 0x61, 0xf6, 0x40, 0x93, 0x2b, 0x08, 0xfd, 0x1e, 0x2f, 0x10, 0x09, 0xc4,
	0xb9, 0xdf, 0x37, 0x7f, 0x0f, 0xa6, 0x38, 0x61, 0xb2, 0x93, 0xff, 0x9c, 0x65, 0x21, 0x79, 0xf9,
	0x2b, 0x82, 0x5e, 0x08, 0xfa, 0xda, 0xb2, 0x18, 0x42, 0x2f, 0x63, 0xc7, 0xdb, 0x1f, 0x2c, 0xf1,
	0x28, 0xb0, 0xb9, 0x64, 0xed, 0x95, 0xaf, 0x31, 0x12, 0x08, 0xc3, 0x85, 0xf1, 0x6a, 0xad, 0x72,
	0xf2, 0xca, 0x68, 0xec, 0x88, 0x27, 0x00, 0xb7, 0x5c, 0x50, 0x5e, 0xaf, 0x16, 0x74, 0x8f, 0xb1,
	0x38, 0x85, 0xd1, 0xdd, 0x87, 0xb7, 0xab, 0xab, 0x0d, 0x95, 0x25, 0xeb, 0x82, 0xf1, 0x58, 0x8c,
	0xa0, 0x77, 0xbd, 0xb7, 0xa4, 0x25, 0x4b, 0xfc, 0x1d, 0x89, 0x53, 0x18, 0x5e, 0xef, 0x2d, 0x6f,
	0x55, 0xc5, 0xda, 0x53, 0x89, 0x7f, 0x22, 0xd1, 0x85, 0xf8, 0xf3, 0xed, 0x12, 0x47, 0x01, 0x2c,
	0xe7, 0x77, 0x28, 0x02, 0x98, 0xdd, 0xcc, 0xf0, 0x59, 0xe8, 0x3f, 0xbb, 0x99, 0x7d, 0x0a, 0xd5,
	0x38, 0xbd, 0x5c, 0xb4, 0x3e, 0x73, 0x23, 0x59, 0x8c, 0xa0, 0xff, 0x45, 0x4b, 0x5e, 0x2b, 0xdd,
	0x18, 0x1d, 0x40, 0x37, 0xe3, 0xef, 0x3b, 0x76, 0x1e, 0xa3, 0xc3, 0x0c, 0xd6, 0x68, 0xc7, 0xd8,
	0x09, 0xa9, 0xe5, 0x2e, 0xcf, 0xd9, 0x39, 0x8c, 0x03, 0xf9, 0x48, 0xaa, 0xdc, 0x6d, 0x19, 0x8f,
	0xa7, 0xaf, 0xdb, 0x35, 0x87, 0x69, 0x55, 0xce, 0xe2, 0x25, 0x9c, 0x6c, 0x48, 0xcb, 0x92, 0x45,
	0xef, 0x61, 0x51, 0xcf, 0xff, 0xa1, 0xe4, 0xe8, 0xfd, 0xf9, 0xd7, 0x57, 0x15, 0x15, 0x15, 0xa5,
	0x6b, 0x2e, 0xd2, 0x82, 0x3c, 0xff, 0xa4, 0x3a, 0x3d, 0x88, 0x9b, 0x3b, 0x48, 0xdb, 0x93, 0xf8,
	0x76, 0xd2, 0xbc, 0x6f, 0xfe, 0x06, 0x00, 0x00, 0xff, 0xff, 0x3b, 0x14, 0xc0, 0x54, 0x23, 0x02,
	0x00, 0x00,
}
