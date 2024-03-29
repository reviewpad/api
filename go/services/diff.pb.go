// Copyright 2023 Explore.dev Unipessoal Lda. All Rights Reserved.
// Use of this source code is governed by a license that can be
// found in the LICENSE file.

// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.30.0
// 	protoc        v3.21.12
// source: services/diff.proto

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

type GetDefinedSymbolsDiffRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	RepoUri string            `protobuf:"bytes,1,opt,name=repo_uri,json=repoUri,proto3" json:"repo_uri,omitempty"`
	Old     *entities.Symbols `protobuf:"bytes,2,opt,name=old,proto3" json:"old,omitempty"`
	New     *entities.Symbols `protobuf:"bytes,3,opt,name=new,proto3" json:"new,omitempty"`
	GitDiff *entities.GitDiff `protobuf:"bytes,4,opt,name=git_diff,json=gitDiff,proto3" json:"git_diff,omitempty"`
}

func (x *GetDefinedSymbolsDiffRequest) Reset() {
	*x = GetDefinedSymbolsDiffRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_services_diff_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetDefinedSymbolsDiffRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetDefinedSymbolsDiffRequest) ProtoMessage() {}

func (x *GetDefinedSymbolsDiffRequest) ProtoReflect() protoreflect.Message {
	mi := &file_services_diff_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetDefinedSymbolsDiffRequest.ProtoReflect.Descriptor instead.
func (*GetDefinedSymbolsDiffRequest) Descriptor() ([]byte, []int) {
	return file_services_diff_proto_rawDescGZIP(), []int{0}
}

func (x *GetDefinedSymbolsDiffRequest) GetRepoUri() string {
	if x != nil {
		return x.RepoUri
	}
	return ""
}

func (x *GetDefinedSymbolsDiffRequest) GetOld() *entities.Symbols {
	if x != nil {
		return x.Old
	}
	return nil
}

func (x *GetDefinedSymbolsDiffRequest) GetNew() *entities.Symbols {
	if x != nil {
		return x.New
	}
	return nil
}

func (x *GetDefinedSymbolsDiffRequest) GetGitDiff() *entities.GitDiff {
	if x != nil {
		return x.GitDiff
	}
	return nil
}

type GetDefinedSymbolsDiffReply struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Symbols map[string]*entities.SymbolDiff `protobuf:"bytes,1,rep,name=symbols,proto3" json:"symbols,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"` //Map from symbol id to symbol
	Roots   []string                        `protobuf:"bytes,2,rep,name=roots,proto3" json:"roots,omitempty"`                                                                                             //Roots of the symbol tree
}

func (x *GetDefinedSymbolsDiffReply) Reset() {
	*x = GetDefinedSymbolsDiffReply{}
	if protoimpl.UnsafeEnabled {
		mi := &file_services_diff_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetDefinedSymbolsDiffReply) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetDefinedSymbolsDiffReply) ProtoMessage() {}

func (x *GetDefinedSymbolsDiffReply) ProtoReflect() protoreflect.Message {
	mi := &file_services_diff_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetDefinedSymbolsDiffReply.ProtoReflect.Descriptor instead.
func (*GetDefinedSymbolsDiffReply) Descriptor() ([]byte, []int) {
	return file_services_diff_proto_rawDescGZIP(), []int{1}
}

func (x *GetDefinedSymbolsDiffReply) GetSymbols() map[string]*entities.SymbolDiff {
	if x != nil {
		return x.Symbols
	}
	return nil
}

func (x *GetDefinedSymbolsDiffReply) GetRoots() []string {
	if x != nil {
		return x.Roots
	}
	return nil
}

var File_services_diff_proto protoreflect.FileDescriptor

var file_services_diff_proto_rawDesc = []byte{
	0x0a, 0x13, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x2f, 0x64, 0x69, 0x66, 0x66, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x08, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x1a,
	0x17, 0x65, 0x6e, 0x74, 0x69, 0x74, 0x69, 0x65, 0x73, 0x2f, 0x6c, 0x61, 0x6e, 0x67, 0x75, 0x61,
	0x67, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1c, 0x65, 0x6e, 0x74, 0x69, 0x74, 0x69,
	0x65, 0x73, 0x2f, 0x73, 0x65, 0x6d, 0x61, 0x6e, 0x74, 0x69, 0x63, 0x5f, 0x64, 0x69, 0x66, 0x66,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x17, 0x65, 0x6e, 0x74, 0x69, 0x74, 0x69, 0x65, 0x73,
	0x2f, 0x67, 0x69, 0x74, 0x5f, 0x64, 0x69, 0x66, 0x66, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22,
	0xb1, 0x01, 0x0a, 0x1c, 0x47, 0x65, 0x74, 0x44, 0x65, 0x66, 0x69, 0x6e, 0x65, 0x64, 0x53, 0x79,
	0x6d, 0x62, 0x6f, 0x6c, 0x73, 0x44, 0x69, 0x66, 0x66, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x12, 0x19, 0x0a, 0x08, 0x72, 0x65, 0x70, 0x6f, 0x5f, 0x75, 0x72, 0x69, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x07, 0x72, 0x65, 0x70, 0x6f, 0x55, 0x72, 0x69, 0x12, 0x23, 0x0a, 0x03, 0x6f,
	0x6c, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x11, 0x2e, 0x65, 0x6e, 0x74, 0x69, 0x74,
	0x69, 0x65, 0x73, 0x2e, 0x53, 0x79, 0x6d, 0x62, 0x6f, 0x6c, 0x73, 0x52, 0x03, 0x6f, 0x6c, 0x64,
	0x12, 0x23, 0x0a, 0x03, 0x6e, 0x65, 0x77, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x11, 0x2e,
	0x65, 0x6e, 0x74, 0x69, 0x74, 0x69, 0x65, 0x73, 0x2e, 0x53, 0x79, 0x6d, 0x62, 0x6f, 0x6c, 0x73,
	0x52, 0x03, 0x6e, 0x65, 0x77, 0x12, 0x2c, 0x0a, 0x08, 0x67, 0x69, 0x74, 0x5f, 0x64, 0x69, 0x66,
	0x66, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x11, 0x2e, 0x65, 0x6e, 0x74, 0x69, 0x74, 0x69,
	0x65, 0x73, 0x2e, 0x47, 0x69, 0x74, 0x44, 0x69, 0x66, 0x66, 0x52, 0x07, 0x67, 0x69, 0x74, 0x44,
	0x69, 0x66, 0x66, 0x22, 0xd1, 0x01, 0x0a, 0x1a, 0x47, 0x65, 0x74, 0x44, 0x65, 0x66, 0x69, 0x6e,
	0x65, 0x64, 0x53, 0x79, 0x6d, 0x62, 0x6f, 0x6c, 0x73, 0x44, 0x69, 0x66, 0x66, 0x52, 0x65, 0x70,
	0x6c, 0x79, 0x12, 0x4b, 0x0a, 0x07, 0x73, 0x79, 0x6d, 0x62, 0x6f, 0x6c, 0x73, 0x18, 0x01, 0x20,
	0x03, 0x28, 0x0b, 0x32, 0x31, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x2e, 0x47,
	0x65, 0x74, 0x44, 0x65, 0x66, 0x69, 0x6e, 0x65, 0x64, 0x53, 0x79, 0x6d, 0x62, 0x6f, 0x6c, 0x73,
	0x44, 0x69, 0x66, 0x66, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x2e, 0x53, 0x79, 0x6d, 0x62, 0x6f, 0x6c,
	0x73, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x52, 0x07, 0x73, 0x79, 0x6d, 0x62, 0x6f, 0x6c, 0x73, 0x12,
	0x14, 0x0a, 0x05, 0x72, 0x6f, 0x6f, 0x74, 0x73, 0x18, 0x02, 0x20, 0x03, 0x28, 0x09, 0x52, 0x05,
	0x72, 0x6f, 0x6f, 0x74, 0x73, 0x1a, 0x50, 0x0a, 0x0c, 0x53, 0x79, 0x6d, 0x62, 0x6f, 0x6c, 0x73,
	0x45, 0x6e, 0x74, 0x72, 0x79, 0x12, 0x10, 0x0a, 0x03, 0x6b, 0x65, 0x79, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x03, 0x6b, 0x65, 0x79, 0x12, 0x2a, 0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x14, 0x2e, 0x65, 0x6e, 0x74, 0x69, 0x74, 0x69, 0x65,
	0x73, 0x2e, 0x53, 0x79, 0x6d, 0x62, 0x6f, 0x6c, 0x44, 0x69, 0x66, 0x66, 0x52, 0x05, 0x76, 0x61,
	0x6c, 0x75, 0x65, 0x3a, 0x02, 0x38, 0x01, 0x32, 0x6f, 0x0a, 0x04, 0x44, 0x69, 0x66, 0x66, 0x12,
	0x67, 0x0a, 0x15, 0x47, 0x65, 0x74, 0x44, 0x65, 0x66, 0x69, 0x6e, 0x65, 0x64, 0x53, 0x79, 0x6d,
	0x62, 0x6f, 0x6c, 0x73, 0x44, 0x69, 0x66, 0x66, 0x12, 0x26, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69,
	0x63, 0x65, 0x73, 0x2e, 0x47, 0x65, 0x74, 0x44, 0x65, 0x66, 0x69, 0x6e, 0x65, 0x64, 0x53, 0x79,
	0x6d, 0x62, 0x6f, 0x6c, 0x73, 0x44, 0x69, 0x66, 0x66, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x1a, 0x24, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x2e, 0x47, 0x65, 0x74, 0x44,
	0x65, 0x66, 0x69, 0x6e, 0x65, 0x64, 0x53, 0x79, 0x6d, 0x62, 0x6f, 0x6c, 0x73, 0x44, 0x69, 0x66,
	0x66, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x22, 0x00, 0x42, 0x26, 0x5a, 0x24, 0x67, 0x69, 0x74, 0x68,
	0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x72, 0x65, 0x76, 0x69, 0x65, 0x77, 0x70, 0x61, 0x64,
	0x2f, 0x61, 0x70, 0x69, 0x2f, 0x67, 0x6f, 0x2f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73,
	0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_services_diff_proto_rawDescOnce sync.Once
	file_services_diff_proto_rawDescData = file_services_diff_proto_rawDesc
)

func file_services_diff_proto_rawDescGZIP() []byte {
	file_services_diff_proto_rawDescOnce.Do(func() {
		file_services_diff_proto_rawDescData = protoimpl.X.CompressGZIP(file_services_diff_proto_rawDescData)
	})
	return file_services_diff_proto_rawDescData
}

var file_services_diff_proto_msgTypes = make([]protoimpl.MessageInfo, 3)
var file_services_diff_proto_goTypes = []interface{}{
	(*GetDefinedSymbolsDiffRequest)(nil), // 0: services.GetDefinedSymbolsDiffRequest
	(*GetDefinedSymbolsDiffReply)(nil),   // 1: services.GetDefinedSymbolsDiffReply
	nil,                                  // 2: services.GetDefinedSymbolsDiffReply.SymbolsEntry
	(*entities.Symbols)(nil),             // 3: entities.Symbols
	(*entities.GitDiff)(nil),             // 4: entities.GitDiff
	(*entities.SymbolDiff)(nil),          // 5: entities.SymbolDiff
}
var file_services_diff_proto_depIdxs = []int32{
	3, // 0: services.GetDefinedSymbolsDiffRequest.old:type_name -> entities.Symbols
	3, // 1: services.GetDefinedSymbolsDiffRequest.new:type_name -> entities.Symbols
	4, // 2: services.GetDefinedSymbolsDiffRequest.git_diff:type_name -> entities.GitDiff
	2, // 3: services.GetDefinedSymbolsDiffReply.symbols:type_name -> services.GetDefinedSymbolsDiffReply.SymbolsEntry
	5, // 4: services.GetDefinedSymbolsDiffReply.SymbolsEntry.value:type_name -> entities.SymbolDiff
	0, // 5: services.Diff.GetDefinedSymbolsDiff:input_type -> services.GetDefinedSymbolsDiffRequest
	1, // 6: services.Diff.GetDefinedSymbolsDiff:output_type -> services.GetDefinedSymbolsDiffReply
	6, // [6:7] is the sub-list for method output_type
	5, // [5:6] is the sub-list for method input_type
	5, // [5:5] is the sub-list for extension type_name
	5, // [5:5] is the sub-list for extension extendee
	0, // [0:5] is the sub-list for field type_name
}

func init() { file_services_diff_proto_init() }
func file_services_diff_proto_init() {
	if File_services_diff_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_services_diff_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetDefinedSymbolsDiffRequest); i {
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
		file_services_diff_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetDefinedSymbolsDiffReply); i {
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
			RawDescriptor: file_services_diff_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   3,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_services_diff_proto_goTypes,
		DependencyIndexes: file_services_diff_proto_depIdxs,
		MessageInfos:      file_services_diff_proto_msgTypes,
	}.Build()
	File_services_diff_proto = out.File
	file_services_diff_proto_rawDesc = nil
	file_services_diff_proto_goTypes = nil
	file_services_diff_proto_depIdxs = nil
}
