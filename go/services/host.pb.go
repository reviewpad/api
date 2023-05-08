// Copyright 2023 Explore.dev Unipessoal Lda. All Rights Reserved.
// Use of this source code is governed by a license that can be
// found in the LICENSE file.

// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.30.0
// 	protoc        v3.21.12
// source: services/host.proto

package services

import (
	codehost "github.com/reviewpad/api/go/codehost"
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

type PostGeneralCommentRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Host             entities.Host           `protobuf:"varint,1,opt,name=host,proto3,enum=entities.Host" json:"host,omitempty"`
	HostUri          string                  `protobuf:"bytes,2,opt,name=host_uri,json=hostUri,proto3" json:"host_uri,omitempty"`
	Slug             string                  `protobuf:"bytes,3,opt,name=slug,proto3" json:"slug,omitempty"`
	ExternalRepoId   string                  `protobuf:"bytes,4,opt,name=external_repo_id,json=externalRepoId,proto3" json:"external_repo_id,omitempty"`
	ReviewNumber     int32                   `protobuf:"varint,5,opt,name=review_number,json=reviewNumber,proto3" json:"review_number,omitempty"`
	ExternalReviewId string                  `protobuf:"bytes,6,opt,name=external_review_id,json=externalReviewId,proto3" json:"external_review_id,omitempty"`
	AccessToken      string                  `protobuf:"bytes,7,opt,name=access_token,json=accessToken,proto3" json:"access_token,omitempty"`
	Comment          *codehost.ReviewComment `protobuf:"bytes,8,opt,name=comment,proto3" json:"comment,omitempty"`
}

func (x *PostGeneralCommentRequest) Reset() {
	*x = PostGeneralCommentRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_services_host_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PostGeneralCommentRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PostGeneralCommentRequest) ProtoMessage() {}

func (x *PostGeneralCommentRequest) ProtoReflect() protoreflect.Message {
	mi := &file_services_host_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PostGeneralCommentRequest.ProtoReflect.Descriptor instead.
func (*PostGeneralCommentRequest) Descriptor() ([]byte, []int) {
	return file_services_host_proto_rawDescGZIP(), []int{0}
}

func (x *PostGeneralCommentRequest) GetHost() entities.Host {
	if x != nil {
		return x.Host
	}
	return entities.Host(0)
}

func (x *PostGeneralCommentRequest) GetHostUri() string {
	if x != nil {
		return x.HostUri
	}
	return ""
}

func (x *PostGeneralCommentRequest) GetSlug() string {
	if x != nil {
		return x.Slug
	}
	return ""
}

func (x *PostGeneralCommentRequest) GetExternalRepoId() string {
	if x != nil {
		return x.ExternalRepoId
	}
	return ""
}

func (x *PostGeneralCommentRequest) GetReviewNumber() int32 {
	if x != nil {
		return x.ReviewNumber
	}
	return 0
}

func (x *PostGeneralCommentRequest) GetExternalReviewId() string {
	if x != nil {
		return x.ExternalReviewId
	}
	return ""
}

func (x *PostGeneralCommentRequest) GetAccessToken() string {
	if x != nil {
		return x.AccessToken
	}
	return ""
}

func (x *PostGeneralCommentRequest) GetComment() *codehost.ReviewComment {
	if x != nil {
		return x.Comment
	}
	return nil
}

type PostGeneralCommentReply struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Comment *codehost.ReviewComment `protobuf:"bytes,1,opt,name=comment,proto3" json:"comment,omitempty"`
}

func (x *PostGeneralCommentReply) Reset() {
	*x = PostGeneralCommentReply{}
	if protoimpl.UnsafeEnabled {
		mi := &file_services_host_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PostGeneralCommentReply) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PostGeneralCommentReply) ProtoMessage() {}

func (x *PostGeneralCommentReply) ProtoReflect() protoreflect.Message {
	mi := &file_services_host_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PostGeneralCommentReply.ProtoReflect.Descriptor instead.
func (*PostGeneralCommentReply) Descriptor() ([]byte, []int) {
	return file_services_host_proto_rawDescGZIP(), []int{1}
}

func (x *PostGeneralCommentReply) GetComment() *codehost.ReviewComment {
	if x != nil {
		return x.Comment
	}
	return nil
}

type GetPullRequestRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Slug           string        `protobuf:"bytes,1,opt,name=slug,proto3" json:"slug,omitempty"`
	ExternalRepoId string        `protobuf:"bytes,2,opt,name=external_repo_id,json=externalRepoId,proto3" json:"external_repo_id,omitempty"`
	Host           entities.Host `protobuf:"varint,3,opt,name=host,proto3,enum=entities.Host" json:"host,omitempty"`
	HostUri        string        `protobuf:"bytes,4,opt,name=host_uri,json=hostUri,proto3" json:"host_uri,omitempty"`
	Number         int64         `protobuf:"varint,5,opt,name=number,proto3" json:"number,omitempty"`
	AccessToken    string        `protobuf:"bytes,6,opt,name=access_token,json=accessToken,proto3" json:"access_token,omitempty"`
}

func (x *GetPullRequestRequest) Reset() {
	*x = GetPullRequestRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_services_host_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetPullRequestRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetPullRequestRequest) ProtoMessage() {}

func (x *GetPullRequestRequest) ProtoReflect() protoreflect.Message {
	mi := &file_services_host_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetPullRequestRequest.ProtoReflect.Descriptor instead.
func (*GetPullRequestRequest) Descriptor() ([]byte, []int) {
	return file_services_host_proto_rawDescGZIP(), []int{2}
}

func (x *GetPullRequestRequest) GetSlug() string {
	if x != nil {
		return x.Slug
	}
	return ""
}

func (x *GetPullRequestRequest) GetExternalRepoId() string {
	if x != nil {
		return x.ExternalRepoId
	}
	return ""
}

func (x *GetPullRequestRequest) GetHost() entities.Host {
	if x != nil {
		return x.Host
	}
	return entities.Host(0)
}

func (x *GetPullRequestRequest) GetHostUri() string {
	if x != nil {
		return x.HostUri
	}
	return ""
}

func (x *GetPullRequestRequest) GetNumber() int64 {
	if x != nil {
		return x.Number
	}
	return 0
}

func (x *GetPullRequestRequest) GetAccessToken() string {
	if x != nil {
		return x.AccessToken
	}
	return ""
}

type GetPullRequestReply struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	PullRequest *codehost.PullRequest `protobuf:"bytes,1,opt,name=pull_request,json=pullRequest,proto3" json:"pull_request,omitempty"`
}

func (x *GetPullRequestReply) Reset() {
	*x = GetPullRequestReply{}
	if protoimpl.UnsafeEnabled {
		mi := &file_services_host_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetPullRequestReply) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetPullRequestReply) ProtoMessage() {}

func (x *GetPullRequestReply) ProtoReflect() protoreflect.Message {
	mi := &file_services_host_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetPullRequestReply.ProtoReflect.Descriptor instead.
func (*GetPullRequestReply) Descriptor() ([]byte, []int) {
	return file_services_host_proto_rawDescGZIP(), []int{3}
}

func (x *GetPullRequestReply) GetPullRequest() *codehost.PullRequest {
	if x != nil {
		return x.PullRequest
	}
	return nil
}

type GetPullRequestFilesRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Slug           string        `protobuf:"bytes,1,opt,name=slug,proto3" json:"slug,omitempty"`
	ExternalRepoId string        `protobuf:"bytes,2,opt,name=external_repo_id,json=externalRepoId,proto3" json:"external_repo_id,omitempty"`
	Host           entities.Host `protobuf:"varint,3,opt,name=host,proto3,enum=entities.Host" json:"host,omitempty"`
	HostUri        string        `protobuf:"bytes,4,opt,name=host_uri,json=hostUri,proto3" json:"host_uri,omitempty"`
	Number         int64         `protobuf:"varint,5,opt,name=number,proto3" json:"number,omitempty"`
	AccessToken    string        `protobuf:"bytes,6,opt,name=access_token,json=accessToken,proto3" json:"access_token,omitempty"`
}

func (x *GetPullRequestFilesRequest) Reset() {
	*x = GetPullRequestFilesRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_services_host_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetPullRequestFilesRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetPullRequestFilesRequest) ProtoMessage() {}

func (x *GetPullRequestFilesRequest) ProtoReflect() protoreflect.Message {
	mi := &file_services_host_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetPullRequestFilesRequest.ProtoReflect.Descriptor instead.
func (*GetPullRequestFilesRequest) Descriptor() ([]byte, []int) {
	return file_services_host_proto_rawDescGZIP(), []int{4}
}

func (x *GetPullRequestFilesRequest) GetSlug() string {
	if x != nil {
		return x.Slug
	}
	return ""
}

func (x *GetPullRequestFilesRequest) GetExternalRepoId() string {
	if x != nil {
		return x.ExternalRepoId
	}
	return ""
}

func (x *GetPullRequestFilesRequest) GetHost() entities.Host {
	if x != nil {
		return x.Host
	}
	return entities.Host(0)
}

func (x *GetPullRequestFilesRequest) GetHostUri() string {
	if x != nil {
		return x.HostUri
	}
	return ""
}

func (x *GetPullRequestFilesRequest) GetNumber() int64 {
	if x != nil {
		return x.Number
	}
	return 0
}

func (x *GetPullRequestFilesRequest) GetAccessToken() string {
	if x != nil {
		return x.AccessToken
	}
	return ""
}

type GetPullRequestFilesReply struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Files []*codehost.File `protobuf:"bytes,1,rep,name=files,proto3" json:"files,omitempty"`
}

func (x *GetPullRequestFilesReply) Reset() {
	*x = GetPullRequestFilesReply{}
	if protoimpl.UnsafeEnabled {
		mi := &file_services_host_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetPullRequestFilesReply) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetPullRequestFilesReply) ProtoMessage() {}

func (x *GetPullRequestFilesReply) ProtoReflect() protoreflect.Message {
	mi := &file_services_host_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetPullRequestFilesReply.ProtoReflect.Descriptor instead.
func (*GetPullRequestFilesReply) Descriptor() ([]byte, []int) {
	return file_services_host_proto_rawDescGZIP(), []int{5}
}

func (x *GetPullRequestFilesReply) GetFiles() []*codehost.File {
	if x != nil {
		return x.Files
	}
	return nil
}

var File_services_host_proto protoreflect.FileDescriptor

var file_services_host_proto_rawDesc = []byte{
	0x0a, 0x13, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x2f, 0x68, 0x6f, 0x73, 0x74, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x08, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x1a,
	0x18, 0x65, 0x6e, 0x74, 0x69, 0x74, 0x69, 0x65, 0x73, 0x2f, 0x63, 0x6f, 0x64, 0x65, 0x5f, 0x68,
	0x6f, 0x73, 0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x16, 0x63, 0x6f, 0x64, 0x65, 0x68,
	0x6f, 0x73, 0x74, 0x2f, 0x63, 0x6f, 0x6d, 0x6d, 0x65, 0x6e, 0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x1a, 0x1b, 0x63, 0x6f, 0x64, 0x65, 0x68, 0x6f, 0x73, 0x74, 0x2f, 0x70, 0x75, 0x6c, 0x6c,
	0x5f, 0x72, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x13,
	0x63, 0x6f, 0x64, 0x65, 0x68, 0x6f, 0x73, 0x74, 0x2f, 0x66, 0x69, 0x6c, 0x65, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x22, 0xc1, 0x02, 0x0a, 0x19, 0x50, 0x6f, 0x73, 0x74, 0x47, 0x65, 0x6e, 0x65,
	0x72, 0x61, 0x6c, 0x43, 0x6f, 0x6d, 0x6d, 0x65, 0x6e, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x12, 0x22, 0x0a, 0x04, 0x68, 0x6f, 0x73, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0e, 0x32,
	0x0e, 0x2e, 0x65, 0x6e, 0x74, 0x69, 0x74, 0x69, 0x65, 0x73, 0x2e, 0x48, 0x6f, 0x73, 0x74, 0x52,
	0x04, 0x68, 0x6f, 0x73, 0x74, 0x12, 0x19, 0x0a, 0x08, 0x68, 0x6f, 0x73, 0x74, 0x5f, 0x75, 0x72,
	0x69, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x68, 0x6f, 0x73, 0x74, 0x55, 0x72, 0x69,
	0x12, 0x12, 0x0a, 0x04, 0x73, 0x6c, 0x75, 0x67, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04,
	0x73, 0x6c, 0x75, 0x67, 0x12, 0x28, 0x0a, 0x10, 0x65, 0x78, 0x74, 0x65, 0x72, 0x6e, 0x61, 0x6c,
	0x5f, 0x72, 0x65, 0x70, 0x6f, 0x5f, 0x69, 0x64, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0e,
	0x65, 0x78, 0x74, 0x65, 0x72, 0x6e, 0x61, 0x6c, 0x52, 0x65, 0x70, 0x6f, 0x49, 0x64, 0x12, 0x23,
	0x0a, 0x0d, 0x72, 0x65, 0x76, 0x69, 0x65, 0x77, 0x5f, 0x6e, 0x75, 0x6d, 0x62, 0x65, 0x72, 0x18,
	0x05, 0x20, 0x01, 0x28, 0x05, 0x52, 0x0c, 0x72, 0x65, 0x76, 0x69, 0x65, 0x77, 0x4e, 0x75, 0x6d,
	0x62, 0x65, 0x72, 0x12, 0x2c, 0x0a, 0x12, 0x65, 0x78, 0x74, 0x65, 0x72, 0x6e, 0x61, 0x6c, 0x5f,
	0x72, 0x65, 0x76, 0x69, 0x65, 0x77, 0x5f, 0x69, 0x64, 0x18, 0x06, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x10, 0x65, 0x78, 0x74, 0x65, 0x72, 0x6e, 0x61, 0x6c, 0x52, 0x65, 0x76, 0x69, 0x65, 0x77, 0x49,
	0x64, 0x12, 0x21, 0x0a, 0x0c, 0x61, 0x63, 0x63, 0x65, 0x73, 0x73, 0x5f, 0x74, 0x6f, 0x6b, 0x65,
	0x6e, 0x18, 0x07, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x61, 0x63, 0x63, 0x65, 0x73, 0x73, 0x54,
	0x6f, 0x6b, 0x65, 0x6e, 0x12, 0x31, 0x0a, 0x07, 0x63, 0x6f, 0x6d, 0x6d, 0x65, 0x6e, 0x74, 0x18,
	0x08, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x17, 0x2e, 0x63, 0x6f, 0x64, 0x65, 0x68, 0x6f, 0x73, 0x74,
	0x2e, 0x52, 0x65, 0x76, 0x69, 0x65, 0x77, 0x43, 0x6f, 0x6d, 0x6d, 0x65, 0x6e, 0x74, 0x52, 0x07,
	0x63, 0x6f, 0x6d, 0x6d, 0x65, 0x6e, 0x74, 0x22, 0x4c, 0x0a, 0x17, 0x50, 0x6f, 0x73, 0x74, 0x47,
	0x65, 0x6e, 0x65, 0x72, 0x61, 0x6c, 0x43, 0x6f, 0x6d, 0x6d, 0x65, 0x6e, 0x74, 0x52, 0x65, 0x70,
	0x6c, 0x79, 0x12, 0x31, 0x0a, 0x07, 0x63, 0x6f, 0x6d, 0x6d, 0x65, 0x6e, 0x74, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x0b, 0x32, 0x17, 0x2e, 0x63, 0x6f, 0x64, 0x65, 0x68, 0x6f, 0x73, 0x74, 0x2e, 0x52,
	0x65, 0x76, 0x69, 0x65, 0x77, 0x43, 0x6f, 0x6d, 0x6d, 0x65, 0x6e, 0x74, 0x52, 0x07, 0x63, 0x6f,
	0x6d, 0x6d, 0x65, 0x6e, 0x74, 0x22, 0xcf, 0x01, 0x0a, 0x15, 0x47, 0x65, 0x74, 0x50, 0x75, 0x6c,
	0x6c, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12,
	0x12, 0x0a, 0x04, 0x73, 0x6c, 0x75, 0x67, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x73,
	0x6c, 0x75, 0x67, 0x12, 0x28, 0x0a, 0x10, 0x65, 0x78, 0x74, 0x65, 0x72, 0x6e, 0x61, 0x6c, 0x5f,
	0x72, 0x65, 0x70, 0x6f, 0x5f, 0x69, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0e, 0x65,
	0x78, 0x74, 0x65, 0x72, 0x6e, 0x61, 0x6c, 0x52, 0x65, 0x70, 0x6f, 0x49, 0x64, 0x12, 0x22, 0x0a,
	0x04, 0x68, 0x6f, 0x73, 0x74, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x0e, 0x2e, 0x65, 0x6e,
	0x74, 0x69, 0x74, 0x69, 0x65, 0x73, 0x2e, 0x48, 0x6f, 0x73, 0x74, 0x52, 0x04, 0x68, 0x6f, 0x73,
	0x74, 0x12, 0x19, 0x0a, 0x08, 0x68, 0x6f, 0x73, 0x74, 0x5f, 0x75, 0x72, 0x69, 0x18, 0x04, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x07, 0x68, 0x6f, 0x73, 0x74, 0x55, 0x72, 0x69, 0x12, 0x16, 0x0a, 0x06,
	0x6e, 0x75, 0x6d, 0x62, 0x65, 0x72, 0x18, 0x05, 0x20, 0x01, 0x28, 0x03, 0x52, 0x06, 0x6e, 0x75,
	0x6d, 0x62, 0x65, 0x72, 0x12, 0x21, 0x0a, 0x0c, 0x61, 0x63, 0x63, 0x65, 0x73, 0x73, 0x5f, 0x74,
	0x6f, 0x6b, 0x65, 0x6e, 0x18, 0x06, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x61, 0x63, 0x63, 0x65,
	0x73, 0x73, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x22, 0x4f, 0x0a, 0x13, 0x47, 0x65, 0x74, 0x50, 0x75,
	0x6c, 0x6c, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x12, 0x38,
	0x0a, 0x0c, 0x70, 0x75, 0x6c, 0x6c, 0x5f, 0x72, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x0b, 0x32, 0x15, 0x2e, 0x63, 0x6f, 0x64, 0x65, 0x68, 0x6f, 0x73, 0x74, 0x2e,
	0x50, 0x75, 0x6c, 0x6c, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x52, 0x0b, 0x70, 0x75, 0x6c,
	0x6c, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x22, 0xd4, 0x01, 0x0a, 0x1a, 0x47, 0x65, 0x74,
	0x50, 0x75, 0x6c, 0x6c, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x46, 0x69, 0x6c, 0x65, 0x73,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x12, 0x0a, 0x04, 0x73, 0x6c, 0x75, 0x67, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x73, 0x6c, 0x75, 0x67, 0x12, 0x28, 0x0a, 0x10, 0x65,
	0x78, 0x74, 0x65, 0x72, 0x6e, 0x61, 0x6c, 0x5f, 0x72, 0x65, 0x70, 0x6f, 0x5f, 0x69, 0x64, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0e, 0x65, 0x78, 0x74, 0x65, 0x72, 0x6e, 0x61, 0x6c, 0x52,
	0x65, 0x70, 0x6f, 0x49, 0x64, 0x12, 0x22, 0x0a, 0x04, 0x68, 0x6f, 0x73, 0x74, 0x18, 0x03, 0x20,
	0x01, 0x28, 0x0e, 0x32, 0x0e, 0x2e, 0x65, 0x6e, 0x74, 0x69, 0x74, 0x69, 0x65, 0x73, 0x2e, 0x48,
	0x6f, 0x73, 0x74, 0x52, 0x04, 0x68, 0x6f, 0x73, 0x74, 0x12, 0x19, 0x0a, 0x08, 0x68, 0x6f, 0x73,
	0x74, 0x5f, 0x75, 0x72, 0x69, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x68, 0x6f, 0x73,
	0x74, 0x55, 0x72, 0x69, 0x12, 0x16, 0x0a, 0x06, 0x6e, 0x75, 0x6d, 0x62, 0x65, 0x72, 0x18, 0x05,
	0x20, 0x01, 0x28, 0x03, 0x52, 0x06, 0x6e, 0x75, 0x6d, 0x62, 0x65, 0x72, 0x12, 0x21, 0x0a, 0x0c,
	0x61, 0x63, 0x63, 0x65, 0x73, 0x73, 0x5f, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x18, 0x06, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x0b, 0x61, 0x63, 0x63, 0x65, 0x73, 0x73, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x22,
	0x40, 0x0a, 0x18, 0x47, 0x65, 0x74, 0x50, 0x75, 0x6c, 0x6c, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x46, 0x69, 0x6c, 0x65, 0x73, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x12, 0x24, 0x0a, 0x05, 0x66,
	0x69, 0x6c, 0x65, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x0e, 0x2e, 0x63, 0x6f, 0x64,
	0x65, 0x68, 0x6f, 0x73, 0x74, 0x2e, 0x46, 0x69, 0x6c, 0x65, 0x52, 0x05, 0x66, 0x69, 0x6c, 0x65,
	0x73, 0x32, 0x9d, 0x02, 0x0a, 0x04, 0x48, 0x6f, 0x73, 0x74, 0x12, 0x5e, 0x0a, 0x12, 0x50, 0x6f,
	0x73, 0x74, 0x47, 0x65, 0x6e, 0x65, 0x72, 0x61, 0x6c, 0x43, 0x6f, 0x6d, 0x6d, 0x65, 0x6e, 0x74,
	0x12, 0x23, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x2e, 0x50, 0x6f, 0x73, 0x74,
	0x47, 0x65, 0x6e, 0x65, 0x72, 0x61, 0x6c, 0x43, 0x6f, 0x6d, 0x6d, 0x65, 0x6e, 0x74, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x21, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73,
	0x2e, 0x50, 0x6f, 0x73, 0x74, 0x47, 0x65, 0x6e, 0x65, 0x72, 0x61, 0x6c, 0x43, 0x6f, 0x6d, 0x6d,
	0x65, 0x6e, 0x74, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x22, 0x00, 0x12, 0x52, 0x0a, 0x0e, 0x47, 0x65,
	0x74, 0x50, 0x75, 0x6c, 0x6c, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x1f, 0x2e, 0x73,
	0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x2e, 0x47, 0x65, 0x74, 0x50, 0x75, 0x6c, 0x6c, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1d, 0x2e,
	0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x2e, 0x47, 0x65, 0x74, 0x50, 0x75, 0x6c, 0x6c,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x22, 0x00, 0x12, 0x61,
	0x0a, 0x13, 0x47, 0x65, 0x74, 0x50, 0x75, 0x6c, 0x6c, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x46, 0x69, 0x6c, 0x65, 0x73, 0x12, 0x24, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73,
	0x2e, 0x47, 0x65, 0x74, 0x50, 0x75, 0x6c, 0x6c, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x46,
	0x69, 0x6c, 0x65, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x22, 0x2e, 0x73, 0x65,
	0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x2e, 0x47, 0x65, 0x74, 0x50, 0x75, 0x6c, 0x6c, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x46, 0x69, 0x6c, 0x65, 0x73, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x22,
	0x00, 0x42, 0x26, 0x5a, 0x24, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f,
	0x72, 0x65, 0x76, 0x69, 0x65, 0x77, 0x70, 0x61, 0x64, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x67, 0x6f,
	0x2f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x33,
}

var (
	file_services_host_proto_rawDescOnce sync.Once
	file_services_host_proto_rawDescData = file_services_host_proto_rawDesc
)

func file_services_host_proto_rawDescGZIP() []byte {
	file_services_host_proto_rawDescOnce.Do(func() {
		file_services_host_proto_rawDescData = protoimpl.X.CompressGZIP(file_services_host_proto_rawDescData)
	})
	return file_services_host_proto_rawDescData
}

var file_services_host_proto_msgTypes = make([]protoimpl.MessageInfo, 6)
var file_services_host_proto_goTypes = []interface{}{
	(*PostGeneralCommentRequest)(nil),  // 0: services.PostGeneralCommentRequest
	(*PostGeneralCommentReply)(nil),    // 1: services.PostGeneralCommentReply
	(*GetPullRequestRequest)(nil),      // 2: services.GetPullRequestRequest
	(*GetPullRequestReply)(nil),        // 3: services.GetPullRequestReply
	(*GetPullRequestFilesRequest)(nil), // 4: services.GetPullRequestFilesRequest
	(*GetPullRequestFilesReply)(nil),   // 5: services.GetPullRequestFilesReply
	(entities.Host)(0),                 // 6: entities.Host
	(*codehost.ReviewComment)(nil),     // 7: codehost.ReviewComment
	(*codehost.PullRequest)(nil),       // 8: codehost.PullRequest
	(*codehost.File)(nil),              // 9: codehost.File
}
var file_services_host_proto_depIdxs = []int32{
	6,  // 0: services.PostGeneralCommentRequest.host:type_name -> entities.Host
	7,  // 1: services.PostGeneralCommentRequest.comment:type_name -> codehost.ReviewComment
	7,  // 2: services.PostGeneralCommentReply.comment:type_name -> codehost.ReviewComment
	6,  // 3: services.GetPullRequestRequest.host:type_name -> entities.Host
	8,  // 4: services.GetPullRequestReply.pull_request:type_name -> codehost.PullRequest
	6,  // 5: services.GetPullRequestFilesRequest.host:type_name -> entities.Host
	9,  // 6: services.GetPullRequestFilesReply.files:type_name -> codehost.File
	0,  // 7: services.Host.PostGeneralComment:input_type -> services.PostGeneralCommentRequest
	2,  // 8: services.Host.GetPullRequest:input_type -> services.GetPullRequestRequest
	4,  // 9: services.Host.GetPullRequestFiles:input_type -> services.GetPullRequestFilesRequest
	1,  // 10: services.Host.PostGeneralComment:output_type -> services.PostGeneralCommentReply
	3,  // 11: services.Host.GetPullRequest:output_type -> services.GetPullRequestReply
	5,  // 12: services.Host.GetPullRequestFiles:output_type -> services.GetPullRequestFilesReply
	10, // [10:13] is the sub-list for method output_type
	7,  // [7:10] is the sub-list for method input_type
	7,  // [7:7] is the sub-list for extension type_name
	7,  // [7:7] is the sub-list for extension extendee
	0,  // [0:7] is the sub-list for field type_name
}

func init() { file_services_host_proto_init() }
func file_services_host_proto_init() {
	if File_services_host_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_services_host_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*PostGeneralCommentRequest); i {
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
		file_services_host_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*PostGeneralCommentReply); i {
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
		file_services_host_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetPullRequestRequest); i {
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
		file_services_host_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetPullRequestReply); i {
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
		file_services_host_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetPullRequestFilesRequest); i {
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
		file_services_host_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetPullRequestFilesReply); i {
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
			RawDescriptor: file_services_host_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   6,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_services_host_proto_goTypes,
		DependencyIndexes: file_services_host_proto_depIdxs,
		MessageInfos:      file_services_host_proto_msgTypes,
	}.Build()
	File_services_host_proto = out.File
	file_services_host_proto_rawDesc = nil
	file_services_host_proto_goTypes = nil
	file_services_host_proto_depIdxs = nil
}