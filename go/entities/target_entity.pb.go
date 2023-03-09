// Copyright 2023 Explore.dev Unipessoal Lda. All Rights Reserved.
// Use of this source code is governed by a license that can be
// found in the LICENSE file.

// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.26.0
// 	protoc        v3.21.4
// source: entities/target_entity.proto

package entities

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

type TargetEntityKind int32

const (
	TargetEntityKind_PULL_REQUEST TargetEntityKind = 0
	TargetEntityKind_ISSUE        TargetEntityKind = 1
)

// Enum value maps for TargetEntityKind.
var (
	TargetEntityKind_name = map[int32]string{
		0: "PULL_REQUEST",
		1: "ISSUE",
	}
	TargetEntityKind_value = map[string]int32{
		"PULL_REQUEST": 0,
		"ISSUE":        1,
	}
)

func (x TargetEntityKind) Enum() *TargetEntityKind {
	p := new(TargetEntityKind)
	*p = x
	return p
}

func (x TargetEntityKind) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (TargetEntityKind) Descriptor() protoreflect.EnumDescriptor {
	return file_entities_target_entity_proto_enumTypes[0].Descriptor()
}

func (TargetEntityKind) Type() protoreflect.EnumType {
	return &file_entities_target_entity_proto_enumTypes[0]
}

func (x TargetEntityKind) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use TargetEntityKind.Descriptor instead.
func (TargetEntityKind) EnumDescriptor() ([]byte, []int) {
	return file_entities_target_entity_proto_rawDescGZIP(), []int{0}
}

type TargetEntity struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Kind   TargetEntityKind `protobuf:"varint,1,opt,name=kind,proto3,enum=entities.TargetEntityKind" json:"kind,omitempty"`
	Number int32            `protobuf:"varint,2,opt,name=number,proto3" json:"number,omitempty"`
	Owner  string           `protobuf:"bytes,3,opt,name=owner,proto3" json:"owner,omitempty"`
	Repo   string           `protobuf:"bytes,4,opt,name=repo,proto3" json:"repo,omitempty"`
}

func (x *TargetEntity) Reset() {
	*x = TargetEntity{}
	if protoimpl.UnsafeEnabled {
		mi := &file_entities_target_entity_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *TargetEntity) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*TargetEntity) ProtoMessage() {}

func (x *TargetEntity) ProtoReflect() protoreflect.Message {
	mi := &file_entities_target_entity_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use TargetEntity.ProtoReflect.Descriptor instead.
func (*TargetEntity) Descriptor() ([]byte, []int) {
	return file_entities_target_entity_proto_rawDescGZIP(), []int{0}
}

func (x *TargetEntity) GetKind() TargetEntityKind {
	if x != nil {
		return x.Kind
	}
	return TargetEntityKind_PULL_REQUEST
}

func (x *TargetEntity) GetNumber() int32 {
	if x != nil {
		return x.Number
	}
	return 0
}

func (x *TargetEntity) GetOwner() string {
	if x != nil {
		return x.Owner
	}
	return ""
}

func (x *TargetEntity) GetRepo() string {
	if x != nil {
		return x.Repo
	}
	return ""
}

var File_entities_target_entity_proto protoreflect.FileDescriptor

var file_entities_target_entity_proto_rawDesc = []byte{
	0x0a, 0x1c, 0x65, 0x6e, 0x74, 0x69, 0x74, 0x69, 0x65, 0x73, 0x2f, 0x74, 0x61, 0x72, 0x67, 0x65,
	0x74, 0x5f, 0x65, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x08,
	0x65, 0x6e, 0x74, 0x69, 0x74, 0x69, 0x65, 0x73, 0x22, 0x80, 0x01, 0x0a, 0x0c, 0x54, 0x61, 0x72,
	0x67, 0x65, 0x74, 0x45, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x12, 0x2e, 0x0a, 0x04, 0x6b, 0x69, 0x6e,
	0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x1a, 0x2e, 0x65, 0x6e, 0x74, 0x69, 0x74, 0x69,
	0x65, 0x73, 0x2e, 0x54, 0x61, 0x72, 0x67, 0x65, 0x74, 0x45, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x4b,
	0x69, 0x6e, 0x64, 0x52, 0x04, 0x6b, 0x69, 0x6e, 0x64, 0x12, 0x16, 0x0a, 0x06, 0x6e, 0x75, 0x6d,
	0x62, 0x65, 0x72, 0x18, 0x02, 0x20, 0x01, 0x28, 0x05, 0x52, 0x06, 0x6e, 0x75, 0x6d, 0x62, 0x65,
	0x72, 0x12, 0x14, 0x0a, 0x05, 0x6f, 0x77, 0x6e, 0x65, 0x72, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x05, 0x6f, 0x77, 0x6e, 0x65, 0x72, 0x12, 0x12, 0x0a, 0x04, 0x72, 0x65, 0x70, 0x6f, 0x18,
	0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x72, 0x65, 0x70, 0x6f, 0x2a, 0x2f, 0x0a, 0x10, 0x54,
	0x61, 0x72, 0x67, 0x65, 0x74, 0x45, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x4b, 0x69, 0x6e, 0x64, 0x12,
	0x10, 0x0a, 0x0c, 0x50, 0x55, 0x4c, 0x4c, 0x5f, 0x52, 0x45, 0x51, 0x55, 0x45, 0x53, 0x54, 0x10,
	0x00, 0x12, 0x09, 0x0a, 0x05, 0x49, 0x53, 0x53, 0x55, 0x45, 0x10, 0x01, 0x42, 0x26, 0x5a, 0x24,
	0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x72, 0x65, 0x76, 0x69, 0x65,
	0x77, 0x70, 0x61, 0x64, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x67, 0x6f, 0x2f, 0x65, 0x6e, 0x74, 0x69,
	0x74, 0x69, 0x65, 0x73, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_entities_target_entity_proto_rawDescOnce sync.Once
	file_entities_target_entity_proto_rawDescData = file_entities_target_entity_proto_rawDesc
)

func file_entities_target_entity_proto_rawDescGZIP() []byte {
	file_entities_target_entity_proto_rawDescOnce.Do(func() {
		file_entities_target_entity_proto_rawDescData = protoimpl.X.CompressGZIP(file_entities_target_entity_proto_rawDescData)
	})
	return file_entities_target_entity_proto_rawDescData
}

var file_entities_target_entity_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_entities_target_entity_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_entities_target_entity_proto_goTypes = []interface{}{
	(TargetEntityKind)(0), // 0: entities.TargetEntityKind
	(*TargetEntity)(nil),  // 1: entities.TargetEntity
}
var file_entities_target_entity_proto_depIdxs = []int32{
	0, // 0: entities.TargetEntity.kind:type_name -> entities.TargetEntityKind
	1, // [1:1] is the sub-list for method output_type
	1, // [1:1] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_entities_target_entity_proto_init() }
func file_entities_target_entity_proto_init() {
	if File_entities_target_entity_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_entities_target_entity_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*TargetEntity); i {
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
			RawDescriptor: file_entities_target_entity_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_entities_target_entity_proto_goTypes,
		DependencyIndexes: file_entities_target_entity_proto_depIdxs,
		EnumInfos:         file_entities_target_entity_proto_enumTypes,
		MessageInfos:      file_entities_target_entity_proto_msgTypes,
	}.Build()
	File_entities_target_entity_proto = out.File
	file_entities_target_entity_proto_rawDesc = nil
	file_entities_target_entity_proto_goTypes = nil
	file_entities_target_entity_proto_depIdxs = nil
}
