// Code generated by protoc-gen-go. DO NOT EDIT.
// source: indexer.proto

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

type IndexRequest struct {
	// states to reindex
	States []*protos.State `protobuf:"bytes,1,rep,name=states,proto3" json:"states,omitempty"`
	// network_id of the states
	NetworkId            string   `protobuf:"bytes,2,opt,name=network_id,json=networkId,proto3" json:"network_id,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *IndexRequest) Reset()         { *m = IndexRequest{} }
func (m *IndexRequest) String() string { return proto.CompactTextString(m) }
func (*IndexRequest) ProtoMessage()    {}
func (*IndexRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_2b06a290ab031ed6, []int{0}
}

func (m *IndexRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_IndexRequest.Unmarshal(m, b)
}
func (m *IndexRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_IndexRequest.Marshal(b, m, deterministic)
}
func (m *IndexRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_IndexRequest.Merge(m, src)
}
func (m *IndexRequest) XXX_Size() int {
	return xxx_messageInfo_IndexRequest.Size(m)
}
func (m *IndexRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_IndexRequest.DiscardUnknown(m)
}

var xxx_messageInfo_IndexRequest proto.InternalMessageInfo

func (m *IndexRequest) GetStates() []*protos.State {
	if m != nil {
		return m.States
	}
	return nil
}

func (m *IndexRequest) GetNetworkId() string {
	if m != nil {
		return m.NetworkId
	}
	return ""
}

type IndexResponse struct {
	// state_errors are errors experienced trying to index specific pieces of state.
	StateErrors          []*protos.IDAndError `protobuf:"bytes,1,rep,name=state_errors,json=stateErrors,proto3" json:"state_errors,omitempty"`
	XXX_NoUnkeyedLiteral struct{}             `json:"-"`
	XXX_unrecognized     []byte               `json:"-"`
	XXX_sizecache        int32                `json:"-"`
}

func (m *IndexResponse) Reset()         { *m = IndexResponse{} }
func (m *IndexResponse) String() string { return proto.CompactTextString(m) }
func (*IndexResponse) ProtoMessage()    {}
func (*IndexResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_2b06a290ab031ed6, []int{1}
}

func (m *IndexResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_IndexResponse.Unmarshal(m, b)
}
func (m *IndexResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_IndexResponse.Marshal(b, m, deterministic)
}
func (m *IndexResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_IndexResponse.Merge(m, src)
}
func (m *IndexResponse) XXX_Size() int {
	return xxx_messageInfo_IndexResponse.Size(m)
}
func (m *IndexResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_IndexResponse.DiscardUnknown(m)
}

var xxx_messageInfo_IndexResponse proto.InternalMessageInfo

func (m *IndexResponse) GetStateErrors() []*protos.IDAndError {
	if m != nil {
		return m.StateErrors
	}
	return nil
}

type PrepareReindexRequest struct {
	// indexer_id being reindexed
	IndexerId string `protobuf:"bytes,1,opt,name=indexer_id,json=indexerId,proto3" json:"indexer_id,omitempty"`
	// from_version is the indexer's current (actual) version
	FromVersion uint32 `protobuf:"varint,2,opt,name=from_version,json=fromVersion,proto3" json:"from_version,omitempty"`
	// to_version is the indexer's future (desired) version
	ToVersion uint32 `protobuf:"varint,3,opt,name=to_version,json=toVersion,proto3" json:"to_version,omitempty"`
	// is_first is true iff this is the first time this indexer is being reindexed
	IsFirst              bool     `protobuf:"varint,4,opt,name=is_first,json=isFirst,proto3" json:"is_first,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *PrepareReindexRequest) Reset()         { *m = PrepareReindexRequest{} }
func (m *PrepareReindexRequest) String() string { return proto.CompactTextString(m) }
func (*PrepareReindexRequest) ProtoMessage()    {}
func (*PrepareReindexRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_2b06a290ab031ed6, []int{2}
}

func (m *PrepareReindexRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_PrepareReindexRequest.Unmarshal(m, b)
}
func (m *PrepareReindexRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_PrepareReindexRequest.Marshal(b, m, deterministic)
}
func (m *PrepareReindexRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_PrepareReindexRequest.Merge(m, src)
}
func (m *PrepareReindexRequest) XXX_Size() int {
	return xxx_messageInfo_PrepareReindexRequest.Size(m)
}
func (m *PrepareReindexRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_PrepareReindexRequest.DiscardUnknown(m)
}

var xxx_messageInfo_PrepareReindexRequest proto.InternalMessageInfo

func (m *PrepareReindexRequest) GetIndexerId() string {
	if m != nil {
		return m.IndexerId
	}
	return ""
}

func (m *PrepareReindexRequest) GetFromVersion() uint32 {
	if m != nil {
		return m.FromVersion
	}
	return 0
}

func (m *PrepareReindexRequest) GetToVersion() uint32 {
	if m != nil {
		return m.ToVersion
	}
	return 0
}

func (m *PrepareReindexRequest) GetIsFirst() bool {
	if m != nil {
		return m.IsFirst
	}
	return false
}

type PrepareReindexResponse struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *PrepareReindexResponse) Reset()         { *m = PrepareReindexResponse{} }
func (m *PrepareReindexResponse) String() string { return proto.CompactTextString(m) }
func (*PrepareReindexResponse) ProtoMessage()    {}
func (*PrepareReindexResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_2b06a290ab031ed6, []int{3}
}

func (m *PrepareReindexResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_PrepareReindexResponse.Unmarshal(m, b)
}
func (m *PrepareReindexResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_PrepareReindexResponse.Marshal(b, m, deterministic)
}
func (m *PrepareReindexResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_PrepareReindexResponse.Merge(m, src)
}
func (m *PrepareReindexResponse) XXX_Size() int {
	return xxx_messageInfo_PrepareReindexResponse.Size(m)
}
func (m *PrepareReindexResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_PrepareReindexResponse.DiscardUnknown(m)
}

var xxx_messageInfo_PrepareReindexResponse proto.InternalMessageInfo

type CompleteReindexRequest struct {
	// indexer_id being reindexed
	IndexerId string `protobuf:"bytes,1,opt,name=indexer_id,json=indexerId,proto3" json:"indexer_id,omitempty"`
	// from_version is the indexer's current (actual) version
	FromVersion uint32 `protobuf:"varint,2,opt,name=from_version,json=fromVersion,proto3" json:"from_version,omitempty"`
	// to_version is the indexer's future (desired) version
	ToVersion            uint32   `protobuf:"varint,3,opt,name=to_version,json=toVersion,proto3" json:"to_version,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *CompleteReindexRequest) Reset()         { *m = CompleteReindexRequest{} }
func (m *CompleteReindexRequest) String() string { return proto.CompactTextString(m) }
func (*CompleteReindexRequest) ProtoMessage()    {}
func (*CompleteReindexRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_2b06a290ab031ed6, []int{4}
}

func (m *CompleteReindexRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CompleteReindexRequest.Unmarshal(m, b)
}
func (m *CompleteReindexRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CompleteReindexRequest.Marshal(b, m, deterministic)
}
func (m *CompleteReindexRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CompleteReindexRequest.Merge(m, src)
}
func (m *CompleteReindexRequest) XXX_Size() int {
	return xxx_messageInfo_CompleteReindexRequest.Size(m)
}
func (m *CompleteReindexRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_CompleteReindexRequest.DiscardUnknown(m)
}

var xxx_messageInfo_CompleteReindexRequest proto.InternalMessageInfo

func (m *CompleteReindexRequest) GetIndexerId() string {
	if m != nil {
		return m.IndexerId
	}
	return ""
}

func (m *CompleteReindexRequest) GetFromVersion() uint32 {
	if m != nil {
		return m.FromVersion
	}
	return 0
}

func (m *CompleteReindexRequest) GetToVersion() uint32 {
	if m != nil {
		return m.ToVersion
	}
	return 0
}

type CompleteReindexResponse struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *CompleteReindexResponse) Reset()         { *m = CompleteReindexResponse{} }
func (m *CompleteReindexResponse) String() string { return proto.CompactTextString(m) }
func (*CompleteReindexResponse) ProtoMessage()    {}
func (*CompleteReindexResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_2b06a290ab031ed6, []int{5}
}

func (m *CompleteReindexResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CompleteReindexResponse.Unmarshal(m, b)
}
func (m *CompleteReindexResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CompleteReindexResponse.Marshal(b, m, deterministic)
}
func (m *CompleteReindexResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CompleteReindexResponse.Merge(m, src)
}
func (m *CompleteReindexResponse) XXX_Size() int {
	return xxx_messageInfo_CompleteReindexResponse.Size(m)
}
func (m *CompleteReindexResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_CompleteReindexResponse.DiscardUnknown(m)
}

var xxx_messageInfo_CompleteReindexResponse proto.InternalMessageInfo

func init() {
	proto.RegisterType((*IndexRequest)(nil), "magma.orc8r.state.IndexRequest")
	proto.RegisterType((*IndexResponse)(nil), "magma.orc8r.state.IndexResponse")
	proto.RegisterType((*PrepareReindexRequest)(nil), "magma.orc8r.state.PrepareReindexRequest")
	proto.RegisterType((*PrepareReindexResponse)(nil), "magma.orc8r.state.PrepareReindexResponse")
	proto.RegisterType((*CompleteReindexRequest)(nil), "magma.orc8r.state.CompleteReindexRequest")
	proto.RegisterType((*CompleteReindexResponse)(nil), "magma.orc8r.state.CompleteReindexResponse")
}

func init() { proto.RegisterFile("indexer.proto", fileDescriptor_2b06a290ab031ed6) }

var fileDescriptor_2b06a290ab031ed6 = []byte{
	// 387 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xbc, 0x53, 0xcb, 0x4e, 0xe3, 0x30,
	0x14, 0x9d, 0xb4, 0x33, 0x7d, 0xdc, 0xb6, 0x33, 0x1a, 0x4b, 0xd3, 0xa6, 0x91, 0xaa, 0x09, 0x59,
	0x85, 0x2e, 0x52, 0xd4, 0x6e, 0x10, 0x3b, 0x9e, 0x52, 0x04, 0x0b, 0x14, 0x24, 0x24, 0xd8, 0x44,
	0xa1, 0x71, 0x2b, 0x03, 0x89, 0x83, 0x6d, 0x0a, 0x7c, 0x07, 0xbf, 0xc1, 0x47, 0x22, 0xdb, 0x69,
	0xd5, 0x47, 0x90, 0xba, 0x62, 0x55, 0xfb, 0x9e, 0x73, 0xee, 0x39, 0xf7, 0xd6, 0x81, 0x16, 0x49,
	0x63, 0xfc, 0x8a, 0x99, 0x97, 0x31, 0x2a, 0x28, 0xfa, 0x9b, 0x44, 0xd3, 0x24, 0xf2, 0x28, 0x1b,
	0xef, 0x33, 0x8f, 0x8b, 0x48, 0x60, 0xab, 0xa7, 0x2e, 0x03, 0x85, 0xf3, 0x01, 0xc7, 0x6c, 0x46,
	0xc6, 0x78, 0xb4, 0x37, 0xd2, 0x0a, 0xcb, 0x5c, 0x85, 0xa5, 0x44, 0x23, 0xce, 0x0d, 0x34, 0x7d,
	0xd9, 0x3c, 0xc0, 0x4f, 0xcf, 0x98, 0x0b, 0xd4, 0x87, 0x8a, 0x82, 0xb9, 0x69, 0xd8, 0x65, 0xb7,
	0x31, 0x44, 0xde, 0xb2, 0xd9, 0x95, 0x84, 0x82, 0x9c, 0x81, 0x7a, 0x00, 0x29, 0x16, 0x2f, 0x94,
	0x3d, 0x84, 0x24, 0x36, 0x4b, 0xb6, 0xe1, 0xd6, 0x83, 0x7a, 0x5e, 0xf1, 0x63, 0xe7, 0x1c, 0x5a,
	0x79, 0x6b, 0x9e, 0xd1, 0x94, 0x63, 0x74, 0x00, 0x4d, 0xa5, 0x0c, 0x31, 0x63, 0x94, 0xcd, 0x1d,
	0x3a, 0x2b, 0x0e, 0xfe, 0xc9, 0x61, 0x1a, 0x9f, 0x4a, 0x3c, 0x68, 0x28, 0xb2, 0x3a, 0x73, 0xe7,
	0xdd, 0x80, 0x7f, 0x97, 0x0c, 0x67, 0x11, 0xc3, 0x01, 0x26, 0xcb, 0x89, 0x7b, 0x00, 0xf9, 0x7a,
	0x64, 0x0a, 0x43, 0xa7, 0xc8, 0x2b, 0x7e, 0x8c, 0x76, 0xa0, 0x39, 0x61, 0x34, 0x09, 0x67, 0x98,
	0x71, 0x42, 0x53, 0x15, 0xb3, 0x15, 0x34, 0x64, 0xed, 0x5a, 0x97, 0x64, 0x07, 0x41, 0x17, 0x84,
	0xb2, 0x22, 0xd4, 0x05, 0x9d, 0xc3, 0x5d, 0xa8, 0x11, 0x1e, 0x4e, 0x08, 0xe3, 0xc2, 0xfc, 0x69,
	0x1b, 0x6e, 0x2d, 0xa8, 0x12, 0x7e, 0x26, 0xaf, 0x8e, 0x09, 0xed, 0xf5, 0x50, 0x7a, 0x56, 0xe7,
	0x0d, 0xda, 0xc7, 0x34, 0xc9, 0x1e, 0xb1, 0xf8, 0xee, 0xbc, 0x4e, 0x17, 0x3a, 0x1b, 0xd6, 0x3a,
	0xd5, 0xf0, 0xa3, 0x04, 0x55, 0x5f, 0x5b, 0xa1, 0x0b, 0xf8, 0xa5, 0x8e, 0xe8, 0xbf, 0xb7, 0xf1,
	0x9e, 0xbc, 0xe5, 0x37, 0x61, 0xd9, 0x5f, 0x13, 0xf2, 0x69, 0x7f, 0xa0, 0x29, 0xfc, 0x5e, 0xdd,
	0x04, 0x72, 0x0b, 0x54, 0x85, 0xff, 0xa0, 0xb5, 0xbb, 0x05, 0x73, 0x61, 0x74, 0x0f, 0x7f, 0xd6,
	0xa6, 0x43, 0x45, 0xfa, 0xe2, 0xe5, 0x5b, 0xfd, 0x6d, 0xa8, 0x73, 0xaf, 0xa3, 0xda, 0x6d, 0x45,
	0x7f, 0x32, 0x77, 0xfa, 0x77, 0xf4, 0x19, 0x00, 0x00, 0xff, 0xff, 0x0f, 0xc7, 0x26, 0xc7, 0x8a,
	0x03, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConnInterface

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion6

// IndexerClient is the client API for Indexer service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type IndexerClient interface {
	// Index a set of states by forwarding to locally-registered indexers.
	Index(ctx context.Context, in *IndexRequest, opts ...grpc.CallOption) (*IndexResponse, error)
	// PrepareReindex of a particular indexer.
	PrepareReindex(ctx context.Context, in *PrepareReindexRequest, opts ...grpc.CallOption) (*PrepareReindexResponse, error)
	// CompleteReindex of a particular indexer.
	CompleteReindex(ctx context.Context, in *CompleteReindexRequest, opts ...grpc.CallOption) (*CompleteReindexResponse, error)
}

type indexerClient struct {
	cc grpc.ClientConnInterface
}

func NewIndexerClient(cc grpc.ClientConnInterface) IndexerClient {
	return &indexerClient{cc}
}

func (c *indexerClient) Index(ctx context.Context, in *IndexRequest, opts ...grpc.CallOption) (*IndexResponse, error) {
	out := new(IndexResponse)
	err := c.cc.Invoke(ctx, "/magma.orc8r.state.Indexer/Index", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *indexerClient) PrepareReindex(ctx context.Context, in *PrepareReindexRequest, opts ...grpc.CallOption) (*PrepareReindexResponse, error) {
	out := new(PrepareReindexResponse)
	err := c.cc.Invoke(ctx, "/magma.orc8r.state.Indexer/PrepareReindex", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *indexerClient) CompleteReindex(ctx context.Context, in *CompleteReindexRequest, opts ...grpc.CallOption) (*CompleteReindexResponse, error) {
	out := new(CompleteReindexResponse)
	err := c.cc.Invoke(ctx, "/magma.orc8r.state.Indexer/CompleteReindex", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// IndexerServer is the server API for Indexer service.
type IndexerServer interface {
	// Index a set of states by forwarding to locally-registered indexers.
	Index(context.Context, *IndexRequest) (*IndexResponse, error)
	// PrepareReindex of a particular indexer.
	PrepareReindex(context.Context, *PrepareReindexRequest) (*PrepareReindexResponse, error)
	// CompleteReindex of a particular indexer.
	CompleteReindex(context.Context, *CompleteReindexRequest) (*CompleteReindexResponse, error)
}

// UnimplementedIndexerServer can be embedded to have forward compatible implementations.
type UnimplementedIndexerServer struct {
}

func (*UnimplementedIndexerServer) Index(ctx context.Context, req *IndexRequest) (*IndexResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Index not implemented")
}
func (*UnimplementedIndexerServer) PrepareReindex(ctx context.Context, req *PrepareReindexRequest) (*PrepareReindexResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method PrepareReindex not implemented")
}
func (*UnimplementedIndexerServer) CompleteReindex(ctx context.Context, req *CompleteReindexRequest) (*CompleteReindexResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CompleteReindex not implemented")
}

func RegisterIndexerServer(s *grpc.Server, srv IndexerServer) {
	s.RegisterService(&_Indexer_serviceDesc, srv)
}

func _Indexer_Index_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(IndexRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(IndexerServer).Index(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/magma.orc8r.state.Indexer/Index",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(IndexerServer).Index(ctx, req.(*IndexRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Indexer_PrepareReindex_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(PrepareReindexRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(IndexerServer).PrepareReindex(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/magma.orc8r.state.Indexer/PrepareReindex",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(IndexerServer).PrepareReindex(ctx, req.(*PrepareReindexRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Indexer_CompleteReindex_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CompleteReindexRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(IndexerServer).CompleteReindex(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/magma.orc8r.state.Indexer/CompleteReindex",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(IndexerServer).CompleteReindex(ctx, req.(*CompleteReindexRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _Indexer_serviceDesc = grpc.ServiceDesc{
	ServiceName: "magma.orc8r.state.Indexer",
	HandlerType: (*IndexerServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Index",
			Handler:    _Indexer_Index_Handler,
		},
		{
			MethodName: "PrepareReindex",
			Handler:    _Indexer_PrepareReindex_Handler,
		},
		{
			MethodName: "CompleteReindex",
			Handler:    _Indexer_CompleteReindex_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "indexer.proto",
}
