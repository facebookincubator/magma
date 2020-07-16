// Code generated by protoc-gen-go. DO NOT EDIT.
// source: storage.proto

package storage

import (
	fmt "fmt"
	math "math"

	proto "github.com/golang/protobuf/proto"
	timestamp "github.com/golang/protobuf/ptypes/timestamp"
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

// TestCase is an end-to-end test case.
type TestCase struct {
	// pk uniquely identifies a test case
	Pk int64 `protobuf:"varint,1,opt,name=pk,proto3" json:"pk,omitempty"`
	// type of the test case which identifies its state machine and how to
	// deserialize its configuration
	TestCaseType string `protobuf:"bytes,2,opt,name=testCaseType,proto3" json:"testCaseType,omitempty"`
	// serialized configuration for the test case
	TestConfig []byte `protobuf:"bytes,10,opt,name=testConfig,proto3" json:"testConfig,omitempty"`
	// flag indicating if the test case is currently being run by a worker
	IsCurrentlyExecuting bool `protobuf:"varint,20,opt,name=isCurrentlyExecuting,proto3" json:"isCurrentlyExecuting,omitempty"`
	// timestamp of the last time the test case was claimed for execution
	LastExecutionTime *timestamp.Timestamp `protobuf:"bytes,21,opt,name=lastExecutionTime,proto3" json:"lastExecutionTime,omitempty"`
	// current state machine state of the test case
	State string `protobuf:"bytes,30,opt,name=state,proto3" json:"state,omitempty"`
	// error message, if any, for the test case
	Error string `protobuf:"bytes,31,opt,name=error,proto3" json:"error,omitempty"`
	// next scheduled runtime for the test case if it is currently idle
	NextScheduledTime    *timestamp.Timestamp `protobuf:"bytes,32,opt,name=nextScheduledTime,proto3" json:"nextScheduledTime,omitempty"`
	XXX_NoUnkeyedLiteral struct{}             `json:"-"`
	XXX_unrecognized     []byte               `json:"-"`
	XXX_sizecache        int32                `json:"-"`
}

func (m *TestCase) Reset()         { *m = TestCase{} }
func (m *TestCase) String() string { return proto.CompactTextString(m) }
func (*TestCase) ProtoMessage()    {}
func (*TestCase) Descriptor() ([]byte, []int) {
	return fileDescriptor_0d2c4ccf1453ffdb, []int{0}
}

func (m *TestCase) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_TestCase.Unmarshal(m, b)
}
func (m *TestCase) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_TestCase.Marshal(b, m, deterministic)
}
func (m *TestCase) XXX_Merge(src proto.Message) {
	xxx_messageInfo_TestCase.Merge(m, src)
}
func (m *TestCase) XXX_Size() int {
	return xxx_messageInfo_TestCase.Size(m)
}
func (m *TestCase) XXX_DiscardUnknown() {
	xxx_messageInfo_TestCase.DiscardUnknown(m)
}

var xxx_messageInfo_TestCase proto.InternalMessageInfo

func (m *TestCase) GetPk() int64 {
	if m != nil {
		return m.Pk
	}
	return 0
}

func (m *TestCase) GetTestCaseType() string {
	if m != nil {
		return m.TestCaseType
	}
	return ""
}

func (m *TestCase) GetTestConfig() []byte {
	if m != nil {
		return m.TestConfig
	}
	return nil
}

func (m *TestCase) GetIsCurrentlyExecuting() bool {
	if m != nil {
		return m.IsCurrentlyExecuting
	}
	return false
}

func (m *TestCase) GetLastExecutionTime() *timestamp.Timestamp {
	if m != nil {
		return m.LastExecutionTime
	}
	return nil
}

func (m *TestCase) GetState() string {
	if m != nil {
		return m.State
	}
	return ""
}

func (m *TestCase) GetError() string {
	if m != nil {
		return m.Error
	}
	return ""
}

func (m *TestCase) GetNextScheduledTime() *timestamp.Timestamp {
	if m != nil {
		return m.NextScheduledTime
	}
	return nil
}

// MutableTestCase encapsulates the set of fields available to clients for
// modification. See TestCase for documentation on fields.
type MutableTestCase struct {
	Pk                   int64    `protobuf:"varint,1,opt,name=pk,proto3" json:"pk,omitempty"`
	TestCaseType         string   `protobuf:"bytes,2,opt,name=testCaseType,proto3" json:"testCaseType,omitempty"`
	TestConfig           []byte   `protobuf:"bytes,10,opt,name=testConfig,proto3" json:"testConfig,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *MutableTestCase) Reset()         { *m = MutableTestCase{} }
func (m *MutableTestCase) String() string { return proto.CompactTextString(m) }
func (*MutableTestCase) ProtoMessage()    {}
func (*MutableTestCase) Descriptor() ([]byte, []int) {
	return fileDescriptor_0d2c4ccf1453ffdb, []int{1}
}

func (m *MutableTestCase) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_MutableTestCase.Unmarshal(m, b)
}
func (m *MutableTestCase) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_MutableTestCase.Marshal(b, m, deterministic)
}
func (m *MutableTestCase) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MutableTestCase.Merge(m, src)
}
func (m *MutableTestCase) XXX_Size() int {
	return xxx_messageInfo_MutableTestCase.Size(m)
}
func (m *MutableTestCase) XXX_DiscardUnknown() {
	xxx_messageInfo_MutableTestCase.DiscardUnknown(m)
}

var xxx_messageInfo_MutableTestCase proto.InternalMessageInfo

func (m *MutableTestCase) GetPk() int64 {
	if m != nil {
		return m.Pk
	}
	return 0
}

func (m *MutableTestCase) GetTestCaseType() string {
	if m != nil {
		return m.TestCaseType
	}
	return ""
}

func (m *MutableTestCase) GetTestConfig() []byte {
	if m != nil {
		return m.TestConfig
	}
	return nil
}

// CINode is a baremetal CI workload executor
type CINode struct {
	// unique ID for the node (e.g. VPN client ID)
	Id string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	// IP address for the node on the VPN
	VpnIp string `protobuf:"bytes,2,opt,name=vpnIp,proto3" json:"vpnIp,omitempty"`
	// is the node available or not
	Available bool `protobuf:"varint,10,opt,name=available,proto3" json:"available,omitempty"`
	// the last time this node was leased out
	LastLeaseTime        *timestamp.Timestamp `protobuf:"bytes,11,opt,name=lastLeaseTime,proto3" json:"lastLeaseTime,omitempty"`
	XXX_NoUnkeyedLiteral struct{}             `json:"-"`
	XXX_unrecognized     []byte               `json:"-"`
	XXX_sizecache        int32                `json:"-"`
}

func (m *CINode) Reset()         { *m = CINode{} }
func (m *CINode) String() string { return proto.CompactTextString(m) }
func (*CINode) ProtoMessage()    {}
func (*CINode) Descriptor() ([]byte, []int) {
	return fileDescriptor_0d2c4ccf1453ffdb, []int{2}
}

func (m *CINode) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CINode.Unmarshal(m, b)
}
func (m *CINode) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CINode.Marshal(b, m, deterministic)
}
func (m *CINode) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CINode.Merge(m, src)
}
func (m *CINode) XXX_Size() int {
	return xxx_messageInfo_CINode.Size(m)
}
func (m *CINode) XXX_DiscardUnknown() {
	xxx_messageInfo_CINode.DiscardUnknown(m)
}

var xxx_messageInfo_CINode proto.InternalMessageInfo

func (m *CINode) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

func (m *CINode) GetVpnIp() string {
	if m != nil {
		return m.VpnIp
	}
	return ""
}

func (m *CINode) GetAvailable() bool {
	if m != nil {
		return m.Available
	}
	return false
}

func (m *CINode) GetLastLeaseTime() *timestamp.Timestamp {
	if m != nil {
		return m.LastLeaseTime
	}
	return nil
}

// MutableCINode encapsulates the set iof fields available to clients for
// modification. See CINode for documentation on fields.
type MutableCINode struct {
	Id                   string   `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	VpnIP                string   `protobuf:"bytes,2,opt,name=vpnIP,proto3" json:"vpnIP,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *MutableCINode) Reset()         { *m = MutableCINode{} }
func (m *MutableCINode) String() string { return proto.CompactTextString(m) }
func (*MutableCINode) ProtoMessage()    {}
func (*MutableCINode) Descriptor() ([]byte, []int) {
	return fileDescriptor_0d2c4ccf1453ffdb, []int{3}
}

func (m *MutableCINode) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_MutableCINode.Unmarshal(m, b)
}
func (m *MutableCINode) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_MutableCINode.Marshal(b, m, deterministic)
}
func (m *MutableCINode) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MutableCINode.Merge(m, src)
}
func (m *MutableCINode) XXX_Size() int {
	return xxx_messageInfo_MutableCINode.Size(m)
}
func (m *MutableCINode) XXX_DiscardUnknown() {
	xxx_messageInfo_MutableCINode.DiscardUnknown(m)
}

var xxx_messageInfo_MutableCINode proto.InternalMessageInfo

func (m *MutableCINode) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

func (m *MutableCINode) GetVpnIP() string {
	if m != nil {
		return m.VpnIP
	}
	return ""
}

// NodeLease encapsulates a successful node lease. To release the lease on the
// node, the same leaseID must be provided.
type NodeLease struct {
	Id                   string   `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	LeaseID              string   `protobuf:"bytes,2,opt,name=leaseID,proto3" json:"leaseID,omitempty"`
	VpnIP                string   `protobuf:"bytes,10,opt,name=vpnIP,proto3" json:"vpnIP,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *NodeLease) Reset()         { *m = NodeLease{} }
func (m *NodeLease) String() string { return proto.CompactTextString(m) }
func (*NodeLease) ProtoMessage()    {}
func (*NodeLease) Descriptor() ([]byte, []int) {
	return fileDescriptor_0d2c4ccf1453ffdb, []int{4}
}

func (m *NodeLease) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_NodeLease.Unmarshal(m, b)
}
func (m *NodeLease) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_NodeLease.Marshal(b, m, deterministic)
}
func (m *NodeLease) XXX_Merge(src proto.Message) {
	xxx_messageInfo_NodeLease.Merge(m, src)
}
func (m *NodeLease) XXX_Size() int {
	return xxx_messageInfo_NodeLease.Size(m)
}
func (m *NodeLease) XXX_DiscardUnknown() {
	xxx_messageInfo_NodeLease.DiscardUnknown(m)
}

var xxx_messageInfo_NodeLease proto.InternalMessageInfo

func (m *NodeLease) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

func (m *NodeLease) GetLeaseID() string {
	if m != nil {
		return m.LeaseID
	}
	return ""
}

func (m *NodeLease) GetVpnIP() string {
	if m != nil {
		return m.VpnIP
	}
	return ""
}

func init() {
	proto.RegisterType((*TestCase)(nil), "magma.fbinternal.testcontroller.storage.TestCase")
	proto.RegisterType((*MutableTestCase)(nil), "magma.fbinternal.testcontroller.storage.MutableTestCase")
	proto.RegisterType((*CINode)(nil), "magma.fbinternal.testcontroller.storage.CINode")
	proto.RegisterType((*MutableCINode)(nil), "magma.fbinternal.testcontroller.storage.MutableCINode")
	proto.RegisterType((*NodeLease)(nil), "magma.fbinternal.testcontroller.storage.NodeLease")
}

func init() { proto.RegisterFile("storage.proto", fileDescriptor_0d2c4ccf1453ffdb) }

var fileDescriptor_0d2c4ccf1453ffdb = []byte{
	// 421 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xb4, 0x53, 0x5d, 0x8b, 0xd3, 0x40,
	0x14, 0x25, 0x11, 0xd7, 0xf6, 0xee, 0x56, 0x71, 0xa8, 0x10, 0x16, 0xd9, 0x0d, 0x79, 0x31, 0x4f,
	0x09, 0xac, 0x08, 0x8a, 0x2f, 0x62, 0x15, 0x2c, 0x7e, 0xb0, 0xc4, 0x3e, 0xf9, 0x36, 0x49, 0x6e,
	0xe3, 0xb0, 0x93, 0x99, 0x30, 0x73, 0x53, 0x76, 0xff, 0x84, 0x3f, 0xce, 0x5f, 0x24, 0x33, 0x49,
	0x69, 0x4b, 0xc5, 0x7d, 0xf2, 0xf1, 0x9c, 0x3b, 0xe7, 0xdc, 0x73, 0x06, 0x2e, 0xcc, 0x2c, 0x69,
	0xc3, 0x1b, 0xcc, 0x3a, 0xa3, 0x49, 0xb3, 0x17, 0x2d, 0x6f, 0x5a, 0x9e, 0xad, 0x4b, 0xa1, 0x08,
	0x8d, 0xe2, 0x32, 0x23, 0xb4, 0x54, 0x69, 0x45, 0x46, 0x4b, 0x89, 0x26, 0x1b, 0x9f, 0x9f, 0x5f,
	0x36, 0x5a, 0x37, 0x12, 0x73, 0x2f, 0x2b, 0xfb, 0x75, 0x4e, 0xa2, 0x45, 0x4b, 0xbc, 0xed, 0x06,
	0xa7, 0xe4, 0x77, 0x08, 0x93, 0x15, 0x5a, 0x5a, 0x70, 0x8b, 0xec, 0x31, 0x84, 0xdd, 0x4d, 0x14,
	0xc4, 0x41, 0xfa, 0xa0, 0x08, 0xbb, 0x1b, 0x96, 0xc0, 0x19, 0x8d, 0xb3, 0xd5, 0x5d, 0x87, 0x51,
	0x18, 0x07, 0xe9, 0xb4, 0x38, 0xe0, 0xd8, 0x05, 0x80, 0xc7, 0x5a, 0xad, 0x45, 0x13, 0x41, 0x1c,
	0xa4, 0x67, 0xc5, 0x1e, 0xc3, 0xae, 0x60, 0x2e, 0xec, 0xa2, 0x37, 0x06, 0x15, 0xc9, 0xbb, 0x8f,
	0xb7, 0x58, 0xf5, 0x24, 0x54, 0x13, 0xcd, 0xe3, 0x20, 0x9d, 0x14, 0x7f, 0x9d, 0xb1, 0x4f, 0xf0,
	0x54, 0x72, 0x4b, 0x23, 0xa1, 0xd5, 0x4a, 0xb4, 0x18, 0x3d, 0x8b, 0x83, 0xf4, 0xf4, 0xea, 0x3c,
	0x1b, 0x1a, 0x65, 0xdb, 0x46, 0xd9, 0x6a, 0xdb, 0xa8, 0x38, 0x16, 0xb1, 0x39, 0x3c, 0xb4, 0xc4,
	0x09, 0xa3, 0x0b, 0x1f, 0x7d, 0x00, 0x8e, 0x45, 0x63, 0xb4, 0x89, 0x2e, 0x07, 0xd6, 0x03, 0xb7,
	0x55, 0xe1, 0x2d, 0x7d, 0xaf, 0x7e, 0x62, 0xdd, 0x4b, 0xac, 0xfd, 0xd6, 0xf8, 0xfe, 0xad, 0x47,
	0xa2, 0x04, 0xe1, 0xc9, 0xd7, 0x9e, 0x78, 0x29, 0xf1, 0x7f, 0x7e, 0x6d, 0xf2, 0x2b, 0x80, 0x93,
	0xc5, 0xf2, 0x9b, 0xae, 0xbd, 0xbd, 0xa8, 0xbd, 0xfd, 0xb4, 0x08, 0x45, 0xed, 0x1a, 0x6e, 0x3a,
	0xb5, 0xec, 0x46, 0xdf, 0x01, 0xb0, 0xe7, 0x30, 0xe5, 0x1b, 0x2e, 0xa4, 0x4b, 0xe6, 0xfd, 0x26,
	0xc5, 0x8e, 0x60, 0xef, 0x60, 0xe6, 0x3e, 0xf0, 0x0b, 0xba, 0xfd, 0xae, 0xfb, 0xe9, 0xbd, 0xdd,
	0x0f, 0x05, 0xc9, 0x2b, 0x98, 0x8d, 0xbd, 0xff, 0x1d, 0xeb, 0x7a, 0x3f, 0xd6, 0x75, 0xf2, 0x19,
	0xa6, 0xee, 0xb5, 0xf7, 0x39, 0x92, 0x44, 0xf0, 0x48, 0xba, 0xc1, 0xf2, 0xc3, 0x28, 0xda, 0xc2,
	0x9d, 0x19, 0xec, 0x99, 0xbd, 0x7f, 0xfb, 0xe3, 0x8d, 0x36, 0xd5, 0x6b, 0x93, 0xef, 0x8e, 0x23,
	0xaf, 0xa4, 0xee, 0xeb, 0xbc, 0xd1, 0xb9, 0x45, 0xb3, 0x11, 0x15, 0xda, 0xfc, 0xf0, 0x5c, 0xf2,
	0xf1, 0x5c, 0xca, 0x13, 0xdf, 0xf1, 0xe5, 0x9f, 0x00, 0x00, 0x00, 0xff, 0xff, 0xc6, 0x4b, 0x8d,
	0xaf, 0x6f, 0x03, 0x00, 0x00,
}
