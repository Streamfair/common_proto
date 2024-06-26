// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.33.0
// 	protoc        v5.26.1
// source: token/rpc_list_revoked_tokens.proto

package token

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

type ListRevokedTokensRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Limit  int32 `protobuf:"varint,1,opt,name=limit,proto3" json:"limit,omitempty"`
	Offset int32 `protobuf:"varint,2,opt,name=offset,proto3" json:"offset,omitempty"`
}

func (x *ListRevokedTokensRequest) Reset() {
	*x = ListRevokedTokensRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_token_rpc_list_revoked_tokens_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ListRevokedTokensRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListRevokedTokensRequest) ProtoMessage() {}

func (x *ListRevokedTokensRequest) ProtoReflect() protoreflect.Message {
	mi := &file_token_rpc_list_revoked_tokens_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListRevokedTokensRequest.ProtoReflect.Descriptor instead.
func (*ListRevokedTokensRequest) Descriptor() ([]byte, []int) {
	return file_token_rpc_list_revoked_tokens_proto_rawDescGZIP(), []int{0}
}

func (x *ListRevokedTokensRequest) GetLimit() int32 {
	if x != nil {
		return x.Limit
	}
	return 0
}

func (x *ListRevokedTokensRequest) GetOffset() int32 {
	if x != nil {
		return x.Offset
	}
	return 0
}

type ListRevokedTokensResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Tokens []*Token `protobuf:"bytes,1,rep,name=tokens,proto3" json:"tokens,omitempty"`
}

func (x *ListRevokedTokensResponse) Reset() {
	*x = ListRevokedTokensResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_token_rpc_list_revoked_tokens_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ListRevokedTokensResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListRevokedTokensResponse) ProtoMessage() {}

func (x *ListRevokedTokensResponse) ProtoReflect() protoreflect.Message {
	mi := &file_token_rpc_list_revoked_tokens_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListRevokedTokensResponse.ProtoReflect.Descriptor instead.
func (*ListRevokedTokensResponse) Descriptor() ([]byte, []int) {
	return file_token_rpc_list_revoked_tokens_proto_rawDescGZIP(), []int{1}
}

func (x *ListRevokedTokensResponse) GetTokens() []*Token {
	if x != nil {
		return x.Tokens
	}
	return nil
}

var File_token_rpc_list_revoked_tokens_proto protoreflect.FileDescriptor

var file_token_rpc_list_revoked_tokens_proto_rawDesc = []byte{
	0x0a, 0x23, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x2f, 0x72, 0x70, 0x63, 0x5f, 0x6c, 0x69, 0x73, 0x74,
	0x5f, 0x72, 0x65, 0x76, 0x6f, 0x6b, 0x65, 0x64, 0x5f, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x73, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x02, 0x70, 0x62, 0x1a, 0x11, 0x74, 0x6f, 0x6b, 0x65, 0x6e,
	0x2f, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x48, 0x0a, 0x18,
	0x4c, 0x69, 0x73, 0x74, 0x52, 0x65, 0x76, 0x6f, 0x6b, 0x65, 0x64, 0x54, 0x6f, 0x6b, 0x65, 0x6e,
	0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x14, 0x0a, 0x05, 0x6c, 0x69, 0x6d, 0x69,
	0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x05, 0x6c, 0x69, 0x6d, 0x69, 0x74, 0x12, 0x16,
	0x0a, 0x06, 0x6f, 0x66, 0x66, 0x73, 0x65, 0x74, 0x18, 0x02, 0x20, 0x01, 0x28, 0x05, 0x52, 0x06,
	0x6f, 0x66, 0x66, 0x73, 0x65, 0x74, 0x22, 0x3e, 0x0a, 0x19, 0x4c, 0x69, 0x73, 0x74, 0x52, 0x65,
	0x76, 0x6f, 0x6b, 0x65, 0x64, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f,
	0x6e, 0x73, 0x65, 0x12, 0x21, 0x0a, 0x06, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x73, 0x18, 0x01, 0x20,
	0x03, 0x28, 0x0b, 0x32, 0x09, 0x2e, 0x70, 0x62, 0x2e, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x52, 0x06,
	0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x73, 0x42, 0x3a, 0x5a, 0x38, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62,
	0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x53, 0x74, 0x72, 0x65, 0x61, 0x6d, 0x66, 0x61, 0x69, 0x72, 0x2f,
	0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x5f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x54, 0x6f, 0x6b,
	0x65, 0x6e, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2f, 0x70, 0x62, 0x2f, 0x74, 0x6f, 0x6b,
	0x65, 0x6e, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_token_rpc_list_revoked_tokens_proto_rawDescOnce sync.Once
	file_token_rpc_list_revoked_tokens_proto_rawDescData = file_token_rpc_list_revoked_tokens_proto_rawDesc
)

func file_token_rpc_list_revoked_tokens_proto_rawDescGZIP() []byte {
	file_token_rpc_list_revoked_tokens_proto_rawDescOnce.Do(func() {
		file_token_rpc_list_revoked_tokens_proto_rawDescData = protoimpl.X.CompressGZIP(file_token_rpc_list_revoked_tokens_proto_rawDescData)
	})
	return file_token_rpc_list_revoked_tokens_proto_rawDescData
}

var file_token_rpc_list_revoked_tokens_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_token_rpc_list_revoked_tokens_proto_goTypes = []interface{}{
	(*ListRevokedTokensRequest)(nil),  // 0: pb.ListRevokedTokensRequest
	(*ListRevokedTokensResponse)(nil), // 1: pb.ListRevokedTokensResponse
	(*Token)(nil),                     // 2: pb.Token
}
var file_token_rpc_list_revoked_tokens_proto_depIdxs = []int32{
	2, // 0: pb.ListRevokedTokensResponse.tokens:type_name -> pb.Token
	1, // [1:1] is the sub-list for method output_type
	1, // [1:1] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_token_rpc_list_revoked_tokens_proto_init() }
func file_token_rpc_list_revoked_tokens_proto_init() {
	if File_token_rpc_list_revoked_tokens_proto != nil {
		return
	}
	file_token_token_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_token_rpc_list_revoked_tokens_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ListRevokedTokensRequest); i {
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
		file_token_rpc_list_revoked_tokens_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ListRevokedTokensResponse); i {
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
			RawDescriptor: file_token_rpc_list_revoked_tokens_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_token_rpc_list_revoked_tokens_proto_goTypes,
		DependencyIndexes: file_token_rpc_list_revoked_tokens_proto_depIdxs,
		MessageInfos:      file_token_rpc_list_revoked_tokens_proto_msgTypes,
	}.Build()
	File_token_rpc_list_revoked_tokens_proto = out.File
	file_token_rpc_list_revoked_tokens_proto_rawDesc = nil
	file_token_rpc_list_revoked_tokens_proto_goTypes = nil
	file_token_rpc_list_revoked_tokens_proto_depIdxs = nil
}
