// Copyright 2023 Explore.dev Unipessoal Lda. All Rights Reserved.
// Use of this source code is governed by a license that can be
// found in the LICENSE file.

// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.26.0
// 	protoc        v3.21.4
// source: entities/file.proto

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

type FileInfoDiff struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	OldInfo         *FileInfo `protobuf:"bytes,1,opt,name=old_info,json=oldInfo,proto3" json:"old_info,omitempty"`
	NewInfo         *FileInfo `protobuf:"bytes,2,opt,name=new_info,json=newInfo,proto3" json:"new_info,omitempty"`
	NumAddedLines   int32     `protobuf:"varint,3,opt,name=num_added_lines,json=numAddedLines,proto3" json:"num_added_lines,omitempty"`
	NumRemovedLines int32     `protobuf:"varint,4,opt,name=num_removed_lines,json=numRemovedLines,proto3" json:"num_removed_lines,omitempty"`
	IsBinary        bool      `protobuf:"varint,5,opt,name=is_binary,json=isBinary,proto3" json:"is_binary,omitempty"`
}

func (x *FileInfoDiff) Reset() {
	*x = FileInfoDiff{}
	if protoimpl.UnsafeEnabled {
		mi := &file_entities_file_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *FileInfoDiff) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*FileInfoDiff) ProtoMessage() {}

func (x *FileInfoDiff) ProtoReflect() protoreflect.Message {
	mi := &file_entities_file_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use FileInfoDiff.ProtoReflect.Descriptor instead.
func (*FileInfoDiff) Descriptor() ([]byte, []int) {
	return file_entities_file_proto_rawDescGZIP(), []int{0}
}

func (x *FileInfoDiff) GetOldInfo() *FileInfo {
	if x != nil {
		return x.OldInfo
	}
	return nil
}

func (x *FileInfoDiff) GetNewInfo() *FileInfo {
	if x != nil {
		return x.NewInfo
	}
	return nil
}

func (x *FileInfoDiff) GetNumAddedLines() int32 {
	if x != nil {
		return x.NumAddedLines
	}
	return 0
}

func (x *FileInfoDiff) GetNumRemovedLines() int32 {
	if x != nil {
		return x.NumRemovedLines
	}
	return 0
}

func (x *FileInfoDiff) GetIsBinary() bool {
	if x != nil {
		return x.IsBinary
	}
	return false
}

type FileInfo struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Path      string `protobuf:"bytes,1,opt,name=path,proto3" json:"path,omitempty"`
	Extension string `protobuf:"bytes,2,opt,name=extension,proto3" json:"extension,omitempty"`
	BlobId    string `protobuf:"bytes,3,opt,name=blob_id,json=blobId,proto3" json:"blob_id,omitempty"`
	NumLines  int32  `protobuf:"varint,4,opt,name=num_lines,json=numLines,proto3" json:"num_lines,omitempty"`
}

func (x *FileInfo) Reset() {
	*x = FileInfo{}
	if protoimpl.UnsafeEnabled {
		mi := &file_entities_file_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *FileInfo) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*FileInfo) ProtoMessage() {}

func (x *FileInfo) ProtoReflect() protoreflect.Message {
	mi := &file_entities_file_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use FileInfo.ProtoReflect.Descriptor instead.
func (*FileInfo) Descriptor() ([]byte, []int) {
	return file_entities_file_proto_rawDescGZIP(), []int{1}
}

func (x *FileInfo) GetPath() string {
	if x != nil {
		return x.Path
	}
	return ""
}

func (x *FileInfo) GetExtension() string {
	if x != nil {
		return x.Extension
	}
	return ""
}

func (x *FileInfo) GetBlobId() string {
	if x != nil {
		return x.BlobId
	}
	return ""
}

func (x *FileInfo) GetNumLines() int32 {
	if x != nil {
		return x.NumLines
	}
	return 0
}

var File_entities_file_proto protoreflect.FileDescriptor

var file_entities_file_proto_rawDesc = []byte{
	0x0a, 0x13, 0x65, 0x6e, 0x74, 0x69, 0x74, 0x69, 0x65, 0x73, 0x2f, 0x66, 0x69, 0x6c, 0x65, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x08, 0x65, 0x6e, 0x74, 0x69, 0x74, 0x69, 0x65, 0x73, 0x22,
	0xdd, 0x01, 0x0a, 0x0c, 0x46, 0x69, 0x6c, 0x65, 0x49, 0x6e, 0x66, 0x6f, 0x44, 0x69, 0x66, 0x66,
	0x12, 0x2d, 0x0a, 0x08, 0x6f, 0x6c, 0x64, 0x5f, 0x69, 0x6e, 0x66, 0x6f, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x0b, 0x32, 0x12, 0x2e, 0x65, 0x6e, 0x74, 0x69, 0x74, 0x69, 0x65, 0x73, 0x2e, 0x46, 0x69,
	0x6c, 0x65, 0x49, 0x6e, 0x66, 0x6f, 0x52, 0x07, 0x6f, 0x6c, 0x64, 0x49, 0x6e, 0x66, 0x6f, 0x12,
	0x2d, 0x0a, 0x08, 0x6e, 0x65, 0x77, 0x5f, 0x69, 0x6e, 0x66, 0x6f, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x0b, 0x32, 0x12, 0x2e, 0x65, 0x6e, 0x74, 0x69, 0x74, 0x69, 0x65, 0x73, 0x2e, 0x46, 0x69, 0x6c,
	0x65, 0x49, 0x6e, 0x66, 0x6f, 0x52, 0x07, 0x6e, 0x65, 0x77, 0x49, 0x6e, 0x66, 0x6f, 0x12, 0x26,
	0x0a, 0x0f, 0x6e, 0x75, 0x6d, 0x5f, 0x61, 0x64, 0x64, 0x65, 0x64, 0x5f, 0x6c, 0x69, 0x6e, 0x65,
	0x73, 0x18, 0x03, 0x20, 0x01, 0x28, 0x05, 0x52, 0x0d, 0x6e, 0x75, 0x6d, 0x41, 0x64, 0x64, 0x65,
	0x64, 0x4c, 0x69, 0x6e, 0x65, 0x73, 0x12, 0x2a, 0x0a, 0x11, 0x6e, 0x75, 0x6d, 0x5f, 0x72, 0x65,
	0x6d, 0x6f, 0x76, 0x65, 0x64, 0x5f, 0x6c, 0x69, 0x6e, 0x65, 0x73, 0x18, 0x04, 0x20, 0x01, 0x28,
	0x05, 0x52, 0x0f, 0x6e, 0x75, 0x6d, 0x52, 0x65, 0x6d, 0x6f, 0x76, 0x65, 0x64, 0x4c, 0x69, 0x6e,
	0x65, 0x73, 0x12, 0x1b, 0x0a, 0x09, 0x69, 0x73, 0x5f, 0x62, 0x69, 0x6e, 0x61, 0x72, 0x79, 0x18,
	0x05, 0x20, 0x01, 0x28, 0x08, 0x52, 0x08, 0x69, 0x73, 0x42, 0x69, 0x6e, 0x61, 0x72, 0x79, 0x22,
	0x72, 0x0a, 0x08, 0x46, 0x69, 0x6c, 0x65, 0x49, 0x6e, 0x66, 0x6f, 0x12, 0x12, 0x0a, 0x04, 0x70,
	0x61, 0x74, 0x68, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x70, 0x61, 0x74, 0x68, 0x12,
	0x1c, 0x0a, 0x09, 0x65, 0x78, 0x74, 0x65, 0x6e, 0x73, 0x69, 0x6f, 0x6e, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x09, 0x65, 0x78, 0x74, 0x65, 0x6e, 0x73, 0x69, 0x6f, 0x6e, 0x12, 0x17, 0x0a,
	0x07, 0x62, 0x6c, 0x6f, 0x62, 0x5f, 0x69, 0x64, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06,
	0x62, 0x6c, 0x6f, 0x62, 0x49, 0x64, 0x12, 0x1b, 0x0a, 0x09, 0x6e, 0x75, 0x6d, 0x5f, 0x6c, 0x69,
	0x6e, 0x65, 0x73, 0x18, 0x04, 0x20, 0x01, 0x28, 0x05, 0x52, 0x08, 0x6e, 0x75, 0x6d, 0x4c, 0x69,
	0x6e, 0x65, 0x73, 0x42, 0x26, 0x5a, 0x24, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f,
	0x6d, 0x2f, 0x72, 0x65, 0x76, 0x69, 0x65, 0x77, 0x70, 0x61, 0x64, 0x2f, 0x61, 0x70, 0x69, 0x2f,
	0x67, 0x6f, 0x2f, 0x65, 0x6e, 0x74, 0x69, 0x74, 0x69, 0x65, 0x73, 0x62, 0x06, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x33,
}

var (
	file_entities_file_proto_rawDescOnce sync.Once
	file_entities_file_proto_rawDescData = file_entities_file_proto_rawDesc
)

func file_entities_file_proto_rawDescGZIP() []byte {
	file_entities_file_proto_rawDescOnce.Do(func() {
		file_entities_file_proto_rawDescData = protoimpl.X.CompressGZIP(file_entities_file_proto_rawDescData)
	})
	return file_entities_file_proto_rawDescData
}

var file_entities_file_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_entities_file_proto_goTypes = []interface{}{
	(*FileInfoDiff)(nil), // 0: entities.FileInfoDiff
	(*FileInfo)(nil),     // 1: entities.FileInfo
}
var file_entities_file_proto_depIdxs = []int32{
	1, // 0: entities.FileInfoDiff.old_info:type_name -> entities.FileInfo
	1, // 1: entities.FileInfoDiff.new_info:type_name -> entities.FileInfo
	2, // [2:2] is the sub-list for method output_type
	2, // [2:2] is the sub-list for method input_type
	2, // [2:2] is the sub-list for extension type_name
	2, // [2:2] is the sub-list for extension extendee
	0, // [0:2] is the sub-list for field type_name
}

func init() { file_entities_file_proto_init() }
func file_entities_file_proto_init() {
	if File_entities_file_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_entities_file_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*FileInfoDiff); i {
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
		file_entities_file_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*FileInfo); i {
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
			RawDescriptor: file_entities_file_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_entities_file_proto_goTypes,
		DependencyIndexes: file_entities_file_proto_depIdxs,
		MessageInfos:      file_entities_file_proto_msgTypes,
	}.Build()
	File_entities_file_proto = out.File
	file_entities_file_proto_rawDesc = nil
	file_entities_file_proto_goTypes = nil
	file_entities_file_proto_depIdxs = nil
}