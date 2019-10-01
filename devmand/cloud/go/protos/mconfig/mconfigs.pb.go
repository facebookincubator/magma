// Code generated by protoc-gen-go. DO NOT EDIT.
// source: mconfigs.proto

package mconfig // import "orc8r/devmand/cloud/go/protos/mconfig"

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

type DevmandGatewayConfig struct {
	ManagedDevices       map[string]*ManagedDevice `protobuf:"bytes,1,rep,name=managed_devices,json=managedDevices,proto3" json:"managed_devices,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
	XXX_NoUnkeyedLiteral struct{}                  `json:"-"`
	XXX_unrecognized     []byte                    `json:"-"`
	XXX_sizecache        int32                     `json:"-"`
}

func (m *DevmandGatewayConfig) Reset()         { *m = DevmandGatewayConfig{} }
func (m *DevmandGatewayConfig) String() string { return proto.CompactTextString(m) }
func (*DevmandGatewayConfig) ProtoMessage()    {}
func (*DevmandGatewayConfig) Descriptor() ([]byte, []int) {
	return fileDescriptor_mconfigs_89aa0844d919344d, []int{0}
}
func (m *DevmandGatewayConfig) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_DevmandGatewayConfig.Unmarshal(m, b)
}
func (m *DevmandGatewayConfig) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_DevmandGatewayConfig.Marshal(b, m, deterministic)
}
func (dst *DevmandGatewayConfig) XXX_Merge(src proto.Message) {
	xxx_messageInfo_DevmandGatewayConfig.Merge(dst, src)
}
func (m *DevmandGatewayConfig) XXX_Size() int {
	return xxx_messageInfo_DevmandGatewayConfig.Size(m)
}
func (m *DevmandGatewayConfig) XXX_DiscardUnknown() {
	xxx_messageInfo_DevmandGatewayConfig.DiscardUnknown(m)
}

var xxx_messageInfo_DevmandGatewayConfig proto.InternalMessageInfo

func (m *DevmandGatewayConfig) GetManagedDevices() map[string]*ManagedDevice {
	if m != nil {
		return m.ManagedDevices
	}
	return nil
}

type ManagedDevice struct {
	// Use any.Any when we are ready. for now use JSON string
	// google.protobuf.Any device_config = 1;
	DeviceConfig         string    `protobuf:"bytes,1,opt,name=device_config,json=deviceConfig,proto3" json:"device_config,omitempty"`
	Host                 string    `protobuf:"bytes,2,opt,name=host,proto3" json:"host,omitempty"`
	DeviceType           []string  `protobuf:"bytes,3,rep,name=device_type,json=deviceType,proto3" json:"device_type,omitempty"`
	Platform             string    `protobuf:"bytes,4,opt,name=platform,proto3" json:"platform,omitempty"`
	Channels             *Channels `protobuf:"bytes,5,opt,name=channels,proto3" json:"channels,omitempty"`
	XXX_NoUnkeyedLiteral struct{}  `json:"-"`
	XXX_unrecognized     []byte    `json:"-"`
	XXX_sizecache        int32     `json:"-"`
}

func (m *ManagedDevice) Reset()         { *m = ManagedDevice{} }
func (m *ManagedDevice) String() string { return proto.CompactTextString(m) }
func (*ManagedDevice) ProtoMessage()    {}
func (*ManagedDevice) Descriptor() ([]byte, []int) {
	return fileDescriptor_mconfigs_89aa0844d919344d, []int{1}
}
func (m *ManagedDevice) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ManagedDevice.Unmarshal(m, b)
}
func (m *ManagedDevice) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ManagedDevice.Marshal(b, m, deterministic)
}
func (dst *ManagedDevice) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ManagedDevice.Merge(dst, src)
}
func (m *ManagedDevice) XXX_Size() int {
	return xxx_messageInfo_ManagedDevice.Size(m)
}
func (m *ManagedDevice) XXX_DiscardUnknown() {
	xxx_messageInfo_ManagedDevice.DiscardUnknown(m)
}

var xxx_messageInfo_ManagedDevice proto.InternalMessageInfo

func (m *ManagedDevice) GetDeviceConfig() string {
	if m != nil {
		return m.DeviceConfig
	}
	return ""
}

func (m *ManagedDevice) GetHost() string {
	if m != nil {
		return m.Host
	}
	return ""
}

func (m *ManagedDevice) GetDeviceType() []string {
	if m != nil {
		return m.DeviceType
	}
	return nil
}

func (m *ManagedDevice) GetPlatform() string {
	if m != nil {
		return m.Platform
	}
	return ""
}

func (m *ManagedDevice) GetChannels() *Channels {
	if m != nil {
		return m.Channels
	}
	return nil
}

type Channels struct {
	SnmpChannel          *SNMPChannel    `protobuf:"bytes,1,opt,name=snmp_channel,json=snmpChannel,proto3" json:"snmp_channel,omitempty"`
	FrinxChannel         *FrinxChannel   `protobuf:"bytes,2,opt,name=frinx_channel,json=frinxChannel,proto3" json:"frinx_channel,omitempty"`
	CambiumChannel       *CambiumChannel `protobuf:"bytes,3,opt,name=cambium_channel,json=cambiumChannel,proto3" json:"cambium_channel,omitempty"`
	OtherChannel         *OtherChannel   `protobuf:"bytes,4,opt,name=other_channel,json=otherChannel,proto3" json:"other_channel,omitempty"`
	XXX_NoUnkeyedLiteral struct{}        `json:"-"`
	XXX_unrecognized     []byte          `json:"-"`
	XXX_sizecache        int32           `json:"-"`
}

func (m *Channels) Reset()         { *m = Channels{} }
func (m *Channels) String() string { return proto.CompactTextString(m) }
func (*Channels) ProtoMessage()    {}
func (*Channels) Descriptor() ([]byte, []int) {
	return fileDescriptor_mconfigs_89aa0844d919344d, []int{2}
}
func (m *Channels) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Channels.Unmarshal(m, b)
}
func (m *Channels) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Channels.Marshal(b, m, deterministic)
}
func (dst *Channels) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Channels.Merge(dst, src)
}
func (m *Channels) XXX_Size() int {
	return xxx_messageInfo_Channels.Size(m)
}
func (m *Channels) XXX_DiscardUnknown() {
	xxx_messageInfo_Channels.DiscardUnknown(m)
}

var xxx_messageInfo_Channels proto.InternalMessageInfo

func (m *Channels) GetSnmpChannel() *SNMPChannel {
	if m != nil {
		return m.SnmpChannel
	}
	return nil
}

func (m *Channels) GetFrinxChannel() *FrinxChannel {
	if m != nil {
		return m.FrinxChannel
	}
	return nil
}

func (m *Channels) GetCambiumChannel() *CambiumChannel {
	if m != nil {
		return m.CambiumChannel
	}
	return nil
}

func (m *Channels) GetOtherChannel() *OtherChannel {
	if m != nil {
		return m.OtherChannel
	}
	return nil
}

type SNMPChannel struct {
	Version              string   `protobuf:"bytes,1,opt,name=version,proto3" json:"version,omitempty"`
	Community            string   `protobuf:"bytes,2,opt,name=community,proto3" json:"community,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *SNMPChannel) Reset()         { *m = SNMPChannel{} }
func (m *SNMPChannel) String() string { return proto.CompactTextString(m) }
func (*SNMPChannel) ProtoMessage()    {}
func (*SNMPChannel) Descriptor() ([]byte, []int) {
	return fileDescriptor_mconfigs_89aa0844d919344d, []int{3}
}
func (m *SNMPChannel) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_SNMPChannel.Unmarshal(m, b)
}
func (m *SNMPChannel) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_SNMPChannel.Marshal(b, m, deterministic)
}
func (dst *SNMPChannel) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SNMPChannel.Merge(dst, src)
}
func (m *SNMPChannel) XXX_Size() int {
	return xxx_messageInfo_SNMPChannel.Size(m)
}
func (m *SNMPChannel) XXX_DiscardUnknown() {
	xxx_messageInfo_SNMPChannel.DiscardUnknown(m)
}

var xxx_messageInfo_SNMPChannel proto.InternalMessageInfo

func (m *SNMPChannel) GetVersion() string {
	if m != nil {
		return m.Version
	}
	return ""
}

func (m *SNMPChannel) GetCommunity() string {
	if m != nil {
		return m.Community
	}
	return ""
}

type FrinxChannel struct {
	FrinxPort            int32    `protobuf:"varint,1,opt,name=frinx_port,json=frinxPort,proto3" json:"frinx_port,omitempty"`
	Authorization        string   `protobuf:"bytes,2,opt,name=authorization,proto3" json:"authorization,omitempty"`
	Host                 string   `protobuf:"bytes,3,opt,name=host,proto3" json:"host,omitempty"`
	Port                 int32    `protobuf:"varint,4,opt,name=port,proto3" json:"port,omitempty"`
	TransportType        string   `protobuf:"bytes,5,opt,name=transport_type,json=transportType,proto3" json:"transport_type,omitempty"`
	DeviceType           string   `protobuf:"bytes,6,opt,name=device_type,json=deviceType,proto3" json:"device_type,omitempty"`
	DeviceVersion        string   `protobuf:"bytes,7,opt,name=device_version,json=deviceVersion,proto3" json:"device_version,omitempty"`
	Username             string   `protobuf:"bytes,8,opt,name=username,proto3" json:"username,omitempty"`
	Password             string   `protobuf:"bytes,9,opt,name=password,proto3" json:"password,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *FrinxChannel) Reset()         { *m = FrinxChannel{} }
func (m *FrinxChannel) String() string { return proto.CompactTextString(m) }
func (*FrinxChannel) ProtoMessage()    {}
func (*FrinxChannel) Descriptor() ([]byte, []int) {
	return fileDescriptor_mconfigs_89aa0844d919344d, []int{4}
}
func (m *FrinxChannel) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_FrinxChannel.Unmarshal(m, b)
}
func (m *FrinxChannel) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_FrinxChannel.Marshal(b, m, deterministic)
}
func (dst *FrinxChannel) XXX_Merge(src proto.Message) {
	xxx_messageInfo_FrinxChannel.Merge(dst, src)
}
func (m *FrinxChannel) XXX_Size() int {
	return xxx_messageInfo_FrinxChannel.Size(m)
}
func (m *FrinxChannel) XXX_DiscardUnknown() {
	xxx_messageInfo_FrinxChannel.DiscardUnknown(m)
}

var xxx_messageInfo_FrinxChannel proto.InternalMessageInfo

func (m *FrinxChannel) GetFrinxPort() int32 {
	if m != nil {
		return m.FrinxPort
	}
	return 0
}

func (m *FrinxChannel) GetAuthorization() string {
	if m != nil {
		return m.Authorization
	}
	return ""
}

func (m *FrinxChannel) GetHost() string {
	if m != nil {
		return m.Host
	}
	return ""
}

func (m *FrinxChannel) GetPort() int32 {
	if m != nil {
		return m.Port
	}
	return 0
}

func (m *FrinxChannel) GetTransportType() string {
	if m != nil {
		return m.TransportType
	}
	return ""
}

func (m *FrinxChannel) GetDeviceType() string {
	if m != nil {
		return m.DeviceType
	}
	return ""
}

func (m *FrinxChannel) GetDeviceVersion() string {
	if m != nil {
		return m.DeviceVersion
	}
	return ""
}

func (m *FrinxChannel) GetUsername() string {
	if m != nil {
		return m.Username
	}
	return ""
}

func (m *FrinxChannel) GetPassword() string {
	if m != nil {
		return m.Password
	}
	return ""
}

type CambiumChannel struct {
	ClientId             string   `protobuf:"bytes,1,opt,name=client_id,json=clientId,proto3" json:"client_id,omitempty"`
	ClientSecret         string   `protobuf:"bytes,2,opt,name=client_secret,json=clientSecret,proto3" json:"client_secret,omitempty"`
	ClientMac            string   `protobuf:"bytes,3,opt,name=client_mac,json=clientMac,proto3" json:"client_mac,omitempty"`
	ClientIp             string   `protobuf:"bytes,4,opt,name=client_ip,json=clientIp,proto3" json:"client_ip,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *CambiumChannel) Reset()         { *m = CambiumChannel{} }
func (m *CambiumChannel) String() string { return proto.CompactTextString(m) }
func (*CambiumChannel) ProtoMessage()    {}
func (*CambiumChannel) Descriptor() ([]byte, []int) {
	return fileDescriptor_mconfigs_89aa0844d919344d, []int{5}
}
func (m *CambiumChannel) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CambiumChannel.Unmarshal(m, b)
}
func (m *CambiumChannel) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CambiumChannel.Marshal(b, m, deterministic)
}
func (dst *CambiumChannel) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CambiumChannel.Merge(dst, src)
}
func (m *CambiumChannel) XXX_Size() int {
	return xxx_messageInfo_CambiumChannel.Size(m)
}
func (m *CambiumChannel) XXX_DiscardUnknown() {
	xxx_messageInfo_CambiumChannel.DiscardUnknown(m)
}

var xxx_messageInfo_CambiumChannel proto.InternalMessageInfo

func (m *CambiumChannel) GetClientId() string {
	if m != nil {
		return m.ClientId
	}
	return ""
}

func (m *CambiumChannel) GetClientSecret() string {
	if m != nil {
		return m.ClientSecret
	}
	return ""
}

func (m *CambiumChannel) GetClientMac() string {
	if m != nil {
		return m.ClientMac
	}
	return ""
}

func (m *CambiumChannel) GetClientIp() string {
	if m != nil {
		return m.ClientIp
	}
	return ""
}

type OtherChannel struct {
	ChannelProps         map[string]string `protobuf:"bytes,1,rep,name=channel_props,json=channelProps,proto3" json:"channel_props,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
	XXX_NoUnkeyedLiteral struct{}          `json:"-"`
	XXX_unrecognized     []byte            `json:"-"`
	XXX_sizecache        int32             `json:"-"`
}

func (m *OtherChannel) Reset()         { *m = OtherChannel{} }
func (m *OtherChannel) String() string { return proto.CompactTextString(m) }
func (*OtherChannel) ProtoMessage()    {}
func (*OtherChannel) Descriptor() ([]byte, []int) {
	return fileDescriptor_mconfigs_89aa0844d919344d, []int{6}
}
func (m *OtherChannel) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_OtherChannel.Unmarshal(m, b)
}
func (m *OtherChannel) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_OtherChannel.Marshal(b, m, deterministic)
}
func (dst *OtherChannel) XXX_Merge(src proto.Message) {
	xxx_messageInfo_OtherChannel.Merge(dst, src)
}
func (m *OtherChannel) XXX_Size() int {
	return xxx_messageInfo_OtherChannel.Size(m)
}
func (m *OtherChannel) XXX_DiscardUnknown() {
	xxx_messageInfo_OtherChannel.DiscardUnknown(m)
}

var xxx_messageInfo_OtherChannel proto.InternalMessageInfo

func (m *OtherChannel) GetChannelProps() map[string]string {
	if m != nil {
		return m.ChannelProps
	}
	return nil
}

func init() {
	proto.RegisterType((*DevmandGatewayConfig)(nil), "magma.mconfig.DevmandGatewayConfig")
	proto.RegisterMapType((map[string]*ManagedDevice)(nil), "magma.mconfig.DevmandGatewayConfig.ManagedDevicesEntry")
	proto.RegisterType((*ManagedDevice)(nil), "magma.mconfig.ManagedDevice")
	proto.RegisterType((*Channels)(nil), "magma.mconfig.Channels")
	proto.RegisterType((*SNMPChannel)(nil), "magma.mconfig.SNMPChannel")
	proto.RegisterType((*FrinxChannel)(nil), "magma.mconfig.FrinxChannel")
	proto.RegisterType((*CambiumChannel)(nil), "magma.mconfig.CambiumChannel")
	proto.RegisterType((*OtherChannel)(nil), "magma.mconfig.OtherChannel")
	proto.RegisterMapType((map[string]string)(nil), "magma.mconfig.OtherChannel.ChannelPropsEntry")
}

func init() { proto.RegisterFile("mconfigs.proto", fileDescriptor_mconfigs_89aa0844d919344d) }

var fileDescriptor_mconfigs_89aa0844d919344d = []byte{
	// 665 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x7c, 0x54, 0xcd, 0x6e, 0xd3, 0x4a,
	0x14, 0x96, 0xf3, 0xd3, 0xc6, 0x27, 0x4e, 0x7a, 0xef, 0xdc, 0x4a, 0xd7, 0x4a, 0x5b, 0x11, 0x19,
	0x2a, 0xb2, 0x21, 0x91, 0xd2, 0x05, 0x15, 0x12, 0x02, 0xd1, 0x1f, 0xc4, 0xa2, 0x50, 0xb9, 0x88,
	0x05, 0x9b, 0x30, 0x1d, 0x4f, 0x12, 0x8b, 0xcc, 0x8c, 0x35, 0x9e, 0xa4, 0x84, 0x87, 0x60, 0xc7,
	0x0b, 0xc0, 0x5b, 0xf0, 0x26, 0xbc, 0x0d, 0xf2, 0xcc, 0xd8, 0xd8, 0x21, 0x62, 0xe7, 0xf3, 0xcd,
	0x77, 0xbe, 0x39, 0xe7, 0xf8, 0x9b, 0x03, 0x5d, 0x46, 0x04, 0x9f, 0xc6, 0xb3, 0x74, 0x98, 0x48,
	0xa1, 0x04, 0xea, 0x30, 0x3c, 0x63, 0x78, 0x68, 0xd1, 0xe0, 0xa7, 0x03, 0xfb, 0xe7, 0x74, 0xc5,
	0x30, 0x8f, 0x5e, 0x62, 0x45, 0xef, 0xf0, 0xfa, 0x4c, 0x1f, 0xa0, 0x0f, 0xb0, 0xc7, 0x30, 0xc7,
	0x33, 0x1a, 0x4d, 0x22, 0xba, 0x8a, 0x09, 0x4d, 0x7d, 0xa7, 0x5f, 0x1f, 0xb4, 0xc7, 0x8f, 0x87,
	0x15, 0x85, 0xe1, 0xb6, 0xec, 0xe1, 0x95, 0x49, 0x3d, 0x37, 0x99, 0x17, 0x5c, 0xc9, 0x75, 0xd8,
	0x65, 0x15, 0xb0, 0x37, 0x81, 0xff, 0xb6, 0xd0, 0xd0, 0x3f, 0x50, 0xff, 0x48, 0xd7, 0xbe, 0xd3,
	0x77, 0x06, 0x6e, 0x98, 0x7d, 0xa2, 0x31, 0x34, 0x57, 0x78, 0xb1, 0xa4, 0x7e, 0xad, 0xef, 0x0c,
	0xda, 0xe3, 0xc3, 0x8d, 0x02, 0x2a, 0x22, 0xa1, 0xa1, 0x3e, 0xa9, 0x9d, 0x3a, 0xc1, 0x0f, 0x07,
	0x3a, 0x95, 0x43, 0x74, 0x1f, 0x3a, 0xa6, 0x99, 0x89, 0xc9, 0xb5, 0xb7, 0x78, 0x06, 0xb4, 0x9d,
	0x23, 0x68, 0xcc, 0x45, 0xaa, 0xf4, 0x6d, 0x6e, 0xa8, 0xbf, 0xd1, 0x3d, 0x68, 0xdb, 0x44, 0xb5,
	0x4e, 0xa8, 0x5f, 0xef, 0xd7, 0x07, 0x6e, 0x08, 0x06, 0x7a, 0xbb, 0x4e, 0x28, 0xea, 0x41, 0x2b,
	0x59, 0x60, 0x35, 0x15, 0x92, 0xf9, 0x0d, 0x9d, 0x58, 0xc4, 0xe8, 0x04, 0x5a, 0x64, 0x8e, 0x39,
	0xa7, 0x8b, 0xd4, 0x6f, 0xea, 0x16, 0xfe, 0xdf, 0x68, 0xe1, 0xcc, 0x1e, 0x87, 0x05, 0x31, 0xf8,
	0x5a, 0x83, 0x56, 0x0e, 0xa3, 0xa7, 0xe0, 0xa5, 0x9c, 0x25, 0x13, 0x7b, 0xaa, 0xcb, 0x6e, 0x8f,
	0x7b, 0x1b, 0x2a, 0x37, 0xaf, 0xaf, 0xae, 0x6d, 0x4a, 0xd8, 0xce, 0xf8, 0x36, 0x40, 0xcf, 0xa1,
	0x33, 0x95, 0x31, 0xff, 0x54, 0xe4, 0x9b, 0x41, 0x1e, 0x6c, 0xe4, 0x5f, 0x66, 0x9c, 0x5c, 0xc0,
	0x9b, 0x96, 0x22, 0x74, 0x09, 0x7b, 0x04, 0xb3, 0xdb, 0x78, 0xc9, 0x0a, 0x8d, 0xba, 0xd6, 0x38,
	0xda, 0xec, 0xc4, 0xb0, 0x72, 0x95, 0x2e, 0xa9, 0xc4, 0x59, 0x25, 0x42, 0xcd, 0xa9, 0x2c, 0x54,
	0x1a, 0x5b, 0x2b, 0x79, 0x93, 0x71, 0x8a, 0x4a, 0x44, 0x29, 0x0a, 0x2e, 0xa0, 0x5d, 0xea, 0x13,
	0xf9, 0xb0, 0xbb, 0xa2, 0x32, 0x8d, 0x05, 0xb7, 0xff, 0x32, 0x0f, 0xd1, 0x21, 0xb8, 0x44, 0x30,
	0xb6, 0xe4, 0xb1, 0x5a, 0xdb, 0x7f, 0xf9, 0x1b, 0x08, 0xbe, 0xd5, 0xc0, 0x2b, 0xf7, 0x8b, 0x8e,
	0x00, 0xcc, 0x8c, 0x12, 0x21, 0x95, 0xd6, 0x6a, 0x86, 0xae, 0x46, 0xae, 0x85, 0x54, 0xe8, 0x01,
	0x74, 0xf0, 0x52, 0xcd, 0x85, 0x8c, 0x3f, 0x63, 0x95, 0xdd, 0x66, 0x14, 0xab, 0x60, 0x61, 0x9d,
	0x7a, 0xc9, 0x3a, 0x08, 0x1a, 0x5a, 0xb2, 0xa1, 0x25, 0xf5, 0x37, 0x3a, 0x86, 0xae, 0x92, 0x98,
	0xa7, 0x59, 0x60, 0x1c, 0xd5, 0x34, 0x72, 0x05, 0xaa, 0x4d, 0xb5, 0xe1, 0xba, 0x1d, 0xcd, 0x29,
	0xbb, 0xee, 0x18, 0xba, 0x96, 0x90, 0x0f, 0x61, 0xd7, 0xe8, 0x18, 0xf4, 0x9d, 0x1d, 0x45, 0x0f,
	0x5a, 0xcb, 0x94, 0x4a, 0x8e, 0x19, 0xf5, 0x5b, 0xc6, 0x9c, 0x79, 0xac, 0x8d, 0x8b, 0xd3, 0xf4,
	0x4e, 0xc8, 0xc8, 0x77, 0xad, 0x71, 0x6d, 0x1c, 0x7c, 0x71, 0xa0, 0x5b, 0xfd, 0xa1, 0xe8, 0x00,
	0x5c, 0xb2, 0x88, 0x29, 0x57, 0x93, 0x38, 0xb2, 0x13, 0x6f, 0x19, 0xe0, 0x55, 0x94, 0x3d, 0x2f,
	0x7b, 0x98, 0x52, 0x22, 0x69, 0xfe, 0x84, 0x3c, 0x03, 0xde, 0x68, 0x2c, 0x1b, 0xb4, 0x25, 0x31,
	0x4c, 0xec, 0xa4, 0xac, 0xe6, 0x15, 0x26, 0xe5, 0x0b, 0x92, 0xfc, 0x25, 0xd9, 0x0b, 0x92, 0xe0,
	0xbb, 0x03, 0x5e, 0xd9, 0x1b, 0x28, 0x84, 0x8e, 0x75, 0xd2, 0x24, 0x91, 0x22, 0xc9, 0x77, 0xd4,
	0xa3, 0xbf, 0xf8, 0x29, 0x7f, 0x6c, 0xd7, 0x19, 0xdf, 0x6c, 0x26, 0x8f, 0x94, 0xa0, 0xde, 0x33,
	0xf8, 0xf7, 0x0f, 0xca, 0x96, 0xad, 0xb4, 0x5f, 0xde, 0x4a, 0x6e, 0x69, 0xef, 0xbc, 0x78, 0xf8,
	0xfe, 0x58, 0x48, 0x72, 0x2a, 0x47, 0x91, 0x59, 0x8d, 0x23, 0xb2, 0x10, 0xcb, 0x68, 0x34, 0x13,
	0x23, 0xbd, 0x82, 0xd3, 0x91, 0x2d, 0xeb, 0x76, 0x47, 0xc7, 0x27, 0xbf, 0x02, 0x00, 0x00, 0xff,
	0xff, 0xa6, 0xe7, 0x72, 0x51, 0xa4, 0x05, 0x00, 0x00,
}
