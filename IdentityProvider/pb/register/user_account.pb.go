// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.33.0
// 	protoc        v5.26.1
// source: register/user_account.proto

package register

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

type AccountData struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id          int64                  `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	AccType     int32                  `protobuf:"varint,2,opt,name=acc_type,json=accType,proto3" json:"acc_type,omitempty"`
	AccountName string                 `protobuf:"bytes,3,opt,name=account_name,json=accountName,proto3" json:"account_name,omitempty"`
	Owner       string                 `protobuf:"bytes,4,opt,name=owner,proto3" json:"owner,omitempty"`
	Bio         string                 `protobuf:"bytes,5,opt,name=bio,proto3" json:"bio,omitempty"`
	Status      string                 `protobuf:"bytes,6,opt,name=status,proto3" json:"status,omitempty"`
	Plan        int32                  `protobuf:"varint,7,opt,name=plan,proto3" json:"plan,omitempty"`
	AvatarUri   string                 `protobuf:"bytes,8,opt,name=avatar_uri,json=avatarUri,proto3" json:"avatar_uri,omitempty"`
	Plays       int64                  `protobuf:"varint,9,opt,name=plays,proto3" json:"plays,omitempty"`
	Likes       int64                  `protobuf:"varint,10,opt,name=likes,proto3" json:"likes,omitempty"`
	Follows     int64                  `protobuf:"varint,11,opt,name=follows,proto3" json:"follows,omitempty"`
	Shares      int64                  `protobuf:"varint,12,opt,name=shares,proto3" json:"shares,omitempty"`
	CreatedAt   *timestamppb.Timestamp `protobuf:"bytes,13,opt,name=created_at,json=createdAt,proto3" json:"created_at,omitempty"`
	UpdatedAt   *timestamppb.Timestamp `protobuf:"bytes,14,opt,name=updated_at,json=updatedAt,proto3" json:"updated_at,omitempty"`
	AccountType *AccountTypeData       `protobuf:"bytes,15,opt,name=account_type,json=accountType,proto3" json:"account_type,omitempty"`
}

func (x *AccountData) Reset() {
	*x = AccountData{}
	if protoimpl.UnsafeEnabled {
		mi := &file_register_user_account_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AccountData) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AccountData) ProtoMessage() {}

func (x *AccountData) ProtoReflect() protoreflect.Message {
	mi := &file_register_user_account_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AccountData.ProtoReflect.Descriptor instead.
func (*AccountData) Descriptor() ([]byte, []int) {
	return file_register_user_account_proto_rawDescGZIP(), []int{0}
}

func (x *AccountData) GetId() int64 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *AccountData) GetAccType() int32 {
	if x != nil {
		return x.AccType
	}
	return 0
}

func (x *AccountData) GetAccountName() string {
	if x != nil {
		return x.AccountName
	}
	return ""
}

func (x *AccountData) GetOwner() string {
	if x != nil {
		return x.Owner
	}
	return ""
}

func (x *AccountData) GetBio() string {
	if x != nil {
		return x.Bio
	}
	return ""
}

func (x *AccountData) GetStatus() string {
	if x != nil {
		return x.Status
	}
	return ""
}

func (x *AccountData) GetPlan() int32 {
	if x != nil {
		return x.Plan
	}
	return 0
}

func (x *AccountData) GetAvatarUri() string {
	if x != nil {
		return x.AvatarUri
	}
	return ""
}

func (x *AccountData) GetPlays() int64 {
	if x != nil {
		return x.Plays
	}
	return 0
}

func (x *AccountData) GetLikes() int64 {
	if x != nil {
		return x.Likes
	}
	return 0
}

func (x *AccountData) GetFollows() int64 {
	if x != nil {
		return x.Follows
	}
	return 0
}

func (x *AccountData) GetShares() int64 {
	if x != nil {
		return x.Shares
	}
	return 0
}

func (x *AccountData) GetCreatedAt() *timestamppb.Timestamp {
	if x != nil {
		return x.CreatedAt
	}
	return nil
}

func (x *AccountData) GetUpdatedAt() *timestamppb.Timestamp {
	if x != nil {
		return x.UpdatedAt
	}
	return nil
}

func (x *AccountData) GetAccountType() *AccountTypeData {
	if x != nil {
		return x.AccountType
	}
	return nil
}

type AccountTypeData struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id          int64                  `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	AccType     int32                  `protobuf:"varint,2,opt,name=acc_type,json=accType,proto3" json:"acc_type,omitempty"`
	Permissions string                 `protobuf:"bytes,3,opt,name=permissions,proto3" json:"permissions,omitempty"`
	IsArtist    bool                   `protobuf:"varint,4,opt,name=is_artist,json=isArtist,proto3" json:"is_artist,omitempty"`
	IsProducer  bool                   `protobuf:"varint,5,opt,name=is_producer,json=isProducer,proto3" json:"is_producer,omitempty"`
	IsWriter    bool                   `protobuf:"varint,6,opt,name=is_writer,json=isWriter,proto3" json:"is_writer,omitempty"`
	IsLabel     bool                   `protobuf:"varint,7,opt,name=is_label,json=isLabel,proto3" json:"is_label,omitempty"`
	IsUser      bool                   `protobuf:"varint,8,opt,name=is_user,json=isUser,proto3" json:"is_user,omitempty"`
	CreatedAt   *timestamppb.Timestamp `protobuf:"bytes,9,opt,name=created_at,json=createdAt,proto3" json:"created_at,omitempty"`
	UpdatedAt   *timestamppb.Timestamp `protobuf:"bytes,10,opt,name=updated_at,json=updatedAt,proto3" json:"updated_at,omitempty"`
}

func (x *AccountTypeData) Reset() {
	*x = AccountTypeData{}
	if protoimpl.UnsafeEnabled {
		mi := &file_register_user_account_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AccountTypeData) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AccountTypeData) ProtoMessage() {}

func (x *AccountTypeData) ProtoReflect() protoreflect.Message {
	mi := &file_register_user_account_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AccountTypeData.ProtoReflect.Descriptor instead.
func (*AccountTypeData) Descriptor() ([]byte, []int) {
	return file_register_user_account_proto_rawDescGZIP(), []int{1}
}

func (x *AccountTypeData) GetId() int64 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *AccountTypeData) GetAccType() int32 {
	if x != nil {
		return x.AccType
	}
	return 0
}

func (x *AccountTypeData) GetPermissions() string {
	if x != nil {
		return x.Permissions
	}
	return ""
}

func (x *AccountTypeData) GetIsArtist() bool {
	if x != nil {
		return x.IsArtist
	}
	return false
}

func (x *AccountTypeData) GetIsProducer() bool {
	if x != nil {
		return x.IsProducer
	}
	return false
}

func (x *AccountTypeData) GetIsWriter() bool {
	if x != nil {
		return x.IsWriter
	}
	return false
}

func (x *AccountTypeData) GetIsLabel() bool {
	if x != nil {
		return x.IsLabel
	}
	return false
}

func (x *AccountTypeData) GetIsUser() bool {
	if x != nil {
		return x.IsUser
	}
	return false
}

func (x *AccountTypeData) GetCreatedAt() *timestamppb.Timestamp {
	if x != nil {
		return x.CreatedAt
	}
	return nil
}

func (x *AccountTypeData) GetUpdatedAt() *timestamppb.Timestamp {
	if x != nil {
		return x.UpdatedAt
	}
	return nil
}

type UserData struct {
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
	AccountType       int32                  `protobuf:"varint,10,opt,name=account_type,json=accountType,proto3" json:"account_type,omitempty"`
	AccountName       string                 `protobuf:"bytes,11,opt,name=account_name,json=accountName,proto3" json:"account_name,omitempty"`
	AccountBio        string                 `protobuf:"bytes,12,opt,name=account_bio,json=accountBio,proto3" json:"account_bio,omitempty"`
	AccountPlan       int32                  `protobuf:"varint,13,opt,name=account_plan,json=accountPlan,proto3" json:"account_plan,omitempty"`
	AccountAvatarUri  string                 `protobuf:"bytes,14,opt,name=account_avatar_uri,json=accountAvatarUri,proto3" json:"account_avatar_uri,omitempty"`
	LastLoginAt       *timestamppb.Timestamp `protobuf:"bytes,15,opt,name=last_login_at,json=lastLoginAt,proto3" json:"last_login_at,omitempty"`
	UsernameChangedAt *timestamppb.Timestamp `protobuf:"bytes,16,opt,name=username_changed_at,json=usernameChangedAt,proto3" json:"username_changed_at,omitempty"`
	EmailChangedAt    *timestamppb.Timestamp `protobuf:"bytes,17,opt,name=email_changed_at,json=emailChangedAt,proto3" json:"email_changed_at,omitempty"`
	PasswordChangedAt *timestamppb.Timestamp `protobuf:"bytes,18,opt,name=password_changed_at,json=passwordChangedAt,proto3" json:"password_changed_at,omitempty"`
	CreatedAt         *timestamppb.Timestamp `protobuf:"bytes,19,opt,name=created_at,json=createdAt,proto3" json:"created_at,omitempty"`
	UpdatedAt         *timestamppb.Timestamp `protobuf:"bytes,20,opt,name=updated_at,json=updatedAt,proto3" json:"updated_at,omitempty"`
}

func (x *UserData) Reset() {
	*x = UserData{}
	if protoimpl.UnsafeEnabled {
		mi := &file_register_user_account_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UserData) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UserData) ProtoMessage() {}

func (x *UserData) ProtoReflect() protoreflect.Message {
	mi := &file_register_user_account_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UserData.ProtoReflect.Descriptor instead.
func (*UserData) Descriptor() ([]byte, []int) {
	return file_register_user_account_proto_rawDescGZIP(), []int{2}
}

func (x *UserData) GetId() int64 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *UserData) GetUsername() string {
	if x != nil {
		return x.Username
	}
	return ""
}

func (x *UserData) GetFullName() string {
	if x != nil {
		return x.FullName
	}
	return ""
}

func (x *UserData) GetEmail() string {
	if x != nil {
		return x.Email
	}
	return ""
}

func (x *UserData) GetPasswordHash() string {
	if x != nil {
		return x.PasswordHash
	}
	return ""
}

func (x *UserData) GetPasswordSalt() string {
	if x != nil {
		return x.PasswordSalt
	}
	return ""
}

func (x *UserData) GetCountryCode() string {
	if x != nil {
		return x.CountryCode
	}
	return ""
}

func (x *UserData) GetRoleId() int64 {
	if x != nil {
		return x.RoleId
	}
	return 0
}

func (x *UserData) GetStatus() string {
	if x != nil {
		return x.Status
	}
	return ""
}

func (x *UserData) GetAccountType() int32 {
	if x != nil {
		return x.AccountType
	}
	return 0
}

func (x *UserData) GetAccountName() string {
	if x != nil {
		return x.AccountName
	}
	return ""
}

func (x *UserData) GetAccountBio() string {
	if x != nil {
		return x.AccountBio
	}
	return ""
}

func (x *UserData) GetAccountPlan() int32 {
	if x != nil {
		return x.AccountPlan
	}
	return 0
}

func (x *UserData) GetAccountAvatarUri() string {
	if x != nil {
		return x.AccountAvatarUri
	}
	return ""
}

func (x *UserData) GetLastLoginAt() *timestamppb.Timestamp {
	if x != nil {
		return x.LastLoginAt
	}
	return nil
}

func (x *UserData) GetUsernameChangedAt() *timestamppb.Timestamp {
	if x != nil {
		return x.UsernameChangedAt
	}
	return nil
}

func (x *UserData) GetEmailChangedAt() *timestamppb.Timestamp {
	if x != nil {
		return x.EmailChangedAt
	}
	return nil
}

func (x *UserData) GetPasswordChangedAt() *timestamppb.Timestamp {
	if x != nil {
		return x.PasswordChangedAt
	}
	return nil
}

func (x *UserData) GetCreatedAt() *timestamppb.Timestamp {
	if x != nil {
		return x.CreatedAt
	}
	return nil
}

func (x *UserData) GetUpdatedAt() *timestamppb.Timestamp {
	if x != nil {
		return x.UpdatedAt
	}
	return nil
}

type UserAccount struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	User    *UserData    `protobuf:"bytes,1,opt,name=user,proto3" json:"user,omitempty"`
	Account *AccountData `protobuf:"bytes,2,opt,name=account,proto3" json:"account,omitempty"`
}

func (x *UserAccount) Reset() {
	*x = UserAccount{}
	if protoimpl.UnsafeEnabled {
		mi := &file_register_user_account_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UserAccount) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UserAccount) ProtoMessage() {}

func (x *UserAccount) ProtoReflect() protoreflect.Message {
	mi := &file_register_user_account_proto_msgTypes[3]
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
	return file_register_user_account_proto_rawDescGZIP(), []int{3}
}

func (x *UserAccount) GetUser() *UserData {
	if x != nil {
		return x.User
	}
	return nil
}

func (x *UserAccount) GetAccount() *AccountData {
	if x != nil {
		return x.Account
	}
	return nil
}

var File_register_user_account_proto protoreflect.FileDescriptor

var file_register_user_account_proto_rawDesc = []byte{
	0x0a, 0x1b, 0x72, 0x65, 0x67, 0x69, 0x73, 0x74, 0x65, 0x72, 0x2f, 0x75, 0x73, 0x65, 0x72, 0x5f,
	0x61, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x02, 0x70,
	0x62, 0x1a, 0x1f, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62,
	0x75, 0x66, 0x2f, 0x74, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x22, 0xda, 0x03, 0x0a, 0x0b, 0x41, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x44, 0x61,
	0x74, 0x61, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x02,
	0x69, 0x64, 0x12, 0x19, 0x0a, 0x08, 0x61, 0x63, 0x63, 0x5f, 0x74, 0x79, 0x70, 0x65, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x05, 0x52, 0x07, 0x61, 0x63, 0x63, 0x54, 0x79, 0x70, 0x65, 0x12, 0x21, 0x0a,
	0x0c, 0x61, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x03, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x0b, 0x61, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x4e, 0x61, 0x6d, 0x65,
	0x12, 0x14, 0x0a, 0x05, 0x6f, 0x77, 0x6e, 0x65, 0x72, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x05, 0x6f, 0x77, 0x6e, 0x65, 0x72, 0x12, 0x10, 0x0a, 0x03, 0x62, 0x69, 0x6f, 0x18, 0x05, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x03, 0x62, 0x69, 0x6f, 0x12, 0x16, 0x0a, 0x06, 0x73, 0x74, 0x61, 0x74,
	0x75, 0x73, 0x18, 0x06, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73,
	0x12, 0x12, 0x0a, 0x04, 0x70, 0x6c, 0x61, 0x6e, 0x18, 0x07, 0x20, 0x01, 0x28, 0x05, 0x52, 0x04,
	0x70, 0x6c, 0x61, 0x6e, 0x12, 0x1d, 0x0a, 0x0a, 0x61, 0x76, 0x61, 0x74, 0x61, 0x72, 0x5f, 0x75,
	0x72, 0x69, 0x18, 0x08, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x61, 0x76, 0x61, 0x74, 0x61, 0x72,
	0x55, 0x72, 0x69, 0x12, 0x14, 0x0a, 0x05, 0x70, 0x6c, 0x61, 0x79, 0x73, 0x18, 0x09, 0x20, 0x01,
	0x28, 0x03, 0x52, 0x05, 0x70, 0x6c, 0x61, 0x79, 0x73, 0x12, 0x14, 0x0a, 0x05, 0x6c, 0x69, 0x6b,
	0x65, 0x73, 0x18, 0x0a, 0x20, 0x01, 0x28, 0x03, 0x52, 0x05, 0x6c, 0x69, 0x6b, 0x65, 0x73, 0x12,
	0x18, 0x0a, 0x07, 0x66, 0x6f, 0x6c, 0x6c, 0x6f, 0x77, 0x73, 0x18, 0x0b, 0x20, 0x01, 0x28, 0x03,
	0x52, 0x07, 0x66, 0x6f, 0x6c, 0x6c, 0x6f, 0x77, 0x73, 0x12, 0x16, 0x0a, 0x06, 0x73, 0x68, 0x61,
	0x72, 0x65, 0x73, 0x18, 0x0c, 0x20, 0x01, 0x28, 0x03, 0x52, 0x06, 0x73, 0x68, 0x61, 0x72, 0x65,
	0x73, 0x12, 0x39, 0x0a, 0x0a, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x5f, 0x61, 0x74, 0x18,
	0x0d, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d,
	0x70, 0x52, 0x09, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x41, 0x74, 0x12, 0x39, 0x0a, 0x0a,
	0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x64, 0x5f, 0x61, 0x74, 0x18, 0x0e, 0x20, 0x01, 0x28, 0x0b,
	0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62,
	0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x52, 0x09, 0x75, 0x70,
	0x64, 0x61, 0x74, 0x65, 0x64, 0x41, 0x74, 0x12, 0x36, 0x0a, 0x0c, 0x61, 0x63, 0x63, 0x6f, 0x75,
	0x6e, 0x74, 0x5f, 0x74, 0x79, 0x70, 0x65, 0x18, 0x0f, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x13, 0x2e,
	0x70, 0x62, 0x2e, 0x41, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x54, 0x79, 0x70, 0x65, 0x44, 0x61,
	0x74, 0x61, 0x52, 0x0b, 0x61, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x54, 0x79, 0x70, 0x65, 0x22,
	0xe3, 0x02, 0x0a, 0x0f, 0x41, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x54, 0x79, 0x70, 0x65, 0x44,
	0x61, 0x74, 0x61, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52,
	0x02, 0x69, 0x64, 0x12, 0x19, 0x0a, 0x08, 0x61, 0x63, 0x63, 0x5f, 0x74, 0x79, 0x70, 0x65, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x05, 0x52, 0x07, 0x61, 0x63, 0x63, 0x54, 0x79, 0x70, 0x65, 0x12, 0x20,
	0x0a, 0x0b, 0x70, 0x65, 0x72, 0x6d, 0x69, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x73, 0x18, 0x03, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x0b, 0x70, 0x65, 0x72, 0x6d, 0x69, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x73,
	0x12, 0x1b, 0x0a, 0x09, 0x69, 0x73, 0x5f, 0x61, 0x72, 0x74, 0x69, 0x73, 0x74, 0x18, 0x04, 0x20,
	0x01, 0x28, 0x08, 0x52, 0x08, 0x69, 0x73, 0x41, 0x72, 0x74, 0x69, 0x73, 0x74, 0x12, 0x1f, 0x0a,
	0x0b, 0x69, 0x73, 0x5f, 0x70, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x65, 0x72, 0x18, 0x05, 0x20, 0x01,
	0x28, 0x08, 0x52, 0x0a, 0x69, 0x73, 0x50, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x65, 0x72, 0x12, 0x1b,
	0x0a, 0x09, 0x69, 0x73, 0x5f, 0x77, 0x72, 0x69, 0x74, 0x65, 0x72, 0x18, 0x06, 0x20, 0x01, 0x28,
	0x08, 0x52, 0x08, 0x69, 0x73, 0x57, 0x72, 0x69, 0x74, 0x65, 0x72, 0x12, 0x19, 0x0a, 0x08, 0x69,
	0x73, 0x5f, 0x6c, 0x61, 0x62, 0x65, 0x6c, 0x18, 0x07, 0x20, 0x01, 0x28, 0x08, 0x52, 0x07, 0x69,
	0x73, 0x4c, 0x61, 0x62, 0x65, 0x6c, 0x12, 0x17, 0x0a, 0x07, 0x69, 0x73, 0x5f, 0x75, 0x73, 0x65,
	0x72, 0x18, 0x08, 0x20, 0x01, 0x28, 0x08, 0x52, 0x06, 0x69, 0x73, 0x55, 0x73, 0x65, 0x72, 0x12,
	0x39, 0x0a, 0x0a, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x5f, 0x61, 0x74, 0x18, 0x09, 0x20,
	0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x52,
	0x09, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x41, 0x74, 0x12, 0x39, 0x0a, 0x0a, 0x75, 0x70,
	0x64, 0x61, 0x74, 0x65, 0x64, 0x5f, 0x61, 0x74, 0x18, 0x0a, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a,
	0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66,
	0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x52, 0x09, 0x75, 0x70, 0x64, 0x61,
	0x74, 0x65, 0x64, 0x41, 0x74, 0x22, 0xd3, 0x06, 0x0a, 0x08, 0x55, 0x73, 0x65, 0x72, 0x44, 0x61,
	0x74, 0x61, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x02,
	0x69, 0x64, 0x12, 0x1a, 0x0a, 0x08, 0x75, 0x73, 0x65, 0x72, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x75, 0x73, 0x65, 0x72, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x1b,
	0x0a, 0x09, 0x66, 0x75, 0x6c, 0x6c, 0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x08, 0x66, 0x75, 0x6c, 0x6c, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x14, 0x0a, 0x05, 0x65,
	0x6d, 0x61, 0x69, 0x6c, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x65, 0x6d, 0x61, 0x69,
	0x6c, 0x12, 0x23, 0x0a, 0x0d, 0x70, 0x61, 0x73, 0x73, 0x77, 0x6f, 0x72, 0x64, 0x5f, 0x68, 0x61,
	0x73, 0x68, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0c, 0x70, 0x61, 0x73, 0x73, 0x77, 0x6f,
	0x72, 0x64, 0x48, 0x61, 0x73, 0x68, 0x12, 0x23, 0x0a, 0x0d, 0x70, 0x61, 0x73, 0x73, 0x77, 0x6f,
	0x72, 0x64, 0x5f, 0x73, 0x61, 0x6c, 0x74, 0x18, 0x06, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0c, 0x70,
	0x61, 0x73, 0x73, 0x77, 0x6f, 0x72, 0x64, 0x53, 0x61, 0x6c, 0x74, 0x12, 0x21, 0x0a, 0x0c, 0x63,
	0x6f, 0x75, 0x6e, 0x74, 0x72, 0x79, 0x5f, 0x63, 0x6f, 0x64, 0x65, 0x18, 0x07, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x0b, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x72, 0x79, 0x43, 0x6f, 0x64, 0x65, 0x12, 0x17,
	0x0a, 0x07, 0x72, 0x6f, 0x6c, 0x65, 0x5f, 0x69, 0x64, 0x18, 0x08, 0x20, 0x01, 0x28, 0x03, 0x52,
	0x06, 0x72, 0x6f, 0x6c, 0x65, 0x49, 0x64, 0x12, 0x16, 0x0a, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75,
	0x73, 0x18, 0x09, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x12,
	0x21, 0x0a, 0x0c, 0x61, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x5f, 0x74, 0x79, 0x70, 0x65, 0x18,
	0x0a, 0x20, 0x01, 0x28, 0x05, 0x52, 0x0b, 0x61, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x54, 0x79,
	0x70, 0x65, 0x12, 0x21, 0x0a, 0x0c, 0x61, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x5f, 0x6e, 0x61,
	0x6d, 0x65, 0x18, 0x0b, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x61, 0x63, 0x63, 0x6f, 0x75, 0x6e,
	0x74, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x1f, 0x0a, 0x0b, 0x61, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74,
	0x5f, 0x62, 0x69, 0x6f, 0x18, 0x0c, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x61, 0x63, 0x63, 0x6f,
	0x75, 0x6e, 0x74, 0x42, 0x69, 0x6f, 0x12, 0x21, 0x0a, 0x0c, 0x61, 0x63, 0x63, 0x6f, 0x75, 0x6e,
	0x74, 0x5f, 0x70, 0x6c, 0x61, 0x6e, 0x18, 0x0d, 0x20, 0x01, 0x28, 0x05, 0x52, 0x0b, 0x61, 0x63,
	0x63, 0x6f, 0x75, 0x6e, 0x74, 0x50, 0x6c, 0x61, 0x6e, 0x12, 0x2c, 0x0a, 0x12, 0x61, 0x63, 0x63,
	0x6f, 0x75, 0x6e, 0x74, 0x5f, 0x61, 0x76, 0x61, 0x74, 0x61, 0x72, 0x5f, 0x75, 0x72, 0x69, 0x18,
	0x0e, 0x20, 0x01, 0x28, 0x09, 0x52, 0x10, 0x61, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x41, 0x76,
	0x61, 0x74, 0x61, 0x72, 0x55, 0x72, 0x69, 0x12, 0x3e, 0x0a, 0x0d, 0x6c, 0x61, 0x73, 0x74, 0x5f,
	0x6c, 0x6f, 0x67, 0x69, 0x6e, 0x5f, 0x61, 0x74, 0x18, 0x0f, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a,
	0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66,
	0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x52, 0x0b, 0x6c, 0x61, 0x73, 0x74,
	0x4c, 0x6f, 0x67, 0x69, 0x6e, 0x41, 0x74, 0x12, 0x4a, 0x0a, 0x13, 0x75, 0x73, 0x65, 0x72, 0x6e,
	0x61, 0x6d, 0x65, 0x5f, 0x63, 0x68, 0x61, 0x6e, 0x67, 0x65, 0x64, 0x5f, 0x61, 0x74, 0x18, 0x10,
	0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70,
	0x52, 0x11, 0x75, 0x73, 0x65, 0x72, 0x6e, 0x61, 0x6d, 0x65, 0x43, 0x68, 0x61, 0x6e, 0x67, 0x65,
	0x64, 0x41, 0x74, 0x12, 0x44, 0x0a, 0x10, 0x65, 0x6d, 0x61, 0x69, 0x6c, 0x5f, 0x63, 0x68, 0x61,
	0x6e, 0x67, 0x65, 0x64, 0x5f, 0x61, 0x74, 0x18, 0x11, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e,
	0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e,
	0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x52, 0x0e, 0x65, 0x6d, 0x61, 0x69, 0x6c,
	0x43, 0x68, 0x61, 0x6e, 0x67, 0x65, 0x64, 0x41, 0x74, 0x12, 0x4a, 0x0a, 0x13, 0x70, 0x61, 0x73,
	0x73, 0x77, 0x6f, 0x72, 0x64, 0x5f, 0x63, 0x68, 0x61, 0x6e, 0x67, 0x65, 0x64, 0x5f, 0x61, 0x74,
	0x18, 0x12, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61,
	0x6d, 0x70, 0x52, 0x11, 0x70, 0x61, 0x73, 0x73, 0x77, 0x6f, 0x72, 0x64, 0x43, 0x68, 0x61, 0x6e,
	0x67, 0x65, 0x64, 0x41, 0x74, 0x12, 0x39, 0x0a, 0x0a, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64,
	0x5f, 0x61, 0x74, 0x18, 0x13, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67,
	0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65,
	0x73, 0x74, 0x61, 0x6d, 0x70, 0x52, 0x09, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x41, 0x74,
	0x12, 0x39, 0x0a, 0x0a, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x64, 0x5f, 0x61, 0x74, 0x18, 0x14,
	0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70,
	0x52, 0x09, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x64, 0x41, 0x74, 0x22, 0x5a, 0x0a, 0x0b, 0x55,
	0x73, 0x65, 0x72, 0x41, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x12, 0x20, 0x0a, 0x04, 0x75, 0x73,
	0x65, 0x72, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0c, 0x2e, 0x70, 0x62, 0x2e, 0x55, 0x73,
	0x65, 0x72, 0x44, 0x61, 0x74, 0x61, 0x52, 0x04, 0x75, 0x73, 0x65, 0x72, 0x12, 0x29, 0x0a, 0x07,
	0x61, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0f, 0x2e,
	0x70, 0x62, 0x2e, 0x41, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x44, 0x61, 0x74, 0x61, 0x52, 0x07,
	0x61, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x42, 0x41, 0x5a, 0x3f, 0x67, 0x69, 0x74, 0x68, 0x75,
	0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x53, 0x74, 0x72, 0x65, 0x61, 0x6d, 0x66, 0x61, 0x69, 0x72,
	0x2f, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x5f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x49, 0x64,
	0x65, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x50, 0x72, 0x6f, 0x76, 0x69, 0x64, 0x65, 0x72, 0x2f, 0x70,
	0x62, 0x2f, 0x72, 0x65, 0x67, 0x69, 0x73, 0x74, 0x65, 0x72, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x33,
}

var (
	file_register_user_account_proto_rawDescOnce sync.Once
	file_register_user_account_proto_rawDescData = file_register_user_account_proto_rawDesc
)

func file_register_user_account_proto_rawDescGZIP() []byte {
	file_register_user_account_proto_rawDescOnce.Do(func() {
		file_register_user_account_proto_rawDescData = protoimpl.X.CompressGZIP(file_register_user_account_proto_rawDescData)
	})
	return file_register_user_account_proto_rawDescData
}

var file_register_user_account_proto_msgTypes = make([]protoimpl.MessageInfo, 4)
var file_register_user_account_proto_goTypes = []interface{}{
	(*AccountData)(nil),           // 0: pb.AccountData
	(*AccountTypeData)(nil),       // 1: pb.AccountTypeData
	(*UserData)(nil),              // 2: pb.UserData
	(*UserAccount)(nil),           // 3: pb.UserAccount
	(*timestamppb.Timestamp)(nil), // 4: google.protobuf.Timestamp
}
var file_register_user_account_proto_depIdxs = []int32{
	4,  // 0: pb.AccountData.created_at:type_name -> google.protobuf.Timestamp
	4,  // 1: pb.AccountData.updated_at:type_name -> google.protobuf.Timestamp
	1,  // 2: pb.AccountData.account_type:type_name -> pb.AccountTypeData
	4,  // 3: pb.AccountTypeData.created_at:type_name -> google.protobuf.Timestamp
	4,  // 4: pb.AccountTypeData.updated_at:type_name -> google.protobuf.Timestamp
	4,  // 5: pb.UserData.last_login_at:type_name -> google.protobuf.Timestamp
	4,  // 6: pb.UserData.username_changed_at:type_name -> google.protobuf.Timestamp
	4,  // 7: pb.UserData.email_changed_at:type_name -> google.protobuf.Timestamp
	4,  // 8: pb.UserData.password_changed_at:type_name -> google.protobuf.Timestamp
	4,  // 9: pb.UserData.created_at:type_name -> google.protobuf.Timestamp
	4,  // 10: pb.UserData.updated_at:type_name -> google.protobuf.Timestamp
	2,  // 11: pb.UserAccount.user:type_name -> pb.UserData
	0,  // 12: pb.UserAccount.account:type_name -> pb.AccountData
	13, // [13:13] is the sub-list for method output_type
	13, // [13:13] is the sub-list for method input_type
	13, // [13:13] is the sub-list for extension type_name
	13, // [13:13] is the sub-list for extension extendee
	0,  // [0:13] is the sub-list for field type_name
}

func init() { file_register_user_account_proto_init() }
func file_register_user_account_proto_init() {
	if File_register_user_account_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_register_user_account_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AccountData); i {
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
		file_register_user_account_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AccountTypeData); i {
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
		file_register_user_account_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UserData); i {
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
		file_register_user_account_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
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
			RawDescriptor: file_register_user_account_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   4,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_register_user_account_proto_goTypes,
		DependencyIndexes: file_register_user_account_proto_depIdxs,
		MessageInfos:      file_register_user_account_proto_msgTypes,
	}.Build()
	File_register_user_account_proto = out.File
	file_register_user_account_proto_rawDesc = nil
	file_register_user_account_proto_goTypes = nil
	file_register_user_account_proto_depIdxs = nil
}
