// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.32.0
// 	protoc        v4.25.2
// source: token/rpc_update_token.proto

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

type UpdateTokenRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id        int64   `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	UserId    *int64  `protobuf:"varint,2,opt,name=user_id,json=userId,proto3,oneof" json:"user_id,omitempty"`
	TokenType *string `protobuf:"bytes,3,opt,name=token_type,json=tokenType,proto3,oneof" json:"token_type,omitempty"`
}

func (x *UpdateTokenRequest) Reset() {
	*x = UpdateTokenRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_token_rpc_update_token_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UpdateTokenRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UpdateTokenRequest) ProtoMessage() {}

func (x *UpdateTokenRequest) ProtoReflect() protoreflect.Message {
	mi := &file_token_rpc_update_token_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UpdateTokenRequest.ProtoReflect.Descriptor instead.
func (*UpdateTokenRequest) Descriptor() ([]byte, []int) {
	return file_token_rpc_update_token_proto_rawDescGZIP(), []int{0}
}

func (x *UpdateTokenRequest) GetId() int64 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *UpdateTokenRequest) GetUserId() int64 {
	if x != nil && x.UserId != nil {
		return *x.UserId
	}
	return 0
}

func (x *UpdateTokenRequest) GetTokenType() string {
	if x != nil && x.TokenType != nil {
		return *x.TokenType
	}
	return ""
}

type UpdateTokenResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Token *Token `protobuf:"bytes,1,opt,name=token,proto3" json:"token,omitempty"`
}

func (x *UpdateTokenResponse) Reset() {
	*x = UpdateTokenResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_token_rpc_update_token_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UpdateTokenResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UpdateTokenResponse) ProtoMessage() {}

func (x *UpdateTokenResponse) ProtoReflect() protoreflect.Message {
	mi := &file_token_rpc_update_token_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UpdateTokenResponse.ProtoReflect.Descriptor instead.
func (*UpdateTokenResponse) Descriptor() ([]byte, []int) {
	return file_token_rpc_update_token_proto_rawDescGZIP(), []int{1}
}

func (x *UpdateTokenResponse) GetToken() *Token {
	if x != nil {
		return x.Token
	}
	return nil
}

var File_token_rpc_update_token_proto protoreflect.FileDescriptor

var file_token_rpc_update_token_proto_rawDesc = []byte{
	0x0a, 0x1c, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x2f, 0x72, 0x70, 0x63, 0x5f, 0x75, 0x70, 0x64, 0x61,
	0x74, 0x65, 0x5f, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x02,
	0x70, 0x62, 0x1a, 0x11, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x2f, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x81, 0x01, 0x0a, 0x12, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65,
	0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x0e, 0x0a, 0x02,
	0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x02, 0x69, 0x64, 0x12, 0x1c, 0x0a, 0x07,
	0x75, 0x73, 0x65, 0x72, 0x5f, 0x69, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x03, 0x48, 0x00, 0x52,
	0x06, 0x75, 0x73, 0x65, 0x72, 0x49, 0x64, 0x88, 0x01, 0x01, 0x12, 0x22, 0x0a, 0x0a, 0x74, 0x6f,
	0x6b, 0x65, 0x6e, 0x5f, 0x74, 0x79, 0x70, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x48, 0x01,
	0x52, 0x09, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x54, 0x79, 0x70, 0x65, 0x88, 0x01, 0x01, 0x42, 0x0a,
	0x0a, 0x08, 0x5f, 0x75, 0x73, 0x65, 0x72, 0x5f, 0x69, 0x64, 0x42, 0x0d, 0x0a, 0x0b, 0x5f, 0x74,
	0x6f, 0x6b, 0x65, 0x6e, 0x5f, 0x74, 0x79, 0x70, 0x65, 0x22, 0x36, 0x0a, 0x13, 0x55, 0x70, 0x64,
	0x61, 0x74, 0x65, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65,
	0x12, 0x1f, 0x0a, 0x05, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32,
	0x09, 0x2e, 0x70, 0x62, 0x2e, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x52, 0x05, 0x74, 0x6f, 0x6b, 0x65,
	0x6e, 0x42, 0x3a, 0x5a, 0x38, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f,
	0x53, 0x74, 0x72, 0x65, 0x61, 0x6d, 0x66, 0x61, 0x69, 0x72, 0x2f, 0x63, 0x6f, 0x6d, 0x6d, 0x6f,
	0x6e, 0x5f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x53, 0x65, 0x72,
	0x76, 0x69, 0x63, 0x65, 0x2f, 0x70, 0x62, 0x2f, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x62, 0x06, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_token_rpc_update_token_proto_rawDescOnce sync.Once
	file_token_rpc_update_token_proto_rawDescData = file_token_rpc_update_token_proto_rawDesc
)

func file_token_rpc_update_token_proto_rawDescGZIP() []byte {
	file_token_rpc_update_token_proto_rawDescOnce.Do(func() {
		file_token_rpc_update_token_proto_rawDescData = protoimpl.X.CompressGZIP(file_token_rpc_update_token_proto_rawDescData)
	})
	return file_token_rpc_update_token_proto_rawDescData
}

var file_token_rpc_update_token_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_token_rpc_update_token_proto_goTypes = []interface{}{
	(*UpdateTokenRequest)(nil),  // 0: pb.UpdateTokenRequest
	(*UpdateTokenResponse)(nil), // 1: pb.UpdateTokenResponse
	(*Token)(nil),               // 2: pb.Token
}
var file_token_rpc_update_token_proto_depIdxs = []int32{
	2, // 0: pb.UpdateTokenResponse.token:type_name -> pb.Token
	1, // [1:1] is the sub-list for method output_type
	1, // [1:1] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_token_rpc_update_token_proto_init() }
func file_token_rpc_update_token_proto_init() {
	if File_token_rpc_update_token_proto != nil {
		return
	}
	file_token_token_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_token_rpc_update_token_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UpdateTokenRequest); i {
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
		file_token_rpc_update_token_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UpdateTokenResponse); i {
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
	file_token_rpc_update_token_proto_msgTypes[0].OneofWrappers = []interface{}{}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_token_rpc_update_token_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_token_rpc_update_token_proto_goTypes,
		DependencyIndexes: file_token_rpc_update_token_proto_depIdxs,
		MessageInfos:      file_token_rpc_update_token_proto_msgTypes,
	}.Build()
	File_token_rpc_update_token_proto = out.File
	file_token_rpc_update_token_proto_rawDesc = nil
	file_token_rpc_update_token_proto_goTypes = nil
	file_token_rpc_update_token_proto_depIdxs = nil
}
