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
// 	protoc-gen-go v1.27.1
// 	protoc        (unknown)
// source: orc8r/protos/service_status.proto

package protos

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

// ServiceResult enumeration as defined in service "result" by systemd
type ServiceExitStatus_ServiceResult int32

const (
	ServiceExitStatus_UNUSED          ServiceExitStatus_ServiceResult = 0
	ServiceExitStatus_SUCCESS         ServiceExitStatus_ServiceResult = 1
	ServiceExitStatus_PROTOCOL        ServiceExitStatus_ServiceResult = 2
	ServiceExitStatus_TIMEOUT         ServiceExitStatus_ServiceResult = 3
	ServiceExitStatus_EXIT_CODE       ServiceExitStatus_ServiceResult = 4
	ServiceExitStatus_SIGNAL          ServiceExitStatus_ServiceResult = 5
	ServiceExitStatus_CORE_DUMP       ServiceExitStatus_ServiceResult = 6
	ServiceExitStatus_WATCHDOG        ServiceExitStatus_ServiceResult = 7
	ServiceExitStatus_START_LIMIT_HIT ServiceExitStatus_ServiceResult = 8
	ServiceExitStatus_RESOURCES       ServiceExitStatus_ServiceResult = 9
)

// Enum value maps for ServiceExitStatus_ServiceResult.
var (
	ServiceExitStatus_ServiceResult_name = map[int32]string{
		0: "UNUSED",
		1: "SUCCESS",
		2: "PROTOCOL",
		3: "TIMEOUT",
		4: "EXIT_CODE",
		5: "SIGNAL",
		6: "CORE_DUMP",
		7: "WATCHDOG",
		8: "START_LIMIT_HIT",
		9: "RESOURCES",
	}
	ServiceExitStatus_ServiceResult_value = map[string]int32{
		"UNUSED":          0,
		"SUCCESS":         1,
		"PROTOCOL":        2,
		"TIMEOUT":         3,
		"EXIT_CODE":       4,
		"SIGNAL":          5,
		"CORE_DUMP":       6,
		"WATCHDOG":        7,
		"START_LIMIT_HIT": 8,
		"RESOURCES":       9,
	}
)

func (x ServiceExitStatus_ServiceResult) Enum() *ServiceExitStatus_ServiceResult {
	p := new(ServiceExitStatus_ServiceResult)
	*p = x
	return p
}

func (x ServiceExitStatus_ServiceResult) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (ServiceExitStatus_ServiceResult) Descriptor() protoreflect.EnumDescriptor {
	return file_orc8r_protos_service_status_proto_enumTypes[0].Descriptor()
}

func (ServiceExitStatus_ServiceResult) Type() protoreflect.EnumType {
	return &file_orc8r_protos_service_status_proto_enumTypes[0]
}

func (x ServiceExitStatus_ServiceResult) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use ServiceExitStatus_ServiceResult.Descriptor instead.
func (ServiceExitStatus_ServiceResult) EnumDescriptor() ([]byte, []int) {
	return file_orc8r_protos_service_status_proto_rawDescGZIP(), []int{0, 0}
}

// ExitCode enumeration as defined in service "result" by systemd
type ServiceExitStatus_ExitCode int32

const (
	ServiceExitStatus_UNUSED_EXIT_CODE ServiceExitStatus_ExitCode = 0
	ServiceExitStatus_EXITED           ServiceExitStatus_ExitCode = 1
	ServiceExitStatus_KILLED           ServiceExitStatus_ExitCode = 2
	ServiceExitStatus_DUMPED           ServiceExitStatus_ExitCode = 3
)

// Enum value maps for ServiceExitStatus_ExitCode.
var (
	ServiceExitStatus_ExitCode_name = map[int32]string{
		0: "UNUSED_EXIT_CODE",
		1: "EXITED",
		2: "KILLED",
		3: "DUMPED",
	}
	ServiceExitStatus_ExitCode_value = map[string]int32{
		"UNUSED_EXIT_CODE": 0,
		"EXITED":           1,
		"KILLED":           2,
		"DUMPED":           3,
	}
)

func (x ServiceExitStatus_ExitCode) Enum() *ServiceExitStatus_ExitCode {
	p := new(ServiceExitStatus_ExitCode)
	*p = x
	return p
}

func (x ServiceExitStatus_ExitCode) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (ServiceExitStatus_ExitCode) Descriptor() protoreflect.EnumDescriptor {
	return file_orc8r_protos_service_status_proto_enumTypes[1].Descriptor()
}

func (ServiceExitStatus_ExitCode) Type() protoreflect.EnumType {
	return &file_orc8r_protos_service_status_proto_enumTypes[1]
}

func (x ServiceExitStatus_ExitCode) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use ServiceExitStatus_ExitCode.Descriptor instead.
func (ServiceExitStatus_ExitCode) EnumDescriptor() ([]byte, []int) {
	return file_orc8r_protos_service_status_proto_rawDescGZIP(), []int{0, 1}
}

type ServiceExitStatus struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	LatestServiceResult ServiceExitStatus_ServiceResult `protobuf:"varint,1,opt,name=latest_service_result,json=latestServiceResult,proto3,enum=magma.orc8r.ServiceExitStatus_ServiceResult" json:"latest_service_result,omitempty"`
	LatestExitCode      ServiceExitStatus_ExitCode      `protobuf:"varint,2,opt,name=latest_exit_code,json=latestExitCode,proto3,enum=magma.orc8r.ServiceExitStatus_ExitCode" json:"latest_exit_code,omitempty"`
	// Optional return code returned by the service during exit
	LatestRc uint32 `protobuf:"varint,3,opt,name=latest_rc,json=latestRc,proto3" json:"latest_rc,omitempty"`
	// Clean exit, e.g. SIGNKILL
	NumCleanExits uint32 `protobuf:"varint,4,opt,name=num_clean_exits,json=numCleanExits,proto3" json:"num_clean_exits,omitempty"`
	// Unclean exit e.g. CORE_DUMP or non zero exit code.
	NumFailExits uint32 `protobuf:"varint,5,opt,name=num_fail_exits,json=numFailExits,proto3" json:"num_fail_exits,omitempty"`
}

func (x *ServiceExitStatus) Reset() {
	*x = ServiceExitStatus{}
	if protoimpl.UnsafeEnabled {
		mi := &file_orc8r_protos_service_status_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ServiceExitStatus) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ServiceExitStatus) ProtoMessage() {}

func (x *ServiceExitStatus) ProtoReflect() protoreflect.Message {
	mi := &file_orc8r_protos_service_status_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ServiceExitStatus.ProtoReflect.Descriptor instead.
func (*ServiceExitStatus) Descriptor() ([]byte, []int) {
	return file_orc8r_protos_service_status_proto_rawDescGZIP(), []int{0}
}

func (x *ServiceExitStatus) GetLatestServiceResult() ServiceExitStatus_ServiceResult {
	if x != nil {
		return x.LatestServiceResult
	}
	return ServiceExitStatus_UNUSED
}

func (x *ServiceExitStatus) GetLatestExitCode() ServiceExitStatus_ExitCode {
	if x != nil {
		return x.LatestExitCode
	}
	return ServiceExitStatus_UNUSED_EXIT_CODE
}

func (x *ServiceExitStatus) GetLatestRc() uint32 {
	if x != nil {
		return x.LatestRc
	}
	return 0
}

func (x *ServiceExitStatus) GetNumCleanExits() uint32 {
	if x != nil {
		return x.NumCleanExits
	}
	return 0
}

func (x *ServiceExitStatus) GetNumFailExits() uint32 {
	if x != nil {
		return x.NumFailExits
	}
	return 0
}

var File_orc8r_protos_service_status_proto protoreflect.FileDescriptor

var file_orc8r_protos_service_status_proto_rawDesc = []byte{
	0x0a, 0x21, 0x6f, 0x72, 0x63, 0x38, 0x72, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x73, 0x2f, 0x73,
	0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x5f, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x12, 0x0b, 0x6d, 0x61, 0x67, 0x6d, 0x61, 0x2e, 0x6f, 0x72, 0x63, 0x38, 0x72,
	0x22, 0x9b, 0x04, 0x0a, 0x11, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x45, 0x78, 0x69, 0x74,
	0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x12, 0x60, 0x0a, 0x15, 0x6c, 0x61, 0x74, 0x65, 0x73, 0x74,
	0x5f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x5f, 0x72, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x2c, 0x2e, 0x6d, 0x61, 0x67, 0x6d, 0x61, 0x2e, 0x6f, 0x72,
	0x63, 0x38, 0x72, 0x2e, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x45, 0x78, 0x69, 0x74, 0x53,
	0x74, 0x61, 0x74, 0x75, 0x73, 0x2e, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x52, 0x65, 0x73,
	0x75, 0x6c, 0x74, 0x52, 0x13, 0x6c, 0x61, 0x74, 0x65, 0x73, 0x74, 0x53, 0x65, 0x72, 0x76, 0x69,
	0x63, 0x65, 0x52, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x12, 0x51, 0x0a, 0x10, 0x6c, 0x61, 0x74, 0x65,
	0x73, 0x74, 0x5f, 0x65, 0x78, 0x69, 0x74, 0x5f, 0x63, 0x6f, 0x64, 0x65, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x0e, 0x32, 0x27, 0x2e, 0x6d, 0x61, 0x67, 0x6d, 0x61, 0x2e, 0x6f, 0x72, 0x63, 0x38, 0x72,
	0x2e, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x45, 0x78, 0x69, 0x74, 0x53, 0x74, 0x61, 0x74,
	0x75, 0x73, 0x2e, 0x45, 0x78, 0x69, 0x74, 0x43, 0x6f, 0x64, 0x65, 0x52, 0x0e, 0x6c, 0x61, 0x74,
	0x65, 0x73, 0x74, 0x45, 0x78, 0x69, 0x74, 0x43, 0x6f, 0x64, 0x65, 0x12, 0x1b, 0x0a, 0x09, 0x6c,
	0x61, 0x74, 0x65, 0x73, 0x74, 0x5f, 0x72, 0x63, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x08,
	0x6c, 0x61, 0x74, 0x65, 0x73, 0x74, 0x52, 0x63, 0x12, 0x26, 0x0a, 0x0f, 0x6e, 0x75, 0x6d, 0x5f,
	0x63, 0x6c, 0x65, 0x61, 0x6e, 0x5f, 0x65, 0x78, 0x69, 0x74, 0x73, 0x18, 0x04, 0x20, 0x01, 0x28,
	0x0d, 0x52, 0x0d, 0x6e, 0x75, 0x6d, 0x43, 0x6c, 0x65, 0x61, 0x6e, 0x45, 0x78, 0x69, 0x74, 0x73,
	0x12, 0x24, 0x0a, 0x0e, 0x6e, 0x75, 0x6d, 0x5f, 0x66, 0x61, 0x69, 0x6c, 0x5f, 0x65, 0x78, 0x69,
	0x74, 0x73, 0x18, 0x05, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x0c, 0x6e, 0x75, 0x6d, 0x46, 0x61, 0x69,
	0x6c, 0x45, 0x78, 0x69, 0x74, 0x73, 0x22, 0x9f, 0x01, 0x0a, 0x0d, 0x53, 0x65, 0x72, 0x76, 0x69,
	0x63, 0x65, 0x52, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x12, 0x0a, 0x0a, 0x06, 0x55, 0x4e, 0x55, 0x53,
	0x45, 0x44, 0x10, 0x00, 0x12, 0x0b, 0x0a, 0x07, 0x53, 0x55, 0x43, 0x43, 0x45, 0x53, 0x53, 0x10,
	0x01, 0x12, 0x0c, 0x0a, 0x08, 0x50, 0x52, 0x4f, 0x54, 0x4f, 0x43, 0x4f, 0x4c, 0x10, 0x02, 0x12,
	0x0b, 0x0a, 0x07, 0x54, 0x49, 0x4d, 0x45, 0x4f, 0x55, 0x54, 0x10, 0x03, 0x12, 0x0d, 0x0a, 0x09,
	0x45, 0x58, 0x49, 0x54, 0x5f, 0x43, 0x4f, 0x44, 0x45, 0x10, 0x04, 0x12, 0x0a, 0x0a, 0x06, 0x53,
	0x49, 0x47, 0x4e, 0x41, 0x4c, 0x10, 0x05, 0x12, 0x0d, 0x0a, 0x09, 0x43, 0x4f, 0x52, 0x45, 0x5f,
	0x44, 0x55, 0x4d, 0x50, 0x10, 0x06, 0x12, 0x0c, 0x0a, 0x08, 0x57, 0x41, 0x54, 0x43, 0x48, 0x44,
	0x4f, 0x47, 0x10, 0x07, 0x12, 0x13, 0x0a, 0x0f, 0x53, 0x54, 0x41, 0x52, 0x54, 0x5f, 0x4c, 0x49,
	0x4d, 0x49, 0x54, 0x5f, 0x48, 0x49, 0x54, 0x10, 0x08, 0x12, 0x0d, 0x0a, 0x09, 0x52, 0x45, 0x53,
	0x4f, 0x55, 0x52, 0x43, 0x45, 0x53, 0x10, 0x09, 0x22, 0x44, 0x0a, 0x08, 0x45, 0x78, 0x69, 0x74,
	0x43, 0x6f, 0x64, 0x65, 0x12, 0x14, 0x0a, 0x10, 0x55, 0x4e, 0x55, 0x53, 0x45, 0x44, 0x5f, 0x45,
	0x58, 0x49, 0x54, 0x5f, 0x43, 0x4f, 0x44, 0x45, 0x10, 0x00, 0x12, 0x0a, 0x0a, 0x06, 0x45, 0x58,
	0x49, 0x54, 0x45, 0x44, 0x10, 0x01, 0x12, 0x0a, 0x0a, 0x06, 0x4b, 0x49, 0x4c, 0x4c, 0x45, 0x44,
	0x10, 0x02, 0x12, 0x0a, 0x0a, 0x06, 0x44, 0x55, 0x4d, 0x50, 0x45, 0x44, 0x10, 0x03, 0x42, 0x1b,
	0x5a, 0x19, 0x6d, 0x61, 0x67, 0x6d, 0x61, 0x2f, 0x6f, 0x72, 0x63, 0x38, 0x72, 0x2f, 0x6c, 0x69,
	0x62, 0x2f, 0x67, 0x6f, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x73, 0x62, 0x06, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x33,
}

var (
	file_orc8r_protos_service_status_proto_rawDescOnce sync.Once
	file_orc8r_protos_service_status_proto_rawDescData = file_orc8r_protos_service_status_proto_rawDesc
)

func file_orc8r_protos_service_status_proto_rawDescGZIP() []byte {
	file_orc8r_protos_service_status_proto_rawDescOnce.Do(func() {
		file_orc8r_protos_service_status_proto_rawDescData = protoimpl.X.CompressGZIP(file_orc8r_protos_service_status_proto_rawDescData)
	})
	return file_orc8r_protos_service_status_proto_rawDescData
}

var file_orc8r_protos_service_status_proto_enumTypes = make([]protoimpl.EnumInfo, 2)
var file_orc8r_protos_service_status_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_orc8r_protos_service_status_proto_goTypes = []interface{}{
	(ServiceExitStatus_ServiceResult)(0), // 0: magma.orc8r.ServiceExitStatus.ServiceResult
	(ServiceExitStatus_ExitCode)(0),      // 1: magma.orc8r.ServiceExitStatus.ExitCode
	(*ServiceExitStatus)(nil),            // 2: magma.orc8r.ServiceExitStatus
}
var file_orc8r_protos_service_status_proto_depIdxs = []int32{
	0, // 0: magma.orc8r.ServiceExitStatus.latest_service_result:type_name -> magma.orc8r.ServiceExitStatus.ServiceResult
	1, // 1: magma.orc8r.ServiceExitStatus.latest_exit_code:type_name -> magma.orc8r.ServiceExitStatus.ExitCode
	2, // [2:2] is the sub-list for method output_type
	2, // [2:2] is the sub-list for method input_type
	2, // [2:2] is the sub-list for extension type_name
	2, // [2:2] is the sub-list for extension extendee
	0, // [0:2] is the sub-list for field type_name
}

func init() { file_orc8r_protos_service_status_proto_init() }
func file_orc8r_protos_service_status_proto_init() {
	if File_orc8r_protos_service_status_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_orc8r_protos_service_status_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ServiceExitStatus); i {
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
			RawDescriptor: file_orc8r_protos_service_status_proto_rawDesc,
			NumEnums:      2,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_orc8r_protos_service_status_proto_goTypes,
		DependencyIndexes: file_orc8r_protos_service_status_proto_depIdxs,
		EnumInfos:         file_orc8r_protos_service_status_proto_enumTypes,
		MessageInfos:      file_orc8r_protos_service_status_proto_msgTypes,
	}.Build()
	File_orc8r_protos_service_status_proto = out.File
	file_orc8r_protos_service_status_proto_rawDesc = nil
	file_orc8r_protos_service_status_proto_goTypes = nil
	file_orc8r_protos_service_status_proto_depIdxs = nil
}
