// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.33.0
// 	protoc        v5.26.1
// source: account_type/rpc_create_account_type.proto

package account_type

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

type CreateAccountTypeRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Type        string `protobuf:"bytes,1,opt,name=type,proto3" json:"type,omitempty"`
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
		mi := &file_account_type_rpc_create_account_type_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreateAccountTypeRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateAccountTypeRequest) ProtoMessage() {}

func (x *CreateAccountTypeRequest) ProtoReflect() protoreflect.Message {
	mi := &file_account_type_rpc_create_account_type_proto_msgTypes[0]
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
	return file_account_type_rpc_create_account_type_proto_rawDescGZIP(), []int{0}
}

func (x *CreateAccountTypeRequest) GetType() string {
	if x != nil {
		return x.Type
	}
	return ""
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

type CreateAccountTypeResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Account *AccountType `protobuf:"bytes,1,opt,name=account,proto3" json:"account,omitempty"`
}

func (x *CreateAccountTypeResponse) Reset() {
	*x = CreateAccountTypeResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_account_type_rpc_create_account_type_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreateAccountTypeResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateAccountTypeResponse) ProtoMessage() {}

func (x *CreateAccountTypeResponse) ProtoReflect() protoreflect.Message {
	mi := &file_account_type_rpc_create_account_type_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateAccountTypeResponse.ProtoReflect.Descriptor instead.
func (*CreateAccountTypeResponse) Descriptor() ([]byte, []int) {
	return file_account_type_rpc_create_account_type_proto_rawDescGZIP(), []int{1}
}

func (x *CreateAccountTypeResponse) GetAccount() *AccountType {
	if x != nil {
		return x.Account
	}
	return nil
}

var File_account_type_rpc_create_account_type_proto protoreflect.FileDescriptor

var file_account_type_rpc_create_account_type_proto_rawDesc = []byte{
	0x0a, 0x2a, 0x61, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x5f, 0x74, 0x79, 0x70, 0x65, 0x2f, 0x72,
	0x70, 0x63, 0x5f, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x5f, 0x61, 0x63, 0x63, 0x6f, 0x75, 0x6e,
	0x74, 0x5f, 0x74, 0x79, 0x70, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x02, 0x70, 0x62,
	0x1a, 0x1f, 0x61, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x5f, 0x74, 0x79, 0x70, 0x65, 0x2f, 0x61,
	0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x5f, 0x74, 0x79, 0x70, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x22, 0xdf, 0x01, 0x0a, 0x18, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x41, 0x63, 0x63, 0x6f,
	0x75, 0x6e, 0x74, 0x54, 0x79, 0x70, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x12,
	0x0a, 0x04, 0x74, 0x79, 0x70, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x74, 0x79,
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
	0x73, 0x65, 0x72, 0x22, 0x46, 0x0a, 0x19, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x41, 0x63, 0x63,
	0x6f, 0x75, 0x6e, 0x74, 0x54, 0x79, 0x70, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65,
	0x12, 0x29, 0x0a, 0x07, 0x61, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x0b, 0x32, 0x0f, 0x2e, 0x70, 0x62, 0x2e, 0x41, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x54, 0x79,
	0x70, 0x65, 0x52, 0x07, 0x61, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x42, 0x43, 0x5a, 0x41, 0x67,
	0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x53, 0x74, 0x72, 0x65, 0x61, 0x6d,
	0x66, 0x61, 0x69, 0x72, 0x2f, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x5f, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x2f, 0x41, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65,
	0x2f, 0x70, 0x62, 0x2f, 0x61, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x5f, 0x74, 0x79, 0x70, 0x65,
	0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_account_type_rpc_create_account_type_proto_rawDescOnce sync.Once
	file_account_type_rpc_create_account_type_proto_rawDescData = file_account_type_rpc_create_account_type_proto_rawDesc
)

func file_account_type_rpc_create_account_type_proto_rawDescGZIP() []byte {
	file_account_type_rpc_create_account_type_proto_rawDescOnce.Do(func() {
		file_account_type_rpc_create_account_type_proto_rawDescData = protoimpl.X.CompressGZIP(file_account_type_rpc_create_account_type_proto_rawDescData)
	})
	return file_account_type_rpc_create_account_type_proto_rawDescData
}

var file_account_type_rpc_create_account_type_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_account_type_rpc_create_account_type_proto_goTypes = []interface{}{
	(*CreateAccountTypeRequest)(nil),  // 0: pb.CreateAccountTypeRequest
	(*CreateAccountTypeResponse)(nil), // 1: pb.CreateAccountTypeResponse
	(*AccountType)(nil),               // 2: pb.AccountType
}
var file_account_type_rpc_create_account_type_proto_depIdxs = []int32{
	2, // 0: pb.CreateAccountTypeResponse.account:type_name -> pb.AccountType
	1, // [1:1] is the sub-list for method output_type
	1, // [1:1] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_account_type_rpc_create_account_type_proto_init() }
func file_account_type_rpc_create_account_type_proto_init() {
	if File_account_type_rpc_create_account_type_proto != nil {
		return
	}
	file_account_type_account_type_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_account_type_rpc_create_account_type_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
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
		file_account_type_rpc_create_account_type_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CreateAccountTypeResponse); i {
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
			RawDescriptor: file_account_type_rpc_create_account_type_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_account_type_rpc_create_account_type_proto_goTypes,
		DependencyIndexes: file_account_type_rpc_create_account_type_proto_depIdxs,
		MessageInfos:      file_account_type_rpc_create_account_type_proto_msgTypes,
	}.Build()
	File_account_type_rpc_create_account_type_proto = out.File
	file_account_type_rpc_create_account_type_proto_rawDesc = nil
	file_account_type_rpc_create_account_type_proto_goTypes = nil
	file_account_type_rpc_create_account_type_proto_depIdxs = nil
}
