// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.33.0
// 	protoc        v5.26.1
// source: refresh_token/rpc_list_revoked_refresh_tokens.proto

package refresh_token

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

type ListRevokedRefreshTokensRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Limit  int32 `protobuf:"varint,1,opt,name=limit,proto3" json:"limit,omitempty"`
	Offset int32 `protobuf:"varint,2,opt,name=offset,proto3" json:"offset,omitempty"`
}

func (x *ListRevokedRefreshTokensRequest) Reset() {
	*x = ListRevokedRefreshTokensRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_refresh_token_rpc_list_revoked_refresh_tokens_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ListRevokedRefreshTokensRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListRevokedRefreshTokensRequest) ProtoMessage() {}

func (x *ListRevokedRefreshTokensRequest) ProtoReflect() protoreflect.Message {
	mi := &file_refresh_token_rpc_list_revoked_refresh_tokens_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListRevokedRefreshTokensRequest.ProtoReflect.Descriptor instead.
func (*ListRevokedRefreshTokensRequest) Descriptor() ([]byte, []int) {
	return file_refresh_token_rpc_list_revoked_refresh_tokens_proto_rawDescGZIP(), []int{0}
}

func (x *ListRevokedRefreshTokensRequest) GetLimit() int32 {
	if x != nil {
		return x.Limit
	}
	return 0
}

func (x *ListRevokedRefreshTokensRequest) GetOffset() int32 {
	if x != nil {
		return x.Offset
	}
	return 0
}

type ListRevokedRefreshTokensResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	RefreshTokens []*RefreshToken `protobuf:"bytes,1,rep,name=refresh_tokens,json=refreshTokens,proto3" json:"refresh_tokens,omitempty"`
}

func (x *ListRevokedRefreshTokensResponse) Reset() {
	*x = ListRevokedRefreshTokensResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_refresh_token_rpc_list_revoked_refresh_tokens_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ListRevokedRefreshTokensResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListRevokedRefreshTokensResponse) ProtoMessage() {}

func (x *ListRevokedRefreshTokensResponse) ProtoReflect() protoreflect.Message {
	mi := &file_refresh_token_rpc_list_revoked_refresh_tokens_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListRevokedRefreshTokensResponse.ProtoReflect.Descriptor instead.
func (*ListRevokedRefreshTokensResponse) Descriptor() ([]byte, []int) {
	return file_refresh_token_rpc_list_revoked_refresh_tokens_proto_rawDescGZIP(), []int{1}
}

func (x *ListRevokedRefreshTokensResponse) GetRefreshTokens() []*RefreshToken {
	if x != nil {
		return x.RefreshTokens
	}
	return nil
}

var File_refresh_token_rpc_list_revoked_refresh_tokens_proto protoreflect.FileDescriptor

var file_refresh_token_rpc_list_revoked_refresh_tokens_proto_rawDesc = []byte{
	0x0a, 0x33, 0x72, 0x65, 0x66, 0x72, 0x65, 0x73, 0x68, 0x5f, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x2f,
	0x72, 0x70, 0x63, 0x5f, 0x6c, 0x69, 0x73, 0x74, 0x5f, 0x72, 0x65, 0x76, 0x6f, 0x6b, 0x65, 0x64,
	0x5f, 0x72, 0x65, 0x66, 0x72, 0x65, 0x73, 0x68, 0x5f, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x73, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x02, 0x70, 0x62, 0x1a, 0x21, 0x72, 0x65, 0x66, 0x72, 0x65,
	0x73, 0x68, 0x5f, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x2f, 0x72, 0x65, 0x66, 0x72, 0x65, 0x73, 0x68,
	0x5f, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x4f, 0x0a, 0x1f,
	0x4c, 0x69, 0x73, 0x74, 0x52, 0x65, 0x76, 0x6f, 0x6b, 0x65, 0x64, 0x52, 0x65, 0x66, 0x72, 0x65,
	0x73, 0x68, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12,
	0x14, 0x0a, 0x05, 0x6c, 0x69, 0x6d, 0x69, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x05,
	0x6c, 0x69, 0x6d, 0x69, 0x74, 0x12, 0x16, 0x0a, 0x06, 0x6f, 0x66, 0x66, 0x73, 0x65, 0x74, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x05, 0x52, 0x06, 0x6f, 0x66, 0x66, 0x73, 0x65, 0x74, 0x22, 0x5b, 0x0a,
	0x20, 0x4c, 0x69, 0x73, 0x74, 0x52, 0x65, 0x76, 0x6f, 0x6b, 0x65, 0x64, 0x52, 0x65, 0x66, 0x72,
	0x65, 0x73, 0x68, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73,
	0x65, 0x12, 0x37, 0x0a, 0x0e, 0x72, 0x65, 0x66, 0x72, 0x65, 0x73, 0x68, 0x5f, 0x74, 0x6f, 0x6b,
	0x65, 0x6e, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x10, 0x2e, 0x70, 0x62, 0x2e, 0x52,
	0x65, 0x66, 0x72, 0x65, 0x73, 0x68, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x52, 0x0d, 0x72, 0x65, 0x66,
	0x72, 0x65, 0x73, 0x68, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x73, 0x42, 0x42, 0x5a, 0x40, 0x67, 0x69,
	0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x53, 0x74, 0x72, 0x65, 0x61, 0x6d, 0x66,
	0x61, 0x69, 0x72, 0x2f, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x5f, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x2f, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2f, 0x70, 0x62,
	0x2f, 0x72, 0x65, 0x66, 0x72, 0x65, 0x73, 0x68, 0x5f, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x62, 0x06,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_refresh_token_rpc_list_revoked_refresh_tokens_proto_rawDescOnce sync.Once
	file_refresh_token_rpc_list_revoked_refresh_tokens_proto_rawDescData = file_refresh_token_rpc_list_revoked_refresh_tokens_proto_rawDesc
)

func file_refresh_token_rpc_list_revoked_refresh_tokens_proto_rawDescGZIP() []byte {
	file_refresh_token_rpc_list_revoked_refresh_tokens_proto_rawDescOnce.Do(func() {
		file_refresh_token_rpc_list_revoked_refresh_tokens_proto_rawDescData = protoimpl.X.CompressGZIP(file_refresh_token_rpc_list_revoked_refresh_tokens_proto_rawDescData)
	})
	return file_refresh_token_rpc_list_revoked_refresh_tokens_proto_rawDescData
}

var file_refresh_token_rpc_list_revoked_refresh_tokens_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_refresh_token_rpc_list_revoked_refresh_tokens_proto_goTypes = []interface{}{
	(*ListRevokedRefreshTokensRequest)(nil),  // 0: pb.ListRevokedRefreshTokensRequest
	(*ListRevokedRefreshTokensResponse)(nil), // 1: pb.ListRevokedRefreshTokensResponse
	(*RefreshToken)(nil),                     // 2: pb.RefreshToken
}
var file_refresh_token_rpc_list_revoked_refresh_tokens_proto_depIdxs = []int32{
	2, // 0: pb.ListRevokedRefreshTokensResponse.refresh_tokens:type_name -> pb.RefreshToken
	1, // [1:1] is the sub-list for method output_type
	1, // [1:1] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_refresh_token_rpc_list_revoked_refresh_tokens_proto_init() }
func file_refresh_token_rpc_list_revoked_refresh_tokens_proto_init() {
	if File_refresh_token_rpc_list_revoked_refresh_tokens_proto != nil {
		return
	}
	file_refresh_token_refresh_token_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_refresh_token_rpc_list_revoked_refresh_tokens_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ListRevokedRefreshTokensRequest); i {
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
		file_refresh_token_rpc_list_revoked_refresh_tokens_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ListRevokedRefreshTokensResponse); i {
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
			RawDescriptor: file_refresh_token_rpc_list_revoked_refresh_tokens_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_refresh_token_rpc_list_revoked_refresh_tokens_proto_goTypes,
		DependencyIndexes: file_refresh_token_rpc_list_revoked_refresh_tokens_proto_depIdxs,
		MessageInfos:      file_refresh_token_rpc_list_revoked_refresh_tokens_proto_msgTypes,
	}.Build()
	File_refresh_token_rpc_list_revoked_refresh_tokens_proto = out.File
	file_refresh_token_rpc_list_revoked_refresh_tokens_proto_rawDesc = nil
	file_refresh_token_rpc_list_revoked_refresh_tokens_proto_goTypes = nil
	file_refresh_token_rpc_list_revoked_refresh_tokens_proto_depIdxs = nil
}
