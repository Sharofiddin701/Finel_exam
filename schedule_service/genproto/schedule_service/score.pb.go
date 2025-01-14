// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.33.0
// 	protoc        v3.21.12
// source: score.proto

package schedule_service

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

type ScoreEmpty struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *ScoreEmpty) Reset() {
	*x = ScoreEmpty{}
	if protoimpl.UnsafeEnabled {
		mi := &file_score_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ScoreEmpty) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ScoreEmpty) ProtoMessage() {}

func (x *ScoreEmpty) ProtoReflect() protoreflect.Message {
	mi := &file_score_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ScoreEmpty.ProtoReflect.Descriptor instead.
func (*ScoreEmpty) Descriptor() ([]byte, []int) {
	return file_score_proto_rawDescGZIP(), []int{0}
}

type ScorePrimaryKey struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
}

func (x *ScorePrimaryKey) Reset() {
	*x = ScorePrimaryKey{}
	if protoimpl.UnsafeEnabled {
		mi := &file_score_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ScorePrimaryKey) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ScorePrimaryKey) ProtoMessage() {}

func (x *ScorePrimaryKey) ProtoReflect() protoreflect.Message {
	mi := &file_score_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ScorePrimaryKey.ProtoReflect.Descriptor instead.
func (*ScorePrimaryKey) Descriptor() ([]byte, []int) {
	return file_score_proto_rawDescGZIP(), []int{1}
}

func (x *ScorePrimaryKey) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

type CreateScore struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	TaskId    string `protobuf:"bytes,1,opt,name=task_id,json=taskId,proto3" json:"task_id,omitempty"`
	StudentId string `protobuf:"bytes,2,opt,name=student_id,json=studentId,proto3" json:"student_id,omitempty"`
	Score     int64  `protobuf:"varint,3,opt,name=score,proto3" json:"score,omitempty"`
}

func (x *CreateScore) Reset() {
	*x = CreateScore{}
	if protoimpl.UnsafeEnabled {
		mi := &file_score_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreateScore) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateScore) ProtoMessage() {}

func (x *CreateScore) ProtoReflect() protoreflect.Message {
	mi := &file_score_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateScore.ProtoReflect.Descriptor instead.
func (*CreateScore) Descriptor() ([]byte, []int) {
	return file_score_proto_rawDescGZIP(), []int{2}
}

func (x *CreateScore) GetTaskId() string {
	if x != nil {
		return x.TaskId
	}
	return ""
}

func (x *CreateScore) GetStudentId() string {
	if x != nil {
		return x.StudentId
	}
	return ""
}

func (x *CreateScore) GetScore() int64 {
	if x != nil {
		return x.Score
	}
	return 0
}

type Score struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id        string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	TaskId    string `protobuf:"bytes,2,opt,name=task_id,json=taskId,proto3" json:"task_id,omitempty"`
	StudentId string `protobuf:"bytes,3,opt,name=student_id,json=studentId,proto3" json:"student_id,omitempty"`
	Score     int64  `protobuf:"varint,4,opt,name=score,proto3" json:"score,omitempty"`
	CreatedAt string `protobuf:"bytes,5,opt,name=created_at,json=createdAt,proto3" json:"created_at,omitempty"`
	UpdatedAt string `protobuf:"bytes,6,opt,name=updated_at,json=updatedAt,proto3" json:"updated_at,omitempty"`
}

func (x *Score) Reset() {
	*x = Score{}
	if protoimpl.UnsafeEnabled {
		mi := &file_score_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Score) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Score) ProtoMessage() {}

func (x *Score) ProtoReflect() protoreflect.Message {
	mi := &file_score_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Score.ProtoReflect.Descriptor instead.
func (*Score) Descriptor() ([]byte, []int) {
	return file_score_proto_rawDescGZIP(), []int{3}
}

func (x *Score) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *Score) GetTaskId() string {
	if x != nil {
		return x.TaskId
	}
	return ""
}

func (x *Score) GetStudentId() string {
	if x != nil {
		return x.StudentId
	}
	return ""
}

func (x *Score) GetScore() int64 {
	if x != nil {
		return x.Score
	}
	return 0
}

func (x *Score) GetCreatedAt() string {
	if x != nil {
		return x.CreatedAt
	}
	return ""
}

func (x *Score) GetUpdatedAt() string {
	if x != nil {
		return x.UpdatedAt
	}
	return ""
}

type UpdateScore struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id        string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	TaskId    string `protobuf:"bytes,2,opt,name=task_id,json=taskId,proto3" json:"task_id,omitempty"`
	StudentId string `protobuf:"bytes,3,opt,name=student_id,json=studentId,proto3" json:"student_id,omitempty"`
	Score     int64  `protobuf:"varint,4,opt,name=score,proto3" json:"score,omitempty"`
}

func (x *UpdateScore) Reset() {
	*x = UpdateScore{}
	if protoimpl.UnsafeEnabled {
		mi := &file_score_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UpdateScore) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UpdateScore) ProtoMessage() {}

func (x *UpdateScore) ProtoReflect() protoreflect.Message {
	mi := &file_score_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UpdateScore.ProtoReflect.Descriptor instead.
func (*UpdateScore) Descriptor() ([]byte, []int) {
	return file_score_proto_rawDescGZIP(), []int{4}
}

func (x *UpdateScore) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *UpdateScore) GetTaskId() string {
	if x != nil {
		return x.TaskId
	}
	return ""
}

func (x *UpdateScore) GetStudentId() string {
	if x != nil {
		return x.StudentId
	}
	return ""
}

func (x *UpdateScore) GetScore() int64 {
	if x != nil {
		return x.Score
	}
	return 0
}

type GetListScoreRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Offset int64  `protobuf:"varint,1,opt,name=offset,proto3" json:"offset,omitempty"`
	Limit  int64  `protobuf:"varint,2,opt,name=limit,proto3" json:"limit,omitempty"`
	Search string `protobuf:"bytes,3,opt,name=search,proto3" json:"search,omitempty"`
}

func (x *GetListScoreRequest) Reset() {
	*x = GetListScoreRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_score_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetListScoreRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetListScoreRequest) ProtoMessage() {}

func (x *GetListScoreRequest) ProtoReflect() protoreflect.Message {
	mi := &file_score_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetListScoreRequest.ProtoReflect.Descriptor instead.
func (*GetListScoreRequest) Descriptor() ([]byte, []int) {
	return file_score_proto_rawDescGZIP(), []int{5}
}

func (x *GetListScoreRequest) GetOffset() int64 {
	if x != nil {
		return x.Offset
	}
	return 0
}

func (x *GetListScoreRequest) GetLimit() int64 {
	if x != nil {
		return x.Limit
	}
	return 0
}

func (x *GetListScoreRequest) GetSearch() string {
	if x != nil {
		return x.Search
	}
	return ""
}

type GetListScoreResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Count  int64    `protobuf:"varint,1,opt,name=count,proto3" json:"count,omitempty"`
	Scores []*Score `protobuf:"bytes,2,rep,name=Scores,proto3" json:"Scores,omitempty"`
}

func (x *GetListScoreResponse) Reset() {
	*x = GetListScoreResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_score_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetListScoreResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetListScoreResponse) ProtoMessage() {}

func (x *GetListScoreResponse) ProtoReflect() protoreflect.Message {
	mi := &file_score_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetListScoreResponse.ProtoReflect.Descriptor instead.
func (*GetListScoreResponse) Descriptor() ([]byte, []int) {
	return file_score_proto_rawDescGZIP(), []int{6}
}

func (x *GetListScoreResponse) GetCount() int64 {
	if x != nil {
		return x.Count
	}
	return 0
}

func (x *GetListScoreResponse) GetScores() []*Score {
	if x != nil {
		return x.Scores
	}
	return nil
}

var File_score_proto protoreflect.FileDescriptor

var file_score_proto_rawDesc = []byte{
	0x0a, 0x0b, 0x73, 0x63, 0x6f, 0x72, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x10, 0x73,
	0x63, 0x68, 0x65, 0x64, 0x75, 0x6c, 0x65, 0x5f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x22,
	0x0c, 0x0a, 0x0a, 0x53, 0x63, 0x6f, 0x72, 0x65, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x22, 0x21, 0x0a,
	0x0f, 0x53, 0x63, 0x6f, 0x72, 0x65, 0x50, 0x72, 0x69, 0x6d, 0x61, 0x72, 0x79, 0x4b, 0x65, 0x79,
	0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64,
	0x22, 0x5b, 0x0a, 0x0b, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x53, 0x63, 0x6f, 0x72, 0x65, 0x12,
	0x17, 0x0a, 0x07, 0x74, 0x61, 0x73, 0x6b, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x06, 0x74, 0x61, 0x73, 0x6b, 0x49, 0x64, 0x12, 0x1d, 0x0a, 0x0a, 0x73, 0x74, 0x75, 0x64,
	0x65, 0x6e, 0x74, 0x5f, 0x69, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x73, 0x74,
	0x75, 0x64, 0x65, 0x6e, 0x74, 0x49, 0x64, 0x12, 0x14, 0x0a, 0x05, 0x73, 0x63, 0x6f, 0x72, 0x65,
	0x18, 0x03, 0x20, 0x01, 0x28, 0x03, 0x52, 0x05, 0x73, 0x63, 0x6f, 0x72, 0x65, 0x22, 0xa3, 0x01,
	0x0a, 0x05, 0x53, 0x63, 0x6f, 0x72, 0x65, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x12, 0x17, 0x0a, 0x07, 0x74, 0x61, 0x73, 0x6b, 0x5f,
	0x69, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x74, 0x61, 0x73, 0x6b, 0x49, 0x64,
	0x12, 0x1d, 0x0a, 0x0a, 0x73, 0x74, 0x75, 0x64, 0x65, 0x6e, 0x74, 0x5f, 0x69, 0x64, 0x18, 0x03,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x73, 0x74, 0x75, 0x64, 0x65, 0x6e, 0x74, 0x49, 0x64, 0x12,
	0x14, 0x0a, 0x05, 0x73, 0x63, 0x6f, 0x72, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x03, 0x52, 0x05,
	0x73, 0x63, 0x6f, 0x72, 0x65, 0x12, 0x1d, 0x0a, 0x0a, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64,
	0x5f, 0x61, 0x74, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x63, 0x72, 0x65, 0x61, 0x74,
	0x65, 0x64, 0x41, 0x74, 0x12, 0x1d, 0x0a, 0x0a, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x64, 0x5f,
	0x61, 0x74, 0x18, 0x06, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65,
	0x64, 0x41, 0x74, 0x22, 0x6b, 0x0a, 0x0b, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x53, 0x63, 0x6f,
	0x72, 0x65, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02,
	0x69, 0x64, 0x12, 0x17, 0x0a, 0x07, 0x74, 0x61, 0x73, 0x6b, 0x5f, 0x69, 0x64, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x06, 0x74, 0x61, 0x73, 0x6b, 0x49, 0x64, 0x12, 0x1d, 0x0a, 0x0a, 0x73,
	0x74, 0x75, 0x64, 0x65, 0x6e, 0x74, 0x5f, 0x69, 0x64, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x09, 0x73, 0x74, 0x75, 0x64, 0x65, 0x6e, 0x74, 0x49, 0x64, 0x12, 0x14, 0x0a, 0x05, 0x73, 0x63,
	0x6f, 0x72, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x03, 0x52, 0x05, 0x73, 0x63, 0x6f, 0x72, 0x65,
	0x22, 0x5b, 0x0a, 0x13, 0x47, 0x65, 0x74, 0x4c, 0x69, 0x73, 0x74, 0x53, 0x63, 0x6f, 0x72, 0x65,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x16, 0x0a, 0x06, 0x6f, 0x66, 0x66, 0x73, 0x65,
	0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x06, 0x6f, 0x66, 0x66, 0x73, 0x65, 0x74, 0x12,
	0x14, 0x0a, 0x05, 0x6c, 0x69, 0x6d, 0x69, 0x74, 0x18, 0x02, 0x20, 0x01, 0x28, 0x03, 0x52, 0x05,
	0x6c, 0x69, 0x6d, 0x69, 0x74, 0x12, 0x16, 0x0a, 0x06, 0x73, 0x65, 0x61, 0x72, 0x63, 0x68, 0x18,
	0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x73, 0x65, 0x61, 0x72, 0x63, 0x68, 0x22, 0x5d, 0x0a,
	0x14, 0x47, 0x65, 0x74, 0x4c, 0x69, 0x73, 0x74, 0x53, 0x63, 0x6f, 0x72, 0x65, 0x52, 0x65, 0x73,
	0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x14, 0x0a, 0x05, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x03, 0x52, 0x05, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x12, 0x2f, 0x0a, 0x06, 0x53,
	0x63, 0x6f, 0x72, 0x65, 0x73, 0x18, 0x02, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x17, 0x2e, 0x73, 0x63,
	0x68, 0x65, 0x64, 0x75, 0x6c, 0x65, 0x5f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x53,
	0x63, 0x6f, 0x72, 0x65, 0x52, 0x06, 0x53, 0x63, 0x6f, 0x72, 0x65, 0x73, 0x32, 0x88, 0x03, 0x0a,
	0x0c, 0x53, 0x63, 0x6f, 0x72, 0x65, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x42, 0x0a,
	0x06, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x12, 0x1d, 0x2e, 0x73, 0x63, 0x68, 0x65, 0x64, 0x75,
	0x6c, 0x65, 0x5f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x43, 0x72, 0x65, 0x61, 0x74,
	0x65, 0x53, 0x63, 0x6f, 0x72, 0x65, 0x1a, 0x17, 0x2e, 0x73, 0x63, 0x68, 0x65, 0x64, 0x75, 0x6c,
	0x65, 0x5f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x53, 0x63, 0x6f, 0x72, 0x65, 0x22,
	0x00, 0x12, 0x47, 0x0a, 0x07, 0x47, 0x65, 0x74, 0x42, 0x79, 0x49, 0x44, 0x12, 0x21, 0x2e, 0x73,
	0x63, 0x68, 0x65, 0x64, 0x75, 0x6c, 0x65, 0x5f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e,
	0x53, 0x63, 0x6f, 0x72, 0x65, 0x50, 0x72, 0x69, 0x6d, 0x61, 0x72, 0x79, 0x4b, 0x65, 0x79, 0x1a,
	0x17, 0x2e, 0x73, 0x63, 0x68, 0x65, 0x64, 0x75, 0x6c, 0x65, 0x5f, 0x73, 0x65, 0x72, 0x76, 0x69,
	0x63, 0x65, 0x2e, 0x53, 0x63, 0x6f, 0x72, 0x65, 0x22, 0x00, 0x12, 0x5a, 0x0a, 0x07, 0x47, 0x65,
	0x74, 0x4c, 0x69, 0x73, 0x74, 0x12, 0x25, 0x2e, 0x73, 0x63, 0x68, 0x65, 0x64, 0x75, 0x6c, 0x65,
	0x5f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x47, 0x65, 0x74, 0x4c, 0x69, 0x73, 0x74,
	0x53, 0x63, 0x6f, 0x72, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x26, 0x2e, 0x73,
	0x63, 0x68, 0x65, 0x64, 0x75, 0x6c, 0x65, 0x5f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e,
	0x47, 0x65, 0x74, 0x4c, 0x69, 0x73, 0x74, 0x53, 0x63, 0x6f, 0x72, 0x65, 0x52, 0x65, 0x73, 0x70,
	0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x42, 0x0a, 0x06, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65,
	0x12, 0x1d, 0x2e, 0x73, 0x63, 0x68, 0x65, 0x64, 0x75, 0x6c, 0x65, 0x5f, 0x73, 0x65, 0x72, 0x76,
	0x69, 0x63, 0x65, 0x2e, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x53, 0x63, 0x6f, 0x72, 0x65, 0x1a,
	0x17, 0x2e, 0x73, 0x63, 0x68, 0x65, 0x64, 0x75, 0x6c, 0x65, 0x5f, 0x73, 0x65, 0x72, 0x76, 0x69,
	0x63, 0x65, 0x2e, 0x53, 0x63, 0x6f, 0x72, 0x65, 0x22, 0x00, 0x12, 0x4b, 0x0a, 0x06, 0x44, 0x65,
	0x6c, 0x65, 0x74, 0x65, 0x12, 0x21, 0x2e, 0x73, 0x63, 0x68, 0x65, 0x64, 0x75, 0x6c, 0x65, 0x5f,
	0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x53, 0x63, 0x6f, 0x72, 0x65, 0x50, 0x72, 0x69,
	0x6d, 0x61, 0x72, 0x79, 0x4b, 0x65, 0x79, 0x1a, 0x1c, 0x2e, 0x73, 0x63, 0x68, 0x65, 0x64, 0x75,
	0x6c, 0x65, 0x5f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x53, 0x63, 0x6f, 0x72, 0x65,
	0x45, 0x6d, 0x70, 0x74, 0x79, 0x22, 0x00, 0x42, 0x1b, 0x5a, 0x19, 0x67, 0x65, 0x6e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x2f, 0x73, 0x63, 0x68, 0x65, 0x64, 0x75, 0x6c, 0x65, 0x5f, 0x73, 0x65, 0x72,
	0x76, 0x69, 0x63, 0x65, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_score_proto_rawDescOnce sync.Once
	file_score_proto_rawDescData = file_score_proto_rawDesc
)

func file_score_proto_rawDescGZIP() []byte {
	file_score_proto_rawDescOnce.Do(func() {
		file_score_proto_rawDescData = protoimpl.X.CompressGZIP(file_score_proto_rawDescData)
	})
	return file_score_proto_rawDescData
}

var file_score_proto_msgTypes = make([]protoimpl.MessageInfo, 7)
var file_score_proto_goTypes = []interface{}{
	(*ScoreEmpty)(nil),           // 0: schedule_service.ScoreEmpty
	(*ScorePrimaryKey)(nil),      // 1: schedule_service.ScorePrimaryKey
	(*CreateScore)(nil),          // 2: schedule_service.CreateScore
	(*Score)(nil),                // 3: schedule_service.Score
	(*UpdateScore)(nil),          // 4: schedule_service.UpdateScore
	(*GetListScoreRequest)(nil),  // 5: schedule_service.GetListScoreRequest
	(*GetListScoreResponse)(nil), // 6: schedule_service.GetListScoreResponse
}
var file_score_proto_depIdxs = []int32{
	3, // 0: schedule_service.GetListScoreResponse.Scores:type_name -> schedule_service.Score
	2, // 1: schedule_service.ScoreService.Create:input_type -> schedule_service.CreateScore
	1, // 2: schedule_service.ScoreService.GetByID:input_type -> schedule_service.ScorePrimaryKey
	5, // 3: schedule_service.ScoreService.GetList:input_type -> schedule_service.GetListScoreRequest
	4, // 4: schedule_service.ScoreService.Update:input_type -> schedule_service.UpdateScore
	1, // 5: schedule_service.ScoreService.Delete:input_type -> schedule_service.ScorePrimaryKey
	3, // 6: schedule_service.ScoreService.Create:output_type -> schedule_service.Score
	3, // 7: schedule_service.ScoreService.GetByID:output_type -> schedule_service.Score
	6, // 8: schedule_service.ScoreService.GetList:output_type -> schedule_service.GetListScoreResponse
	3, // 9: schedule_service.ScoreService.Update:output_type -> schedule_service.Score
	0, // 10: schedule_service.ScoreService.Delete:output_type -> schedule_service.ScoreEmpty
	6, // [6:11] is the sub-list for method output_type
	1, // [1:6] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_score_proto_init() }
func file_score_proto_init() {
	if File_score_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_score_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ScoreEmpty); i {
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
		file_score_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ScorePrimaryKey); i {
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
		file_score_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CreateScore); i {
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
		file_score_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Score); i {
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
		file_score_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UpdateScore); i {
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
		file_score_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetListScoreRequest); i {
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
		file_score_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetListScoreResponse); i {
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
			RawDescriptor: file_score_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   7,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_score_proto_goTypes,
		DependencyIndexes: file_score_proto_depIdxs,
		MessageInfos:      file_score_proto_msgTypes,
	}.Build()
	File_score_proto = out.File
	file_score_proto_rawDesc = nil
	file_score_proto_goTypes = nil
	file_score_proto_depIdxs = nil
}
