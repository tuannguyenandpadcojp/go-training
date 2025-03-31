// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.33.0
// 	protoc        (unknown)
// source: v1/client_service.proto

package v1

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

type UserAttributes struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// User ID of the user making the request
	UserId string `protobuf:"bytes,1,opt,name=user_id,json=userId,proto3" json:"user_id,omitempty"`
	// Client ID of the client that the user is belonging to (main client)
	ClientId string `protobuf:"bytes,2,opt,name=client_id,json=clientId,proto3" json:"client_id,omitempty"`
	// Client ID of the client being requested
	CurrentClientId string `protobuf:"bytes,3,opt,name=current_client_id,json=currentClientId,proto3" json:"current_client_id,omitempty"`
}

func (x *UserAttributes) Reset() {
	*x = UserAttributes{}
	if protoimpl.UnsafeEnabled {
		mi := &file_v1_client_service_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UserAttributes) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UserAttributes) ProtoMessage() {}

func (x *UserAttributes) ProtoReflect() protoreflect.Message {
	mi := &file_v1_client_service_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UserAttributes.ProtoReflect.Descriptor instead.
func (*UserAttributes) Descriptor() ([]byte, []int) {
	return file_v1_client_service_proto_rawDescGZIP(), []int{0}
}

func (x *UserAttributes) GetUserId() string {
	if x != nil {
		return x.UserId
	}
	return ""
}

func (x *UserAttributes) GetClientId() string {
	if x != nil {
		return x.ClientId
	}
	return ""
}

func (x *UserAttributes) GetCurrentClientId() string {
	if x != nil {
		return x.CurrentClientId
	}
	return ""
}

type GetClientRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	UserAttributes *UserAttributes `protobuf:"bytes,1,opt,name=user_attributes,json=userAttributes,proto3" json:"user_attributes,omitempty"`
}

func (x *GetClientRequest) Reset() {
	*x = GetClientRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_v1_client_service_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetClientRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetClientRequest) ProtoMessage() {}

func (x *GetClientRequest) ProtoReflect() protoreflect.Message {
	mi := &file_v1_client_service_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetClientRequest.ProtoReflect.Descriptor instead.
func (*GetClientRequest) Descriptor() ([]byte, []int) {
	return file_v1_client_service_proto_rawDescGZIP(), []int{1}
}

func (x *GetClientRequest) GetUserAttributes() *UserAttributes {
	if x != nil {
		return x.UserAttributes
	}
	return nil
}

type GetClientResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Client *Client `protobuf:"bytes,1,opt,name=client,proto3" json:"client,omitempty"`
}

func (x *GetClientResponse) Reset() {
	*x = GetClientResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_v1_client_service_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetClientResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetClientResponse) ProtoMessage() {}

func (x *GetClientResponse) ProtoReflect() protoreflect.Message {
	mi := &file_v1_client_service_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetClientResponse.ProtoReflect.Descriptor instead.
func (*GetClientResponse) Descriptor() ([]byte, []int) {
	return file_v1_client_service_proto_rawDescGZIP(), []int{2}
}

func (x *GetClientResponse) GetClient() *Client {
	if x != nil {
		return x.Client
	}
	return nil
}

var File_v1_client_service_proto protoreflect.FileDescriptor

var file_v1_client_service_proto_rawDesc = []byte{
	0x0a, 0x17, 0x76, 0x31, 0x2f, 0x63, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x5f, 0x73, 0x65, 0x72, 0x76,
	0x69, 0x63, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x05, 0x74, 0x6d, 0x2e, 0x76, 0x31,
	0x1a, 0x0f, 0x76, 0x31, 0x2f, 0x63, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x22, 0x72, 0x0a, 0x0e, 0x55, 0x73, 0x65, 0x72, 0x41, 0x74, 0x74, 0x72, 0x69, 0x62, 0x75,
	0x74, 0x65, 0x73, 0x12, 0x17, 0x0a, 0x07, 0x75, 0x73, 0x65, 0x72, 0x5f, 0x69, 0x64, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x75, 0x73, 0x65, 0x72, 0x49, 0x64, 0x12, 0x1b, 0x0a, 0x09,
	0x63, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x5f, 0x69, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x08, 0x63, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x49, 0x64, 0x12, 0x2a, 0x0a, 0x11, 0x63, 0x75, 0x72,
	0x72, 0x65, 0x6e, 0x74, 0x5f, 0x63, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x5f, 0x69, 0x64, 0x18, 0x03,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x0f, 0x63, 0x75, 0x72, 0x72, 0x65, 0x6e, 0x74, 0x43, 0x6c, 0x69,
	0x65, 0x6e, 0x74, 0x49, 0x64, 0x22, 0x52, 0x0a, 0x10, 0x47, 0x65, 0x74, 0x43, 0x6c, 0x69, 0x65,
	0x6e, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x3e, 0x0a, 0x0f, 0x75, 0x73, 0x65,
	0x72, 0x5f, 0x61, 0x74, 0x74, 0x72, 0x69, 0x62, 0x75, 0x74, 0x65, 0x73, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x0b, 0x32, 0x15, 0x2e, 0x74, 0x6d, 0x2e, 0x76, 0x31, 0x2e, 0x55, 0x73, 0x65, 0x72, 0x41,
	0x74, 0x74, 0x72, 0x69, 0x62, 0x75, 0x74, 0x65, 0x73, 0x52, 0x0e, 0x75, 0x73, 0x65, 0x72, 0x41,
	0x74, 0x74, 0x72, 0x69, 0x62, 0x75, 0x74, 0x65, 0x73, 0x22, 0x3a, 0x0a, 0x11, 0x47, 0x65, 0x74,
	0x43, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x25,
	0x0a, 0x06, 0x63, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0d,
	0x2e, 0x74, 0x6d, 0x2e, 0x76, 0x31, 0x2e, 0x43, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x52, 0x06, 0x63,
	0x6c, 0x69, 0x65, 0x6e, 0x74, 0x32, 0x4f, 0x0a, 0x0d, 0x43, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x53,
	0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x3e, 0x0a, 0x09, 0x47, 0x65, 0x74, 0x43, 0x6c, 0x69,
	0x65, 0x6e, 0x74, 0x12, 0x17, 0x2e, 0x74, 0x6d, 0x2e, 0x76, 0x31, 0x2e, 0x47, 0x65, 0x74, 0x43,
	0x6c, 0x69, 0x65, 0x6e, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x18, 0x2e, 0x74,
	0x6d, 0x2e, 0x76, 0x31, 0x2e, 0x47, 0x65, 0x74, 0x43, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x52, 0x65,
	0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x42, 0x4a, 0x5a, 0x48, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62,
	0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x74, 0x75, 0x61, 0x6e, 0x6e, 0x67, 0x75, 0x79, 0x65, 0x6e, 0x61,
	0x6e, 0x64, 0x70, 0x61, 0x64, 0x63, 0x6f, 0x6a, 0x70, 0x2f, 0x67, 0x6f, 0x2d, 0x74, 0x72, 0x61,
	0x69, 0x6e, 0x69, 0x6e, 0x67, 0x2f, 0x77, 0x65, 0x65, 0x6b, 0x37, 0x2f, 0x64, 0x61, 0x79, 0x31,
	0x2f, 0x69, 0x6e, 0x74, 0x65, 0x72, 0x6e, 0x61, 0x6c, 0x2f, 0x70, 0x62, 0x2f, 0x76, 0x31, 0x3b,
	0x76, 0x31, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_v1_client_service_proto_rawDescOnce sync.Once
	file_v1_client_service_proto_rawDescData = file_v1_client_service_proto_rawDesc
)

func file_v1_client_service_proto_rawDescGZIP() []byte {
	file_v1_client_service_proto_rawDescOnce.Do(func() {
		file_v1_client_service_proto_rawDescData = protoimpl.X.CompressGZIP(file_v1_client_service_proto_rawDescData)
	})
	return file_v1_client_service_proto_rawDescData
}

var file_v1_client_service_proto_msgTypes = make([]protoimpl.MessageInfo, 3)
var file_v1_client_service_proto_goTypes = []interface{}{
	(*UserAttributes)(nil),    // 0: tm.v1.UserAttributes
	(*GetClientRequest)(nil),  // 1: tm.v1.GetClientRequest
	(*GetClientResponse)(nil), // 2: tm.v1.GetClientResponse
	(*Client)(nil),            // 3: tm.v1.Client
}
var file_v1_client_service_proto_depIdxs = []int32{
	0, // 0: tm.v1.GetClientRequest.user_attributes:type_name -> tm.v1.UserAttributes
	3, // 1: tm.v1.GetClientResponse.client:type_name -> tm.v1.Client
	1, // 2: tm.v1.ClientService.GetClient:input_type -> tm.v1.GetClientRequest
	2, // 3: tm.v1.ClientService.GetClient:output_type -> tm.v1.GetClientResponse
	3, // [3:4] is the sub-list for method output_type
	2, // [2:3] is the sub-list for method input_type
	2, // [2:2] is the sub-list for extension type_name
	2, // [2:2] is the sub-list for extension extendee
	0, // [0:2] is the sub-list for field type_name
}

func init() { file_v1_client_service_proto_init() }
func file_v1_client_service_proto_init() {
	if File_v1_client_service_proto != nil {
		return
	}
	file_v1_client_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_v1_client_service_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UserAttributes); i {
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
		file_v1_client_service_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetClientRequest); i {
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
		file_v1_client_service_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetClientResponse); i {
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
			RawDescriptor: file_v1_client_service_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   3,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_v1_client_service_proto_goTypes,
		DependencyIndexes: file_v1_client_service_proto_depIdxs,
		MessageInfos:      file_v1_client_service_proto_msgTypes,
	}.Build()
	File_v1_client_service_proto = out.File
	file_v1_client_service_proto_rawDesc = nil
	file_v1_client_service_proto_goTypes = nil
	file_v1_client_service_proto_depIdxs = nil
}
