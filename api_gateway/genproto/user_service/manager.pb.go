// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.33.0
// 	protoc        v3.21.12
// source: manager.proto

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

type ManagerEmpty struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *ManagerEmpty) Reset() {
	*x = ManagerEmpty{}
	if protoimpl.UnsafeEnabled {
		mi := &file_manager_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ManagerEmpty) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ManagerEmpty) ProtoMessage() {}

func (x *ManagerEmpty) ProtoReflect() protoreflect.Message {
	mi := &file_manager_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ManagerEmpty.ProtoReflect.Descriptor instead.
func (*ManagerEmpty) Descriptor() ([]byte, []int) {
	return file_manager_proto_rawDescGZIP(), []int{0}
}

type ManagerPrimaryKey struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
}

func (x *ManagerPrimaryKey) Reset() {
	*x = ManagerPrimaryKey{}
	if protoimpl.UnsafeEnabled {
		mi := &file_manager_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ManagerPrimaryKey) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ManagerPrimaryKey) ProtoMessage() {}

func (x *ManagerPrimaryKey) ProtoReflect() protoreflect.Message {
	mi := &file_manager_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ManagerPrimaryKey.ProtoReflect.Descriptor instead.
func (*ManagerPrimaryKey) Descriptor() ([]byte, []int) {
	return file_manager_proto_rawDescGZIP(), []int{1}
}

func (x *ManagerPrimaryKey) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

type CreateManager struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	FullName string  `protobuf:"bytes,1,opt,name=full_name,json=fullName,proto3" json:"full_name,omitempty"`
	Phone    string  `protobuf:"bytes,2,opt,name=phone,proto3" json:"phone,omitempty"`
	Salary   float64 `protobuf:"fixed64,3,opt,name=salary,proto3" json:"salary,omitempty"`
	Password string  `protobuf:"bytes,4,opt,name=password,proto3" json:"password,omitempty"`
	BranchId string  `protobuf:"bytes,5,opt,name=branch_id,json=branchId,proto3" json:"branch_id,omitempty"`
}

func (x *CreateManager) Reset() {
	*x = CreateManager{}
	if protoimpl.UnsafeEnabled {
		mi := &file_manager_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreateManager) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateManager) ProtoMessage() {}

func (x *CreateManager) ProtoReflect() protoreflect.Message {
	mi := &file_manager_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateManager.ProtoReflect.Descriptor instead.
func (*CreateManager) Descriptor() ([]byte, []int) {
	return file_manager_proto_rawDescGZIP(), []int{2}
}

func (x *CreateManager) GetFullName() string {
	if x != nil {
		return x.FullName
	}
	return ""
}

func (x *CreateManager) GetPhone() string {
	if x != nil {
		return x.Phone
	}
	return ""
}

func (x *CreateManager) GetSalary() float64 {
	if x != nil {
		return x.Salary
	}
	return 0
}

func (x *CreateManager) GetPassword() string {
	if x != nil {
		return x.Password
	}
	return ""
}

func (x *CreateManager) GetBranchId() string {
	if x != nil {
		return x.BranchId
	}
	return ""
}

type Manager struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id        string  `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	FullName  string  `protobuf:"bytes,2,opt,name=full_name,json=fullName,proto3" json:"full_name,omitempty"`
	Phone     string  `protobuf:"bytes,3,opt,name=phone,proto3" json:"phone,omitempty"`
	Salary    float64 `protobuf:"fixed64,4,opt,name=salary,proto3" json:"salary,omitempty"`
	Password  string  `protobuf:"bytes,5,opt,name=password,proto3" json:"password,omitempty"`
	Login     string  `protobuf:"bytes,6,opt,name=login,proto3" json:"login,omitempty"`
	BranchId  string  `protobuf:"bytes,7,opt,name=branch_id,json=branchId,proto3" json:"branch_id,omitempty"`
	RoleId    string  `protobuf:"bytes,8,opt,name=role_id,json=roleId,proto3" json:"role_id,omitempty"`
	CreatedAt string  `protobuf:"bytes,9,opt,name=created_at,json=createdAt,proto3" json:"created_at,omitempty"`
	UpdatedAt string  `protobuf:"bytes,10,opt,name=updated_at,json=updatedAt,proto3" json:"updated_at,omitempty"`
}

func (x *Manager) Reset() {
	*x = Manager{}
	if protoimpl.UnsafeEnabled {
		mi := &file_manager_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Manager) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Manager) ProtoMessage() {}

func (x *Manager) ProtoReflect() protoreflect.Message {
	mi := &file_manager_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Manager.ProtoReflect.Descriptor instead.
func (*Manager) Descriptor() ([]byte, []int) {
	return file_manager_proto_rawDescGZIP(), []int{3}
}

func (x *Manager) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *Manager) GetFullName() string {
	if x != nil {
		return x.FullName
	}
	return ""
}

func (x *Manager) GetPhone() string {
	if x != nil {
		return x.Phone
	}
	return ""
}

func (x *Manager) GetSalary() float64 {
	if x != nil {
		return x.Salary
	}
	return 0
}

func (x *Manager) GetPassword() string {
	if x != nil {
		return x.Password
	}
	return ""
}

func (x *Manager) GetLogin() string {
	if x != nil {
		return x.Login
	}
	return ""
}

func (x *Manager) GetBranchId() string {
	if x != nil {
		return x.BranchId
	}
	return ""
}

func (x *Manager) GetRoleId() string {
	if x != nil {
		return x.RoleId
	}
	return ""
}

func (x *Manager) GetCreatedAt() string {
	if x != nil {
		return x.CreatedAt
	}
	return ""
}

func (x *Manager) GetUpdatedAt() string {
	if x != nil {
		return x.UpdatedAt
	}
	return ""
}

type UpdateManager struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id       string  `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	FullName string  `protobuf:"bytes,2,opt,name=full_name,json=fullName,proto3" json:"full_name,omitempty"`
	Phone    string  `protobuf:"bytes,3,opt,name=phone,proto3" json:"phone,omitempty"`
	Salary   float64 `protobuf:"fixed64,4,opt,name=salary,proto3" json:"salary,omitempty"`
	Password string  `protobuf:"bytes,5,opt,name=password,proto3" json:"password,omitempty"`
	Login    string  `protobuf:"bytes,6,opt,name=login,proto3" json:"login,omitempty"`
	BranchId string  `protobuf:"bytes,7,opt,name=branch_id,json=branchId,proto3" json:"branch_id,omitempty"`
}

func (x *UpdateManager) Reset() {
	*x = UpdateManager{}
	if protoimpl.UnsafeEnabled {
		mi := &file_manager_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UpdateManager) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UpdateManager) ProtoMessage() {}

func (x *UpdateManager) ProtoReflect() protoreflect.Message {
	mi := &file_manager_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UpdateManager.ProtoReflect.Descriptor instead.
func (*UpdateManager) Descriptor() ([]byte, []int) {
	return file_manager_proto_rawDescGZIP(), []int{4}
}

func (x *UpdateManager) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *UpdateManager) GetFullName() string {
	if x != nil {
		return x.FullName
	}
	return ""
}

func (x *UpdateManager) GetPhone() string {
	if x != nil {
		return x.Phone
	}
	return ""
}

func (x *UpdateManager) GetSalary() float64 {
	if x != nil {
		return x.Salary
	}
	return 0
}

func (x *UpdateManager) GetPassword() string {
	if x != nil {
		return x.Password
	}
	return ""
}

func (x *UpdateManager) GetLogin() string {
	if x != nil {
		return x.Login
	}
	return ""
}

func (x *UpdateManager) GetBranchId() string {
	if x != nil {
		return x.BranchId
	}
	return ""
}

type GetListManagerRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Offset int64  `protobuf:"varint,1,opt,name=offset,proto3" json:"offset,omitempty"`
	Limit  int64  `protobuf:"varint,2,opt,name=limit,proto3" json:"limit,omitempty"`
	Search string `protobuf:"bytes,3,opt,name=search,proto3" json:"search,omitempty"`
}

func (x *GetListManagerRequest) Reset() {
	*x = GetListManagerRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_manager_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetListManagerRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetListManagerRequest) ProtoMessage() {}

func (x *GetListManagerRequest) ProtoReflect() protoreflect.Message {
	mi := &file_manager_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetListManagerRequest.ProtoReflect.Descriptor instead.
func (*GetListManagerRequest) Descriptor() ([]byte, []int) {
	return file_manager_proto_rawDescGZIP(), []int{5}
}

func (x *GetListManagerRequest) GetOffset() int64 {
	if x != nil {
		return x.Offset
	}
	return 0
}

func (x *GetListManagerRequest) GetLimit() int64 {
	if x != nil {
		return x.Limit
	}
	return 0
}

func (x *GetListManagerRequest) GetSearch() string {
	if x != nil {
		return x.Search
	}
	return ""
}

type GetListManagerResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Count    int64      `protobuf:"varint,1,opt,name=count,proto3" json:"count,omitempty"`
	Managers []*Manager `protobuf:"bytes,2,rep,name=Managers,proto3" json:"Managers,omitempty"`
}

func (x *GetListManagerResponse) Reset() {
	*x = GetListManagerResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_manager_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetListManagerResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetListManagerResponse) ProtoMessage() {}

func (x *GetListManagerResponse) ProtoReflect() protoreflect.Message {
	mi := &file_manager_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetListManagerResponse.ProtoReflect.Descriptor instead.
func (*GetListManagerResponse) Descriptor() ([]byte, []int) {
	return file_manager_proto_rawDescGZIP(), []int{6}
}

func (x *GetListManagerResponse) GetCount() int64 {
	if x != nil {
		return x.Count
	}
	return 0
}

func (x *GetListManagerResponse) GetManagers() []*Manager {
	if x != nil {
		return x.Managers
	}
	return nil
}

var File_manager_proto protoreflect.FileDescriptor

var file_manager_proto_rawDesc = []byte{
	0x0a, 0x0d, 0x6d, 0x61, 0x6e, 0x61, 0x67, 0x65, 0x72, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12,
	0x0c, 0x75, 0x73, 0x65, 0x72, 0x5f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x22, 0x0e, 0x0a,
	0x0c, 0x4d, 0x61, 0x6e, 0x61, 0x67, 0x65, 0x72, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x22, 0x23, 0x0a,
	0x11, 0x4d, 0x61, 0x6e, 0x61, 0x67, 0x65, 0x72, 0x50, 0x72, 0x69, 0x6d, 0x61, 0x72, 0x79, 0x4b,
	0x65, 0x79, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02,
	0x69, 0x64, 0x22, 0x93, 0x01, 0x0a, 0x0d, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x4d, 0x61, 0x6e,
	0x61, 0x67, 0x65, 0x72, 0x12, 0x1b, 0x0a, 0x09, 0x66, 0x75, 0x6c, 0x6c, 0x5f, 0x6e, 0x61, 0x6d,
	0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x66, 0x75, 0x6c, 0x6c, 0x4e, 0x61, 0x6d,
	0x65, 0x12, 0x14, 0x0a, 0x05, 0x70, 0x68, 0x6f, 0x6e, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x05, 0x70, 0x68, 0x6f, 0x6e, 0x65, 0x12, 0x16, 0x0a, 0x06, 0x73, 0x61, 0x6c, 0x61, 0x72,
	0x79, 0x18, 0x03, 0x20, 0x01, 0x28, 0x01, 0x52, 0x06, 0x73, 0x61, 0x6c, 0x61, 0x72, 0x79, 0x12,
	0x1a, 0x0a, 0x08, 0x70, 0x61, 0x73, 0x73, 0x77, 0x6f, 0x72, 0x64, 0x18, 0x04, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x08, 0x70, 0x61, 0x73, 0x73, 0x77, 0x6f, 0x72, 0x64, 0x12, 0x1b, 0x0a, 0x09, 0x62,
	0x72, 0x61, 0x6e, 0x63, 0x68, 0x5f, 0x69, 0x64, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08,
	0x62, 0x72, 0x61, 0x6e, 0x63, 0x68, 0x49, 0x64, 0x22, 0x8a, 0x02, 0x0a, 0x07, 0x4d, 0x61, 0x6e,
	0x61, 0x67, 0x65, 0x72, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x02, 0x69, 0x64, 0x12, 0x1b, 0x0a, 0x09, 0x66, 0x75, 0x6c, 0x6c, 0x5f, 0x6e, 0x61, 0x6d,
	0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x66, 0x75, 0x6c, 0x6c, 0x4e, 0x61, 0x6d,
	0x65, 0x12, 0x14, 0x0a, 0x05, 0x70, 0x68, 0x6f, 0x6e, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x05, 0x70, 0x68, 0x6f, 0x6e, 0x65, 0x12, 0x16, 0x0a, 0x06, 0x73, 0x61, 0x6c, 0x61, 0x72,
	0x79, 0x18, 0x04, 0x20, 0x01, 0x28, 0x01, 0x52, 0x06, 0x73, 0x61, 0x6c, 0x61, 0x72, 0x79, 0x12,
	0x1a, 0x0a, 0x08, 0x70, 0x61, 0x73, 0x73, 0x77, 0x6f, 0x72, 0x64, 0x18, 0x05, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x08, 0x70, 0x61, 0x73, 0x73, 0x77, 0x6f, 0x72, 0x64, 0x12, 0x14, 0x0a, 0x05, 0x6c,
	0x6f, 0x67, 0x69, 0x6e, 0x18, 0x06, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x6c, 0x6f, 0x67, 0x69,
	0x6e, 0x12, 0x1b, 0x0a, 0x09, 0x62, 0x72, 0x61, 0x6e, 0x63, 0x68, 0x5f, 0x69, 0x64, 0x18, 0x07,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x62, 0x72, 0x61, 0x6e, 0x63, 0x68, 0x49, 0x64, 0x12, 0x17,
	0x0a, 0x07, 0x72, 0x6f, 0x6c, 0x65, 0x5f, 0x69, 0x64, 0x18, 0x08, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x06, 0x72, 0x6f, 0x6c, 0x65, 0x49, 0x64, 0x12, 0x1d, 0x0a, 0x0a, 0x63, 0x72, 0x65, 0x61, 0x74,
	0x65, 0x64, 0x5f, 0x61, 0x74, 0x18, 0x09, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x63, 0x72, 0x65,
	0x61, 0x74, 0x65, 0x64, 0x41, 0x74, 0x12, 0x1d, 0x0a, 0x0a, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65,
	0x64, 0x5f, 0x61, 0x74, 0x18, 0x0a, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x75, 0x70, 0x64, 0x61,
	0x74, 0x65, 0x64, 0x41, 0x74, 0x22, 0xb9, 0x01, 0x0a, 0x0d, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65,
	0x4d, 0x61, 0x6e, 0x61, 0x67, 0x65, 0x72, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x12, 0x1b, 0x0a, 0x09, 0x66, 0x75, 0x6c, 0x6c, 0x5f,
	0x6e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x66, 0x75, 0x6c, 0x6c,
	0x4e, 0x61, 0x6d, 0x65, 0x12, 0x14, 0x0a, 0x05, 0x70, 0x68, 0x6f, 0x6e, 0x65, 0x18, 0x03, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x05, 0x70, 0x68, 0x6f, 0x6e, 0x65, 0x12, 0x16, 0x0a, 0x06, 0x73, 0x61,
	0x6c, 0x61, 0x72, 0x79, 0x18, 0x04, 0x20, 0x01, 0x28, 0x01, 0x52, 0x06, 0x73, 0x61, 0x6c, 0x61,
	0x72, 0x79, 0x12, 0x1a, 0x0a, 0x08, 0x70, 0x61, 0x73, 0x73, 0x77, 0x6f, 0x72, 0x64, 0x18, 0x05,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x70, 0x61, 0x73, 0x73, 0x77, 0x6f, 0x72, 0x64, 0x12, 0x14,
	0x0a, 0x05, 0x6c, 0x6f, 0x67, 0x69, 0x6e, 0x18, 0x06, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x6c,
	0x6f, 0x67, 0x69, 0x6e, 0x12, 0x1b, 0x0a, 0x09, 0x62, 0x72, 0x61, 0x6e, 0x63, 0x68, 0x5f, 0x69,
	0x64, 0x18, 0x07, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x62, 0x72, 0x61, 0x6e, 0x63, 0x68, 0x49,
	0x64, 0x22, 0x5d, 0x0a, 0x15, 0x47, 0x65, 0x74, 0x4c, 0x69, 0x73, 0x74, 0x4d, 0x61, 0x6e, 0x61,
	0x67, 0x65, 0x72, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x16, 0x0a, 0x06, 0x6f, 0x66,
	0x66, 0x73, 0x65, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x06, 0x6f, 0x66, 0x66, 0x73,
	0x65, 0x74, 0x12, 0x14, 0x0a, 0x05, 0x6c, 0x69, 0x6d, 0x69, 0x74, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x03, 0x52, 0x05, 0x6c, 0x69, 0x6d, 0x69, 0x74, 0x12, 0x16, 0x0a, 0x06, 0x73, 0x65, 0x61, 0x72,
	0x63, 0x68, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x73, 0x65, 0x61, 0x72, 0x63, 0x68,
	0x22, 0x61, 0x0a, 0x16, 0x47, 0x65, 0x74, 0x4c, 0x69, 0x73, 0x74, 0x4d, 0x61, 0x6e, 0x61, 0x67,
	0x65, 0x72, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x14, 0x0a, 0x05, 0x63, 0x6f,
	0x75, 0x6e, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x05, 0x63, 0x6f, 0x75, 0x6e, 0x74,
	0x12, 0x31, 0x0a, 0x08, 0x4d, 0x61, 0x6e, 0x61, 0x67, 0x65, 0x72, 0x73, 0x18, 0x02, 0x20, 0x03,
	0x28, 0x0b, 0x32, 0x15, 0x2e, 0x75, 0x73, 0x65, 0x72, 0x5f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63,
	0x65, 0x2e, 0x4d, 0x61, 0x6e, 0x61, 0x67, 0x65, 0x72, 0x52, 0x08, 0x4d, 0x61, 0x6e, 0x61, 0x67,
	0x65, 0x72, 0x73, 0x32, 0xf6, 0x02, 0x0a, 0x0e, 0x4d, 0x61, 0x6e, 0x61, 0x67, 0x65, 0x72, 0x53,
	0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x3e, 0x0a, 0x06, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65,
	0x12, 0x1b, 0x2e, 0x75, 0x73, 0x65, 0x72, 0x5f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e,
	0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x4d, 0x61, 0x6e, 0x61, 0x67, 0x65, 0x72, 0x1a, 0x15, 0x2e,
	0x75, 0x73, 0x65, 0x72, 0x5f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x4d, 0x61, 0x6e,
	0x61, 0x67, 0x65, 0x72, 0x22, 0x00, 0x12, 0x43, 0x0a, 0x07, 0x47, 0x65, 0x74, 0x42, 0x79, 0x49,
	0x44, 0x12, 0x1f, 0x2e, 0x75, 0x73, 0x65, 0x72, 0x5f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65,
	0x2e, 0x4d, 0x61, 0x6e, 0x61, 0x67, 0x65, 0x72, 0x50, 0x72, 0x69, 0x6d, 0x61, 0x72, 0x79, 0x4b,
	0x65, 0x79, 0x1a, 0x15, 0x2e, 0x75, 0x73, 0x65, 0x72, 0x5f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63,
	0x65, 0x2e, 0x4d, 0x61, 0x6e, 0x61, 0x67, 0x65, 0x72, 0x22, 0x00, 0x12, 0x56, 0x0a, 0x07, 0x47,
	0x65, 0x74, 0x4c, 0x69, 0x73, 0x74, 0x12, 0x23, 0x2e, 0x75, 0x73, 0x65, 0x72, 0x5f, 0x73, 0x65,
	0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x47, 0x65, 0x74, 0x4c, 0x69, 0x73, 0x74, 0x4d, 0x61, 0x6e,
	0x61, 0x67, 0x65, 0x72, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x24, 0x2e, 0x75, 0x73,
	0x65, 0x72, 0x5f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x47, 0x65, 0x74, 0x4c, 0x69,
	0x73, 0x74, 0x4d, 0x61, 0x6e, 0x61, 0x67, 0x65, 0x72, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73,
	0x65, 0x22, 0x00, 0x12, 0x3e, 0x0a, 0x06, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x12, 0x1b, 0x2e,
	0x75, 0x73, 0x65, 0x72, 0x5f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x55, 0x70, 0x64,
	0x61, 0x74, 0x65, 0x4d, 0x61, 0x6e, 0x61, 0x67, 0x65, 0x72, 0x1a, 0x15, 0x2e, 0x75, 0x73, 0x65,
	0x72, 0x5f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x4d, 0x61, 0x6e, 0x61, 0x67, 0x65,
	0x72, 0x22, 0x00, 0x12, 0x47, 0x0a, 0x06, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x12, 0x1f, 0x2e,
	0x75, 0x73, 0x65, 0x72, 0x5f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x4d, 0x61, 0x6e,
	0x61, 0x67, 0x65, 0x72, 0x50, 0x72, 0x69, 0x6d, 0x61, 0x72, 0x79, 0x4b, 0x65, 0x79, 0x1a, 0x1a,
	0x2e, 0x75, 0x73, 0x65, 0x72, 0x5f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x4d, 0x61,
	0x6e, 0x61, 0x67, 0x65, 0x72, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x22, 0x00, 0x42, 0x17, 0x5a, 0x15,
	0x67, 0x65, 0x6e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x75, 0x73, 0x65, 0x72, 0x5f, 0x73, 0x65,
	0x72, 0x76, 0x69, 0x63, 0x65, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_manager_proto_rawDescOnce sync.Once
	file_manager_proto_rawDescData = file_manager_proto_rawDesc
)

func file_manager_proto_rawDescGZIP() []byte {
	file_manager_proto_rawDescOnce.Do(func() {
		file_manager_proto_rawDescData = protoimpl.X.CompressGZIP(file_manager_proto_rawDescData)
	})
	return file_manager_proto_rawDescData
}

var file_manager_proto_msgTypes = make([]protoimpl.MessageInfo, 7)
var file_manager_proto_goTypes = []interface{}{
	(*ManagerEmpty)(nil),           // 0: user_service.ManagerEmpty
	(*ManagerPrimaryKey)(nil),      // 1: user_service.ManagerPrimaryKey
	(*CreateManager)(nil),          // 2: user_service.CreateManager
	(*Manager)(nil),                // 3: user_service.Manager
	(*UpdateManager)(nil),          // 4: user_service.UpdateManager
	(*GetListManagerRequest)(nil),  // 5: user_service.GetListManagerRequest
	(*GetListManagerResponse)(nil), // 6: user_service.GetListManagerResponse
}
var file_manager_proto_depIdxs = []int32{
	3, // 0: user_service.GetListManagerResponse.Managers:type_name -> user_service.Manager
	2, // 1: user_service.ManagerService.Create:input_type -> user_service.CreateManager
	1, // 2: user_service.ManagerService.GetByID:input_type -> user_service.ManagerPrimaryKey
	5, // 3: user_service.ManagerService.GetList:input_type -> user_service.GetListManagerRequest
	4, // 4: user_service.ManagerService.Update:input_type -> user_service.UpdateManager
	1, // 5: user_service.ManagerService.Delete:input_type -> user_service.ManagerPrimaryKey
	3, // 6: user_service.ManagerService.Create:output_type -> user_service.Manager
	3, // 7: user_service.ManagerService.GetByID:output_type -> user_service.Manager
	6, // 8: user_service.ManagerService.GetList:output_type -> user_service.GetListManagerResponse
	3, // 9: user_service.ManagerService.Update:output_type -> user_service.Manager
	0, // 10: user_service.ManagerService.Delete:output_type -> user_service.ManagerEmpty
	6, // [6:11] is the sub-list for method output_type
	1, // [1:6] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_manager_proto_init() }
func file_manager_proto_init() {
	if File_manager_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_manager_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ManagerEmpty); i {
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
		file_manager_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ManagerPrimaryKey); i {
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
		file_manager_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CreateManager); i {
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
		file_manager_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Manager); i {
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
		file_manager_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UpdateManager); i {
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
		file_manager_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetListManagerRequest); i {
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
		file_manager_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetListManagerResponse); i {
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
			RawDescriptor: file_manager_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   7,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_manager_proto_goTypes,
		DependencyIndexes: file_manager_proto_depIdxs,
		MessageInfos:      file_manager_proto_msgTypes,
	}.Build()
	File_manager_proto = out.File
	file_manager_proto_rawDesc = nil
	file_manager_proto_goTypes = nil
	file_manager_proto_depIdxs = nil
}
