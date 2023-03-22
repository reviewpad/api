// Copyright 2023 Explore.dev Unipessoal Lda. All Rights Reserved.
// Use of this source code is governed by a license that can be
// found in the LICENSE file.

// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.30.0
// 	protoc        v3.21.12
// source: codehost/comments.proto

package codehost

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

type CommentSeverity int32

const (
	CommentSeverity_LOW    CommentSeverity = 0
	CommentSeverity_MEDIUM CommentSeverity = 1
	CommentSeverity_HIGH   CommentSeverity = 2
)

// Enum value maps for CommentSeverity.
var (
	CommentSeverity_name = map[int32]string{
		0: "LOW",
		1: "MEDIUM",
		2: "HIGH",
	}
	CommentSeverity_value = map[string]int32{
		"LOW":    0,
		"MEDIUM": 1,
		"HIGH":   2,
	}
)

func (x CommentSeverity) Enum() *CommentSeverity {
	p := new(CommentSeverity)
	*p = x
	return p
}

func (x CommentSeverity) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (CommentSeverity) Descriptor() protoreflect.EnumDescriptor {
	return file_codehost_comments_proto_enumTypes[0].Descriptor()
}

func (CommentSeverity) Type() protoreflect.EnumType {
	return &file_codehost_comments_proto_enumTypes[0]
}

func (x CommentSeverity) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use CommentSeverity.Descriptor instead.
func (CommentSeverity) EnumDescriptor() ([]byte, []int) {
	return file_codehost_comments_proto_rawDescGZIP(), []int{0}
}

type CommentType int32

const (
	CommentType_UNKNOWN CommentType = 0
	CommentType_DIFF    CommentType = 1
	CommentType_GENERAL CommentType = 2
)

// Enum value maps for CommentType.
var (
	CommentType_name = map[int32]string{
		0: "UNKNOWN",
		1: "DIFF",
		2: "GENERAL",
	}
	CommentType_value = map[string]int32{
		"UNKNOWN": 0,
		"DIFF":    1,
		"GENERAL": 2,
	}
)

func (x CommentType) Enum() *CommentType {
	p := new(CommentType)
	*p = x
	return p
}

func (x CommentType) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (CommentType) Descriptor() protoreflect.EnumDescriptor {
	return file_codehost_comments_proto_enumTypes[1].Descriptor()
}

func (CommentType) Type() protoreflect.EnumType {
	return &file_codehost_comments_proto_enumTypes[1]
}

func (x CommentType) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use CommentType.Descriptor instead.
func (CommentType) EnumDescriptor() ([]byte, []int) {
	return file_codehost_comments_proto_rawDescGZIP(), []int{1}
}

type CommentLineType int32

const (
	CommentLineType_LINE_ADDED   CommentLineType = 0
	CommentLineType_LINE_REMOVED CommentLineType = 1
	CommentLineType_CONTEXT_LINE CommentLineType = 2
)

// Enum value maps for CommentLineType.
var (
	CommentLineType_name = map[int32]string{
		0: "LINE_ADDED",
		1: "LINE_REMOVED",
		2: "CONTEXT_LINE",
	}
	CommentLineType_value = map[string]int32{
		"LINE_ADDED":   0,
		"LINE_REMOVED": 1,
		"CONTEXT_LINE": 2,
	}
)

func (x CommentLineType) Enum() *CommentLineType {
	p := new(CommentLineType)
	*p = x
	return p
}

func (x CommentLineType) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (CommentLineType) Descriptor() protoreflect.EnumDescriptor {
	return file_codehost_comments_proto_enumTypes[2].Descriptor()
}

func (CommentLineType) Type() protoreflect.EnumType {
	return &file_codehost_comments_proto_enumTypes[2]
}

func (x CommentLineType) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use CommentLineType.Descriptor instead.
func (CommentLineType) EnumDescriptor() ([]byte, []int) {
	return file_codehost_comments_proto_rawDescGZIP(), []int{2}
}

type ReviewComment struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	// In GitLab / Bitbucket, you can have threads in the general comments
	ReplyTo      string                 `protobuf:"bytes,2,opt,name=reply_to,json=replyTo,proto3" json:"reply_to,omitempty"`
	Author       *User                  `protobuf:"bytes,3,opt,name=author,proto3" json:"author,omitempty"`
	Severity     CommentSeverity        `protobuf:"varint,4,opt,name=severity,proto3,enum=codehost.CommentSeverity" json:"severity,omitempty"`
	Body         string                 `protobuf:"bytes,5,opt,name=body,proto3" json:"body,omitempty"`
	CodeReviewId string                 `protobuf:"bytes,6,opt,name=code_review_id,json=codeReviewId,proto3" json:"code_review_id,omitempty"`
	UserId       string                 `protobuf:"bytes,7,opt,name=user_id,json=userId,proto3" json:"user_id,omitempty"`
	UserReviewId string                 `protobuf:"bytes,8,opt,name=user_review_id,json=userReviewId,proto3" json:"user_review_id,omitempty"`
	Location     *ReviewCommentLocation `protobuf:"bytes,9,opt,name=location,proto3" json:"location,omitempty"`
	Resolved     *bool                  `protobuf:"varint,10,opt,name=resolved,proto3,oneof" json:"resolved,omitempty"`
	Type         CommentType            `protobuf:"varint,11,opt,name=type,proto3,enum=codehost.CommentType" json:"type,omitempty"`
	IsOutdated   bool                   `protobuf:"varint,12,opt,name=is_outdated,json=isOutdated,proto3" json:"is_outdated,omitempty"`
	Context      *entities.ChangesBlock `protobuf:"bytes,13,opt,name=context,proto3" json:"context,omitempty"`
	FileInfoDiff *entities.FileInfoDiff `protobuf:"bytes,14,opt,name=file_info_diff,json=fileInfoDiff,proto3" json:"file_info_diff,omitempty"`
	LineType     CommentLineType        `protobuf:"varint,15,opt,name=line_type,json=lineType,proto3,enum=codehost.CommentLineType" json:"line_type,omitempty"`
	// Original information
	OriginalLocation     *ReviewCommentLocation    `protobuf:"bytes,16,opt,name=original_location,json=originalLocation,proto3" json:"original_location,omitempty"`
	ExternalId           string                    `protobuf:"bytes,17,opt,name=external_id,json=externalId,proto3" json:"external_id,omitempty"`
	ExternalNodeId       string                    `protobuf:"bytes,18,opt,name=external_node_id,json=externalNodeId,proto3" json:"external_node_id,omitempty"`
	ExternalReplyTo      string                    `protobuf:"bytes,19,opt,name=external_reply_to,json=externalReplyTo,proto3" json:"external_reply_to,omitempty"`
	ExternalReplyToNode  string                    `protobuf:"bytes,20,opt,name=external_reply_to_node,json=externalReplyToNode,proto3" json:"external_reply_to_node,omitempty"`
	ExternalDiscussionId string                    `protobuf:"bytes,21,opt,name=external_discussion_id,json=externalDiscussionId,proto3" json:"external_discussion_id,omitempty"`
	ExternalCreatedAt    int64                     `protobuf:"varint,22,opt,name=external_created_at,json=externalCreatedAt,proto3" json:"external_created_at,omitempty"`
	ExternalUpdatedAt    int64                     `protobuf:"varint,23,opt,name=external_updated_at,json=externalUpdatedAt,proto3" json:"external_updated_at,omitempty"`
	ExternalReviewId     string                    `protobuf:"bytes,24,opt,name=external_review_id,json=externalReviewId,proto3" json:"external_review_id,omitempty"`
	CreatedAt            int64                     `protobuf:"varint,25,opt,name=created_at,json=createdAt,proto3" json:"created_at,omitempty"`
	UpdatedAt            int64                     `protobuf:"varint,26,opt,name=updated_at,json=updatedAt,proto3" json:"updated_at,omitempty"`
	Diff                 *entities.StaticFileDiffs `protobuf:"bytes,27,opt,name=diff,proto3" json:"diff,omitempty"`
	RelatedCommentsId    []string                  `protobuf:"bytes,28,rep,name=related_comments_id,json=relatedCommentsId,proto3" json:"related_comments_id,omitempty"`
}

func (x *ReviewComment) Reset() {
	*x = ReviewComment{}
	if protoimpl.UnsafeEnabled {
		mi := &file_codehost_comments_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ReviewComment) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ReviewComment) ProtoMessage() {}

func (x *ReviewComment) ProtoReflect() protoreflect.Message {
	mi := &file_codehost_comments_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ReviewComment.ProtoReflect.Descriptor instead.
func (*ReviewComment) Descriptor() ([]byte, []int) {
	return file_codehost_comments_proto_rawDescGZIP(), []int{0}
}

func (x *ReviewComment) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *ReviewComment) GetReplyTo() string {
	if x != nil {
		return x.ReplyTo
	}
	return ""
}

func (x *ReviewComment) GetAuthor() *User {
	if x != nil {
		return x.Author
	}
	return nil
}

func (x *ReviewComment) GetSeverity() CommentSeverity {
	if x != nil {
		return x.Severity
	}
	return CommentSeverity_LOW
}

func (x *ReviewComment) GetBody() string {
	if x != nil {
		return x.Body
	}
	return ""
}

func (x *ReviewComment) GetCodeReviewId() string {
	if x != nil {
		return x.CodeReviewId
	}
	return ""
}

func (x *ReviewComment) GetUserId() string {
	if x != nil {
		return x.UserId
	}
	return ""
}

func (x *ReviewComment) GetUserReviewId() string {
	if x != nil {
		return x.UserReviewId
	}
	return ""
}

func (x *ReviewComment) GetLocation() *ReviewCommentLocation {
	if x != nil {
		return x.Location
	}
	return nil
}

func (x *ReviewComment) GetResolved() bool {
	if x != nil && x.Resolved != nil {
		return *x.Resolved
	}
	return false
}

func (x *ReviewComment) GetType() CommentType {
	if x != nil {
		return x.Type
	}
	return CommentType_UNKNOWN
}

func (x *ReviewComment) GetIsOutdated() bool {
	if x != nil {
		return x.IsOutdated
	}
	return false
}

func (x *ReviewComment) GetContext() *entities.ChangesBlock {
	if x != nil {
		return x.Context
	}
	return nil
}

func (x *ReviewComment) GetFileInfoDiff() *entities.FileInfoDiff {
	if x != nil {
		return x.FileInfoDiff
	}
	return nil
}

func (x *ReviewComment) GetLineType() CommentLineType {
	if x != nil {
		return x.LineType
	}
	return CommentLineType_LINE_ADDED
}

func (x *ReviewComment) GetOriginalLocation() *ReviewCommentLocation {
	if x != nil {
		return x.OriginalLocation
	}
	return nil
}

func (x *ReviewComment) GetExternalId() string {
	if x != nil {
		return x.ExternalId
	}
	return ""
}

func (x *ReviewComment) GetExternalNodeId() string {
	if x != nil {
		return x.ExternalNodeId
	}
	return ""
}

func (x *ReviewComment) GetExternalReplyTo() string {
	if x != nil {
		return x.ExternalReplyTo
	}
	return ""
}

func (x *ReviewComment) GetExternalReplyToNode() string {
	if x != nil {
		return x.ExternalReplyToNode
	}
	return ""
}

func (x *ReviewComment) GetExternalDiscussionId() string {
	if x != nil {
		return x.ExternalDiscussionId
	}
	return ""
}

func (x *ReviewComment) GetExternalCreatedAt() int64 {
	if x != nil {
		return x.ExternalCreatedAt
	}
	return 0
}

func (x *ReviewComment) GetExternalUpdatedAt() int64 {
	if x != nil {
		return x.ExternalUpdatedAt
	}
	return 0
}

func (x *ReviewComment) GetExternalReviewId() string {
	if x != nil {
		return x.ExternalReviewId
	}
	return ""
}

func (x *ReviewComment) GetCreatedAt() int64 {
	if x != nil {
		return x.CreatedAt
	}
	return 0
}

func (x *ReviewComment) GetUpdatedAt() int64 {
	if x != nil {
		return x.UpdatedAt
	}
	return 0
}

func (x *ReviewComment) GetDiff() *entities.StaticFileDiffs {
	if x != nil {
		return x.Diff
	}
	return nil
}

func (x *ReviewComment) GetRelatedCommentsId() []string {
	if x != nil {
		return x.RelatedCommentsId
	}
	return nil
}

type ReviewCommentLocation struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	NewPath      string `protobuf:"bytes,1,opt,name=new_path,json=newPath,proto3" json:"new_path,omitempty"`
	OldPath      string `protobuf:"bytes,2,opt,name=old_path,json=oldPath,proto3" json:"old_path,omitempty"`
	BaseCommitId string `protobuf:"bytes,3,opt,name=base_commit_id,json=baseCommitId,proto3" json:"base_commit_id,omitempty"`
	HeadCommitId string `protobuf:"bytes,4,opt,name=head_commit_id,json=headCommitId,proto3" json:"head_commit_id,omitempty"`
	NewLine      int32  `protobuf:"varint,5,opt,name=new_line,json=newLine,proto3" json:"new_line,omitempty"`
	OldLine      int32  `protobuf:"varint,6,opt,name=old_line,json=oldLine,proto3" json:"old_line,omitempty"`
}

func (x *ReviewCommentLocation) Reset() {
	*x = ReviewCommentLocation{}
	if protoimpl.UnsafeEnabled {
		mi := &file_codehost_comments_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ReviewCommentLocation) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ReviewCommentLocation) ProtoMessage() {}

func (x *ReviewCommentLocation) ProtoReflect() protoreflect.Message {
	mi := &file_codehost_comments_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ReviewCommentLocation.ProtoReflect.Descriptor instead.
func (*ReviewCommentLocation) Descriptor() ([]byte, []int) {
	return file_codehost_comments_proto_rawDescGZIP(), []int{1}
}

func (x *ReviewCommentLocation) GetNewPath() string {
	if x != nil {
		return x.NewPath
	}
	return ""
}

func (x *ReviewCommentLocation) GetOldPath() string {
	if x != nil {
		return x.OldPath
	}
	return ""
}

func (x *ReviewCommentLocation) GetBaseCommitId() string {
	if x != nil {
		return x.BaseCommitId
	}
	return ""
}

func (x *ReviewCommentLocation) GetHeadCommitId() string {
	if x != nil {
		return x.HeadCommitId
	}
	return ""
}

func (x *ReviewCommentLocation) GetNewLine() int32 {
	if x != nil {
		return x.NewLine
	}
	return 0
}

func (x *ReviewCommentLocation) GetOldLine() int32 {
	if x != nil {
		return x.OldLine
	}
	return 0
}

type ReviewFileComments struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Comments []*ReviewComment `protobuf:"bytes,1,rep,name=comments,proto3" json:"comments,omitempty"`
}

func (x *ReviewFileComments) Reset() {
	*x = ReviewFileComments{}
	if protoimpl.UnsafeEnabled {
		mi := &file_codehost_comments_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ReviewFileComments) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ReviewFileComments) ProtoMessage() {}

func (x *ReviewFileComments) ProtoReflect() protoreflect.Message {
	mi := &file_codehost_comments_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ReviewFileComments.ProtoReflect.Descriptor instead.
func (*ReviewFileComments) Descriptor() ([]byte, []int) {
	return file_codehost_comments_proto_rawDescGZIP(), []int{2}
}

func (x *ReviewFileComments) GetComments() []*ReviewComment {
	if x != nil {
		return x.Comments
	}
	return nil
}

type ReviewDiscussionCountBySeverity struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	BySeverity map[string]*ReviewDiscussionCount `protobuf:"bytes,1,rep,name=by_severity,json=bySeverity,proto3" json:"by_severity,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
}

func (x *ReviewDiscussionCountBySeverity) Reset() {
	*x = ReviewDiscussionCountBySeverity{}
	if protoimpl.UnsafeEnabled {
		mi := &file_codehost_comments_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ReviewDiscussionCountBySeverity) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ReviewDiscussionCountBySeverity) ProtoMessage() {}

func (x *ReviewDiscussionCountBySeverity) ProtoReflect() protoreflect.Message {
	mi := &file_codehost_comments_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ReviewDiscussionCountBySeverity.ProtoReflect.Descriptor instead.
func (*ReviewDiscussionCountBySeverity) Descriptor() ([]byte, []int) {
	return file_codehost_comments_proto_rawDescGZIP(), []int{3}
}

func (x *ReviewDiscussionCountBySeverity) GetBySeverity() map[string]*ReviewDiscussionCount {
	if x != nil {
		return x.BySeverity
	}
	return nil
}

type ReviewDiscussionCount struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Total    int32 `protobuf:"varint,1,opt,name=total,proto3" json:"total,omitempty"`
	Resolved int32 `protobuf:"varint,2,opt,name=resolved,proto3" json:"resolved,omitempty"`
}

func (x *ReviewDiscussionCount) Reset() {
	*x = ReviewDiscussionCount{}
	if protoimpl.UnsafeEnabled {
		mi := &file_codehost_comments_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ReviewDiscussionCount) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ReviewDiscussionCount) ProtoMessage() {}

func (x *ReviewDiscussionCount) ProtoReflect() protoreflect.Message {
	mi := &file_codehost_comments_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ReviewDiscussionCount.ProtoReflect.Descriptor instead.
func (*ReviewDiscussionCount) Descriptor() ([]byte, []int) {
	return file_codehost_comments_proto_rawDescGZIP(), []int{4}
}

func (x *ReviewDiscussionCount) GetTotal() int32 {
	if x != nil {
		return x.Total
	}
	return 0
}

func (x *ReviewDiscussionCount) GetResolved() int32 {
	if x != nil {
		return x.Resolved
	}
	return 0
}

var File_codehost_comments_proto protoreflect.FileDescriptor

var file_codehost_comments_proto_rawDesc = []byte{
	0x0a, 0x17, 0x63, 0x6f, 0x64, 0x65, 0x68, 0x6f, 0x73, 0x74, 0x2f, 0x63, 0x6f, 0x6d, 0x6d, 0x65,
	0x6e, 0x74, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x08, 0x63, 0x6f, 0x64, 0x65, 0x68,
	0x6f, 0x73, 0x74, 0x1a, 0x13, 0x63, 0x6f, 0x64, 0x65, 0x68, 0x6f, 0x73, 0x74, 0x2f, 0x75, 0x73,
	0x65, 0x72, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x13, 0x65, 0x6e, 0x74, 0x69, 0x74, 0x69,
	0x65, 0x73, 0x2f, 0x66, 0x69, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x17, 0x65,
	0x6e, 0x74, 0x69, 0x74, 0x69, 0x65, 0x73, 0x2f, 0x67, 0x69, 0x74, 0x5f, 0x64, 0x69, 0x66, 0x66,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1c, 0x65, 0x6e, 0x74, 0x69, 0x74, 0x69, 0x65, 0x73,
	0x2f, 0x73, 0x65, 0x6d, 0x61, 0x6e, 0x74, 0x69, 0x63, 0x5f, 0x64, 0x69, 0x66, 0x66, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x22, 0xcc, 0x09, 0x0a, 0x0d, 0x52, 0x65, 0x76, 0x69, 0x65, 0x77, 0x43,
	0x6f, 0x6d, 0x6d, 0x65, 0x6e, 0x74, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x12, 0x19, 0x0a, 0x08, 0x72, 0x65, 0x70, 0x6c, 0x79, 0x5f,
	0x74, 0x6f, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x72, 0x65, 0x70, 0x6c, 0x79, 0x54,
	0x6f, 0x12, 0x26, 0x0a, 0x06, 0x61, 0x75, 0x74, 0x68, 0x6f, 0x72, 0x18, 0x03, 0x20, 0x01, 0x28,
	0x0b, 0x32, 0x0e, 0x2e, 0x63, 0x6f, 0x64, 0x65, 0x68, 0x6f, 0x73, 0x74, 0x2e, 0x55, 0x73, 0x65,
	0x72, 0x52, 0x06, 0x61, 0x75, 0x74, 0x68, 0x6f, 0x72, 0x12, 0x35, 0x0a, 0x08, 0x73, 0x65, 0x76,
	0x65, 0x72, 0x69, 0x74, 0x79, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x19, 0x2e, 0x63, 0x6f,
	0x64, 0x65, 0x68, 0x6f, 0x73, 0x74, 0x2e, 0x43, 0x6f, 0x6d, 0x6d, 0x65, 0x6e, 0x74, 0x53, 0x65,
	0x76, 0x65, 0x72, 0x69, 0x74, 0x79, 0x52, 0x08, 0x73, 0x65, 0x76, 0x65, 0x72, 0x69, 0x74, 0x79,
	0x12, 0x12, 0x0a, 0x04, 0x62, 0x6f, 0x64, 0x79, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04,
	0x62, 0x6f, 0x64, 0x79, 0x12, 0x24, 0x0a, 0x0e, 0x63, 0x6f, 0x64, 0x65, 0x5f, 0x72, 0x65, 0x76,
	0x69, 0x65, 0x77, 0x5f, 0x69, 0x64, 0x18, 0x06, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0c, 0x63, 0x6f,
	0x64, 0x65, 0x52, 0x65, 0x76, 0x69, 0x65, 0x77, 0x49, 0x64, 0x12, 0x17, 0x0a, 0x07, 0x75, 0x73,
	0x65, 0x72, 0x5f, 0x69, 0x64, 0x18, 0x07, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x75, 0x73, 0x65,
	0x72, 0x49, 0x64, 0x12, 0x24, 0x0a, 0x0e, 0x75, 0x73, 0x65, 0x72, 0x5f, 0x72, 0x65, 0x76, 0x69,
	0x65, 0x77, 0x5f, 0x69, 0x64, 0x18, 0x08, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0c, 0x75, 0x73, 0x65,
	0x72, 0x52, 0x65, 0x76, 0x69, 0x65, 0x77, 0x49, 0x64, 0x12, 0x3b, 0x0a, 0x08, 0x6c, 0x6f, 0x63,
	0x61, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x09, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1f, 0x2e, 0x63, 0x6f,
	0x64, 0x65, 0x68, 0x6f, 0x73, 0x74, 0x2e, 0x52, 0x65, 0x76, 0x69, 0x65, 0x77, 0x43, 0x6f, 0x6d,
	0x6d, 0x65, 0x6e, 0x74, 0x4c, 0x6f, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x08, 0x6c, 0x6f,
	0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x1f, 0x0a, 0x08, 0x72, 0x65, 0x73, 0x6f, 0x6c, 0x76,
	0x65, 0x64, 0x18, 0x0a, 0x20, 0x01, 0x28, 0x08, 0x48, 0x00, 0x52, 0x08, 0x72, 0x65, 0x73, 0x6f,
	0x6c, 0x76, 0x65, 0x64, 0x88, 0x01, 0x01, 0x12, 0x29, 0x0a, 0x04, 0x74, 0x79, 0x70, 0x65, 0x18,
	0x0b, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x15, 0x2e, 0x63, 0x6f, 0x64, 0x65, 0x68, 0x6f, 0x73, 0x74,
	0x2e, 0x43, 0x6f, 0x6d, 0x6d, 0x65, 0x6e, 0x74, 0x54, 0x79, 0x70, 0x65, 0x52, 0x04, 0x74, 0x79,
	0x70, 0x65, 0x12, 0x1f, 0x0a, 0x0b, 0x69, 0x73, 0x5f, 0x6f, 0x75, 0x74, 0x64, 0x61, 0x74, 0x65,
	0x64, 0x18, 0x0c, 0x20, 0x01, 0x28, 0x08, 0x52, 0x0a, 0x69, 0x73, 0x4f, 0x75, 0x74, 0x64, 0x61,
	0x74, 0x65, 0x64, 0x12, 0x30, 0x0a, 0x07, 0x63, 0x6f, 0x6e, 0x74, 0x65, 0x78, 0x74, 0x18, 0x0d,
	0x20, 0x01, 0x28, 0x0b, 0x32, 0x16, 0x2e, 0x65, 0x6e, 0x74, 0x69, 0x74, 0x69, 0x65, 0x73, 0x2e,
	0x43, 0x68, 0x61, 0x6e, 0x67, 0x65, 0x73, 0x42, 0x6c, 0x6f, 0x63, 0x6b, 0x52, 0x07, 0x63, 0x6f,
	0x6e, 0x74, 0x65, 0x78, 0x74, 0x12, 0x3c, 0x0a, 0x0e, 0x66, 0x69, 0x6c, 0x65, 0x5f, 0x69, 0x6e,
	0x66, 0x6f, 0x5f, 0x64, 0x69, 0x66, 0x66, 0x18, 0x0e, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x16, 0x2e,
	0x65, 0x6e, 0x74, 0x69, 0x74, 0x69, 0x65, 0x73, 0x2e, 0x46, 0x69, 0x6c, 0x65, 0x49, 0x6e, 0x66,
	0x6f, 0x44, 0x69, 0x66, 0x66, 0x52, 0x0c, 0x66, 0x69, 0x6c, 0x65, 0x49, 0x6e, 0x66, 0x6f, 0x44,
	0x69, 0x66, 0x66, 0x12, 0x36, 0x0a, 0x09, 0x6c, 0x69, 0x6e, 0x65, 0x5f, 0x74, 0x79, 0x70, 0x65,
	0x18, 0x0f, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x19, 0x2e, 0x63, 0x6f, 0x64, 0x65, 0x68, 0x6f, 0x73,
	0x74, 0x2e, 0x43, 0x6f, 0x6d, 0x6d, 0x65, 0x6e, 0x74, 0x4c, 0x69, 0x6e, 0x65, 0x54, 0x79, 0x70,
	0x65, 0x52, 0x08, 0x6c, 0x69, 0x6e, 0x65, 0x54, 0x79, 0x70, 0x65, 0x12, 0x4c, 0x0a, 0x11, 0x6f,
	0x72, 0x69, 0x67, 0x69, 0x6e, 0x61, 0x6c, 0x5f, 0x6c, 0x6f, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e,
	0x18, 0x10, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1f, 0x2e, 0x63, 0x6f, 0x64, 0x65, 0x68, 0x6f, 0x73,
	0x74, 0x2e, 0x52, 0x65, 0x76, 0x69, 0x65, 0x77, 0x43, 0x6f, 0x6d, 0x6d, 0x65, 0x6e, 0x74, 0x4c,
	0x6f, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x10, 0x6f, 0x72, 0x69, 0x67, 0x69, 0x6e, 0x61,
	0x6c, 0x4c, 0x6f, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x1f, 0x0a, 0x0b, 0x65, 0x78, 0x74,
	0x65, 0x72, 0x6e, 0x61, 0x6c, 0x5f, 0x69, 0x64, 0x18, 0x11, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a,
	0x65, 0x78, 0x74, 0x65, 0x72, 0x6e, 0x61, 0x6c, 0x49, 0x64, 0x12, 0x28, 0x0a, 0x10, 0x65, 0x78,
	0x74, 0x65, 0x72, 0x6e, 0x61, 0x6c, 0x5f, 0x6e, 0x6f, 0x64, 0x65, 0x5f, 0x69, 0x64, 0x18, 0x12,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x0e, 0x65, 0x78, 0x74, 0x65, 0x72, 0x6e, 0x61, 0x6c, 0x4e, 0x6f,
	0x64, 0x65, 0x49, 0x64, 0x12, 0x2a, 0x0a, 0x11, 0x65, 0x78, 0x74, 0x65, 0x72, 0x6e, 0x61, 0x6c,
	0x5f, 0x72, 0x65, 0x70, 0x6c, 0x79, 0x5f, 0x74, 0x6f, 0x18, 0x13, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x0f, 0x65, 0x78, 0x74, 0x65, 0x72, 0x6e, 0x61, 0x6c, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x54, 0x6f,
	0x12, 0x33, 0x0a, 0x16, 0x65, 0x78, 0x74, 0x65, 0x72, 0x6e, 0x61, 0x6c, 0x5f, 0x72, 0x65, 0x70,
	0x6c, 0x79, 0x5f, 0x74, 0x6f, 0x5f, 0x6e, 0x6f, 0x64, 0x65, 0x18, 0x14, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x13, 0x65, 0x78, 0x74, 0x65, 0x72, 0x6e, 0x61, 0x6c, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x54,
	0x6f, 0x4e, 0x6f, 0x64, 0x65, 0x12, 0x34, 0x0a, 0x16, 0x65, 0x78, 0x74, 0x65, 0x72, 0x6e, 0x61,
	0x6c, 0x5f, 0x64, 0x69, 0x73, 0x63, 0x75, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x5f, 0x69, 0x64, 0x18,
	0x15, 0x20, 0x01, 0x28, 0x09, 0x52, 0x14, 0x65, 0x78, 0x74, 0x65, 0x72, 0x6e, 0x61, 0x6c, 0x44,
	0x69, 0x73, 0x63, 0x75, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x49, 0x64, 0x12, 0x2e, 0x0a, 0x13, 0x65,
	0x78, 0x74, 0x65, 0x72, 0x6e, 0x61, 0x6c, 0x5f, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x5f,
	0x61, 0x74, 0x18, 0x16, 0x20, 0x01, 0x28, 0x03, 0x52, 0x11, 0x65, 0x78, 0x74, 0x65, 0x72, 0x6e,
	0x61, 0x6c, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x41, 0x74, 0x12, 0x2e, 0x0a, 0x13, 0x65,
	0x78, 0x74, 0x65, 0x72, 0x6e, 0x61, 0x6c, 0x5f, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x64, 0x5f,
	0x61, 0x74, 0x18, 0x17, 0x20, 0x01, 0x28, 0x03, 0x52, 0x11, 0x65, 0x78, 0x74, 0x65, 0x72, 0x6e,
	0x61, 0x6c, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x64, 0x41, 0x74, 0x12, 0x2c, 0x0a, 0x12, 0x65,
	0x78, 0x74, 0x65, 0x72, 0x6e, 0x61, 0x6c, 0x5f, 0x72, 0x65, 0x76, 0x69, 0x65, 0x77, 0x5f, 0x69,
	0x64, 0x18, 0x18, 0x20, 0x01, 0x28, 0x09, 0x52, 0x10, 0x65, 0x78, 0x74, 0x65, 0x72, 0x6e, 0x61,
	0x6c, 0x52, 0x65, 0x76, 0x69, 0x65, 0x77, 0x49, 0x64, 0x12, 0x1d, 0x0a, 0x0a, 0x63, 0x72, 0x65,
	0x61, 0x74, 0x65, 0x64, 0x5f, 0x61, 0x74, 0x18, 0x19, 0x20, 0x01, 0x28, 0x03, 0x52, 0x09, 0x63,
	0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x41, 0x74, 0x12, 0x1d, 0x0a, 0x0a, 0x75, 0x70, 0x64, 0x61,
	0x74, 0x65, 0x64, 0x5f, 0x61, 0x74, 0x18, 0x1a, 0x20, 0x01, 0x28, 0x03, 0x52, 0x09, 0x75, 0x70,
	0x64, 0x61, 0x74, 0x65, 0x64, 0x41, 0x74, 0x12, 0x2d, 0x0a, 0x04, 0x64, 0x69, 0x66, 0x66, 0x18,
	0x1b, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x19, 0x2e, 0x65, 0x6e, 0x74, 0x69, 0x74, 0x69, 0x65, 0x73,
	0x2e, 0x53, 0x74, 0x61, 0x74, 0x69, 0x63, 0x46, 0x69, 0x6c, 0x65, 0x44, 0x69, 0x66, 0x66, 0x73,
	0x52, 0x04, 0x64, 0x69, 0x66, 0x66, 0x12, 0x2e, 0x0a, 0x13, 0x72, 0x65, 0x6c, 0x61, 0x74, 0x65,
	0x64, 0x5f, 0x63, 0x6f, 0x6d, 0x6d, 0x65, 0x6e, 0x74, 0x73, 0x5f, 0x69, 0x64, 0x18, 0x1c, 0x20,
	0x03, 0x28, 0x09, 0x52, 0x11, 0x72, 0x65, 0x6c, 0x61, 0x74, 0x65, 0x64, 0x43, 0x6f, 0x6d, 0x6d,
	0x65, 0x6e, 0x74, 0x73, 0x49, 0x64, 0x42, 0x0b, 0x0a, 0x09, 0x5f, 0x72, 0x65, 0x73, 0x6f, 0x6c,
	0x76, 0x65, 0x64, 0x22, 0xcf, 0x01, 0x0a, 0x15, 0x52, 0x65, 0x76, 0x69, 0x65, 0x77, 0x43, 0x6f,
	0x6d, 0x6d, 0x65, 0x6e, 0x74, 0x4c, 0x6f, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x19, 0x0a,
	0x08, 0x6e, 0x65, 0x77, 0x5f, 0x70, 0x61, 0x74, 0x68, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x07, 0x6e, 0x65, 0x77, 0x50, 0x61, 0x74, 0x68, 0x12, 0x19, 0x0a, 0x08, 0x6f, 0x6c, 0x64, 0x5f,
	0x70, 0x61, 0x74, 0x68, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x6f, 0x6c, 0x64, 0x50,
	0x61, 0x74, 0x68, 0x12, 0x24, 0x0a, 0x0e, 0x62, 0x61, 0x73, 0x65, 0x5f, 0x63, 0x6f, 0x6d, 0x6d,
	0x69, 0x74, 0x5f, 0x69, 0x64, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0c, 0x62, 0x61, 0x73,
	0x65, 0x43, 0x6f, 0x6d, 0x6d, 0x69, 0x74, 0x49, 0x64, 0x12, 0x24, 0x0a, 0x0e, 0x68, 0x65, 0x61,
	0x64, 0x5f, 0x63, 0x6f, 0x6d, 0x6d, 0x69, 0x74, 0x5f, 0x69, 0x64, 0x18, 0x04, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x0c, 0x68, 0x65, 0x61, 0x64, 0x43, 0x6f, 0x6d, 0x6d, 0x69, 0x74, 0x49, 0x64, 0x12,
	0x19, 0x0a, 0x08, 0x6e, 0x65, 0x77, 0x5f, 0x6c, 0x69, 0x6e, 0x65, 0x18, 0x05, 0x20, 0x01, 0x28,
	0x05, 0x52, 0x07, 0x6e, 0x65, 0x77, 0x4c, 0x69, 0x6e, 0x65, 0x12, 0x19, 0x0a, 0x08, 0x6f, 0x6c,
	0x64, 0x5f, 0x6c, 0x69, 0x6e, 0x65, 0x18, 0x06, 0x20, 0x01, 0x28, 0x05, 0x52, 0x07, 0x6f, 0x6c,
	0x64, 0x4c, 0x69, 0x6e, 0x65, 0x22, 0x49, 0x0a, 0x12, 0x52, 0x65, 0x76, 0x69, 0x65, 0x77, 0x46,
	0x69, 0x6c, 0x65, 0x43, 0x6f, 0x6d, 0x6d, 0x65, 0x6e, 0x74, 0x73, 0x12, 0x33, 0x0a, 0x08, 0x63,
	0x6f, 0x6d, 0x6d, 0x65, 0x6e, 0x74, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x17, 0x2e,
	0x63, 0x6f, 0x64, 0x65, 0x68, 0x6f, 0x73, 0x74, 0x2e, 0x52, 0x65, 0x76, 0x69, 0x65, 0x77, 0x43,
	0x6f, 0x6d, 0x6d, 0x65, 0x6e, 0x74, 0x52, 0x08, 0x63, 0x6f, 0x6d, 0x6d, 0x65, 0x6e, 0x74, 0x73,
	0x22, 0xdd, 0x01, 0x0a, 0x1f, 0x52, 0x65, 0x76, 0x69, 0x65, 0x77, 0x44, 0x69, 0x73, 0x63, 0x75,
	0x73, 0x73, 0x69, 0x6f, 0x6e, 0x43, 0x6f, 0x75, 0x6e, 0x74, 0x42, 0x79, 0x53, 0x65, 0x76, 0x65,
	0x72, 0x69, 0x74, 0x79, 0x12, 0x5a, 0x0a, 0x0b, 0x62, 0x79, 0x5f, 0x73, 0x65, 0x76, 0x65, 0x72,
	0x69, 0x74, 0x79, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x39, 0x2e, 0x63, 0x6f, 0x64, 0x65,
	0x68, 0x6f, 0x73, 0x74, 0x2e, 0x52, 0x65, 0x76, 0x69, 0x65, 0x77, 0x44, 0x69, 0x73, 0x63, 0x75,
	0x73, 0x73, 0x69, 0x6f, 0x6e, 0x43, 0x6f, 0x75, 0x6e, 0x74, 0x42, 0x79, 0x53, 0x65, 0x76, 0x65,
	0x72, 0x69, 0x74, 0x79, 0x2e, 0x42, 0x79, 0x53, 0x65, 0x76, 0x65, 0x72, 0x69, 0x74, 0x79, 0x45,
	0x6e, 0x74, 0x72, 0x79, 0x52, 0x0a, 0x62, 0x79, 0x53, 0x65, 0x76, 0x65, 0x72, 0x69, 0x74, 0x79,
	0x1a, 0x5e, 0x0a, 0x0f, 0x42, 0x79, 0x53, 0x65, 0x76, 0x65, 0x72, 0x69, 0x74, 0x79, 0x45, 0x6e,
	0x74, 0x72, 0x79, 0x12, 0x10, 0x0a, 0x03, 0x6b, 0x65, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x03, 0x6b, 0x65, 0x79, 0x12, 0x35, 0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x0b, 0x32, 0x1f, 0x2e, 0x63, 0x6f, 0x64, 0x65, 0x68, 0x6f, 0x73, 0x74, 0x2e,
	0x52, 0x65, 0x76, 0x69, 0x65, 0x77, 0x44, 0x69, 0x73, 0x63, 0x75, 0x73, 0x73, 0x69, 0x6f, 0x6e,
	0x43, 0x6f, 0x75, 0x6e, 0x74, 0x52, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x3a, 0x02, 0x38, 0x01,
	0x22, 0x49, 0x0a, 0x15, 0x52, 0x65, 0x76, 0x69, 0x65, 0x77, 0x44, 0x69, 0x73, 0x63, 0x75, 0x73,
	0x73, 0x69, 0x6f, 0x6e, 0x43, 0x6f, 0x75, 0x6e, 0x74, 0x12, 0x14, 0x0a, 0x05, 0x74, 0x6f, 0x74,
	0x61, 0x6c, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x05, 0x74, 0x6f, 0x74, 0x61, 0x6c, 0x12,
	0x1a, 0x0a, 0x08, 0x72, 0x65, 0x73, 0x6f, 0x6c, 0x76, 0x65, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x05, 0x52, 0x08, 0x72, 0x65, 0x73, 0x6f, 0x6c, 0x76, 0x65, 0x64, 0x2a, 0x30, 0x0a, 0x0f, 0x43,
	0x6f, 0x6d, 0x6d, 0x65, 0x6e, 0x74, 0x53, 0x65, 0x76, 0x65, 0x72, 0x69, 0x74, 0x79, 0x12, 0x07,
	0x0a, 0x03, 0x4c, 0x4f, 0x57, 0x10, 0x00, 0x12, 0x0a, 0x0a, 0x06, 0x4d, 0x45, 0x44, 0x49, 0x55,
	0x4d, 0x10, 0x01, 0x12, 0x08, 0x0a, 0x04, 0x48, 0x49, 0x47, 0x48, 0x10, 0x02, 0x2a, 0x31, 0x0a,
	0x0b, 0x43, 0x6f, 0x6d, 0x6d, 0x65, 0x6e, 0x74, 0x54, 0x79, 0x70, 0x65, 0x12, 0x0b, 0x0a, 0x07,
	0x55, 0x4e, 0x4b, 0x4e, 0x4f, 0x57, 0x4e, 0x10, 0x00, 0x12, 0x08, 0x0a, 0x04, 0x44, 0x49, 0x46,
	0x46, 0x10, 0x01, 0x12, 0x0b, 0x0a, 0x07, 0x47, 0x45, 0x4e, 0x45, 0x52, 0x41, 0x4c, 0x10, 0x02,
	0x2a, 0x45, 0x0a, 0x0f, 0x43, 0x6f, 0x6d, 0x6d, 0x65, 0x6e, 0x74, 0x4c, 0x69, 0x6e, 0x65, 0x54,
	0x79, 0x70, 0x65, 0x12, 0x0e, 0x0a, 0x0a, 0x4c, 0x49, 0x4e, 0x45, 0x5f, 0x41, 0x44, 0x44, 0x45,
	0x44, 0x10, 0x00, 0x12, 0x10, 0x0a, 0x0c, 0x4c, 0x49, 0x4e, 0x45, 0x5f, 0x52, 0x45, 0x4d, 0x4f,
	0x56, 0x45, 0x44, 0x10, 0x01, 0x12, 0x10, 0x0a, 0x0c, 0x43, 0x4f, 0x4e, 0x54, 0x45, 0x58, 0x54,
	0x5f, 0x4c, 0x49, 0x4e, 0x45, 0x10, 0x02, 0x42, 0x26, 0x5a, 0x24, 0x67, 0x69, 0x74, 0x68, 0x75,
	0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x72, 0x65, 0x76, 0x69, 0x65, 0x77, 0x70, 0x61, 0x64, 0x2f,
	0x61, 0x70, 0x69, 0x2f, 0x67, 0x6f, 0x2f, 0x63, 0x6f, 0x64, 0x65, 0x68, 0x6f, 0x73, 0x74, 0x62,
	0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_codehost_comments_proto_rawDescOnce sync.Once
	file_codehost_comments_proto_rawDescData = file_codehost_comments_proto_rawDesc
)

func file_codehost_comments_proto_rawDescGZIP() []byte {
	file_codehost_comments_proto_rawDescOnce.Do(func() {
		file_codehost_comments_proto_rawDescData = protoimpl.X.CompressGZIP(file_codehost_comments_proto_rawDescData)
	})
	return file_codehost_comments_proto_rawDescData
}

var file_codehost_comments_proto_enumTypes = make([]protoimpl.EnumInfo, 3)
var file_codehost_comments_proto_msgTypes = make([]protoimpl.MessageInfo, 6)
var file_codehost_comments_proto_goTypes = []interface{}{
	(CommentSeverity)(0),                    // 0: codehost.CommentSeverity
	(CommentType)(0),                        // 1: codehost.CommentType
	(CommentLineType)(0),                    // 2: codehost.CommentLineType
	(*ReviewComment)(nil),                   // 3: codehost.ReviewComment
	(*ReviewCommentLocation)(nil),           // 4: codehost.ReviewCommentLocation
	(*ReviewFileComments)(nil),              // 5: codehost.ReviewFileComments
	(*ReviewDiscussionCountBySeverity)(nil), // 6: codehost.ReviewDiscussionCountBySeverity
	(*ReviewDiscussionCount)(nil),           // 7: codehost.ReviewDiscussionCount
	nil,                                     // 8: codehost.ReviewDiscussionCountBySeverity.BySeverityEntry
	(*User)(nil),                            // 9: codehost.User
	(*entities.ChangesBlock)(nil),           // 10: entities.ChangesBlock
	(*entities.FileInfoDiff)(nil),           // 11: entities.FileInfoDiff
	(*entities.StaticFileDiffs)(nil),        // 12: entities.StaticFileDiffs
}
var file_codehost_comments_proto_depIdxs = []int32{
	9,  // 0: codehost.ReviewComment.author:type_name -> codehost.User
	0,  // 1: codehost.ReviewComment.severity:type_name -> codehost.CommentSeverity
	4,  // 2: codehost.ReviewComment.location:type_name -> codehost.ReviewCommentLocation
	1,  // 3: codehost.ReviewComment.type:type_name -> codehost.CommentType
	10, // 4: codehost.ReviewComment.context:type_name -> entities.ChangesBlock
	11, // 5: codehost.ReviewComment.file_info_diff:type_name -> entities.FileInfoDiff
	2,  // 6: codehost.ReviewComment.line_type:type_name -> codehost.CommentLineType
	4,  // 7: codehost.ReviewComment.original_location:type_name -> codehost.ReviewCommentLocation
	12, // 8: codehost.ReviewComment.diff:type_name -> entities.StaticFileDiffs
	3,  // 9: codehost.ReviewFileComments.comments:type_name -> codehost.ReviewComment
	8,  // 10: codehost.ReviewDiscussionCountBySeverity.by_severity:type_name -> codehost.ReviewDiscussionCountBySeverity.BySeverityEntry
	7,  // 11: codehost.ReviewDiscussionCountBySeverity.BySeverityEntry.value:type_name -> codehost.ReviewDiscussionCount
	12, // [12:12] is the sub-list for method output_type
	12, // [12:12] is the sub-list for method input_type
	12, // [12:12] is the sub-list for extension type_name
	12, // [12:12] is the sub-list for extension extendee
	0,  // [0:12] is the sub-list for field type_name
}

func init() { file_codehost_comments_proto_init() }
func file_codehost_comments_proto_init() {
	if File_codehost_comments_proto != nil {
		return
	}
	file_codehost_user_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_codehost_comments_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ReviewComment); i {
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
		file_codehost_comments_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ReviewCommentLocation); i {
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
		file_codehost_comments_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ReviewFileComments); i {
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
		file_codehost_comments_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ReviewDiscussionCountBySeverity); i {
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
		file_codehost_comments_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ReviewDiscussionCount); i {
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
	file_codehost_comments_proto_msgTypes[0].OneofWrappers = []interface{}{}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_codehost_comments_proto_rawDesc,
			NumEnums:      3,
			NumMessages:   6,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_codehost_comments_proto_goTypes,
		DependencyIndexes: file_codehost_comments_proto_depIdxs,
		EnumInfos:         file_codehost_comments_proto_enumTypes,
		MessageInfos:      file_codehost_comments_proto_msgTypes,
	}.Build()
	File_codehost_comments_proto = out.File
	file_codehost_comments_proto_rawDesc = nil
	file_codehost_comments_proto_goTypes = nil
	file_codehost_comments_proto_depIdxs = nil
}