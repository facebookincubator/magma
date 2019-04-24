// Code generated by protoc-gen-go. DO NOT EDIT.
// source: lte/protos/enodebd.proto

package protos // import "magma/lte/cloud/go/protos"

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"
import protos "magma/orc8r/cloud/go/protos"

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

// --------------------------------------------------------------------------
// Message Definitions for TR-069 message injection. This is used for manual
// testing of the TR-069 server.
// --------------------------------------------------------------------------
type GetParameterRequest struct {
	// Serial ID of eNodeB. Uniquely identifies the eNodeB.
	DeviceSerial string `protobuf:"bytes,1,opt,name=device_serial,json=deviceSerial,proto3" json:"device_serial,omitempty"`
	// Fully qualified parameter name, e.g:
	// InternetGatewayDevice.LANDevice.1.Hosts.
	ParameterName        string   `protobuf:"bytes,2,opt,name=parameter_name,json=parameterName,proto3" json:"parameter_name,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *GetParameterRequest) Reset()         { *m = GetParameterRequest{} }
func (m *GetParameterRequest) String() string { return proto.CompactTextString(m) }
func (*GetParameterRequest) ProtoMessage()    {}
func (*GetParameterRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_enodebd_b6ad9ca24c20de46, []int{0}
}
func (m *GetParameterRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetParameterRequest.Unmarshal(m, b)
}
func (m *GetParameterRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetParameterRequest.Marshal(b, m, deterministic)
}
func (dst *GetParameterRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetParameterRequest.Merge(dst, src)
}
func (m *GetParameterRequest) XXX_Size() int {
	return xxx_messageInfo_GetParameterRequest.Size(m)
}
func (m *GetParameterRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_GetParameterRequest.DiscardUnknown(m)
}

var xxx_messageInfo_GetParameterRequest proto.InternalMessageInfo

func (m *GetParameterRequest) GetDeviceSerial() string {
	if m != nil {
		return m.DeviceSerial
	}
	return ""
}

func (m *GetParameterRequest) GetParameterName() string {
	if m != nil {
		return m.ParameterName
	}
	return ""
}

type NameValue struct {
	Name string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	// Note: parameter value is always passed back as string. Up to calling
	// function to determine type
	Value                string   `protobuf:"bytes,2,opt,name=value,proto3" json:"value,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *NameValue) Reset()         { *m = NameValue{} }
func (m *NameValue) String() string { return proto.CompactTextString(m) }
func (*NameValue) ProtoMessage()    {}
func (*NameValue) Descriptor() ([]byte, []int) {
	return fileDescriptor_enodebd_b6ad9ca24c20de46, []int{1}
}
func (m *NameValue) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_NameValue.Unmarshal(m, b)
}
func (m *NameValue) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_NameValue.Marshal(b, m, deterministic)
}
func (dst *NameValue) XXX_Merge(src proto.Message) {
	xxx_messageInfo_NameValue.Merge(dst, src)
}
func (m *NameValue) XXX_Size() int {
	return xxx_messageInfo_NameValue.Size(m)
}
func (m *NameValue) XXX_DiscardUnknown() {
	xxx_messageInfo_NameValue.DiscardUnknown(m)
}

var xxx_messageInfo_NameValue proto.InternalMessageInfo

func (m *NameValue) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *NameValue) GetValue() string {
	if m != nil {
		return m.Value
	}
	return ""
}

type GetParameterResponse struct {
	DeviceSerial         string       `protobuf:"bytes,1,opt,name=device_serial,json=deviceSerial,proto3" json:"device_serial,omitempty"`
	Parameters           []*NameValue `protobuf:"bytes,2,rep,name=parameters,proto3" json:"parameters,omitempty"`
	XXX_NoUnkeyedLiteral struct{}     `json:"-"`
	XXX_unrecognized     []byte       `json:"-"`
	XXX_sizecache        int32        `json:"-"`
}

func (m *GetParameterResponse) Reset()         { *m = GetParameterResponse{} }
func (m *GetParameterResponse) String() string { return proto.CompactTextString(m) }
func (*GetParameterResponse) ProtoMessage()    {}
func (*GetParameterResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_enodebd_b6ad9ca24c20de46, []int{2}
}
func (m *GetParameterResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetParameterResponse.Unmarshal(m, b)
}
func (m *GetParameterResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetParameterResponse.Marshal(b, m, deterministic)
}
func (dst *GetParameterResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetParameterResponse.Merge(dst, src)
}
func (m *GetParameterResponse) XXX_Size() int {
	return xxx_messageInfo_GetParameterResponse.Size(m)
}
func (m *GetParameterResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_GetParameterResponse.DiscardUnknown(m)
}

var xxx_messageInfo_GetParameterResponse proto.InternalMessageInfo

func (m *GetParameterResponse) GetDeviceSerial() string {
	if m != nil {
		return m.DeviceSerial
	}
	return ""
}

func (m *GetParameterResponse) GetParameters() []*NameValue {
	if m != nil {
		return m.Parameters
	}
	return nil
}

type SetParameterRequest struct {
	// Serial ID of eNodeB. Uniquely identifies the eNodeB.
	DeviceSerial string `protobuf:"bytes,1,opt,name=device_serial,json=deviceSerial,proto3" json:"device_serial,omitempty"`
	// Fully qualified parameter name, e.g:
	// InternetGatewayDevice.LANDevice.1.Hosts.
	ParameterName string `protobuf:"bytes,2,opt,name=parameter_name,json=parameterName,proto3" json:"parameter_name,omitempty"`
	// Data values for each data type
	//
	// Types that are valid to be assigned to Value:
	//	*SetParameterRequest_ValueInt
	//	*SetParameterRequest_ValueString
	//	*SetParameterRequest_ValueBool
	Value isSetParameterRequest_Value `protobuf_oneof:"value"`
	// Key to be used at ACS discretion to determine when parameter was last
	// updated
	ParameterKey         string   `protobuf:"bytes,6,opt,name=parameter_key,json=parameterKey,proto3" json:"parameter_key,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *SetParameterRequest) Reset()         { *m = SetParameterRequest{} }
func (m *SetParameterRequest) String() string { return proto.CompactTextString(m) }
func (*SetParameterRequest) ProtoMessage()    {}
func (*SetParameterRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_enodebd_b6ad9ca24c20de46, []int{3}
}
func (m *SetParameterRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_SetParameterRequest.Unmarshal(m, b)
}
func (m *SetParameterRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_SetParameterRequest.Marshal(b, m, deterministic)
}
func (dst *SetParameterRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SetParameterRequest.Merge(dst, src)
}
func (m *SetParameterRequest) XXX_Size() int {
	return xxx_messageInfo_SetParameterRequest.Size(m)
}
func (m *SetParameterRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_SetParameterRequest.DiscardUnknown(m)
}

var xxx_messageInfo_SetParameterRequest proto.InternalMessageInfo

func (m *SetParameterRequest) GetDeviceSerial() string {
	if m != nil {
		return m.DeviceSerial
	}
	return ""
}

func (m *SetParameterRequest) GetParameterName() string {
	if m != nil {
		return m.ParameterName
	}
	return ""
}

type isSetParameterRequest_Value interface {
	isSetParameterRequest_Value()
}

type SetParameterRequest_ValueInt struct {
	ValueInt int32 `protobuf:"varint,3,opt,name=value_int,json=valueInt,proto3,oneof"`
}

type SetParameterRequest_ValueString struct {
	ValueString string `protobuf:"bytes,4,opt,name=value_string,json=valueString,proto3,oneof"`
}

type SetParameterRequest_ValueBool struct {
	ValueBool bool `protobuf:"varint,5,opt,name=value_bool,json=valueBool,proto3,oneof"`
}

func (*SetParameterRequest_ValueInt) isSetParameterRequest_Value() {}

func (*SetParameterRequest_ValueString) isSetParameterRequest_Value() {}

func (*SetParameterRequest_ValueBool) isSetParameterRequest_Value() {}

func (m *SetParameterRequest) GetValue() isSetParameterRequest_Value {
	if m != nil {
		return m.Value
	}
	return nil
}

func (m *SetParameterRequest) GetValueInt() int32 {
	if x, ok := m.GetValue().(*SetParameterRequest_ValueInt); ok {
		return x.ValueInt
	}
	return 0
}

func (m *SetParameterRequest) GetValueString() string {
	if x, ok := m.GetValue().(*SetParameterRequest_ValueString); ok {
		return x.ValueString
	}
	return ""
}

func (m *SetParameterRequest) GetValueBool() bool {
	if x, ok := m.GetValue().(*SetParameterRequest_ValueBool); ok {
		return x.ValueBool
	}
	return false
}

func (m *SetParameterRequest) GetParameterKey() string {
	if m != nil {
		return m.ParameterKey
	}
	return ""
}

// XXX_OneofFuncs is for the internal use of the proto package.
func (*SetParameterRequest) XXX_OneofFuncs() (func(msg proto.Message, b *proto.Buffer) error, func(msg proto.Message, tag, wire int, b *proto.Buffer) (bool, error), func(msg proto.Message) (n int), []interface{}) {
	return _SetParameterRequest_OneofMarshaler, _SetParameterRequest_OneofUnmarshaler, _SetParameterRequest_OneofSizer, []interface{}{
		(*SetParameterRequest_ValueInt)(nil),
		(*SetParameterRequest_ValueString)(nil),
		(*SetParameterRequest_ValueBool)(nil),
	}
}

func _SetParameterRequest_OneofMarshaler(msg proto.Message, b *proto.Buffer) error {
	m := msg.(*SetParameterRequest)
	// value
	switch x := m.Value.(type) {
	case *SetParameterRequest_ValueInt:
		b.EncodeVarint(3<<3 | proto.WireVarint)
		b.EncodeVarint(uint64(x.ValueInt))
	case *SetParameterRequest_ValueString:
		b.EncodeVarint(4<<3 | proto.WireBytes)
		b.EncodeStringBytes(x.ValueString)
	case *SetParameterRequest_ValueBool:
		t := uint64(0)
		if x.ValueBool {
			t = 1
		}
		b.EncodeVarint(5<<3 | proto.WireVarint)
		b.EncodeVarint(t)
	case nil:
	default:
		return fmt.Errorf("SetParameterRequest.Value has unexpected type %T", x)
	}
	return nil
}

func _SetParameterRequest_OneofUnmarshaler(msg proto.Message, tag, wire int, b *proto.Buffer) (bool, error) {
	m := msg.(*SetParameterRequest)
	switch tag {
	case 3: // value.value_int
		if wire != proto.WireVarint {
			return true, proto.ErrInternalBadWireType
		}
		x, err := b.DecodeVarint()
		m.Value = &SetParameterRequest_ValueInt{int32(x)}
		return true, err
	case 4: // value.value_string
		if wire != proto.WireBytes {
			return true, proto.ErrInternalBadWireType
		}
		x, err := b.DecodeStringBytes()
		m.Value = &SetParameterRequest_ValueString{x}
		return true, err
	case 5: // value.value_bool
		if wire != proto.WireVarint {
			return true, proto.ErrInternalBadWireType
		}
		x, err := b.DecodeVarint()
		m.Value = &SetParameterRequest_ValueBool{x != 0}
		return true, err
	default:
		return false, nil
	}
}

func _SetParameterRequest_OneofSizer(msg proto.Message) (n int) {
	m := msg.(*SetParameterRequest)
	// value
	switch x := m.Value.(type) {
	case *SetParameterRequest_ValueInt:
		n += 1 // tag and wire
		n += proto.SizeVarint(uint64(x.ValueInt))
	case *SetParameterRequest_ValueString:
		n += 1 // tag and wire
		n += proto.SizeVarint(uint64(len(x.ValueString)))
		n += len(x.ValueString)
	case *SetParameterRequest_ValueBool:
		n += 1 // tag and wire
		n += 1
	case nil:
	default:
		panic(fmt.Sprintf("proto: unexpected type %T in oneof", x))
	}
	return n
}

type EnodebIdentity struct {
	// Serial ID of eNodeB. Uniquely identifies the eNodeB.
	DeviceSerial         string   `protobuf:"bytes,1,opt,name=device_serial,json=deviceSerial,proto3" json:"device_serial,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *EnodebIdentity) Reset()         { *m = EnodebIdentity{} }
func (m *EnodebIdentity) String() string { return proto.CompactTextString(m) }
func (*EnodebIdentity) ProtoMessage()    {}
func (*EnodebIdentity) Descriptor() ([]byte, []int) {
	return fileDescriptor_enodebd_b6ad9ca24c20de46, []int{4}
}
func (m *EnodebIdentity) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_EnodebIdentity.Unmarshal(m, b)
}
func (m *EnodebIdentity) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_EnodebIdentity.Marshal(b, m, deterministic)
}
func (dst *EnodebIdentity) XXX_Merge(src proto.Message) {
	xxx_messageInfo_EnodebIdentity.Merge(dst, src)
}
func (m *EnodebIdentity) XXX_Size() int {
	return xxx_messageInfo_EnodebIdentity.Size(m)
}
func (m *EnodebIdentity) XXX_DiscardUnknown() {
	xxx_messageInfo_EnodebIdentity.DiscardUnknown(m)
}

var xxx_messageInfo_EnodebIdentity proto.InternalMessageInfo

func (m *EnodebIdentity) GetDeviceSerial() string {
	if m != nil {
		return m.DeviceSerial
	}
	return ""
}

type AllEnodebStatus struct {
	EnbStatusList        []*SingleEnodebStatus `protobuf:"bytes,1,rep,name=enb_status_list,json=enbStatusList,proto3" json:"enb_status_list,omitempty"`
	XXX_NoUnkeyedLiteral struct{}              `json:"-"`
	XXX_unrecognized     []byte                `json:"-"`
	XXX_sizecache        int32                 `json:"-"`
}

func (m *AllEnodebStatus) Reset()         { *m = AllEnodebStatus{} }
func (m *AllEnodebStatus) String() string { return proto.CompactTextString(m) }
func (*AllEnodebStatus) ProtoMessage()    {}
func (*AllEnodebStatus) Descriptor() ([]byte, []int) {
	return fileDescriptor_enodebd_b6ad9ca24c20de46, []int{5}
}
func (m *AllEnodebStatus) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_AllEnodebStatus.Unmarshal(m, b)
}
func (m *AllEnodebStatus) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_AllEnodebStatus.Marshal(b, m, deterministic)
}
func (dst *AllEnodebStatus) XXX_Merge(src proto.Message) {
	xxx_messageInfo_AllEnodebStatus.Merge(dst, src)
}
func (m *AllEnodebStatus) XXX_Size() int {
	return xxx_messageInfo_AllEnodebStatus.Size(m)
}
func (m *AllEnodebStatus) XXX_DiscardUnknown() {
	xxx_messageInfo_AllEnodebStatus.DiscardUnknown(m)
}

var xxx_messageInfo_AllEnodebStatus proto.InternalMessageInfo

func (m *AllEnodebStatus) GetEnbStatusList() []*SingleEnodebStatus {
	if m != nil {
		return m.EnbStatusList
	}
	return nil
}

type SingleEnodebStatus struct {
	DeviceSerial         string   `protobuf:"bytes,1,opt,name=device_serial,json=deviceSerial,proto3" json:"device_serial,omitempty"`
	IpAddress            string   `protobuf:"bytes,2,opt,name=ip_address,json=ipAddress,proto3" json:"ip_address,omitempty"`
	Connected            string   `protobuf:"bytes,3,opt,name=connected,proto3" json:"connected,omitempty"`
	Configured           string   `protobuf:"bytes,4,opt,name=configured,proto3" json:"configured,omitempty"`
	OpstateEnabled       string   `protobuf:"bytes,5,opt,name=opstate_enabled,json=opstateEnabled,proto3" json:"opstate_enabled,omitempty"`
	RfTxOn               string   `protobuf:"bytes,6,opt,name=rf_tx_on,json=rfTxOn,proto3" json:"rf_tx_on,omitempty"`
	GpsConnected         string   `protobuf:"bytes,7,opt,name=gps_connected,json=gpsConnected,proto3" json:"gps_connected,omitempty"`
	PtpConnected         string   `protobuf:"bytes,8,opt,name=ptp_connected,json=ptpConnected,proto3" json:"ptp_connected,omitempty"`
	MmeConnected         string   `protobuf:"bytes,9,opt,name=mme_connected,json=mmeConnected,proto3" json:"mme_connected,omitempty"`
	GpsLongitude         string   `protobuf:"bytes,10,opt,name=gps_longitude,json=gpsLongitude,proto3" json:"gps_longitude,omitempty"`
	GpsLatitude          string   `protobuf:"bytes,11,opt,name=gps_latitude,json=gpsLatitude,proto3" json:"gps_latitude,omitempty"`
	FsmState             string   `protobuf:"bytes,12,opt,name=fsm_state,json=fsmState,proto3" json:"fsm_state,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *SingleEnodebStatus) Reset()         { *m = SingleEnodebStatus{} }
func (m *SingleEnodebStatus) String() string { return proto.CompactTextString(m) }
func (*SingleEnodebStatus) ProtoMessage()    {}
func (*SingleEnodebStatus) Descriptor() ([]byte, []int) {
	return fileDescriptor_enodebd_b6ad9ca24c20de46, []int{6}
}
func (m *SingleEnodebStatus) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_SingleEnodebStatus.Unmarshal(m, b)
}
func (m *SingleEnodebStatus) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_SingleEnodebStatus.Marshal(b, m, deterministic)
}
func (dst *SingleEnodebStatus) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SingleEnodebStatus.Merge(dst, src)
}
func (m *SingleEnodebStatus) XXX_Size() int {
	return xxx_messageInfo_SingleEnodebStatus.Size(m)
}
func (m *SingleEnodebStatus) XXX_DiscardUnknown() {
	xxx_messageInfo_SingleEnodebStatus.DiscardUnknown(m)
}

var xxx_messageInfo_SingleEnodebStatus proto.InternalMessageInfo

func (m *SingleEnodebStatus) GetDeviceSerial() string {
	if m != nil {
		return m.DeviceSerial
	}
	return ""
}

func (m *SingleEnodebStatus) GetIpAddress() string {
	if m != nil {
		return m.IpAddress
	}
	return ""
}

func (m *SingleEnodebStatus) GetConnected() string {
	if m != nil {
		return m.Connected
	}
	return ""
}

func (m *SingleEnodebStatus) GetConfigured() string {
	if m != nil {
		return m.Configured
	}
	return ""
}

func (m *SingleEnodebStatus) GetOpstateEnabled() string {
	if m != nil {
		return m.OpstateEnabled
	}
	return ""
}

func (m *SingleEnodebStatus) GetRfTxOn() string {
	if m != nil {
		return m.RfTxOn
	}
	return ""
}

func (m *SingleEnodebStatus) GetGpsConnected() string {
	if m != nil {
		return m.GpsConnected
	}
	return ""
}

func (m *SingleEnodebStatus) GetPtpConnected() string {
	if m != nil {
		return m.PtpConnected
	}
	return ""
}

func (m *SingleEnodebStatus) GetMmeConnected() string {
	if m != nil {
		return m.MmeConnected
	}
	return ""
}

func (m *SingleEnodebStatus) GetGpsLongitude() string {
	if m != nil {
		return m.GpsLongitude
	}
	return ""
}

func (m *SingleEnodebStatus) GetGpsLatitude() string {
	if m != nil {
		return m.GpsLatitude
	}
	return ""
}

func (m *SingleEnodebStatus) GetFsmState() string {
	if m != nil {
		return m.FsmState
	}
	return ""
}

func init() {
	proto.RegisterType((*GetParameterRequest)(nil), "magma.lte.GetParameterRequest")
	proto.RegisterType((*NameValue)(nil), "magma.lte.NameValue")
	proto.RegisterType((*GetParameterResponse)(nil), "magma.lte.GetParameterResponse")
	proto.RegisterType((*SetParameterRequest)(nil), "magma.lte.SetParameterRequest")
	proto.RegisterType((*EnodebIdentity)(nil), "magma.lte.EnodebIdentity")
	proto.RegisterType((*AllEnodebStatus)(nil), "magma.lte.AllEnodebStatus")
	proto.RegisterType((*SingleEnodebStatus)(nil), "magma.lte.SingleEnodebStatus")
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// EnodebdClient is the client API for Enodebd service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type EnodebdClient interface {
	// Sends GetParameterValues message to ENodeB. TR-069 supports multiple
	// parameter names per message, but only one is supported here.
	GetParameter(ctx context.Context, in *GetParameterRequest, opts ...grpc.CallOption) (*GetParameterResponse, error)
	// Sends SetParameterValues message to ENodeB. TR-069 supports multiple
	// parameter names per message, but only one is supported here.
	SetParameter(ctx context.Context, in *SetParameterRequest, opts ...grpc.CallOption) (*protos.Void, error)
	// Configure eNodeB based on enodebd config file
	Configure(ctx context.Context, in *EnodebIdentity, opts ...grpc.CallOption) (*protos.Void, error)
	// Reboot eNodeB
	Reboot(ctx context.Context, in *EnodebIdentity, opts ...grpc.CallOption) (*protos.Void, error)
	// Get current status
	GetStatus(ctx context.Context, in *protos.Void, opts ...grpc.CallOption) (*protos.ServiceStatus, error)
	// Get status info for all connected eNodeB devices
	GetAllEnodebStatus(ctx context.Context, in *protos.Void, opts ...grpc.CallOption) (*AllEnodebStatus, error)
	// Get status info of a single connected eNodeB device
	GetEnodebStatus(ctx context.Context, in *EnodebIdentity, opts ...grpc.CallOption) (*SingleEnodebStatus, error)
}

type enodebdClient struct {
	cc *grpc.ClientConn
}

func NewEnodebdClient(cc *grpc.ClientConn) EnodebdClient {
	return &enodebdClient{cc}
}

func (c *enodebdClient) GetParameter(ctx context.Context, in *GetParameterRequest, opts ...grpc.CallOption) (*GetParameterResponse, error) {
	out := new(GetParameterResponse)
	err := c.cc.Invoke(ctx, "/magma.lte.Enodebd/GetParameter", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *enodebdClient) SetParameter(ctx context.Context, in *SetParameterRequest, opts ...grpc.CallOption) (*protos.Void, error) {
	out := new(protos.Void)
	err := c.cc.Invoke(ctx, "/magma.lte.Enodebd/SetParameter", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *enodebdClient) Configure(ctx context.Context, in *EnodebIdentity, opts ...grpc.CallOption) (*protos.Void, error) {
	out := new(protos.Void)
	err := c.cc.Invoke(ctx, "/magma.lte.Enodebd/Configure", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *enodebdClient) Reboot(ctx context.Context, in *EnodebIdentity, opts ...grpc.CallOption) (*protos.Void, error) {
	out := new(protos.Void)
	err := c.cc.Invoke(ctx, "/magma.lte.Enodebd/Reboot", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *enodebdClient) GetStatus(ctx context.Context, in *protos.Void, opts ...grpc.CallOption) (*protos.ServiceStatus, error) {
	out := new(protos.ServiceStatus)
	err := c.cc.Invoke(ctx, "/magma.lte.Enodebd/GetStatus", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *enodebdClient) GetAllEnodebStatus(ctx context.Context, in *protos.Void, opts ...grpc.CallOption) (*AllEnodebStatus, error) {
	out := new(AllEnodebStatus)
	err := c.cc.Invoke(ctx, "/magma.lte.Enodebd/GetAllEnodebStatus", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *enodebdClient) GetEnodebStatus(ctx context.Context, in *EnodebIdentity, opts ...grpc.CallOption) (*SingleEnodebStatus, error) {
	out := new(SingleEnodebStatus)
	err := c.cc.Invoke(ctx, "/magma.lte.Enodebd/GetEnodebStatus", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// EnodebdServer is the server API for Enodebd service.
type EnodebdServer interface {
	// Sends GetParameterValues message to ENodeB. TR-069 supports multiple
	// parameter names per message, but only one is supported here.
	GetParameter(context.Context, *GetParameterRequest) (*GetParameterResponse, error)
	// Sends SetParameterValues message to ENodeB. TR-069 supports multiple
	// parameter names per message, but only one is supported here.
	SetParameter(context.Context, *SetParameterRequest) (*protos.Void, error)
	// Configure eNodeB based on enodebd config file
	Configure(context.Context, *EnodebIdentity) (*protos.Void, error)
	// Reboot eNodeB
	Reboot(context.Context, *EnodebIdentity) (*protos.Void, error)
	// Get current status
	GetStatus(context.Context, *protos.Void) (*protos.ServiceStatus, error)
	// Get status info for all connected eNodeB devices
	GetAllEnodebStatus(context.Context, *protos.Void) (*AllEnodebStatus, error)
	// Get status info of a single connected eNodeB device
	GetEnodebStatus(context.Context, *EnodebIdentity) (*SingleEnodebStatus, error)
}

func RegisterEnodebdServer(s *grpc.Server, srv EnodebdServer) {
	s.RegisterService(&_Enodebd_serviceDesc, srv)
}

func _Enodebd_GetParameter_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetParameterRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(EnodebdServer).GetParameter(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/magma.lte.Enodebd/GetParameter",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(EnodebdServer).GetParameter(ctx, req.(*GetParameterRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Enodebd_SetParameter_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SetParameterRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(EnodebdServer).SetParameter(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/magma.lte.Enodebd/SetParameter",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(EnodebdServer).SetParameter(ctx, req.(*SetParameterRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Enodebd_Configure_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(EnodebIdentity)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(EnodebdServer).Configure(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/magma.lte.Enodebd/Configure",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(EnodebdServer).Configure(ctx, req.(*EnodebIdentity))
	}
	return interceptor(ctx, in, info, handler)
}

func _Enodebd_Reboot_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(EnodebIdentity)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(EnodebdServer).Reboot(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/magma.lte.Enodebd/Reboot",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(EnodebdServer).Reboot(ctx, req.(*EnodebIdentity))
	}
	return interceptor(ctx, in, info, handler)
}

func _Enodebd_GetStatus_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(protos.Void)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(EnodebdServer).GetStatus(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/magma.lte.Enodebd/GetStatus",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(EnodebdServer).GetStatus(ctx, req.(*protos.Void))
	}
	return interceptor(ctx, in, info, handler)
}

func _Enodebd_GetAllEnodebStatus_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(protos.Void)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(EnodebdServer).GetAllEnodebStatus(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/magma.lte.Enodebd/GetAllEnodebStatus",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(EnodebdServer).GetAllEnodebStatus(ctx, req.(*protos.Void))
	}
	return interceptor(ctx, in, info, handler)
}

func _Enodebd_GetEnodebStatus_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(EnodebIdentity)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(EnodebdServer).GetEnodebStatus(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/magma.lte.Enodebd/GetEnodebStatus",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(EnodebdServer).GetEnodebStatus(ctx, req.(*EnodebIdentity))
	}
	return interceptor(ctx, in, info, handler)
}

var _Enodebd_serviceDesc = grpc.ServiceDesc{
	ServiceName: "magma.lte.Enodebd",
	HandlerType: (*EnodebdServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetParameter",
			Handler:    _Enodebd_GetParameter_Handler,
		},
		{
			MethodName: "SetParameter",
			Handler:    _Enodebd_SetParameter_Handler,
		},
		{
			MethodName: "Configure",
			Handler:    _Enodebd_Configure_Handler,
		},
		{
			MethodName: "Reboot",
			Handler:    _Enodebd_Reboot_Handler,
		},
		{
			MethodName: "GetStatus",
			Handler:    _Enodebd_GetStatus_Handler,
		},
		{
			MethodName: "GetAllEnodebStatus",
			Handler:    _Enodebd_GetAllEnodebStatus_Handler,
		},
		{
			MethodName: "GetEnodebStatus",
			Handler:    _Enodebd_GetEnodebStatus_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "lte/protos/enodebd.proto",
}

func init() { proto.RegisterFile("lte/protos/enodebd.proto", fileDescriptor_enodebd_b6ad9ca24c20de46) }

var fileDescriptor_enodebd_b6ad9ca24c20de46 = []byte{
	// 703 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xb4, 0x55, 0xdd, 0x4e, 0xdb, 0x4a,
	0x10, 0x26, 0x04, 0x42, 0x3c, 0x09, 0x44, 0x67, 0xe1, 0xc2, 0x84, 0x13, 0xc8, 0xf1, 0x51, 0xd5,
	0x5c, 0x25, 0x15, 0x94, 0xaa, 0xed, 0x1d, 0x20, 0x14, 0x10, 0xa8, 0x54, 0x76, 0x85, 0xaa, 0xde,
	0x58, 0x8e, 0x3d, 0xb1, 0xac, 0xda, 0xbb, 0xc6, 0xbb, 0x41, 0xf0, 0x50, 0x7d, 0xa5, 0x3e, 0x45,
	0x1f, 0xa0, 0xda, 0x5d, 0xc7, 0xd8, 0xfc, 0x89, 0x5e, 0xf4, 0xca, 0xde, 0x6f, 0xbe, 0x99, 0xfd,
	0x66, 0x66, 0x67, 0x17, 0xcc, 0x58, 0xe0, 0x28, 0xcd, 0x98, 0x60, 0x7c, 0x84, 0x94, 0x05, 0x38,
	0x09, 0x86, 0x6a, 0x49, 0x8c, 0xc4, 0x0b, 0x13, 0x6f, 0x18, 0x0b, 0xec, 0xf6, 0x58, 0xe6, 0xbf,
	0xcf, 0xe6, 0x34, 0x8e, 0xd9, 0x75, 0xe4, 0xe3, 0xde, 0x9b, 0x3d, 0xcd, 0xec, 0x6e, 0x56, 0xcc,
	0x3e, 0x4b, 0x12, 0x46, 0xb5, 0xc9, 0xf2, 0x60, 0x7d, 0x8c, 0xe2, 0xb3, 0x97, 0x79, 0x09, 0x0a,
	0xcc, 0x6c, 0xbc, 0x9a, 0x21, 0x17, 0xe4, 0x7f, 0x58, 0x0d, 0x50, 0x06, 0x71, 0x39, 0x66, 0x91,
	0x17, 0x9b, 0xb5, 0x7e, 0x6d, 0x60, 0xd8, 0x6d, 0x0d, 0x3a, 0x0a, 0x23, 0xaf, 0x60, 0x2d, 0x9d,
	0x3b, 0xba, 0xd4, 0x4b, 0xd0, 0x5c, 0x54, 0xac, 0xd5, 0x02, 0xfd, 0xe4, 0x25, 0x68, 0xed, 0x83,
	0x21, 0xbf, 0x97, 0x5e, 0x3c, 0x43, 0x42, 0x60, 0x49, 0x31, 0x75, 0x3c, 0xf5, 0x4f, 0x36, 0x60,
	0xf9, 0x5a, 0x1a, 0x73, 0x77, 0xbd, 0xb0, 0xae, 0x60, 0xa3, 0xaa, 0x8c, 0xa7, 0x8c, 0x72, 0x7c,
	0x99, 0xb4, 0xb7, 0x00, 0x85, 0x08, 0x6e, 0x2e, 0xf6, 0xeb, 0x83, 0xd6, 0xee, 0xc6, 0xb0, 0x28,
	0xd8, 0xb0, 0x10, 0x64, 0x97, 0x78, 0xd6, 0xaf, 0x1a, 0xac, 0x3b, 0x7f, 0xb7, 0x1a, 0xa4, 0x07,
	0x86, 0xca, 0xcf, 0x8d, 0xa8, 0x30, 0xeb, 0xfd, 0xda, 0x60, 0xf9, 0x64, 0xc1, 0x6e, 0x2a, 0xe8,
	0x94, 0xca, 0xad, 0xda, 0xda, 0xcc, 0x45, 0x16, 0xd1, 0xd0, 0x5c, 0x92, 0x31, 0x4e, 0x16, 0xec,
	0x96, 0x42, 0x1d, 0x05, 0x92, 0x1d, 0x00, 0x4d, 0x9a, 0x30, 0x16, 0x9b, 0xcb, 0xfd, 0xda, 0xa0,
	0x79, 0xb2, 0x60, 0xeb, 0xb8, 0x87, 0x8c, 0xc5, 0x52, 0xf0, 0x9d, 0x96, 0xef, 0x78, 0x6b, 0x36,
	0xb4, 0xe0, 0x02, 0x3c, 0xc3, 0xdb, 0xc3, 0x95, 0xbc, 0xec, 0xd6, 0x3e, 0xac, 0x1d, 0xab, 0x93,
	0x75, 0x1a, 0x20, 0x15, 0x91, 0xb8, 0x7d, 0x51, 0xc2, 0xd6, 0x57, 0xe8, 0x1c, 0xc4, 0xb1, 0xf6,
	0x74, 0x84, 0x27, 0x66, 0x9c, 0x1c, 0x43, 0x07, 0xe9, 0xc4, 0xe5, 0x6a, 0xe5, 0xc6, 0x11, 0x17,
	0x66, 0x4d, 0xd5, 0xbe, 0x57, 0xaa, 0xbd, 0x13, 0xd1, 0x30, 0xc6, 0xb2, 0x9f, 0xbd, 0x8a, 0x34,
	0xff, 0x3d, 0x8f, 0xb8, 0xb0, 0x7e, 0xd4, 0x81, 0x3c, 0x64, 0xbd, 0xac, 0x0d, 0x3d, 0x80, 0x28,
	0x75, 0xbd, 0x20, 0xc8, 0x90, 0xf3, 0xbc, 0x05, 0x46, 0x94, 0x1e, 0x68, 0x80, 0xfc, 0x0b, 0x86,
	0xcf, 0x28, 0x45, 0x5f, 0x60, 0xa0, 0xca, 0x6f, 0xd8, 0x77, 0x00, 0xd9, 0x06, 0xf0, 0x19, 0x9d,
	0x46, 0xe1, 0x2c, 0xc3, 0x40, 0xd7, 0xde, 0x2e, 0x21, 0xe4, 0x35, 0x74, 0x58, 0x2a, 0xb3, 0x43,
	0x17, 0xa9, 0x37, 0x89, 0x31, 0x50, 0xd5, 0x37, 0xec, 0xb5, 0x1c, 0x3e, 0xd6, 0x28, 0x31, 0xa1,
	0x99, 0x4d, 0x5d, 0x71, 0xe3, 0x32, 0x9a, 0xd7, 0xbe, 0x91, 0x4d, 0xbf, 0xdc, 0x5c, 0x50, 0x99,
	0x44, 0x98, 0x72, 0xf7, 0x4e, 0xc4, 0x8a, 0x4e, 0x22, 0x4c, 0xf9, 0x51, 0xa1, 0x43, 0xf6, 0x4f,
	0xa4, 0x25, 0x52, 0x33, 0xef, 0x9f, 0x48, 0x2b, 0xa4, 0x24, 0xc1, 0x12, 0xc9, 0xd0, 0xa4, 0x24,
	0xc1, 0x0a, 0x49, 0x6e, 0x17, 0x33, 0x1a, 0x46, 0x62, 0x16, 0xa0, 0x09, 0xc5, 0x76, 0xe7, 0x73,
	0x8c, 0xfc, 0x07, 0x6d, 0x45, 0xf2, 0x84, 0xe6, 0xb4, 0x14, 0xa7, 0x25, 0x39, 0x39, 0x44, 0xb6,
	0xc0, 0x98, 0xf2, 0x44, 0x75, 0x16, 0xcd, 0xb6, 0xb2, 0x37, 0xa7, 0x3c, 0x91, 0x9d, 0xc1, 0xdd,
	0x9f, 0x75, 0x58, 0xd1, 0x9d, 0x0a, 0xc8, 0x05, 0xb4, 0xcb, 0x63, 0x4b, 0xb6, 0x4b, 0x9d, 0x7f,
	0xe4, 0xa6, 0xe9, 0xee, 0x3c, 0x69, 0xcf, 0xe7, 0xfd, 0x00, 0xda, 0xce, 0x53, 0x01, 0x1f, 0x19,
	0xd6, 0xee, 0x3f, 0xb9, 0x5d, 0xdd, 0x79, 0xc3, 0x4b, 0x16, 0x05, 0xe4, 0x03, 0x18, 0x47, 0xf3,
	0x26, 0x92, 0xcd, 0x92, 0x7f, 0xf5, 0xd8, 0x3f, 0xe6, 0xfa, 0x0e, 0x1a, 0x36, 0x4e, 0x18, 0x13,
	0x7f, 0xe8, 0xf7, 0x11, 0x8c, 0x31, 0x8a, 0xfc, 0xe0, 0x3e, 0xb4, 0x77, 0xbb, 0x15, 0xc8, 0xd1,
	0x37, 0x76, 0x4e, 0x3f, 0x02, 0x32, 0x46, 0x71, 0x7f, 0xb6, 0x9e, 0x09, 0x22, 0x25, 0xdd, 0xa7,
	0x9f, 0x41, 0x67, 0x8c, 0xa2, 0x02, 0x3d, 0x93, 0xc1, 0xf3, 0xf3, 0x79, 0xb8, 0xf5, 0x6d, 0x53,
	0xd9, 0x47, 0xf2, 0x31, 0xf2, 0x63, 0x36, 0x0b, 0x46, 0x21, 0xcb, 0xdf, 0x93, 0x49, 0x43, 0x7d,
	0xf7, 0x7e, 0x07, 0x00, 0x00, 0xff, 0xff, 0x8e, 0xfb, 0xd4, 0x44, 0xaa, 0x06, 0x00, 0x00,
}
