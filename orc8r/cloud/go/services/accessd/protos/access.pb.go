// Code generated by protoc-gen-go. DO NOT EDIT.
// source: access.proto

package protos

import (
	context "context"
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
	protos "magma/orc8r/cloud/go/protos"
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

// All permission definitions are used as bitmasks & should be 2**N
type AccessControl_Permission int32

const (
	AccessControl_NONE  AccessControl_Permission = 0
	AccessControl_READ  AccessControl_Permission = 1
	AccessControl_WRITE AccessControl_Permission = 2
)

var AccessControl_Permission_name = map[int32]string{
	0: "NONE",
	1: "READ",
	2: "WRITE",
}

var AccessControl_Permission_value = map[string]int32{
	"NONE":  0,
	"READ":  1,
	"WRITE": 2,
}

func (x AccessControl_Permission) String() string {
	return proto.EnumName(AccessControl_Permission_name, int32(x))
}

func (AccessControl_Permission) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_a098e900d2c3a6f2, []int{0, 0}
}

// Access Control Data Structures & Definitions
type AccessControl struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *AccessControl) Reset()         { *m = AccessControl{} }
func (m *AccessControl) String() string { return proto.CompactTextString(m) }
func (*AccessControl) ProtoMessage()    {}
func (*AccessControl) Descriptor() ([]byte, []int) {
	return fileDescriptor_a098e900d2c3a6f2, []int{0}
}

func (m *AccessControl) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_AccessControl.Unmarshal(m, b)
}
func (m *AccessControl) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_AccessControl.Marshal(b, m, deterministic)
}
func (m *AccessControl) XXX_Merge(src proto.Message) {
	xxx_messageInfo_AccessControl.Merge(m, src)
}
func (m *AccessControl) XXX_Size() int {
	return xxx_messageInfo_AccessControl.Size(m)
}
func (m *AccessControl) XXX_DiscardUnknown() {
	xxx_messageInfo_AccessControl.DiscardUnknown(m)
}

var xxx_messageInfo_AccessControl proto.InternalMessageInfo

// "Managed/monitored" entity record
type AccessControl_Entity struct {
	Id                   *protos.Identity         `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Permissions          AccessControl_Permission `protobuf:"varint,2,opt,name=permissions,proto3,enum=magma.orc8r.accessd.AccessControl_Permission" json:"permissions,omitempty"`
	XXX_NoUnkeyedLiteral struct{}                 `json:"-"`
	XXX_unrecognized     []byte                   `json:"-"`
	XXX_sizecache        int32                    `json:"-"`
}

func (m *AccessControl_Entity) Reset()         { *m = AccessControl_Entity{} }
func (m *AccessControl_Entity) String() string { return proto.CompactTextString(m) }
func (*AccessControl_Entity) ProtoMessage()    {}
func (*AccessControl_Entity) Descriptor() ([]byte, []int) {
	return fileDescriptor_a098e900d2c3a6f2, []int{0, 0}
}

func (m *AccessControl_Entity) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_AccessControl_Entity.Unmarshal(m, b)
}
func (m *AccessControl_Entity) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_AccessControl_Entity.Marshal(b, m, deterministic)
}
func (m *AccessControl_Entity) XXX_Merge(src proto.Message) {
	xxx_messageInfo_AccessControl_Entity.Merge(m, src)
}
func (m *AccessControl_Entity) XXX_Size() int {
	return xxx_messageInfo_AccessControl_Entity.Size(m)
}
func (m *AccessControl_Entity) XXX_DiscardUnknown() {
	xxx_messageInfo_AccessControl_Entity.DiscardUnknown(m)
}

var xxx_messageInfo_AccessControl_Entity proto.InternalMessageInfo

func (m *AccessControl_Entity) GetId() *protos.Identity {
	if m != nil {
		return m.Id
	}
	return nil
}

func (m *AccessControl_Entity) GetPermissions() AccessControl_Permission {
	if m != nil {
		return m.Permissions
	}
	return AccessControl_NONE
}

// Operator's Access Control List (map)
type AccessControl_List struct {
	Operator             *protos.Identity                 `protobuf:"bytes,1,opt,name=operator,proto3" json:"operator,omitempty"`
	Entities             map[string]*AccessControl_Entity `protobuf:"bytes,2,rep,name=entities,proto3" json:"entities,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
	XXX_NoUnkeyedLiteral struct{}                         `json:"-"`
	XXX_unrecognized     []byte                           `json:"-"`
	XXX_sizecache        int32                            `json:"-"`
}

func (m *AccessControl_List) Reset()         { *m = AccessControl_List{} }
func (m *AccessControl_List) String() string { return proto.CompactTextString(m) }
func (*AccessControl_List) ProtoMessage()    {}
func (*AccessControl_List) Descriptor() ([]byte, []int) {
	return fileDescriptor_a098e900d2c3a6f2, []int{0, 1}
}

func (m *AccessControl_List) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_AccessControl_List.Unmarshal(m, b)
}
func (m *AccessControl_List) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_AccessControl_List.Marshal(b, m, deterministic)
}
func (m *AccessControl_List) XXX_Merge(src proto.Message) {
	xxx_messageInfo_AccessControl_List.Merge(m, src)
}
func (m *AccessControl_List) XXX_Size() int {
	return xxx_messageInfo_AccessControl_List.Size(m)
}
func (m *AccessControl_List) XXX_DiscardUnknown() {
	xxx_messageInfo_AccessControl_List.DiscardUnknown(m)
}

var xxx_messageInfo_AccessControl_List proto.InternalMessageInfo

func (m *AccessControl_List) GetOperator() *protos.Identity {
	if m != nil {
		return m.Operator
	}
	return nil
}

func (m *AccessControl_List) GetEntities() map[string]*AccessControl_Entity {
	if m != nil {
		return m.Entities
	}
	return nil
}

// RPC Request/Responce used to 1) manage AND 2) check permissions
// 1. When Adding or Modifying permissions entities will represent managed
// entities Operator's permissions
// 2. When verifying permissions, entities will represent a list of
// Identities and their corresponding permissions requested by the operator
type AccessControl_ListRequest struct {
	Operator             *protos.Identity        `protobuf:"bytes,1,opt,name=operator,proto3" json:"operator,omitempty"`
	Entities             []*AccessControl_Entity `protobuf:"bytes,2,rep,name=entities,proto3" json:"entities,omitempty"`
	XXX_NoUnkeyedLiteral struct{}                `json:"-"`
	XXX_unrecognized     []byte                  `json:"-"`
	XXX_sizecache        int32                   `json:"-"`
}

func (m *AccessControl_ListRequest) Reset()         { *m = AccessControl_ListRequest{} }
func (m *AccessControl_ListRequest) String() string { return proto.CompactTextString(m) }
func (*AccessControl_ListRequest) ProtoMessage()    {}
func (*AccessControl_ListRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_a098e900d2c3a6f2, []int{0, 2}
}

func (m *AccessControl_ListRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_AccessControl_ListRequest.Unmarshal(m, b)
}
func (m *AccessControl_ListRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_AccessControl_ListRequest.Marshal(b, m, deterministic)
}
func (m *AccessControl_ListRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_AccessControl_ListRequest.Merge(m, src)
}
func (m *AccessControl_ListRequest) XXX_Size() int {
	return xxx_messageInfo_AccessControl_ListRequest.Size(m)
}
func (m *AccessControl_ListRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_AccessControl_ListRequest.DiscardUnknown(m)
}

var xxx_messageInfo_AccessControl_ListRequest proto.InternalMessageInfo

func (m *AccessControl_ListRequest) GetOperator() *protos.Identity {
	if m != nil {
		return m.Operator
	}
	return nil
}

func (m *AccessControl_ListRequest) GetEntities() []*AccessControl_Entity {
	if m != nil {
		return m.Entities
	}
	return nil
}

// RPC Request used to verify permissions for operator on a given entity
type AccessControl_PermissionsRequest struct {
	Operator             *protos.Identity `protobuf:"bytes,1,opt,name=operator,proto3" json:"operator,omitempty"`
	Entity               *protos.Identity `protobuf:"bytes,2,opt,name=entity,proto3" json:"entity,omitempty"`
	XXX_NoUnkeyedLiteral struct{}         `json:"-"`
	XXX_unrecognized     []byte           `json:"-"`
	XXX_sizecache        int32            `json:"-"`
}

func (m *AccessControl_PermissionsRequest) Reset()         { *m = AccessControl_PermissionsRequest{} }
func (m *AccessControl_PermissionsRequest) String() string { return proto.CompactTextString(m) }
func (*AccessControl_PermissionsRequest) ProtoMessage()    {}
func (*AccessControl_PermissionsRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_a098e900d2c3a6f2, []int{0, 3}
}

func (m *AccessControl_PermissionsRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_AccessControl_PermissionsRequest.Unmarshal(m, b)
}
func (m *AccessControl_PermissionsRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_AccessControl_PermissionsRequest.Marshal(b, m, deterministic)
}
func (m *AccessControl_PermissionsRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_AccessControl_PermissionsRequest.Merge(m, src)
}
func (m *AccessControl_PermissionsRequest) XXX_Size() int {
	return xxx_messageInfo_AccessControl_PermissionsRequest.Size(m)
}
func (m *AccessControl_PermissionsRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_AccessControl_PermissionsRequest.DiscardUnknown(m)
}

var xxx_messageInfo_AccessControl_PermissionsRequest proto.InternalMessageInfo

func (m *AccessControl_PermissionsRequest) GetOperator() *protos.Identity {
	if m != nil {
		return m.Operator
	}
	return nil
}

func (m *AccessControl_PermissionsRequest) GetEntity() *protos.Identity {
	if m != nil {
		return m.Entity
	}
	return nil
}

type AccessControl_Lists struct {
	Acls                 []*AccessControl_List `protobuf:"bytes,1,rep,name=acls,proto3" json:"acls,omitempty"`
	XXX_NoUnkeyedLiteral struct{}              `json:"-"`
	XXX_unrecognized     []byte                `json:"-"`
	XXX_sizecache        int32                 `json:"-"`
}

func (m *AccessControl_Lists) Reset()         { *m = AccessControl_Lists{} }
func (m *AccessControl_Lists) String() string { return proto.CompactTextString(m) }
func (*AccessControl_Lists) ProtoMessage()    {}
func (*AccessControl_Lists) Descriptor() ([]byte, []int) {
	return fileDescriptor_a098e900d2c3a6f2, []int{0, 4}
}

func (m *AccessControl_Lists) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_AccessControl_Lists.Unmarshal(m, b)
}
func (m *AccessControl_Lists) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_AccessControl_Lists.Marshal(b, m, deterministic)
}
func (m *AccessControl_Lists) XXX_Merge(src proto.Message) {
	xxx_messageInfo_AccessControl_Lists.Merge(m, src)
}
func (m *AccessControl_Lists) XXX_Size() int {
	return xxx_messageInfo_AccessControl_Lists.Size(m)
}
func (m *AccessControl_Lists) XXX_DiscardUnknown() {
	xxx_messageInfo_AccessControl_Lists.DiscardUnknown(m)
}

var xxx_messageInfo_AccessControl_Lists proto.InternalMessageInfo

func (m *AccessControl_Lists) GetAcls() []*AccessControl_List {
	if m != nil {
		return m.Acls
	}
	return nil
}

func init() {
	proto.RegisterEnum("magma.orc8r.accessd.AccessControl_Permission", AccessControl_Permission_name, AccessControl_Permission_value)
	proto.RegisterType((*AccessControl)(nil), "magma.orc8r.accessd.AccessControl")
	proto.RegisterType((*AccessControl_Entity)(nil), "magma.orc8r.accessd.AccessControl.Entity")
	proto.RegisterType((*AccessControl_List)(nil), "magma.orc8r.accessd.AccessControl.List")
	proto.RegisterMapType((map[string]*AccessControl_Entity)(nil), "magma.orc8r.accessd.AccessControl.List.EntitiesEntry")
	proto.RegisterType((*AccessControl_ListRequest)(nil), "magma.orc8r.accessd.AccessControl.ListRequest")
	proto.RegisterType((*AccessControl_PermissionsRequest)(nil), "magma.orc8r.accessd.AccessControl.PermissionsRequest")
	proto.RegisterType((*AccessControl_Lists)(nil), "magma.orc8r.accessd.AccessControl.Lists")
}

func init() { proto.RegisterFile("access.proto", fileDescriptor_a098e900d2c3a6f2) }

var fileDescriptor_a098e900d2c3a6f2 = []byte{
	// 531 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xac, 0x55, 0xe1, 0x6b, 0xd3, 0x40,
	0x14, 0x4f, 0xba, 0xb5, 0x74, 0x2f, 0x6b, 0x89, 0x4f, 0x85, 0x7a, 0x7e, 0x19, 0x01, 0xb1, 0x22,
	0xcb, 0xb0, 0x32, 0x18, 0x53, 0xd0, 0xda, 0x06, 0x19, 0xd4, 0x55, 0x4f, 0xe7, 0x60, 0xdf, 0x62,
	0x72, 0xce, 0xb0, 0x26, 0x57, 0x73, 0xb7, 0x41, 0xbf, 0xf9, 0xcd, 0x7f, 0xd8, 0x8f, 0x22, 0x92,
	0xbb, 0xac, 0x4d, 0x58, 0x37, 0x52, 0xdd, 0xa7, 0x5c, 0xee, 0xfd, 0x7e, 0xbf, 0xf7, 0x7b, 0xef,
	0x5d, 0x2e, 0xb0, 0xe9, 0x07, 0x01, 0x13, 0xc2, 0x9d, 0xa6, 0x5c, 0x72, 0xbc, 0x1b, 0xfb, 0xa7,
	0xb1, 0xef, 0xf2, 0x34, 0xd8, 0x4b, 0x5d, 0x1d, 0x09, 0xc9, 0x03, 0xf5, 0xba, 0xa3, 0x10, 0x62,
	0x27, 0xe0, 0x71, 0xcc, 0x13, 0x8d, 0x27, 0x0f, 0x4b, 0xa1, 0x28, 0x64, 0x89, 0x8c, 0xe4, 0x4c,
	0x07, 0x9d, 0x3f, 0x75, 0x68, 0xf5, 0x95, 0xc6, 0x80, 0x27, 0x32, 0xe5, 0x13, 0xf2, 0xc3, 0x84,
	0x86, 0xa7, 0x20, 0xf8, 0x08, 0x6a, 0x51, 0xd8, 0x31, 0xb7, 0xcc, 0xae, 0xd5, 0xbb, 0xef, 0x16,
	0xd3, 0x1e, 0xe4, 0x2a, 0xb4, 0x16, 0x85, 0x38, 0x06, 0x6b, 0xca, 0xd2, 0x38, 0x12, 0x22, 0xe2,
	0x89, 0xe8, 0xd4, 0xb6, 0xcc, 0x6e, 0xbb, 0xb7, 0xed, 0x2e, 0xb1, 0xe9, 0x96, 0x52, 0xb9, 0xef,
	0xe7, 0x2c, 0x5a, 0x54, 0x20, 0xbf, 0x4c, 0x58, 0x1f, 0x45, 0x42, 0xe2, 0x33, 0x68, 0xf2, 0x29,
	0x4b, 0x7d, 0xc9, 0xd3, 0x9b, 0x6d, 0xcc, 0x61, 0xf8, 0x01, 0x9a, 0x6a, 0x2f, 0x62, 0x99, 0x93,
	0xb5, 0xae, 0xd5, 0xdb, 0xad, 0xe0, 0x24, 0xcb, 0xe6, 0x7a, 0x39, 0xcf, 0x4b, 0x64, 0x3a, 0xa3,
	0x73, 0x19, 0xf2, 0x15, 0x5a, 0xa5, 0x10, 0xda, 0xb0, 0x76, 0xc6, 0x66, 0xca, 0xd1, 0x06, 0xcd,
	0x96, 0xf8, 0x0a, 0xea, 0x17, 0xfe, 0xe4, 0x9c, 0xa9, 0xe2, 0xad, 0xde, 0x93, 0x0a, 0x29, 0x75,
	0x8f, 0xa9, 0xe6, 0xed, 0xd7, 0xf6, 0x4c, 0xf2, 0xd3, 0x04, 0x2b, 0x33, 0x42, 0xd9, 0xf7, 0x73,
	0xf6, 0x6f, 0xd5, 0x7b, 0x57, 0xaa, 0x5f, 0xc1, 0xca, 0xa2, 0xe2, 0x0b, 0xc0, 0xc5, 0x6c, 0xc4,
	0x7f, 0xf8, 0xd9, 0x86, 0x86, 0xde, 0xcb, 0x1b, 0x73, 0x0d, 0x21, 0x07, 0x91, 0x21, 0xd4, 0xb3,
	0x06, 0x08, 0x7c, 0x01, 0xeb, 0x7e, 0x30, 0x11, 0x1d, 0x53, 0xd5, 0xf0, 0xb8, 0xe2, 0x04, 0xa9,
	0x22, 0x39, 0x4f, 0x01, 0x16, 0xee, 0xb1, 0x09, 0xeb, 0x87, 0xe3, 0x43, 0xcf, 0x36, 0xb2, 0x15,
	0xf5, 0xfa, 0x43, 0xdb, 0xc4, 0x0d, 0xa8, 0x1f, 0xd3, 0x83, 0x4f, 0x9e, 0x5d, 0xeb, 0xfd, 0xae,
	0xc3, 0xbd, 0x92, 0xd2, 0x3b, 0x3f, 0xf1, 0x4f, 0x59, 0x8a, 0x14, 0xac, 0x8f, 0x4c, 0x8e, 0x2f,
	0x2b, 0x71, 0xab, 0x7a, 0xd0, 0xcd, 0x22, 0x77, 0x4a, 0xf8, 0xcf, 0x3c, 0x0a, 0x1d, 0x03, 0x8f,
	0xa0, 0x7d, 0x34, 0x0d, 0x7d, 0xc9, 0x6e, 0x57, 0xf6, 0x25, 0xb4, 0x87, 0x6c, 0xc2, 0x0a, 0xb2,
	0xcb, 0xfb, 0xbc, 0x9c, 0x4d, 0xa1, 0xfd, 0x76, 0x51, 0x68, 0x7f, 0x30, 0xba, 0x8e, 0x5d, 0x75,
	0x0c, 0x8e, 0x81, 0x27, 0x60, 0x17, 0x34, 0x45, 0x7f, 0x30, 0x12, 0x48, 0x96, 0xaa, 0x2a, 0x06,
	0xe9, 0x56, 0x94, 0x16, 0x8e, 0x81, 0x52, 0xf9, 0x2d, 0x9c, 0x4f, 0xdc, 0x5d, 0xe9, 0xae, 0xb9,
	0x3c, 0xcf, 0xa4, 0xfa, 0xa7, 0xe1, 0x18, 0x78, 0x0c, 0xf6, 0xe0, 0x1b, 0x0b, 0xce, 0x8a, 0x79,
	0x6f, 0x65, 0x78, 0xaf, 0xa1, 0x95, 0x61, 0xe6, 0xbd, 0xc2, 0xab, 0x28, 0x72, 0x43, 0xeb, 0x1c,
	0x03, 0xf7, 0x61, 0x53, 0x8f, 0x3f, 0xbf, 0xb6, 0x57, 0x18, 0xfe, 0x9b, 0xe6, 0x49, 0x43, 0xff,
	0x18, 0xbe, 0xe8, 0xe7, 0xf3, 0xbf, 0x01, 0x00, 0x00, 0xff, 0xff, 0x3e, 0x04, 0xce, 0x33, 0x6d,
	0x06, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// AccessControlManagerClient is the client API for AccessControlManager service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type AccessControlManagerClient interface {
	// Overwrites Permissions for operator Identity to manage others
	// Request includes ACL to set for the Operator
	// If the Operator doesn't exist - creates a new operator with the given ACL
	SetOperator(ctx context.Context, in *AccessControl_ListRequest, opts ...grpc.CallOption) (*protos.Void, error)
	// Adds Permissions for one Identity to manage others
	// Request includes ACL to add (append to the existing ACL) for the Operator
	UpdateOperator(ctx context.Context, in *AccessControl_ListRequest, opts ...grpc.CallOption) (*protos.Void, error)
	// Removes all operator's permissions (the entire operator's ACL)
	DeleteOperator(ctx context.Context, in *protos.Identity, opts ...grpc.CallOption) (*protos.Void, error)
	// Returns the managing Identity's permissions list
	GetOperatorACL(ctx context.Context, in *protos.Identity, opts ...grpc.CallOption) (*AccessControl_List, error)
	// Returns the managing Identity's permissions list
	GetOperatorsACLs(ctx context.Context, in *protos.Identity_List, opts ...grpc.CallOption) (*AccessControl_Lists, error)
	// Returns the managing Identity's permissions for a given entity
	// NOTE: Takes into account wildcards for the entity's type in the ACL
	GetPermissions(ctx context.Context, in *AccessControl_PermissionsRequest, opts ...grpc.CallOption) (*AccessControl_Entity, error)
	// CheckPermissions verifies Operator permissions for a list of given
	// Identities. AccessControl.ListRequest.entities is a list of
	// Identities and their corresponding permissions requested by the operator
	// CheckPermissions will return success only if all requested permissions
	// are satisfied (AND logic)
	// Intended to be used for multi-Identity requests such as Network Identity
	// AND REST API Identity, etc.
	CheckPermissions(ctx context.Context, in *AccessControl_ListRequest, opts ...grpc.CallOption) (*protos.Void, error)
	// Lists all globally registered operators on the cloud
	ListOperators(ctx context.Context, in *protos.Void, opts ...grpc.CallOption) (*protos.Identity_List, error)
	// Cleanup a given entity from all Operators' ACLs
	DeleteEntity(ctx context.Context, in *protos.Identity, opts ...grpc.CallOption) (*protos.Void, error)
}

type accessControlManagerClient struct {
	cc *grpc.ClientConn
}

func NewAccessControlManagerClient(cc *grpc.ClientConn) AccessControlManagerClient {
	return &accessControlManagerClient{cc}
}

func (c *accessControlManagerClient) SetOperator(ctx context.Context, in *AccessControl_ListRequest, opts ...grpc.CallOption) (*protos.Void, error) {
	out := new(protos.Void)
	err := c.cc.Invoke(ctx, "/magma.orc8r.accessd.AccessControlManager/SetOperator", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *accessControlManagerClient) UpdateOperator(ctx context.Context, in *AccessControl_ListRequest, opts ...grpc.CallOption) (*protos.Void, error) {
	out := new(protos.Void)
	err := c.cc.Invoke(ctx, "/magma.orc8r.accessd.AccessControlManager/UpdateOperator", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *accessControlManagerClient) DeleteOperator(ctx context.Context, in *protos.Identity, opts ...grpc.CallOption) (*protos.Void, error) {
	out := new(protos.Void)
	err := c.cc.Invoke(ctx, "/magma.orc8r.accessd.AccessControlManager/DeleteOperator", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *accessControlManagerClient) GetOperatorACL(ctx context.Context, in *protos.Identity, opts ...grpc.CallOption) (*AccessControl_List, error) {
	out := new(AccessControl_List)
	err := c.cc.Invoke(ctx, "/magma.orc8r.accessd.AccessControlManager/GetOperatorACL", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *accessControlManagerClient) GetOperatorsACLs(ctx context.Context, in *protos.Identity_List, opts ...grpc.CallOption) (*AccessControl_Lists, error) {
	out := new(AccessControl_Lists)
	err := c.cc.Invoke(ctx, "/magma.orc8r.accessd.AccessControlManager/GetOperatorsACLs", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *accessControlManagerClient) GetPermissions(ctx context.Context, in *AccessControl_PermissionsRequest, opts ...grpc.CallOption) (*AccessControl_Entity, error) {
	out := new(AccessControl_Entity)
	err := c.cc.Invoke(ctx, "/magma.orc8r.accessd.AccessControlManager/GetPermissions", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *accessControlManagerClient) CheckPermissions(ctx context.Context, in *AccessControl_ListRequest, opts ...grpc.CallOption) (*protos.Void, error) {
	out := new(protos.Void)
	err := c.cc.Invoke(ctx, "/magma.orc8r.accessd.AccessControlManager/CheckPermissions", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *accessControlManagerClient) ListOperators(ctx context.Context, in *protos.Void, opts ...grpc.CallOption) (*protos.Identity_List, error) {
	out := new(protos.Identity_List)
	err := c.cc.Invoke(ctx, "/magma.orc8r.accessd.AccessControlManager/ListOperators", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *accessControlManagerClient) DeleteEntity(ctx context.Context, in *protos.Identity, opts ...grpc.CallOption) (*protos.Void, error) {
	out := new(protos.Void)
	err := c.cc.Invoke(ctx, "/magma.orc8r.accessd.AccessControlManager/DeleteEntity", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// AccessControlManagerServer is the server API for AccessControlManager service.
type AccessControlManagerServer interface {
	// Overwrites Permissions for operator Identity to manage others
	// Request includes ACL to set for the Operator
	// If the Operator doesn't exist - creates a new operator with the given ACL
	SetOperator(context.Context, *AccessControl_ListRequest) (*protos.Void, error)
	// Adds Permissions for one Identity to manage others
	// Request includes ACL to add (append to the existing ACL) for the Operator
	UpdateOperator(context.Context, *AccessControl_ListRequest) (*protos.Void, error)
	// Removes all operator's permissions (the entire operator's ACL)
	DeleteOperator(context.Context, *protos.Identity) (*protos.Void, error)
	// Returns the managing Identity's permissions list
	GetOperatorACL(context.Context, *protos.Identity) (*AccessControl_List, error)
	// Returns the managing Identity's permissions list
	GetOperatorsACLs(context.Context, *protos.Identity_List) (*AccessControl_Lists, error)
	// Returns the managing Identity's permissions for a given entity
	// NOTE: Takes into account wildcards for the entity's type in the ACL
	GetPermissions(context.Context, *AccessControl_PermissionsRequest) (*AccessControl_Entity, error)
	// CheckPermissions verifies Operator permissions for a list of given
	// Identities. AccessControl.ListRequest.entities is a list of
	// Identities and their corresponding permissions requested by the operator
	// CheckPermissions will return success only if all requested permissions
	// are satisfied (AND logic)
	// Intended to be used for multi-Identity requests such as Network Identity
	// AND REST API Identity, etc.
	CheckPermissions(context.Context, *AccessControl_ListRequest) (*protos.Void, error)
	// Lists all globally registered operators on the cloud
	ListOperators(context.Context, *protos.Void) (*protos.Identity_List, error)
	// Cleanup a given entity from all Operators' ACLs
	DeleteEntity(context.Context, *protos.Identity) (*protos.Void, error)
}

// UnimplementedAccessControlManagerServer can be embedded to have forward compatible implementations.
type UnimplementedAccessControlManagerServer struct {
}

func (*UnimplementedAccessControlManagerServer) SetOperator(ctx context.Context, req *AccessControl_ListRequest) (*protos.Void, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SetOperator not implemented")
}
func (*UnimplementedAccessControlManagerServer) UpdateOperator(ctx context.Context, req *AccessControl_ListRequest) (*protos.Void, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateOperator not implemented")
}
func (*UnimplementedAccessControlManagerServer) DeleteOperator(ctx context.Context, req *protos.Identity) (*protos.Void, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteOperator not implemented")
}
func (*UnimplementedAccessControlManagerServer) GetOperatorACL(ctx context.Context, req *protos.Identity) (*AccessControl_List, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetOperatorACL not implemented")
}
func (*UnimplementedAccessControlManagerServer) GetOperatorsACLs(ctx context.Context, req *protos.Identity_List) (*AccessControl_Lists, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetOperatorsACLs not implemented")
}
func (*UnimplementedAccessControlManagerServer) GetPermissions(ctx context.Context, req *AccessControl_PermissionsRequest) (*AccessControl_Entity, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetPermissions not implemented")
}
func (*UnimplementedAccessControlManagerServer) CheckPermissions(ctx context.Context, req *AccessControl_ListRequest) (*protos.Void, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CheckPermissions not implemented")
}
func (*UnimplementedAccessControlManagerServer) ListOperators(ctx context.Context, req *protos.Void) (*protos.Identity_List, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ListOperators not implemented")
}
func (*UnimplementedAccessControlManagerServer) DeleteEntity(ctx context.Context, req *protos.Identity) (*protos.Void, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteEntity not implemented")
}

func RegisterAccessControlManagerServer(s *grpc.Server, srv AccessControlManagerServer) {
	s.RegisterService(&_AccessControlManager_serviceDesc, srv)
}

func _AccessControlManager_SetOperator_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(AccessControl_ListRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AccessControlManagerServer).SetOperator(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/magma.orc8r.accessd.AccessControlManager/SetOperator",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AccessControlManagerServer).SetOperator(ctx, req.(*AccessControl_ListRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _AccessControlManager_UpdateOperator_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(AccessControl_ListRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AccessControlManagerServer).UpdateOperator(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/magma.orc8r.accessd.AccessControlManager/UpdateOperator",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AccessControlManagerServer).UpdateOperator(ctx, req.(*AccessControl_ListRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _AccessControlManager_DeleteOperator_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(protos.Identity)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AccessControlManagerServer).DeleteOperator(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/magma.orc8r.accessd.AccessControlManager/DeleteOperator",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AccessControlManagerServer).DeleteOperator(ctx, req.(*protos.Identity))
	}
	return interceptor(ctx, in, info, handler)
}

func _AccessControlManager_GetOperatorACL_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(protos.Identity)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AccessControlManagerServer).GetOperatorACL(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/magma.orc8r.accessd.AccessControlManager/GetOperatorACL",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AccessControlManagerServer).GetOperatorACL(ctx, req.(*protos.Identity))
	}
	return interceptor(ctx, in, info, handler)
}

func _AccessControlManager_GetOperatorsACLs_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(protos.Identity_List)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AccessControlManagerServer).GetOperatorsACLs(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/magma.orc8r.accessd.AccessControlManager/GetOperatorsACLs",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AccessControlManagerServer).GetOperatorsACLs(ctx, req.(*protos.Identity_List))
	}
	return interceptor(ctx, in, info, handler)
}

func _AccessControlManager_GetPermissions_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(AccessControl_PermissionsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AccessControlManagerServer).GetPermissions(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/magma.orc8r.accessd.AccessControlManager/GetPermissions",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AccessControlManagerServer).GetPermissions(ctx, req.(*AccessControl_PermissionsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _AccessControlManager_CheckPermissions_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(AccessControl_ListRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AccessControlManagerServer).CheckPermissions(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/magma.orc8r.accessd.AccessControlManager/CheckPermissions",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AccessControlManagerServer).CheckPermissions(ctx, req.(*AccessControl_ListRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _AccessControlManager_ListOperators_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(protos.Void)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AccessControlManagerServer).ListOperators(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/magma.orc8r.accessd.AccessControlManager/ListOperators",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AccessControlManagerServer).ListOperators(ctx, req.(*protos.Void))
	}
	return interceptor(ctx, in, info, handler)
}

func _AccessControlManager_DeleteEntity_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(protos.Identity)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AccessControlManagerServer).DeleteEntity(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/magma.orc8r.accessd.AccessControlManager/DeleteEntity",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AccessControlManagerServer).DeleteEntity(ctx, req.(*protos.Identity))
	}
	return interceptor(ctx, in, info, handler)
}

var _AccessControlManager_serviceDesc = grpc.ServiceDesc{
	ServiceName: "magma.orc8r.accessd.AccessControlManager",
	HandlerType: (*AccessControlManagerServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "SetOperator",
			Handler:    _AccessControlManager_SetOperator_Handler,
		},
		{
			MethodName: "UpdateOperator",
			Handler:    _AccessControlManager_UpdateOperator_Handler,
		},
		{
			MethodName: "DeleteOperator",
			Handler:    _AccessControlManager_DeleteOperator_Handler,
		},
		{
			MethodName: "GetOperatorACL",
			Handler:    _AccessControlManager_GetOperatorACL_Handler,
		},
		{
			MethodName: "GetOperatorsACLs",
			Handler:    _AccessControlManager_GetOperatorsACLs_Handler,
		},
		{
			MethodName: "GetPermissions",
			Handler:    _AccessControlManager_GetPermissions_Handler,
		},
		{
			MethodName: "CheckPermissions",
			Handler:    _AccessControlManager_CheckPermissions_Handler,
		},
		{
			MethodName: "ListOperators",
			Handler:    _AccessControlManager_ListOperators_Handler,
		},
		{
			MethodName: "DeleteEntity",
			Handler:    _AccessControlManager_DeleteEntity_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "access.proto",
}
