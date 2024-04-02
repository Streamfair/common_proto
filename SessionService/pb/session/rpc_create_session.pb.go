// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.32.0
// 	protoc        v4.25.2
// source: session/rpc_create_session.proto

package session

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

type CreateSessionRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Uuid         string                 `protobuf:"bytes,1,opt,name=uuid,proto3" json:"uuid,omitempty"`
	Username     string                 `protobuf:"bytes,2,opt,name=username,proto3" json:"username,omitempty"`
	RefreshToken string                 `protobuf:"bytes,3,opt,name=refresh_token,json=refreshToken,proto3" json:"refresh_token,omitempty"`
	UserAgent    string                 `protobuf:"bytes,4,opt,name=user_agent,json=userAgent,proto3" json:"user_agent,omitempty"`
	ClientIp     string                 `protobuf:"bytes,5,opt,name=client_ip,json=clientIp,proto3" json:"client_ip,omitempty"`
	IsBlocked    bool                   `protobuf:"varint,6,opt,name=is_blocked,json=isBlocked,proto3" json:"is_blocked,omitempty"`
	ExpiresAt    *timestamppb.Timestamp `protobuf:"bytes,7,opt,name=expires_at,json=expiresAt,proto3" json:"expires_at,omitempty"`
}

func (x *CreateSessionRequest) Reset() {
	*x = CreateSessionRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_session_rpc_create_session_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreateSessionRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateSessionRequest) ProtoMessage() {}

func (x *CreateSessionRequest) ProtoReflect() protoreflect.Message {
	mi := &file_session_rpc_create_session_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateSessionRequest.ProtoReflect.Descriptor instead.
func (*CreateSessionRequest) Descriptor() ([]byte, []int) {
	return file_session_rpc_create_session_proto_rawDescGZIP(), []int{0}
}

func (x *CreateSessionRequest) GetUuid() string {
	if x != nil {
		return x.Uuid
	}
	return ""
}

func (x *CreateSessionRequest) GetUsername() string {
	if x != nil {
		return x.Username
	}
	return ""
}

func (x *CreateSessionRequest) GetRefreshToken() string {
	if x != nil {
		return x.RefreshToken
	}
	return ""
}

func (x *CreateSessionRequest) GetUserAgent() string {
	if x != nil {
		return x.UserAgent
	}
	return ""
}

func (x *CreateSessionRequest) GetClientIp() string {
	if x != nil {
		return x.ClientIp
	}
	return ""
}

func (x *CreateSessionRequest) GetIsBlocked() bool {
	if x != nil {
		return x.IsBlocked
	}
	return false
}

func (x *CreateSessionRequest) GetExpiresAt() *timestamppb.Timestamp {
	if x != nil {
		return x.ExpiresAt
	}
	return nil
}

type CreateSessionResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Session *Session `protobuf:"bytes,1,opt,name=session,proto3" json:"session,omitempty"`
}

func (x *CreateSessionResponse) Reset() {
	*x = CreateSessionResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_session_rpc_create_session_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreateSessionResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateSessionResponse) ProtoMessage() {}

func (x *CreateSessionResponse) ProtoReflect() protoreflect.Message {
	mi := &file_session_rpc_create_session_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateSessionResponse.ProtoReflect.Descriptor instead.
func (*CreateSessionResponse) Descriptor() ([]byte, []int) {
	return file_session_rpc_create_session_proto_rawDescGZIP(), []int{1}
}

func (x *CreateSessionResponse) GetSession() *Session {
	if x != nil {
		return x.Session
	}
	return nil
}

var File_session_rpc_create_session_proto protoreflect.FileDescriptor

var file_session_rpc_create_session_proto_rawDesc = []byte{
	0x0a, 0x20, 0x73, 0x65, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x2f, 0x72, 0x70, 0x63, 0x5f, 0x63, 0x72,
	0x65, 0x61, 0x74, 0x65, 0x5f, 0x73, 0x65, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x12, 0x02, 0x70, 0x62, 0x1a, 0x1f, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x74, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d,
	0x70, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x15, 0x73, 0x65, 0x73, 0x73, 0x69, 0x6f, 0x6e,
	0x2f, 0x73, 0x65, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x81,
	0x02, 0x0a, 0x14, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x53, 0x65, 0x73, 0x73, 0x69, 0x6f, 0x6e,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x12, 0x0a, 0x04, 0x75, 0x75, 0x69, 0x64, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x75, 0x75, 0x69, 0x64, 0x12, 0x1a, 0x0a, 0x08, 0x75,
	0x73, 0x65, 0x72, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x75,
	0x73, 0x65, 0x72, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x23, 0x0a, 0x0d, 0x72, 0x65, 0x66, 0x72, 0x65,
	0x73, 0x68, 0x5f, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0c,
	0x72, 0x65, 0x66, 0x72, 0x65, 0x73, 0x68, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x12, 0x1d, 0x0a, 0x0a,
	0x75, 0x73, 0x65, 0x72, 0x5f, 0x61, 0x67, 0x65, 0x6e, 0x74, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x09, 0x75, 0x73, 0x65, 0x72, 0x41, 0x67, 0x65, 0x6e, 0x74, 0x12, 0x1b, 0x0a, 0x09, 0x63,
	0x6c, 0x69, 0x65, 0x6e, 0x74, 0x5f, 0x69, 0x70, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08,
	0x63, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x49, 0x70, 0x12, 0x1d, 0x0a, 0x0a, 0x69, 0x73, 0x5f, 0x62,
	0x6c, 0x6f, 0x63, 0x6b, 0x65, 0x64, 0x18, 0x06, 0x20, 0x01, 0x28, 0x08, 0x52, 0x09, 0x69, 0x73,
	0x42, 0x6c, 0x6f, 0x63, 0x6b, 0x65, 0x64, 0x12, 0x39, 0x0a, 0x0a, 0x65, 0x78, 0x70, 0x69, 0x72,
	0x65, 0x73, 0x5f, 0x61, 0x74, 0x18, 0x07, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f,
	0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54, 0x69,
	0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x52, 0x09, 0x65, 0x78, 0x70, 0x69, 0x72, 0x65, 0x73,
	0x41, 0x74, 0x22, 0x3e, 0x0a, 0x15, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x53, 0x65, 0x73, 0x73,
	0x69, 0x6f, 0x6e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x25, 0x0a, 0x07, 0x73,
	0x65, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0b, 0x2e, 0x70,
	0x62, 0x2e, 0x53, 0x65, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x52, 0x07, 0x73, 0x65, 0x73, 0x73, 0x69,
	0x6f, 0x6e, 0x42, 0x55, 0x5a, 0x53, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d,
	0x2f, 0x53, 0x74, 0x72, 0x65, 0x61, 0x6d, 0x66, 0x61, 0x69, 0x72, 0x2f, 0x73, 0x74, 0x72, 0x65,
	0x61, 0x6d, 0x66, 0x61, 0x69, 0x72, 0x5f, 0x73, 0x65, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x5f, 0x73,
	0x76, 0x63, 0x2f, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x5f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f,
	0x53, 0x65, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2f, 0x70,
	0x62, 0x2f, 0x73, 0x65, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x33,
}

var (
	file_session_rpc_create_session_proto_rawDescOnce sync.Once
	file_session_rpc_create_session_proto_rawDescData = file_session_rpc_create_session_proto_rawDesc
)

func file_session_rpc_create_session_proto_rawDescGZIP() []byte {
	file_session_rpc_create_session_proto_rawDescOnce.Do(func() {
		file_session_rpc_create_session_proto_rawDescData = protoimpl.X.CompressGZIP(file_session_rpc_create_session_proto_rawDescData)
	})
	return file_session_rpc_create_session_proto_rawDescData
}

var file_session_rpc_create_session_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_session_rpc_create_session_proto_goTypes = []interface{}{
	(*CreateSessionRequest)(nil),  // 0: pb.CreateSessionRequest
	(*CreateSessionResponse)(nil), // 1: pb.CreateSessionResponse
	(*timestamppb.Timestamp)(nil), // 2: google.protobuf.Timestamp
	(*Session)(nil),               // 3: pb.Session
}
var file_session_rpc_create_session_proto_depIdxs = []int32{
	2, // 0: pb.CreateSessionRequest.expires_at:type_name -> google.protobuf.Timestamp
	3, // 1: pb.CreateSessionResponse.session:type_name -> pb.Session
	2, // [2:2] is the sub-list for method output_type
	2, // [2:2] is the sub-list for method input_type
	2, // [2:2] is the sub-list for extension type_name
	2, // [2:2] is the sub-list for extension extendee
	0, // [0:2] is the sub-list for field type_name
}

func init() { file_session_rpc_create_session_proto_init() }
func file_session_rpc_create_session_proto_init() {
	if File_session_rpc_create_session_proto != nil {
		return
	}
	file_session_session_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_session_rpc_create_session_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CreateSessionRequest); i {
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
		file_session_rpc_create_session_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CreateSessionResponse); i {
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
			RawDescriptor: file_session_rpc_create_session_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_session_rpc_create_session_proto_goTypes,
		DependencyIndexes: file_session_rpc_create_session_proto_depIdxs,
		MessageInfos:      file_session_rpc_create_session_proto_msgTypes,
	}.Build()
	File_session_rpc_create_session_proto = out.File
	file_session_rpc_create_session_proto_rawDesc = nil
	file_session_rpc_create_session_proto_goTypes = nil
	file_session_rpc_create_session_proto_depIdxs = nil
}
