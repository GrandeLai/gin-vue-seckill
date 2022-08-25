// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v3.13.0
// source: admin_user.proto

package admin_user

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

type AdminUserRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Username string `protobuf:"bytes,1,opt,name=username,proto3" json:"username,omitempty"`
	Password string `protobuf:"bytes,2,opt,name=password,proto3" json:"password,omitempty"`
}

func (x *AdminUserRequest) Reset() {
	*x = AdminUserRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_admin_user_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AdminUserRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AdminUserRequest) ProtoMessage() {}

func (x *AdminUserRequest) ProtoReflect() protoreflect.Message {
	mi := &file_admin_user_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AdminUserRequest.ProtoReflect.Descriptor instead.
func (*AdminUserRequest) Descriptor() ([]byte, []int) {
	return file_admin_user_proto_rawDescGZIP(), []int{0}
}

func (x *AdminUserRequest) GetUsername() string {
	if x != nil {
		return x.Username
	}
	return ""
}

func (x *AdminUserRequest) GetPassword() string {
	if x != nil {
		return x.Password
	}
	return ""
}

type AdminUserResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Code     int32  `protobuf:"varint,1,opt,name=code,proto3" json:"code,omitempty"`
	Msg      string `protobuf:"bytes,2,opt,name=msg,proto3" json:"msg,omitempty"`
	UserName string `protobuf:"bytes,3,opt,name=user_name,json=userName,proto3" json:"user_name,omitempty"`
}

func (x *AdminUserResponse) Reset() {
	*x = AdminUserResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_admin_user_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AdminUserResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AdminUserResponse) ProtoMessage() {}

func (x *AdminUserResponse) ProtoReflect() protoreflect.Message {
	mi := &file_admin_user_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AdminUserResponse.ProtoReflect.Descriptor instead.
func (*AdminUserResponse) Descriptor() ([]byte, []int) {
	return file_admin_user_proto_rawDescGZIP(), []int{1}
}

func (x *AdminUserResponse) GetCode() int32 {
	if x != nil {
		return x.Code
	}
	return 0
}

func (x *AdminUserResponse) GetMsg() string {
	if x != nil {
		return x.Msg
	}
	return ""
}

func (x *AdminUserResponse) GetUserName() string {
	if x != nil {
		return x.UserName
	}
	return ""
}

type FrontUserListRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	CurrentPage int32 `protobuf:"varint,1,opt,name=currentPage,proto3" json:"currentPage,omitempty"`
	PageSize    int32 `protobuf:"varint,2,opt,name=pageSize,proto3" json:"pageSize,omitempty"`
}

func (x *FrontUserListRequest) Reset() {
	*x = FrontUserListRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_admin_user_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *FrontUserListRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*FrontUserListRequest) ProtoMessage() {}

func (x *FrontUserListRequest) ProtoReflect() protoreflect.Message {
	mi := &file_admin_user_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use FrontUserListRequest.ProtoReflect.Descriptor instead.
func (*FrontUserListRequest) Descriptor() ([]byte, []int) {
	return file_admin_user_proto_rawDescGZIP(), []int{2}
}

func (x *FrontUserListRequest) GetCurrentPage() int32 {
	if x != nil {
		return x.CurrentPage
	}
	return 0
}

func (x *FrontUserListRequest) GetPageSize() int32 {
	if x != nil {
		return x.PageSize
	}
	return 0
}

type FrontUserListResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Code          int32              `protobuf:"varint,1,opt,name=code,proto3" json:"code,omitempty"`
	Msg           string             `protobuf:"bytes,2,opt,name=msg,proto3" json:"msg,omitempty"`
	Frontuserlist []*FrontUserDetail `protobuf:"bytes,3,rep,name=frontuserlist,proto3" json:"frontuserlist,omitempty"`
	Total         int32              `protobuf:"varint,4,opt,name=total,proto3" json:"total,omitempty"`
	Current       int32              `protobuf:"varint,5,opt,name=current,proto3" json:"current,omitempty"`
	PageSize      int32              `protobuf:"varint,6,opt,name=page_size,json=pageSize,proto3" json:"page_size,omitempty"`
}

func (x *FrontUserListResponse) Reset() {
	*x = FrontUserListResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_admin_user_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *FrontUserListResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*FrontUserListResponse) ProtoMessage() {}

func (x *FrontUserListResponse) ProtoReflect() protoreflect.Message {
	mi := &file_admin_user_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use FrontUserListResponse.ProtoReflect.Descriptor instead.
func (*FrontUserListResponse) Descriptor() ([]byte, []int) {
	return file_admin_user_proto_rawDescGZIP(), []int{3}
}

func (x *FrontUserListResponse) GetCode() int32 {
	if x != nil {
		return x.Code
	}
	return 0
}

func (x *FrontUserListResponse) GetMsg() string {
	if x != nil {
		return x.Msg
	}
	return ""
}

func (x *FrontUserListResponse) GetFrontuserlist() []*FrontUserDetail {
	if x != nil {
		return x.Frontuserlist
	}
	return nil
}

func (x *FrontUserListResponse) GetTotal() int32 {
	if x != nil {
		return x.Total
	}
	return 0
}

func (x *FrontUserListResponse) GetCurrent() int32 {
	if x != nil {
		return x.Current
	}
	return 0
}

func (x *FrontUserListResponse) GetPageSize() int32 {
	if x != nil {
		return x.PageSize
	}
	return 0
}

type FrontUserDetail struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Email      string `protobuf:"bytes,1,opt,name=email,proto3" json:"email,omitempty"`
	Desc       string `protobuf:"bytes,2,opt,name=desc,proto3" json:"desc,omitempty"`
	Status     string `protobuf:"bytes,3,opt,name=status,proto3" json:"status,omitempty"`
	CreateTime string `protobuf:"bytes,4,opt,name=create_time,json=createTime,proto3" json:"create_time,omitempty"`
}

func (x *FrontUserDetail) Reset() {
	*x = FrontUserDetail{}
	if protoimpl.UnsafeEnabled {
		mi := &file_admin_user_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *FrontUserDetail) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*FrontUserDetail) ProtoMessage() {}

func (x *FrontUserDetail) ProtoReflect() protoreflect.Message {
	mi := &file_admin_user_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use FrontUserDetail.ProtoReflect.Descriptor instead.
func (*FrontUserDetail) Descriptor() ([]byte, []int) {
	return file_admin_user_proto_rawDescGZIP(), []int{4}
}

func (x *FrontUserDetail) GetEmail() string {
	if x != nil {
		return x.Email
	}
	return ""
}

func (x *FrontUserDetail) GetDesc() string {
	if x != nil {
		return x.Desc
	}
	return ""
}

func (x *FrontUserDetail) GetStatus() string {
	if x != nil {
		return x.Status
	}
	return ""
}

func (x *FrontUserDetail) GetCreateTime() string {
	if x != nil {
		return x.CreateTime
	}
	return ""
}

var File_admin_user_proto protoreflect.FileDescriptor

var file_admin_user_proto_rawDesc = []byte{
	0x0a, 0x10, 0x61, 0x64, 0x6d, 0x69, 0x6e, 0x5f, 0x75, 0x73, 0x65, 0x72, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x12, 0x0d, 0x6a, 0x64, 0x61, 0x77, 0x2e, 0x75, 0x73, 0x65, 0x72, 0x2e, 0x73, 0x72,
	0x76, 0x22, 0x4a, 0x0a, 0x10, 0x41, 0x64, 0x6d, 0x69, 0x6e, 0x55, 0x73, 0x65, 0x72, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x1a, 0x0a, 0x08, 0x75, 0x73, 0x65, 0x72, 0x6e, 0x61, 0x6d,
	0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x75, 0x73, 0x65, 0x72, 0x6e, 0x61, 0x6d,
	0x65, 0x12, 0x1a, 0x0a, 0x08, 0x70, 0x61, 0x73, 0x73, 0x77, 0x6f, 0x72, 0x64, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x08, 0x70, 0x61, 0x73, 0x73, 0x77, 0x6f, 0x72, 0x64, 0x22, 0x56, 0x0a,
	0x11, 0x41, 0x64, 0x6d, 0x69, 0x6e, 0x55, 0x73, 0x65, 0x72, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e,
	0x73, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x63, 0x6f, 0x64, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05,
	0x52, 0x04, 0x63, 0x6f, 0x64, 0x65, 0x12, 0x10, 0x0a, 0x03, 0x6d, 0x73, 0x67, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x03, 0x6d, 0x73, 0x67, 0x12, 0x1b, 0x0a, 0x09, 0x75, 0x73, 0x65, 0x72,
	0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x75, 0x73, 0x65,
	0x72, 0x4e, 0x61, 0x6d, 0x65, 0x22, 0x54, 0x0a, 0x14, 0x46, 0x72, 0x6f, 0x6e, 0x74, 0x55, 0x73,
	0x65, 0x72, 0x4c, 0x69, 0x73, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x20, 0x0a,
	0x0b, 0x63, 0x75, 0x72, 0x72, 0x65, 0x6e, 0x74, 0x50, 0x61, 0x67, 0x65, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x05, 0x52, 0x0b, 0x63, 0x75, 0x72, 0x72, 0x65, 0x6e, 0x74, 0x50, 0x61, 0x67, 0x65, 0x12,
	0x1a, 0x0a, 0x08, 0x70, 0x61, 0x67, 0x65, 0x53, 0x69, 0x7a, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x05, 0x52, 0x08, 0x70, 0x61, 0x67, 0x65, 0x53, 0x69, 0x7a, 0x65, 0x22, 0xd0, 0x01, 0x0a, 0x15,
	0x46, 0x72, 0x6f, 0x6e, 0x74, 0x55, 0x73, 0x65, 0x72, 0x4c, 0x69, 0x73, 0x74, 0x52, 0x65, 0x73,
	0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x63, 0x6f, 0x64, 0x65, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x05, 0x52, 0x04, 0x63, 0x6f, 0x64, 0x65, 0x12, 0x10, 0x0a, 0x03, 0x6d, 0x73, 0x67,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x6d, 0x73, 0x67, 0x12, 0x44, 0x0a, 0x0d, 0x66,
	0x72, 0x6f, 0x6e, 0x74, 0x75, 0x73, 0x65, 0x72, 0x6c, 0x69, 0x73, 0x74, 0x18, 0x03, 0x20, 0x03,
	0x28, 0x0b, 0x32, 0x1e, 0x2e, 0x6a, 0x64, 0x61, 0x77, 0x2e, 0x75, 0x73, 0x65, 0x72, 0x2e, 0x73,
	0x72, 0x76, 0x2e, 0x46, 0x72, 0x6f, 0x6e, 0x74, 0x55, 0x73, 0x65, 0x72, 0x44, 0x65, 0x74, 0x61,
	0x69, 0x6c, 0x52, 0x0d, 0x66, 0x72, 0x6f, 0x6e, 0x74, 0x75, 0x73, 0x65, 0x72, 0x6c, 0x69, 0x73,
	0x74, 0x12, 0x14, 0x0a, 0x05, 0x74, 0x6f, 0x74, 0x61, 0x6c, 0x18, 0x04, 0x20, 0x01, 0x28, 0x05,
	0x52, 0x05, 0x74, 0x6f, 0x74, 0x61, 0x6c, 0x12, 0x18, 0x0a, 0x07, 0x63, 0x75, 0x72, 0x72, 0x65,
	0x6e, 0x74, 0x18, 0x05, 0x20, 0x01, 0x28, 0x05, 0x52, 0x07, 0x63, 0x75, 0x72, 0x72, 0x65, 0x6e,
	0x74, 0x12, 0x1b, 0x0a, 0x09, 0x70, 0x61, 0x67, 0x65, 0x5f, 0x73, 0x69, 0x7a, 0x65, 0x18, 0x06,
	0x20, 0x01, 0x28, 0x05, 0x52, 0x08, 0x70, 0x61, 0x67, 0x65, 0x53, 0x69, 0x7a, 0x65, 0x22, 0x74,
	0x0a, 0x0f, 0x46, 0x72, 0x6f, 0x6e, 0x74, 0x55, 0x73, 0x65, 0x72, 0x44, 0x65, 0x74, 0x61, 0x69,
	0x6c, 0x12, 0x14, 0x0a, 0x05, 0x65, 0x6d, 0x61, 0x69, 0x6c, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x05, 0x65, 0x6d, 0x61, 0x69, 0x6c, 0x12, 0x12, 0x0a, 0x04, 0x64, 0x65, 0x73, 0x63, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x64, 0x65, 0x73, 0x63, 0x12, 0x16, 0x0a, 0x06, 0x73,
	0x74, 0x61, 0x74, 0x75, 0x73, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x73, 0x74, 0x61,
	0x74, 0x75, 0x73, 0x12, 0x1f, 0x0a, 0x0b, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x5f, 0x74, 0x69,
	0x6d, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65,
	0x54, 0x69, 0x6d, 0x65, 0x32, 0xc3, 0x01, 0x0a, 0x09, 0x41, 0x64, 0x6d, 0x69, 0x6e, 0x55, 0x73,
	0x65, 0x72, 0x12, 0x55, 0x0a, 0x0e, 0x41, 0x64, 0x6d, 0x69, 0x6e, 0x55, 0x73, 0x65, 0x72, 0x4c,
	0x6f, 0x67, 0x69, 0x6e, 0x12, 0x1f, 0x2e, 0x6a, 0x64, 0x61, 0x77, 0x2e, 0x75, 0x73, 0x65, 0x72,
	0x2e, 0x73, 0x72, 0x76, 0x2e, 0x41, 0x64, 0x6d, 0x69, 0x6e, 0x55, 0x73, 0x65, 0x72, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x20, 0x2e, 0x6a, 0x64, 0x61, 0x77, 0x2e, 0x75, 0x73, 0x65,
	0x72, 0x2e, 0x73, 0x72, 0x76, 0x2e, 0x41, 0x64, 0x6d, 0x69, 0x6e, 0x55, 0x73, 0x65, 0x72, 0x52,
	0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x5f, 0x0a, 0x10, 0x47, 0x65, 0x74,
	0x46, 0x72, 0x6f, 0x6e, 0x74, 0x55, 0x73, 0x65, 0x72, 0x4c, 0x69, 0x73, 0x74, 0x12, 0x23, 0x2e,
	0x6a, 0x64, 0x61, 0x77, 0x2e, 0x75, 0x73, 0x65, 0x72, 0x2e, 0x73, 0x72, 0x76, 0x2e, 0x46, 0x72,
	0x6f, 0x6e, 0x74, 0x55, 0x73, 0x65, 0x72, 0x4c, 0x69, 0x73, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x1a, 0x24, 0x2e, 0x6a, 0x64, 0x61, 0x77, 0x2e, 0x75, 0x73, 0x65, 0x72, 0x2e, 0x73,
	0x72, 0x76, 0x2e, 0x46, 0x72, 0x6f, 0x6e, 0x74, 0x55, 0x73, 0x65, 0x72, 0x4c, 0x69, 0x73, 0x74,
	0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x42, 0x17, 0x5a, 0x15, 0x2e, 0x2f,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x3b, 0x6a, 0x64, 0x61, 0x77, 0x5f, 0x75, 0x73, 0x65, 0x72, 0x5f,
	0x73, 0x72, 0x76, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_admin_user_proto_rawDescOnce sync.Once
	file_admin_user_proto_rawDescData = file_admin_user_proto_rawDesc
)

func file_admin_user_proto_rawDescGZIP() []byte {
	file_admin_user_proto_rawDescOnce.Do(func() {
		file_admin_user_proto_rawDescData = protoimpl.X.CompressGZIP(file_admin_user_proto_rawDescData)
	})
	return file_admin_user_proto_rawDescData
}

var file_admin_user_proto_msgTypes = make([]protoimpl.MessageInfo, 5)
var file_admin_user_proto_goTypes = []interface{}{
	(*AdminUserRequest)(nil),      // 0: jdaw.user.srv.AdminUserRequest
	(*AdminUserResponse)(nil),     // 1: jdaw.user.srv.AdminUserResponse
	(*FrontUserListRequest)(nil),  // 2: jdaw.user.srv.FrontUserListRequest
	(*FrontUserListResponse)(nil), // 3: jdaw.user.srv.FrontUserListResponse
	(*FrontUserDetail)(nil),       // 4: jdaw.user.srv.FrontUserDetail
}
var file_admin_user_proto_depIdxs = []int32{
	4, // 0: jdaw.user.srv.FrontUserListResponse.frontuserlist:type_name -> jdaw.user.srv.FrontUserDetail
	0, // 1: jdaw.user.srv.AdminUser.AdminUserLogin:input_type -> jdaw.user.srv.AdminUserRequest
	2, // 2: jdaw.user.srv.AdminUser.GetFrontUserList:input_type -> jdaw.user.srv.FrontUserListRequest
	1, // 3: jdaw.user.srv.AdminUser.AdminUserLogin:output_type -> jdaw.user.srv.AdminUserResponse
	3, // 4: jdaw.user.srv.AdminUser.GetFrontUserList:output_type -> jdaw.user.srv.FrontUserListResponse
	3, // [3:5] is the sub-list for method output_type
	1, // [1:3] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_admin_user_proto_init() }
func file_admin_user_proto_init() {
	if File_admin_user_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_admin_user_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AdminUserRequest); i {
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
		file_admin_user_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AdminUserResponse); i {
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
		file_admin_user_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*FrontUserListRequest); i {
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
		file_admin_user_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*FrontUserListResponse); i {
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
		file_admin_user_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*FrontUserDetail); i {
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
			RawDescriptor: file_admin_user_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   5,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_admin_user_proto_goTypes,
		DependencyIndexes: file_admin_user_proto_depIdxs,
		MessageInfos:      file_admin_user_proto_msgTypes,
	}.Build()
	File_admin_user_proto = out.File
	file_admin_user_proto_rawDesc = nil
	file_admin_user_proto_goTypes = nil
	file_admin_user_proto_depIdxs = nil
}