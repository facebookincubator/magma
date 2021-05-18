//
//Copyright 2020 The Magma Authors.
//
//This source code is licensed under the BSD-style license found in the
//LICENSE file in the root directory of this source tree.
//
//Unless required by applicable law or agreed to in writing, software
//distributed under the License is distributed on an "AS IS" BASIS,
//WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
//See the License for the specific language governing permissions and
//limitations under the License.

// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.26.0
// 	protoc        (unknown)
// source: authorization.proto

package protos

import (
	context "context"
	reflect "reflect"
	sync "sync"

	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type CoaResponseCoaResponseTypeEnum int32

const (
	CoaResponse_NAK CoaResponseCoaResponseTypeEnum = 0
	CoaResponse_ACK CoaResponseCoaResponseTypeEnum = 1
)

// Enum value maps for CoaResponseCoaResponseTypeEnum.
var (
	CoaResponseCoaResponseTypeEnum_name = map[int32]string{
		0: "NAK",
		1: "ACK",
	}
	CoaResponseCoaResponseTypeEnum_value = map[string]int32{
		"NAK": 0,
		"ACK": 1,
	}
)

func (x CoaResponseCoaResponseTypeEnum) Enum() *CoaResponseCoaResponseTypeEnum {
	p := new(CoaResponseCoaResponseTypeEnum)
	*p = x
	return p
}

func (x CoaResponseCoaResponseTypeEnum) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (CoaResponseCoaResponseTypeEnum) Descriptor() protoreflect.EnumDescriptor {
	return file_authorization_proto_enumTypes[0].Descriptor()
}

func (CoaResponseCoaResponseTypeEnum) Type() protoreflect.EnumType {
	return &file_authorization_proto_enumTypes[0]
}

func (x CoaResponseCoaResponseTypeEnum) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use CoaResponseCoaResponseTypeEnum.Descriptor instead.
func (CoaResponseCoaResponseTypeEnum) EnumDescriptor() ([]byte, []int) {
	return file_authorization_proto_rawDescGZIP(), []int{2, 0}
}

// update_request with usages & included context
type ChangeRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Ctx              *Context `protobuf:"bytes,1,opt,name=ctx,proto3" json:"ctx,omitempty"`
	JsonTrficClasses string   `protobuf:"bytes,2,opt,name=json_trfic_classes,json=jsonTrficClasses,proto3" json:"json_trfic_classes,omitempty"`
}

func (x *ChangeRequest) Reset() {
	*x = ChangeRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_authorization_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ChangeRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ChangeRequest) ProtoMessage() {}

func (x *ChangeRequest) ProtoReflect() protoreflect.Message {
	mi := &file_authorization_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ChangeRequest.ProtoReflect.Descriptor instead.
func (*ChangeRequest) Descriptor() ([]byte, []int) {
	return file_authorization_proto_rawDescGZIP(), []int{0}
}

func (x *ChangeRequest) GetCtx() *Context {
	if x != nil {
		return x.Ctx
	}
	return nil
}

func (x *ChangeRequest) GetJsonTrficClasses() string {
	if x != nil {
		return x.JsonTrficClasses
	}
	return ""
}

type DisconnectRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Ctx *Context `protobuf:"bytes,1,opt,name=ctx,proto3" json:"ctx,omitempty"`
}

func (x *DisconnectRequest) Reset() {
	*x = DisconnectRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_authorization_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DisconnectRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DisconnectRequest) ProtoMessage() {}

func (x *DisconnectRequest) ProtoReflect() protoreflect.Message {
	mi := &file_authorization_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DisconnectRequest.ProtoReflect.Descriptor instead.
func (*DisconnectRequest) Descriptor() ([]byte, []int) {
	return file_authorization_proto_rawDescGZIP(), []int{1}
}

func (x *DisconnectRequest) GetCtx() *Context {
	if x != nil {
		return x.Ctx
	}
	return nil
}

type CoaResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	CoaResponseType CoaResponseCoaResponseTypeEnum `protobuf:"varint,1,opt,name=coa_response_type,json=coaResponseType,proto3,enum=aaa.protos.CoaResponseCoaResponseTypeEnum" json:"coa_response_type,omitempty"`
	Ctx             *Context                       `protobuf:"bytes,2,opt,name=ctx,proto3" json:"ctx,omitempty"`
}

func (x *CoaResponse) Reset() {
	*x = CoaResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_authorization_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CoaResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CoaResponse) ProtoMessage() {}

func (x *CoaResponse) ProtoReflect() protoreflect.Message {
	mi := &file_authorization_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CoaResponse.ProtoReflect.Descriptor instead.
func (*CoaResponse) Descriptor() ([]byte, []int) {
	return file_authorization_proto_rawDescGZIP(), []int{2}
}

func (x *CoaResponse) GetCoaResponseType() CoaResponseCoaResponseTypeEnum {
	if x != nil {
		return x.CoaResponseType
	}
	return CoaResponse_NAK
}

func (x *CoaResponse) GetCtx() *Context {
	if x != nil {
		return x.Ctx
	}
	return nil
}

var File_authorization_proto protoreflect.FileDescriptor

var file_authorization_proto_rawDesc = []byte{
	0x0a, 0x13, 0x61, 0x75, 0x74, 0x68, 0x6f, 0x72, 0x69, 0x7a, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x0a, 0x61, 0x61, 0x61, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x73, 0x1a, 0x0d, 0x63, 0x6f, 0x6e, 0x74, 0x65, 0x78, 0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x22, 0x65, 0x0a, 0x0e, 0x63, 0x68, 0x61, 0x6e, 0x67, 0x65, 0x5f, 0x72, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x12, 0x25, 0x0a, 0x03, 0x63, 0x74, 0x78, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32,
	0x13, 0x2e, 0x61, 0x61, 0x61, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x73, 0x2e, 0x63, 0x6f, 0x6e,
	0x74, 0x65, 0x78, 0x74, 0x52, 0x03, 0x63, 0x74, 0x78, 0x12, 0x2c, 0x0a, 0x12, 0x6a, 0x73, 0x6f,
	0x6e, 0x5f, 0x74, 0x72, 0x66, 0x69, 0x63, 0x5f, 0x63, 0x6c, 0x61, 0x73, 0x73, 0x65, 0x73, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x10, 0x6a, 0x73, 0x6f, 0x6e, 0x54, 0x72, 0x66, 0x69, 0x63,
	0x43, 0x6c, 0x61, 0x73, 0x73, 0x65, 0x73, 0x22, 0x3b, 0x0a, 0x12, 0x64, 0x69, 0x73, 0x63, 0x6f,
	0x6e, 0x6e, 0x65, 0x63, 0x74, 0x5f, 0x72, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x25, 0x0a,
	0x03, 0x63, 0x74, 0x78, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x13, 0x2e, 0x61, 0x61, 0x61,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x73, 0x2e, 0x63, 0x6f, 0x6e, 0x74, 0x65, 0x78, 0x74, 0x52,
	0x03, 0x63, 0x74, 0x78, 0x22, 0xbe, 0x01, 0x0a, 0x0c, 0x63, 0x6f, 0x61, 0x5f, 0x72, 0x65, 0x73,
	0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x5b, 0x0a, 0x11, 0x63, 0x6f, 0x61, 0x5f, 0x72, 0x65, 0x73,
	0x70, 0x6f, 0x6e, 0x73, 0x65, 0x5f, 0x74, 0x79, 0x70, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0e,
	0x32, 0x2f, 0x2e, 0x61, 0x61, 0x61, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x73, 0x2e, 0x63, 0x6f,
	0x61, 0x5f, 0x72, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x2e, 0x63, 0x6f, 0x61, 0x5f, 0x72,
	0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x5f, 0x74, 0x79, 0x70, 0x65, 0x5f, 0x65, 0x6e, 0x75,
	0x6d, 0x52, 0x0f, 0x63, 0x6f, 0x61, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x54, 0x79,
	0x70, 0x65, 0x12, 0x25, 0x0a, 0x03, 0x63, 0x74, 0x78, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32,
	0x13, 0x2e, 0x61, 0x61, 0x61, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x73, 0x2e, 0x63, 0x6f, 0x6e,
	0x74, 0x65, 0x78, 0x74, 0x52, 0x03, 0x63, 0x74, 0x78, 0x22, 0x2a, 0x0a, 0x16, 0x63, 0x6f, 0x61,
	0x5f, 0x72, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x5f, 0x74, 0x79, 0x70, 0x65, 0x5f, 0x65,
	0x6e, 0x75, 0x6d, 0x12, 0x07, 0x0a, 0x03, 0x4e, 0x41, 0x4b, 0x10, 0x00, 0x12, 0x07, 0x0a, 0x03,
	0x41, 0x43, 0x4b, 0x10, 0x01, 0x32, 0x9b, 0x01, 0x0a, 0x0d, 0x61, 0x75, 0x74, 0x68, 0x6f, 0x72,
	0x69, 0x7a, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x40, 0x0a, 0x06, 0x63, 0x68, 0x61, 0x6e, 0x67,
	0x65, 0x12, 0x1a, 0x2e, 0x61, 0x61, 0x61, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x73, 0x2e, 0x63,
	0x68, 0x61, 0x6e, 0x67, 0x65, 0x5f, 0x72, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x18, 0x2e,
	0x61, 0x61, 0x61, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x73, 0x2e, 0x63, 0x6f, 0x61, 0x5f, 0x72,
	0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x48, 0x0a, 0x0a, 0x64, 0x69, 0x73,
	0x63, 0x6f, 0x6e, 0x6e, 0x65, 0x63, 0x74, 0x12, 0x1e, 0x2e, 0x61, 0x61, 0x61, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x73, 0x2e, 0x64, 0x69, 0x73, 0x63, 0x6f, 0x6e, 0x6e, 0x65, 0x63, 0x74, 0x5f,
	0x72, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x18, 0x2e, 0x61, 0x61, 0x61, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x73, 0x2e, 0x63, 0x6f, 0x61, 0x5f, 0x72, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73,
	0x65, 0x22, 0x00, 0x42, 0x27, 0x5a, 0x25, 0x6d, 0x61, 0x67, 0x6d, 0x61, 0x2f, 0x66, 0x65, 0x67,
	0x2f, 0x67, 0x61, 0x74, 0x65, 0x77, 0x61, 0x79, 0x2f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65,
	0x73, 0x2f, 0x61, 0x61, 0x61, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x73, 0x62, 0x06, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_authorization_proto_rawDescOnce sync.Once
	file_authorization_proto_rawDescData = file_authorization_proto_rawDesc
)

func file_authorization_proto_rawDescGZIP() []byte {
	file_authorization_proto_rawDescOnce.Do(func() {
		file_authorization_proto_rawDescData = protoimpl.X.CompressGZIP(file_authorization_proto_rawDescData)
	})
	return file_authorization_proto_rawDescData
}

var file_authorization_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_authorization_proto_msgTypes = make([]protoimpl.MessageInfo, 3)
var file_authorization_proto_goTypes = []interface{}{
	(CoaResponseCoaResponseTypeEnum)(0), // 0: aaa.protos.coa_response.coa_response_type_enum
	(*ChangeRequest)(nil),               // 1: aaa.protos.change_request
	(*DisconnectRequest)(nil),           // 2: aaa.protos.disconnect_request
	(*CoaResponse)(nil),                 // 3: aaa.protos.coa_response
	(*Context)(nil),                     // 4: aaa.protos.context
}
var file_authorization_proto_depIdxs = []int32{
	4, // 0: aaa.protos.change_request.ctx:type_name -> aaa.protos.context
	4, // 1: aaa.protos.disconnect_request.ctx:type_name -> aaa.protos.context
	0, // 2: aaa.protos.coa_response.coa_response_type:type_name -> aaa.protos.coa_response.coa_response_type_enum
	4, // 3: aaa.protos.coa_response.ctx:type_name -> aaa.protos.context
	1, // 4: aaa.protos.authorization.change:input_type -> aaa.protos.change_request
	2, // 5: aaa.protos.authorization.disconnect:input_type -> aaa.protos.disconnect_request
	3, // 6: aaa.protos.authorization.change:output_type -> aaa.protos.coa_response
	3, // 7: aaa.protos.authorization.disconnect:output_type -> aaa.protos.coa_response
	6, // [6:8] is the sub-list for method output_type
	4, // [4:6] is the sub-list for method input_type
	4, // [4:4] is the sub-list for extension type_name
	4, // [4:4] is the sub-list for extension extendee
	0, // [0:4] is the sub-list for field type_name
}

func init() { file_authorization_proto_init() }
func file_authorization_proto_init() {
	if File_authorization_proto != nil {
		return
	}
	file_context_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_authorization_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ChangeRequest); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_authorization_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DisconnectRequest); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_authorization_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CoaResponse); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_authorization_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   3,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_authorization_proto_goTypes,
		DependencyIndexes: file_authorization_proto_depIdxs,
		EnumInfos:         file_authorization_proto_enumTypes,
		MessageInfos:      file_authorization_proto_msgTypes,
	}.Build()
	File_authorization_proto = out.File
	file_authorization_proto_rawDesc = nil
	file_authorization_proto_goTypes = nil
	file_authorization_proto_depIdxs = nil
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConnInterface

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion6

// AuthorizationClient is the client API for Authorization service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type AuthorizationClient interface {
	Change(ctx context.Context, in *ChangeRequest, opts ...grpc.CallOption) (*CoaResponse, error)
	Disconnect(ctx context.Context, in *DisconnectRequest, opts ...grpc.CallOption) (*CoaResponse, error)
}

type authorizationClient struct {
	cc grpc.ClientConnInterface
}

func NewAuthorizationClient(cc grpc.ClientConnInterface) AuthorizationClient {
	return &authorizationClient{cc}
}

func (c *authorizationClient) Change(ctx context.Context, in *ChangeRequest, opts ...grpc.CallOption) (*CoaResponse, error) {
	out := new(CoaResponse)
	err := c.cc.Invoke(ctx, "/aaa.protos.authorization/change", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *authorizationClient) Disconnect(ctx context.Context, in *DisconnectRequest, opts ...grpc.CallOption) (*CoaResponse, error) {
	out := new(CoaResponse)
	err := c.cc.Invoke(ctx, "/aaa.protos.authorization/disconnect", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// AuthorizationServer is the server API for Authorization service.
type AuthorizationServer interface {
	Change(context.Context, *ChangeRequest) (*CoaResponse, error)
	Disconnect(context.Context, *DisconnectRequest) (*CoaResponse, error)
}

// UnimplementedAuthorizationServer can be embedded to have forward compatible implementations.
type UnimplementedAuthorizationServer struct {
}

func (*UnimplementedAuthorizationServer) Change(context.Context, *ChangeRequest) (*CoaResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Change not implemented")
}
func (*UnimplementedAuthorizationServer) Disconnect(context.Context, *DisconnectRequest) (*CoaResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Disconnect not implemented")
}

func RegisterAuthorizationServer(s *grpc.Server, srv AuthorizationServer) {
	s.RegisterService(&_Authorization_serviceDesc, srv)
}

func _Authorization_Change_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ChangeRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AuthorizationServer).Change(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/aaa.protos.authorization/Change",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AuthorizationServer).Change(ctx, req.(*ChangeRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Authorization_Disconnect_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DisconnectRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AuthorizationServer).Disconnect(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/aaa.protos.authorization/Disconnect",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AuthorizationServer).Disconnect(ctx, req.(*DisconnectRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _Authorization_serviceDesc = grpc.ServiceDesc{
	ServiceName: "aaa.protos.authorization",
	HandlerType: (*AuthorizationServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "change",
			Handler:    _Authorization_Change_Handler,
		},
		{
			MethodName: "disconnect",
			Handler:    _Authorization_Disconnect_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "authorization.proto",
}
