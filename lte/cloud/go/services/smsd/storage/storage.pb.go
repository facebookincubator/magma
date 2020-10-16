// Code generated by protoc-gen-go. DO NOT EDIT.
// source: storage.proto

package storage

import (
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	timestamp "github.com/golang/protobuf/ptypes/timestamp"
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

type MessageStatus int32

const (
	MessageStatus_WAITING   MessageStatus = 0
	MessageStatus_DELIVERED MessageStatus = 1
	MessageStatus_FAILED    MessageStatus = 2
)

var MessageStatus_name = map[int32]string{
	0: "WAITING",
	1: "DELIVERED",
	2: "FAILED",
}

var MessageStatus_value = map[string]int32{
	"WAITING":   0,
	"DELIVERED": 1,
	"FAILED":    2,
}

func (x MessageStatus) String() string {
	return proto.EnumName(MessageStatus_name, int32(x))
}

func (MessageStatus) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_0d2c4ccf1453ffdb, []int{0}
}

// SMS represents a message tracked by the smsd service
type SMS struct {
	// pk uniquely identifies an SMS message (generated unique key)
	Pk string `protobuf:"bytes,1,opt,name=pk,proto3" json:"pk,omitempty"`
	// delivery status of the SMS
	Status MessageStatus `protobuf:"varint,2,opt,name=status,proto3,enum=magma.lte.smsd.storage.MessageStatus" json:"status,omitempty"`
	// destination for the message
	Imsi string `protobuf:"bytes,10,opt,name=imsi,proto3" json:"imsi,omitempty"`
	// source MSISDN for the mesasge, to be encoded into the SMS
	SourceMsisdn string `protobuf:"bytes,11,opt,name=sourceMsisdn,proto3" json:"sourceMsisdn,omitempty"`
	// the desired message content of the SMS
	Message string `protobuf:"bytes,12,opt,name=message,proto3" json:"message,omitempty"`
	// time at which the message was created in the system
	CreatedTime *timestamp.Timestamp `protobuf:"bytes,20,opt,name=createdTime,proto3" json:"createdTime,omitempty"`
	// time that we last tried delivering this message. if status is delivered,
	// this will be the delivery time
	LastDeliveryAttemptTime *timestamp.Timestamp `protobuf:"bytes,21,opt,name=lastDeliveryAttemptTime,proto3" json:"lastDeliveryAttemptTime,omitempty"`
	// number of times we've attempted to send this SMS
	AttemptCount uint32 `protobuf:"varint,22,opt,name=attemptCount,proto3" json:"attemptCount,omitempty"`
	// error message from the most recent failed delivery attempt
	DeliveryError string `protobuf:"bytes,23,opt,name=deliveryError,proto3" json:"deliveryError,omitempty"`
	// Internal field which holds the reference numbers assigned to an SMS
	// which is in flight.
	// Value is a bytearray because one message could result in multiple SMSs
	// being sent to the UE.
	// If this field is non-empty, the message is in flight. Otherwise, it has
	// yet to be delivered to an AGW.
	RefNums              []byte   `protobuf:"bytes,30,opt,name=refNums,proto3" json:"refNums,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *SMS) Reset()         { *m = SMS{} }
func (m *SMS) String() string { return proto.CompactTextString(m) }
func (*SMS) ProtoMessage()    {}
func (*SMS) Descriptor() ([]byte, []int) {
	return fileDescriptor_0d2c4ccf1453ffdb, []int{0}
}

func (m *SMS) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_SMS.Unmarshal(m, b)
}
func (m *SMS) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_SMS.Marshal(b, m, deterministic)
}
func (m *SMS) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SMS.Merge(m, src)
}
func (m *SMS) XXX_Size() int {
	return xxx_messageInfo_SMS.Size(m)
}
func (m *SMS) XXX_DiscardUnknown() {
	xxx_messageInfo_SMS.DiscardUnknown(m)
}

var xxx_messageInfo_SMS proto.InternalMessageInfo

func (m *SMS) GetPk() string {
	if m != nil {
		return m.Pk
	}
	return ""
}

func (m *SMS) GetStatus() MessageStatus {
	if m != nil {
		return m.Status
	}
	return MessageStatus_WAITING
}

func (m *SMS) GetImsi() string {
	if m != nil {
		return m.Imsi
	}
	return ""
}

func (m *SMS) GetSourceMsisdn() string {
	if m != nil {
		return m.SourceMsisdn
	}
	return ""
}

func (m *SMS) GetMessage() string {
	if m != nil {
		return m.Message
	}
	return ""
}

func (m *SMS) GetCreatedTime() *timestamp.Timestamp {
	if m != nil {
		return m.CreatedTime
	}
	return nil
}

func (m *SMS) GetLastDeliveryAttemptTime() *timestamp.Timestamp {
	if m != nil {
		return m.LastDeliveryAttemptTime
	}
	return nil
}

func (m *SMS) GetAttemptCount() uint32 {
	if m != nil {
		return m.AttemptCount
	}
	return 0
}

func (m *SMS) GetDeliveryError() string {
	if m != nil {
		return m.DeliveryError
	}
	return ""
}

func (m *SMS) GetRefNums() []byte {
	if m != nil {
		return m.RefNums
	}
	return nil
}

// MutableSMS encapsulates the state that service clients are allowed to set.
type MutableSMS struct {
	Imsi                 string   `protobuf:"bytes,1,opt,name=imsi,proto3" json:"imsi,omitempty"`
	SourceMsisdn         string   `protobuf:"bytes,2,opt,name=sourceMsisdn,proto3" json:"sourceMsisdn,omitempty"`
	Message              string   `protobuf:"bytes,3,opt,name=message,proto3" json:"message,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *MutableSMS) Reset()         { *m = MutableSMS{} }
func (m *MutableSMS) String() string { return proto.CompactTextString(m) }
func (*MutableSMS) ProtoMessage()    {}
func (*MutableSMS) Descriptor() ([]byte, []int) {
	return fileDescriptor_0d2c4ccf1453ffdb, []int{1}
}

func (m *MutableSMS) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_MutableSMS.Unmarshal(m, b)
}
func (m *MutableSMS) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_MutableSMS.Marshal(b, m, deterministic)
}
func (m *MutableSMS) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MutableSMS.Merge(m, src)
}
func (m *MutableSMS) XXX_Size() int {
	return xxx_messageInfo_MutableSMS.Size(m)
}
func (m *MutableSMS) XXX_DiscardUnknown() {
	xxx_messageInfo_MutableSMS.DiscardUnknown(m)
}

var xxx_messageInfo_MutableSMS proto.InternalMessageInfo

func (m *MutableSMS) GetImsi() string {
	if m != nil {
		return m.Imsi
	}
	return ""
}

func (m *MutableSMS) GetSourceMsisdn() string {
	if m != nil {
		return m.SourceMsisdn
	}
	return ""
}

func (m *MutableSMS) GetMessage() string {
	if m != nil {
		return m.Message
	}
	return ""
}

func init() {
	proto.RegisterEnum("magma.lte.smsd.storage.MessageStatus", MessageStatus_name, MessageStatus_value)
	proto.RegisterType((*SMS)(nil), "magma.lte.smsd.storage.SMS")
	proto.RegisterType((*MutableSMS)(nil), "magma.lte.smsd.storage.MutableSMS")
}

func init() { proto.RegisterFile("storage.proto", fileDescriptor_0d2c4ccf1453ffdb) }

var fileDescriptor_0d2c4ccf1453ffdb = []byte{
	// 376 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x84, 0x92, 0x4f, 0xab, 0xd3, 0x40,
	0x14, 0xc5, 0x9d, 0x54, 0x5a, 0x7a, 0xd3, 0x3c, 0xca, 0xa0, 0xef, 0x0d, 0x5d, 0x68, 0x28, 0x0a,
	0xc1, 0xc5, 0x14, 0xea, 0xc2, 0x8d, 0x2e, 0xaa, 0x89, 0x12, 0x68, 0xba, 0x48, 0x83, 0x82, 0x0b,
	0x61, 0xda, 0x4c, 0x43, 0x68, 0xa6, 0x13, 0x66, 0x26, 0x82, 0x5f, 0xc6, 0xcf, 0x2a, 0x9d, 0x24,
	0x60, 0x40, 0xfb, 0x76, 0xb9, 0x7f, 0xce, 0x09, 0xbf, 0x73, 0x07, 0x3c, 0x6d, 0xa4, 0x62, 0x05,
	0xa7, 0xb5, 0x92, 0x46, 0xe2, 0x7b, 0xc1, 0x0a, 0xc1, 0x68, 0x65, 0x38, 0xd5, 0x42, 0xe7, 0xb4,
	0x9b, 0x2e, 0x5e, 0x16, 0x52, 0x16, 0x15, 0x5f, 0xd9, 0xad, 0x43, 0x73, 0x5a, 0x99, 0x52, 0x70,
	0x6d, 0x98, 0xa8, 0x5b, 0xe1, 0xf2, 0xf7, 0x08, 0x46, 0xfb, 0x64, 0x8f, 0xef, 0xc0, 0xa9, 0xcf,
	0x04, 0xf9, 0x28, 0x98, 0xa6, 0x4e, 0x7d, 0xc6, 0x1f, 0x60, 0xac, 0x0d, 0x33, 0x8d, 0x26, 0x8e,
	0x8f, 0x82, 0xbb, 0xf5, 0x6b, 0xfa, 0xef, 0x3f, 0xd0, 0x84, 0x6b, 0xcd, 0x0a, 0xbe, 0xb7, 0xcb,
	0x69, 0x27, 0xc2, 0x18, 0x9e, 0x96, 0x42, 0x97, 0x04, 0xac, 0xa1, 0xfd, 0xc6, 0x4b, 0x98, 0x69,
	0xd9, 0xa8, 0x23, 0x4f, 0x74, 0xa9, 0xf3, 0x0b, 0x71, 0xed, 0x6c, 0xd0, 0xc3, 0x04, 0x26, 0xa2,
	0x35, 0x24, 0x33, 0x3b, 0xee, 0x4b, 0xfc, 0x1e, 0xdc, 0xa3, 0xe2, 0xcc, 0xf0, 0x3c, 0x2b, 0x05,
	0x27, 0xcf, 0x7c, 0x14, 0xb8, 0xeb, 0x05, 0x6d, 0xf9, 0x68, 0xcf, 0x47, 0xb3, 0x9e, 0x2f, 0xfd,
	0x7b, 0x1d, 0x67, 0xf0, 0x50, 0x31, 0x6d, 0x42, 0x5e, 0x95, 0x3f, 0xb9, 0xfa, 0xb5, 0x31, 0x86,
	0x8b, 0xda, 0x58, 0xa7, 0xe7, 0x8f, 0x3a, 0xfd, 0x4f, 0x7a, 0x25, 0x62, 0x6d, 0xf9, 0x49, 0x36,
	0x17, 0x43, 0xee, 0x7d, 0x14, 0x78, 0xe9, 0xa0, 0x87, 0x5f, 0x81, 0x97, 0x77, 0xd2, 0x48, 0x29,
	0xa9, 0xc8, 0x83, 0xe5, 0x1a, 0x36, 0xaf, 0xdc, 0x8a, 0x9f, 0x76, 0x8d, 0xd0, 0xe4, 0x85, 0x8f,
	0x82, 0x59, 0xda, 0x97, 0xcb, 0x1f, 0x00, 0x49, 0x63, 0xd8, 0xa1, 0xe2, 0xd7, 0x33, 0xf5, 0xb9,
	0xa2, 0x1b, 0xb9, 0x3a, 0xb7, 0x73, 0x1d, 0x0d, 0x72, 0x7d, 0xf3, 0x0e, 0xbc, 0xc1, 0x09, 0xb1,
	0x0b, 0x93, 0x6f, 0x9b, 0x38, 0x8b, 0x77, 0x5f, 0xe6, 0x4f, 0xb0, 0x07, 0xd3, 0x30, 0xda, 0xc6,
	0x5f, 0xa3, 0x34, 0x0a, 0xe7, 0x08, 0x03, 0x8c, 0x3f, 0x6f, 0xe2, 0x6d, 0x14, 0xce, 0x9d, 0x8f,
	0xd3, 0xef, 0x93, 0xee, 0x0d, 0x1c, 0xc6, 0x36, 0xb4, 0xb7, 0x7f, 0x02, 0x00, 0x00, 0xff, 0xff,
	0xc3, 0x50, 0xc7, 0x85, 0x95, 0x02, 0x00, 0x00,
}
