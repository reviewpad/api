// Copyright 2023 Explore.dev Unipessoal Lda. All Rights Reserved.
// Use of this source code is governed by a license that can be
// found in the LICENSE file.

// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.30.0
// 	protoc        v3.21.12
// source: codehost/code_review.proto

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

type UserReviewStatus int32

const (
	UserReviewStatus_USER_REQUESTED_CHANGES UserReviewStatus = 0
	UserReviewStatus_USER_APPROVED          UserReviewStatus = 1
	UserReviewStatus_USER_COMMENTED         UserReviewStatus = 2
	UserReviewStatus_PENDING                UserReviewStatus = 3
	UserReviewStatus_DISMISSED              UserReviewStatus = 4
)

// Enum value maps for UserReviewStatus.
var (
	UserReviewStatus_name = map[int32]string{
		0: "USER_REQUESTED_CHANGES",
		1: "USER_APPROVED",
		2: "USER_COMMENTED",
		3: "PENDING",
		4: "DISMISSED",
	}
	UserReviewStatus_value = map[string]int32{
		"USER_REQUESTED_CHANGES": 0,
		"USER_APPROVED":          1,
		"USER_COMMENTED":         2,
		"PENDING":                3,
		"DISMISSED":              4,
	}
)

func (x UserReviewStatus) Enum() *UserReviewStatus {
	p := new(UserReviewStatus)
	*p = x
	return p
}

func (x UserReviewStatus) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (UserReviewStatus) Descriptor() protoreflect.EnumDescriptor {
	return file_codehost_code_review_proto_enumTypes[0].Descriptor()
}

func (UserReviewStatus) Type() protoreflect.EnumType {
	return &file_codehost_code_review_proto_enumTypes[0]
}

func (x UserReviewStatus) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use UserReviewStatus.Descriptor instead.
func (UserReviewStatus) EnumDescriptor() ([]byte, []int) {
	return file_codehost_code_review_proto_rawDescGZIP(), []int{0}
}

type ReviewEvent int32

const (
	ReviewEvent_APPROVE         ReviewEvent = 0
	ReviewEvent_REQUEST_CHANGES ReviewEvent = 1
	ReviewEvent_COMMENT         ReviewEvent = 2
	ReviewEvent_UNAPPROVE       ReviewEvent = 3
)

// Enum value maps for ReviewEvent.
var (
	ReviewEvent_name = map[int32]string{
		0: "APPROVE",
		1: "REQUEST_CHANGES",
		2: "COMMENT",
		3: "UNAPPROVE",
	}
	ReviewEvent_value = map[string]int32{
		"APPROVE":         0,
		"REQUEST_CHANGES": 1,
		"COMMENT":         2,
		"UNAPPROVE":       3,
	}
)

func (x ReviewEvent) Enum() *ReviewEvent {
	p := new(ReviewEvent)
	*p = x
	return p
}

func (x ReviewEvent) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (ReviewEvent) Descriptor() protoreflect.EnumDescriptor {
	return file_codehost_code_review_proto_enumTypes[1].Descriptor()
}

func (ReviewEvent) Type() protoreflect.EnumType {
	return &file_codehost_code_review_proto_enumTypes[1]
}

func (x ReviewEvent) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use ReviewEvent.Descriptor instead.
func (ReviewEvent) EnumDescriptor() ([]byte, []int) {
	return file_codehost_code_review_proto_rawDescGZIP(), []int{1}
}

type UserReview struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id           string           `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	ReviewerId   string           `protobuf:"bytes,2,opt,name=reviewer_id,json=reviewerId,proto3" json:"reviewer_id,omitempty"`
	BaseCommitId string           `protobuf:"bytes,3,opt,name=base_commit_id,json=baseCommitId,proto3" json:"base_commit_id,omitempty"`
	HeadCommitId string           `protobuf:"bytes,4,opt,name=head_commit_id,json=headCommitId,proto3" json:"head_commit_id,omitempty"`
	Body         string           `protobuf:"bytes,5,opt,name=body,proto3" json:"body,omitempty"`
	Status       UserReviewStatus `protobuf:"varint,6,opt,name=status,proto3,enum=codehost.UserReviewStatus" json:"status,omitempty"`
	CreatedAt    int64            `protobuf:"varint,7,opt,name=created_at,json=createdAt,proto3" json:"created_at,omitempty"`
	UpdatedAt    int64            `protobuf:"varint,8,opt,name=updated_at,json=updatedAt,proto3" json:"updated_at,omitempty"`
	ExternalId   string           `protobuf:"bytes,9,opt,name=external_id,json=externalId,proto3" json:"external_id,omitempty"`
}

func (x *UserReview) Reset() {
	*x = UserReview{}
	if protoimpl.UnsafeEnabled {
		mi := &file_codehost_code_review_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UserReview) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UserReview) ProtoMessage() {}

func (x *UserReview) ProtoReflect() protoreflect.Message {
	mi := &file_codehost_code_review_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UserReview.ProtoReflect.Descriptor instead.
func (*UserReview) Descriptor() ([]byte, []int) {
	return file_codehost_code_review_proto_rawDescGZIP(), []int{0}
}

func (x *UserReview) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *UserReview) GetReviewerId() string {
	if x != nil {
		return x.ReviewerId
	}
	return ""
}

func (x *UserReview) GetBaseCommitId() string {
	if x != nil {
		return x.BaseCommitId
	}
	return ""
}

func (x *UserReview) GetHeadCommitId() string {
	if x != nil {
		return x.HeadCommitId
	}
	return ""
}

func (x *UserReview) GetBody() string {
	if x != nil {
		return x.Body
	}
	return ""
}

func (x *UserReview) GetStatus() UserReviewStatus {
	if x != nil {
		return x.Status
	}
	return UserReviewStatus_USER_REQUESTED_CHANGES
}

func (x *UserReview) GetCreatedAt() int64 {
	if x != nil {
		return x.CreatedAt
	}
	return 0
}

func (x *UserReview) GetUpdatedAt() int64 {
	if x != nil {
		return x.UpdatedAt
	}
	return 0
}

func (x *UserReview) GetExternalId() string {
	if x != nil {
		return x.ExternalId
	}
	return ""
}

var File_codehost_code_review_proto protoreflect.FileDescriptor

var file_codehost_code_review_proto_rawDesc = []byte{
	0x0a, 0x1a, 0x63, 0x6f, 0x64, 0x65, 0x68, 0x6f, 0x73, 0x74, 0x2f, 0x63, 0x6f, 0x64, 0x65, 0x5f,
	0x72, 0x65, 0x76, 0x69, 0x65, 0x77, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x08, 0x63, 0x6f,
	0x64, 0x65, 0x68, 0x6f, 0x73, 0x74, 0x22, 0xb0, 0x02, 0x0a, 0x0a, 0x55, 0x73, 0x65, 0x72, 0x52,
	0x65, 0x76, 0x69, 0x65, 0x77, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x02, 0x69, 0x64, 0x12, 0x1f, 0x0a, 0x0b, 0x72, 0x65, 0x76, 0x69, 0x65, 0x77, 0x65,
	0x72, 0x5f, 0x69, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x72, 0x65, 0x76, 0x69,
	0x65, 0x77, 0x65, 0x72, 0x49, 0x64, 0x12, 0x24, 0x0a, 0x0e, 0x62, 0x61, 0x73, 0x65, 0x5f, 0x63,
	0x6f, 0x6d, 0x6d, 0x69, 0x74, 0x5f, 0x69, 0x64, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0c,
	0x62, 0x61, 0x73, 0x65, 0x43, 0x6f, 0x6d, 0x6d, 0x69, 0x74, 0x49, 0x64, 0x12, 0x24, 0x0a, 0x0e,
	0x68, 0x65, 0x61, 0x64, 0x5f, 0x63, 0x6f, 0x6d, 0x6d, 0x69, 0x74, 0x5f, 0x69, 0x64, 0x18, 0x04,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x0c, 0x68, 0x65, 0x61, 0x64, 0x43, 0x6f, 0x6d, 0x6d, 0x69, 0x74,
	0x49, 0x64, 0x12, 0x12, 0x0a, 0x04, 0x62, 0x6f, 0x64, 0x79, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x04, 0x62, 0x6f, 0x64, 0x79, 0x12, 0x32, 0x0a, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73,
	0x18, 0x06, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x1a, 0x2e, 0x63, 0x6f, 0x64, 0x65, 0x68, 0x6f, 0x73,
	0x74, 0x2e, 0x55, 0x73, 0x65, 0x72, 0x52, 0x65, 0x76, 0x69, 0x65, 0x77, 0x53, 0x74, 0x61, 0x74,
	0x75, 0x73, 0x52, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x12, 0x1d, 0x0a, 0x0a, 0x63, 0x72,
	0x65, 0x61, 0x74, 0x65, 0x64, 0x5f, 0x61, 0x74, 0x18, 0x07, 0x20, 0x01, 0x28, 0x03, 0x52, 0x09,
	0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x41, 0x74, 0x12, 0x1d, 0x0a, 0x0a, 0x75, 0x70, 0x64,
	0x61, 0x74, 0x65, 0x64, 0x5f, 0x61, 0x74, 0x18, 0x08, 0x20, 0x01, 0x28, 0x03, 0x52, 0x09, 0x75,
	0x70, 0x64, 0x61, 0x74, 0x65, 0x64, 0x41, 0x74, 0x12, 0x1f, 0x0a, 0x0b, 0x65, 0x78, 0x74, 0x65,
	0x72, 0x6e, 0x61, 0x6c, 0x5f, 0x69, 0x64, 0x18, 0x09, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x65,
	0x78, 0x74, 0x65, 0x72, 0x6e, 0x61, 0x6c, 0x49, 0x64, 0x2a, 0x71, 0x0a, 0x10, 0x55, 0x73, 0x65,
	0x72, 0x52, 0x65, 0x76, 0x69, 0x65, 0x77, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x12, 0x1a, 0x0a,
	0x16, 0x55, 0x53, 0x45, 0x52, 0x5f, 0x52, 0x45, 0x51, 0x55, 0x45, 0x53, 0x54, 0x45, 0x44, 0x5f,
	0x43, 0x48, 0x41, 0x4e, 0x47, 0x45, 0x53, 0x10, 0x00, 0x12, 0x11, 0x0a, 0x0d, 0x55, 0x53, 0x45,
	0x52, 0x5f, 0x41, 0x50, 0x50, 0x52, 0x4f, 0x56, 0x45, 0x44, 0x10, 0x01, 0x12, 0x12, 0x0a, 0x0e,
	0x55, 0x53, 0x45, 0x52, 0x5f, 0x43, 0x4f, 0x4d, 0x4d, 0x45, 0x4e, 0x54, 0x45, 0x44, 0x10, 0x02,
	0x12, 0x0b, 0x0a, 0x07, 0x50, 0x45, 0x4e, 0x44, 0x49, 0x4e, 0x47, 0x10, 0x03, 0x12, 0x0d, 0x0a,
	0x09, 0x44, 0x49, 0x53, 0x4d, 0x49, 0x53, 0x53, 0x45, 0x44, 0x10, 0x04, 0x2a, 0x4b, 0x0a, 0x0b,
	0x52, 0x65, 0x76, 0x69, 0x65, 0x77, 0x45, 0x76, 0x65, 0x6e, 0x74, 0x12, 0x0b, 0x0a, 0x07, 0x41,
	0x50, 0x50, 0x52, 0x4f, 0x56, 0x45, 0x10, 0x00, 0x12, 0x13, 0x0a, 0x0f, 0x52, 0x45, 0x51, 0x55,
	0x45, 0x53, 0x54, 0x5f, 0x43, 0x48, 0x41, 0x4e, 0x47, 0x45, 0x53, 0x10, 0x01, 0x12, 0x0b, 0x0a,
	0x07, 0x43, 0x4f, 0x4d, 0x4d, 0x45, 0x4e, 0x54, 0x10, 0x02, 0x12, 0x0d, 0x0a, 0x09, 0x55, 0x4e,
	0x41, 0x50, 0x50, 0x52, 0x4f, 0x56, 0x45, 0x10, 0x03, 0x42, 0x26, 0x5a, 0x24, 0x67, 0x69, 0x74,
	0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x72, 0x65, 0x76, 0x69, 0x65, 0x77, 0x70, 0x61,
	0x64, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x67, 0x6f, 0x2f, 0x63, 0x6f, 0x64, 0x65, 0x68, 0x6f, 0x73,
	0x74, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_codehost_code_review_proto_rawDescOnce sync.Once
	file_codehost_code_review_proto_rawDescData = file_codehost_code_review_proto_rawDesc
)

func file_codehost_code_review_proto_rawDescGZIP() []byte {
	file_codehost_code_review_proto_rawDescOnce.Do(func() {
		file_codehost_code_review_proto_rawDescData = protoimpl.X.CompressGZIP(file_codehost_code_review_proto_rawDescData)
	})
	return file_codehost_code_review_proto_rawDescData
}

var file_codehost_code_review_proto_enumTypes = make([]protoimpl.EnumInfo, 2)
var file_codehost_code_review_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_codehost_code_review_proto_goTypes = []interface{}{
	(UserReviewStatus)(0), // 0: codehost.UserReviewStatus
	(ReviewEvent)(0),      // 1: codehost.ReviewEvent
	(*UserReview)(nil),    // 2: codehost.UserReview
}
var file_codehost_code_review_proto_depIdxs = []int32{
	0, // 0: codehost.UserReview.status:type_name -> codehost.UserReviewStatus
	1, // [1:1] is the sub-list for method output_type
	1, // [1:1] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_codehost_code_review_proto_init() }
func file_codehost_code_review_proto_init() {
	if File_codehost_code_review_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_codehost_code_review_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UserReview); i {
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
			RawDescriptor: file_codehost_code_review_proto_rawDesc,
			NumEnums:      2,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_codehost_code_review_proto_goTypes,
		DependencyIndexes: file_codehost_code_review_proto_depIdxs,
		EnumInfos:         file_codehost_code_review_proto_enumTypes,
		MessageInfos:      file_codehost_code_review_proto_msgTypes,
	}.Build()
	File_codehost_code_review_proto = out.File
	file_codehost_code_review_proto_rawDesc = nil
	file_codehost_code_review_proto_goTypes = nil
	file_codehost_code_review_proto_depIdxs = nil
}
