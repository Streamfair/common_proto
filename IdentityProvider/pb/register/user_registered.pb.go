// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.32.0
// 	protoc        v4.25.2
// source: register/user_registered.proto

package user_registered

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

type UserRegistered struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id                int64                  `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	Username          string                 `protobuf:"bytes,2,opt,name=username,proto3" json:"username,omitempty"`
	FullName          string                 `protobuf:"bytes,3,opt,name=full_name,json=fullName,proto3" json:"full_name,omitempty"`
	Email             string                 `protobuf:"bytes,4,opt,name=email,proto3" json:"email,omitempty"`
	PasswordHash      string                 `protobuf:"bytes,5,opt,name=password_hash,json=passwordHash,proto3" json:"password_hash,omitempty"`
	PasswordSalt      string                 `protobuf:"bytes,6,opt,name=password_salt,json=passwordSalt,proto3" json:"password_salt,omitempty"`
	CountryCode       string                 `protobuf:"bytes,7,opt,name=country_code,json=countryCode,proto3" json:"country_code,omitempty"`
	RoleId            int64                  `protobuf:"varint,8,opt,name=role_id,json=roleId,proto3" json:"role_id,omitempty"`
	Status            string                 `protobuf:"bytes,9,opt,name=status,proto3" json:"status,omitempty"`
	LastLoginAt       *timestamppb.Timestamp `protobuf:"bytes,10,opt,name=last_login_at,json=lastLoginAt,proto3" json:"last_login_at,omitempty"`
	UsernameChangedAt *timestamppb.Timestamp `protobuf:"bytes,11,opt,name=username_changed_at,json=usernameChangedAt,proto3" json:"username_changed_at,omitempty"`
	EmailChangedAt    *timestamppb.Timestamp `protobuf:"bytes,12,opt,name=email_changed_at,json=emailChangedAt,proto3" json:"email_changed_at,omitempty"`
	PasswordChangedAt *timestamppb.Timestamp `protobuf:"bytes,13,opt,name=password_changed_at,json=passwordChangedAt,proto3" json:"password_changed_at,omitempty"`
	CreatedAt         *timestamppb.Timestamp `protobuf:"bytes,14,opt,name=created_at,json=createdAt,proto3" json:"created_at,omitempty"`
	UpdatedAt         *timestamppb.Timestamp `protobuf:"bytes,15,opt,name=updated_at,json=updatedAt,proto3" json:"updated_at,omitempty"`
}

func (x *UserRegistered) Reset() {
	*x = UserRegistered{}
	if protoimpl.UnsafeEnabled {
		mi := &file_register_user_registered_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UserRegistered) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UserRegistered) ProtoMessage() {}

func (x *UserRegistered) ProtoReflect() protoreflect.Message {
	mi := &file_register_user_registered_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UserRegistered.ProtoReflect.Descriptor instead.
func (*UserRegistered) Descriptor() ([]byte, []int) {
	return file_register_user_registered_proto_rawDescGZIP(), []int{0}
}

func (x *UserRegistered) GetId() int64 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *UserRegistered) GetUsername() string {
	if x != nil {
		return x.Username
	}
	return ""
}

func (x *UserRegistered) GetFullName() string {
	if x != nil {
		return x.FullName
	}
	return ""
}

func (x *UserRegistered) GetEmail() string {
	if x != nil {
		return x.Email
	}
	return ""
}

func (x *UserRegistered) GetPasswordHash() string {
	if x != nil {
		return x.PasswordHash
	}
	return ""
}

func (x *UserRegistered) GetPasswordSalt() string {
	if x != nil {
		return x.PasswordSalt
	}
	return ""
}

func (x *UserRegistered) GetCountryCode() string {
	if x != nil {
		return x.CountryCode
	}
	return ""
}

func (x *UserRegistered) GetRoleId() int64 {
	if x != nil {
		return x.RoleId
	}
	return 0
}

func (x *UserRegistered) GetStatus() string {
	if x != nil {
		return x.Status
	}
	return ""
}

func (x *UserRegistered) GetLastLoginAt() *timestamppb.Timestamp {
	if x != nil {
		return x.LastLoginAt
	}
	return nil
}

func (x *UserRegistered) GetUsernameChangedAt() *timestamppb.Timestamp {
	if x != nil {
		return x.UsernameChangedAt
	}
	return nil
}

func (x *UserRegistered) GetEmailChangedAt() *timestamppb.Timestamp {
	if x != nil {
		return x.EmailChangedAt
	}
	return nil
}

func (x *UserRegistered) GetPasswordChangedAt() *timestamppb.Timestamp {
	if x != nil {
		return x.PasswordChangedAt
	}
	return nil
}

func (x *UserRegistered) GetCreatedAt() *timestamppb.Timestamp {
	if x != nil {
		return x.CreatedAt
	}
	return nil
}

func (x *UserRegistered) GetUpdatedAt() *timestamppb.Timestamp {
	if x != nil {
		return x.UpdatedAt
	}
	return nil
}

type UsersRegistered struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id          int64                  `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	Username    string                 `protobuf:"bytes,2,opt,name=username,proto3" json:"username,omitempty"`
	FullName    string                 `protobuf:"bytes,3,opt,name=full_name,json=fullName,proto3" json:"full_name,omitempty"`
	Email       string                 `protobuf:"bytes,4,opt,name=email,proto3" json:"email,omitempty"`
	CountryCode string                 `protobuf:"bytes,5,opt,name=country_code,json=countryCode,proto3" json:"country_code,omitempty"`
	RoleId      int64                  `protobuf:"varint,6,opt,name=role_id,json=roleId,proto3" json:"role_id,omitempty"`
	LastLoginAt *timestamppb.Timestamp `protobuf:"bytes,7,opt,name=last_login_at,json=lastLoginAt,proto3" json:"last_login_at,omitempty"`
	CreatedAt   *timestamppb.Timestamp `protobuf:"bytes,8,opt,name=created_at,json=createdAt,proto3" json:"created_at,omitempty"`
	UpdatedAt   *timestamppb.Timestamp `protobuf:"bytes,9,opt,name=updated_at,json=updatedAt,proto3" json:"updated_at,omitempty"`
}

func (x *UsersRegistered) Reset() {
	*x = UsersRegistered{}
	if protoimpl.UnsafeEnabled {
		mi := &file_register_user_registered_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UsersRegistered) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UsersRegistered) ProtoMessage() {}

func (x *UsersRegistered) ProtoReflect() protoreflect.Message {
	mi := &file_register_user_registered_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UsersRegistered.ProtoReflect.Descriptor instead.
func (*UsersRegistered) Descriptor() ([]byte, []int) {
	return file_register_user_registered_proto_rawDescGZIP(), []int{1}
}

func (x *UsersRegistered) GetId() int64 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *UsersRegistered) GetUsername() string {
	if x != nil {
		return x.Username
	}
	return ""
}

func (x *UsersRegistered) GetFullName() string {
	if x != nil {
		return x.FullName
	}
	return ""
}

func (x *UsersRegistered) GetEmail() string {
	if x != nil {
		return x.Email
	}
	return ""
}

func (x *UsersRegistered) GetCountryCode() string {
	if x != nil {
		return x.CountryCode
	}
	return ""
}

func (x *UsersRegistered) GetRoleId() int64 {
	if x != nil {
		return x.RoleId
	}
	return 0
}

func (x *UsersRegistered) GetLastLoginAt() *timestamppb.Timestamp {
	if x != nil {
		return x.LastLoginAt
	}
	return nil
}

func (x *UsersRegistered) GetCreatedAt() *timestamppb.Timestamp {
	if x != nil {
		return x.CreatedAt
	}
	return nil
}

func (x *UsersRegistered) GetUpdatedAt() *timestamppb.Timestamp {
	if x != nil {
		return x.UpdatedAt
	}
	return nil
}

var File_register_user_registered_proto protoreflect.FileDescriptor

var file_register_user_registered_proto_rawDesc = []byte{
	0x0a, 0x1e, 0x72, 0x65, 0x67, 0x69, 0x73, 0x74, 0x65, 0x72, 0x2f, 0x75, 0x73, 0x65, 0x72, 0x5f,
	0x72, 0x65, 0x67, 0x69, 0x73, 0x74, 0x65, 0x72, 0x65, 0x64, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x12, 0x02, 0x70, 0x62, 0x1a, 0x1f, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x74, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xa2, 0x05, 0x0a, 0x0f, 0x55, 0x73, 0x65, 0x72, 0x5f, 0x72,
	0x65, 0x67, 0x69, 0x73, 0x74, 0x65, 0x72, 0x65, 0x64, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x02, 0x69, 0x64, 0x12, 0x1a, 0x0a, 0x08, 0x75, 0x73, 0x65,
	0x72, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x75, 0x73, 0x65,
	0x72, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x1b, 0x0a, 0x09, 0x66, 0x75, 0x6c, 0x6c, 0x5f, 0x6e, 0x61,
	0x6d, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x66, 0x75, 0x6c, 0x6c, 0x4e, 0x61,
	0x6d, 0x65, 0x12, 0x14, 0x0a, 0x05, 0x65, 0x6d, 0x61, 0x69, 0x6c, 0x18, 0x04, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x05, 0x65, 0x6d, 0x61, 0x69, 0x6c, 0x12, 0x23, 0x0a, 0x0d, 0x70, 0x61, 0x73, 0x73,
	0x77, 0x6f, 0x72, 0x64, 0x5f, 0x68, 0x61, 0x73, 0x68, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x0c, 0x70, 0x61, 0x73, 0x73, 0x77, 0x6f, 0x72, 0x64, 0x48, 0x61, 0x73, 0x68, 0x12, 0x23, 0x0a,
	0x0d, 0x70, 0x61, 0x73, 0x73, 0x77, 0x6f, 0x72, 0x64, 0x5f, 0x73, 0x61, 0x6c, 0x74, 0x18, 0x06,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x0c, 0x70, 0x61, 0x73, 0x73, 0x77, 0x6f, 0x72, 0x64, 0x53, 0x61,
	0x6c, 0x74, 0x12, 0x21, 0x0a, 0x0c, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x72, 0x79, 0x5f, 0x63, 0x6f,
	0x64, 0x65, 0x18, 0x07, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x72,
	0x79, 0x43, 0x6f, 0x64, 0x65, 0x12, 0x17, 0x0a, 0x07, 0x72, 0x6f, 0x6c, 0x65, 0x5f, 0x69, 0x64,
	0x18, 0x08, 0x20, 0x01, 0x28, 0x03, 0x52, 0x06, 0x72, 0x6f, 0x6c, 0x65, 0x49, 0x64, 0x12, 0x16,
	0x0a, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x18, 0x09, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06,
	0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x12, 0x3e, 0x0a, 0x0d, 0x6c, 0x61, 0x73, 0x74, 0x5f, 0x6c,
	0x6f, 0x67, 0x69, 0x6e, 0x5f, 0x61, 0x74, 0x18, 0x0a, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e,
	0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e,
	0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x52, 0x0b, 0x6c, 0x61, 0x73, 0x74, 0x4c,
	0x6f, 0x67, 0x69, 0x6e, 0x41, 0x74, 0x12, 0x4a, 0x0a, 0x13, 0x75, 0x73, 0x65, 0x72, 0x6e, 0x61,
	0x6d, 0x65, 0x5f, 0x63, 0x68, 0x61, 0x6e, 0x67, 0x65, 0x64, 0x5f, 0x61, 0x74, 0x18, 0x0b, 0x20,
	0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x52,
	0x11, 0x75, 0x73, 0x65, 0x72, 0x6e, 0x61, 0x6d, 0x65, 0x43, 0x68, 0x61, 0x6e, 0x67, 0x65, 0x64,
	0x41, 0x74, 0x12, 0x44, 0x0a, 0x10, 0x65, 0x6d, 0x61, 0x69, 0x6c, 0x5f, 0x63, 0x68, 0x61, 0x6e,
	0x67, 0x65, 0x64, 0x5f, 0x61, 0x74, 0x18, 0x0c, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67,
	0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54,
	0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x52, 0x0e, 0x65, 0x6d, 0x61, 0x69, 0x6c, 0x43,
	0x68, 0x61, 0x6e, 0x67, 0x65, 0x64, 0x41, 0x74, 0x12, 0x4a, 0x0a, 0x13, 0x70, 0x61, 0x73, 0x73,
	0x77, 0x6f, 0x72, 0x64, 0x5f, 0x63, 0x68, 0x61, 0x6e, 0x67, 0x65, 0x64, 0x5f, 0x61, 0x74, 0x18,
	0x0d, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d,
	0x70, 0x52, 0x11, 0x70, 0x61, 0x73, 0x73, 0x77, 0x6f, 0x72, 0x64, 0x43, 0x68, 0x61, 0x6e, 0x67,
	0x65, 0x64, 0x41, 0x74, 0x12, 0x39, 0x0a, 0x0a, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x5f,
	0x61, 0x74, 0x18, 0x0e, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c,
	0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73,
	0x74, 0x61, 0x6d, 0x70, 0x52, 0x09, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x41, 0x74, 0x12,
	0x39, 0x0a, 0x0a, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x64, 0x5f, 0x61, 0x74, 0x18, 0x0f, 0x20,
	0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x52,
	0x09, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x64, 0x41, 0x74, 0x22, 0xe3, 0x02, 0x0a, 0x10, 0x55,
	0x73, 0x65, 0x72, 0x73, 0x5f, 0x72, 0x65, 0x67, 0x69, 0x73, 0x74, 0x65, 0x72, 0x65, 0x64, 0x12,
	0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x02, 0x69, 0x64, 0x12,
	0x1a, 0x0a, 0x08, 0x75, 0x73, 0x65, 0x72, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x08, 0x75, 0x73, 0x65, 0x72, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x1b, 0x0a, 0x09, 0x66,
	0x75, 0x6c, 0x6c, 0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08,
	0x66, 0x75, 0x6c, 0x6c, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x14, 0x0a, 0x05, 0x65, 0x6d, 0x61, 0x69,
	0x6c, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x65, 0x6d, 0x61, 0x69, 0x6c, 0x12, 0x21,
	0x0a, 0x0c, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x72, 0x79, 0x5f, 0x63, 0x6f, 0x64, 0x65, 0x18, 0x05,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x72, 0x79, 0x43, 0x6f, 0x64,
	0x65, 0x12, 0x17, 0x0a, 0x07, 0x72, 0x6f, 0x6c, 0x65, 0x5f, 0x69, 0x64, 0x18, 0x06, 0x20, 0x01,
	0x28, 0x03, 0x52, 0x06, 0x72, 0x6f, 0x6c, 0x65, 0x49, 0x64, 0x12, 0x3e, 0x0a, 0x0d, 0x6c, 0x61,
	0x73, 0x74, 0x5f, 0x6c, 0x6f, 0x67, 0x69, 0x6e, 0x5f, 0x61, 0x74, 0x18, 0x07, 0x20, 0x01, 0x28,
	0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x62, 0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x52, 0x0b, 0x6c,
	0x61, 0x73, 0x74, 0x4c, 0x6f, 0x67, 0x69, 0x6e, 0x41, 0x74, 0x12, 0x39, 0x0a, 0x0a, 0x63, 0x72,
	0x65, 0x61, 0x74, 0x65, 0x64, 0x5f, 0x61, 0x74, 0x18, 0x08, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a,
	0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66,
	0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x52, 0x09, 0x63, 0x72, 0x65, 0x61,
	0x74, 0x65, 0x64, 0x41, 0x74, 0x12, 0x39, 0x0a, 0x0a, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x64,
	0x5f, 0x61, 0x74, 0x18, 0x09, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67,
	0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65,
	0x73, 0x74, 0x61, 0x6d, 0x70, 0x52, 0x09, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x64, 0x41, 0x74,
	0x42, 0x48, 0x5a, 0x46, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x53,
	0x74, 0x72, 0x65, 0x61, 0x6d, 0x66, 0x61, 0x69, 0x72, 0x2f, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e,
	0x5f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x49, 0x64, 0x65, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x50,
	0x72, 0x6f, 0x76, 0x69, 0x64, 0x65, 0x72, 0x2f, 0x70, 0x62, 0x2f, 0x75, 0x73, 0x65, 0x72, 0x5f,
	0x72, 0x65, 0x67, 0x69, 0x73, 0x74, 0x65, 0x72, 0x65, 0x64, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x33,
}

var (
	file_register_user_registered_proto_rawDescOnce sync.Once
	file_register_user_registered_proto_rawDescData = file_register_user_registered_proto_rawDesc
)

func file_register_user_registered_proto_rawDescGZIP() []byte {
	file_register_user_registered_proto_rawDescOnce.Do(func() {
		file_register_user_registered_proto_rawDescData = protoimpl.X.CompressGZIP(file_register_user_registered_proto_rawDescData)
	})
	return file_register_user_registered_proto_rawDescData
}

var file_register_user_registered_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_register_user_registered_proto_goTypes = []interface{}{
	(*UserRegistered)(nil),        // 0: pb.User_registered
	(*UsersRegistered)(nil),       // 1: pb.Users_registered
	(*timestamppb.Timestamp)(nil), // 2: google.protobuf.Timestamp
}
var file_register_user_registered_proto_depIdxs = []int32{
	2, // 0: pb.User_registered.last_login_at:type_name -> google.protobuf.Timestamp
	2, // 1: pb.User_registered.username_changed_at:type_name -> google.protobuf.Timestamp
	2, // 2: pb.User_registered.email_changed_at:type_name -> google.protobuf.Timestamp
	2, // 3: pb.User_registered.password_changed_at:type_name -> google.protobuf.Timestamp
	2, // 4: pb.User_registered.created_at:type_name -> google.protobuf.Timestamp
	2, // 5: pb.User_registered.updated_at:type_name -> google.protobuf.Timestamp
	2, // 6: pb.Users_registered.last_login_at:type_name -> google.protobuf.Timestamp
	2, // 7: pb.Users_registered.created_at:type_name -> google.protobuf.Timestamp
	2, // 8: pb.Users_registered.updated_at:type_name -> google.protobuf.Timestamp
	9, // [9:9] is the sub-list for method output_type
	9, // [9:9] is the sub-list for method input_type
	9, // [9:9] is the sub-list for extension type_name
	9, // [9:9] is the sub-list for extension extendee
	0, // [0:9] is the sub-list for field type_name
}

func init() { file_register_user_registered_proto_init() }
func file_register_user_registered_proto_init() {
	if File_register_user_registered_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_register_user_registered_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UserRegistered); i {
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
		file_register_user_registered_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UsersRegistered); i {
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
			RawDescriptor: file_register_user_registered_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_register_user_registered_proto_goTypes,
		DependencyIndexes: file_register_user_registered_proto_depIdxs,
		MessageInfos:      file_register_user_registered_proto_msgTypes,
	}.Build()
	File_register_user_registered_proto = out.File
	file_register_user_registered_proto_rawDesc = nil
	file_register_user_registered_proto_goTypes = nil
	file_register_user_registered_proto_depIdxs = nil
}
