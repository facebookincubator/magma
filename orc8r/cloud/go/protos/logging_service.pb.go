// Code generated by protoc-gen-go. DO NOT EDIT.
// source: orc8r/protos/logging_service.proto

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

// Where to log to. Currently only supports scribe.
type LoggerDestination int32

const (
	LoggerDestination_SCRIBE LoggerDestination = 0
)

var LoggerDestination_name = map[int32]string{
	0: "SCRIBE",
}

var LoggerDestination_value = map[string]int32{
	"SCRIBE": 0,
}

func (x LoggerDestination) String() string {
	return proto.EnumName(LoggerDestination_name, int32(x))
}

func (LoggerDestination) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_e52e10d1e5f65add, []int{0}
}

type LogEntry struct {
	// category of the log entry
	Category string `protobuf:"bytes,1,opt,name=category,proto3" json:"category,omitempty"`
	// required unix timestamp in seconds of the entry
	Time int64 `protobuf:"varint,3,opt,name=time,proto3" json:"time,omitempty"`
	// optional hardware id of the gateway if the LogEntry comes from a gateway
	HwId string `protobuf:"bytes,4,opt,name=hw_id,json=hwId,proto3" json:"hw_id,omitempty"`
	// optinoal map of normal(string) values
	NormalMap map[string]string `protobuf:"bytes,5,rep,name=normal_map,json=normalMap,proto3" json:"normal_map,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
	// optional map of int values
	IntMap map[string]int64 `protobuf:"bytes,6,rep,name=int_map,json=intMap,proto3" json:"int_map,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"varint,2,opt,name=value,proto3"`
	// optional a set of string values, usually used for *gk_on* or *quick_experiment*
	TagSet []string `protobuf:"bytes,7,rep,name=tag_set,json=tagSet,proto3" json:"tag_set,omitempty"`
	// optional a vector of strings, usually used for stack traces
	Normvector           []string `protobuf:"bytes,8,rep,name=normvector,proto3" json:"normvector,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *LogEntry) Reset()         { *m = LogEntry{} }
func (m *LogEntry) String() string { return proto.CompactTextString(m) }
func (*LogEntry) ProtoMessage()    {}
func (*LogEntry) Descriptor() ([]byte, []int) {
	return fileDescriptor_e52e10d1e5f65add, []int{0}
}

func (m *LogEntry) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_LogEntry.Unmarshal(m, b)
}
func (m *LogEntry) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_LogEntry.Marshal(b, m, deterministic)
}
func (m *LogEntry) XXX_Merge(src proto.Message) {
	xxx_messageInfo_LogEntry.Merge(m, src)
}
func (m *LogEntry) XXX_Size() int {
	return xxx_messageInfo_LogEntry.Size(m)
}
func (m *LogEntry) XXX_DiscardUnknown() {
	xxx_messageInfo_LogEntry.DiscardUnknown(m)
}

var xxx_messageInfo_LogEntry proto.InternalMessageInfo

func (m *LogEntry) GetCategory() string {
	if m != nil {
		return m.Category
	}
	return ""
}

func (m *LogEntry) GetTime() int64 {
	if m != nil {
		return m.Time
	}
	return 0
}

func (m *LogEntry) GetHwId() string {
	if m != nil {
		return m.HwId
	}
	return ""
}

func (m *LogEntry) GetNormalMap() map[string]string {
	if m != nil {
		return m.NormalMap
	}
	return nil
}

func (m *LogEntry) GetIntMap() map[string]int64 {
	if m != nil {
		return m.IntMap
	}
	return nil
}

func (m *LogEntry) GetTagSet() []string {
	if m != nil {
		return m.TagSet
	}
	return nil
}

func (m *LogEntry) GetNormvector() []string {
	if m != nil {
		return m.Normvector
	}
	return nil
}

type LogRequest struct {
	Entries              []*LogEntry       `protobuf:"bytes,1,rep,name=Entries,proto3" json:"Entries,omitempty"`
	Destination          LoggerDestination `protobuf:"varint,2,opt,name=Destination,proto3,enum=magma.orc8r.LoggerDestination" json:"Destination,omitempty"`
	XXX_NoUnkeyedLiteral struct{}          `json:"-"`
	XXX_unrecognized     []byte            `json:"-"`
	XXX_sizecache        int32             `json:"-"`
}

func (m *LogRequest) Reset()         { *m = LogRequest{} }
func (m *LogRequest) String() string { return proto.CompactTextString(m) }
func (*LogRequest) ProtoMessage()    {}
func (*LogRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_e52e10d1e5f65add, []int{1}
}

func (m *LogRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_LogRequest.Unmarshal(m, b)
}
func (m *LogRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_LogRequest.Marshal(b, m, deterministic)
}
func (m *LogRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_LogRequest.Merge(m, src)
}
func (m *LogRequest) XXX_Size() int {
	return xxx_messageInfo_LogRequest.Size(m)
}
func (m *LogRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_LogRequest.DiscardUnknown(m)
}

var xxx_messageInfo_LogRequest proto.InternalMessageInfo

func (m *LogRequest) GetEntries() []*LogEntry {
	if m != nil {
		return m.Entries
	}
	return nil
}

func (m *LogRequest) GetDestination() LoggerDestination {
	if m != nil {
		return m.Destination
	}
	return LoggerDestination_SCRIBE
}

func init() {
	proto.RegisterEnum("magma.orc8r.LoggerDestination", LoggerDestination_name, LoggerDestination_value)
	proto.RegisterType((*LogEntry)(nil), "magma.orc8r.LogEntry")
	proto.RegisterMapType((map[string]int64)(nil), "magma.orc8r.LogEntry.IntMapEntry")
	proto.RegisterMapType((map[string]string)(nil), "magma.orc8r.LogEntry.NormalMapEntry")
	proto.RegisterType((*LogRequest)(nil), "magma.orc8r.LogRequest")
}

func init() { proto.RegisterFile("orc8r/protos/logging_service.proto", fileDescriptor_e52e10d1e5f65add) }

var fileDescriptor_e52e10d1e5f65add = []byte{
	// 415 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x84, 0x52, 0xc1, 0x6e, 0x13, 0x31,
	0x10, 0xed, 0xd6, 0xc9, 0x26, 0x99, 0x48, 0x51, 0x3b, 0x80, 0x6a, 0x16, 0x51, 0x42, 0xc4, 0x21,
	0xe2, 0xb0, 0x2b, 0xa5, 0x97, 0x52, 0x71, 0x40, 0x2d, 0x39, 0x44, 0x0a, 0x1c, 0x1c, 0x89, 0x03,
	0x97, 0xc8, 0x6c, 0x2c, 0xd7, 0x22, 0xeb, 0x09, 0x5e, 0x27, 0x55, 0x4f, 0x48, 0x7c, 0x39, 0x8a,
	0xb7, 0xa9, 0x36, 0x25, 0x12, 0x27, 0xfb, 0x8d, 0xdf, 0x7b, 0x9e, 0x19, 0x3d, 0x18, 0x90, 0xcb,
	0x2f, 0x5d, 0xb6, 0x72, 0xe4, 0xa9, 0xcc, 0x96, 0xa4, 0xb5, 0xb1, 0x7a, 0x5e, 0x2a, 0xb7, 0x31,
	0xb9, 0x4a, 0x43, 0x19, 0xbb, 0x85, 0xd4, 0x85, 0x4c, 0x03, 0x33, 0x79, 0xb9, 0x27, 0xc8, 0xa9,
	0x28, 0xc8, 0x56, 0xbc, 0xc1, 0x1f, 0x06, 0xed, 0x29, 0xe9, 0xb1, 0xf5, 0xee, 0x1e, 0x13, 0x68,
	0xe7, 0xd2, 0x2b, 0x4d, 0xee, 0x9e, 0x47, 0xfd, 0x68, 0xd8, 0x11, 0x8f, 0x18, 0x11, 0x1a, 0xde,
	0x14, 0x8a, 0xb3, 0x7e, 0x34, 0x64, 0x22, 0xdc, 0xf1, 0x19, 0x34, 0x6f, 0xef, 0xe6, 0x66, 0xc1,
	0x1b, 0x81, 0xdc, 0xb8, 0xbd, 0x9b, 0x2c, 0xf0, 0x06, 0xc0, 0x92, 0x2b, 0xe4, 0x72, 0x5e, 0xc8,
	0x15, 0x6f, 0xf6, 0xd9, 0xb0, 0x3b, 0x7a, 0x97, 0xd6, 0xda, 0x49, 0x77, 0xff, 0xa5, 0x5f, 0x03,
	0xef, 0x8b, 0x5c, 0x05, 0x28, 0x3a, 0x76, 0x87, 0xf1, 0x0a, 0x5a, 0xc6, 0xfa, 0xe0, 0x10, 0x07,
	0x87, 0xb7, 0x87, 0x1d, 0x26, 0xd6, 0x3f, 0xca, 0x63, 0x13, 0x00, 0x9e, 0x41, 0xcb, 0xcb, 0xed,
	0x3e, 0x3c, 0x6f, 0xf5, 0xd9, 0xb0, 0x23, 0x62, 0x2f, 0xf5, 0x4c, 0x79, 0x3c, 0xaf, 0x3a, 0xdb,
	0xa8, 0xdc, 0x93, 0xe3, 0xed, 0xf0, 0x56, 0xab, 0x24, 0x1f, 0xa1, 0xb7, 0xdf, 0x11, 0x9e, 0x00,
	0xfb, 0xa9, 0x76, 0xbb, 0xd8, 0x5e, 0xf1, 0x39, 0x34, 0x37, 0x72, 0xb9, 0x56, 0xfc, 0x38, 0xd4,
	0x2a, 0x70, 0x75, 0x7c, 0x19, 0x25, 0x1f, 0xa0, 0x5b, 0xeb, 0xe6, 0x7f, 0x52, 0x56, 0x93, 0x0e,
	0x7e, 0x03, 0x4c, 0x49, 0x0b, 0xf5, 0x6b, 0xad, 0x4a, 0x8f, 0x19, 0xb4, 0xb6, 0x16, 0x46, 0x95,
	0x3c, 0x0a, 0xb3, 0xbf, 0x38, 0x38, 0xbb, 0xd8, 0xb1, 0xf0, 0x13, 0x74, 0x3f, 0xab, 0xd2, 0x1b,
	0x2b, 0xbd, 0x21, 0x1b, 0xec, 0x7b, 0xa3, 0xf3, 0xa7, 0x22, 0xad, 0x5c, 0x8d, 0x25, 0xea, 0x92,
	0xf7, 0x6f, 0xe0, 0xf4, 0x1f, 0x06, 0x02, 0xc4, 0xb3, 0x1b, 0x31, 0xb9, 0x1e, 0x9f, 0x1c, 0x8d,
	0xc6, 0xd0, 0x9b, 0x56, 0x39, 0x9b, 0x55, 0x31, 0xc3, 0x0b, 0x60, 0x53, 0xd2, 0x78, 0xf6, 0xf4,
	0x9b, 0x87, 0x29, 0x92, 0xd3, 0xbd, 0x87, 0x6f, 0x64, 0x16, 0x83, 0xa3, 0xeb, 0xd7, 0xdf, 0x5f,
	0x85, 0x6a, 0x56, 0x05, 0x32, 0x5f, 0xd2, 0x7a, 0x91, 0x69, 0x7a, 0x48, 0xe6, 0x8f, 0x38, 0x9c,
	0x17, 0x7f, 0x03, 0x00, 0x00, 0xff, 0xff, 0x92, 0xa0, 0x62, 0xe2, 0xe1, 0x02, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// LoggingServiceClient is the client API for LoggingService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type LoggingServiceClient interface {
	// Log a list of LogEntry.
	Log(ctx context.Context, in *LogRequest, opts ...grpc.CallOption) (*Void, error)
}

type loggingServiceClient struct {
	cc *grpc.ClientConn
}

func NewLoggingServiceClient(cc *grpc.ClientConn) LoggingServiceClient {
	return &loggingServiceClient{cc}
}

func (c *loggingServiceClient) Log(ctx context.Context, in *LogRequest, opts ...grpc.CallOption) (*Void, error) {
	out := new(Void)
	err := c.cc.Invoke(ctx, "/magma.orc8r.LoggingService/Log", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// LoggingServiceServer is the server API for LoggingService service.
type LoggingServiceServer interface {
	// Log a list of LogEntry.
	Log(context.Context, *LogRequest) (*Void, error)
}

// UnimplementedLoggingServiceServer can be embedded to have forward compatible implementations.
type UnimplementedLoggingServiceServer struct {
}

func (*UnimplementedLoggingServiceServer) Log(ctx context.Context, req *LogRequest) (*Void, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Log not implemented")
}

func RegisterLoggingServiceServer(s *grpc.Server, srv LoggingServiceServer) {
	s.RegisterService(&_LoggingService_serviceDesc, srv)
}

func _LoggingService_Log_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(LogRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(LoggingServiceServer).Log(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/magma.orc8r.LoggingService/Log",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(LoggingServiceServer).Log(ctx, req.(*LogRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _LoggingService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "magma.orc8r.LoggingService",
	HandlerType: (*LoggingServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Log",
			Handler:    _LoggingService_Log_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "orc8r/protos/logging_service.proto",
}
