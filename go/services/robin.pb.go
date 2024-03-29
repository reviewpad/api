// Copyright 2023 Explore.dev Unipessoal Lda. All Rights Reserved.
// Use of this source code is governed by a license that can be
// found in the LICENSE file.

// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.30.0
// 	protoc        v3.21.12
// source: services/robin.proto

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

type ChatInCodeHostRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Token  string                 `protobuf:"bytes,1,opt,name=token,proto3" json:"token,omitempty"`
	Target *entities.TargetEntity `protobuf:"bytes,2,opt,name=target,proto3" json:"target,omitempty"`
}

func (x *ChatInCodeHostRequest) Reset() {
	*x = ChatInCodeHostRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_services_robin_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ChatInCodeHostRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ChatInCodeHostRequest) ProtoMessage() {}

func (x *ChatInCodeHostRequest) ProtoReflect() protoreflect.Message {
	mi := &file_services_robin_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ChatInCodeHostRequest.ProtoReflect.Descriptor instead.
func (*ChatInCodeHostRequest) Descriptor() ([]byte, []int) {
	return file_services_robin_proto_rawDescGZIP(), []int{0}
}

func (x *ChatInCodeHostRequest) GetToken() string {
	if x != nil {
		return x.Token
	}
	return ""
}

func (x *ChatInCodeHostRequest) GetTarget() *entities.TargetEntity {
	if x != nil {
		return x.Target
	}
	return nil
}

type ChatInCodeHostReply struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Reply string `protobuf:"bytes,1,opt,name=reply,proto3" json:"reply,omitempty"`
}

func (x *ChatInCodeHostReply) Reset() {
	*x = ChatInCodeHostReply{}
	if protoimpl.UnsafeEnabled {
		mi := &file_services_robin_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ChatInCodeHostReply) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ChatInCodeHostReply) ProtoMessage() {}

func (x *ChatInCodeHostReply) ProtoReflect() protoreflect.Message {
	mi := &file_services_robin_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ChatInCodeHostReply.ProtoReflect.Descriptor instead.
func (*ChatInCodeHostReply) Descriptor() ([]byte, []int) {
	return file_services_robin_proto_rawDescGZIP(), []int{1}
}

func (x *ChatInCodeHostReply) GetReply() string {
	if x != nil {
		return x.Reply
	}
	return ""
}

type ExplainRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Token     string                 `protobuf:"bytes,1,opt,name=token,proto3" json:"token,omitempty"`
	Target    *entities.TargetEntity `protobuf:"bytes,2,opt,name=target,proto3" json:"target,omitempty"`
	CommentId string                 `protobuf:"bytes,3,opt,name=comment_id,json=commentId,proto3" json:"comment_id,omitempty"`
	Act       bool                   `protobuf:"varint,4,opt,name=act,proto3" json:"act,omitempty"`
	Model     string                 `protobuf:"bytes,5,opt,name=model,proto3" json:"model,omitempty"`
}

func (x *ExplainRequest) Reset() {
	*x = ExplainRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_services_robin_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ExplainRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ExplainRequest) ProtoMessage() {}

func (x *ExplainRequest) ProtoReflect() protoreflect.Message {
	mi := &file_services_robin_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ExplainRequest.ProtoReflect.Descriptor instead.
func (*ExplainRequest) Descriptor() ([]byte, []int) {
	return file_services_robin_proto_rawDescGZIP(), []int{2}
}

func (x *ExplainRequest) GetToken() string {
	if x != nil {
		return x.Token
	}
	return ""
}

func (x *ExplainRequest) GetTarget() *entities.TargetEntity {
	if x != nil {
		return x.Target
	}
	return nil
}

func (x *ExplainRequest) GetCommentId() string {
	if x != nil {
		return x.CommentId
	}
	return ""
}

func (x *ExplainRequest) GetAct() bool {
	if x != nil {
		return x.Act
	}
	return false
}

func (x *ExplainRequest) GetModel() string {
	if x != nil {
		return x.Model
	}
	return ""
}

type ExplainReply struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Explanation string `protobuf:"bytes,1,opt,name=explanation,proto3" json:"explanation,omitempty"`
}

func (x *ExplainReply) Reset() {
	*x = ExplainReply{}
	if protoimpl.UnsafeEnabled {
		mi := &file_services_robin_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ExplainReply) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ExplainReply) ProtoMessage() {}

func (x *ExplainReply) ProtoReflect() protoreflect.Message {
	mi := &file_services_robin_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ExplainReply.ProtoReflect.Descriptor instead.
func (*ExplainReply) Descriptor() ([]byte, []int) {
	return file_services_robin_proto_rawDescGZIP(), []int{3}
}

func (x *ExplainReply) GetExplanation() string {
	if x != nil {
		return x.Explanation
	}
	return ""
}

type PromptRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Token  string                 `protobuf:"bytes,1,opt,name=token,proto3" json:"token,omitempty"`
	Target *entities.TargetEntity `protobuf:"bytes,2,opt,name=target,proto3" json:"target,omitempty"`
	Prompt string                 `protobuf:"bytes,3,opt,name=prompt,proto3" json:"prompt,omitempty"`
	Act    bool                   `protobuf:"varint,4,opt,name=act,proto3" json:"act,omitempty"`
	Model  string                 `protobuf:"bytes,5,opt,name=model,proto3" json:"model,omitempty"`
}

func (x *PromptRequest) Reset() {
	*x = PromptRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_services_robin_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PromptRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PromptRequest) ProtoMessage() {}

func (x *PromptRequest) ProtoReflect() protoreflect.Message {
	mi := &file_services_robin_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PromptRequest.ProtoReflect.Descriptor instead.
func (*PromptRequest) Descriptor() ([]byte, []int) {
	return file_services_robin_proto_rawDescGZIP(), []int{4}
}

func (x *PromptRequest) GetToken() string {
	if x != nil {
		return x.Token
	}
	return ""
}

func (x *PromptRequest) GetTarget() *entities.TargetEntity {
	if x != nil {
		return x.Target
	}
	return nil
}

func (x *PromptRequest) GetPrompt() string {
	if x != nil {
		return x.Prompt
	}
	return ""
}

func (x *PromptRequest) GetAct() bool {
	if x != nil {
		return x.Act
	}
	return false
}

func (x *PromptRequest) GetModel() string {
	if x != nil {
		return x.Model
	}
	return ""
}

type PromptReply struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Reply string `protobuf:"bytes,1,opt,name=reply,proto3" json:"reply,omitempty"`
}

func (x *PromptReply) Reset() {
	*x = PromptReply{}
	if protoimpl.UnsafeEnabled {
		mi := &file_services_robin_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PromptReply) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PromptReply) ProtoMessage() {}

func (x *PromptReply) ProtoReflect() protoreflect.Message {
	mi := &file_services_robin_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PromptReply.ProtoReflect.Descriptor instead.
func (*PromptReply) Descriptor() ([]byte, []int) {
	return file_services_robin_proto_rawDescGZIP(), []int{5}
}

func (x *PromptReply) GetReply() string {
	if x != nil {
		return x.Reply
	}
	return ""
}

type RawPromptRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Token        string                 `protobuf:"bytes,1,opt,name=token,proto3" json:"token,omitempty"`
	Target       *entities.TargetEntity `protobuf:"bytes,2,opt,name=target,proto3" json:"target,omitempty"`
	SystemPrompt string                 `protobuf:"bytes,3,opt,name=system_prompt,json=systemPrompt,proto3" json:"system_prompt,omitempty"`
	UserPrompt   string                 `protobuf:"bytes,4,opt,name=user_prompt,json=userPrompt,proto3" json:"user_prompt,omitempty"`
	Act          bool                   `protobuf:"varint,5,opt,name=act,proto3" json:"act,omitempty"`
	Model        string                 `protobuf:"bytes,6,opt,name=model,proto3" json:"model,omitempty"`
}

func (x *RawPromptRequest) Reset() {
	*x = RawPromptRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_services_robin_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RawPromptRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RawPromptRequest) ProtoMessage() {}

func (x *RawPromptRequest) ProtoReflect() protoreflect.Message {
	mi := &file_services_robin_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RawPromptRequest.ProtoReflect.Descriptor instead.
func (*RawPromptRequest) Descriptor() ([]byte, []int) {
	return file_services_robin_proto_rawDescGZIP(), []int{6}
}

func (x *RawPromptRequest) GetToken() string {
	if x != nil {
		return x.Token
	}
	return ""
}

func (x *RawPromptRequest) GetTarget() *entities.TargetEntity {
	if x != nil {
		return x.Target
	}
	return nil
}

func (x *RawPromptRequest) GetSystemPrompt() string {
	if x != nil {
		return x.SystemPrompt
	}
	return ""
}

func (x *RawPromptRequest) GetUserPrompt() string {
	if x != nil {
		return x.UserPrompt
	}
	return ""
}

func (x *RawPromptRequest) GetAct() bool {
	if x != nil {
		return x.Act
	}
	return false
}

func (x *RawPromptRequest) GetModel() string {
	if x != nil {
		return x.Model
	}
	return ""
}

type SummarizeRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Token    string                 `protobuf:"bytes,1,opt,name=token,proto3" json:"token,omitempty"`
	Target   *entities.TargetEntity `protobuf:"bytes,2,opt,name=target,proto3" json:"target,omitempty"`
	Act      bool                   `protobuf:"varint,3,opt,name=act,proto3" json:"act,omitempty"`
	Extended bool                   `protobuf:"varint,4,opt,name=extended,proto3" json:"extended,omitempty"`
	Model    string                 `protobuf:"bytes,5,opt,name=model,proto3" json:"model,omitempty"`
}

func (x *SummarizeRequest) Reset() {
	*x = SummarizeRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_services_robin_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SummarizeRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SummarizeRequest) ProtoMessage() {}

func (x *SummarizeRequest) ProtoReflect() protoreflect.Message {
	mi := &file_services_robin_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SummarizeRequest.ProtoReflect.Descriptor instead.
func (*SummarizeRequest) Descriptor() ([]byte, []int) {
	return file_services_robin_proto_rawDescGZIP(), []int{7}
}

func (x *SummarizeRequest) GetToken() string {
	if x != nil {
		return x.Token
	}
	return ""
}

func (x *SummarizeRequest) GetTarget() *entities.TargetEntity {
	if x != nil {
		return x.Target
	}
	return nil
}

func (x *SummarizeRequest) GetAct() bool {
	if x != nil {
		return x.Act
	}
	return false
}

func (x *SummarizeRequest) GetExtended() bool {
	if x != nil {
		return x.Extended
	}
	return false
}

func (x *SummarizeRequest) GetModel() string {
	if x != nil {
		return x.Model
	}
	return ""
}

type SummarizeReply struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Summary string `protobuf:"bytes,1,opt,name=summary,proto3" json:"summary,omitempty"`
}

func (x *SummarizeReply) Reset() {
	*x = SummarizeReply{}
	if protoimpl.UnsafeEnabled {
		mi := &file_services_robin_proto_msgTypes[8]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SummarizeReply) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SummarizeReply) ProtoMessage() {}

func (x *SummarizeReply) ProtoReflect() protoreflect.Message {
	mi := &file_services_robin_proto_msgTypes[8]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SummarizeReply.ProtoReflect.Descriptor instead.
func (*SummarizeReply) Descriptor() ([]byte, []int) {
	return file_services_robin_proto_rawDescGZIP(), []int{8}
}

func (x *SummarizeReply) GetSummary() string {
	if x != nil {
		return x.Summary
	}
	return ""
}

var File_services_robin_proto protoreflect.FileDescriptor

var file_services_robin_proto_rawDesc = []byte{
	0x0a, 0x14, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x2f, 0x72, 0x6f, 0x62, 0x69, 0x6e,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x08, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73,
	0x1a, 0x1c, 0x65, 0x6e, 0x74, 0x69, 0x74, 0x69, 0x65, 0x73, 0x2f, 0x74, 0x61, 0x72, 0x67, 0x65,
	0x74, 0x5f, 0x65, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x5d,
	0x0a, 0x15, 0x43, 0x68, 0x61, 0x74, 0x49, 0x6e, 0x43, 0x6f, 0x64, 0x65, 0x48, 0x6f, 0x73, 0x74,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x14, 0x0a, 0x05, 0x74, 0x6f, 0x6b, 0x65, 0x6e,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x12, 0x2e, 0x0a,
	0x06, 0x74, 0x61, 0x72, 0x67, 0x65, 0x74, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x16, 0x2e,
	0x65, 0x6e, 0x74, 0x69, 0x74, 0x69, 0x65, 0x73, 0x2e, 0x54, 0x61, 0x72, 0x67, 0x65, 0x74, 0x45,
	0x6e, 0x74, 0x69, 0x74, 0x79, 0x52, 0x06, 0x74, 0x61, 0x72, 0x67, 0x65, 0x74, 0x22, 0x2b, 0x0a,
	0x13, 0x43, 0x68, 0x61, 0x74, 0x49, 0x6e, 0x43, 0x6f, 0x64, 0x65, 0x48, 0x6f, 0x73, 0x74, 0x52,
	0x65, 0x70, 0x6c, 0x79, 0x12, 0x14, 0x0a, 0x05, 0x72, 0x65, 0x70, 0x6c, 0x79, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x05, 0x72, 0x65, 0x70, 0x6c, 0x79, 0x22, 0x9d, 0x01, 0x0a, 0x0e, 0x45,
	0x78, 0x70, 0x6c, 0x61, 0x69, 0x6e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x14, 0x0a,
	0x05, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x74, 0x6f,
	0x6b, 0x65, 0x6e, 0x12, 0x2e, 0x0a, 0x06, 0x74, 0x61, 0x72, 0x67, 0x65, 0x74, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x0b, 0x32, 0x16, 0x2e, 0x65, 0x6e, 0x74, 0x69, 0x74, 0x69, 0x65, 0x73, 0x2e, 0x54,
	0x61, 0x72, 0x67, 0x65, 0x74, 0x45, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x52, 0x06, 0x74, 0x61, 0x72,
	0x67, 0x65, 0x74, 0x12, 0x1d, 0x0a, 0x0a, 0x63, 0x6f, 0x6d, 0x6d, 0x65, 0x6e, 0x74, 0x5f, 0x69,
	0x64, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x63, 0x6f, 0x6d, 0x6d, 0x65, 0x6e, 0x74,
	0x49, 0x64, 0x12, 0x10, 0x0a, 0x03, 0x61, 0x63, 0x74, 0x18, 0x04, 0x20, 0x01, 0x28, 0x08, 0x52,
	0x03, 0x61, 0x63, 0x74, 0x12, 0x14, 0x0a, 0x05, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x18, 0x05, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x05, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x22, 0x30, 0x0a, 0x0c, 0x45, 0x78,
	0x70, 0x6c, 0x61, 0x69, 0x6e, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x12, 0x20, 0x0a, 0x0b, 0x65, 0x78,
	0x70, 0x6c, 0x61, 0x6e, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x0b, 0x65, 0x78, 0x70, 0x6c, 0x61, 0x6e, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x22, 0x95, 0x01, 0x0a,
	0x0d, 0x50, 0x72, 0x6f, 0x6d, 0x70, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x14,
	0x0a, 0x05, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x74,
	0x6f, 0x6b, 0x65, 0x6e, 0x12, 0x2e, 0x0a, 0x06, 0x74, 0x61, 0x72, 0x67, 0x65, 0x74, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x0b, 0x32, 0x16, 0x2e, 0x65, 0x6e, 0x74, 0x69, 0x74, 0x69, 0x65, 0x73, 0x2e,
	0x54, 0x61, 0x72, 0x67, 0x65, 0x74, 0x45, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x52, 0x06, 0x74, 0x61,
	0x72, 0x67, 0x65, 0x74, 0x12, 0x16, 0x0a, 0x06, 0x70, 0x72, 0x6f, 0x6d, 0x70, 0x74, 0x18, 0x03,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x70, 0x72, 0x6f, 0x6d, 0x70, 0x74, 0x12, 0x10, 0x0a, 0x03,
	0x61, 0x63, 0x74, 0x18, 0x04, 0x20, 0x01, 0x28, 0x08, 0x52, 0x03, 0x61, 0x63, 0x74, 0x12, 0x14,
	0x0a, 0x05, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x6d,
	0x6f, 0x64, 0x65, 0x6c, 0x22, 0x23, 0x0a, 0x0b, 0x50, 0x72, 0x6f, 0x6d, 0x70, 0x74, 0x52, 0x65,
	0x70, 0x6c, 0x79, 0x12, 0x14, 0x0a, 0x05, 0x72, 0x65, 0x70, 0x6c, 0x79, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x05, 0x72, 0x65, 0x70, 0x6c, 0x79, 0x22, 0xc6, 0x01, 0x0a, 0x10, 0x52, 0x61,
	0x77, 0x50, 0x72, 0x6f, 0x6d, 0x70, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x14,
	0x0a, 0x05, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x74,
	0x6f, 0x6b, 0x65, 0x6e, 0x12, 0x2e, 0x0a, 0x06, 0x74, 0x61, 0x72, 0x67, 0x65, 0x74, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x0b, 0x32, 0x16, 0x2e, 0x65, 0x6e, 0x74, 0x69, 0x74, 0x69, 0x65, 0x73, 0x2e,
	0x54, 0x61, 0x72, 0x67, 0x65, 0x74, 0x45, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x52, 0x06, 0x74, 0x61,
	0x72, 0x67, 0x65, 0x74, 0x12, 0x23, 0x0a, 0x0d, 0x73, 0x79, 0x73, 0x74, 0x65, 0x6d, 0x5f, 0x70,
	0x72, 0x6f, 0x6d, 0x70, 0x74, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0c, 0x73, 0x79, 0x73,
	0x74, 0x65, 0x6d, 0x50, 0x72, 0x6f, 0x6d, 0x70, 0x74, 0x12, 0x1f, 0x0a, 0x0b, 0x75, 0x73, 0x65,
	0x72, 0x5f, 0x70, 0x72, 0x6f, 0x6d, 0x70, 0x74, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a,
	0x75, 0x73, 0x65, 0x72, 0x50, 0x72, 0x6f, 0x6d, 0x70, 0x74, 0x12, 0x10, 0x0a, 0x03, 0x61, 0x63,
	0x74, 0x18, 0x05, 0x20, 0x01, 0x28, 0x08, 0x52, 0x03, 0x61, 0x63, 0x74, 0x12, 0x14, 0x0a, 0x05,
	0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x18, 0x06, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x6d, 0x6f, 0x64,
	0x65, 0x6c, 0x22, 0x9c, 0x01, 0x0a, 0x10, 0x53, 0x75, 0x6d, 0x6d, 0x61, 0x72, 0x69, 0x7a, 0x65,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x14, 0x0a, 0x05, 0x74, 0x6f, 0x6b, 0x65, 0x6e,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x12, 0x2e, 0x0a,
	0x06, 0x74, 0x61, 0x72, 0x67, 0x65, 0x74, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x16, 0x2e,
	0x65, 0x6e, 0x74, 0x69, 0x74, 0x69, 0x65, 0x73, 0x2e, 0x54, 0x61, 0x72, 0x67, 0x65, 0x74, 0x45,
	0x6e, 0x74, 0x69, 0x74, 0x79, 0x52, 0x06, 0x74, 0x61, 0x72, 0x67, 0x65, 0x74, 0x12, 0x10, 0x0a,
	0x03, 0x61, 0x63, 0x74, 0x18, 0x03, 0x20, 0x01, 0x28, 0x08, 0x52, 0x03, 0x61, 0x63, 0x74, 0x12,
	0x1a, 0x0a, 0x08, 0x65, 0x78, 0x74, 0x65, 0x6e, 0x64, 0x65, 0x64, 0x18, 0x04, 0x20, 0x01, 0x28,
	0x08, 0x52, 0x08, 0x65, 0x78, 0x74, 0x65, 0x6e, 0x64, 0x65, 0x64, 0x12, 0x14, 0x0a, 0x05, 0x6d,
	0x6f, 0x64, 0x65, 0x6c, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x6d, 0x6f, 0x64, 0x65,
	0x6c, 0x22, 0x2a, 0x0a, 0x0e, 0x53, 0x75, 0x6d, 0x6d, 0x61, 0x72, 0x69, 0x7a, 0x65, 0x52, 0x65,
	0x70, 0x6c, 0x79, 0x12, 0x18, 0x0a, 0x07, 0x73, 0x75, 0x6d, 0x6d, 0x61, 0x72, 0x79, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x73, 0x75, 0x6d, 0x6d, 0x61, 0x72, 0x79, 0x32, 0xdd, 0x02,
	0x0a, 0x05, 0x52, 0x6f, 0x62, 0x69, 0x6e, 0x12, 0x52, 0x0a, 0x0e, 0x43, 0x68, 0x61, 0x74, 0x49,
	0x6e, 0x43, 0x6f, 0x64, 0x65, 0x48, 0x6f, 0x73, 0x74, 0x12, 0x1f, 0x2e, 0x73, 0x65, 0x72, 0x76,
	0x69, 0x63, 0x65, 0x73, 0x2e, 0x43, 0x68, 0x61, 0x74, 0x49, 0x6e, 0x43, 0x6f, 0x64, 0x65, 0x48,
	0x6f, 0x73, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1d, 0x2e, 0x73, 0x65, 0x72,
	0x76, 0x69, 0x63, 0x65, 0x73, 0x2e, 0x43, 0x68, 0x61, 0x74, 0x49, 0x6e, 0x43, 0x6f, 0x64, 0x65,
	0x48, 0x6f, 0x73, 0x74, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x22, 0x00, 0x12, 0x3d, 0x0a, 0x07, 0x45,
	0x78, 0x70, 0x6c, 0x61, 0x69, 0x6e, 0x12, 0x18, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65,
	0x73, 0x2e, 0x45, 0x78, 0x70, 0x6c, 0x61, 0x69, 0x6e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x1a, 0x16, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x2e, 0x45, 0x78, 0x70, 0x6c,
	0x61, 0x69, 0x6e, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x22, 0x00, 0x12, 0x3a, 0x0a, 0x06, 0x50, 0x72,
	0x6f, 0x6d, 0x70, 0x74, 0x12, 0x17, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x2e,
	0x50, 0x72, 0x6f, 0x6d, 0x70, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x15, 0x2e,
	0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x2e, 0x50, 0x72, 0x6f, 0x6d, 0x70, 0x74, 0x52,
	0x65, 0x70, 0x6c, 0x79, 0x22, 0x00, 0x12, 0x40, 0x0a, 0x09, 0x52, 0x61, 0x77, 0x50, 0x72, 0x6f,
	0x6d, 0x70, 0x74, 0x12, 0x1a, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x2e, 0x52,
	0x61, 0x77, 0x50, 0x72, 0x6f, 0x6d, 0x70, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a,
	0x15, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x2e, 0x50, 0x72, 0x6f, 0x6d, 0x70,
	0x74, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x22, 0x00, 0x12, 0x43, 0x0a, 0x09, 0x53, 0x75, 0x6d, 0x6d,
	0x61, 0x72, 0x69, 0x7a, 0x65, 0x12, 0x1a, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73,
	0x2e, 0x53, 0x75, 0x6d, 0x6d, 0x61, 0x72, 0x69, 0x7a, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x1a, 0x18, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x2e, 0x53, 0x75, 0x6d,
	0x6d, 0x61, 0x72, 0x69, 0x7a, 0x65, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x22, 0x00, 0x42, 0x26, 0x5a,
	0x24, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x72, 0x65, 0x76, 0x69,
	0x65, 0x77, 0x70, 0x61, 0x64, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x67, 0x6f, 0x2f, 0x73, 0x65, 0x72,
	0x76, 0x69, 0x63, 0x65, 0x73, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_services_robin_proto_rawDescOnce sync.Once
	file_services_robin_proto_rawDescData = file_services_robin_proto_rawDesc
)

func file_services_robin_proto_rawDescGZIP() []byte {
	file_services_robin_proto_rawDescOnce.Do(func() {
		file_services_robin_proto_rawDescData = protoimpl.X.CompressGZIP(file_services_robin_proto_rawDescData)
	})
	return file_services_robin_proto_rawDescData
}

var file_services_robin_proto_msgTypes = make([]protoimpl.MessageInfo, 9)
var file_services_robin_proto_goTypes = []interface{}{
	(*ChatInCodeHostRequest)(nil), // 0: services.ChatInCodeHostRequest
	(*ChatInCodeHostReply)(nil),   // 1: services.ChatInCodeHostReply
	(*ExplainRequest)(nil),        // 2: services.ExplainRequest
	(*ExplainReply)(nil),          // 3: services.ExplainReply
	(*PromptRequest)(nil),         // 4: services.PromptRequest
	(*PromptReply)(nil),           // 5: services.PromptReply
	(*RawPromptRequest)(nil),      // 6: services.RawPromptRequest
	(*SummarizeRequest)(nil),      // 7: services.SummarizeRequest
	(*SummarizeReply)(nil),        // 8: services.SummarizeReply
	(*entities.TargetEntity)(nil), // 9: entities.TargetEntity
}
var file_services_robin_proto_depIdxs = []int32{
	9,  // 0: services.ChatInCodeHostRequest.target:type_name -> entities.TargetEntity
	9,  // 1: services.ExplainRequest.target:type_name -> entities.TargetEntity
	9,  // 2: services.PromptRequest.target:type_name -> entities.TargetEntity
	9,  // 3: services.RawPromptRequest.target:type_name -> entities.TargetEntity
	9,  // 4: services.SummarizeRequest.target:type_name -> entities.TargetEntity
	0,  // 5: services.Robin.ChatInCodeHost:input_type -> services.ChatInCodeHostRequest
	2,  // 6: services.Robin.Explain:input_type -> services.ExplainRequest
	4,  // 7: services.Robin.Prompt:input_type -> services.PromptRequest
	6,  // 8: services.Robin.RawPrompt:input_type -> services.RawPromptRequest
	7,  // 9: services.Robin.Summarize:input_type -> services.SummarizeRequest
	1,  // 10: services.Robin.ChatInCodeHost:output_type -> services.ChatInCodeHostReply
	3,  // 11: services.Robin.Explain:output_type -> services.ExplainReply
	5,  // 12: services.Robin.Prompt:output_type -> services.PromptReply
	5,  // 13: services.Robin.RawPrompt:output_type -> services.PromptReply
	8,  // 14: services.Robin.Summarize:output_type -> services.SummarizeReply
	10, // [10:15] is the sub-list for method output_type
	5,  // [5:10] is the sub-list for method input_type
	5,  // [5:5] is the sub-list for extension type_name
	5,  // [5:5] is the sub-list for extension extendee
	0,  // [0:5] is the sub-list for field type_name
}

func init() { file_services_robin_proto_init() }
func file_services_robin_proto_init() {
	if File_services_robin_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_services_robin_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ChatInCodeHostRequest); i {
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
		file_services_robin_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ChatInCodeHostReply); i {
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
		file_services_robin_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ExplainRequest); i {
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
		file_services_robin_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ExplainReply); i {
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
		file_services_robin_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*PromptRequest); i {
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
		file_services_robin_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*PromptReply); i {
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
		file_services_robin_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*RawPromptRequest); i {
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
		file_services_robin_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SummarizeRequest); i {
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
		file_services_robin_proto_msgTypes[8].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SummarizeReply); i {
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
			RawDescriptor: file_services_robin_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   9,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_services_robin_proto_goTypes,
		DependencyIndexes: file_services_robin_proto_depIdxs,
		MessageInfos:      file_services_robin_proto_msgTypes,
	}.Build()
	File_services_robin_proto = out.File
	file_services_robin_proto_rawDesc = nil
	file_services_robin_proto_goTypes = nil
	file_services_robin_proto_depIdxs = nil
}
