// Copyright 2022 Explore.dev Unipessoal Lda. All Rights Reserved.
// Use of this source code is governed by a license that can be
// found in the LICENSE file.

// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.30.0
// 	protoc        v3.21.12
// source: services/semantic.proto

package services

import (
	entities "github.com/reviewpad/api/go/entities"
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

type GetSymbolsRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Uri      string                    `protobuf:"bytes,1,opt,name=uri,proto3" json:"uri,omitempty"`
	CommitId string                    `protobuf:"bytes,2,opt,name=commit_id,json=commitId,proto3" json:"commit_id,omitempty"`
	BlobId   string                    `protobuf:"bytes,3,opt,name=blob_id,json=blobId,proto3" json:"blob_id,omitempty"`
	Filepath string                    `protobuf:"bytes,4,opt,name=filepath,proto3" json:"filepath,omitempty"`
	Blob     []byte                    `protobuf:"bytes,5,opt,name=blob,proto3" json:"blob,omitempty"`
	Diff     *entities.ResolveFileDiff `protobuf:"bytes,6,opt,name=diff,proto3" json:"diff,omitempty"`
}

func (x *GetSymbolsRequest) Reset() {
	*x = GetSymbolsRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_services_semantic_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetSymbolsRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetSymbolsRequest) ProtoMessage() {}

func (x *GetSymbolsRequest) ProtoReflect() protoreflect.Message {
	mi := &file_services_semantic_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetSymbolsRequest.ProtoReflect.Descriptor instead.
func (*GetSymbolsRequest) Descriptor() ([]byte, []int) {
	return file_services_semantic_proto_rawDescGZIP(), []int{0}
}

func (x *GetSymbolsRequest) GetUri() string {
	if x != nil {
		return x.Uri
	}
	return ""
}

func (x *GetSymbolsRequest) GetCommitId() string {
	if x != nil {
		return x.CommitId
	}
	return ""
}

func (x *GetSymbolsRequest) GetBlobId() string {
	if x != nil {
		return x.BlobId
	}
	return ""
}

func (x *GetSymbolsRequest) GetFilepath() string {
	if x != nil {
		return x.Filepath
	}
	return ""
}

func (x *GetSymbolsRequest) GetBlob() []byte {
	if x != nil {
		return x.Blob
	}
	return nil
}

func (x *GetSymbolsRequest) GetDiff() *entities.ResolveFileDiff {
	if x != nil {
		return x.Diff
	}
	return nil
}

type GetSymbolsReply struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Symbols *entities.Symbols `protobuf:"bytes,1,opt,name=symbols,proto3" json:"symbols,omitempty"`
}

func (x *GetSymbolsReply) Reset() {
	*x = GetSymbolsReply{}
	if protoimpl.UnsafeEnabled {
		mi := &file_services_semantic_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetSymbolsReply) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetSymbolsReply) ProtoMessage() {}

func (x *GetSymbolsReply) ProtoReflect() protoreflect.Message {
	mi := &file_services_semantic_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetSymbolsReply.ProtoReflect.Descriptor instead.
func (*GetSymbolsReply) Descriptor() ([]byte, []int) {
	return file_services_semantic_proto_rawDescGZIP(), []int{1}
}

func (x *GetSymbolsReply) GetSymbols() *entities.Symbols {
	if x != nil {
		return x.Symbols
	}
	return nil
}

var File_services_semantic_proto protoreflect.FileDescriptor

var file_services_semantic_proto_rawDesc = []byte{
	0x0a, 0x17, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x2f, 0x73, 0x65, 0x6d, 0x61, 0x6e,
	0x74, 0x69, 0x63, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x08, 0x73, 0x65, 0x72, 0x76, 0x69,
	0x63, 0x65, 0x73, 0x1a, 0x17, 0x65, 0x6e, 0x74, 0x69, 0x74, 0x69, 0x65, 0x73, 0x2f, 0x6c, 0x61,
	0x6e, 0x67, 0x75, 0x61, 0x67, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x13, 0x65, 0x6e,
	0x74, 0x69, 0x74, 0x69, 0x65, 0x73, 0x2f, 0x64, 0x69, 0x66, 0x66, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x22, 0xba, 0x01, 0x0a, 0x11, 0x47, 0x65, 0x74, 0x53, 0x79, 0x6d, 0x62, 0x6f, 0x6c, 0x73,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x10, 0x0a, 0x03, 0x75, 0x72, 0x69, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x75, 0x72, 0x69, 0x12, 0x1b, 0x0a, 0x09, 0x63, 0x6f, 0x6d,
	0x6d, 0x69, 0x74, 0x5f, 0x69, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x63, 0x6f,
	0x6d, 0x6d, 0x69, 0x74, 0x49, 0x64, 0x12, 0x17, 0x0a, 0x07, 0x62, 0x6c, 0x6f, 0x62, 0x5f, 0x69,
	0x64, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x62, 0x6c, 0x6f, 0x62, 0x49, 0x64, 0x12,
	0x1a, 0x0a, 0x08, 0x66, 0x69, 0x6c, 0x65, 0x70, 0x61, 0x74, 0x68, 0x18, 0x04, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x08, 0x66, 0x69, 0x6c, 0x65, 0x70, 0x61, 0x74, 0x68, 0x12, 0x12, 0x0a, 0x04, 0x62,
	0x6c, 0x6f, 0x62, 0x18, 0x05, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x04, 0x62, 0x6c, 0x6f, 0x62, 0x12,
	0x2d, 0x0a, 0x04, 0x64, 0x69, 0x66, 0x66, 0x18, 0x06, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x19, 0x2e,
	0x65, 0x6e, 0x74, 0x69, 0x74, 0x69, 0x65, 0x73, 0x2e, 0x52, 0x65, 0x73, 0x6f, 0x6c, 0x76, 0x65,
	0x46, 0x69, 0x6c, 0x65, 0x44, 0x69, 0x66, 0x66, 0x52, 0x04, 0x64, 0x69, 0x66, 0x66, 0x22, 0x3e,
	0x0a, 0x0f, 0x47, 0x65, 0x74, 0x53, 0x79, 0x6d, 0x62, 0x6f, 0x6c, 0x73, 0x52, 0x65, 0x70, 0x6c,
	0x79, 0x12, 0x2b, 0x0a, 0x07, 0x73, 0x79, 0x6d, 0x62, 0x6f, 0x6c, 0x73, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x0b, 0x32, 0x11, 0x2e, 0x65, 0x6e, 0x74, 0x69, 0x74, 0x69, 0x65, 0x73, 0x2e, 0x53, 0x79,
	0x6d, 0x62, 0x6f, 0x6c, 0x73, 0x52, 0x07, 0x73, 0x79, 0x6d, 0x62, 0x6f, 0x6c, 0x73, 0x32, 0x52,
	0x0a, 0x08, 0x53, 0x65, 0x6d, 0x61, 0x6e, 0x74, 0x69, 0x63, 0x12, 0x46, 0x0a, 0x0a, 0x47, 0x65,
	0x74, 0x53, 0x79, 0x6d, 0x62, 0x6f, 0x6c, 0x73, 0x12, 0x1b, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69,
	0x63, 0x65, 0x73, 0x2e, 0x47, 0x65, 0x74, 0x53, 0x79, 0x6d, 0x62, 0x6f, 0x6c, 0x73, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x19, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73,
	0x2e, 0x47, 0x65, 0x74, 0x53, 0x79, 0x6d, 0x62, 0x6f, 0x6c, 0x73, 0x52, 0x65, 0x70, 0x6c, 0x79,
	0x22, 0x00, 0x42, 0x26, 0x5a, 0x24, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d,
	0x2f, 0x72, 0x65, 0x76, 0x69, 0x65, 0x77, 0x70, 0x61, 0x64, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x67,
	0x6f, 0x2f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x33,
}

var (
	file_services_semantic_proto_rawDescOnce sync.Once
	file_services_semantic_proto_rawDescData = file_services_semantic_proto_rawDesc
)

func file_services_semantic_proto_rawDescGZIP() []byte {
	file_services_semantic_proto_rawDescOnce.Do(func() {
		file_services_semantic_proto_rawDescData = protoimpl.X.CompressGZIP(file_services_semantic_proto_rawDescData)
	})
	return file_services_semantic_proto_rawDescData
}

var file_services_semantic_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_services_semantic_proto_goTypes = []interface{}{
	(*GetSymbolsRequest)(nil),        // 0: services.GetSymbolsRequest
	(*GetSymbolsReply)(nil),          // 1: services.GetSymbolsReply
	(*entities.ResolveFileDiff)(nil), // 2: entities.ResolveFileDiff
	(*entities.Symbols)(nil),         // 3: entities.Symbols
}
var file_services_semantic_proto_depIdxs = []int32{
	2, // 0: services.GetSymbolsRequest.diff:type_name -> entities.ResolveFileDiff
	3, // 1: services.GetSymbolsReply.symbols:type_name -> entities.Symbols
	0, // 2: services.Semantic.GetSymbols:input_type -> services.GetSymbolsRequest
	1, // 3: services.Semantic.GetSymbols:output_type -> services.GetSymbolsReply
	3, // [3:4] is the sub-list for method output_type
	2, // [2:3] is the sub-list for method input_type
	2, // [2:2] is the sub-list for extension type_name
	2, // [2:2] is the sub-list for extension extendee
	0, // [0:2] is the sub-list for field type_name
}

func init() { file_services_semantic_proto_init() }
func file_services_semantic_proto_init() {
	if File_services_semantic_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_services_semantic_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetSymbolsRequest); i {
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
		file_services_semantic_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetSymbolsReply); i {
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
			RawDescriptor: file_services_semantic_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_services_semantic_proto_goTypes,
		DependencyIndexes: file_services_semantic_proto_depIdxs,
		MessageInfos:      file_services_semantic_proto_msgTypes,
	}.Build()
	File_services_semantic_proto = out.File
	file_services_semantic_proto_rawDesc = nil
	file_services_semantic_proto_goTypes = nil
	file_services_semantic_proto_depIdxs = nil
}
