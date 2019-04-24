// Code generated by protoc-gen-go. DO NOT EDIT.
// source: orc8r/protos/state.proto

package protos // import "magma/orc8r/cloud/go/protos"

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

type StateID struct {
	Type                 string   `protobuf:"bytes,1,opt,name=type,proto3" json:"type,omitempty"`
	DeviceID             string   `protobuf:"bytes,2,opt,name=deviceID,proto3" json:"deviceID,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *StateID) Reset()         { *m = StateID{} }
func (m *StateID) String() string { return proto.CompactTextString(m) }
func (*StateID) ProtoMessage()    {}
func (*StateID) Descriptor() ([]byte, []int) {
	return fileDescriptor_state_40967c0a9a3f0140, []int{0}
}
func (m *StateID) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_StateID.Unmarshal(m, b)
}
func (m *StateID) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_StateID.Marshal(b, m, deterministic)
}
func (dst *StateID) XXX_Merge(src proto.Message) {
	xxx_messageInfo_StateID.Merge(dst, src)
}
func (m *StateID) XXX_Size() int {
	return xxx_messageInfo_StateID.Size(m)
}
func (m *StateID) XXX_DiscardUnknown() {
	xxx_messageInfo_StateID.DiscardUnknown(m)
}

var xxx_messageInfo_StateID proto.InternalMessageInfo

func (m *StateID) GetType() string {
	if m != nil {
		return m.Type
	}
	return ""
}

func (m *StateID) GetDeviceID() string {
	if m != nil {
		return m.DeviceID
	}
	return ""
}

type GetStatesRequest struct {
	NetworkID            string     `protobuf:"bytes,1,opt,name=networkID,proto3" json:"networkID,omitempty"`
	Ids                  []*StateID `protobuf:"bytes,2,rep,name=ids,proto3" json:"ids,omitempty"`
	XXX_NoUnkeyedLiteral struct{}   `json:"-"`
	XXX_unrecognized     []byte     `json:"-"`
	XXX_sizecache        int32      `json:"-"`
}

func (m *GetStatesRequest) Reset()         { *m = GetStatesRequest{} }
func (m *GetStatesRequest) String() string { return proto.CompactTextString(m) }
func (*GetStatesRequest) ProtoMessage()    {}
func (*GetStatesRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_state_40967c0a9a3f0140, []int{1}
}
func (m *GetStatesRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetStatesRequest.Unmarshal(m, b)
}
func (m *GetStatesRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetStatesRequest.Marshal(b, m, deterministic)
}
func (dst *GetStatesRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetStatesRequest.Merge(dst, src)
}
func (m *GetStatesRequest) XXX_Size() int {
	return xxx_messageInfo_GetStatesRequest.Size(m)
}
func (m *GetStatesRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_GetStatesRequest.DiscardUnknown(m)
}

var xxx_messageInfo_GetStatesRequest proto.InternalMessageInfo

func (m *GetStatesRequest) GetNetworkID() string {
	if m != nil {
		return m.NetworkID
	}
	return ""
}

func (m *GetStatesRequest) GetIds() []*StateID {
	if m != nil {
		return m.Ids
	}
	return nil
}

type GetStatesResponse struct {
	States               []*State `protobuf:"bytes,1,rep,name=states,proto3" json:"states,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *GetStatesResponse) Reset()         { *m = GetStatesResponse{} }
func (m *GetStatesResponse) String() string { return proto.CompactTextString(m) }
func (*GetStatesResponse) ProtoMessage()    {}
func (*GetStatesResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_state_40967c0a9a3f0140, []int{2}
}
func (m *GetStatesResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetStatesResponse.Unmarshal(m, b)
}
func (m *GetStatesResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetStatesResponse.Marshal(b, m, deterministic)
}
func (dst *GetStatesResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetStatesResponse.Merge(dst, src)
}
func (m *GetStatesResponse) XXX_Size() int {
	return xxx_messageInfo_GetStatesResponse.Size(m)
}
func (m *GetStatesResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_GetStatesResponse.DiscardUnknown(m)
}

var xxx_messageInfo_GetStatesResponse proto.InternalMessageInfo

func (m *GetStatesResponse) GetStates() []*State {
	if m != nil {
		return m.States
	}
	return nil
}

type ReportStatesRequest struct {
	States               []*State `protobuf:"bytes,1,rep,name=states,proto3" json:"states,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ReportStatesRequest) Reset()         { *m = ReportStatesRequest{} }
func (m *ReportStatesRequest) String() string { return proto.CompactTextString(m) }
func (*ReportStatesRequest) ProtoMessage()    {}
func (*ReportStatesRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_state_40967c0a9a3f0140, []int{3}
}
func (m *ReportStatesRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ReportStatesRequest.Unmarshal(m, b)
}
func (m *ReportStatesRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ReportStatesRequest.Marshal(b, m, deterministic)
}
func (dst *ReportStatesRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ReportStatesRequest.Merge(dst, src)
}
func (m *ReportStatesRequest) XXX_Size() int {
	return xxx_messageInfo_ReportStatesRequest.Size(m)
}
func (m *ReportStatesRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_ReportStatesRequest.DiscardUnknown(m)
}

var xxx_messageInfo_ReportStatesRequest proto.InternalMessageInfo

func (m *ReportStatesRequest) GetStates() []*State {
	if m != nil {
		return m.States
	}
	return nil
}

type DeleteStatesRequest struct {
	NetworkID            string     `protobuf:"bytes,1,opt,name=networkID,proto3" json:"networkID,omitempty"`
	Ids                  []*StateID `protobuf:"bytes,2,rep,name=ids,proto3" json:"ids,omitempty"`
	XXX_NoUnkeyedLiteral struct{}   `json:"-"`
	XXX_unrecognized     []byte     `json:"-"`
	XXX_sizecache        int32      `json:"-"`
}

func (m *DeleteStatesRequest) Reset()         { *m = DeleteStatesRequest{} }
func (m *DeleteStatesRequest) String() string { return proto.CompactTextString(m) }
func (*DeleteStatesRequest) ProtoMessage()    {}
func (*DeleteStatesRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_state_40967c0a9a3f0140, []int{4}
}
func (m *DeleteStatesRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_DeleteStatesRequest.Unmarshal(m, b)
}
func (m *DeleteStatesRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_DeleteStatesRequest.Marshal(b, m, deterministic)
}
func (dst *DeleteStatesRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_DeleteStatesRequest.Merge(dst, src)
}
func (m *DeleteStatesRequest) XXX_Size() int {
	return xxx_messageInfo_DeleteStatesRequest.Size(m)
}
func (m *DeleteStatesRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_DeleteStatesRequest.DiscardUnknown(m)
}

var xxx_messageInfo_DeleteStatesRequest proto.InternalMessageInfo

func (m *DeleteStatesRequest) GetNetworkID() string {
	if m != nil {
		return m.NetworkID
	}
	return ""
}

func (m *DeleteStatesRequest) GetIds() []*StateID {
	if m != nil {
		return m.Ids
	}
	return nil
}

func init() {
	proto.RegisterType((*StateID)(nil), "magma.orc8r.StateID")
	proto.RegisterType((*GetStatesRequest)(nil), "magma.orc8r.GetStatesRequest")
	proto.RegisterType((*GetStatesResponse)(nil), "magma.orc8r.GetStatesResponse")
	proto.RegisterType((*ReportStatesRequest)(nil), "magma.orc8r.ReportStatesRequest")
	proto.RegisterType((*DeleteStatesRequest)(nil), "magma.orc8r.DeleteStatesRequest")
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// StateServiceClient is the client API for StateService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type StateServiceClient interface {
	GetStates(ctx context.Context, in *GetStatesRequest, opts ...grpc.CallOption) (*GetStatesResponse, error)
	ReportStates(ctx context.Context, in *ReportStatesRequest, opts ...grpc.CallOption) (*Void, error)
	DeleteStates(ctx context.Context, in *DeleteStatesRequest, opts ...grpc.CallOption) (*Void, error)
}

type stateServiceClient struct {
	cc *grpc.ClientConn
}

func NewStateServiceClient(cc *grpc.ClientConn) StateServiceClient {
	return &stateServiceClient{cc}
}

func (c *stateServiceClient) GetStates(ctx context.Context, in *GetStatesRequest, opts ...grpc.CallOption) (*GetStatesResponse, error) {
	out := new(GetStatesResponse)
	err := c.cc.Invoke(ctx, "/magma.orc8r.StateService/GetStates", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *stateServiceClient) ReportStates(ctx context.Context, in *ReportStatesRequest, opts ...grpc.CallOption) (*Void, error) {
	out := new(Void)
	err := c.cc.Invoke(ctx, "/magma.orc8r.StateService/ReportStates", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *stateServiceClient) DeleteStates(ctx context.Context, in *DeleteStatesRequest, opts ...grpc.CallOption) (*Void, error) {
	out := new(Void)
	err := c.cc.Invoke(ctx, "/magma.orc8r.StateService/DeleteStates", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// StateServiceServer is the server API for StateService service.
type StateServiceServer interface {
	GetStates(context.Context, *GetStatesRequest) (*GetStatesResponse, error)
	ReportStates(context.Context, *ReportStatesRequest) (*Void, error)
	DeleteStates(context.Context, *DeleteStatesRequest) (*Void, error)
}

func RegisterStateServiceServer(s *grpc.Server, srv StateServiceServer) {
	s.RegisterService(&_StateService_serviceDesc, srv)
}

func _StateService_GetStates_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetStatesRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(StateServiceServer).GetStates(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/magma.orc8r.StateService/GetStates",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(StateServiceServer).GetStates(ctx, req.(*GetStatesRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _StateService_ReportStates_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ReportStatesRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(StateServiceServer).ReportStates(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/magma.orc8r.StateService/ReportStates",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(StateServiceServer).ReportStates(ctx, req.(*ReportStatesRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _StateService_DeleteStates_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeleteStatesRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(StateServiceServer).DeleteStates(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/magma.orc8r.StateService/DeleteStates",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(StateServiceServer).DeleteStates(ctx, req.(*DeleteStatesRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _StateService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "magma.orc8r.StateService",
	HandlerType: (*StateServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetStates",
			Handler:    _StateService_GetStates_Handler,
		},
		{
			MethodName: "ReportStates",
			Handler:    _StateService_ReportStates_Handler,
		},
		{
			MethodName: "DeleteStates",
			Handler:    _StateService_DeleteStates_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "orc8r/protos/state.proto",
}

func init() { proto.RegisterFile("orc8r/protos/state.proto", fileDescriptor_state_40967c0a9a3f0140) }

var fileDescriptor_state_40967c0a9a3f0140 = []byte{
	// 322 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xb4, 0x52, 0xc1, 0x4e, 0x83, 0x40,
	0x14, 0x94, 0xd6, 0x54, 0xfb, 0xda, 0x83, 0x7d, 0xf5, 0x80, 0x68, 0x4d, 0xc3, 0xc1, 0x34, 0x1e,
	0xc0, 0xb4, 0x17, 0x3d, 0x19, 0x0d, 0xc6, 0x34, 0xf1, 0x44, 0x13, 0x63, 0xf4, 0x84, 0xf0, 0xd2,
	0x10, 0x0b, 0x8b, 0xbb, 0x5b, 0x8d, 0xbf, 0xec, 0x57, 0x18, 0x76, 0x49, 0x05, 0xc4, 0x83, 0x07,
	0x4f, 0xb0, 0xf3, 0x66, 0x26, 0x3b, 0xb3, 0x0f, 0x4c, 0xc6, 0xc3, 0x73, 0xee, 0x66, 0x9c, 0x49,
	0x26, 0x5c, 0x21, 0x03, 0x49, 0x8e, 0x3a, 0x60, 0x2f, 0x09, 0x96, 0x49, 0xe0, 0xa8, 0xb9, 0x75,
	0x50, 0xa1, 0x85, 0x2c, 0x49, 0x58, 0xaa, 0x79, 0xd6, 0xa8, 0xea, 0x40, 0xfc, 0x2d, 0x0e, 0x69,
	0x76, 0x36, 0xd3, 0x63, 0xfb, 0x02, 0x76, 0x16, 0xb9, 0xeb, 0xdc, 0x43, 0x84, 0x6d, 0xf9, 0x91,
	0x91, 0x69, 0x8c, 0x8d, 0x49, 0xd7, 0x57, 0xff, 0x68, 0xc1, 0x6e, 0x44, 0xb9, 0x62, 0xee, 0x99,
	0x2d, 0x85, 0x6f, 0xce, 0xf6, 0x03, 0xec, 0xdd, 0x92, 0x54, 0x6a, 0xe1, 0xd3, 0xeb, 0x9a, 0x84,
	0xc4, 0x23, 0xe8, 0xa6, 0x24, 0xdf, 0x19, 0x7f, 0x99, 0x7b, 0x85, 0xd1, 0x37, 0x80, 0x27, 0xd0,
	0x8e, 0x23, 0x61, 0xb6, 0xc6, 0xed, 0x49, 0x6f, 0xba, 0xef, 0x94, 0x12, 0x38, 0xc5, 0x25, 0xfc,
	0x9c, 0x60, 0x5f, 0xc2, 0xa0, 0xe4, 0x2c, 0x32, 0x96, 0x0a, 0xc2, 0x53, 0xe8, 0xa8, 0xfc, 0xc2,
	0x34, 0x94, 0x1e, 0x7f, 0xea, 0xfd, 0x82, 0x61, 0x5f, 0xc1, 0xd0, 0xa7, 0x8c, 0xf1, 0xda, 0xed,
	0xfe, 0x62, 0xf1, 0x04, 0x43, 0x8f, 0x56, 0x24, 0xe9, 0x1f, 0x02, 0x4e, 0x3f, 0x0d, 0xe8, 0x2b,
	0x60, 0xa1, 0xdf, 0x03, 0xef, 0xa0, 0xbb, 0x49, 0x8c, 0xa3, 0x8a, 0xb0, 0xde, 0xb1, 0x75, 0xfc,
	0xdb, 0x58, 0x17, 0x65, 0x6f, 0xe1, 0x0d, 0xf4, 0xcb, 0xf1, 0x71, 0x5c, 0x51, 0x34, 0x34, 0x63,
	0x0d, 0x2a, 0x8c, 0x7b, 0x16, 0x47, 0xda, 0xa6, 0x5c, 0x41, 0xcd, 0xa6, 0xa1, 0x9d, 0x46, 0x9b,
	0xeb, 0xd1, 0xe3, 0xa1, 0x42, 0x5d, 0xbd, 0x89, 0xe1, 0x8a, 0xad, 0x23, 0x77, 0xc9, 0x8a, 0x95,
	0x7c, 0xee, 0xa8, 0xef, 0xec, 0x2b, 0x00, 0x00, 0xff, 0xff, 0x0c, 0xd7, 0xb9, 0x90, 0xeb, 0x02,
	0x00, 0x00,
}
