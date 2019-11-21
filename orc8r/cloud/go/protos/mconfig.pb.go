// Code generated by protoc-gen-go. DO NOT EDIT.
// source: orc8r/protos/mconfig.proto

package protos

import (
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	any "github.com/golang/protobuf/ptypes/any"
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
// GatewayConfigs structure is a container for all Access Gateway's (AG) Cloud
// Managed Configs (CMC). Each and every field of GatewayConfigs represents
// one AG service config
// --------------------------------------------------------------------------
// NOTE: a service config field name (control_proxy, enodebd, etc.) must match
//       the corresponding gateway service's name exactly
type GatewayConfigs struct {
	ConfigsByKey         map[string]*any.Any     `protobuf:"bytes,10,rep,name=configs_by_key,json=configsByKey,proto3" json:"configs_by_key,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
	Metadata             *GatewayConfigsMetadata `protobuf:"bytes,11,opt,name=metadata,proto3" json:"metadata,omitempty"`
	XXX_NoUnkeyedLiteral struct{}                `json:"-"`
	XXX_unrecognized     []byte                  `json:"-"`
	XXX_sizecache        int32                   `json:"-"`
}

func (m *GatewayConfigs) Reset()         { *m = GatewayConfigs{} }
func (m *GatewayConfigs) String() string { return proto.CompactTextString(m) }
func (*GatewayConfigs) ProtoMessage()    {}
func (*GatewayConfigs) Descriptor() ([]byte, []int) {
	return fileDescriptor_8d300a2840d2449e, []int{0}
}

func (m *GatewayConfigs) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GatewayConfigs.Unmarshal(m, b)
}
func (m *GatewayConfigs) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GatewayConfigs.Marshal(b, m, deterministic)
}
func (m *GatewayConfigs) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GatewayConfigs.Merge(m, src)
}
func (m *GatewayConfigs) XXX_Size() int {
	return xxx_messageInfo_GatewayConfigs.Size(m)
}
func (m *GatewayConfigs) XXX_DiscardUnknown() {
	xxx_messageInfo_GatewayConfigs.DiscardUnknown(m)
}

var xxx_messageInfo_GatewayConfigs proto.InternalMessageInfo

func (m *GatewayConfigs) GetConfigsByKey() map[string]*any.Any {
	if m != nil {
		return m.ConfigsByKey
	}
	return nil
}

func (m *GatewayConfigs) GetMetadata() *GatewayConfigsMetadata {
	if m != nil {
		return m.Metadata
	}
	return nil
}

// Metadata about the configs.
type GatewayConfigsMetadata struct {
	// Unix timestamp of Cloud at the time of config generation.
	CreatedAt            uint64   `protobuf:"varint,11,opt,name=created_at,json=createdAt,proto3" json:"created_at,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *GatewayConfigsMetadata) Reset()         { *m = GatewayConfigsMetadata{} }
func (m *GatewayConfigsMetadata) String() string { return proto.CompactTextString(m) }
func (*GatewayConfigsMetadata) ProtoMessage()    {}
func (*GatewayConfigsMetadata) Descriptor() ([]byte, []int) {
	return fileDescriptor_8d300a2840d2449e, []int{1}
}

func (m *GatewayConfigsMetadata) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GatewayConfigsMetadata.Unmarshal(m, b)
}
func (m *GatewayConfigsMetadata) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GatewayConfigsMetadata.Marshal(b, m, deterministic)
}
func (m *GatewayConfigsMetadata) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GatewayConfigsMetadata.Merge(m, src)
}
func (m *GatewayConfigsMetadata) XXX_Size() int {
	return xxx_messageInfo_GatewayConfigsMetadata.Size(m)
}
func (m *GatewayConfigsMetadata) XXX_DiscardUnknown() {
	xxx_messageInfo_GatewayConfigsMetadata.DiscardUnknown(m)
}

var xxx_messageInfo_GatewayConfigsMetadata proto.InternalMessageInfo

func (m *GatewayConfigsMetadata) GetCreatedAt() uint64 {
	if m != nil {
		return m.CreatedAt
	}
	return 0
}

// Wraps a gateway config and a stream offset that the config was computed
// from
type OffsetGatewayConfigs struct {
	Configs              *GatewayConfigs `protobuf:"bytes,1,opt,name=configs,proto3" json:"configs,omitempty"`
	Offset               int64           `protobuf:"varint,2,opt,name=offset,proto3" json:"offset,omitempty"`
	XXX_NoUnkeyedLiteral struct{}        `json:"-"`
	XXX_unrecognized     []byte          `json:"-"`
	XXX_sizecache        int32           `json:"-"`
}

func (m *OffsetGatewayConfigs) Reset()         { *m = OffsetGatewayConfigs{} }
func (m *OffsetGatewayConfigs) String() string { return proto.CompactTextString(m) }
func (*OffsetGatewayConfigs) ProtoMessage()    {}
func (*OffsetGatewayConfigs) Descriptor() ([]byte, []int) {
	return fileDescriptor_8d300a2840d2449e, []int{2}
}

func (m *OffsetGatewayConfigs) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_OffsetGatewayConfigs.Unmarshal(m, b)
}
func (m *OffsetGatewayConfigs) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_OffsetGatewayConfigs.Marshal(b, m, deterministic)
}
func (m *OffsetGatewayConfigs) XXX_Merge(src proto.Message) {
	xxx_messageInfo_OffsetGatewayConfigs.Merge(m, src)
}
func (m *OffsetGatewayConfigs) XXX_Size() int {
	return xxx_messageInfo_OffsetGatewayConfigs.Size(m)
}
func (m *OffsetGatewayConfigs) XXX_DiscardUnknown() {
	xxx_messageInfo_OffsetGatewayConfigs.DiscardUnknown(m)
}

var xxx_messageInfo_OffsetGatewayConfigs proto.InternalMessageInfo

func (m *OffsetGatewayConfigs) GetConfigs() *GatewayConfigs {
	if m != nil {
		return m.Configs
	}
	return nil
}

func (m *OffsetGatewayConfigs) GetOffset() int64 {
	if m != nil {
		return m.Offset
	}
	return 0
}

// Stream request passed as extra args to the streaming mconfig streamer policy.
// Contains a single field, the offset of the mconfig currently stored on
// the device.
type MconfigStreamRequest struct {
	Offset               int64    `protobuf:"varint,1,opt,name=offset,proto3" json:"offset,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *MconfigStreamRequest) Reset()         { *m = MconfigStreamRequest{} }
func (m *MconfigStreamRequest) String() string { return proto.CompactTextString(m) }
func (*MconfigStreamRequest) ProtoMessage()    {}
func (*MconfigStreamRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_8d300a2840d2449e, []int{3}
}

func (m *MconfigStreamRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_MconfigStreamRequest.Unmarshal(m, b)
}
func (m *MconfigStreamRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_MconfigStreamRequest.Marshal(b, m, deterministic)
}
func (m *MconfigStreamRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MconfigStreamRequest.Merge(m, src)
}
func (m *MconfigStreamRequest) XXX_Size() int {
	return xxx_messageInfo_MconfigStreamRequest.Size(m)
}
func (m *MconfigStreamRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_MconfigStreamRequest.DiscardUnknown(m)
}

var xxx_messageInfo_MconfigStreamRequest proto.InternalMessageInfo

func (m *MconfigStreamRequest) GetOffset() int64 {
	if m != nil {
		return m.Offset
	}
	return 0
}

func init() {
	proto.RegisterType((*GatewayConfigs)(nil), "magma.orc8r.GatewayConfigs")
	proto.RegisterMapType((map[string]*any.Any)(nil), "magma.orc8r.GatewayConfigs.ConfigsByKeyEntry")
	proto.RegisterType((*GatewayConfigsMetadata)(nil), "magma.orc8r.GatewayConfigsMetadata")
	proto.RegisterType((*OffsetGatewayConfigs)(nil), "magma.orc8r.OffsetGatewayConfigs")
	proto.RegisterType((*MconfigStreamRequest)(nil), "magma.orc8r.MconfigStreamRequest")
}

func init() { proto.RegisterFile("orc8r/protos/mconfig.proto", fileDescriptor_8d300a2840d2449e) }

var fileDescriptor_8d300a2840d2449e = []byte{
	// 327 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x7c, 0x92, 0xcf, 0x4b, 0xc3, 0x30,
	0x14, 0xc7, 0xe9, 0xa6, 0xd3, 0xbd, 0xca, 0xd0, 0x30, 0xc6, 0xdc, 0x18, 0x8c, 0x7a, 0x19, 0x82,
	0x29, 0x4c, 0xc4, 0xe1, 0x45, 0x36, 0x11, 0x0f, 0x32, 0x84, 0x0c, 0x2f, 0x5e, 0xc6, 0x5b, 0xf7,
	0x5a, 0xc4, 0xb5, 0xd1, 0x36, 0x55, 0xf2, 0xc7, 0x0b, 0xb2, 0xa4, 0x93, 0xfa, 0x83, 0x9d, 0xd2,
	0x84, 0xcf, 0xfb, 0xfe, 0x28, 0x0f, 0x3a, 0x32, 0x0d, 0x46, 0xa9, 0xff, 0x9a, 0x4a, 0x25, 0x33,
	0x3f, 0x0e, 0x64, 0x12, 0x3e, 0x47, 0xdc, 0x5c, 0x99, 0x1b, 0x63, 0x14, 0x23, 0x37, 0x44, 0xe7,
	0x38, 0x92, 0x32, 0x5a, 0x91, 0x25, 0x17, 0x79, 0xe8, 0x63, 0xa2, 0x2d, 0xe7, 0x7d, 0x3a, 0xd0,
	0xb8, 0x43, 0x45, 0x1f, 0xa8, 0x6f, 0xcc, 0x7c, 0xc6, 0x66, 0xd0, 0xb0, 0x52, 0xd9, 0x7c, 0xa1,
	0xe7, 0x2f, 0xa4, 0xdb, 0xd0, 0xaf, 0x0e, 0xdc, 0xe1, 0x19, 0x2f, 0x69, 0xf2, 0x9f, 0x43, 0xbc,
	0x38, 0x27, 0xfa, 0x9e, 0xf4, 0x6d, 0xa2, 0x52, 0x2d, 0x0e, 0x82, 0xd2, 0x13, 0xbb, 0x86, 0xfd,
	0x98, 0x14, 0x2e, 0x51, 0x61, 0xdb, 0xed, 0x3b, 0x03, 0x77, 0x78, 0xb2, 0x45, 0x6e, 0x5a, 0xa0,
	0xe2, 0x7b, 0xa8, 0xf3, 0x08, 0x47, 0x7f, 0x3c, 0xd8, 0x21, 0x54, 0xd7, 0xf9, 0x9c, 0xbe, 0x33,
	0xa8, 0x8b, 0xf5, 0x27, 0x3b, 0x85, 0xdd, 0x77, 0x5c, 0xe5, 0xd4, 0xae, 0x18, 0x93, 0x26, 0xb7,
	0xd5, 0xf9, 0xa6, 0x3a, 0x1f, 0x27, 0x5a, 0x58, 0xe4, 0xaa, 0x32, 0x72, 0xbc, 0x4b, 0x68, 0xfd,
	0x6f, 0xcd, 0x7a, 0x00, 0x41, 0x4a, 0xa8, 0x68, 0x39, 0x47, 0x65, 0x32, 0xef, 0x88, 0x7a, 0xf1,
	0x32, 0x56, 0x1e, 0x41, 0xf3, 0x21, 0x0c, 0x33, 0x52, 0xbf, 0xfe, 0xde, 0x05, 0xec, 0x15, 0xc5,
	0x4d, 0x2c, 0x77, 0xd8, 0xdd, 0xd2, 0x53, 0x6c, 0x58, 0xd6, 0x82, 0x9a, 0x34, 0x72, 0x26, 0x78,
	0x55, 0x14, 0x37, 0x8f, 0x43, 0x73, 0x6a, 0x99, 0x99, 0x4a, 0x09, 0x63, 0x41, 0x6f, 0x39, 0x65,
	0xaa, 0xc4, 0x3b, 0x65, 0x7e, 0xd2, 0x7b, 0xea, 0x1a, 0x3b, 0xdf, 0xee, 0x46, 0xb0, 0x92, 0xf9,
	0xd2, 0x8f, 0x64, 0xb1, 0x24, 0x8b, 0x9a, 0x39, 0xcf, 0xbf, 0x02, 0x00, 0x00, 0xff, 0xff, 0xbf,
	0xb0, 0x68, 0x96, 0x3b, 0x02, 0x00, 0x00,
}
