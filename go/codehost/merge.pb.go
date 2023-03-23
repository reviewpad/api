// Copyright (C) 2023 Explore.dev, Unipessoal Lda - All Rights Reserved
// Unauthorized copying of this file, via any medium is strictly prohibited
// Proprietary and confidential

// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.30.0
// 	protoc        v3.21.12
// source: codehost/merge.proto

package codehost

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

type MergeMethod int32

const (
	MergeMethod_MERGE                    MergeMethod = 0 // Creates merge commit; blocks operation if there are conflicts
	MergeMethod_REBASE                   MergeMethod = 1 // Tries to rebase source branch into target branch; blocks operation if there are (intermediate) conflicts
	MergeMethod_SQUASH                   MergeMethod = 2 // Creates single squash commit; blocks operation if there are conflicts
	MergeMethod_FAST_FORWARD             MergeMethod = 3 // Tries to fast-forward target branch to source branch; creates merge commit otherwise if there are no conflicts
	MergeMethod_FAST_FORWARD_ONLY        MergeMethod = 4 // Tries to fast-forward target branch to source branch if up-to-date; blocks operation if outdated
	MergeMethod_REBASE_AND_MERGE         MergeMethod = 5 // Same as REBASE but with a merge commit in the end
	MergeMethod_SQUASH_FAST_FORWARD_ONLY MergeMethod = 6 // Same as SQUASH, except operation is blocked if source branch is outdated
)

// Enum value maps for MergeMethod.
var (
	MergeMethod_name = map[int32]string{
		0: "MERGE",
		1: "REBASE",
		2: "SQUASH",
		3: "FAST_FORWARD",
		4: "FAST_FORWARD_ONLY",
		5: "REBASE_AND_MERGE",
		6: "SQUASH_FAST_FORWARD_ONLY",
	}
	MergeMethod_value = map[string]int32{
		"MERGE":                    0,
		"REBASE":                   1,
		"SQUASH":                   2,
		"FAST_FORWARD":             3,
		"FAST_FORWARD_ONLY":        4,
		"REBASE_AND_MERGE":         5,
		"SQUASH_FAST_FORWARD_ONLY": 6,
	}
)

func (x MergeMethod) Enum() *MergeMethod {
	p := new(MergeMethod)
	*p = x
	return p
}

func (x MergeMethod) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (MergeMethod) Descriptor() protoreflect.EnumDescriptor {
	return file_codehost_merge_proto_enumTypes[0].Descriptor()
}

func (MergeMethod) Type() protoreflect.EnumType {
	return &file_codehost_merge_proto_enumTypes[0]
}

func (x MergeMethod) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use MergeMethod.Descriptor instead.
func (MergeMethod) EnumDescriptor() ([]byte, []int) {
	return file_codehost_merge_proto_rawDescGZIP(), []int{0}
}

type DeleteHeadOption int32

const (
	DeleteHeadOption_PRESERVE DeleteHeadOption = 0
	DeleteHeadOption_DELETE   DeleteHeadOption = 1
)

// Enum value maps for DeleteHeadOption.
var (
	DeleteHeadOption_name = map[int32]string{
		0: "PRESERVE",
		1: "DELETE",
	}
	DeleteHeadOption_value = map[string]int32{
		"PRESERVE": 0,
		"DELETE":   1,
	}
)

func (x DeleteHeadOption) Enum() *DeleteHeadOption {
	p := new(DeleteHeadOption)
	*p = x
	return p
}

func (x DeleteHeadOption) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (DeleteHeadOption) Descriptor() protoreflect.EnumDescriptor {
	return file_codehost_merge_proto_enumTypes[1].Descriptor()
}

func (DeleteHeadOption) Type() protoreflect.EnumType {
	return &file_codehost_merge_proto_enumTypes[1]
}

func (x DeleteHeadOption) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use DeleteHeadOption.Descriptor instead.
func (DeleteHeadOption) EnumDescriptor() ([]byte, []int) {
	return file_codehost_merge_proto_rawDescGZIP(), []int{1}
}

type MergeMethodOptions struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	MergeMethod     MergeMethod `protobuf:"varint,1,opt,name=merge_method,json=mergeMethod,proto3,enum=codehost.MergeMethod" json:"merge_method,omitempty"`
	RequiresMessage bool        `protobuf:"varint,2,opt,name=requires_message,json=requiresMessage,proto3" json:"requires_message,omitempty"`
}

func (x *MergeMethodOptions) Reset() {
	*x = MergeMethodOptions{}
	if protoimpl.UnsafeEnabled {
		mi := &file_codehost_merge_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *MergeMethodOptions) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*MergeMethodOptions) ProtoMessage() {}

func (x *MergeMethodOptions) ProtoReflect() protoreflect.Message {
	mi := &file_codehost_merge_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use MergeMethodOptions.ProtoReflect.Descriptor instead.
func (*MergeMethodOptions) Descriptor() ([]byte, []int) {
	return file_codehost_merge_proto_rawDescGZIP(), []int{0}
}

func (x *MergeMethodOptions) GetMergeMethod() MergeMethod {
	if x != nil {
		return x.MergeMethod
	}
	return MergeMethod_MERGE
}

func (x *MergeMethodOptions) GetRequiresMessage() bool {
	if x != nil {
		return x.RequiresMessage
	}
	return false
}

var File_codehost_merge_proto protoreflect.FileDescriptor

var file_codehost_merge_proto_rawDesc = []byte{
	0x0a, 0x14, 0x63, 0x6f, 0x64, 0x65, 0x68, 0x6f, 0x73, 0x74, 0x2f, 0x6d, 0x65, 0x72, 0x67, 0x65,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x08, 0x63, 0x6f, 0x64, 0x65, 0x68, 0x6f, 0x73, 0x74,
	0x22, 0x79, 0x0a, 0x12, 0x4d, 0x65, 0x72, 0x67, 0x65, 0x4d, 0x65, 0x74, 0x68, 0x6f, 0x64, 0x4f,
	0x70, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x12, 0x38, 0x0a, 0x0c, 0x6d, 0x65, 0x72, 0x67, 0x65, 0x5f,
	0x6d, 0x65, 0x74, 0x68, 0x6f, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x15, 0x2e, 0x63,
	0x6f, 0x64, 0x65, 0x68, 0x6f, 0x73, 0x74, 0x2e, 0x4d, 0x65, 0x72, 0x67, 0x65, 0x4d, 0x65, 0x74,
	0x68, 0x6f, 0x64, 0x52, 0x0b, 0x6d, 0x65, 0x72, 0x67, 0x65, 0x4d, 0x65, 0x74, 0x68, 0x6f, 0x64,
	0x12, 0x29, 0x0a, 0x10, 0x72, 0x65, 0x71, 0x75, 0x69, 0x72, 0x65, 0x73, 0x5f, 0x6d, 0x65, 0x73,
	0x73, 0x61, 0x67, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x08, 0x52, 0x0f, 0x72, 0x65, 0x71, 0x75,
	0x69, 0x72, 0x65, 0x73, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x2a, 0x8d, 0x01, 0x0a, 0x0b,
	0x4d, 0x65, 0x72, 0x67, 0x65, 0x4d, 0x65, 0x74, 0x68, 0x6f, 0x64, 0x12, 0x09, 0x0a, 0x05, 0x4d,
	0x45, 0x52, 0x47, 0x45, 0x10, 0x00, 0x12, 0x0a, 0x0a, 0x06, 0x52, 0x45, 0x42, 0x41, 0x53, 0x45,
	0x10, 0x01, 0x12, 0x0a, 0x0a, 0x06, 0x53, 0x51, 0x55, 0x41, 0x53, 0x48, 0x10, 0x02, 0x12, 0x10,
	0x0a, 0x0c, 0x46, 0x41, 0x53, 0x54, 0x5f, 0x46, 0x4f, 0x52, 0x57, 0x41, 0x52, 0x44, 0x10, 0x03,
	0x12, 0x15, 0x0a, 0x11, 0x46, 0x41, 0x53, 0x54, 0x5f, 0x46, 0x4f, 0x52, 0x57, 0x41, 0x52, 0x44,
	0x5f, 0x4f, 0x4e, 0x4c, 0x59, 0x10, 0x04, 0x12, 0x14, 0x0a, 0x10, 0x52, 0x45, 0x42, 0x41, 0x53,
	0x45, 0x5f, 0x41, 0x4e, 0x44, 0x5f, 0x4d, 0x45, 0x52, 0x47, 0x45, 0x10, 0x05, 0x12, 0x1c, 0x0a,
	0x18, 0x53, 0x51, 0x55, 0x41, 0x53, 0x48, 0x5f, 0x46, 0x41, 0x53, 0x54, 0x5f, 0x46, 0x4f, 0x52,
	0x57, 0x41, 0x52, 0x44, 0x5f, 0x4f, 0x4e, 0x4c, 0x59, 0x10, 0x06, 0x2a, 0x2c, 0x0a, 0x10, 0x44,
	0x65, 0x6c, 0x65, 0x74, 0x65, 0x48, 0x65, 0x61, 0x64, 0x4f, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x12,
	0x0c, 0x0a, 0x08, 0x50, 0x52, 0x45, 0x53, 0x45, 0x52, 0x56, 0x45, 0x10, 0x00, 0x12, 0x0a, 0x0a,
	0x06, 0x44, 0x45, 0x4c, 0x45, 0x54, 0x45, 0x10, 0x01, 0x42, 0x26, 0x5a, 0x24, 0x67, 0x69, 0x74,
	0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x72, 0x65, 0x76, 0x69, 0x65, 0x77, 0x70, 0x61,
	0x64, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x67, 0x6f, 0x2f, 0x63, 0x6f, 0x64, 0x65, 0x68, 0x6f, 0x73,
	0x74, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_codehost_merge_proto_rawDescOnce sync.Once
	file_codehost_merge_proto_rawDescData = file_codehost_merge_proto_rawDesc
)

func file_codehost_merge_proto_rawDescGZIP() []byte {
	file_codehost_merge_proto_rawDescOnce.Do(func() {
		file_codehost_merge_proto_rawDescData = protoimpl.X.CompressGZIP(file_codehost_merge_proto_rawDescData)
	})
	return file_codehost_merge_proto_rawDescData
}

var file_codehost_merge_proto_enumTypes = make([]protoimpl.EnumInfo, 2)
var file_codehost_merge_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_codehost_merge_proto_goTypes = []interface{}{
	(MergeMethod)(0),           // 0: codehost.MergeMethod
	(DeleteHeadOption)(0),      // 1: codehost.DeleteHeadOption
	(*MergeMethodOptions)(nil), // 2: codehost.MergeMethodOptions
}
var file_codehost_merge_proto_depIdxs = []int32{
	0, // 0: codehost.MergeMethodOptions.merge_method:type_name -> codehost.MergeMethod
	1, // [1:1] is the sub-list for method output_type
	1, // [1:1] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_codehost_merge_proto_init() }
func file_codehost_merge_proto_init() {
	if File_codehost_merge_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_codehost_merge_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*MergeMethodOptions); i {
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
			RawDescriptor: file_codehost_merge_proto_rawDesc,
			NumEnums:      2,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_codehost_merge_proto_goTypes,
		DependencyIndexes: file_codehost_merge_proto_depIdxs,
		EnumInfos:         file_codehost_merge_proto_enumTypes,
		MessageInfos:      file_codehost_merge_proto_msgTypes,
	}.Build()
	File_codehost_merge_proto = out.File
	file_codehost_merge_proto_rawDesc = nil
	file_codehost_merge_proto_goTypes = nil
	file_codehost_merge_proto_depIdxs = nil
}
