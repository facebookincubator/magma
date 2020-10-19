// Code generated by protoc-gen-go. DO NOT EDIT.
// source: lte/protos/sctpd.proto

package protos

import (
	context "context"
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
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

type InitRes_InitResult int32

const (
	InitRes_INIT_UNKNOWN InitRes_InitResult = 0
	InitRes_INIT_OK      InitRes_InitResult = 1
	InitRes_INIT_FAIL    InitRes_InitResult = 2
)

var InitRes_InitResult_name = map[int32]string{
	0: "INIT_UNKNOWN",
	1: "INIT_OK",
	2: "INIT_FAIL",
}

var InitRes_InitResult_value = map[string]int32{
	"INIT_UNKNOWN": 0,
	"INIT_OK":      1,
	"INIT_FAIL":    2,
}

func (x InitRes_InitResult) String() string {
	return proto.EnumName(InitRes_InitResult_name, int32(x))
}

func (InitRes_InitResult) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_4b79271ac29ed95c, []int{1, 0}
}

type SendDlRes_SendDlResult int32

const (
	SendDlRes_SEND_DL_UNKNOWN SendDlRes_SendDlResult = 0
	SendDlRes_SEND_DL_OK      SendDlRes_SendDlResult = 1
	SendDlRes_SEND_DL_FAIL    SendDlRes_SendDlResult = 2
)

var SendDlRes_SendDlResult_name = map[int32]string{
	0: "SEND_DL_UNKNOWN",
	1: "SEND_DL_OK",
	2: "SEND_DL_FAIL",
}

var SendDlRes_SendDlResult_value = map[string]int32{
	"SEND_DL_UNKNOWN": 0,
	"SEND_DL_OK":      1,
	"SEND_DL_FAIL":    2,
}

func (x SendDlRes_SendDlResult) String() string {
	return proto.EnumName(SendDlRes_SendDlResult_name, int32(x))
}

func (SendDlRes_SendDlResult) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_4b79271ac29ed95c, []int{3, 0}
}

// InitReq - request for sctpd to initialize sctp connection
type InitReq struct {
	UseIpv4              bool     `protobuf:"varint,1,opt,name=use_ipv4,json=useIpv4,proto3" json:"use_ipv4,omitempty"`
	UseIpv6              bool     `protobuf:"varint,2,opt,name=use_ipv6,json=useIpv6,proto3" json:"use_ipv6,omitempty"`
	Ipv4Addrs            []string `protobuf:"bytes,3,rep,name=ipv4_addrs,json=ipv4Addrs,proto3" json:"ipv4_addrs,omitempty"`
	Ipv6Addrs            []string `protobuf:"bytes,4,rep,name=ipv6_addrs,json=ipv6Addrs,proto3" json:"ipv6_addrs,omitempty"`
	Port                 uint32   `protobuf:"varint,5,opt,name=port,proto3" json:"port,omitempty"`
	Ppid                 uint32   `protobuf:"varint,6,opt,name=ppid,proto3" json:"ppid,omitempty"`
	ForceRestart         bool     `protobuf:"varint,7,opt,name=force_restart,json=forceRestart,proto3" json:"force_restart,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *InitReq) Reset()         { *m = InitReq{} }
func (m *InitReq) String() string { return proto.CompactTextString(m) }
func (*InitReq) ProtoMessage()    {}
func (*InitReq) Descriptor() ([]byte, []int) {
	return fileDescriptor_4b79271ac29ed95c, []int{0}
}

func (m *InitReq) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_InitReq.Unmarshal(m, b)
}
func (m *InitReq) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_InitReq.Marshal(b, m, deterministic)
}
func (m *InitReq) XXX_Merge(src proto.Message) {
	xxx_messageInfo_InitReq.Merge(m, src)
}
func (m *InitReq) XXX_Size() int {
	return xxx_messageInfo_InitReq.Size(m)
}
func (m *InitReq) XXX_DiscardUnknown() {
	xxx_messageInfo_InitReq.DiscardUnknown(m)
}

var xxx_messageInfo_InitReq proto.InternalMessageInfo

func (m *InitReq) GetUseIpv4() bool {
	if m != nil {
		return m.UseIpv4
	}
	return false
}

func (m *InitReq) GetUseIpv6() bool {
	if m != nil {
		return m.UseIpv6
	}
	return false
}

func (m *InitReq) GetIpv4Addrs() []string {
	if m != nil {
		return m.Ipv4Addrs
	}
	return nil
}

func (m *InitReq) GetIpv6Addrs() []string {
	if m != nil {
		return m.Ipv6Addrs
	}
	return nil
}

func (m *InitReq) GetPort() uint32 {
	if m != nil {
		return m.Port
	}
	return 0
}

func (m *InitReq) GetPpid() uint32 {
	if m != nil {
		return m.Ppid
	}
	return 0
}

func (m *InitReq) GetForceRestart() bool {
	if m != nil {
		return m.ForceRestart
	}
	return false
}

// InitRes - response with status of sctp initialization
type InitRes struct {
	Result               InitRes_InitResult `protobuf:"varint,1,opt,name=result,proto3,enum=magma.sctpd.InitRes_InitResult" json:"result,omitempty"`
	XXX_NoUnkeyedLiteral struct{}           `json:"-"`
	XXX_unrecognized     []byte             `json:"-"`
	XXX_sizecache        int32              `json:"-"`
}

func (m *InitRes) Reset()         { *m = InitRes{} }
func (m *InitRes) String() string { return proto.CompactTextString(m) }
func (*InitRes) ProtoMessage()    {}
func (*InitRes) Descriptor() ([]byte, []int) {
	return fileDescriptor_4b79271ac29ed95c, []int{1}
}

func (m *InitRes) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_InitRes.Unmarshal(m, b)
}
func (m *InitRes) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_InitRes.Marshal(b, m, deterministic)
}
func (m *InitRes) XXX_Merge(src proto.Message) {
	xxx_messageInfo_InitRes.Merge(m, src)
}
func (m *InitRes) XXX_Size() int {
	return xxx_messageInfo_InitRes.Size(m)
}
func (m *InitRes) XXX_DiscardUnknown() {
	xxx_messageInfo_InitRes.DiscardUnknown(m)
}

var xxx_messageInfo_InitRes proto.InternalMessageInfo

func (m *InitRes) GetResult() InitRes_InitResult {
	if m != nil {
		return m.Result
	}
	return InitRes_INIT_UNKNOWN
}

// SendDlReq - requests a downlink packet to be sent to eNB
type SendDlReq struct {
	Ppid                 uint32   `protobuf:"varint,1,opt,name=ppid,proto3" json:"ppid,omitempty"`
	AssocId              uint32   `protobuf:"varint,2,opt,name=assoc_id,json=assocId,proto3" json:"assoc_id,omitempty"`
	Stream               uint32   `protobuf:"varint,3,opt,name=stream,proto3" json:"stream,omitempty"`
	Payload              []byte   `protobuf:"bytes,4,opt,name=payload,proto3" json:"payload,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *SendDlReq) Reset()         { *m = SendDlReq{} }
func (m *SendDlReq) String() string { return proto.CompactTextString(m) }
func (*SendDlReq) ProtoMessage()    {}
func (*SendDlReq) Descriptor() ([]byte, []int) {
	return fileDescriptor_4b79271ac29ed95c, []int{2}
}

func (m *SendDlReq) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_SendDlReq.Unmarshal(m, b)
}
func (m *SendDlReq) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_SendDlReq.Marshal(b, m, deterministic)
}
func (m *SendDlReq) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SendDlReq.Merge(m, src)
}
func (m *SendDlReq) XXX_Size() int {
	return xxx_messageInfo_SendDlReq.Size(m)
}
func (m *SendDlReq) XXX_DiscardUnknown() {
	xxx_messageInfo_SendDlReq.DiscardUnknown(m)
}

var xxx_messageInfo_SendDlReq proto.InternalMessageInfo

func (m *SendDlReq) GetPpid() uint32 {
	if m != nil {
		return m.Ppid
	}
	return 0
}

func (m *SendDlReq) GetAssocId() uint32 {
	if m != nil {
		return m.AssocId
	}
	return 0
}

func (m *SendDlReq) GetStream() uint32 {
	if m != nil {
		return m.Stream
	}
	return 0
}

func (m *SendDlReq) GetPayload() []byte {
	if m != nil {
		return m.Payload
	}
	return nil
}

// SendDlRes - response with status of downlink packet send
type SendDlRes struct {
	Result               SendDlRes_SendDlResult `protobuf:"varint,1,opt,name=result,proto3,enum=magma.sctpd.SendDlRes_SendDlResult" json:"result,omitempty"`
	XXX_NoUnkeyedLiteral struct{}               `json:"-"`
	XXX_unrecognized     []byte                 `json:"-"`
	XXX_sizecache        int32                  `json:"-"`
}

func (m *SendDlRes) Reset()         { *m = SendDlRes{} }
func (m *SendDlRes) String() string { return proto.CompactTextString(m) }
func (*SendDlRes) ProtoMessage()    {}
func (*SendDlRes) Descriptor() ([]byte, []int) {
	return fileDescriptor_4b79271ac29ed95c, []int{3}
}

func (m *SendDlRes) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_SendDlRes.Unmarshal(m, b)
}
func (m *SendDlRes) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_SendDlRes.Marshal(b, m, deterministic)
}
func (m *SendDlRes) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SendDlRes.Merge(m, src)
}
func (m *SendDlRes) XXX_Size() int {
	return xxx_messageInfo_SendDlRes.Size(m)
}
func (m *SendDlRes) XXX_DiscardUnknown() {
	xxx_messageInfo_SendDlRes.DiscardUnknown(m)
}

var xxx_messageInfo_SendDlRes proto.InternalMessageInfo

func (m *SendDlRes) GetResult() SendDlRes_SendDlResult {
	if m != nil {
		return m.Result
	}
	return SendDlRes_SEND_DL_UNKNOWN
}

// SendUlReq - requests an uplink packet to be sent to MME
type SendUlReq struct {
	Ppid                 uint32   `protobuf:"varint,1,opt,name=ppid,proto3" json:"ppid,omitempty"`
	AssocId              uint32   `protobuf:"varint,2,opt,name=assoc_id,json=assocId,proto3" json:"assoc_id,omitempty"`
	Stream               uint32   `protobuf:"varint,3,opt,name=stream,proto3" json:"stream,omitempty"`
	Payload              []byte   `protobuf:"bytes,4,opt,name=payload,proto3" json:"payload,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *SendUlReq) Reset()         { *m = SendUlReq{} }
func (m *SendUlReq) String() string { return proto.CompactTextString(m) }
func (*SendUlReq) ProtoMessage()    {}
func (*SendUlReq) Descriptor() ([]byte, []int) {
	return fileDescriptor_4b79271ac29ed95c, []int{4}
}

func (m *SendUlReq) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_SendUlReq.Unmarshal(m, b)
}
func (m *SendUlReq) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_SendUlReq.Marshal(b, m, deterministic)
}
func (m *SendUlReq) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SendUlReq.Merge(m, src)
}
func (m *SendUlReq) XXX_Size() int {
	return xxx_messageInfo_SendUlReq.Size(m)
}
func (m *SendUlReq) XXX_DiscardUnknown() {
	xxx_messageInfo_SendUlReq.DiscardUnknown(m)
}

var xxx_messageInfo_SendUlReq proto.InternalMessageInfo

func (m *SendUlReq) GetPpid() uint32 {
	if m != nil {
		return m.Ppid
	}
	return 0
}

func (m *SendUlReq) GetAssocId() uint32 {
	if m != nil {
		return m.AssocId
	}
	return 0
}

func (m *SendUlReq) GetStream() uint32 {
	if m != nil {
		return m.Stream
	}
	return 0
}

func (m *SendUlReq) GetPayload() []byte {
	if m != nil {
		return m.Payload
	}
	return nil
}

// SendUlRes - response for SendUlReq, present for forwards compat
type SendUlRes struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *SendUlRes) Reset()         { *m = SendUlRes{} }
func (m *SendUlRes) String() string { return proto.CompactTextString(m) }
func (*SendUlRes) ProtoMessage()    {}
func (*SendUlRes) Descriptor() ([]byte, []int) {
	return fileDescriptor_4b79271ac29ed95c, []int{5}
}

func (m *SendUlRes) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_SendUlRes.Unmarshal(m, b)
}
func (m *SendUlRes) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_SendUlRes.Marshal(b, m, deterministic)
}
func (m *SendUlRes) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SendUlRes.Merge(m, src)
}
func (m *SendUlRes) XXX_Size() int {
	return xxx_messageInfo_SendUlRes.Size(m)
}
func (m *SendUlRes) XXX_DiscardUnknown() {
	xxx_messageInfo_SendUlRes.DiscardUnknown(m)
}

var xxx_messageInfo_SendUlRes proto.InternalMessageInfo

// NewAssocReq - request to notify MME of new eNB association
type NewAssocReq struct {
	Ppid                 uint32   `protobuf:"varint,1,opt,name=ppid,proto3" json:"ppid,omitempty"`
	AssocId              uint32   `protobuf:"varint,2,opt,name=assoc_id,json=assocId,proto3" json:"assoc_id,omitempty"`
	Instreams            uint32   `protobuf:"varint,3,opt,name=instreams,proto3" json:"instreams,omitempty"`
	Outstreams           uint32   `protobuf:"varint,4,opt,name=outstreams,proto3" json:"outstreams,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *NewAssocReq) Reset()         { *m = NewAssocReq{} }
func (m *NewAssocReq) String() string { return proto.CompactTextString(m) }
func (*NewAssocReq) ProtoMessage()    {}
func (*NewAssocReq) Descriptor() ([]byte, []int) {
	return fileDescriptor_4b79271ac29ed95c, []int{6}
}

func (m *NewAssocReq) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_NewAssocReq.Unmarshal(m, b)
}
func (m *NewAssocReq) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_NewAssocReq.Marshal(b, m, deterministic)
}
func (m *NewAssocReq) XXX_Merge(src proto.Message) {
	xxx_messageInfo_NewAssocReq.Merge(m, src)
}
func (m *NewAssocReq) XXX_Size() int {
	return xxx_messageInfo_NewAssocReq.Size(m)
}
func (m *NewAssocReq) XXX_DiscardUnknown() {
	xxx_messageInfo_NewAssocReq.DiscardUnknown(m)
}

var xxx_messageInfo_NewAssocReq proto.InternalMessageInfo

func (m *NewAssocReq) GetPpid() uint32 {
	if m != nil {
		return m.Ppid
	}
	return 0
}

func (m *NewAssocReq) GetAssocId() uint32 {
	if m != nil {
		return m.AssocId
	}
	return 0
}

func (m *NewAssocReq) GetInstreams() uint32 {
	if m != nil {
		return m.Instreams
	}
	return 0
}

func (m *NewAssocReq) GetOutstreams() uint32 {
	if m != nil {
		return m.Outstreams
	}
	return 0
}

// NewAssocRes - response for NewAssocReq, present for forwards compat
type NewAssocRes struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *NewAssocRes) Reset()         { *m = NewAssocRes{} }
func (m *NewAssocRes) String() string { return proto.CompactTextString(m) }
func (*NewAssocRes) ProtoMessage()    {}
func (*NewAssocRes) Descriptor() ([]byte, []int) {
	return fileDescriptor_4b79271ac29ed95c, []int{7}
}

func (m *NewAssocRes) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_NewAssocRes.Unmarshal(m, b)
}
func (m *NewAssocRes) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_NewAssocRes.Marshal(b, m, deterministic)
}
func (m *NewAssocRes) XXX_Merge(src proto.Message) {
	xxx_messageInfo_NewAssocRes.Merge(m, src)
}
func (m *NewAssocRes) XXX_Size() int {
	return xxx_messageInfo_NewAssocRes.Size(m)
}
func (m *NewAssocRes) XXX_DiscardUnknown() {
	xxx_messageInfo_NewAssocRes.DiscardUnknown(m)
}

var xxx_messageInfo_NewAssocRes proto.InternalMessageInfo

// CloseAssocReq - request to notify MME of a closing/resetting assocation
type CloseAssocReq struct {
	Ppid                 uint32   `protobuf:"varint,1,opt,name=ppid,proto3" json:"ppid,omitempty"`
	AssocId              uint32   `protobuf:"varint,2,opt,name=assoc_id,json=assocId,proto3" json:"assoc_id,omitempty"`
	IsReset              bool     `protobuf:"varint,3,opt,name=is_reset,json=isReset,proto3" json:"is_reset,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *CloseAssocReq) Reset()         { *m = CloseAssocReq{} }
func (m *CloseAssocReq) String() string { return proto.CompactTextString(m) }
func (*CloseAssocReq) ProtoMessage()    {}
func (*CloseAssocReq) Descriptor() ([]byte, []int) {
	return fileDescriptor_4b79271ac29ed95c, []int{8}
}

func (m *CloseAssocReq) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CloseAssocReq.Unmarshal(m, b)
}
func (m *CloseAssocReq) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CloseAssocReq.Marshal(b, m, deterministic)
}
func (m *CloseAssocReq) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CloseAssocReq.Merge(m, src)
}
func (m *CloseAssocReq) XXX_Size() int {
	return xxx_messageInfo_CloseAssocReq.Size(m)
}
func (m *CloseAssocReq) XXX_DiscardUnknown() {
	xxx_messageInfo_CloseAssocReq.DiscardUnknown(m)
}

var xxx_messageInfo_CloseAssocReq proto.InternalMessageInfo

func (m *CloseAssocReq) GetPpid() uint32 {
	if m != nil {
		return m.Ppid
	}
	return 0
}

func (m *CloseAssocReq) GetAssocId() uint32 {
	if m != nil {
		return m.AssocId
	}
	return 0
}

func (m *CloseAssocReq) GetIsReset() bool {
	if m != nil {
		return m.IsReset
	}
	return false
}

// CloseAssocRes - response for CloseAssocReq, present for forwards compat
type CloseAssocRes struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *CloseAssocRes) Reset()         { *m = CloseAssocRes{} }
func (m *CloseAssocRes) String() string { return proto.CompactTextString(m) }
func (*CloseAssocRes) ProtoMessage()    {}
func (*CloseAssocRes) Descriptor() ([]byte, []int) {
	return fileDescriptor_4b79271ac29ed95c, []int{9}
}

func (m *CloseAssocRes) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CloseAssocRes.Unmarshal(m, b)
}
func (m *CloseAssocRes) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CloseAssocRes.Marshal(b, m, deterministic)
}
func (m *CloseAssocRes) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CloseAssocRes.Merge(m, src)
}
func (m *CloseAssocRes) XXX_Size() int {
	return xxx_messageInfo_CloseAssocRes.Size(m)
}
func (m *CloseAssocRes) XXX_DiscardUnknown() {
	xxx_messageInfo_CloseAssocRes.DiscardUnknown(m)
}

var xxx_messageInfo_CloseAssocRes proto.InternalMessageInfo

func init() {
	proto.RegisterEnum("magma.sctpd.InitRes_InitResult", InitRes_InitResult_name, InitRes_InitResult_value)
	proto.RegisterEnum("magma.sctpd.SendDlRes_SendDlResult", SendDlRes_SendDlResult_name, SendDlRes_SendDlResult_value)
	proto.RegisterType((*InitReq)(nil), "magma.sctpd.InitReq")
	proto.RegisterType((*InitRes)(nil), "magma.sctpd.InitRes")
	proto.RegisterType((*SendDlReq)(nil), "magma.sctpd.SendDlReq")
	proto.RegisterType((*SendDlRes)(nil), "magma.sctpd.SendDlRes")
	proto.RegisterType((*SendUlReq)(nil), "magma.sctpd.SendUlReq")
	proto.RegisterType((*SendUlRes)(nil), "magma.sctpd.SendUlRes")
	proto.RegisterType((*NewAssocReq)(nil), "magma.sctpd.NewAssocReq")
	proto.RegisterType((*NewAssocRes)(nil), "magma.sctpd.NewAssocRes")
	proto.RegisterType((*CloseAssocReq)(nil), "magma.sctpd.CloseAssocReq")
	proto.RegisterType((*CloseAssocRes)(nil), "magma.sctpd.CloseAssocRes")
}

func init() { proto.RegisterFile("lte/protos/sctpd.proto", fileDescriptor_4b79271ac29ed95c) }

var fileDescriptor_4b79271ac29ed95c = []byte{
	// 579 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xbc, 0x94, 0x4f, 0x6f, 0xda, 0x4c,
	0x10, 0xc6, 0x71, 0xe0, 0xc5, 0x66, 0xc0, 0x09, 0xda, 0xb7, 0x42, 0x86, 0xfe, 0x43, 0x9b, 0x0b,
	0x27, 0x90, 0x28, 0xa2, 0x52, 0x7a, 0x29, 0x2d, 0x89, 0x64, 0x25, 0x72, 0x24, 0x53, 0x54, 0xb5,
	0x17, 0xcb, 0xc5, 0xdb, 0xc8, 0xaa, 0xc1, 0xc6, 0xb3, 0x4e, 0xd4, 0x1e, 0x2a, 0xf5, 0x13, 0xf4,
	0x7b, 0xf5, 0xda, 0x2f, 0x54, 0xed, 0x62, 0xc7, 0x76, 0x05, 0x97, 0x1c, 0x7a, 0xb2, 0xe7, 0xf7,
	0x8c, 0x67, 0xe6, 0xd9, 0x1d, 0x19, 0x3a, 0x01, 0x67, 0xa3, 0x28, 0x0e, 0x79, 0x88, 0x23, 0x5c,
	0xf1, 0xc8, 0x1b, 0xca, 0x80, 0x34, 0xd7, 0xee, 0xcd, 0xda, 0x1d, 0x4a, 0x44, 0x7f, 0x29, 0xa0,
	0x9a, 0x1b, 0x9f, 0xdb, 0x6c, 0x4b, 0xba, 0xa0, 0x25, 0xc8, 0x1c, 0x3f, 0xba, 0x9d, 0x18, 0x4a,
	0x5f, 0x19, 0x68, 0xb6, 0x9a, 0x20, 0x33, 0xa3, 0xdb, 0x49, 0x41, 0x9a, 0x1a, 0x47, 0x45, 0x69,
	0x4a, 0x9e, 0x02, 0x88, 0x2f, 0x1c, 0xd7, 0xf3, 0x62, 0x34, 0xaa, 0xfd, 0xea, 0xa0, 0x61, 0x37,
	0x04, 0x99, 0x09, 0x90, 0xca, 0xd3, 0x54, 0xae, 0xdd, 0xcb, 0xd3, 0x9d, 0x4c, 0xa0, 0x16, 0x85,
	0x31, 0x37, 0xfe, 0xeb, 0x2b, 0x03, 0xdd, 0x96, 0xef, 0x92, 0x45, 0xbe, 0x67, 0xd4, 0x53, 0x16,
	0xf9, 0x1e, 0x39, 0x05, 0xfd, 0x73, 0x18, 0xaf, 0x98, 0x13, 0x33, 0xe4, 0x6e, 0xcc, 0x0d, 0x55,
	0x4e, 0xd1, 0x92, 0xd0, 0xde, 0x31, 0xfa, 0x3d, 0xf3, 0x82, 0xe4, 0x25, 0xd4, 0x63, 0x86, 0x49,
	0xc0, 0xa5, 0x93, 0xe3, 0xf1, 0xf3, 0x61, 0xc1, 0xf5, 0x30, 0xcd, 0xca, 0x9e, 0x49, 0xc0, 0xed,
	0x34, 0x9d, 0x9e, 0x01, 0xe4, 0x94, 0xb4, 0xa1, 0x65, 0x5a, 0xe6, 0x3b, 0x67, 0x69, 0x5d, 0x5a,
	0xd7, 0xef, 0xad, 0x76, 0x85, 0x34, 0x41, 0x95, 0xe4, 0xfa, 0xb2, 0xad, 0x10, 0x1d, 0x1a, 0x32,
	0xb8, 0x98, 0x99, 0x57, 0xed, 0x23, 0x1a, 0x40, 0x63, 0xc1, 0x36, 0xde, 0x3c, 0x10, 0xa7, 0x99,
	0xb9, 0x50, 0x0a, 0x2e, 0xba, 0xa0, 0xb9, 0x88, 0xe1, 0xca, 0xf1, 0x3d, 0x79, 0x8c, 0xba, 0xad,
	0xca, 0xd8, 0xf4, 0x48, 0x07, 0xea, 0xc8, 0x63, 0xe6, 0xae, 0x8d, 0xaa, 0x14, 0xd2, 0x88, 0x18,
	0xa0, 0x46, 0xee, 0xd7, 0x20, 0x74, 0x3d, 0xa3, 0xd6, 0x57, 0x06, 0x2d, 0x3b, 0x0b, 0xe9, 0x4f,
	0x25, 0x6f, 0x87, 0xe4, 0xd5, 0x5f, 0x86, 0x4f, 0x4b, 0x86, 0xef, 0xf3, 0xf2, 0xb7, 0xa2, 0xe9,
	0x73, 0x68, 0x15, 0x39, 0xf9, 0x1f, 0x4e, 0x16, 0xe7, 0xd6, 0xdc, 0x99, 0x5f, 0x15, 0x9c, 0x1f,
	0x03, 0x64, 0x50, 0x9a, 0x6f, 0x43, 0x2b, 0x8b, 0xcb, 0xfe, 0x97, 0xff, 0xc6, 0x7f, 0x33, 0xef,
	0x86, 0xf4, 0x1b, 0x34, 0x2d, 0x76, 0x37, 0x13, 0xc5, 0x1e, 0xd0, 0xfc, 0x09, 0x34, 0xfc, 0xcd,
	0xae, 0x21, 0xa6, 0xfd, 0x73, 0x40, 0x9e, 0x01, 0x84, 0x09, 0xcf, 0xe4, 0x9a, 0x94, 0x0b, 0x84,
	0xea, 0xc5, 0xde, 0x48, 0x3f, 0x80, 0xfe, 0x36, 0x08, 0x91, 0x3d, 0x74, 0x98, 0x2e, 0x68, 0x3e,
	0x8a, 0x3d, 0x67, 0x5c, 0xce, 0xa2, 0xd9, 0xaa, 0x8f, 0xb6, 0x08, 0xe9, 0x49, 0xb9, 0x34, 0x8e,
	0x7f, 0x28, 0xa0, 0x2f, 0xc4, 0x0d, 0xcf, 0xc3, 0xbb, 0x4d, 0xe0, 0x6f, 0xbe, 0x90, 0x09, 0xd4,
	0xc4, 0xfe, 0x92, 0x47, 0x7b, 0x16, 0x7e, 0xdb, 0xdb, 0x47, 0x91, 0x56, 0xc8, 0x19, 0xd4, 0x77,
	0x0b, 0x40, 0x3a, 0x7b, 0xf7, 0x66, 0xdb, 0xdb, 0xcf, 0x91, 0x56, 0xc6, 0xbf, 0x15, 0x68, 0xca,
	0x19, 0x96, 0x91, 0x9c, 0x20, 0xad, 0xb5, 0xdc, 0x57, 0x6b, 0x79, 0xa0, 0xd6, 0xee, 0x12, 0x2b,
	0xe4, 0x35, 0x68, 0xd9, 0x51, 0x12, 0xa3, 0x94, 0x55, 0xb8, 0xdd, 0xde, 0x21, 0x45, 0x54, 0xb8,
	0x00, 0xc8, 0x8f, 0x88, 0xf4, 0x4a, 0x99, 0xa5, 0x6b, 0xe9, 0x1d, 0xd6, 0x90, 0x56, 0xde, 0x3c,
	0xfe, 0xd8, 0x95, 0xf2, 0x48, 0xfc, 0x45, 0x57, 0x41, 0x98, 0x78, 0xa3, 0x9b, 0x30, 0xfd, 0x9d,
	0x7e, 0xaa, 0xcb, 0xe7, 0x8b, 0x3f, 0x01, 0x00, 0x00, 0xff, 0xff, 0x7c, 0x82, 0x8b, 0x17, 0x63,
	0x05, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConnInterface

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion6

// SctpdDownlinkClient is the client API for SctpdDownlink service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type SctpdDownlinkClient interface {
	// Init - initialize sctp connection according to InitReq
	// @param InitReq request specifying desired sctp configuration
	// @return InitRes response w/ init success status
	Init(ctx context.Context, in *InitReq, opts ...grpc.CallOption) (*InitRes, error)
	// SendDl - send a downlink packet to eNB
	// @param SendDlReq request specifying packet data and destination
	// @return SendDlRes response w/ send success status
	SendDl(ctx context.Context, in *SendDlReq, opts ...grpc.CallOption) (*SendDlRes, error)
}

type sctpdDownlinkClient struct {
	cc grpc.ClientConnInterface
}

func NewSctpdDownlinkClient(cc grpc.ClientConnInterface) SctpdDownlinkClient {
	return &sctpdDownlinkClient{cc}
}

func (c *sctpdDownlinkClient) Init(ctx context.Context, in *InitReq, opts ...grpc.CallOption) (*InitRes, error) {
	out := new(InitRes)
	err := c.cc.Invoke(ctx, "/magma.sctpd.SctpdDownlink/Init", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *sctpdDownlinkClient) SendDl(ctx context.Context, in *SendDlReq, opts ...grpc.CallOption) (*SendDlRes, error) {
	out := new(SendDlRes)
	err := c.cc.Invoke(ctx, "/magma.sctpd.SctpdDownlink/SendDl", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// SctpdDownlinkServer is the server API for SctpdDownlink service.
type SctpdDownlinkServer interface {
	// Init - initialize sctp connection according to InitReq
	// @param InitReq request specifying desired sctp configuration
	// @return InitRes response w/ init success status
	Init(context.Context, *InitReq) (*InitRes, error)
	// SendDl - send a downlink packet to eNB
	// @param SendDlReq request specifying packet data and destination
	// @return SendDlRes response w/ send success status
	SendDl(context.Context, *SendDlReq) (*SendDlRes, error)
}

// UnimplementedSctpdDownlinkServer can be embedded to have forward compatible implementations.
type UnimplementedSctpdDownlinkServer struct {
}

func (*UnimplementedSctpdDownlinkServer) Init(ctx context.Context, req *InitReq) (*InitRes, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Init not implemented")
}
func (*UnimplementedSctpdDownlinkServer) SendDl(ctx context.Context, req *SendDlReq) (*SendDlRes, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SendDl not implemented")
}

func RegisterSctpdDownlinkServer(s *grpc.Server, srv SctpdDownlinkServer) {
	s.RegisterService(&_SctpdDownlink_serviceDesc, srv)
}

func _SctpdDownlink_Init_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(InitReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SctpdDownlinkServer).Init(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/magma.sctpd.SctpdDownlink/Init",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SctpdDownlinkServer).Init(ctx, req.(*InitReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _SctpdDownlink_SendDl_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SendDlReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SctpdDownlinkServer).SendDl(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/magma.sctpd.SctpdDownlink/SendDl",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SctpdDownlinkServer).SendDl(ctx, req.(*SendDlReq))
	}
	return interceptor(ctx, in, info, handler)
}

var _SctpdDownlink_serviceDesc = grpc.ServiceDesc{
	ServiceName: "magma.sctpd.SctpdDownlink",
	HandlerType: (*SctpdDownlinkServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Init",
			Handler:    _SctpdDownlink_Init_Handler,
		},
		{
			MethodName: "SendDl",
			Handler:    _SctpdDownlink_SendDl_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "lte/protos/sctpd.proto",
}

// SctpdUplinkClient is the client API for SctpdUplink service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type SctpdUplinkClient interface {
	// SendUl - send an uplink packet to MME
	// @param SendUlReq request specifying packet data and destination
	// @return SendUlRes void response object
	SendUl(ctx context.Context, in *SendUlReq, opts ...grpc.CallOption) (*SendUlRes, error)
	// NewAssoc - notify MME of new eNB association
	// @param NewAssocReq request specifying new association's information
	// @return NewAssocRes void response object
	NewAssoc(ctx context.Context, in *NewAssocReq, opts ...grpc.CallOption) (*NewAssocRes, error)
	// CloseAssoc - notify MME of closing/reseting eNB assocation
	// @param CloseAssocReq request specifying closing assocation and close type
	// @return CloseAssocRes void response object
	CloseAssoc(ctx context.Context, in *CloseAssocReq, opts ...grpc.CallOption) (*CloseAssocRes, error)
}

type sctpdUplinkClient struct {
	cc grpc.ClientConnInterface
}

func NewSctpdUplinkClient(cc grpc.ClientConnInterface) SctpdUplinkClient {
	return &sctpdUplinkClient{cc}
}

func (c *sctpdUplinkClient) SendUl(ctx context.Context, in *SendUlReq, opts ...grpc.CallOption) (*SendUlRes, error) {
	out := new(SendUlRes)
	err := c.cc.Invoke(ctx, "/magma.sctpd.SctpdUplink/SendUl", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *sctpdUplinkClient) NewAssoc(ctx context.Context, in *NewAssocReq, opts ...grpc.CallOption) (*NewAssocRes, error) {
	out := new(NewAssocRes)
	err := c.cc.Invoke(ctx, "/magma.sctpd.SctpdUplink/NewAssoc", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *sctpdUplinkClient) CloseAssoc(ctx context.Context, in *CloseAssocReq, opts ...grpc.CallOption) (*CloseAssocRes, error) {
	out := new(CloseAssocRes)
	err := c.cc.Invoke(ctx, "/magma.sctpd.SctpdUplink/CloseAssoc", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// SctpdUplinkServer is the server API for SctpdUplink service.
type SctpdUplinkServer interface {
	// SendUl - send an uplink packet to MME
	// @param SendUlReq request specifying packet data and destination
	// @return SendUlRes void response object
	SendUl(context.Context, *SendUlReq) (*SendUlRes, error)
	// NewAssoc - notify MME of new eNB association
	// @param NewAssocReq request specifying new association's information
	// @return NewAssocRes void response object
	NewAssoc(context.Context, *NewAssocReq) (*NewAssocRes, error)
	// CloseAssoc - notify MME of closing/reseting eNB assocation
	// @param CloseAssocReq request specifying closing assocation and close type
	// @return CloseAssocRes void response object
	CloseAssoc(context.Context, *CloseAssocReq) (*CloseAssocRes, error)
}

// UnimplementedSctpdUplinkServer can be embedded to have forward compatible implementations.
type UnimplementedSctpdUplinkServer struct {
}

func (*UnimplementedSctpdUplinkServer) SendUl(ctx context.Context, req *SendUlReq) (*SendUlRes, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SendUl not implemented")
}
func (*UnimplementedSctpdUplinkServer) NewAssoc(ctx context.Context, req *NewAssocReq) (*NewAssocRes, error) {
	return nil, status.Errorf(codes.Unimplemented, "method NewAssoc not implemented")
}
func (*UnimplementedSctpdUplinkServer) CloseAssoc(ctx context.Context, req *CloseAssocReq) (*CloseAssocRes, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CloseAssoc not implemented")
}

func RegisterSctpdUplinkServer(s *grpc.Server, srv SctpdUplinkServer) {
	s.RegisterService(&_SctpdUplink_serviceDesc, srv)
}

func _SctpdUplink_SendUl_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SendUlReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SctpdUplinkServer).SendUl(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/magma.sctpd.SctpdUplink/SendUl",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SctpdUplinkServer).SendUl(ctx, req.(*SendUlReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _SctpdUplink_NewAssoc_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(NewAssocReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SctpdUplinkServer).NewAssoc(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/magma.sctpd.SctpdUplink/NewAssoc",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SctpdUplinkServer).NewAssoc(ctx, req.(*NewAssocReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _SctpdUplink_CloseAssoc_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CloseAssocReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SctpdUplinkServer).CloseAssoc(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/magma.sctpd.SctpdUplink/CloseAssoc",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SctpdUplinkServer).CloseAssoc(ctx, req.(*CloseAssocReq))
	}
	return interceptor(ctx, in, info, handler)
}

var _SctpdUplink_serviceDesc = grpc.ServiceDesc{
	ServiceName: "magma.sctpd.SctpdUplink",
	HandlerType: (*SctpdUplinkServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "SendUl",
			Handler:    _SctpdUplink_SendUl_Handler,
		},
		{
			MethodName: "NewAssoc",
			Handler:    _SctpdUplink_NewAssoc_Handler,
		},
		{
			MethodName: "CloseAssoc",
			Handler:    _SctpdUplink_CloseAssoc_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "lte/protos/sctpd.proto",
}
