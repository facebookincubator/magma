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
// 	protoc-gen-go v1.26.0
// 	protoc        (unknown)
// source: protos/eap_provider.proto

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

type EapAkaSubtype int32

const (
	EapAkaSubtype_eap_aka_subtype_undefined   EapAkaSubtype = 0
	EapAkaSubtype_aka_challenge               EapAkaSubtype = 1
	EapAkaSubtype_aka_authentication_reject   EapAkaSubtype = 2
	EapAkaSubtype_aka_synchronization_failure EapAkaSubtype = 4
	EapAkaSubtype_SubtypeIdentity             EapAkaSubtype = 5
	EapAkaSubtype_SubtypeNotification         EapAkaSubtype = 12
	EapAkaSubtype_aka_reauthentication        EapAkaSubtype = 13
	EapAkaSubtype_aka_cient_error             EapAkaSubtype = 14
)

// Enum value maps for EapAkaSubtype.
var (
	EapAkaSubtype_name = map[int32]string{
		0:  "eap_aka_subtype_undefined",
		1:  "aka_challenge",
		2:  "aka_authentication_reject",
		4:  "aka_synchronization_failure",
		5:  "SubtypeIdentity",
		12: "SubtypeNotification",
		13: "aka_reauthentication",
		14: "aka_cient_error",
	}
	EapAkaSubtype_value = map[string]int32{
		"eap_aka_subtype_undefined":   0,
		"aka_challenge":               1,
		"aka_authentication_reject":   2,
		"aka_synchronization_failure": 4,
		"SubtypeIdentity":             5,
		"SubtypeNotification":         12,
		"aka_reauthentication":        13,
		"aka_cient_error":             14,
	}
)

func (x EapAkaSubtype) Enum() *EapAkaSubtype {
	p := new(EapAkaSubtype)
	*p = x
	return p
}

func (x EapAkaSubtype) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (EapAkaSubtype) Descriptor() protoreflect.EnumDescriptor {
	return file_protos_eap_provider_proto_enumTypes[0].Descriptor()
}

func (EapAkaSubtype) Type() protoreflect.EnumType {
	return &file_protos_eap_provider_proto_enumTypes[0]
}

func (x EapAkaSubtype) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use EapAkaSubtype.Descriptor instead.
func (EapAkaSubtype) EnumDescriptor() ([]byte, []int) {
	return file_protos_eap_provider_proto_rawDescGZIP(), []int{0}
}

var File_protos_eap_provider_proto protoreflect.FileDescriptor

var file_protos_eap_provider_proto_rawDesc = []byte{
	0x0a, 0x19, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x73, 0x2f, 0x65, 0x61, 0x70, 0x5f, 0x70, 0x72, 0x6f,
	0x76, 0x69, 0x64, 0x65, 0x72, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x03, 0x65, 0x61, 0x70,
	0x2a, 0xe0, 0x01, 0x0a, 0x0f, 0x65, 0x61, 0x70, 0x5f, 0x61, 0x6b, 0x61, 0x5f, 0x73, 0x75, 0x62,
	0x74, 0x79, 0x70, 0x65, 0x12, 0x1d, 0x0a, 0x19, 0x65, 0x61, 0x70, 0x5f, 0x61, 0x6b, 0x61, 0x5f,
	0x73, 0x75, 0x62, 0x74, 0x79, 0x70, 0x65, 0x5f, 0x75, 0x6e, 0x64, 0x65, 0x66, 0x69, 0x6e, 0x65,
	0x64, 0x10, 0x00, 0x12, 0x11, 0x0a, 0x0d, 0x61, 0x6b, 0x61, 0x5f, 0x63, 0x68, 0x61, 0x6c, 0x6c,
	0x65, 0x6e, 0x67, 0x65, 0x10, 0x01, 0x12, 0x1d, 0x0a, 0x19, 0x61, 0x6b, 0x61, 0x5f, 0x61, 0x75,
	0x74, 0x68, 0x65, 0x6e, 0x74, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x5f, 0x72, 0x65, 0x6a,
	0x65, 0x63, 0x74, 0x10, 0x02, 0x12, 0x1f, 0x0a, 0x1b, 0x61, 0x6b, 0x61, 0x5f, 0x73, 0x79, 0x6e,
	0x63, 0x68, 0x72, 0x6f, 0x6e, 0x69, 0x7a, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x5f, 0x66, 0x61, 0x69,
	0x6c, 0x75, 0x72, 0x65, 0x10, 0x04, 0x12, 0x13, 0x0a, 0x0f, 0x53, 0x75, 0x62, 0x74, 0x79, 0x70,
	0x65, 0x49, 0x64, 0x65, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x10, 0x05, 0x12, 0x17, 0x0a, 0x13, 0x53,
	0x75, 0x62, 0x74, 0x79, 0x70, 0x65, 0x4e, 0x6f, 0x74, 0x69, 0x66, 0x69, 0x63, 0x61, 0x74, 0x69,
	0x6f, 0x6e, 0x10, 0x0c, 0x12, 0x18, 0x0a, 0x14, 0x61, 0x6b, 0x61, 0x5f, 0x72, 0x65, 0x61, 0x75,
	0x74, 0x68, 0x65, 0x6e, 0x74, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x10, 0x0d, 0x12, 0x13,
	0x0a, 0x0f, 0x61, 0x6b, 0x61, 0x5f, 0x63, 0x69, 0x65, 0x6e, 0x74, 0x5f, 0x65, 0x72, 0x72, 0x6f,
	0x72, 0x10, 0x0e, 0x42, 0x30, 0x5a, 0x2e, 0x6d, 0x61, 0x67, 0x6d, 0x61, 0x2f, 0x66, 0x65, 0x67,
	0x2f, 0x67, 0x61, 0x74, 0x65, 0x77, 0x61, 0x79, 0x2f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65,
	0x73, 0x2f, 0x65, 0x61, 0x70, 0x2f, 0x70, 0x72, 0x6f, 0x76, 0x69, 0x64, 0x65, 0x72, 0x2f, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x73, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_protos_eap_provider_proto_rawDescOnce sync.Once
	file_protos_eap_provider_proto_rawDescData = file_protos_eap_provider_proto_rawDesc
)

func file_protos_eap_provider_proto_rawDescGZIP() []byte {
	file_protos_eap_provider_proto_rawDescOnce.Do(func() {
		file_protos_eap_provider_proto_rawDescData = protoimpl.X.CompressGZIP(file_protos_eap_provider_proto_rawDescData)
	})
	return file_protos_eap_provider_proto_rawDescData
}

var file_protos_eap_provider_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_protos_eap_provider_proto_goTypes = []interface{}{
	(EapAkaSubtype)(0), // 0: eap.eap_aka_subtype
}
var file_protos_eap_provider_proto_depIdxs = []int32{
	0, // [0:0] is the sub-list for method output_type
	0, // [0:0] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_protos_eap_provider_proto_init() }
func file_protos_eap_provider_proto_init() {
	if File_protos_eap_provider_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_protos_eap_provider_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   0,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_protos_eap_provider_proto_goTypes,
		DependencyIndexes: file_protos_eap_provider_proto_depIdxs,
		EnumInfos:         file_protos_eap_provider_proto_enumTypes,
	}.Build()
	File_protos_eap_provider_proto = out.File
	file_protos_eap_provider_proto_rawDesc = nil
	file_protos_eap_provider_proto_goTypes = nil
	file_protos_eap_provider_proto_depIdxs = nil
}
