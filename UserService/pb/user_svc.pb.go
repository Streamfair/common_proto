// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.32.0
// 	protoc        v4.25.2
// source: user_svc.proto

package pb

import (
	user "github.com/Streamfair/streamfair_user_svc/_common_proto/UserService/pb/user"
	_ "github.com/grpc-ecosystem/grpc-gateway/v2/protoc-gen-openapiv2/options"
	_ "google.golang.org/genproto/googleapis/api/annotations"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	emptypb "google.golang.org/protobuf/types/known/emptypb"
	reflect "reflect"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

var File_user_svc_proto protoreflect.FileDescriptor

var file_user_svc_proto_rawDesc = []byte{
	0x0a, 0x0e, 0x75, 0x73, 0x65, 0x72, 0x5f, 0x73, 0x76, 0x63, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x12, 0x02, 0x70, 0x62, 0x1a, 0x1c, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x61, 0x70, 0x69,
	0x2f, 0x61, 0x6e, 0x6e, 0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x1a, 0x1b, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x62, 0x75, 0x66, 0x2f, 0x65, 0x6d, 0x70, 0x74, 0x79, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x63, 0x2d, 0x67, 0x65, 0x6e, 0x2d, 0x6f, 0x70, 0x65, 0x6e,
	0x61, 0x70, 0x69, 0x76, 0x32, 0x2f, 0x6f, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2f, 0x61, 0x6e,
	0x6e, 0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a,
	0x1a, 0x75, 0x73, 0x65, 0x72, 0x2f, 0x72, 0x70, 0x63, 0x5f, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65,
	0x5f, 0x75, 0x73, 0x65, 0x72, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x20, 0x75, 0x73, 0x65,
	0x72, 0x2f, 0x72, 0x70, 0x63, 0x5f, 0x64, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x5f, 0x75, 0x73, 0x65,
	0x72, 0x5f, 0x62, 0x79, 0x5f, 0x69, 0x64, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x23, 0x75,
	0x73, 0x65, 0x72, 0x2f, 0x72, 0x70, 0x63, 0x5f, 0x64, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x5f, 0x75,
	0x73, 0x65, 0x72, 0x5f, 0x62, 0x79, 0x5f, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x1a, 0x1d, 0x75, 0x73, 0x65, 0x72, 0x2f, 0x72, 0x70, 0x63, 0x5f, 0x67, 0x65, 0x74,
	0x5f, 0x75, 0x73, 0x65, 0x72, 0x5f, 0x62, 0x79, 0x5f, 0x69, 0x64, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x1a, 0x20, 0x75, 0x73, 0x65, 0x72, 0x2f, 0x72, 0x70, 0x63, 0x5f, 0x67, 0x65, 0x74, 0x5f,
	0x75, 0x73, 0x65, 0x72, 0x5f, 0x62, 0x79, 0x5f, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x1a, 0x19, 0x75, 0x73, 0x65, 0x72, 0x2f, 0x72, 0x70, 0x63, 0x5f, 0x6c, 0x69,
	0x73, 0x74, 0x5f, 0x75, 0x73, 0x65, 0x72, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1a,
	0x75, 0x73, 0x65, 0x72, 0x2f, 0x72, 0x70, 0x63, 0x5f, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x5f,
	0x75, 0x73, 0x65, 0x72, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x32, 0xcc, 0x08, 0x0a, 0x0b, 0x55,
	0x73, 0x65, 0x72, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x8d, 0x01, 0x0a, 0x0a, 0x43,
	0x72, 0x65, 0x61, 0x74, 0x65, 0x55, 0x73, 0x65, 0x72, 0x12, 0x15, 0x2e, 0x70, 0x62, 0x2e, 0x43,
	0x72, 0x65, 0x61, 0x74, 0x65, 0x55, 0x73, 0x65, 0x72, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x1a, 0x16, 0x2e, 0x70, 0x62, 0x2e, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x55, 0x73, 0x65, 0x72,
	0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x50, 0x92, 0x41, 0x28, 0x12, 0x10, 0x61,
	0x64, 0x64, 0x20, 0x73, 0x75, 0x6d, 0x6d, 0x61, 0x72, 0x79, 0x20, 0x68, 0x65, 0x72, 0x65, 0x1a,
	0x14, 0x61, 0x64, 0x64, 0x20, 0x64, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e,
	0x20, 0x68, 0x65, 0x72, 0x65, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x1f, 0x3a, 0x01, 0x2a, 0x22, 0x1a,
	0x2f, 0x73, 0x74, 0x72, 0x65, 0x61, 0x6d, 0x66, 0x61, 0x69, 0x72, 0x2f, 0x76, 0x31, 0x2f, 0x63,
	0x72, 0x65, 0x61, 0x74, 0x65, 0x5f, 0x75, 0x73, 0x65, 0x72, 0x12, 0x9d, 0x01, 0x0a, 0x0e, 0x44,
	0x65, 0x6c, 0x65, 0x74, 0x65, 0x55, 0x73, 0x65, 0x72, 0x42, 0x79, 0x49, 0x64, 0x12, 0x19, 0x2e,
	0x70, 0x62, 0x2e, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x55, 0x73, 0x65, 0x72, 0x42, 0x79, 0x49,
	0x64, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x16, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c,
	0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79,
	0x22, 0x58, 0x92, 0x41, 0x28, 0x12, 0x10, 0x61, 0x64, 0x64, 0x20, 0x73, 0x75, 0x6d, 0x6d, 0x61,
	0x72, 0x79, 0x20, 0x68, 0x65, 0x72, 0x65, 0x1a, 0x14, 0x61, 0x64, 0x64, 0x20, 0x64, 0x65, 0x73,
	0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x20, 0x68, 0x65, 0x72, 0x65, 0x82, 0xd3, 0xe4,
	0x93, 0x02, 0x27, 0x2a, 0x25, 0x2f, 0x73, 0x74, 0x72, 0x65, 0x61, 0x6d, 0x66, 0x61, 0x69, 0x72,
	0x2f, 0x76, 0x31, 0x2f, 0x64, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x5f, 0x75, 0x73, 0x65, 0x72, 0x5f,
	0x62, 0x79, 0x5f, 0x69, 0x64, 0x2f, 0x7b, 0x69, 0x64, 0x7d, 0x12, 0xac, 0x01, 0x0a, 0x11, 0x44,
	0x65, 0x6c, 0x65, 0x74, 0x65, 0x55, 0x73, 0x65, 0x72, 0x42, 0x79, 0x56, 0x61, 0x6c, 0x75, 0x65,
	0x12, 0x1c, 0x2e, 0x70, 0x62, 0x2e, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x55, 0x73, 0x65, 0x72,
	0x42, 0x79, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x16,
	0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66,
	0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x22, 0x61, 0x92, 0x41, 0x28, 0x12, 0x10, 0x61, 0x64, 0x64,
	0x20, 0x73, 0x75, 0x6d, 0x6d, 0x61, 0x72, 0x79, 0x20, 0x68, 0x65, 0x72, 0x65, 0x1a, 0x14, 0x61,
	0x64, 0x64, 0x20, 0x64, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x20, 0x68,
	0x65, 0x72, 0x65, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x30, 0x2a, 0x2e, 0x2f, 0x73, 0x74, 0x72, 0x65,
	0x61, 0x6d, 0x66, 0x61, 0x69, 0x72, 0x2f, 0x76, 0x31, 0x2f, 0x64, 0x65, 0x6c, 0x65, 0x74, 0x65,
	0x5f, 0x75, 0x73, 0x65, 0x72, 0x5f, 0x62, 0x79, 0x5f, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x2f, 0x7b,
	0x75, 0x73, 0x65, 0x72, 0x6e, 0x61, 0x6d, 0x65, 0x7d, 0x12, 0x95, 0x01, 0x0a, 0x0b, 0x47, 0x65,
	0x74, 0x55, 0x73, 0x65, 0x72, 0x42, 0x79, 0x49, 0x64, 0x12, 0x16, 0x2e, 0x70, 0x62, 0x2e, 0x47,
	0x65, 0x74, 0x55, 0x73, 0x65, 0x72, 0x42, 0x79, 0x49, 0x64, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x1a, 0x17, 0x2e, 0x70, 0x62, 0x2e, 0x47, 0x65, 0x74, 0x55, 0x73, 0x65, 0x72, 0x42, 0x79,
	0x49, 0x64, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x55, 0x92, 0x41, 0x28, 0x12,
	0x10, 0x61, 0x64, 0x64, 0x20, 0x73, 0x75, 0x6d, 0x6d, 0x61, 0x72, 0x79, 0x20, 0x68, 0x65, 0x72,
	0x65, 0x1a, 0x14, 0x61, 0x64, 0x64, 0x20, 0x64, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69,
	0x6f, 0x6e, 0x20, 0x68, 0x65, 0x72, 0x65, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x24, 0x12, 0x22, 0x2f,
	0x73, 0x74, 0x72, 0x65, 0x61, 0x6d, 0x66, 0x61, 0x69, 0x72, 0x2f, 0x76, 0x31, 0x2f, 0x67, 0x65,
	0x74, 0x5f, 0x75, 0x73, 0x65, 0x72, 0x5f, 0x62, 0x79, 0x5f, 0x69, 0x64, 0x2f, 0x7b, 0x69, 0x64,
	0x7d, 0x12, 0xa7, 0x01, 0x0a, 0x0e, 0x47, 0x65, 0x74, 0x55, 0x73, 0x65, 0x72, 0x42, 0x79, 0x56,
	0x61, 0x6c, 0x75, 0x65, 0x12, 0x19, 0x2e, 0x70, 0x62, 0x2e, 0x47, 0x65, 0x74, 0x55, 0x73, 0x65,
	0x72, 0x42, 0x79, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a,
	0x1a, 0x2e, 0x70, 0x62, 0x2e, 0x47, 0x65, 0x74, 0x55, 0x73, 0x65, 0x72, 0x42, 0x79, 0x56, 0x61,
	0x6c, 0x75, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x5e, 0x92, 0x41, 0x28,
	0x12, 0x10, 0x61, 0x64, 0x64, 0x20, 0x73, 0x75, 0x6d, 0x6d, 0x61, 0x72, 0x79, 0x20, 0x68, 0x65,
	0x72, 0x65, 0x1a, 0x14, 0x61, 0x64, 0x64, 0x20, 0x64, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74,
	0x69, 0x6f, 0x6e, 0x20, 0x68, 0x65, 0x72, 0x65, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x2d, 0x12, 0x2b,
	0x2f, 0x73, 0x74, 0x72, 0x65, 0x61, 0x6d, 0x66, 0x61, 0x69, 0x72, 0x2f, 0x76, 0x31, 0x2f, 0x67,
	0x65, 0x74, 0x5f, 0x75, 0x73, 0x65, 0x72, 0x5f, 0x62, 0x79, 0x5f, 0x76, 0x61, 0x6c, 0x75, 0x65,
	0x2f, 0x7b, 0x75, 0x73, 0x65, 0x72, 0x6e, 0x61, 0x6d, 0x65, 0x7d, 0x12, 0x86, 0x01, 0x0a, 0x09,
	0x4c, 0x69, 0x73, 0x74, 0x55, 0x73, 0x65, 0x72, 0x73, 0x12, 0x14, 0x2e, 0x70, 0x62, 0x2e, 0x4c,
	0x69, 0x73, 0x74, 0x55, 0x73, 0x65, 0x72, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a,
	0x15, 0x2e, 0x70, 0x62, 0x2e, 0x4c, 0x69, 0x73, 0x74, 0x55, 0x73, 0x65, 0x72, 0x73, 0x52, 0x65,
	0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x4c, 0x92, 0x41, 0x28, 0x12, 0x10, 0x61, 0x64, 0x64,
	0x20, 0x73, 0x75, 0x6d, 0x6d, 0x61, 0x72, 0x79, 0x20, 0x68, 0x65, 0x72, 0x65, 0x1a, 0x14, 0x61,
	0x64, 0x64, 0x20, 0x64, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x20, 0x68,
	0x65, 0x72, 0x65, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x1b, 0x22, 0x19, 0x2f, 0x73, 0x74, 0x72, 0x65,
	0x61, 0x6d, 0x66, 0x61, 0x69, 0x72, 0x2f, 0x76, 0x31, 0x2f, 0x6c, 0x69, 0x73, 0x74, 0x5f, 0x75,
	0x73, 0x65, 0x72, 0x73, 0x12, 0x92, 0x01, 0x0a, 0x0a, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x55,
	0x73, 0x65, 0x72, 0x12, 0x15, 0x2e, 0x70, 0x62, 0x2e, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x55,
	0x73, 0x65, 0x72, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x16, 0x2e, 0x70, 0x62, 0x2e,
	0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x55, 0x73, 0x65, 0x72, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e,
	0x73, 0x65, 0x22, 0x55, 0x92, 0x41, 0x28, 0x12, 0x10, 0x61, 0x64, 0x64, 0x20, 0x73, 0x75, 0x6d,
	0x6d, 0x61, 0x72, 0x79, 0x20, 0x68, 0x65, 0x72, 0x65, 0x1a, 0x14, 0x61, 0x64, 0x64, 0x20, 0x64,
	0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x20, 0x68, 0x65, 0x72, 0x65, 0x82,
	0xd3, 0xe4, 0x93, 0x02, 0x24, 0x3a, 0x01, 0x2a, 0x32, 0x1f, 0x2f, 0x73, 0x74, 0x72, 0x65, 0x61,
	0x6d, 0x66, 0x61, 0x69, 0x72, 0x2f, 0x76, 0x31, 0x2f, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x5f,
	0x75, 0x73, 0x65, 0x72, 0x2f, 0x7b, 0x69, 0x64, 0x7d, 0x42, 0xa9, 0x01, 0x92, 0x41, 0x5e, 0x12,
	0x5c, 0x0a, 0x17, 0x53, 0x74, 0x72, 0x65, 0x61, 0x6d, 0x66, 0x61, 0x69, 0x72, 0x20, 0x55, 0x73,
	0x65, 0x72, 0x20, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x22, 0x3c, 0x0a, 0x0a, 0x53, 0x74,
	0x72, 0x65, 0x61, 0x6d, 0x66, 0x61, 0x69, 0x72, 0x12, 0x1d, 0x68, 0x74, 0x74, 0x70, 0x73, 0x3a,
	0x2f, 0x2f, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x53, 0x74, 0x72,
	0x65, 0x61, 0x6d, 0x66, 0x61, 0x69, 0x72, 0x1a, 0x0f, 0x6e, 0x65, 0x6c, 0x69, 0x78, 0x40, 0x70,
	0x6f, 0x73, 0x74, 0x65, 0x6f, 0x2e, 0x64, 0x65, 0x32, 0x03, 0x31, 0x2e, 0x31, 0x5a, 0x46, 0x67,
	0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x53, 0x74, 0x72, 0x65, 0x61, 0x6d,
	0x66, 0x61, 0x69, 0x72, 0x2f, 0x73, 0x74, 0x72, 0x65, 0x61, 0x6d, 0x66, 0x61, 0x69, 0x72, 0x5f,
	0x75, 0x73, 0x65, 0x72, 0x5f, 0x73, 0x76, 0x63, 0x2f, 0x5f, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e,
	0x5f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x55, 0x73, 0x65, 0x72, 0x53, 0x65, 0x72, 0x76, 0x69,
	0x63, 0x65, 0x2f, 0x70, 0x62, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var file_user_svc_proto_goTypes = []interface{}{
	(*user.CreateUserRequest)(nil),        // 0: pb.CreateUserRequest
	(*user.DeleteUserByIdRequest)(nil),    // 1: pb.DeleteUserByIdRequest
	(*user.DeleteUserByValueRequest)(nil), // 2: pb.DeleteUserByValueRequest
	(*user.GetUserByIdRequest)(nil),       // 3: pb.GetUserByIdRequest
	(*user.GetUserByValueRequest)(nil),    // 4: pb.GetUserByValueRequest
	(*user.ListUsersRequest)(nil),         // 5: pb.ListUsersRequest
	(*user.UpdateUserRequest)(nil),        // 6: pb.UpdateUserRequest
	(*user.CreateUserResponse)(nil),       // 7: pb.CreateUserResponse
	(*emptypb.Empty)(nil),                 // 8: google.protobuf.Empty
	(*user.GetUserByIdResponse)(nil),      // 9: pb.GetUserByIdResponse
	(*user.GetUserByValueResponse)(nil),   // 10: pb.GetUserByValueResponse
	(*user.ListUsersResponse)(nil),        // 11: pb.ListUsersResponse
	(*user.UpdateUserResponse)(nil),       // 12: pb.UpdateUserResponse
}
var file_user_svc_proto_depIdxs = []int32{
	0,  // 0: pb.UserService.CreateUser:input_type -> pb.CreateUserRequest
	1,  // 1: pb.UserService.DeleteUserById:input_type -> pb.DeleteUserByIdRequest
	2,  // 2: pb.UserService.DeleteUserByValue:input_type -> pb.DeleteUserByValueRequest
	3,  // 3: pb.UserService.GetUserById:input_type -> pb.GetUserByIdRequest
	4,  // 4: pb.UserService.GetUserByValue:input_type -> pb.GetUserByValueRequest
	5,  // 5: pb.UserService.ListUsers:input_type -> pb.ListUsersRequest
	6,  // 6: pb.UserService.UpdateUser:input_type -> pb.UpdateUserRequest
	7,  // 7: pb.UserService.CreateUser:output_type -> pb.CreateUserResponse
	8,  // 8: pb.UserService.DeleteUserById:output_type -> google.protobuf.Empty
	8,  // 9: pb.UserService.DeleteUserByValue:output_type -> google.protobuf.Empty
	9,  // 10: pb.UserService.GetUserById:output_type -> pb.GetUserByIdResponse
	10, // 11: pb.UserService.GetUserByValue:output_type -> pb.GetUserByValueResponse
	11, // 12: pb.UserService.ListUsers:output_type -> pb.ListUsersResponse
	12, // 13: pb.UserService.UpdateUser:output_type -> pb.UpdateUserResponse
	7,  // [7:14] is the sub-list for method output_type
	0,  // [0:7] is the sub-list for method input_type
	0,  // [0:0] is the sub-list for extension type_name
	0,  // [0:0] is the sub-list for extension extendee
	0,  // [0:0] is the sub-list for field type_name
}

func init() { file_user_svc_proto_init() }
func file_user_svc_proto_init() {
	if File_user_svc_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_user_svc_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   0,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_user_svc_proto_goTypes,
		DependencyIndexes: file_user_svc_proto_depIdxs,
	}.Build()
	File_user_svc_proto = out.File
	file_user_svc_proto_rawDesc = nil
	file_user_svc_proto_goTypes = nil
	file_user_svc_proto_depIdxs = nil
}
