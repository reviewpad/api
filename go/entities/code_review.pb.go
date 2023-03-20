// Copyright 2023 Explore.dev Unipessoal Lda. All Rights Reserved.
// Use of this source code is governed by a license that can be
// found in the LICENSE file.

// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.30.0
// 	protoc        v3.21.12
// source: entities/code_review.proto

package entities

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	timestamppb "google.golang.org/protobuf/types/known/timestamppb"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type CodeReviewStatus int32

const (
	CodeReviewStatus_OPEN   CodeReviewStatus = 0
	CodeReviewStatus_CLOSED CodeReviewStatus = 1
)

// Enum value maps for CodeReviewStatus.
var (
	CodeReviewStatus_name = map[int32]string{
		0: "OPEN",
		1: "CLOSED",
	}
	CodeReviewStatus_value = map[string]int32{
		"OPEN":   0,
		"CLOSED": 1,
	}
)

func (x CodeReviewStatus) Enum() *CodeReviewStatus {
	p := new(CodeReviewStatus)
	*p = x
	return p
}

func (x CodeReviewStatus) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (CodeReviewStatus) Descriptor() protoreflect.EnumDescriptor {
	return file_entities_code_review_proto_enumTypes[0].Descriptor()
}

func (CodeReviewStatus) Type() protoreflect.EnumType {
	return &file_entities_code_review_proto_enumTypes[0]
}

func (x CodeReviewStatus) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use CodeReviewStatus.Descriptor instead.
func (CodeReviewStatus) EnumDescriptor() ([]byte, []int) {
	return file_entities_code_review_proto_rawDescGZIP(), []int{0}
}

type CodeReview struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// The unique identifier of the code review. usually the id of the object we get for the GraphQL API
	Id                 string                 `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Number             int64                  `protobuf:"varint,2,opt,name=number,proto3" json:"number,omitempty"`
	Title              string                 `protobuf:"bytes,3,opt,name=title,proto3" json:"title,omitempty"`
	Description        string                 `protobuf:"bytes,4,opt,name=description,proto3" json:"description,omitempty"`
	Status             CodeReviewStatus       `protobuf:"varint,5,opt,name=status,proto3,enum=entities.CodeReviewStatus" json:"status,omitempty"`
	IsDraft            bool                   `protobuf:"varint,6,opt,name=is_draft,json=isDraft,proto3" json:"is_draft,omitempty"`
	IsMerged           bool                   `protobuf:"varint,7,opt,name=is_merged,json=isMerged,proto3" json:"is_merged,omitempty"`
	IsClosed           bool                   `protobuf:"varint,8,opt,name=is_closed,json=isClosed,proto3" json:"is_closed,omitempty"`
	IsRebaseable       bool                   `protobuf:"varint,9,opt,name=is_rebaseable,json=isRebaseable,proto3" json:"is_rebaseable,omitempty"`
	CommentsCount      int64                  `protobuf:"varint,10,opt,name=comments_count,json=commentsCount,proto3" json:"comments_count,omitempty"`
	CommitsCount       int64                  `protobuf:"varint,11,opt,name=commits_count,json=commitsCount,proto3" json:"commits_count,omitempty"`
	AdditionsCount     int64                  `protobuf:"varint,12,opt,name=additions_count,json=additionsCount,proto3" json:"additions_count,omitempty"`
	DeletionsCount     int64                  `protobuf:"varint,13,opt,name=deletions_count,json=deletionsCount,proto3" json:"deletions_count,omitempty"`
	ChangedFilesCount  int64                  `protobuf:"varint,14,opt,name=changed_files_count,json=changedFilesCount,proto3" json:"changed_files_count,omitempty"`
	Url                string                 `protobuf:"bytes,15,opt,name=url,proto3" json:"url,omitempty"`
	ClosedAt           *timestamppb.Timestamp `protobuf:"bytes,16,opt,name=closed_at,json=closedAt,proto3" json:"closed_at,omitempty"`
	MergedAt           *timestamppb.Timestamp `protobuf:"bytes,17,opt,name=merged_at,json=mergedAt,proto3" json:"merged_at,omitempty"`
	CreatedAt          *timestamppb.Timestamp `protobuf:"bytes,18,opt,name=created_at,json=createdAt,proto3" json:"created_at,omitempty"`
	UpdatedAt          *timestamppb.Timestamp `protobuf:"bytes,19,opt,name=updated_at,json=updatedAt,proto3" json:"updated_at,omitempty"`
	Base               *Branch                `protobuf:"bytes,20,opt,name=base,proto3" json:"base,omitempty"`
	Head               *Branch                `protobuf:"bytes,21,opt,name=head,proto3" json:"head,omitempty"`
	Author             *ExternalUser          `protobuf:"bytes,22,opt,name=author,proto3" json:"author,omitempty"`
	Assignees          []*ExternalUser        `protobuf:"bytes,23,rep,name=assignees,proto3" json:"assignees,omitempty"`
	Labels             []*Label               `protobuf:"bytes,24,rep,name=labels,proto3" json:"labels,omitempty"`
	Milestone          *Milestone             `protobuf:"bytes,25,opt,name=milestone,proto3" json:"milestone,omitempty"`
	RequestedReviewers *RequestedReviewers    `protobuf:"bytes,26,opt,name=requested_reviewers,json=requestedReviewers,proto3" json:"requested_reviewers,omitempty"`
	// A JSON encoded presentation of the raw code review object
	// usually the response we get from the REST API of the provider
	Raw string `protobuf:"bytes,27,opt,name=raw,proto3" json:"raw,omitempty"`
}

func (x *CodeReview) Reset() {
	*x = CodeReview{}
	if protoimpl.UnsafeEnabled {
		mi := &file_entities_code_review_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CodeReview) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CodeReview) ProtoMessage() {}

func (x *CodeReview) ProtoReflect() protoreflect.Message {
	mi := &file_entities_code_review_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CodeReview.ProtoReflect.Descriptor instead.
func (*CodeReview) Descriptor() ([]byte, []int) {
	return file_entities_code_review_proto_rawDescGZIP(), []int{0}
}

func (x *CodeReview) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *CodeReview) GetNumber() int64 {
	if x != nil {
		return x.Number
	}
	return 0
}

func (x *CodeReview) GetTitle() string {
	if x != nil {
		return x.Title
	}
	return ""
}

func (x *CodeReview) GetDescription() string {
	if x != nil {
		return x.Description
	}
	return ""
}

func (x *CodeReview) GetStatus() CodeReviewStatus {
	if x != nil {
		return x.Status
	}
	return CodeReviewStatus_OPEN
}

func (x *CodeReview) GetIsDraft() bool {
	if x != nil {
		return x.IsDraft
	}
	return false
}

func (x *CodeReview) GetIsMerged() bool {
	if x != nil {
		return x.IsMerged
	}
	return false
}

func (x *CodeReview) GetIsClosed() bool {
	if x != nil {
		return x.IsClosed
	}
	return false
}

func (x *CodeReview) GetIsRebaseable() bool {
	if x != nil {
		return x.IsRebaseable
	}
	return false
}

func (x *CodeReview) GetCommentsCount() int64 {
	if x != nil {
		return x.CommentsCount
	}
	return 0
}

func (x *CodeReview) GetCommitsCount() int64 {
	if x != nil {
		return x.CommitsCount
	}
	return 0
}

func (x *CodeReview) GetAdditionsCount() int64 {
	if x != nil {
		return x.AdditionsCount
	}
	return 0
}

func (x *CodeReview) GetDeletionsCount() int64 {
	if x != nil {
		return x.DeletionsCount
	}
	return 0
}

func (x *CodeReview) GetChangedFilesCount() int64 {
	if x != nil {
		return x.ChangedFilesCount
	}
	return 0
}

func (x *CodeReview) GetUrl() string {
	if x != nil {
		return x.Url
	}
	return ""
}

func (x *CodeReview) GetClosedAt() *timestamppb.Timestamp {
	if x != nil {
		return x.ClosedAt
	}
	return nil
}

func (x *CodeReview) GetMergedAt() *timestamppb.Timestamp {
	if x != nil {
		return x.MergedAt
	}
	return nil
}

func (x *CodeReview) GetCreatedAt() *timestamppb.Timestamp {
	if x != nil {
		return x.CreatedAt
	}
	return nil
}

func (x *CodeReview) GetUpdatedAt() *timestamppb.Timestamp {
	if x != nil {
		return x.UpdatedAt
	}
	return nil
}

func (x *CodeReview) GetBase() *Branch {
	if x != nil {
		return x.Base
	}
	return nil
}

func (x *CodeReview) GetHead() *Branch {
	if x != nil {
		return x.Head
	}
	return nil
}

func (x *CodeReview) GetAuthor() *ExternalUser {
	if x != nil {
		return x.Author
	}
	return nil
}

func (x *CodeReview) GetAssignees() []*ExternalUser {
	if x != nil {
		return x.Assignees
	}
	return nil
}

func (x *CodeReview) GetLabels() []*Label {
	if x != nil {
		return x.Labels
	}
	return nil
}

func (x *CodeReview) GetMilestone() *Milestone {
	if x != nil {
		return x.Milestone
	}
	return nil
}

func (x *CodeReview) GetRequestedReviewers() *RequestedReviewers {
	if x != nil {
		return x.RequestedReviewers
	}
	return nil
}

func (x *CodeReview) GetRaw() string {
	if x != nil {
		return x.Raw
	}
	return ""
}

type RequestedReviewers struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Users []*ExternalUser `protobuf:"bytes,1,rep,name=users,proto3" json:"users,omitempty"`
	Teams []*Team         `protobuf:"bytes,2,rep,name=teams,proto3" json:"teams,omitempty"`
}

func (x *RequestedReviewers) Reset() {
	*x = RequestedReviewers{}
	if protoimpl.UnsafeEnabled {
		mi := &file_entities_code_review_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RequestedReviewers) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RequestedReviewers) ProtoMessage() {}

func (x *RequestedReviewers) ProtoReflect() protoreflect.Message {
	mi := &file_entities_code_review_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RequestedReviewers.ProtoReflect.Descriptor instead.
func (*RequestedReviewers) Descriptor() ([]byte, []int) {
	return file_entities_code_review_proto_rawDescGZIP(), []int{1}
}

func (x *RequestedReviewers) GetUsers() []*ExternalUser {
	if x != nil {
		return x.Users
	}
	return nil
}

func (x *RequestedReviewers) GetTeams() []*Team {
	if x != nil {
		return x.Teams
	}
	return nil
}

var File_entities_code_review_proto protoreflect.FileDescriptor

var file_entities_code_review_proto_rawDesc = []byte{
	0x0a, 0x1a, 0x65, 0x6e, 0x74, 0x69, 0x74, 0x69, 0x65, 0x73, 0x2f, 0x63, 0x6f, 0x64, 0x65, 0x5f,
	0x72, 0x65, 0x76, 0x69, 0x65, 0x77, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x08, 0x65, 0x6e,
	0x74, 0x69, 0x74, 0x69, 0x65, 0x73, 0x1a, 0x1c, 0x65, 0x6e, 0x74, 0x69, 0x74, 0x69, 0x65, 0x73,
	0x2f, 0x65, 0x78, 0x74, 0x65, 0x72, 0x6e, 0x61, 0x6c, 0x5f, 0x75, 0x73, 0x65, 0x72, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x14, 0x65, 0x6e, 0x74, 0x69, 0x74, 0x69, 0x65, 0x73, 0x2f, 0x6c,
	0x61, 0x62, 0x65, 0x6c, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x14, 0x65, 0x6e, 0x74, 0x69,
	0x74, 0x69, 0x65, 0x73, 0x2f, 0x74, 0x65, 0x61, 0x6d, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x1a, 0x18, 0x65, 0x6e, 0x74, 0x69, 0x74, 0x69, 0x65, 0x73, 0x2f, 0x6d, 0x69, 0x6c, 0x65, 0x73,
	0x74, 0x6f, 0x6e, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x15, 0x65, 0x6e, 0x74, 0x69,
	0x74, 0x69, 0x65, 0x73, 0x2f, 0x62, 0x72, 0x61, 0x6e, 0x63, 0x68, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x1a, 0x1f, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62,
	0x75, 0x66, 0x2f, 0x74, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x22, 0xd1, 0x08, 0x0a, 0x0a, 0x43, 0x6f, 0x64, 0x65, 0x52, 0x65, 0x76, 0x69, 0x65,
	0x77, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69,
	0x64, 0x12, 0x16, 0x0a, 0x06, 0x6e, 0x75, 0x6d, 0x62, 0x65, 0x72, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x03, 0x52, 0x06, 0x6e, 0x75, 0x6d, 0x62, 0x65, 0x72, 0x12, 0x14, 0x0a, 0x05, 0x74, 0x69, 0x74,
	0x6c, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x74, 0x69, 0x74, 0x6c, 0x65, 0x12,
	0x20, 0x0a, 0x0b, 0x64, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x04,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x64, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f,
	0x6e, 0x12, 0x32, 0x0a, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x18, 0x05, 0x20, 0x01, 0x28,
	0x0e, 0x32, 0x1a, 0x2e, 0x65, 0x6e, 0x74, 0x69, 0x74, 0x69, 0x65, 0x73, 0x2e, 0x43, 0x6f, 0x64,
	0x65, 0x52, 0x65, 0x76, 0x69, 0x65, 0x77, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x52, 0x06, 0x73,
	0x74, 0x61, 0x74, 0x75, 0x73, 0x12, 0x19, 0x0a, 0x08, 0x69, 0x73, 0x5f, 0x64, 0x72, 0x61, 0x66,
	0x74, 0x18, 0x06, 0x20, 0x01, 0x28, 0x08, 0x52, 0x07, 0x69, 0x73, 0x44, 0x72, 0x61, 0x66, 0x74,
	0x12, 0x1b, 0x0a, 0x09, 0x69, 0x73, 0x5f, 0x6d, 0x65, 0x72, 0x67, 0x65, 0x64, 0x18, 0x07, 0x20,
	0x01, 0x28, 0x08, 0x52, 0x08, 0x69, 0x73, 0x4d, 0x65, 0x72, 0x67, 0x65, 0x64, 0x12, 0x1b, 0x0a,
	0x09, 0x69, 0x73, 0x5f, 0x63, 0x6c, 0x6f, 0x73, 0x65, 0x64, 0x18, 0x08, 0x20, 0x01, 0x28, 0x08,
	0x52, 0x08, 0x69, 0x73, 0x43, 0x6c, 0x6f, 0x73, 0x65, 0x64, 0x12, 0x23, 0x0a, 0x0d, 0x69, 0x73,
	0x5f, 0x72, 0x65, 0x62, 0x61, 0x73, 0x65, 0x61, 0x62, 0x6c, 0x65, 0x18, 0x09, 0x20, 0x01, 0x28,
	0x08, 0x52, 0x0c, 0x69, 0x73, 0x52, 0x65, 0x62, 0x61, 0x73, 0x65, 0x61, 0x62, 0x6c, 0x65, 0x12,
	0x25, 0x0a, 0x0e, 0x63, 0x6f, 0x6d, 0x6d, 0x65, 0x6e, 0x74, 0x73, 0x5f, 0x63, 0x6f, 0x75, 0x6e,
	0x74, 0x18, 0x0a, 0x20, 0x01, 0x28, 0x03, 0x52, 0x0d, 0x63, 0x6f, 0x6d, 0x6d, 0x65, 0x6e, 0x74,
	0x73, 0x43, 0x6f, 0x75, 0x6e, 0x74, 0x12, 0x23, 0x0a, 0x0d, 0x63, 0x6f, 0x6d, 0x6d, 0x69, 0x74,
	0x73, 0x5f, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x18, 0x0b, 0x20, 0x01, 0x28, 0x03, 0x52, 0x0c, 0x63,
	0x6f, 0x6d, 0x6d, 0x69, 0x74, 0x73, 0x43, 0x6f, 0x75, 0x6e, 0x74, 0x12, 0x27, 0x0a, 0x0f, 0x61,
	0x64, 0x64, 0x69, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x5f, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x18, 0x0c,
	0x20, 0x01, 0x28, 0x03, 0x52, 0x0e, 0x61, 0x64, 0x64, 0x69, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x43,
	0x6f, 0x75, 0x6e, 0x74, 0x12, 0x27, 0x0a, 0x0f, 0x64, 0x65, 0x6c, 0x65, 0x74, 0x69, 0x6f, 0x6e,
	0x73, 0x5f, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x18, 0x0d, 0x20, 0x01, 0x28, 0x03, 0x52, 0x0e, 0x64,
	0x65, 0x6c, 0x65, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x43, 0x6f, 0x75, 0x6e, 0x74, 0x12, 0x2e, 0x0a,
	0x13, 0x63, 0x68, 0x61, 0x6e, 0x67, 0x65, 0x64, 0x5f, 0x66, 0x69, 0x6c, 0x65, 0x73, 0x5f, 0x63,
	0x6f, 0x75, 0x6e, 0x74, 0x18, 0x0e, 0x20, 0x01, 0x28, 0x03, 0x52, 0x11, 0x63, 0x68, 0x61, 0x6e,
	0x67, 0x65, 0x64, 0x46, 0x69, 0x6c, 0x65, 0x73, 0x43, 0x6f, 0x75, 0x6e, 0x74, 0x12, 0x10, 0x0a,
	0x03, 0x75, 0x72, 0x6c, 0x18, 0x0f, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x75, 0x72, 0x6c, 0x12,
	0x37, 0x0a, 0x09, 0x63, 0x6c, 0x6f, 0x73, 0x65, 0x64, 0x5f, 0x61, 0x74, 0x18, 0x10, 0x20, 0x01,
	0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x52, 0x08,
	0x63, 0x6c, 0x6f, 0x73, 0x65, 0x64, 0x41, 0x74, 0x12, 0x37, 0x0a, 0x09, 0x6d, 0x65, 0x72, 0x67,
	0x65, 0x64, 0x5f, 0x61, 0x74, 0x18, 0x11, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f,
	0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54, 0x69,
	0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x52, 0x08, 0x6d, 0x65, 0x72, 0x67, 0x65, 0x64, 0x41,
	0x74, 0x12, 0x39, 0x0a, 0x0a, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x5f, 0x61, 0x74, 0x18,
	0x12, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d,
	0x70, 0x52, 0x09, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x41, 0x74, 0x12, 0x39, 0x0a, 0x0a,
	0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x64, 0x5f, 0x61, 0x74, 0x18, 0x13, 0x20, 0x01, 0x28, 0x0b,
	0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62,
	0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x52, 0x09, 0x75, 0x70,
	0x64, 0x61, 0x74, 0x65, 0x64, 0x41, 0x74, 0x12, 0x24, 0x0a, 0x04, 0x62, 0x61, 0x73, 0x65, 0x18,
	0x14, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x10, 0x2e, 0x65, 0x6e, 0x74, 0x69, 0x74, 0x69, 0x65, 0x73,
	0x2e, 0x42, 0x72, 0x61, 0x6e, 0x63, 0x68, 0x52, 0x04, 0x62, 0x61, 0x73, 0x65, 0x12, 0x24, 0x0a,
	0x04, 0x68, 0x65, 0x61, 0x64, 0x18, 0x15, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x10, 0x2e, 0x65, 0x6e,
	0x74, 0x69, 0x74, 0x69, 0x65, 0x73, 0x2e, 0x42, 0x72, 0x61, 0x6e, 0x63, 0x68, 0x52, 0x04, 0x68,
	0x65, 0x61, 0x64, 0x12, 0x2e, 0x0a, 0x06, 0x61, 0x75, 0x74, 0x68, 0x6f, 0x72, 0x18, 0x16, 0x20,
	0x01, 0x28, 0x0b, 0x32, 0x16, 0x2e, 0x65, 0x6e, 0x74, 0x69, 0x74, 0x69, 0x65, 0x73, 0x2e, 0x45,
	0x78, 0x74, 0x65, 0x72, 0x6e, 0x61, 0x6c, 0x55, 0x73, 0x65, 0x72, 0x52, 0x06, 0x61, 0x75, 0x74,
	0x68, 0x6f, 0x72, 0x12, 0x34, 0x0a, 0x09, 0x61, 0x73, 0x73, 0x69, 0x67, 0x6e, 0x65, 0x65, 0x73,
	0x18, 0x17, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x16, 0x2e, 0x65, 0x6e, 0x74, 0x69, 0x74, 0x69, 0x65,
	0x73, 0x2e, 0x45, 0x78, 0x74, 0x65, 0x72, 0x6e, 0x61, 0x6c, 0x55, 0x73, 0x65, 0x72, 0x52, 0x09,
	0x61, 0x73, 0x73, 0x69, 0x67, 0x6e, 0x65, 0x65, 0x73, 0x12, 0x27, 0x0a, 0x06, 0x6c, 0x61, 0x62,
	0x65, 0x6c, 0x73, 0x18, 0x18, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x0f, 0x2e, 0x65, 0x6e, 0x74, 0x69,
	0x74, 0x69, 0x65, 0x73, 0x2e, 0x4c, 0x61, 0x62, 0x65, 0x6c, 0x52, 0x06, 0x6c, 0x61, 0x62, 0x65,
	0x6c, 0x73, 0x12, 0x31, 0x0a, 0x09, 0x6d, 0x69, 0x6c, 0x65, 0x73, 0x74, 0x6f, 0x6e, 0x65, 0x18,
	0x19, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x13, 0x2e, 0x65, 0x6e, 0x74, 0x69, 0x74, 0x69, 0x65, 0x73,
	0x2e, 0x4d, 0x69, 0x6c, 0x65, 0x73, 0x74, 0x6f, 0x6e, 0x65, 0x52, 0x09, 0x6d, 0x69, 0x6c, 0x65,
	0x73, 0x74, 0x6f, 0x6e, 0x65, 0x12, 0x4d, 0x0a, 0x13, 0x72, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x65, 0x64, 0x5f, 0x72, 0x65, 0x76, 0x69, 0x65, 0x77, 0x65, 0x72, 0x73, 0x18, 0x1a, 0x20, 0x01,
	0x28, 0x0b, 0x32, 0x1c, 0x2e, 0x65, 0x6e, 0x74, 0x69, 0x74, 0x69, 0x65, 0x73, 0x2e, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x65, 0x64, 0x52, 0x65, 0x76, 0x69, 0x65, 0x77, 0x65, 0x72, 0x73,
	0x52, 0x12, 0x72, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x65, 0x64, 0x52, 0x65, 0x76, 0x69, 0x65,
	0x77, 0x65, 0x72, 0x73, 0x12, 0x10, 0x0a, 0x03, 0x72, 0x61, 0x77, 0x18, 0x1b, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x03, 0x72, 0x61, 0x77, 0x22, 0x68, 0x0a, 0x12, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x65, 0x64, 0x52, 0x65, 0x76, 0x69, 0x65, 0x77, 0x65, 0x72, 0x73, 0x12, 0x2c, 0x0a, 0x05,
	0x75, 0x73, 0x65, 0x72, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x16, 0x2e, 0x65, 0x6e,
	0x74, 0x69, 0x74, 0x69, 0x65, 0x73, 0x2e, 0x45, 0x78, 0x74, 0x65, 0x72, 0x6e, 0x61, 0x6c, 0x55,
	0x73, 0x65, 0x72, 0x52, 0x05, 0x75, 0x73, 0x65, 0x72, 0x73, 0x12, 0x24, 0x0a, 0x05, 0x74, 0x65,
	0x61, 0x6d, 0x73, 0x18, 0x02, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x0e, 0x2e, 0x65, 0x6e, 0x74, 0x69,
	0x74, 0x69, 0x65, 0x73, 0x2e, 0x54, 0x65, 0x61, 0x6d, 0x52, 0x05, 0x74, 0x65, 0x61, 0x6d, 0x73,
	0x2a, 0x28, 0x0a, 0x10, 0x43, 0x6f, 0x64, 0x65, 0x52, 0x65, 0x76, 0x69, 0x65, 0x77, 0x53, 0x74,
	0x61, 0x74, 0x75, 0x73, 0x12, 0x08, 0x0a, 0x04, 0x4f, 0x50, 0x45, 0x4e, 0x10, 0x00, 0x12, 0x0a,
	0x0a, 0x06, 0x43, 0x4c, 0x4f, 0x53, 0x45, 0x44, 0x10, 0x01, 0x42, 0x26, 0x5a, 0x24, 0x67, 0x69,
	0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x72, 0x65, 0x76, 0x69, 0x65, 0x77, 0x70,
	0x61, 0x64, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x67, 0x6f, 0x2f, 0x65, 0x6e, 0x74, 0x69, 0x74, 0x69,
	0x65, 0x73, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_entities_code_review_proto_rawDescOnce sync.Once
	file_entities_code_review_proto_rawDescData = file_entities_code_review_proto_rawDesc
)

func file_entities_code_review_proto_rawDescGZIP() []byte {
	file_entities_code_review_proto_rawDescOnce.Do(func() {
		file_entities_code_review_proto_rawDescData = protoimpl.X.CompressGZIP(file_entities_code_review_proto_rawDescData)
	})
	return file_entities_code_review_proto_rawDescData
}

var file_entities_code_review_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_entities_code_review_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_entities_code_review_proto_goTypes = []interface{}{
	(CodeReviewStatus)(0),         // 0: entities.CodeReviewStatus
	(*CodeReview)(nil),            // 1: entities.CodeReview
	(*RequestedReviewers)(nil),    // 2: entities.RequestedReviewers
	(*timestamppb.Timestamp)(nil), // 3: google.protobuf.Timestamp
	(*Branch)(nil),                // 4: entities.Branch
	(*ExternalUser)(nil),          // 5: entities.ExternalUser
	(*Label)(nil),                 // 6: entities.Label
	(*Milestone)(nil),             // 7: entities.Milestone
	(*Team)(nil),                  // 8: entities.Team
}
var file_entities_code_review_proto_depIdxs = []int32{
	0,  // 0: entities.CodeReview.status:type_name -> entities.CodeReviewStatus
	3,  // 1: entities.CodeReview.closed_at:type_name -> google.protobuf.Timestamp
	3,  // 2: entities.CodeReview.merged_at:type_name -> google.protobuf.Timestamp
	3,  // 3: entities.CodeReview.created_at:type_name -> google.protobuf.Timestamp
	3,  // 4: entities.CodeReview.updated_at:type_name -> google.protobuf.Timestamp
	4,  // 5: entities.CodeReview.base:type_name -> entities.Branch
	4,  // 6: entities.CodeReview.head:type_name -> entities.Branch
	5,  // 7: entities.CodeReview.author:type_name -> entities.ExternalUser
	5,  // 8: entities.CodeReview.assignees:type_name -> entities.ExternalUser
	6,  // 9: entities.CodeReview.labels:type_name -> entities.Label
	7,  // 10: entities.CodeReview.milestone:type_name -> entities.Milestone
	2,  // 11: entities.CodeReview.requested_reviewers:type_name -> entities.RequestedReviewers
	5,  // 12: entities.RequestedReviewers.users:type_name -> entities.ExternalUser
	8,  // 13: entities.RequestedReviewers.teams:type_name -> entities.Team
	14, // [14:14] is the sub-list for method output_type
	14, // [14:14] is the sub-list for method input_type
	14, // [14:14] is the sub-list for extension type_name
	14, // [14:14] is the sub-list for extension extendee
	0,  // [0:14] is the sub-list for field type_name
}

func init() { file_entities_code_review_proto_init() }
func file_entities_code_review_proto_init() {
	if File_entities_code_review_proto != nil {
		return
	}
	file_entities_external_user_proto_init()
	file_entities_label_proto_init()
	file_entities_teams_proto_init()
	file_entities_milestone_proto_init()
	file_entities_branch_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_entities_code_review_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CodeReview); i {
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
		file_entities_code_review_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*RequestedReviewers); i {
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
			RawDescriptor: file_entities_code_review_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_entities_code_review_proto_goTypes,
		DependencyIndexes: file_entities_code_review_proto_depIdxs,
		EnumInfos:         file_entities_code_review_proto_enumTypes,
		MessageInfos:      file_entities_code_review_proto_msgTypes,
	}.Build()
	File_entities_code_review_proto = out.File
	file_entities_code_review_proto_rawDesc = nil
	file_entities_code_review_proto_goTypes = nil
	file_entities_code_review_proto_depIdxs = nil
}
