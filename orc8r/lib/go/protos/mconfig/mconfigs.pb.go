// Code generated by protoc-gen-go. DO NOT EDIT.
// source: orc8r/protos/mconfig/mconfigs.proto

package mconfig

import (
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
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

//------------------------------------------------------------------------------
// Control Proxy configs
//------------------------------------------------------------------------------
type ControlProxy struct {
	LogLevel             protos.LogLevel `protobuf:"varint,1,opt,name=log_level,json=logLevel,proto3,enum=magma.orc8r.LogLevel" json:"log_level,omitempty"`
	XXX_NoUnkeyedLiteral struct{}        `json:"-"`
	XXX_unrecognized     []byte          `json:"-"`
	XXX_sizecache        int32           `json:"-"`
}

func (m *ControlProxy) Reset()         { *m = ControlProxy{} }
func (m *ControlProxy) String() string { return proto.CompactTextString(m) }
func (*ControlProxy) ProtoMessage()    {}
func (*ControlProxy) Descriptor() ([]byte, []int) {
	return fileDescriptor_9618f358f05ec5b9, []int{0}
}

func (m *ControlProxy) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ControlProxy.Unmarshal(m, b)
}
func (m *ControlProxy) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ControlProxy.Marshal(b, m, deterministic)
}
func (m *ControlProxy) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ControlProxy.Merge(m, src)
}
func (m *ControlProxy) XXX_Size() int {
	return xxx_messageInfo_ControlProxy.Size(m)
}
func (m *ControlProxy) XXX_DiscardUnknown() {
	xxx_messageInfo_ControlProxy.DiscardUnknown(m)
}

var xxx_messageInfo_ControlProxy proto.InternalMessageInfo

func (m *ControlProxy) GetLogLevel() protos.LogLevel {
	if m != nil {
		return m.LogLevel
	}
	return protos.LogLevel_DEBUG
}

//------------------------------------------------------------------------------
// DnsD configs
//------------------------------------------------------------------------------
type DnsD struct {
	LogLevel             protos.LogLevel                 `protobuf:"varint,1,opt,name=log_level,json=logLevel,proto3,enum=magma.orc8r.LogLevel" json:"log_level,omitempty"`
	EnableCaching        bool                            `protobuf:"varint,2,opt,name=enable_caching,json=enableCaching,proto3" json:"enable_caching,omitempty"`
	LocalTTL             int32                           `protobuf:"varint,3,opt,name=localTTL,proto3" json:"localTTL,omitempty"`
	Records              []*NetworkDNSConfigRecordsItems `protobuf:"bytes,4,rep,name=records,proto3" json:"records,omitempty"`
	DhcpEnabled          bool                            `protobuf:"varint,5,opt,name=dhcp_enabled,json=dhcpEnabled,proto3" json:"dhcp_enabled,omitempty"`
	XXX_NoUnkeyedLiteral struct{}                        `json:"-"`
	XXX_unrecognized     []byte                          `json:"-"`
	XXX_sizecache        int32                           `json:"-"`
}

func (m *DnsD) Reset()         { *m = DnsD{} }
func (m *DnsD) String() string { return proto.CompactTextString(m) }
func (*DnsD) ProtoMessage()    {}
func (*DnsD) Descriptor() ([]byte, []int) {
	return fileDescriptor_9618f358f05ec5b9, []int{1}
}

func (m *DnsD) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_DnsD.Unmarshal(m, b)
}
func (m *DnsD) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_DnsD.Marshal(b, m, deterministic)
}
func (m *DnsD) XXX_Merge(src proto.Message) {
	xxx_messageInfo_DnsD.Merge(m, src)
}
func (m *DnsD) XXX_Size() int {
	return xxx_messageInfo_DnsD.Size(m)
}
func (m *DnsD) XXX_DiscardUnknown() {
	xxx_messageInfo_DnsD.DiscardUnknown(m)
}

var xxx_messageInfo_DnsD proto.InternalMessageInfo

func (m *DnsD) GetLogLevel() protos.LogLevel {
	if m != nil {
		return m.LogLevel
	}
	return protos.LogLevel_DEBUG
}

func (m *DnsD) GetEnableCaching() bool {
	if m != nil {
		return m.EnableCaching
	}
	return false
}

func (m *DnsD) GetLocalTTL() int32 {
	if m != nil {
		return m.LocalTTL
	}
	return 0
}

func (m *DnsD) GetRecords() []*NetworkDNSConfigRecordsItems {
	if m != nil {
		return m.Records
	}
	return nil
}

func (m *DnsD) GetDhcpEnabled() bool {
	if m != nil {
		return m.DhcpEnabled
	}
	return false
}

type NetworkDNSConfigRecordsItems struct {
	ARecord              []string `protobuf:"bytes,1,rep,name=a_record,json=aRecord,proto3" json:"a_record,omitempty"`
	AaaaRecord           []string `protobuf:"bytes,2,rep,name=aaaa_record,json=aaaaRecord,proto3" json:"aaaa_record,omitempty"`
	CnameRecord          []string `protobuf:"bytes,3,rep,name=cname_record,json=cnameRecord,proto3" json:"cname_record,omitempty"`
	Domain               string   `protobuf:"bytes,4,opt,name=domain,proto3" json:"domain,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *NetworkDNSConfigRecordsItems) Reset()         { *m = NetworkDNSConfigRecordsItems{} }
func (m *NetworkDNSConfigRecordsItems) String() string { return proto.CompactTextString(m) }
func (*NetworkDNSConfigRecordsItems) ProtoMessage()    {}
func (*NetworkDNSConfigRecordsItems) Descriptor() ([]byte, []int) {
	return fileDescriptor_9618f358f05ec5b9, []int{2}
}

func (m *NetworkDNSConfigRecordsItems) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_NetworkDNSConfigRecordsItems.Unmarshal(m, b)
}
func (m *NetworkDNSConfigRecordsItems) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_NetworkDNSConfigRecordsItems.Marshal(b, m, deterministic)
}
func (m *NetworkDNSConfigRecordsItems) XXX_Merge(src proto.Message) {
	xxx_messageInfo_NetworkDNSConfigRecordsItems.Merge(m, src)
}
func (m *NetworkDNSConfigRecordsItems) XXX_Size() int {
	return xxx_messageInfo_NetworkDNSConfigRecordsItems.Size(m)
}
func (m *NetworkDNSConfigRecordsItems) XXX_DiscardUnknown() {
	xxx_messageInfo_NetworkDNSConfigRecordsItems.DiscardUnknown(m)
}

var xxx_messageInfo_NetworkDNSConfigRecordsItems proto.InternalMessageInfo

func (m *NetworkDNSConfigRecordsItems) GetARecord() []string {
	if m != nil {
		return m.ARecord
	}
	return nil
}

func (m *NetworkDNSConfigRecordsItems) GetAaaaRecord() []string {
	if m != nil {
		return m.AaaaRecord
	}
	return nil
}

func (m *NetworkDNSConfigRecordsItems) GetCnameRecord() []string {
	if m != nil {
		return m.CnameRecord
	}
	return nil
}

func (m *NetworkDNSConfigRecordsItems) GetDomain() string {
	if m != nil {
		return m.Domain
	}
	return ""
}

type ImageSpec struct {
	Name                 string   `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	Order                int64    `protobuf:"varint,2,opt,name=order,proto3" json:"order,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ImageSpec) Reset()         { *m = ImageSpec{} }
func (m *ImageSpec) String() string { return proto.CompactTextString(m) }
func (*ImageSpec) ProtoMessage()    {}
func (*ImageSpec) Descriptor() ([]byte, []int) {
	return fileDescriptor_9618f358f05ec5b9, []int{3}
}

func (m *ImageSpec) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ImageSpec.Unmarshal(m, b)
}
func (m *ImageSpec) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ImageSpec.Marshal(b, m, deterministic)
}
func (m *ImageSpec) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ImageSpec.Merge(m, src)
}
func (m *ImageSpec) XXX_Size() int {
	return xxx_messageInfo_ImageSpec.Size(m)
}
func (m *ImageSpec) XXX_DiscardUnknown() {
	xxx_messageInfo_ImageSpec.DiscardUnknown(m)
}

var xxx_messageInfo_ImageSpec proto.InternalMessageInfo

func (m *ImageSpec) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *ImageSpec) GetOrder() int64 {
	if m != nil {
		return m.Order
	}
	return 0
}

type MagmaD struct {
	LogLevel protos.LogLevel `protobuf:"varint,1,opt,name=log_level,json=logLevel,proto3,enum=magma.orc8r.LogLevel" json:"log_level,omitempty"`
	// Interval for the gateways to send checkin rpc calls to the cloud.
	CheckinInterval int32 `protobuf:"varint,2,opt,name=checkin_interval,json=checkinInterval,proto3" json:"checkin_interval,omitempty"`
	// Checkin rpc timeout
	CheckinTimeout int32 `protobuf:"varint,3,opt,name=checkin_timeout,json=checkinTimeout,proto3" json:"checkin_timeout,omitempty"`
	// Enables autoupgrading of the magma package
	AutoupgradeEnabled bool `protobuf:"varint,4,opt,name=autoupgrade_enabled,json=autoupgradeEnabled,proto3" json:"autoupgrade_enabled,omitempty"`
	// Interval to poll for package upgrades
	AutoupgradePollInterval int32 `protobuf:"varint,5,opt,name=autoupgrade_poll_interval,json=autoupgradePollInterval,proto3" json:"autoupgrade_poll_interval,omitempty"`
	// The magma package version the gateway should upgrade to
	PackageVersion string `protobuf:"bytes,6,opt,name=package_version,json=packageVersion,proto3" json:"package_version,omitempty"`
	// List of upgrade images
	Images []*ImageSpec `protobuf:"bytes,7,rep,name=images,proto3" json:"images,omitempty"`
	// For streamer, should be left unused by gateway
	TierId       string          `protobuf:"bytes,8,opt,name=tier_id,json=tierId,proto3" json:"tier_id,omitempty"`
	FeatureFlags map[string]bool `protobuf:"bytes,9,rep,name=feature_flags,json=featureFlags,proto3" json:"feature_flags,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"varint,2,opt,name=value,proto3"`
	// List of dynamic_services
	DynamicServices      []string `protobuf:"bytes,10,rep,name=dynamic_services,json=dynamicServices,proto3" json:"dynamic_services,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *MagmaD) Reset()         { *m = MagmaD{} }
func (m *MagmaD) String() string { return proto.CompactTextString(m) }
func (*MagmaD) ProtoMessage()    {}
func (*MagmaD) Descriptor() ([]byte, []int) {
	return fileDescriptor_9618f358f05ec5b9, []int{4}
}

func (m *MagmaD) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_MagmaD.Unmarshal(m, b)
}
func (m *MagmaD) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_MagmaD.Marshal(b, m, deterministic)
}
func (m *MagmaD) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MagmaD.Merge(m, src)
}
func (m *MagmaD) XXX_Size() int {
	return xxx_messageInfo_MagmaD.Size(m)
}
func (m *MagmaD) XXX_DiscardUnknown() {
	xxx_messageInfo_MagmaD.DiscardUnknown(m)
}

var xxx_messageInfo_MagmaD proto.InternalMessageInfo

func (m *MagmaD) GetLogLevel() protos.LogLevel {
	if m != nil {
		return m.LogLevel
	}
	return protos.LogLevel_DEBUG
}

func (m *MagmaD) GetCheckinInterval() int32 {
	if m != nil {
		return m.CheckinInterval
	}
	return 0
}

func (m *MagmaD) GetCheckinTimeout() int32 {
	if m != nil {
		return m.CheckinTimeout
	}
	return 0
}

func (m *MagmaD) GetAutoupgradeEnabled() bool {
	if m != nil {
		return m.AutoupgradeEnabled
	}
	return false
}

func (m *MagmaD) GetAutoupgradePollInterval() int32 {
	if m != nil {
		return m.AutoupgradePollInterval
	}
	return 0
}

func (m *MagmaD) GetPackageVersion() string {
	if m != nil {
		return m.PackageVersion
	}
	return ""
}

func (m *MagmaD) GetImages() []*ImageSpec {
	if m != nil {
		return m.Images
	}
	return nil
}

func (m *MagmaD) GetTierId() string {
	if m != nil {
		return m.TierId
	}
	return ""
}

func (m *MagmaD) GetFeatureFlags() map[string]bool {
	if m != nil {
		return m.FeatureFlags
	}
	return nil
}

func (m *MagmaD) GetDynamicServices() []string {
	if m != nil {
		return m.DynamicServices
	}
	return nil
}

//------------------------------------------------------------------------------
// EventD configs
//------------------------------------------------------------------------------
type EventD struct {
	LogLevel protos.LogLevel `protobuf:"varint,1,opt,name=log_level,json=logLevel,proto3,enum=magma.orc8r.LogLevel" json:"log_level,omitempty"`
	// The verbosity level for events.
	// All events less than or equal to this verbosity will be logged.
	EventVerbosity       int32    `protobuf:"varint,2,opt,name=event_verbosity,json=eventVerbosity,proto3" json:"event_verbosity,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *EventD) Reset()         { *m = EventD{} }
func (m *EventD) String() string { return proto.CompactTextString(m) }
func (*EventD) ProtoMessage()    {}
func (*EventD) Descriptor() ([]byte, []int) {
	return fileDescriptor_9618f358f05ec5b9, []int{5}
}

func (m *EventD) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_EventD.Unmarshal(m, b)
}
func (m *EventD) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_EventD.Marshal(b, m, deterministic)
}
func (m *EventD) XXX_Merge(src proto.Message) {
	xxx_messageInfo_EventD.Merge(m, src)
}
func (m *EventD) XXX_Size() int {
	return xxx_messageInfo_EventD.Size(m)
}
func (m *EventD) XXX_DiscardUnknown() {
	xxx_messageInfo_EventD.DiscardUnknown(m)
}

var xxx_messageInfo_EventD proto.InternalMessageInfo

func (m *EventD) GetLogLevel() protos.LogLevel {
	if m != nil {
		return m.LogLevel
	}
	return protos.LogLevel_DEBUG
}

func (m *EventD) GetEventVerbosity() int32 {
	if m != nil {
		return m.EventVerbosity
	}
	return 0
}

type DirectoryD struct {
	LogLevel             protos.LogLevel `protobuf:"varint,1,opt,name=log_level,json=logLevel,proto3,enum=magma.orc8r.LogLevel" json:"log_level,omitempty"`
	XXX_NoUnkeyedLiteral struct{}        `json:"-"`
	XXX_unrecognized     []byte          `json:"-"`
	XXX_sizecache        int32           `json:"-"`
}

func (m *DirectoryD) Reset()         { *m = DirectoryD{} }
func (m *DirectoryD) String() string { return proto.CompactTextString(m) }
func (*DirectoryD) ProtoMessage()    {}
func (*DirectoryD) Descriptor() ([]byte, []int) {
	return fileDescriptor_9618f358f05ec5b9, []int{6}
}

func (m *DirectoryD) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_DirectoryD.Unmarshal(m, b)
}
func (m *DirectoryD) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_DirectoryD.Marshal(b, m, deterministic)
}
func (m *DirectoryD) XXX_Merge(src proto.Message) {
	xxx_messageInfo_DirectoryD.Merge(m, src)
}
func (m *DirectoryD) XXX_Size() int {
	return xxx_messageInfo_DirectoryD.Size(m)
}
func (m *DirectoryD) XXX_DiscardUnknown() {
	xxx_messageInfo_DirectoryD.DiscardUnknown(m)
}

var xxx_messageInfo_DirectoryD proto.InternalMessageInfo

func (m *DirectoryD) GetLogLevel() protos.LogLevel {
	if m != nil {
		return m.LogLevel
	}
	return protos.LogLevel_DEBUG
}

//------------------------------------------------------------------------------
// MetricsD configs
//------------------------------------------------------------------------------
type MetricsD struct {
	LogLevel             protos.LogLevel `protobuf:"varint,1,opt,name=log_level,json=logLevel,proto3,enum=magma.orc8r.LogLevel" json:"log_level,omitempty"`
	XXX_NoUnkeyedLiteral struct{}        `json:"-"`
	XXX_unrecognized     []byte          `json:"-"`
	XXX_sizecache        int32           `json:"-"`
}

func (m *MetricsD) Reset()         { *m = MetricsD{} }
func (m *MetricsD) String() string { return proto.CompactTextString(m) }
func (*MetricsD) ProtoMessage()    {}
func (*MetricsD) Descriptor() ([]byte, []int) {
	return fileDescriptor_9618f358f05ec5b9, []int{7}
}

func (m *MetricsD) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_MetricsD.Unmarshal(m, b)
}
func (m *MetricsD) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_MetricsD.Marshal(b, m, deterministic)
}
func (m *MetricsD) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MetricsD.Merge(m, src)
}
func (m *MetricsD) XXX_Size() int {
	return xxx_messageInfo_MetricsD.Size(m)
}
func (m *MetricsD) XXX_DiscardUnknown() {
	xxx_messageInfo_MetricsD.DiscardUnknown(m)
}

var xxx_messageInfo_MetricsD proto.InternalMessageInfo

func (m *MetricsD) GetLogLevel() protos.LogLevel {
	if m != nil {
		return m.LogLevel
	}
	return protos.LogLevel_DEBUG
}

//------------------------------------------------------------------------------
// State configs
//------------------------------------------------------------------------------
type State struct {
	LogLevel             protos.LogLevel `protobuf:"varint,1,opt,name=log_level,json=logLevel,proto3,enum=magma.orc8r.LogLevel" json:"log_level,omitempty"`
	XXX_NoUnkeyedLiteral struct{}        `json:"-"`
	XXX_unrecognized     []byte          `json:"-"`
	XXX_sizecache        int32           `json:"-"`
}

func (m *State) Reset()         { *m = State{} }
func (m *State) String() string { return proto.CompactTextString(m) }
func (*State) ProtoMessage()    {}
func (*State) Descriptor() ([]byte, []int) {
	return fileDescriptor_9618f358f05ec5b9, []int{8}
}

func (m *State) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_State.Unmarshal(m, b)
}
func (m *State) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_State.Marshal(b, m, deterministic)
}
func (m *State) XXX_Merge(src proto.Message) {
	xxx_messageInfo_State.Merge(m, src)
}
func (m *State) XXX_Size() int {
	return xxx_messageInfo_State.Size(m)
}
func (m *State) XXX_DiscardUnknown() {
	xxx_messageInfo_State.DiscardUnknown(m)
}

var xxx_messageInfo_State proto.InternalMessageInfo

func (m *State) GetLogLevel() protos.LogLevel {
	if m != nil {
		return m.LogLevel
	}
	return protos.LogLevel_DEBUG
}

//------------------------------------------------------------------------------
// Fluent Bit configs
//------------------------------------------------------------------------------
type FluentBit struct {
	ExtraTags            map[string]string `protobuf:"bytes,1,rep,name=extra_tags,json=extraTags,proto3" json:"extra_tags,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
	ThrottleRate         uint32            `protobuf:"varint,10,opt,name=throttle_rate,json=throttleRate,proto3" json:"throttle_rate,omitempty"`
	ThrottleWindow       uint32            `protobuf:"varint,11,opt,name=throttle_window,json=throttleWindow,proto3" json:"throttle_window,omitempty"`
	ThrottleInterval     string            `protobuf:"bytes,12,opt,name=throttle_interval,json=throttleInterval,proto3" json:"throttle_interval,omitempty"`
	FilesByTag           map[string]string `protobuf:"bytes,20,rep,name=files_by_tag,json=filesByTag,proto3" json:"files_by_tag,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
	XXX_NoUnkeyedLiteral struct{}          `json:"-"`
	XXX_unrecognized     []byte            `json:"-"`
	XXX_sizecache        int32             `json:"-"`
}

func (m *FluentBit) Reset()         { *m = FluentBit{} }
func (m *FluentBit) String() string { return proto.CompactTextString(m) }
func (*FluentBit) ProtoMessage()    {}
func (*FluentBit) Descriptor() ([]byte, []int) {
	return fileDescriptor_9618f358f05ec5b9, []int{9}
}

func (m *FluentBit) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_FluentBit.Unmarshal(m, b)
}
func (m *FluentBit) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_FluentBit.Marshal(b, m, deterministic)
}
func (m *FluentBit) XXX_Merge(src proto.Message) {
	xxx_messageInfo_FluentBit.Merge(m, src)
}
func (m *FluentBit) XXX_Size() int {
	return xxx_messageInfo_FluentBit.Size(m)
}
func (m *FluentBit) XXX_DiscardUnknown() {
	xxx_messageInfo_FluentBit.DiscardUnknown(m)
}

var xxx_messageInfo_FluentBit proto.InternalMessageInfo

func (m *FluentBit) GetExtraTags() map[string]string {
	if m != nil {
		return m.ExtraTags
	}
	return nil
}

func (m *FluentBit) GetThrottleRate() uint32 {
	if m != nil {
		return m.ThrottleRate
	}
	return 0
}

func (m *FluentBit) GetThrottleWindow() uint32 {
	if m != nil {
		return m.ThrottleWindow
	}
	return 0
}

func (m *FluentBit) GetThrottleInterval() string {
	if m != nil {
		return m.ThrottleInterval
	}
	return ""
}

func (m *FluentBit) GetFilesByTag() map[string]string {
	if m != nil {
		return m.FilesByTag
	}
	return nil
}

func init() {
	proto.RegisterType((*ControlProxy)(nil), "magma.mconfig.ControlProxy")
	proto.RegisterType((*DnsD)(nil), "magma.mconfig.DnsD")
	proto.RegisterType((*NetworkDNSConfigRecordsItems)(nil), "magma.mconfig.NetworkDNSConfigRecordsItems")
	proto.RegisterType((*ImageSpec)(nil), "magma.mconfig.ImageSpec")
	proto.RegisterType((*MagmaD)(nil), "magma.mconfig.MagmaD")
	proto.RegisterMapType((map[string]bool)(nil), "magma.mconfig.MagmaD.FeatureFlagsEntry")
	proto.RegisterType((*EventD)(nil), "magma.mconfig.EventD")
	proto.RegisterType((*DirectoryD)(nil), "magma.mconfig.DirectoryD")
	proto.RegisterType((*MetricsD)(nil), "magma.mconfig.MetricsD")
	proto.RegisterType((*State)(nil), "magma.mconfig.State")
	proto.RegisterType((*FluentBit)(nil), "magma.mconfig.FluentBit")
	proto.RegisterMapType((map[string]string)(nil), "magma.mconfig.FluentBit.ExtraTagsEntry")
	proto.RegisterMapType((map[string]string)(nil), "magma.mconfig.FluentBit.FilesByTagEntry")
}

func init() {
	proto.RegisterFile("orc8r/protos/mconfig/mconfigs.proto", fileDescriptor_9618f358f05ec5b9)
}

var fileDescriptor_9618f358f05ec5b9 = []byte{
	// 823 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x9c, 0x95, 0xfb, 0x6a, 0xe3, 0xc6,
	0x17, 0xc7, 0x51, 0x7c, 0x89, 0x75, 0x7c, 0xcb, 0xce, 0x6f, 0x7f, 0x5d, 0xc5, 0x14, 0xea, 0x28,
	0x94, 0xb8, 0x2c, 0xd8, 0x25, 0xa5, 0xb0, 0x6c, 0xaf, 0x24, 0xb1, 0x21, 0x25, 0xbb, 0x2c, 0x8a,
	0xd9, 0x42, 0xff, 0x11, 0x63, 0xe9, 0x58, 0x19, 0x32, 0xd2, 0x98, 0xd1, 0xd8, 0x59, 0x3f, 0x48,
	0xe9, 0x43, 0xf4, 0xc1, 0xfa, 0x1a, 0x65, 0x66, 0x24, 0xd5, 0x49, 0x6f, 0xd4, 0x7f, 0x59, 0xe7,
	0x33, 0xdf, 0x73, 0x34, 0x73, 0xbe, 0x67, 0x2c, 0x38, 0x15, 0x32, 0x7a, 0x25, 0x27, 0x2b, 0x29,
	0x94, 0xc8, 0x27, 0x69, 0x24, 0xb2, 0x25, 0x4b, 0xca, 0xdf, 0x7c, 0x6c, 0x38, 0xe9, 0xa6, 0x34,
	0x49, 0xe9, 0xb8, 0xa0, 0x83, 0xe3, 0x47, 0x39, 0x91, 0x48, 0x53, 0x91, 0x59, 0xa5, 0x7f, 0x01,
	0x9d, 0x4b, 0x91, 0x29, 0x29, 0xf8, 0x3b, 0x29, 0x3e, 0x6c, 0xc9, 0x39, 0xb8, 0x5c, 0x24, 0x21,
	0xc7, 0x0d, 0x72, 0xcf, 0x19, 0x3a, 0xa3, 0xde, 0xf9, 0xff, 0xc7, 0xb6, 0x9a, 0x29, 0x32, 0xbe,
	0x11, 0xc9, 0x8d, 0x5e, 0x0c, 0x5a, 0xbc, 0x78, 0xf2, 0x7f, 0x73, 0xa0, 0x7e, 0x95, 0xe5, 0x57,
	0xfb, 0x24, 0x93, 0x4f, 0xa1, 0x87, 0x19, 0x5d, 0x70, 0x0c, 0x23, 0x1a, 0xdd, 0xb1, 0x2c, 0xf1,
	0x0e, 0x86, 0xce, 0xa8, 0x15, 0x74, 0x2d, 0xbd, 0xb4, 0x90, 0x0c, 0xa0, 0xc5, 0x45, 0x44, 0xf9,
	0x7c, 0x7e, 0xe3, 0xd5, 0x86, 0xce, 0xa8, 0x11, 0x54, 0x31, 0x99, 0xc2, 0xa1, 0xc4, 0x48, 0xc8,
	0x38, 0xf7, 0xea, 0xc3, 0xda, 0xa8, 0x7d, 0xfe, 0x72, 0xfc, 0xe8, 0xfc, 0xe3, 0xb7, 0xa8, 0x1e,
	0x84, 0xbc, 0xbf, 0x7a, 0x7b, 0x7b, 0x69, 0x40, 0x60, 0xd5, 0xd7, 0x0a, 0xd3, 0x3c, 0x28, 0x73,
	0xc9, 0x09, 0x74, 0xe2, 0xbb, 0x68, 0x15, 0xda, 0x17, 0xc7, 0x5e, 0xc3, 0xec, 0xa3, 0xad, 0xd9,
	0xd4, 0x22, 0xff, 0x67, 0x07, 0x3e, 0xfe, 0xa7, 0x62, 0xe4, 0x18, 0x5a, 0x34, 0xb4, 0x05, 0x3d,
	0x67, 0x58, 0x1b, 0xb9, 0xc1, 0x21, 0xb5, 0x02, 0xf2, 0x09, 0xb4, 0x29, 0xa5, 0xd5, 0xea, 0x81,
	0x59, 0x05, 0x8d, 0x0a, 0xc1, 0x09, 0x74, 0xa2, 0x8c, 0xa6, 0x58, 0x2a, 0x6a, 0x46, 0xd1, 0x36,
	0xac, 0x90, 0x7c, 0x04, 0xcd, 0x58, 0xa4, 0x94, 0x65, 0x5e, 0x7d, 0xe8, 0x8c, 0xdc, 0xa0, 0x88,
	0xfc, 0x2f, 0xc1, 0xbd, 0x4e, 0x69, 0x82, 0xb7, 0x2b, 0x8c, 0x08, 0x81, 0xba, 0x4e, 0x31, 0x06,
	0xb8, 0x81, 0x79, 0x26, 0xcf, 0xa1, 0x21, 0x64, 0x8c, 0xd2, 0x34, 0xb7, 0x16, 0xd8, 0xc0, 0xff,
	0xb5, 0x0e, 0xcd, 0x37, 0xba, 0x53, 0xfb, 0x59, 0xf7, 0x19, 0x1c, 0x45, 0x77, 0x18, 0xdd, 0xb3,
	0x2c, 0x64, 0x99, 0x42, 0xb9, 0xa1, 0xdc, 0xd4, 0x6f, 0x04, 0xfd, 0x82, 0x5f, 0x17, 0x98, 0x9c,
	0x41, 0x89, 0x42, 0xc5, 0x52, 0x14, 0x6b, 0x55, 0xb8, 0xd8, 0x2b, 0xf0, 0xdc, 0x52, 0x32, 0x81,
	0xff, 0xd1, 0xb5, 0x12, 0xeb, 0x55, 0x22, 0x69, 0x8c, 0x95, 0x17, 0x75, 0xe3, 0x05, 0xd9, 0x59,
	0x2a, 0x2c, 0x21, 0xaf, 0xe1, 0x78, 0x37, 0x61, 0x25, 0x38, 0xff, 0x63, 0x37, 0x0d, 0xf3, 0x8e,
	0x17, 0x3b, 0x82, 0x77, 0x82, 0xf3, 0xdd, 0x5d, 0xad, 0x68, 0x74, 0x4f, 0x13, 0x0c, 0x37, 0x28,
	0x73, 0x26, 0x32, 0xaf, 0x69, 0x9a, 0xd6, 0x2b, 0xf0, 0x7b, 0x4b, 0xc9, 0xe7, 0xd0, 0x64, 0xba,
	0xbf, 0xb9, 0x77, 0x68, 0x06, 0xcc, 0x7b, 0x32, 0x60, 0x55, 0xf3, 0x83, 0x42, 0x47, 0x5e, 0xc0,
	0xa1, 0x62, 0x28, 0x43, 0x16, 0x7b, 0x2d, 0x6b, 0x95, 0x0e, 0xaf, 0x63, 0x72, 0x03, 0xdd, 0x25,
	0x52, 0xb5, 0x96, 0x18, 0x2e, 0x39, 0x4d, 0x72, 0xcf, 0x35, 0x15, 0xcf, 0x9e, 0x54, 0xb4, 0xb6,
	0x8c, 0x67, 0x56, 0x3a, 0xd3, 0xca, 0x69, 0xa6, 0xe4, 0x36, 0xe8, 0x2c, 0x77, 0x90, 0xb6, 0x20,
	0xde, 0x66, 0x34, 0x65, 0x51, 0x98, 0xa3, 0xdc, 0xb0, 0x08, 0x73, 0x0f, 0xcc, 0xdc, 0xf4, 0x0b,
	0x7e, 0x5b, 0xe0, 0xc1, 0x77, 0xf0, 0xec, 0x4f, 0xd5, 0xc8, 0x11, 0xd4, 0xee, 0x71, 0x5b, 0x8c,
	0x8a, 0x7e, 0xd4, 0x93, 0xb2, 0xa1, 0x7c, 0x8d, 0xc5, 0x35, 0xb4, 0xc1, 0xeb, 0x83, 0x57, 0x8e,
	0x8f, 0xd0, 0x9c, 0x6e, 0x30, 0x53, 0xfb, 0x0d, 0xcb, 0x19, 0xf4, 0x51, 0x67, 0xeb, 0x4e, 0x2f,
	0x44, 0xce, 0xd4, 0xb6, 0x98, 0x95, 0x9e, 0xc1, 0xef, 0x4b, 0xea, 0x7f, 0x0f, 0x70, 0xc5, 0x24,
	0x46, 0x4a, 0xc8, 0xed, 0x5e, 0xaf, 0xf2, 0xbf, 0x85, 0xd6, 0x1b, 0x54, 0x92, 0x45, 0xfb, 0xfd,
	0x25, 0xf9, 0x5f, 0x41, 0xe3, 0x56, 0x51, 0x85, 0x7b, 0x25, 0xff, 0x52, 0x03, 0x77, 0xc6, 0xd7,
	0x98, 0xa9, 0x0b, 0xa6, 0xc8, 0x0c, 0x00, 0x3f, 0x28, 0x49, 0x43, 0xa5, 0xad, 0x76, 0xfe, 0xd2,
	0xea, 0x4a, 0x3d, 0x9e, 0x6a, 0xe9, 0xbc, 0xb2, 0xda, 0xc5, 0x32, 0x26, 0xa7, 0xd0, 0x55, 0x77,
	0x52, 0x28, 0xc5, 0x31, 0x94, 0x54, 0xa1, 0x07, 0x43, 0x67, 0xd4, 0x0d, 0x3a, 0x25, 0x0c, 0xf4,
	0x76, 0xcf, 0xa0, 0x5f, 0x89, 0x1e, 0x58, 0x16, 0x8b, 0x07, 0xaf, 0x6d, 0x64, 0xbd, 0x12, 0xff,
	0x68, 0x28, 0x79, 0x09, 0xcf, 0x2a, 0x61, 0x75, 0x57, 0x3a, 0x66, 0x06, 0x8e, 0xca, 0x85, 0xea,
	0x92, 0xfc, 0x00, 0x9d, 0x25, 0xe3, 0x98, 0x87, 0x8b, 0xad, 0x3e, 0x85, 0xf7, 0xdc, 0x1c, 0x62,
	0xf4, 0xb7, 0x87, 0x98, 0x69, 0xf1, 0xc5, 0x76, 0x4e, 0x13, 0x7b, 0x0a, 0x58, 0x56, 0x60, 0xf0,
	0x35, 0xf4, 0x1e, 0x9f, 0xf1, 0xdf, 0x06, 0xd0, 0xdd, 0x19, 0xc0, 0xc1, 0x37, 0xd0, 0x7f, 0x52,
	0xfc, 0xbf, 0xa4, 0x5f, 0x9c, 0xfe, 0x74, 0x62, 0xf6, 0x3c, 0xb1, 0x5f, 0x43, 0xce, 0x16, 0x93,
	0x44, 0x3c, 0xf9, 0x90, 0x2e, 0x9a, 0x26, 0xfe, 0xe2, 0xf7, 0x00, 0x00, 0x00, 0xff, 0xff, 0x38,
	0xd8, 0x80, 0x8c, 0x67, 0x07, 0x00, 0x00,
}
