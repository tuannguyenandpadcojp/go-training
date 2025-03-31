// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.36.6
// 	protoc        (unknown)
// source: tenant/v1/tenant.proto

package tenantv1

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	reflect "reflect"
	sync "sync"
	unsafe "unsafe"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type Tenant struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Id            string                 `protobuf:"bytes,1,opt,Name=id,proto3" json:"id,omitempty"`
	Name          string                 `protobuf:"bytes,2,opt,Name=Name,proto3" json:"Name,omitempty"`
	Email         string                 `protobuf:"bytes,3,opt,Name=email,proto3" json:"email,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *Tenant) Reset() {
	*x = Tenant{}
	mi := &file_tenant_v1_tenant_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *Tenant) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Tenant) ProtoMessage() {}

func (x *Tenant) ProtoReflect() protoreflect.Message {
	mi := &file_tenant_v1_tenant_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Tenant.ProtoReflect.Descriptor instead.
func (*Tenant) Descriptor() ([]byte, []int) {
	return file_tenant_v1_tenant_proto_rawDescGZIP(), []int{0}
}

func (x *Tenant) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *Tenant) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *Tenant) GetEmail() string {
	if x != nil {
		return x.Email
	}
	return ""
}

type GetTenantResponse struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Tenant        *Tenant                `protobuf:"bytes,1,opt,Name=tenant,proto3" json:"tenant,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *GetTenantResponse) Reset() {
	*x = GetTenantResponse{}
	mi := &file_tenant_v1_tenant_proto_msgTypes[1]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *GetTenantResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetTenantResponse) ProtoMessage() {}

func (x *GetTenantResponse) ProtoReflect() protoreflect.Message {
	mi := &file_tenant_v1_tenant_proto_msgTypes[1]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetTenantResponse.ProtoReflect.Descriptor instead.
func (*GetTenantResponse) Descriptor() ([]byte, []int) {
	return file_tenant_v1_tenant_proto_rawDescGZIP(), []int{1}
}

func (x *GetTenantResponse) GetTenant() *Tenant {
	if x != nil {
		return x.Tenant
	}
	return nil
}

type GetTenantRequest struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Name          string                 `protobuf:"bytes,1,opt,Name=Name,proto3" json:"Name,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *GetTenantRequest) Reset() {
	*x = GetTenantRequest{}
	mi := &file_tenant_v1_tenant_proto_msgTypes[2]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *GetTenantRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetTenantRequest) ProtoMessage() {}

func (x *GetTenantRequest) ProtoReflect() protoreflect.Message {
	mi := &file_tenant_v1_tenant_proto_msgTypes[2]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetTenantRequest.ProtoReflect.Descriptor instead.
func (*GetTenantRequest) Descriptor() ([]byte, []int) {
	return file_tenant_v1_tenant_proto_rawDescGZIP(), []int{2}
}

func (x *GetTenantRequest) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

var File_tenant_v1_tenant_proto protoreflect.FileDescriptor

const file_tenant_v1_tenant_proto_rawDesc = "" +
	"\n" +
	"\x16tenant/v1/tenant.proto\x12\ttenant.v1\"B\n" +
	"\x06Tenant\x12\x0e\n" +
	"\x02id\x18\x01 \x01(\tR\x02id\x12\x12\n" +
	"\x04name\x18\x02 \x01(\tR\x04name\x12\x14\n" +
	"\x05email\x18\x03 \x01(\tR\x05email\">\n" +
	"\x11GetTenantResponse\x12)\n" +
	"\x06tenant\x18\x01 \x01(\v2\x11.tenant.v1.TenantR\x06tenant\"&\n" +
	"\x10GetTenantRequest\x12\x12\n" +
	"\x04name\x18\x01 \x01(\tR\x04name2Y\n" +
	"\rTenantService\x12H\n" +
	"\tGetTenant\x12\x1b.tenant.v1.GetTenantRequest\x1a\x1c.tenant.v1.GetTenantResponse\"\x00BLZJgithub.com/tuannguyenandpadcojp/go-training/lqm/week7/day1/gen/go;tenantv1b\x06proto3"

var (
	file_tenant_v1_tenant_proto_rawDescOnce sync.Once
	file_tenant_v1_tenant_proto_rawDescData []byte
)

func file_tenant_v1_tenant_proto_rawDescGZIP() []byte {
	file_tenant_v1_tenant_proto_rawDescOnce.Do(func() {
		file_tenant_v1_tenant_proto_rawDescData = protoimpl.X.CompressGZIP(unsafe.Slice(unsafe.StringData(file_tenant_v1_tenant_proto_rawDesc), len(file_tenant_v1_tenant_proto_rawDesc)))
	})
	return file_tenant_v1_tenant_proto_rawDescData
}

var file_tenant_v1_tenant_proto_msgTypes = make([]protoimpl.MessageInfo, 3)
var file_tenant_v1_tenant_proto_goTypes = []any{
	(*Tenant)(nil),            // 0: tenant.v1.Tenant
	(*GetTenantResponse)(nil), // 1: tenant.v1.GetTenantResponse
	(*GetTenantRequest)(nil),  // 2: tenant.v1.GetTenantRequest
}
var file_tenant_v1_tenant_proto_depIdxs = []int32{
	0, // 0: tenant.v1.GetTenantResponse.tenant:type_name -> tenant.v1.Tenant
	2, // 1: tenant.v1.TenantService.GetTenant:input_type -> tenant.v1.GetTenantRequest
	1, // 2: tenant.v1.TenantService.GetTenant:output_type -> tenant.v1.GetTenantResponse
	2, // [2:3] is the sub-list for method output_type
	1, // [1:2] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_tenant_v1_tenant_proto_init() }
func file_tenant_v1_tenant_proto_init() {
	if File_tenant_v1_tenant_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: unsafe.Slice(unsafe.StringData(file_tenant_v1_tenant_proto_rawDesc), len(file_tenant_v1_tenant_proto_rawDesc)),
			NumEnums:      0,
			NumMessages:   3,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_tenant_v1_tenant_proto_goTypes,
		DependencyIndexes: file_tenant_v1_tenant_proto_depIdxs,
		MessageInfos:      file_tenant_v1_tenant_proto_msgTypes,
	}.Build()
	File_tenant_v1_tenant_proto = out.File
	file_tenant_v1_tenant_proto_goTypes = nil
	file_tenant_v1_tenant_proto_depIdxs = nil
}
