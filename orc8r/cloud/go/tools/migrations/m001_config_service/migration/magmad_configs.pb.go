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

//

// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.25.0
// 	protoc        (unknown)
// source: magmad_configs.proto

package migration

import (
	proto "github.com/golang/protobuf/proto"
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

// This is a compile-time assertion that a sufficiently up-to-date version
// of the legacy proto package is being used.
const _ = proto.ProtoPackageIsVersion4

// Container message type for namespaced network and gateway configs.
type Config struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Namespaced map of serialized configs
	ConfigsByKey map[string][]byte `protobuf:"bytes,1,rep,name=configs_by_key,json=configsByKey,proto3" json:"configs_by_key,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
}

func (x *Config) Reset() {
	*x = Config{}
	if protoimpl.UnsafeEnabled {
		mi := &file_magmad_configs_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Config) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Config) ProtoMessage() {}

func (x *Config) ProtoReflect() protoreflect.Message {
	mi := &file_magmad_configs_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Config.ProtoReflect.Descriptor instead.
func (*Config) Descriptor() ([]byte, []int) {
	return file_magmad_configs_proto_rawDescGZIP(), []int{0}
}

func (x *Config) GetConfigsByKey() map[string][]byte {
	if x != nil {
		return x.ConfigsByKey
	}
	return nil
}

var File_magmad_configs_proto protoreflect.FileDescriptor

var file_magmad_configs_proto_rawDesc = []byte{
	0x0a, 0x14, 0x6d, 0x61, 0x67, 0x6d, 0x61, 0x64, 0x5f, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x73,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x10, 0x6d, 0x61, 0x67, 0x6d, 0x61, 0x2e, 0x6d, 0x69,
	0x67, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x22, 0x9b, 0x01, 0x0a, 0x06, 0x43, 0x6f, 0x6e,
	0x66, 0x69, 0x67, 0x12, 0x50, 0x0a, 0x0e, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x73, 0x5f, 0x62,
	0x79, 0x5f, 0x6b, 0x65, 0x79, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x2a, 0x2e, 0x6d, 0x61,
	0x67, 0x6d, 0x61, 0x2e, 0x6d, 0x69, 0x67, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x43,
	0x6f, 0x6e, 0x66, 0x69, 0x67, 0x2e, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x73, 0x42, 0x79, 0x4b,
	0x65, 0x79, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x52, 0x0c, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x73,
	0x42, 0x79, 0x4b, 0x65, 0x79, 0x1a, 0x3f, 0x0a, 0x11, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x73,
	0x42, 0x79, 0x4b, 0x65, 0x79, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x12, 0x10, 0x0a, 0x03, 0x6b, 0x65,
	0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x6b, 0x65, 0x79, 0x12, 0x14, 0x0a, 0x05,
	0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x05, 0x76, 0x61, 0x6c,
	0x75, 0x65, 0x3a, 0x02, 0x38, 0x01, 0x42, 0x0b, 0x5a, 0x09, 0x6d, 0x69, 0x67, 0x72, 0x61, 0x74,
	0x69, 0x6f, 0x6e, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_magmad_configs_proto_rawDescOnce sync.Once
	file_magmad_configs_proto_rawDescData = file_magmad_configs_proto_rawDesc
)

func file_magmad_configs_proto_rawDescGZIP() []byte {
	file_magmad_configs_proto_rawDescOnce.Do(func() {
		file_magmad_configs_proto_rawDescData = protoimpl.X.CompressGZIP(file_magmad_configs_proto_rawDescData)
	})
	return file_magmad_configs_proto_rawDescData
}

var file_magmad_configs_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_magmad_configs_proto_goTypes = []interface{}{
	(*Config)(nil), // 0: magma.migrations.Config
	nil,            // 1: magma.migrations.Config.ConfigsByKeyEntry
}
var file_magmad_configs_proto_depIdxs = []int32{
	1, // 0: magma.migrations.Config.configs_by_key:type_name -> magma.migrations.Config.ConfigsByKeyEntry
	1, // [1:1] is the sub-list for method output_type
	1, // [1:1] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_magmad_configs_proto_init() }
func file_magmad_configs_proto_init() {
	if File_magmad_configs_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_magmad_configs_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Config); i {
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
			RawDescriptor: file_magmad_configs_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_magmad_configs_proto_goTypes,
		DependencyIndexes: file_magmad_configs_proto_depIdxs,
		MessageInfos:      file_magmad_configs_proto_msgTypes,
	}.Build()
	File_magmad_configs_proto = out.File
	file_magmad_configs_proto_rawDesc = nil
	file_magmad_configs_proto_goTypes = nil
	file_magmad_configs_proto_depIdxs = nil
}
