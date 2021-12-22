// Copyright 2018 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// Code generated by protoc-gen-go. DO NOT EDIT.
// cmd/protoc-gen-go/testdata/comments/deprecated.proto is a deprecated file.

package comments

import (
	reflect "reflect"
	sync "sync"

	protoreflect "github.com/schattian/protobuf/reflect/protoreflect"
	protoimpl "github.com/schattian/protobuf/runtime/protoimpl"
)

// Deprecated: Do not use.
type DeprecatedEnum int32

const (
	// Deprecated: Do not use.
	DeprecatedEnum_DEPRECATED DeprecatedEnum = 0
)

// Enum value maps for DeprecatedEnum.
var (
	DeprecatedEnum_name = map[int32]string{
		0: "DEPRECATED",
	}
	DeprecatedEnum_value = map[string]int32{
		"DEPRECATED": 0,
	}
)

func (x DeprecatedEnum) Enum() *DeprecatedEnum {
	p := new(DeprecatedEnum)
	*p = x
	return p
}

func (x DeprecatedEnum) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (DeprecatedEnum) Descriptor() protoreflect.EnumDescriptor {
	return file_cmd_protoc_gen_go_testdata_comments_deprecated_proto_enumTypes[0].Descriptor()
}

func (DeprecatedEnum) Type() protoreflect.EnumType {
	return &file_cmd_protoc_gen_go_testdata_comments_deprecated_proto_enumTypes[0]
}

func (x DeprecatedEnum) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use DeprecatedEnum.Descriptor instead.
func (DeprecatedEnum) EnumDescriptor() ([]byte, []int) {
	return file_cmd_protoc_gen_go_testdata_comments_deprecated_proto_rawDescGZIP(), []int{0}
}

// Deprecated: Do not use.
type DeprecatedMessage struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Deprecated: Do not use.
	DeprecatedField string `protobuf:"bytes,1,opt,name=deprecated_field,json=deprecatedField,proto3" json:"deprecated_field,omitempty"`
}

func (x *DeprecatedMessage) Reset() {
	*x = DeprecatedMessage{}
	if protoimpl.UnsafeEnabled {
		mi := &file_cmd_protoc_gen_go_testdata_comments_deprecated_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DeprecatedMessage) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DeprecatedMessage) ProtoMessage() {}

func (x *DeprecatedMessage) ProtoReflect() protoreflect.Message {
	mi := &file_cmd_protoc_gen_go_testdata_comments_deprecated_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DeprecatedMessage.ProtoReflect.Descriptor instead.
func (*DeprecatedMessage) Descriptor() ([]byte, []int) {
	return file_cmd_protoc_gen_go_testdata_comments_deprecated_proto_rawDescGZIP(), []int{0}
}

// Deprecated: Do not use.
func (x *DeprecatedMessage) GetDeprecatedField() string {
	if x != nil {
		return x.DeprecatedField
	}
	return ""
}

var File_cmd_protoc_gen_go_testdata_comments_deprecated_proto protoreflect.FileDescriptor

var file_cmd_protoc_gen_go_testdata_comments_deprecated_proto_rawDesc = []byte{
	0x0a, 0x34, 0x63, 0x6d, 0x64, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x63, 0x2d, 0x67, 0x65, 0x6e,
	0x2d, 0x67, 0x6f, 0x2f, 0x74, 0x65, 0x73, 0x74, 0x64, 0x61, 0x74, 0x61, 0x2f, 0x63, 0x6f, 0x6d,
	0x6d, 0x65, 0x6e, 0x74, 0x73, 0x2f, 0x64, 0x65, 0x70, 0x72, 0x65, 0x63, 0x61, 0x74, 0x65, 0x64,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x17, 0x67, 0x6f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x63, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x65, 0x6e, 0x74, 0x73, 0x22,
	0x46, 0x0a, 0x11, 0x44, 0x65, 0x70, 0x72, 0x65, 0x63, 0x61, 0x74, 0x65, 0x64, 0x4d, 0x65, 0x73,
	0x73, 0x61, 0x67, 0x65, 0x12, 0x2d, 0x0a, 0x10, 0x64, 0x65, 0x70, 0x72, 0x65, 0x63, 0x61, 0x74,
	0x65, 0x64, 0x5f, 0x66, 0x69, 0x65, 0x6c, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x42, 0x02,
	0x18, 0x01, 0x52, 0x0f, 0x64, 0x65, 0x70, 0x72, 0x65, 0x63, 0x61, 0x74, 0x65, 0x64, 0x46, 0x69,
	0x65, 0x6c, 0x64, 0x3a, 0x02, 0x18, 0x01, 0x2a, 0x28, 0x0a, 0x0e, 0x44, 0x65, 0x70, 0x72, 0x65,
	0x63, 0x61, 0x74, 0x65, 0x64, 0x45, 0x6e, 0x75, 0x6d, 0x12, 0x12, 0x0a, 0x0a, 0x44, 0x45, 0x50,
	0x52, 0x45, 0x43, 0x41, 0x54, 0x45, 0x44, 0x10, 0x00, 0x1a, 0x02, 0x08, 0x01, 0x1a, 0x02, 0x18,
	0x01, 0x42, 0x43, 0x5a, 0x3e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x67, 0x6f, 0x6c, 0x61,
	0x6e, 0x67, 0x2e, 0x6f, 0x72, 0x67, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f,
	0x63, 0x6d, 0x64, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x63, 0x2d, 0x67, 0x65, 0x6e, 0x2d, 0x67,
	0x6f, 0x2f, 0x74, 0x65, 0x73, 0x74, 0x64, 0x61, 0x74, 0x61, 0x2f, 0x63, 0x6f, 0x6d, 0x6d, 0x65,
	0x6e, 0x74, 0x73, 0xb8, 0x01, 0x01, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_cmd_protoc_gen_go_testdata_comments_deprecated_proto_rawDescOnce sync.Once
	file_cmd_protoc_gen_go_testdata_comments_deprecated_proto_rawDescData = file_cmd_protoc_gen_go_testdata_comments_deprecated_proto_rawDesc
)

func file_cmd_protoc_gen_go_testdata_comments_deprecated_proto_rawDescGZIP() []byte {
	file_cmd_protoc_gen_go_testdata_comments_deprecated_proto_rawDescOnce.Do(func() {
		file_cmd_protoc_gen_go_testdata_comments_deprecated_proto_rawDescData = protoimpl.X.CompressGZIP(file_cmd_protoc_gen_go_testdata_comments_deprecated_proto_rawDescData)
	})
	return file_cmd_protoc_gen_go_testdata_comments_deprecated_proto_rawDescData
}

var file_cmd_protoc_gen_go_testdata_comments_deprecated_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_cmd_protoc_gen_go_testdata_comments_deprecated_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_cmd_protoc_gen_go_testdata_comments_deprecated_proto_goTypes = []interface{}{
	(DeprecatedEnum)(0),       // 0: goproto.protoc.comments.DeprecatedEnum
	(*DeprecatedMessage)(nil), // 1: goproto.protoc.comments.DeprecatedMessage
}
var file_cmd_protoc_gen_go_testdata_comments_deprecated_proto_depIdxs = []int32{
	0, // [0:0] is the sub-list for method output_type
	0, // [0:0] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_cmd_protoc_gen_go_testdata_comments_deprecated_proto_init() }
func file_cmd_protoc_gen_go_testdata_comments_deprecated_proto_init() {
	if File_cmd_protoc_gen_go_testdata_comments_deprecated_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_cmd_protoc_gen_go_testdata_comments_deprecated_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DeprecatedMessage); i {
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
			RawDescriptor: file_cmd_protoc_gen_go_testdata_comments_deprecated_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_cmd_protoc_gen_go_testdata_comments_deprecated_proto_goTypes,
		DependencyIndexes: file_cmd_protoc_gen_go_testdata_comments_deprecated_proto_depIdxs,
		EnumInfos:         file_cmd_protoc_gen_go_testdata_comments_deprecated_proto_enumTypes,
		MessageInfos:      file_cmd_protoc_gen_go_testdata_comments_deprecated_proto_msgTypes,
	}.Build()
	File_cmd_protoc_gen_go_testdata_comments_deprecated_proto = out.File
	file_cmd_protoc_gen_go_testdata_comments_deprecated_proto_rawDesc = nil
	file_cmd_protoc_gen_go_testdata_comments_deprecated_proto_goTypes = nil
	file_cmd_protoc_gen_go_testdata_comments_deprecated_proto_depIdxs = nil
}
