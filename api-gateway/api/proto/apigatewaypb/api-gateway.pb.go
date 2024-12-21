// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.35.2
// 	protoc        v5.27.0
// source: github.com/wazwki/WearStore/api-gateway/api/proto/api-gateway.proto

package apigatewaypb

import (
	cartpb "github.com/wazwki/WearStore/cart-service/api/proto/cartpb"
	productpb "github.com/wazwki/WearStore/product-service/api/proto/productpb"
	userpb "github.com/wazwki/WearStore/user-service/api/proto/userpb"
	_ "google.golang.org/genproto/googleapis/api/annotations"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	reflect "reflect"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

var File_github_com_wazwki_WearStore_api_gateway_api_proto_api_gateway_proto protoreflect.FileDescriptor

var file_github_com_wazwki_WearStore_api_gateway_api_proto_api_gateway_proto_rawDesc = []byte{
	0x0a, 0x43, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x77, 0x61, 0x7a,
	0x77, 0x6b, 0x69, 0x2f, 0x57, 0x65, 0x61, 0x72, 0x53, 0x74, 0x6f, 0x72, 0x65, 0x2f, 0x61, 0x70,
	0x69, 0x2d, 0x67, 0x61, 0x74, 0x65, 0x77, 0x61, 0x79, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x2f, 0x61, 0x70, 0x69, 0x2d, 0x67, 0x61, 0x74, 0x65, 0x77, 0x61, 0x79, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x0c, 0x61, 0x70, 0x69, 0x67, 0x61, 0x74, 0x65, 0x77, 0x61,
	0x79, 0x70, 0x62, 0x1a, 0x1c, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x61, 0x70, 0x69, 0x2f,
	0x61, 0x6e, 0x6e, 0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x1a, 0x3d, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x77, 0x61,
	0x7a, 0x77, 0x6b, 0x69, 0x2f, 0x57, 0x65, 0x61, 0x72, 0x53, 0x74, 0x6f, 0x72, 0x65, 0x2f, 0x75,
	0x73, 0x65, 0x72, 0x2d, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2f, 0x61, 0x70, 0x69, 0x2f,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x75, 0x73, 0x65, 0x72, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x1a, 0x43, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x77, 0x61, 0x7a,
	0x77, 0x6b, 0x69, 0x2f, 0x57, 0x65, 0x61, 0x72, 0x53, 0x74, 0x6f, 0x72, 0x65, 0x2f, 0x70, 0x72,
	0x6f, 0x64, 0x75, 0x63, 0x74, 0x2d, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2f, 0x61, 0x70,
	0x69, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x70, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x3d, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f,
	0x6d, 0x2f, 0x77, 0x61, 0x7a, 0x77, 0x6b, 0x69, 0x2f, 0x57, 0x65, 0x61, 0x72, 0x53, 0x74, 0x6f,
	0x72, 0x65, 0x2f, 0x63, 0x61, 0x72, 0x74, 0x2d, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2f,
	0x61, 0x70, 0x69, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x63, 0x61, 0x72, 0x74, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x32, 0xd8, 0x0b, 0x0a, 0x11, 0x41, 0x70, 0x69, 0x47, 0x61, 0x74, 0x65,
	0x77, 0x61, 0x79, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x68, 0x0a, 0x0c, 0x52, 0x65,
	0x67, 0x69, 0x73, 0x74, 0x65, 0x72, 0x55, 0x73, 0x65, 0x72, 0x12, 0x1b, 0x2e, 0x75, 0x73, 0x65,
	0x72, 0x70, 0x62, 0x2e, 0x52, 0x65, 0x67, 0x69, 0x73, 0x74, 0x65, 0x72, 0x55, 0x73, 0x65, 0x72,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1c, 0x2e, 0x75, 0x73, 0x65, 0x72, 0x70, 0x62,
	0x2e, 0x52, 0x65, 0x67, 0x69, 0x73, 0x74, 0x65, 0x72, 0x55, 0x73, 0x65, 0x72, 0x52, 0x65, 0x73,
	0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x1d, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x17, 0x3a, 0x01, 0x2a,
	0x22, 0x12, 0x2f, 0x76, 0x31, 0x2f, 0x75, 0x73, 0x65, 0x72, 0x73, 0x2f, 0x72, 0x65, 0x67, 0x69,
	0x73, 0x74, 0x65, 0x72, 0x12, 0x5c, 0x0a, 0x09, 0x4c, 0x6f, 0x67, 0x69, 0x6e, 0x55, 0x73, 0x65,
	0x72, 0x12, 0x18, 0x2e, 0x75, 0x73, 0x65, 0x72, 0x70, 0x62, 0x2e, 0x4c, 0x6f, 0x67, 0x69, 0x6e,
	0x55, 0x73, 0x65, 0x72, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x19, 0x2e, 0x75, 0x73,
	0x65, 0x72, 0x70, 0x62, 0x2e, 0x4c, 0x6f, 0x67, 0x69, 0x6e, 0x55, 0x73, 0x65, 0x72, 0x52, 0x65,
	0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x1a, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x14, 0x3a, 0x01,
	0x2a, 0x22, 0x0f, 0x2f, 0x76, 0x31, 0x2f, 0x75, 0x73, 0x65, 0x72, 0x73, 0x2f, 0x6c, 0x6f, 0x67,
	0x69, 0x6e, 0x12, 0x52, 0x0a, 0x07, 0x47, 0x65, 0x74, 0x55, 0x73, 0x65, 0x72, 0x12, 0x16, 0x2e,
	0x75, 0x73, 0x65, 0x72, 0x70, 0x62, 0x2e, 0x47, 0x65, 0x74, 0x55, 0x73, 0x65, 0x72, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x17, 0x2e, 0x75, 0x73, 0x65, 0x72, 0x70, 0x62, 0x2e, 0x47,
	0x65, 0x74, 0x55, 0x73, 0x65, 0x72, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x16,
	0x82, 0xd3, 0xe4, 0x93, 0x02, 0x10, 0x12, 0x0e, 0x2f, 0x76, 0x31, 0x2f, 0x75, 0x73, 0x65, 0x72,
	0x73, 0x2f, 0x7b, 0x69, 0x64, 0x7d, 0x12, 0x5e, 0x0a, 0x0a, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65,
	0x55, 0x73, 0x65, 0x72, 0x12, 0x19, 0x2e, 0x75, 0x73, 0x65, 0x72, 0x70, 0x62, 0x2e, 0x55, 0x70,
	0x64, 0x61, 0x74, 0x65, 0x55, 0x73, 0x65, 0x72, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a,
	0x1a, 0x2e, 0x75, 0x73, 0x65, 0x72, 0x70, 0x62, 0x2e, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x55,
	0x73, 0x65, 0x72, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x19, 0x82, 0xd3, 0xe4,
	0x93, 0x02, 0x13, 0x3a, 0x01, 0x2a, 0x1a, 0x0e, 0x2f, 0x76, 0x31, 0x2f, 0x75, 0x73, 0x65, 0x72,
	0x73, 0x2f, 0x7b, 0x69, 0x64, 0x7d, 0x12, 0x5b, 0x0a, 0x0a, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65,
	0x55, 0x73, 0x65, 0x72, 0x12, 0x19, 0x2e, 0x75, 0x73, 0x65, 0x72, 0x70, 0x62, 0x2e, 0x44, 0x65,
	0x6c, 0x65, 0x74, 0x65, 0x55, 0x73, 0x65, 0x72, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a,
	0x1a, 0x2e, 0x75, 0x73, 0x65, 0x72, 0x70, 0x62, 0x2e, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x55,
	0x73, 0x65, 0x72, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x16, 0x82, 0xd3, 0xe4,
	0x93, 0x02, 0x10, 0x2a, 0x0e, 0x2f, 0x76, 0x31, 0x2f, 0x75, 0x73, 0x65, 0x72, 0x73, 0x2f, 0x7b,
	0x69, 0x64, 0x7d, 0x12, 0x64, 0x0a, 0x0a, 0x47, 0x65, 0x74, 0x50, 0x72, 0x6f, 0x64, 0x75, 0x63,
	0x74, 0x12, 0x1c, 0x2e, 0x70, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x70, 0x62, 0x2e, 0x47, 0x65,
	0x74, 0x50, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a,
	0x1d, 0x2e, 0x70, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x70, 0x62, 0x2e, 0x47, 0x65, 0x74, 0x50,
	0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x19,
	0x82, 0xd3, 0xe4, 0x93, 0x02, 0x13, 0x12, 0x11, 0x2f, 0x76, 0x31, 0x2f, 0x70, 0x72, 0x6f, 0x64,
	0x75, 0x63, 0x74, 0x73, 0x2f, 0x7b, 0x69, 0x64, 0x7d, 0x12, 0x65, 0x0a, 0x0c, 0x4c, 0x69, 0x73,
	0x74, 0x50, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x73, 0x12, 0x1e, 0x2e, 0x70, 0x72, 0x6f, 0x64,
	0x75, 0x63, 0x74, 0x70, 0x62, 0x2e, 0x4c, 0x69, 0x73, 0x74, 0x50, 0x72, 0x6f, 0x64, 0x75, 0x63,
	0x74, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1f, 0x2e, 0x70, 0x72, 0x6f, 0x64,
	0x75, 0x63, 0x74, 0x70, 0x62, 0x2e, 0x4c, 0x69, 0x73, 0x74, 0x50, 0x72, 0x6f, 0x64, 0x75, 0x63,
	0x74, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x14, 0x82, 0xd3, 0xe4, 0x93,
	0x02, 0x0e, 0x12, 0x0c, 0x2f, 0x76, 0x31, 0x2f, 0x70, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x73,
	0x12, 0x62, 0x0a, 0x0a, 0x41, 0x64, 0x64, 0x50, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x12, 0x1c,
	0x2e, 0x70, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x70, 0x62, 0x2e, 0x41, 0x64, 0x64, 0x50, 0x72,
	0x6f, 0x64, 0x75, 0x63, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1d, 0x2e, 0x70,
	0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x70, 0x62, 0x2e, 0x41, 0x64, 0x64, 0x50, 0x72, 0x6f, 0x64,
	0x75, 0x63, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x17, 0x82, 0xd3, 0xe4,
	0x93, 0x02, 0x11, 0x3a, 0x01, 0x2a, 0x22, 0x0c, 0x2f, 0x76, 0x31, 0x2f, 0x70, 0x72, 0x6f, 0x64,
	0x75, 0x63, 0x74, 0x73, 0x12, 0x70, 0x0a, 0x0d, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x50, 0x72,
	0x6f, 0x64, 0x75, 0x63, 0x74, 0x12, 0x1f, 0x2e, 0x70, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x70,
	0x62, 0x2e, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x50, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x20, 0x2e, 0x70, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74,
	0x70, 0x62, 0x2e, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x50, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74,
	0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x1c, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x16,
	0x3a, 0x01, 0x2a, 0x1a, 0x11, 0x2f, 0x76, 0x31, 0x2f, 0x70, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74,
	0x73, 0x2f, 0x7b, 0x69, 0x64, 0x7d, 0x12, 0x6d, 0x0a, 0x0d, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65,
	0x50, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x12, 0x1f, 0x2e, 0x70, 0x72, 0x6f, 0x64, 0x75, 0x63,
	0x74, 0x70, 0x62, 0x2e, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x50, 0x72, 0x6f, 0x64, 0x75, 0x63,
	0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x20, 0x2e, 0x70, 0x72, 0x6f, 0x64, 0x75,
	0x63, 0x74, 0x70, 0x62, 0x2e, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x50, 0x72, 0x6f, 0x64, 0x75,
	0x63, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x19, 0x82, 0xd3, 0xe4, 0x93,
	0x02, 0x13, 0x2a, 0x11, 0x2f, 0x76, 0x31, 0x2f, 0x70, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x73,
	0x2f, 0x7b, 0x69, 0x64, 0x7d, 0x12, 0x59, 0x0a, 0x09, 0x41, 0x64, 0x64, 0x54, 0x6f, 0x43, 0x61,
	0x72, 0x74, 0x12, 0x18, 0x2e, 0x63, 0x61, 0x72, 0x74, 0x70, 0x62, 0x2e, 0x41, 0x64, 0x64, 0x54,
	0x6f, 0x43, 0x61, 0x72, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x19, 0x2e, 0x63,
	0x61, 0x72, 0x74, 0x70, 0x62, 0x2e, 0x41, 0x64, 0x64, 0x54, 0x6f, 0x43, 0x61, 0x72, 0x74, 0x52,
	0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x17, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x11, 0x3a,
	0x01, 0x2a, 0x22, 0x0c, 0x2f, 0x76, 0x31, 0x2f, 0x63, 0x61, 0x72, 0x74, 0x2f, 0x61, 0x64, 0x64,
	0x12, 0x6b, 0x0a, 0x0e, 0x52, 0x65, 0x6d, 0x6f, 0x76, 0x65, 0x46, 0x72, 0x6f, 0x6d, 0x43, 0x61,
	0x72, 0x74, 0x12, 0x1d, 0x2e, 0x63, 0x61, 0x72, 0x74, 0x70, 0x62, 0x2e, 0x52, 0x65, 0x6d, 0x6f,
	0x76, 0x65, 0x46, 0x72, 0x6f, 0x6d, 0x43, 0x61, 0x72, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x1a, 0x1e, 0x2e, 0x63, 0x61, 0x72, 0x74, 0x70, 0x62, 0x2e, 0x52, 0x65, 0x6d, 0x6f, 0x76,
	0x65, 0x46, 0x72, 0x6f, 0x6d, 0x43, 0x61, 0x72, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73,
	0x65, 0x22, 0x1a, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x14, 0x3a, 0x01, 0x2a, 0x22, 0x0f, 0x2f, 0x76,
	0x31, 0x2f, 0x63, 0x61, 0x72, 0x74, 0x2f, 0x72, 0x65, 0x6d, 0x6f, 0x76, 0x65, 0x12, 0x56, 0x0a,
	0x07, 0x47, 0x65, 0x74, 0x43, 0x61, 0x72, 0x74, 0x12, 0x16, 0x2e, 0x63, 0x61, 0x72, 0x74, 0x70,
	0x62, 0x2e, 0x47, 0x65, 0x74, 0x43, 0x61, 0x72, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x1a, 0x17, 0x2e, 0x63, 0x61, 0x72, 0x74, 0x70, 0x62, 0x2e, 0x47, 0x65, 0x74, 0x43, 0x61, 0x72,
	0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x1a, 0x82, 0xd3, 0xe4, 0x93, 0x02,
	0x14, 0x12, 0x12, 0x2f, 0x76, 0x31, 0x2f, 0x63, 0x61, 0x72, 0x74, 0x2f, 0x7b, 0x75, 0x73, 0x65,
	0x72, 0x5f, 0x69, 0x64, 0x7d, 0x12, 0x5b, 0x0a, 0x09, 0x43, 0x6c, 0x65, 0x61, 0x72, 0x43, 0x61,
	0x72, 0x74, 0x12, 0x18, 0x2e, 0x63, 0x61, 0x72, 0x74, 0x70, 0x62, 0x2e, 0x43, 0x6c, 0x65, 0x61,
	0x72, 0x43, 0x61, 0x72, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x19, 0x2e, 0x63,
	0x61, 0x72, 0x74, 0x70, 0x62, 0x2e, 0x43, 0x6c, 0x65, 0x61, 0x72, 0x43, 0x61, 0x72, 0x74, 0x52,
	0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x19, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x13, 0x3a,
	0x01, 0x2a, 0x22, 0x0e, 0x2f, 0x76, 0x31, 0x2f, 0x63, 0x61, 0x72, 0x74, 0x2f, 0x63, 0x6c, 0x65,
	0x61, 0x72, 0x12, 0x5b, 0x0a, 0x08, 0x43, 0x68, 0x65, 0x63, 0x6b, 0x6f, 0x75, 0x74, 0x12, 0x17,
	0x2e, 0x63, 0x61, 0x72, 0x74, 0x70, 0x62, 0x2e, 0x43, 0x68, 0x65, 0x63, 0x6b, 0x6f, 0x75, 0x74,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x18, 0x2e, 0x63, 0x61, 0x72, 0x74, 0x70, 0x62,
	0x2e, 0x43, 0x68, 0x65, 0x63, 0x6b, 0x6f, 0x75, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73,
	0x65, 0x22, 0x1c, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x16, 0x3a, 0x01, 0x2a, 0x22, 0x11, 0x2f, 0x76,
	0x31, 0x2f, 0x63, 0x61, 0x72, 0x74, 0x2f, 0x63, 0x68, 0x65, 0x63, 0x6b, 0x6f, 0x75, 0x74, 0x42,
	0x40, 0x5a, 0x3e, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x77, 0x61,
	0x7a, 0x77, 0x6b, 0x69, 0x2f, 0x57, 0x65, 0x61, 0x72, 0x53, 0x74, 0x6f, 0x72, 0x65, 0x2f, 0x61,
	0x70, 0x69, 0x2d, 0x67, 0x61, 0x74, 0x65, 0x77, 0x61, 0x79, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x61, 0x70, 0x69, 0x67, 0x61, 0x74, 0x65, 0x77, 0x61, 0x79, 0x70,
	0x62, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var file_github_com_wazwki_WearStore_api_gateway_api_proto_api_gateway_proto_goTypes = []any{
	(*userpb.RegisterUserRequest)(nil),      // 0: userpb.RegisterUserRequest
	(*userpb.LoginUserRequest)(nil),         // 1: userpb.LoginUserRequest
	(*userpb.GetUserRequest)(nil),           // 2: userpb.GetUserRequest
	(*userpb.UpdateUserRequest)(nil),        // 3: userpb.UpdateUserRequest
	(*userpb.DeleteUserRequest)(nil),        // 4: userpb.DeleteUserRequest
	(*productpb.GetProductRequest)(nil),     // 5: productpb.GetProductRequest
	(*productpb.ListProductsRequest)(nil),   // 6: productpb.ListProductsRequest
	(*productpb.AddProductRequest)(nil),     // 7: productpb.AddProductRequest
	(*productpb.UpdateProductRequest)(nil),  // 8: productpb.UpdateProductRequest
	(*productpb.DeleteProductRequest)(nil),  // 9: productpb.DeleteProductRequest
	(*cartpb.AddToCartRequest)(nil),         // 10: cartpb.AddToCartRequest
	(*cartpb.RemoveFromCartRequest)(nil),    // 11: cartpb.RemoveFromCartRequest
	(*cartpb.GetCartRequest)(nil),           // 12: cartpb.GetCartRequest
	(*cartpb.ClearCartRequest)(nil),         // 13: cartpb.ClearCartRequest
	(*cartpb.CheckoutRequest)(nil),          // 14: cartpb.CheckoutRequest
	(*userpb.RegisterUserResponse)(nil),     // 15: userpb.RegisterUserResponse
	(*userpb.LoginUserResponse)(nil),        // 16: userpb.LoginUserResponse
	(*userpb.GetUserResponse)(nil),          // 17: userpb.GetUserResponse
	(*userpb.UpdateUserResponse)(nil),       // 18: userpb.UpdateUserResponse
	(*userpb.DeleteUserResponse)(nil),       // 19: userpb.DeleteUserResponse
	(*productpb.GetProductResponse)(nil),    // 20: productpb.GetProductResponse
	(*productpb.ListProductsResponse)(nil),  // 21: productpb.ListProductsResponse
	(*productpb.AddProductResponse)(nil),    // 22: productpb.AddProductResponse
	(*productpb.UpdateProductResponse)(nil), // 23: productpb.UpdateProductResponse
	(*productpb.DeleteProductResponse)(nil), // 24: productpb.DeleteProductResponse
	(*cartpb.AddToCartResponse)(nil),        // 25: cartpb.AddToCartResponse
	(*cartpb.RemoveFromCartResponse)(nil),   // 26: cartpb.RemoveFromCartResponse
	(*cartpb.GetCartResponse)(nil),          // 27: cartpb.GetCartResponse
	(*cartpb.ClearCartResponse)(nil),        // 28: cartpb.ClearCartResponse
	(*cartpb.CheckoutResponse)(nil),         // 29: cartpb.CheckoutResponse
}
var file_github_com_wazwki_WearStore_api_gateway_api_proto_api_gateway_proto_depIdxs = []int32{
	0,  // 0: apigatewaypb.ApiGatewayService.RegisterUser:input_type -> userpb.RegisterUserRequest
	1,  // 1: apigatewaypb.ApiGatewayService.LoginUser:input_type -> userpb.LoginUserRequest
	2,  // 2: apigatewaypb.ApiGatewayService.GetUser:input_type -> userpb.GetUserRequest
	3,  // 3: apigatewaypb.ApiGatewayService.UpdateUser:input_type -> userpb.UpdateUserRequest
	4,  // 4: apigatewaypb.ApiGatewayService.DeleteUser:input_type -> userpb.DeleteUserRequest
	5,  // 5: apigatewaypb.ApiGatewayService.GetProduct:input_type -> productpb.GetProductRequest
	6,  // 6: apigatewaypb.ApiGatewayService.ListProducts:input_type -> productpb.ListProductsRequest
	7,  // 7: apigatewaypb.ApiGatewayService.AddProduct:input_type -> productpb.AddProductRequest
	8,  // 8: apigatewaypb.ApiGatewayService.UpdateProduct:input_type -> productpb.UpdateProductRequest
	9,  // 9: apigatewaypb.ApiGatewayService.DeleteProduct:input_type -> productpb.DeleteProductRequest
	10, // 10: apigatewaypb.ApiGatewayService.AddToCart:input_type -> cartpb.AddToCartRequest
	11, // 11: apigatewaypb.ApiGatewayService.RemoveFromCart:input_type -> cartpb.RemoveFromCartRequest
	12, // 12: apigatewaypb.ApiGatewayService.GetCart:input_type -> cartpb.GetCartRequest
	13, // 13: apigatewaypb.ApiGatewayService.ClearCart:input_type -> cartpb.ClearCartRequest
	14, // 14: apigatewaypb.ApiGatewayService.Checkout:input_type -> cartpb.CheckoutRequest
	15, // 15: apigatewaypb.ApiGatewayService.RegisterUser:output_type -> userpb.RegisterUserResponse
	16, // 16: apigatewaypb.ApiGatewayService.LoginUser:output_type -> userpb.LoginUserResponse
	17, // 17: apigatewaypb.ApiGatewayService.GetUser:output_type -> userpb.GetUserResponse
	18, // 18: apigatewaypb.ApiGatewayService.UpdateUser:output_type -> userpb.UpdateUserResponse
	19, // 19: apigatewaypb.ApiGatewayService.DeleteUser:output_type -> userpb.DeleteUserResponse
	20, // 20: apigatewaypb.ApiGatewayService.GetProduct:output_type -> productpb.GetProductResponse
	21, // 21: apigatewaypb.ApiGatewayService.ListProducts:output_type -> productpb.ListProductsResponse
	22, // 22: apigatewaypb.ApiGatewayService.AddProduct:output_type -> productpb.AddProductResponse
	23, // 23: apigatewaypb.ApiGatewayService.UpdateProduct:output_type -> productpb.UpdateProductResponse
	24, // 24: apigatewaypb.ApiGatewayService.DeleteProduct:output_type -> productpb.DeleteProductResponse
	25, // 25: apigatewaypb.ApiGatewayService.AddToCart:output_type -> cartpb.AddToCartResponse
	26, // 26: apigatewaypb.ApiGatewayService.RemoveFromCart:output_type -> cartpb.RemoveFromCartResponse
	27, // 27: apigatewaypb.ApiGatewayService.GetCart:output_type -> cartpb.GetCartResponse
	28, // 28: apigatewaypb.ApiGatewayService.ClearCart:output_type -> cartpb.ClearCartResponse
	29, // 29: apigatewaypb.ApiGatewayService.Checkout:output_type -> cartpb.CheckoutResponse
	15, // [15:30] is the sub-list for method output_type
	0,  // [0:15] is the sub-list for method input_type
	0,  // [0:0] is the sub-list for extension type_name
	0,  // [0:0] is the sub-list for extension extendee
	0,  // [0:0] is the sub-list for field type_name
}

func init() { file_github_com_wazwki_WearStore_api_gateway_api_proto_api_gateway_proto_init() }
func file_github_com_wazwki_WearStore_api_gateway_api_proto_api_gateway_proto_init() {
	if File_github_com_wazwki_WearStore_api_gateway_api_proto_api_gateway_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_github_com_wazwki_WearStore_api_gateway_api_proto_api_gateway_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   0,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_github_com_wazwki_WearStore_api_gateway_api_proto_api_gateway_proto_goTypes,
		DependencyIndexes: file_github_com_wazwki_WearStore_api_gateway_api_proto_api_gateway_proto_depIdxs,
	}.Build()
	File_github_com_wazwki_WearStore_api_gateway_api_proto_api_gateway_proto = out.File
	file_github_com_wazwki_WearStore_api_gateway_api_proto_api_gateway_proto_rawDesc = nil
	file_github_com_wazwki_WearStore_api_gateway_api_proto_api_gateway_proto_goTypes = nil
	file_github_com_wazwki_WearStore_api_gateway_api_proto_api_gateway_proto_depIdxs = nil
}
