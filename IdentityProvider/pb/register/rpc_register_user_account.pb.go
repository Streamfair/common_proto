// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.33.0
// 	protoc        v5.26.1
// source: register/rpc_register_user_account.proto

package register

import (
	account "github.com/Streamfair/common_proto/AccountService/pb/account"
	account_type "github.com/Streamfair/common_proto/AccountService/pb/account_type"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	_ "google.golang.org/protobuf/types/known/timestamppb"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type CreateUserRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Username     string `protobuf:"bytes,1,opt,name=username,proto3" json:"username,omitempty"`
	FullName     string `protobuf:"bytes,2,opt,name=full_name,json=fullName,proto3" json:"full_name,omitempty"`
	Email        string `protobuf:"bytes,3,opt,name=email,proto3" json:"email,omitempty"`
	PasswordHash string `protobuf:"bytes,4,opt,name=password_hash,json=passwordHash,proto3" json:"password_hash,omitempty"`
	PasswordSalt string `protobuf:"bytes,5,opt,name=password_salt,json=passwordSalt,proto3" json:"password_salt,omitempty"`
	CountryCode  string `protobuf:"bytes,6,opt,name=country_code,json=countryCode,proto3" json:"country_code,omitempty"`
	RoleId       int64  `protobuf:"varint,7,opt,name=role_id,json=roleId,proto3" json:"role_id,omitempty"`
	Status       string `protobuf:"bytes,8,opt,name=status,proto3" json:"status,omitempty"`
}

func (x *CreateUserRequest) Reset() {
	*x = CreateUserRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_register_rpc_register_user_account_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreateUserRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateUserRequest) ProtoMessage() {}

func (x *CreateUserRequest) ProtoReflect() protoreflect.Message {
	mi := &file_register_rpc_register_user_account_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateUserRequest.ProtoReflect.Descriptor instead.
func (*CreateUserRequest) Descriptor() ([]byte, []int) {
	return file_register_rpc_register_user_account_proto_rawDescGZIP(), []int{0}
}

func (x *CreateUserRequest) GetUsername() string {
	if x != nil {
		return x.Username
	}
	return ""
}

func (x *CreateUserRequest) GetFullName() string {
	if x != nil {
		return x.FullName
	}
	return ""
}

func (x *CreateUserRequest) GetEmail() string {
	if x != nil {
		return x.Email
	}
	return ""
}

func (x *CreateUserRequest) GetPasswordHash() string {
	if x != nil {
		return x.PasswordHash
	}
	return ""
}

func (x *CreateUserRequest) GetPasswordSalt() string {
	if x != nil {
		return x.PasswordSalt
	}
	return ""
}

func (x *CreateUserRequest) GetCountryCode() string {
	if x != nil {
		return x.CountryCode
	}
	return ""
}

func (x *CreateUserRequest) GetRoleId() int64 {
	if x != nil {
		return x.RoleId
	}
	return 0
}

func (x *CreateUserRequest) GetStatus() string {
	if x != nil {
		return x.Status
	}
	return ""
}

type CreateAccountTypeRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Type        int32  `protobuf:"varint,1,opt,name=type,proto3" json:"type,omitempty"`
	Permissions string `protobuf:"bytes,2,opt,name=permissions,proto3" json:"permissions,omitempty"`
	IsArtist    bool   `protobuf:"varint,3,opt,name=is_artist,json=isArtist,proto3" json:"is_artist,omitempty"`
	IsProducer  bool   `protobuf:"varint,4,opt,name=is_producer,json=isProducer,proto3" json:"is_producer,omitempty"`
	IsWriter    bool   `protobuf:"varint,5,opt,name=is_writer,json=isWriter,proto3" json:"is_writer,omitempty"`
	IsLabel     bool   `protobuf:"varint,6,opt,name=is_label,json=isLabel,proto3" json:"is_label,omitempty"`
	IsUser      bool   `protobuf:"varint,7,opt,name=is_user,json=isUser,proto3" json:"is_user,omitempty"`
}

func (x *CreateAccountTypeRequest) Reset() {
	*x = CreateAccountTypeRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_register_rpc_register_user_account_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreateAccountTypeRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateAccountTypeRequest) ProtoMessage() {}

func (x *CreateAccountTypeRequest) ProtoReflect() protoreflect.Message {
	mi := &file_register_rpc_register_user_account_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateAccountTypeRequest.ProtoReflect.Descriptor instead.
func (*CreateAccountTypeRequest) Descriptor() ([]byte, []int) {
	return file_register_rpc_register_user_account_proto_rawDescGZIP(), []int{1}
}

func (x *CreateAccountTypeRequest) GetType() int32 {
	if x != nil {
		return x.Type
	}
	return 0
}

func (x *CreateAccountTypeRequest) GetPermissions() string {
	if x != nil {
		return x.Permissions
	}
	return ""
}

func (x *CreateAccountTypeRequest) GetIsArtist() bool {
	if x != nil {
		return x.IsArtist
	}
	return false
}

func (x *CreateAccountTypeRequest) GetIsProducer() bool {
	if x != nil {
		return x.IsProducer
	}
	return false
}

func (x *CreateAccountTypeRequest) GetIsWriter() bool {
	if x != nil {
		return x.IsWriter
	}
	return false
}

func (x *CreateAccountTypeRequest) GetIsLabel() bool {
	if x != nil {
		return x.IsLabel
	}
	return false
}

func (x *CreateAccountTypeRequest) GetIsUser() bool {
	if x != nil {
		return x.IsUser
	}
	return false
}

type CreateAccountRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	AccountType int32  `protobuf:"varint,1,opt,name=account_type,json=accountType,proto3" json:"account_type,omitempty"`
	AccountName string `protobuf:"bytes,2,opt,name=account_name,json=accountName,proto3" json:"account_name,omitempty"`
	Owner       string `protobuf:"bytes,3,opt,name=owner,proto3" json:"owner,omitempty"`
	Bio         string `protobuf:"bytes,4,opt,name=bio,proto3" json:"bio,omitempty"`
	Status      string `protobuf:"bytes,5,opt,name=status,proto3" json:"status,omitempty"`
	Plan        string `protobuf:"bytes,6,opt,name=plan,proto3" json:"plan,omitempty"`
	AvatarUri   string `protobuf:"bytes,7,opt,name=avatar_uri,json=avatarUri,proto3" json:"avatar_uri,omitempty"`
	Plays       int64  `protobuf:"varint,8,opt,name=plays,proto3" json:"plays,omitempty"`
	Likes       int64  `protobuf:"varint,9,opt,name=likes,proto3" json:"likes,omitempty"`
	Follows     int64  `protobuf:"varint,10,opt,name=follows,proto3" json:"follows,omitempty"`
	Shares      int64  `protobuf:"varint,11,opt,name=shares,proto3" json:"shares,omitempty"`
}

func (x *CreateAccountRequest) Reset() {
	*x = CreateAccountRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_register_rpc_register_user_account_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreateAccountRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateAccountRequest) ProtoMessage() {}

func (x *CreateAccountRequest) ProtoReflect() protoreflect.Message {
	mi := &file_register_rpc_register_user_account_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateAccountRequest.ProtoReflect.Descriptor instead.
func (*CreateAccountRequest) Descriptor() ([]byte, []int) {
	return file_register_rpc_register_user_account_proto_rawDescGZIP(), []int{2}
}

func (x *CreateAccountRequest) GetAccountType() int32 {
	if x != nil {
		return x.AccountType
	}
	return 0
}

func (x *CreateAccountRequest) GetAccountName() string {
	if x != nil {
		return x.AccountName
	}
	return ""
}

func (x *CreateAccountRequest) GetOwner() string {
	if x != nil {
		return x.Owner
	}
	return ""
}

func (x *CreateAccountRequest) GetBio() string {
	if x != nil {
		return x.Bio
	}
	return ""
}

func (x *CreateAccountRequest) GetStatus() string {
	if x != nil {
		return x.Status
	}
	return ""
}

func (x *CreateAccountRequest) GetPlan() string {
	if x != nil {
		return x.Plan
	}
	return ""
}

func (x *CreateAccountRequest) GetAvatarUri() string {
	if x != nil {
		return x.AvatarUri
	}
	return ""
}

func (x *CreateAccountRequest) GetPlays() int64 {
	if x != nil {
		return x.Plays
	}
	return 0
}

func (x *CreateAccountRequest) GetLikes() int64 {
	if x != nil {
		return x.Likes
	}
	return 0
}

func (x *CreateAccountRequest) GetFollows() int64 {
	if x != nil {
		return x.Follows
	}
	return 0
}

func (x *CreateAccountRequest) GetShares() int64 {
	if x != nil {
		return x.Shares
	}
	return 0
}

type RegisterUserAccountRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	User        *CreateUserRequest        `protobuf:"bytes,1,opt,name=user,proto3" json:"user,omitempty"`
	Account     *CreateAccountRequest     `protobuf:"bytes,2,opt,name=account,proto3" json:"account,omitempty"`
	AccountType *CreateAccountTypeRequest `protobuf:"bytes,3,opt,name=account_type,json=accountType,proto3" json:"account_type,omitempty"`
}

func (x *RegisterUserAccountRequest) Reset() {
	*x = RegisterUserAccountRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_register_rpc_register_user_account_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RegisterUserAccountRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RegisterUserAccountRequest) ProtoMessage() {}

func (x *RegisterUserAccountRequest) ProtoReflect() protoreflect.Message {
	mi := &file_register_rpc_register_user_account_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RegisterUserAccountRequest.ProtoReflect.Descriptor instead.
func (*RegisterUserAccountRequest) Descriptor() ([]byte, []int) {
	return file_register_rpc_register_user_account_proto_rawDescGZIP(), []int{3}
}

func (x *RegisterUserAccountRequest) GetUser() *CreateUserRequest {
	if x != nil {
		return x.User
	}
	return nil
}

func (x *RegisterUserAccountRequest) GetAccount() *CreateAccountRequest {
	if x != nil {
		return x.Account
	}
	return nil
}

func (x *RegisterUserAccountRequest) GetAccountType() *CreateAccountTypeRequest {
	if x != nil {
		return x.AccountType
	}
	return nil
}

type UserAccount struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	User        *User                     `protobuf:"bytes,1,opt,name=user,proto3" json:"user,omitempty"`
	Account     *account.Account          `protobuf:"bytes,2,opt,name=account,proto3" json:"account,omitempty"`
	AccountType *account_type.AccountType `protobuf:"bytes,3,opt,name=account_type,json=accountType,proto3" json:"account_type,omitempty"`
}

func (x *UserAccount) Reset() {
	*x = UserAccount{}
	if protoimpl.UnsafeEnabled {
		mi := &file_register_rpc_register_user_account_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UserAccount) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UserAccount) ProtoMessage() {}

func (x *UserAccount) ProtoReflect() protoreflect.Message {
	mi := &file_register_rpc_register_user_account_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UserAccount.ProtoReflect.Descriptor instead.
func (*UserAccount) Descriptor() ([]byte, []int) {
	return file_register_rpc_register_user_account_proto_rawDescGZIP(), []int{4}
}

func (x *UserAccount) GetUser() *User {
	if x != nil {
		return x.User
	}
	return nil
}

func (x *UserAccount) GetAccount() *account.Account {
	if x != nil {
		return x.Account
	}
	return nil
}

func (x *UserAccount) GetAccountType() *account_type.AccountType {
	if x != nil {
		return x.AccountType
	}
	return nil
}

var File_register_rpc_register_user_account_proto protoreflect.FileDescriptor

var file_register_rpc_register_user_account_proto_rawDesc = []byte{
	0x0a, 0x28, 0x72, 0x65, 0x67, 0x69, 0x73, 0x74, 0x65, 0x72, 0x2f, 0x72, 0x70, 0x63, 0x5f, 0x72,
	0x65, 0x67, 0x69, 0x73, 0x74, 0x65, 0x72, 0x5f, 0x75, 0x73, 0x65, 0x72, 0x5f, 0x61, 0x63, 0x63,
	0x6f, 0x75, 0x6e, 0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x02, 0x70, 0x62, 0x1a, 0x1f,
	0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f,
	0x74, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a,
	0x1b, 0x72, 0x65, 0x67, 0x69, 0x73, 0x74, 0x65, 0x72, 0x2f, 0x75, 0x73, 0x65, 0x72, 0x5f, 0x61,
	0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x2a, 0x41, 0x63,
	0x63, 0x6f, 0x75, 0x6e, 0x74, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2f, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x2f, 0x61, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x2f, 0x61, 0x63, 0x63, 0x6f, 0x75,
	0x6e, 0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x34, 0x41, 0x63, 0x63, 0x6f, 0x75, 0x6e,
	0x74, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x61,
	0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x5f, 0x74, 0x79, 0x70, 0x65, 0x2f, 0x61, 0x63, 0x63, 0x6f,
	0x75, 0x6e, 0x74, 0x5f, 0x74, 0x79, 0x70, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x80,
	0x02, 0x0a, 0x11, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x55, 0x73, 0x65, 0x72, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x12, 0x1a, 0x0a, 0x08, 0x75, 0x73, 0x65, 0x72, 0x6e, 0x61, 0x6d, 0x65,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x75, 0x73, 0x65, 0x72, 0x6e, 0x61, 0x6d, 0x65,
	0x12, 0x1b, 0x0a, 0x09, 0x66, 0x75, 0x6c, 0x6c, 0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x08, 0x66, 0x75, 0x6c, 0x6c, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x14, 0x0a,
	0x05, 0x65, 0x6d, 0x61, 0x69, 0x6c, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x65, 0x6d,
	0x61, 0x69, 0x6c, 0x12, 0x23, 0x0a, 0x0d, 0x70, 0x61, 0x73, 0x73, 0x77, 0x6f, 0x72, 0x64, 0x5f,
	0x68, 0x61, 0x73, 0x68, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0c, 0x70, 0x61, 0x73, 0x73,
	0x77, 0x6f, 0x72, 0x64, 0x48, 0x61, 0x73, 0x68, 0x12, 0x23, 0x0a, 0x0d, 0x70, 0x61, 0x73, 0x73,
	0x77, 0x6f, 0x72, 0x64, 0x5f, 0x73, 0x61, 0x6c, 0x74, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x0c, 0x70, 0x61, 0x73, 0x73, 0x77, 0x6f, 0x72, 0x64, 0x53, 0x61, 0x6c, 0x74, 0x12, 0x21, 0x0a,
	0x0c, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x72, 0x79, 0x5f, 0x63, 0x6f, 0x64, 0x65, 0x18, 0x06, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x0b, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x72, 0x79, 0x43, 0x6f, 0x64, 0x65,
	0x12, 0x17, 0x0a, 0x07, 0x72, 0x6f, 0x6c, 0x65, 0x5f, 0x69, 0x64, 0x18, 0x07, 0x20, 0x01, 0x28,
	0x03, 0x52, 0x06, 0x72, 0x6f, 0x6c, 0x65, 0x49, 0x64, 0x12, 0x16, 0x0a, 0x06, 0x73, 0x74, 0x61,
	0x74, 0x75, 0x73, 0x18, 0x08, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75,
	0x73, 0x22, 0xdf, 0x01, 0x0a, 0x18, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x41, 0x63, 0x63, 0x6f,
	0x75, 0x6e, 0x74, 0x54, 0x79, 0x70, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x12,
	0x0a, 0x04, 0x74, 0x79, 0x70, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x04, 0x74, 0x79,
	0x70, 0x65, 0x12, 0x20, 0x0a, 0x0b, 0x70, 0x65, 0x72, 0x6d, 0x69, 0x73, 0x73, 0x69, 0x6f, 0x6e,
	0x73, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x70, 0x65, 0x72, 0x6d, 0x69, 0x73, 0x73,
	0x69, 0x6f, 0x6e, 0x73, 0x12, 0x1b, 0x0a, 0x09, 0x69, 0x73, 0x5f, 0x61, 0x72, 0x74, 0x69, 0x73,
	0x74, 0x18, 0x03, 0x20, 0x01, 0x28, 0x08, 0x52, 0x08, 0x69, 0x73, 0x41, 0x72, 0x74, 0x69, 0x73,
	0x74, 0x12, 0x1f, 0x0a, 0x0b, 0x69, 0x73, 0x5f, 0x70, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x65, 0x72,
	0x18, 0x04, 0x20, 0x01, 0x28, 0x08, 0x52, 0x0a, 0x69, 0x73, 0x50, 0x72, 0x6f, 0x64, 0x75, 0x63,
	0x65, 0x72, 0x12, 0x1b, 0x0a, 0x09, 0x69, 0x73, 0x5f, 0x77, 0x72, 0x69, 0x74, 0x65, 0x72, 0x18,
	0x05, 0x20, 0x01, 0x28, 0x08, 0x52, 0x08, 0x69, 0x73, 0x57, 0x72, 0x69, 0x74, 0x65, 0x72, 0x12,
	0x19, 0x0a, 0x08, 0x69, 0x73, 0x5f, 0x6c, 0x61, 0x62, 0x65, 0x6c, 0x18, 0x06, 0x20, 0x01, 0x28,
	0x08, 0x52, 0x07, 0x69, 0x73, 0x4c, 0x61, 0x62, 0x65, 0x6c, 0x12, 0x17, 0x0a, 0x07, 0x69, 0x73,
	0x5f, 0x75, 0x73, 0x65, 0x72, 0x18, 0x07, 0x20, 0x01, 0x28, 0x08, 0x52, 0x06, 0x69, 0x73, 0x55,
	0x73, 0x65, 0x72, 0x22, 0xad, 0x02, 0x0a, 0x14, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x41, 0x63,
	0x63, 0x6f, 0x75, 0x6e, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x21, 0x0a, 0x0c,
	0x61, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x5f, 0x74, 0x79, 0x70, 0x65, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x05, 0x52, 0x0b, 0x61, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x54, 0x79, 0x70, 0x65, 0x12,
	0x21, 0x0a, 0x0c, 0x61, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x61, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x4e, 0x61,
	0x6d, 0x65, 0x12, 0x14, 0x0a, 0x05, 0x6f, 0x77, 0x6e, 0x65, 0x72, 0x18, 0x03, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x05, 0x6f, 0x77, 0x6e, 0x65, 0x72, 0x12, 0x10, 0x0a, 0x03, 0x62, 0x69, 0x6f, 0x18,
	0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x62, 0x69, 0x6f, 0x12, 0x16, 0x0a, 0x06, 0x73, 0x74,
	0x61, 0x74, 0x75, 0x73, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x73, 0x74, 0x61, 0x74,
	0x75, 0x73, 0x12, 0x12, 0x0a, 0x04, 0x70, 0x6c, 0x61, 0x6e, 0x18, 0x06, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x04, 0x70, 0x6c, 0x61, 0x6e, 0x12, 0x1d, 0x0a, 0x0a, 0x61, 0x76, 0x61, 0x74, 0x61, 0x72,
	0x5f, 0x75, 0x72, 0x69, 0x18, 0x07, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x61, 0x76, 0x61, 0x74,
	0x61, 0x72, 0x55, 0x72, 0x69, 0x12, 0x14, 0x0a, 0x05, 0x70, 0x6c, 0x61, 0x79, 0x73, 0x18, 0x08,
	0x20, 0x01, 0x28, 0x03, 0x52, 0x05, 0x70, 0x6c, 0x61, 0x79, 0x73, 0x12, 0x14, 0x0a, 0x05, 0x6c,
	0x69, 0x6b, 0x65, 0x73, 0x18, 0x09, 0x20, 0x01, 0x28, 0x03, 0x52, 0x05, 0x6c, 0x69, 0x6b, 0x65,
	0x73, 0x12, 0x18, 0x0a, 0x07, 0x66, 0x6f, 0x6c, 0x6c, 0x6f, 0x77, 0x73, 0x18, 0x0a, 0x20, 0x01,
	0x28, 0x03, 0x52, 0x07, 0x66, 0x6f, 0x6c, 0x6c, 0x6f, 0x77, 0x73, 0x12, 0x16, 0x0a, 0x06, 0x73,
	0x68, 0x61, 0x72, 0x65, 0x73, 0x18, 0x0b, 0x20, 0x01, 0x28, 0x03, 0x52, 0x06, 0x73, 0x68, 0x61,
	0x72, 0x65, 0x73, 0x22, 0xbc, 0x01, 0x0a, 0x1a, 0x52, 0x65, 0x67, 0x69, 0x73, 0x74, 0x65, 0x72,
	0x55, 0x73, 0x65, 0x72, 0x41, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x12, 0x29, 0x0a, 0x04, 0x75, 0x73, 0x65, 0x72, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b,
	0x32, 0x15, 0x2e, 0x70, 0x62, 0x2e, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x55, 0x73, 0x65, 0x72,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x52, 0x04, 0x75, 0x73, 0x65, 0x72, 0x12, 0x32, 0x0a,
	0x07, 0x61, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x18,
	0x2e, 0x70, 0x62, 0x2e, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x41, 0x63, 0x63, 0x6f, 0x75, 0x6e,
	0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x52, 0x07, 0x61, 0x63, 0x63, 0x6f, 0x75, 0x6e,
	0x74, 0x12, 0x3f, 0x0a, 0x0c, 0x61, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x5f, 0x74, 0x79, 0x70,
	0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1c, 0x2e, 0x70, 0x62, 0x2e, 0x43, 0x72, 0x65,
	0x61, 0x74, 0x65, 0x41, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x54, 0x79, 0x70, 0x65, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x52, 0x0b, 0x61, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x54, 0x79,
	0x70, 0x65, 0x22, 0x86, 0x01, 0x0a, 0x0b, 0x55, 0x73, 0x65, 0x72, 0x41, 0x63, 0x63, 0x6f, 0x75,
	0x6e, 0x74, 0x12, 0x1c, 0x0a, 0x04, 0x75, 0x73, 0x65, 0x72, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b,
	0x32, 0x08, 0x2e, 0x70, 0x62, 0x2e, 0x55, 0x73, 0x65, 0x72, 0x52, 0x04, 0x75, 0x73, 0x65, 0x72,
	0x12, 0x25, 0x0a, 0x07, 0x61, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x0b, 0x32, 0x0b, 0x2e, 0x70, 0x62, 0x2e, 0x41, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x52, 0x07,
	0x61, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x12, 0x32, 0x0a, 0x0c, 0x61, 0x63, 0x63, 0x6f, 0x75,
	0x6e, 0x74, 0x5f, 0x74, 0x79, 0x70, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0f, 0x2e,
	0x70, 0x62, 0x2e, 0x41, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x54, 0x79, 0x70, 0x65, 0x52, 0x0b,
	0x61, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x54, 0x79, 0x70, 0x65, 0x42, 0x41, 0x5a, 0x3f, 0x67,
	0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x53, 0x74, 0x72, 0x65, 0x61, 0x6d,
	0x66, 0x61, 0x69, 0x72, 0x2f, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x5f, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x2f, 0x49, 0x64, 0x65, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x50, 0x72, 0x6f, 0x76, 0x69, 0x64,
	0x65, 0x72, 0x2f, 0x70, 0x62, 0x2f, 0x72, 0x65, 0x67, 0x69, 0x73, 0x74, 0x65, 0x72, 0x62, 0x06,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_register_rpc_register_user_account_proto_rawDescOnce sync.Once
	file_register_rpc_register_user_account_proto_rawDescData = file_register_rpc_register_user_account_proto_rawDesc
)

func file_register_rpc_register_user_account_proto_rawDescGZIP() []byte {
	file_register_rpc_register_user_account_proto_rawDescOnce.Do(func() {
		file_register_rpc_register_user_account_proto_rawDescData = protoimpl.X.CompressGZIP(file_register_rpc_register_user_account_proto_rawDescData)
	})
	return file_register_rpc_register_user_account_proto_rawDescData
}

var file_register_rpc_register_user_account_proto_msgTypes = make([]protoimpl.MessageInfo, 5)
var file_register_rpc_register_user_account_proto_goTypes = []interface{}{
	(*CreateUserRequest)(nil),          // 0: pb.CreateUserRequest
	(*CreateAccountTypeRequest)(nil),   // 1: pb.CreateAccountTypeRequest
	(*CreateAccountRequest)(nil),       // 2: pb.CreateAccountRequest
	(*RegisterUserAccountRequest)(nil), // 3: pb.RegisterUserAccountRequest
	(*UserAccount)(nil),                // 4: pb.UserAccount
	(*User)(nil),                       // 5: pb.User
	(*account.Account)(nil),            // 6: pb.Account
	(*account_type.AccountType)(nil),   // 7: pb.AccountType
}
var file_register_rpc_register_user_account_proto_depIdxs = []int32{
	0, // 0: pb.RegisterUserAccountRequest.user:type_name -> pb.CreateUserRequest
	2, // 1: pb.RegisterUserAccountRequest.account:type_name -> pb.CreateAccountRequest
	1, // 2: pb.RegisterUserAccountRequest.account_type:type_name -> pb.CreateAccountTypeRequest
	5, // 3: pb.UserAccount.user:type_name -> pb.User
	6, // 4: pb.UserAccount.account:type_name -> pb.Account
	7, // 5: pb.UserAccount.account_type:type_name -> pb.AccountType
	6, // [6:6] is the sub-list for method output_type
	6, // [6:6] is the sub-list for method input_type
	6, // [6:6] is the sub-list for extension type_name
	6, // [6:6] is the sub-list for extension extendee
	0, // [0:6] is the sub-list for field type_name
}

func init() { file_register_rpc_register_user_account_proto_init() }
func file_register_rpc_register_user_account_proto_init() {
	if File_register_rpc_register_user_account_proto != nil {
		return
	}
	file_register_user_account_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_register_rpc_register_user_account_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CreateUserRequest); i {
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
		file_register_rpc_register_user_account_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CreateAccountTypeRequest); i {
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
		file_register_rpc_register_user_account_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CreateAccountRequest); i {
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
		file_register_rpc_register_user_account_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*RegisterUserAccountRequest); i {
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
		file_register_rpc_register_user_account_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UserAccount); i {
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
			RawDescriptor: file_register_rpc_register_user_account_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   5,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_register_rpc_register_user_account_proto_goTypes,
		DependencyIndexes: file_register_rpc_register_user_account_proto_depIdxs,
		MessageInfos:      file_register_rpc_register_user_account_proto_msgTypes,
	}.Build()
	File_register_rpc_register_user_account_proto = out.File
	file_register_rpc_register_user_account_proto_rawDesc = nil
	file_register_rpc_register_user_account_proto_goTypes = nil
	file_register_rpc_register_user_account_proto_depIdxs = nil
}