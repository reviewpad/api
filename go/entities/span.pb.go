// Copyright 2022 Explore.dev Unipessoal Lda. All Rights Reserved.
// Use of this source code is governed by a license that can be
// found in the LICENSE file.

// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.29.0
// 	protoc        v3.21.12
// source: entities/span.proto

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

type Span struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Start *Location `protobuf:"bytes,1,opt,name=start,proto3" json:"start,omitempty"`
	End   *Location `protobuf:"bytes,2,opt,name=end,proto3" json:"end,omitempty"`
}

func (x *Span) Reset() {
	*x = Span{}
	if protoimpl.UnsafeEnabled {
		mi := &file_entities_span_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Span) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Span) ProtoMessage() {}

func (x *Span) ProtoReflect() protoreflect.Message {
	mi := &file_entities_span_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Span.ProtoReflect.Descriptor instead.
func (*Span) Descriptor() ([]byte, []int) {
	return file_entities_span_proto_rawDescGZIP(), []int{0}
}

func (x *Span) GetStart() *Location {
	if x != nil {
		return x.Start
	}
	return nil
}

func (x *Span) GetEnd() *Location {
	if x != nil {
		return x.End
	}
	return nil
}

type Location struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Line   int32 `protobuf:"varint,1,opt,name=line,proto3" json:"line,omitempty"`
	Offset int32 `protobuf:"varint,2,opt,name=offset,proto3" json:"offset,omitempty"`
}

func (x *Location) Reset() {
	*x = Location{}
	if protoimpl.UnsafeEnabled {
		mi := &file_entities_span_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Location) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Location) ProtoMessage() {}

func (x *Location) ProtoReflect() protoreflect.Message {
	mi := &file_entities_span_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Location.ProtoReflect.Descriptor instead.
func (*Location) Descriptor() ([]byte, []int) {
	return file_entities_span_proto_rawDescGZIP(), []int{1}
}

func (x *Location) GetLine() int32 {
	if x != nil {
		return x.Line
	}
	return 0
}

func (x *Location) GetOffset() int32 {
	if x != nil {
		return x.Offset
	}
	return 0
}

type WordSpan struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	StartColumn int32 `protobuf:"varint,1,opt,name=start_column,json=startColumn,proto3" json:"start_column,omitempty"`
	EndColumn   int32 `protobuf:"varint,2,opt,name=end_column,json=endColumn,proto3" json:"end_column,omitempty"`
}

func (x *WordSpan) Reset() {
	*x = WordSpan{}
	if protoimpl.UnsafeEnabled {
		mi := &file_entities_span_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *WordSpan) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*WordSpan) ProtoMessage() {}

func (x *WordSpan) ProtoReflect() protoreflect.Message {
	mi := &file_entities_span_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use WordSpan.ProtoReflect.Descriptor instead.
func (*WordSpan) Descriptor() ([]byte, []int) {
	return file_entities_span_proto_rawDescGZIP(), []int{2}
}

func (x *WordSpan) GetStartColumn() int32 {
	if x != nil {
		return x.StartColumn
	}
	return 0
}

func (x *WordSpan) GetEndColumn() int32 {
	if x != nil {
		return x.EndColumn
	}
	return 0
}

var File_entities_span_proto protoreflect.FileDescriptor

var file_entities_span_proto_rawDesc = []byte{
	0x0a, 0x13, 0x65, 0x6e, 0x74, 0x69, 0x74, 0x69, 0x65, 0x73, 0x2f, 0x73, 0x70, 0x61, 0x6e, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x08, 0x65, 0x6e, 0x74, 0x69, 0x74, 0x69, 0x65, 0x73, 0x22,
	0x56, 0x0a, 0x04, 0x53, 0x70, 0x61, 0x6e, 0x12, 0x28, 0x0a, 0x05, 0x73, 0x74, 0x61, 0x72, 0x74,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x12, 0x2e, 0x65, 0x6e, 0x74, 0x69, 0x74, 0x69, 0x65,
	0x73, 0x2e, 0x4c, 0x6f, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x05, 0x73, 0x74, 0x61, 0x72,
	0x74, 0x12, 0x24, 0x0a, 0x03, 0x65, 0x6e, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x12,
	0x2e, 0x65, 0x6e, 0x74, 0x69, 0x74, 0x69, 0x65, 0x73, 0x2e, 0x4c, 0x6f, 0x63, 0x61, 0x74, 0x69,
	0x6f, 0x6e, 0x52, 0x03, 0x65, 0x6e, 0x64, 0x22, 0x36, 0x0a, 0x08, 0x4c, 0x6f, 0x63, 0x61, 0x74,
	0x69, 0x6f, 0x6e, 0x12, 0x12, 0x0a, 0x04, 0x6c, 0x69, 0x6e, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x05, 0x52, 0x04, 0x6c, 0x69, 0x6e, 0x65, 0x12, 0x16, 0x0a, 0x06, 0x6f, 0x66, 0x66, 0x73, 0x65,
	0x74, 0x18, 0x02, 0x20, 0x01, 0x28, 0x05, 0x52, 0x06, 0x6f, 0x66, 0x66, 0x73, 0x65, 0x74, 0x22,
	0x4c, 0x0a, 0x08, 0x57, 0x6f, 0x72, 0x64, 0x53, 0x70, 0x61, 0x6e, 0x12, 0x21, 0x0a, 0x0c, 0x73,
	0x74, 0x61, 0x72, 0x74, 0x5f, 0x63, 0x6f, 0x6c, 0x75, 0x6d, 0x6e, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x05, 0x52, 0x0b, 0x73, 0x74, 0x61, 0x72, 0x74, 0x43, 0x6f, 0x6c, 0x75, 0x6d, 0x6e, 0x12, 0x1d,
	0x0a, 0x0a, 0x65, 0x6e, 0x64, 0x5f, 0x63, 0x6f, 0x6c, 0x75, 0x6d, 0x6e, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x05, 0x52, 0x09, 0x65, 0x6e, 0x64, 0x43, 0x6f, 0x6c, 0x75, 0x6d, 0x6e, 0x42, 0x26, 0x5a,
	0x24, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x72, 0x65, 0x76, 0x69,
	0x65, 0x77, 0x70, 0x61, 0x64, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x67, 0x6f, 0x2f, 0x65, 0x6e, 0x74,
	0x69, 0x74, 0x69, 0x65, 0x73, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_entities_span_proto_rawDescOnce sync.Once
	file_entities_span_proto_rawDescData = file_entities_span_proto_rawDesc
)

func file_entities_span_proto_rawDescGZIP() []byte {
	file_entities_span_proto_rawDescOnce.Do(func() {
		file_entities_span_proto_rawDescData = protoimpl.X.CompressGZIP(file_entities_span_proto_rawDescData)
	})
	return file_entities_span_proto_rawDescData
}

var file_entities_span_proto_msgTypes = make([]protoimpl.MessageInfo, 3)
var file_entities_span_proto_goTypes = []interface{}{
	(*Span)(nil),     // 0: entities.Span
	(*Location)(nil), // 1: entities.Location
	(*WordSpan)(nil), // 2: entities.WordSpan
}
var file_entities_span_proto_depIdxs = []int32{
	1, // 0: entities.Span.start:type_name -> entities.Location
	1, // 1: entities.Span.end:type_name -> entities.Location
	2, // [2:2] is the sub-list for method output_type
	2, // [2:2] is the sub-list for method input_type
	2, // [2:2] is the sub-list for extension type_name
	2, // [2:2] is the sub-list for extension extendee
	0, // [0:2] is the sub-list for field type_name
}

func init() { file_entities_span_proto_init() }
func file_entities_span_proto_init() {
	if File_entities_span_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_entities_span_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Span); i {
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
		file_entities_span_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Location); i {
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
		file_entities_span_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*WordSpan); i {
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
			RawDescriptor: file_entities_span_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   3,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_entities_span_proto_goTypes,
		DependencyIndexes: file_entities_span_proto_depIdxs,
		MessageInfos:      file_entities_span_proto_msgTypes,
	}.Build()
	File_entities_span_proto = out.File
	file_entities_span_proto_rawDesc = nil
	file_entities_span_proto_goTypes = nil
	file_entities_span_proto_depIdxs = nil
}
