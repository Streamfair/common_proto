// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.32.0
// 	protoc        v4.25.2
// source: user/rpc_delete_user_by_id.proto

package user

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

type DeleteUserByIdRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id int64 `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
}

func (x *DeleteUserByIdRequest) Reset() {
	*x = DeleteUserByIdRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_user_rpc_delete_user_by_id_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DeleteUserByIdRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DeleteUserByIdRequest) ProtoMessage() {}

func (x *DeleteUserByIdRequest) ProtoReflect() protoreflect.Message {
	mi := &file_user_rpc_delete_user_by_id_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DeleteUserByIdRequest.ProtoReflect.Descriptor instead.
func (*DeleteUserByIdRequest) Descriptor() ([]byte, []int) {
	return file_user_rpc_delete_user_by_id_proto_rawDescGZIP(), []int{0}
}

func (x *DeleteUserByIdRequest) GetId() int64 {
	if x != nil {
		return x.Id
	}
	return 0
}

var File_user_rpc_delete_user_by_id_proto protoreflect.FileDescriptor

var file_user_rpc_delete_user_by_id_proto_rawDesc = []byte{
	0x0a, 0x20, 0x75, 0x73, 0x65, 0x72, 0x2f, 0x72, 0x70, 0x63, 0x5f, 0x64, 0x65, 0x6c, 0x65, 0x74,
	0x65, 0x5f, 0x75, 0x73, 0x65, 0x72, 0x5f, 0x62, 0x79, 0x5f, 0x69, 0x64, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x12, 0x02, 0x70, 0x62, 0x22, 0x27, 0x0a, 0x15, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65,
	0x55, 0x73, 0x65, 0x72, 0x42, 0x79, 0x49, 0x64, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12,
	0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x02, 0x69, 0x64, 0x42,
	0x33, 0x5a, 0x31, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x53, 0x74,
	0x72, 0x65, 0x61, 0x6d, 0x66, 0x61, 0x69, 0x72, 0x2f, 0x73, 0x74, 0x72, 0x65, 0x61, 0x6d, 0x66,
	0x61, 0x69, 0x72, 0x5f, 0x75, 0x73, 0x65, 0x72, 0x5f, 0x73, 0x76, 0x63, 0x2f, 0x70, 0x62, 0x2f,
	0x75, 0x73, 0x65, 0x72, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_user_rpc_delete_user_by_id_proto_rawDescOnce sync.Once
	file_user_rpc_delete_user_by_id_proto_rawDescData = file_user_rpc_delete_user_by_id_proto_rawDesc
)

func file_user_rpc_delete_user_by_id_proto_rawDescGZIP() []byte {
	file_user_rpc_delete_user_by_id_proto_rawDescOnce.Do(func() {
		file_user_rpc_delete_user_by_id_proto_rawDescData = protoimpl.X.CompressGZIP(file_user_rpc_delete_user_by_id_proto_rawDescData)
	})
	return file_user_rpc_delete_user_by_id_proto_rawDescData
}

var file_user_rpc_delete_user_by_id_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_user_rpc_delete_user_by_id_proto_goTypes = []interface{}{
	(*DeleteUserByIdRequest)(nil), // 0: pb.DeleteUserByIdRequest
}
var file_user_rpc_delete_user_by_id_proto_depIdxs = []int32{
	0, // [0:0] is the sub-list for method output_type
	0, // [0:0] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_user_rpc_delete_user_by_id_proto_init() }
func file_user_rpc_delete_user_by_id_proto_init() {
	if File_user_rpc_delete_user_by_id_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_user_rpc_delete_user_by_id_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DeleteUserByIdRequest); i {
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
			RawDescriptor: file_user_rpc_delete_user_by_id_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_user_rpc_delete_user_by_id_proto_goTypes,
		DependencyIndexes: file_user_rpc_delete_user_by_id_proto_depIdxs,
		MessageInfos:      file_user_rpc_delete_user_by_id_proto_msgTypes,
	}.Build()
	File_user_rpc_delete_user_by_id_proto = out.File
	file_user_rpc_delete_user_by_id_proto_rawDesc = nil
	file_user_rpc_delete_user_by_id_proto_goTypes = nil
	file_user_rpc_delete_user_by_id_proto_depIdxs = nil
}
