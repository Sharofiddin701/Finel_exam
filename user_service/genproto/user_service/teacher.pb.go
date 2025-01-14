// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.33.0
// 	protoc        v3.21.12
// source: teacher.proto

package user_service

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

type TeacherEmpty struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *TeacherEmpty) Reset() {
	*x = TeacherEmpty{}
	if protoimpl.UnsafeEnabled {
		mi := &file_teacher_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *TeacherEmpty) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*TeacherEmpty) ProtoMessage() {}

func (x *TeacherEmpty) ProtoReflect() protoreflect.Message {
	mi := &file_teacher_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use TeacherEmpty.ProtoReflect.Descriptor instead.
func (*TeacherEmpty) Descriptor() ([]byte, []int) {
	return file_teacher_proto_rawDescGZIP(), []int{0}
}

type TeacherPrimaryKey struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
}

func (x *TeacherPrimaryKey) Reset() {
	*x = TeacherPrimaryKey{}
	if protoimpl.UnsafeEnabled {
		mi := &file_teacher_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *TeacherPrimaryKey) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*TeacherPrimaryKey) ProtoMessage() {}

func (x *TeacherPrimaryKey) ProtoReflect() protoreflect.Message {
	mi := &file_teacher_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use TeacherPrimaryKey.ProtoReflect.Descriptor instead.
func (*TeacherPrimaryKey) Descriptor() ([]byte, []int) {
	return file_teacher_proto_rawDescGZIP(), []int{1}
}

func (x *TeacherPrimaryKey) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

type CreateTeacher struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	FullName   string  `protobuf:"bytes,1,opt,name=full_name,json=fullName,proto3" json:"full_name,omitempty"`
	Phone      string  `protobuf:"bytes,2,opt,name=phone,proto3" json:"phone,omitempty"`
	Password   string  `protobuf:"bytes,3,opt,name=password,proto3" json:"password,omitempty"`
	Salary     float64 `protobuf:"fixed64,4,opt,name=salary,proto3" json:"salary,omitempty"`
	IeltsScore string  `protobuf:"bytes,5,opt,name=ielts_score,json=ieltsScore,proto3" json:"ielts_score,omitempty"`
	BranchId   string  `protobuf:"bytes,6,opt,name=branch_id,json=branchId,proto3" json:"branch_id,omitempty"`
}

func (x *CreateTeacher) Reset() {
	*x = CreateTeacher{}
	if protoimpl.UnsafeEnabled {
		mi := &file_teacher_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreateTeacher) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateTeacher) ProtoMessage() {}

func (x *CreateTeacher) ProtoReflect() protoreflect.Message {
	mi := &file_teacher_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateTeacher.ProtoReflect.Descriptor instead.
func (*CreateTeacher) Descriptor() ([]byte, []int) {
	return file_teacher_proto_rawDescGZIP(), []int{2}
}

func (x *CreateTeacher) GetFullName() string {
	if x != nil {
		return x.FullName
	}
	return ""
}

func (x *CreateTeacher) GetPhone() string {
	if x != nil {
		return x.Phone
	}
	return ""
}

func (x *CreateTeacher) GetPassword() string {
	if x != nil {
		return x.Password
	}
	return ""
}

func (x *CreateTeacher) GetSalary() float64 {
	if x != nil {
		return x.Salary
	}
	return 0
}

func (x *CreateTeacher) GetIeltsScore() string {
	if x != nil {
		return x.IeltsScore
	}
	return ""
}

func (x *CreateTeacher) GetBranchId() string {
	if x != nil {
		return x.BranchId
	}
	return ""
}

type Teacher struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id         string  `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	FullName   string  `protobuf:"bytes,2,opt,name=full_name,json=fullName,proto3" json:"full_name,omitempty"`
	Phone      string  `protobuf:"bytes,3,opt,name=phone,proto3" json:"phone,omitempty"`
	Password   string  `protobuf:"bytes,4,opt,name=password,proto3" json:"password,omitempty"`
	Login      string  `protobuf:"bytes,5,opt,name=login,proto3" json:"login,omitempty"`
	Salary     float64 `protobuf:"fixed64,6,opt,name=salary,proto3" json:"salary,omitempty"`
	IeltsScore string  `protobuf:"bytes,7,opt,name=ielts_score,json=ieltsScore,proto3" json:"ielts_score,omitempty"`
	BranchId   string  `protobuf:"bytes,8,opt,name=branch_id,json=branchId,proto3" json:"branch_id,omitempty"`
	RoleId     string  `protobuf:"bytes,9,opt,name=role_id,json=roleId,proto3" json:"role_id,omitempty"`
	CreatedAt  string  `protobuf:"bytes,10,opt,name=created_at,json=createdAt,proto3" json:"created_at,omitempty"`
	UpdatedAt  string  `protobuf:"bytes,11,opt,name=updated_at,json=updatedAt,proto3" json:"updated_at,omitempty"`
}

func (x *Teacher) Reset() {
	*x = Teacher{}
	if protoimpl.UnsafeEnabled {
		mi := &file_teacher_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Teacher) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Teacher) ProtoMessage() {}

func (x *Teacher) ProtoReflect() protoreflect.Message {
	mi := &file_teacher_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Teacher.ProtoReflect.Descriptor instead.
func (*Teacher) Descriptor() ([]byte, []int) {
	return file_teacher_proto_rawDescGZIP(), []int{3}
}

func (x *Teacher) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *Teacher) GetFullName() string {
	if x != nil {
		return x.FullName
	}
	return ""
}

func (x *Teacher) GetPhone() string {
	if x != nil {
		return x.Phone
	}
	return ""
}

func (x *Teacher) GetPassword() string {
	if x != nil {
		return x.Password
	}
	return ""
}

func (x *Teacher) GetLogin() string {
	if x != nil {
		return x.Login
	}
	return ""
}

func (x *Teacher) GetSalary() float64 {
	if x != nil {
		return x.Salary
	}
	return 0
}

func (x *Teacher) GetIeltsScore() string {
	if x != nil {
		return x.IeltsScore
	}
	return ""
}

func (x *Teacher) GetBranchId() string {
	if x != nil {
		return x.BranchId
	}
	return ""
}

func (x *Teacher) GetRoleId() string {
	if x != nil {
		return x.RoleId
	}
	return ""
}

func (x *Teacher) GetCreatedAt() string {
	if x != nil {
		return x.CreatedAt
	}
	return ""
}

func (x *Teacher) GetUpdatedAt() string {
	if x != nil {
		return x.UpdatedAt
	}
	return ""
}

type UpdateTeacher struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id         string  `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	FullName   string  `protobuf:"bytes,2,opt,name=full_name,json=fullName,proto3" json:"full_name,omitempty"`
	Phone      string  `protobuf:"bytes,3,opt,name=phone,proto3" json:"phone,omitempty"`
	Password   string  `protobuf:"bytes,4,opt,name=password,proto3" json:"password,omitempty"`
	Login      string  `protobuf:"bytes,5,opt,name=login,proto3" json:"login,omitempty"`
	Salary     float64 `protobuf:"fixed64,6,opt,name=salary,proto3" json:"salary,omitempty"`
	IeltsScore string  `protobuf:"bytes,7,opt,name=ielts_score,json=ieltsScore,proto3" json:"ielts_score,omitempty"`
	BranchId   string  `protobuf:"bytes,8,opt,name=branch_id,json=branchId,proto3" json:"branch_id,omitempty"`
}

func (x *UpdateTeacher) Reset() {
	*x = UpdateTeacher{}
	if protoimpl.UnsafeEnabled {
		mi := &file_teacher_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UpdateTeacher) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UpdateTeacher) ProtoMessage() {}

func (x *UpdateTeacher) ProtoReflect() protoreflect.Message {
	mi := &file_teacher_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UpdateTeacher.ProtoReflect.Descriptor instead.
func (*UpdateTeacher) Descriptor() ([]byte, []int) {
	return file_teacher_proto_rawDescGZIP(), []int{4}
}

func (x *UpdateTeacher) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *UpdateTeacher) GetFullName() string {
	if x != nil {
		return x.FullName
	}
	return ""
}

func (x *UpdateTeacher) GetPhone() string {
	if x != nil {
		return x.Phone
	}
	return ""
}

func (x *UpdateTeacher) GetPassword() string {
	if x != nil {
		return x.Password
	}
	return ""
}

func (x *UpdateTeacher) GetLogin() string {
	if x != nil {
		return x.Login
	}
	return ""
}

func (x *UpdateTeacher) GetSalary() float64 {
	if x != nil {
		return x.Salary
	}
	return 0
}

func (x *UpdateTeacher) GetIeltsScore() string {
	if x != nil {
		return x.IeltsScore
	}
	return ""
}

func (x *UpdateTeacher) GetBranchId() string {
	if x != nil {
		return x.BranchId
	}
	return ""
}

type GetListTeacherRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Offset int64  `protobuf:"varint,1,opt,name=offset,proto3" json:"offset,omitempty"`
	Limit  int64  `protobuf:"varint,2,opt,name=limit,proto3" json:"limit,omitempty"`
	Search string `protobuf:"bytes,3,opt,name=search,proto3" json:"search,omitempty"`
}

func (x *GetListTeacherRequest) Reset() {
	*x = GetListTeacherRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_teacher_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetListTeacherRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetListTeacherRequest) ProtoMessage() {}

func (x *GetListTeacherRequest) ProtoReflect() protoreflect.Message {
	mi := &file_teacher_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetListTeacherRequest.ProtoReflect.Descriptor instead.
func (*GetListTeacherRequest) Descriptor() ([]byte, []int) {
	return file_teacher_proto_rawDescGZIP(), []int{5}
}

func (x *GetListTeacherRequest) GetOffset() int64 {
	if x != nil {
		return x.Offset
	}
	return 0
}

func (x *GetListTeacherRequest) GetLimit() int64 {
	if x != nil {
		return x.Limit
	}
	return 0
}

func (x *GetListTeacherRequest) GetSearch() string {
	if x != nil {
		return x.Search
	}
	return ""
}

type GetListTeacherResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Count    int64      `protobuf:"varint,1,opt,name=count,proto3" json:"count,omitempty"`
	Teachers []*Teacher `protobuf:"bytes,2,rep,name=Teachers,proto3" json:"Teachers,omitempty"`
}

func (x *GetListTeacherResponse) Reset() {
	*x = GetListTeacherResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_teacher_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetListTeacherResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetListTeacherResponse) ProtoMessage() {}

func (x *GetListTeacherResponse) ProtoReflect() protoreflect.Message {
	mi := &file_teacher_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetListTeacherResponse.ProtoReflect.Descriptor instead.
func (*GetListTeacherResponse) Descriptor() ([]byte, []int) {
	return file_teacher_proto_rawDescGZIP(), []int{6}
}

func (x *GetListTeacherResponse) GetCount() int64 {
	if x != nil {
		return x.Count
	}
	return 0
}

func (x *GetListTeacherResponse) GetTeachers() []*Teacher {
	if x != nil {
		return x.Teachers
	}
	return nil
}

type TeacherPanelRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	TeacherId string `protobuf:"bytes,1,opt,name=teacher_id,json=teacherId,proto3" json:"teacher_id,omitempty"`
}

func (x *TeacherPanelRequest) Reset() {
	*x = TeacherPanelRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_teacher_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *TeacherPanelRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*TeacherPanelRequest) ProtoMessage() {}

func (x *TeacherPanelRequest) ProtoReflect() protoreflect.Message {
	mi := &file_teacher_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use TeacherPanelRequest.ProtoReflect.Descriptor instead.
func (*TeacherPanelRequest) Descriptor() ([]byte, []int) {
	return file_teacher_proto_rawDescGZIP(), []int{7}
}

func (x *TeacherPanelRequest) GetTeacherId() string {
	if x != nil {
		return x.TeacherId
	}
	return ""
}

var File_teacher_proto protoreflect.FileDescriptor

var file_teacher_proto_rawDesc = []byte{
	0x0a, 0x0d, 0x74, 0x65, 0x61, 0x63, 0x68, 0x65, 0x72, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12,
	0x0c, 0x75, 0x73, 0x65, 0x72, 0x5f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x22, 0x0e, 0x0a,
	0x0c, 0x54, 0x65, 0x61, 0x63, 0x68, 0x65, 0x72, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x22, 0x23, 0x0a,
	0x11, 0x54, 0x65, 0x61, 0x63, 0x68, 0x65, 0x72, 0x50, 0x72, 0x69, 0x6d, 0x61, 0x72, 0x79, 0x4b,
	0x65, 0x79, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02,
	0x69, 0x64, 0x22, 0xb4, 0x01, 0x0a, 0x0d, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x54, 0x65, 0x61,
	0x63, 0x68, 0x65, 0x72, 0x12, 0x1b, 0x0a, 0x09, 0x66, 0x75, 0x6c, 0x6c, 0x5f, 0x6e, 0x61, 0x6d,
	0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x66, 0x75, 0x6c, 0x6c, 0x4e, 0x61, 0x6d,
	0x65, 0x12, 0x14, 0x0a, 0x05, 0x70, 0x68, 0x6f, 0x6e, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x05, 0x70, 0x68, 0x6f, 0x6e, 0x65, 0x12, 0x1a, 0x0a, 0x08, 0x70, 0x61, 0x73, 0x73, 0x77,
	0x6f, 0x72, 0x64, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x70, 0x61, 0x73, 0x73, 0x77,
	0x6f, 0x72, 0x64, 0x12, 0x16, 0x0a, 0x06, 0x73, 0x61, 0x6c, 0x61, 0x72, 0x79, 0x18, 0x04, 0x20,
	0x01, 0x28, 0x01, 0x52, 0x06, 0x73, 0x61, 0x6c, 0x61, 0x72, 0x79, 0x12, 0x1f, 0x0a, 0x0b, 0x69,
	0x65, 0x6c, 0x74, 0x73, 0x5f, 0x73, 0x63, 0x6f, 0x72, 0x65, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x0a, 0x69, 0x65, 0x6c, 0x74, 0x73, 0x53, 0x63, 0x6f, 0x72, 0x65, 0x12, 0x1b, 0x0a, 0x09,
	0x62, 0x72, 0x61, 0x6e, 0x63, 0x68, 0x5f, 0x69, 0x64, 0x18, 0x06, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x08, 0x62, 0x72, 0x61, 0x6e, 0x63, 0x68, 0x49, 0x64, 0x22, 0xab, 0x02, 0x0a, 0x07, 0x54, 0x65,
	0x61, 0x63, 0x68, 0x65, 0x72, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x02, 0x69, 0x64, 0x12, 0x1b, 0x0a, 0x09, 0x66, 0x75, 0x6c, 0x6c, 0x5f, 0x6e, 0x61,
	0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x66, 0x75, 0x6c, 0x6c, 0x4e, 0x61,
	0x6d, 0x65, 0x12, 0x14, 0x0a, 0x05, 0x70, 0x68, 0x6f, 0x6e, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x05, 0x70, 0x68, 0x6f, 0x6e, 0x65, 0x12, 0x1a, 0x0a, 0x08, 0x70, 0x61, 0x73, 0x73,
	0x77, 0x6f, 0x72, 0x64, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x70, 0x61, 0x73, 0x73,
	0x77, 0x6f, 0x72, 0x64, 0x12, 0x14, 0x0a, 0x05, 0x6c, 0x6f, 0x67, 0x69, 0x6e, 0x18, 0x05, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x05, 0x6c, 0x6f, 0x67, 0x69, 0x6e, 0x12, 0x16, 0x0a, 0x06, 0x73, 0x61,
	0x6c, 0x61, 0x72, 0x79, 0x18, 0x06, 0x20, 0x01, 0x28, 0x01, 0x52, 0x06, 0x73, 0x61, 0x6c, 0x61,
	0x72, 0x79, 0x12, 0x1f, 0x0a, 0x0b, 0x69, 0x65, 0x6c, 0x74, 0x73, 0x5f, 0x73, 0x63, 0x6f, 0x72,
	0x65, 0x18, 0x07, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x69, 0x65, 0x6c, 0x74, 0x73, 0x53, 0x63,
	0x6f, 0x72, 0x65, 0x12, 0x1b, 0x0a, 0x09, 0x62, 0x72, 0x61, 0x6e, 0x63, 0x68, 0x5f, 0x69, 0x64,
	0x18, 0x08, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x62, 0x72, 0x61, 0x6e, 0x63, 0x68, 0x49, 0x64,
	0x12, 0x17, 0x0a, 0x07, 0x72, 0x6f, 0x6c, 0x65, 0x5f, 0x69, 0x64, 0x18, 0x09, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x06, 0x72, 0x6f, 0x6c, 0x65, 0x49, 0x64, 0x12, 0x1d, 0x0a, 0x0a, 0x63, 0x72, 0x65,
	0x61, 0x74, 0x65, 0x64, 0x5f, 0x61, 0x74, 0x18, 0x0a, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x63,
	0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x41, 0x74, 0x12, 0x1d, 0x0a, 0x0a, 0x75, 0x70, 0x64, 0x61,
	0x74, 0x65, 0x64, 0x5f, 0x61, 0x74, 0x18, 0x0b, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x75, 0x70,
	0x64, 0x61, 0x74, 0x65, 0x64, 0x41, 0x74, 0x22, 0xda, 0x01, 0x0a, 0x0d, 0x55, 0x70, 0x64, 0x61,
	0x74, 0x65, 0x54, 0x65, 0x61, 0x63, 0x68, 0x65, 0x72, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x12, 0x1b, 0x0a, 0x09, 0x66, 0x75, 0x6c,
	0x6c, 0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x66, 0x75,
	0x6c, 0x6c, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x14, 0x0a, 0x05, 0x70, 0x68, 0x6f, 0x6e, 0x65, 0x18,
	0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x70, 0x68, 0x6f, 0x6e, 0x65, 0x12, 0x1a, 0x0a, 0x08,
	0x70, 0x61, 0x73, 0x73, 0x77, 0x6f, 0x72, 0x64, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08,
	0x70, 0x61, 0x73, 0x73, 0x77, 0x6f, 0x72, 0x64, 0x12, 0x14, 0x0a, 0x05, 0x6c, 0x6f, 0x67, 0x69,
	0x6e, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x6c, 0x6f, 0x67, 0x69, 0x6e, 0x12, 0x16,
	0x0a, 0x06, 0x73, 0x61, 0x6c, 0x61, 0x72, 0x79, 0x18, 0x06, 0x20, 0x01, 0x28, 0x01, 0x52, 0x06,
	0x73, 0x61, 0x6c, 0x61, 0x72, 0x79, 0x12, 0x1f, 0x0a, 0x0b, 0x69, 0x65, 0x6c, 0x74, 0x73, 0x5f,
	0x73, 0x63, 0x6f, 0x72, 0x65, 0x18, 0x07, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x69, 0x65, 0x6c,
	0x74, 0x73, 0x53, 0x63, 0x6f, 0x72, 0x65, 0x12, 0x1b, 0x0a, 0x09, 0x62, 0x72, 0x61, 0x6e, 0x63,
	0x68, 0x5f, 0x69, 0x64, 0x18, 0x08, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x62, 0x72, 0x61, 0x6e,
	0x63, 0x68, 0x49, 0x64, 0x22, 0x5d, 0x0a, 0x15, 0x47, 0x65, 0x74, 0x4c, 0x69, 0x73, 0x74, 0x54,
	0x65, 0x61, 0x63, 0x68, 0x65, 0x72, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x16, 0x0a,
	0x06, 0x6f, 0x66, 0x66, 0x73, 0x65, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x06, 0x6f,
	0x66, 0x66, 0x73, 0x65, 0x74, 0x12, 0x14, 0x0a, 0x05, 0x6c, 0x69, 0x6d, 0x69, 0x74, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x03, 0x52, 0x05, 0x6c, 0x69, 0x6d, 0x69, 0x74, 0x12, 0x16, 0x0a, 0x06, 0x73,
	0x65, 0x61, 0x72, 0x63, 0x68, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x73, 0x65, 0x61,
	0x72, 0x63, 0x68, 0x22, 0x61, 0x0a, 0x16, 0x47, 0x65, 0x74, 0x4c, 0x69, 0x73, 0x74, 0x54, 0x65,
	0x61, 0x63, 0x68, 0x65, 0x72, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x14, 0x0a,
	0x05, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x05, 0x63, 0x6f,
	0x75, 0x6e, 0x74, 0x12, 0x31, 0x0a, 0x08, 0x54, 0x65, 0x61, 0x63, 0x68, 0x65, 0x72, 0x73, 0x18,
	0x02, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x15, 0x2e, 0x75, 0x73, 0x65, 0x72, 0x5f, 0x73, 0x65, 0x72,
	0x76, 0x69, 0x63, 0x65, 0x2e, 0x54, 0x65, 0x61, 0x63, 0x68, 0x65, 0x72, 0x52, 0x08, 0x54, 0x65,
	0x61, 0x63, 0x68, 0x65, 0x72, 0x73, 0x22, 0x34, 0x0a, 0x13, 0x54, 0x65, 0x61, 0x63, 0x68, 0x65,
	0x72, 0x50, 0x61, 0x6e, 0x65, 0x6c, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x1d, 0x0a,
	0x0a, 0x74, 0x65, 0x61, 0x63, 0x68, 0x65, 0x72, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x09, 0x74, 0x65, 0x61, 0x63, 0x68, 0x65, 0x72, 0x49, 0x64, 0x32, 0xf6, 0x02, 0x0a,
	0x0e, 0x54, 0x65, 0x61, 0x63, 0x68, 0x65, 0x72, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12,
	0x3e, 0x0a, 0x06, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x12, 0x1b, 0x2e, 0x75, 0x73, 0x65, 0x72,
	0x5f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x54,
	0x65, 0x61, 0x63, 0x68, 0x65, 0x72, 0x1a, 0x15, 0x2e, 0x75, 0x73, 0x65, 0x72, 0x5f, 0x73, 0x65,
	0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x54, 0x65, 0x61, 0x63, 0x68, 0x65, 0x72, 0x22, 0x00, 0x12,
	0x43, 0x0a, 0x07, 0x47, 0x65, 0x74, 0x42, 0x79, 0x49, 0x44, 0x12, 0x1f, 0x2e, 0x75, 0x73, 0x65,
	0x72, 0x5f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x54, 0x65, 0x61, 0x63, 0x68, 0x65,
	0x72, 0x50, 0x72, 0x69, 0x6d, 0x61, 0x72, 0x79, 0x4b, 0x65, 0x79, 0x1a, 0x15, 0x2e, 0x75, 0x73,
	0x65, 0x72, 0x5f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x54, 0x65, 0x61, 0x63, 0x68,
	0x65, 0x72, 0x22, 0x00, 0x12, 0x56, 0x0a, 0x07, 0x47, 0x65, 0x74, 0x4c, 0x69, 0x73, 0x74, 0x12,
	0x23, 0x2e, 0x75, 0x73, 0x65, 0x72, 0x5f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x47,
	0x65, 0x74, 0x4c, 0x69, 0x73, 0x74, 0x54, 0x65, 0x61, 0x63, 0x68, 0x65, 0x72, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x1a, 0x24, 0x2e, 0x75, 0x73, 0x65, 0x72, 0x5f, 0x73, 0x65, 0x72, 0x76,
	0x69, 0x63, 0x65, 0x2e, 0x47, 0x65, 0x74, 0x4c, 0x69, 0x73, 0x74, 0x54, 0x65, 0x61, 0x63, 0x68,
	0x65, 0x72, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x3e, 0x0a, 0x06,
	0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x12, 0x1b, 0x2e, 0x75, 0x73, 0x65, 0x72, 0x5f, 0x73, 0x65,
	0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x54, 0x65, 0x61, 0x63,
	0x68, 0x65, 0x72, 0x1a, 0x15, 0x2e, 0x75, 0x73, 0x65, 0x72, 0x5f, 0x73, 0x65, 0x72, 0x76, 0x69,
	0x63, 0x65, 0x2e, 0x54, 0x65, 0x61, 0x63, 0x68, 0x65, 0x72, 0x22, 0x00, 0x12, 0x47, 0x0a, 0x06,
	0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x12, 0x1f, 0x2e, 0x75, 0x73, 0x65, 0x72, 0x5f, 0x73, 0x65,
	0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x54, 0x65, 0x61, 0x63, 0x68, 0x65, 0x72, 0x50, 0x72, 0x69,
	0x6d, 0x61, 0x72, 0x79, 0x4b, 0x65, 0x79, 0x1a, 0x1a, 0x2e, 0x75, 0x73, 0x65, 0x72, 0x5f, 0x73,
	0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x54, 0x65, 0x61, 0x63, 0x68, 0x65, 0x72, 0x45, 0x6d,
	0x70, 0x74, 0x79, 0x22, 0x00, 0x42, 0x17, 0x5a, 0x15, 0x67, 0x65, 0x6e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x2f, 0x75, 0x73, 0x65, 0x72, 0x5f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x62, 0x06,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_teacher_proto_rawDescOnce sync.Once
	file_teacher_proto_rawDescData = file_teacher_proto_rawDesc
)

func file_teacher_proto_rawDescGZIP() []byte {
	file_teacher_proto_rawDescOnce.Do(func() {
		file_teacher_proto_rawDescData = protoimpl.X.CompressGZIP(file_teacher_proto_rawDescData)
	})
	return file_teacher_proto_rawDescData
}

var file_teacher_proto_msgTypes = make([]protoimpl.MessageInfo, 8)
var file_teacher_proto_goTypes = []interface{}{
	(*TeacherEmpty)(nil),           // 0: user_service.TeacherEmpty
	(*TeacherPrimaryKey)(nil),      // 1: user_service.TeacherPrimaryKey
	(*CreateTeacher)(nil),          // 2: user_service.CreateTeacher
	(*Teacher)(nil),                // 3: user_service.Teacher
	(*UpdateTeacher)(nil),          // 4: user_service.UpdateTeacher
	(*GetListTeacherRequest)(nil),  // 5: user_service.GetListTeacherRequest
	(*GetListTeacherResponse)(nil), // 6: user_service.GetListTeacherResponse
	(*TeacherPanelRequest)(nil),    // 7: user_service.TeacherPanelRequest
}
var file_teacher_proto_depIdxs = []int32{
	3, // 0: user_service.GetListTeacherResponse.Teachers:type_name -> user_service.Teacher
	2, // 1: user_service.TeacherService.Create:input_type -> user_service.CreateTeacher
	1, // 2: user_service.TeacherService.GetByID:input_type -> user_service.TeacherPrimaryKey
	5, // 3: user_service.TeacherService.GetList:input_type -> user_service.GetListTeacherRequest
	4, // 4: user_service.TeacherService.Update:input_type -> user_service.UpdateTeacher
	1, // 5: user_service.TeacherService.Delete:input_type -> user_service.TeacherPrimaryKey
	3, // 6: user_service.TeacherService.Create:output_type -> user_service.Teacher
	3, // 7: user_service.TeacherService.GetByID:output_type -> user_service.Teacher
	6, // 8: user_service.TeacherService.GetList:output_type -> user_service.GetListTeacherResponse
	3, // 9: user_service.TeacherService.Update:output_type -> user_service.Teacher
	0, // 10: user_service.TeacherService.Delete:output_type -> user_service.TeacherEmpty
	6, // [6:11] is the sub-list for method output_type
	1, // [1:6] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_teacher_proto_init() }
func file_teacher_proto_init() {
	if File_teacher_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_teacher_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*TeacherEmpty); i {
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
		file_teacher_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*TeacherPrimaryKey); i {
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
		file_teacher_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CreateTeacher); i {
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
		file_teacher_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Teacher); i {
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
		file_teacher_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UpdateTeacher); i {
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
		file_teacher_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetListTeacherRequest); i {
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
		file_teacher_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetListTeacherResponse); i {
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
		file_teacher_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*TeacherPanelRequest); i {
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
			RawDescriptor: file_teacher_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   8,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_teacher_proto_goTypes,
		DependencyIndexes: file_teacher_proto_depIdxs,
		MessageInfos:      file_teacher_proto_msgTypes,
	}.Build()
	File_teacher_proto = out.File
	file_teacher_proto_rawDesc = nil
	file_teacher_proto_goTypes = nil
	file_teacher_proto_depIdxs = nil
}
