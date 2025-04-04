// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.36.6
// 	protoc        v3.21.12
// source: user/rpc_update_user.proto

package pb

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

type UpdateUserRequest struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Username      *string                `protobuf:"bytes,1,opt,name=username,proto3,oneof" json:"username,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *UpdateUserRequest) Reset() {
	*x = UpdateUserRequest{}
	mi := &file_user_rpc_update_user_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *UpdateUserRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UpdateUserRequest) ProtoMessage() {}

func (x *UpdateUserRequest) ProtoReflect() protoreflect.Message {
	mi := &file_user_rpc_update_user_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UpdateUserRequest.ProtoReflect.Descriptor instead.
func (*UpdateUserRequest) Descriptor() ([]byte, []int) {
	return file_user_rpc_update_user_proto_rawDescGZIP(), []int{0}
}

func (x *UpdateUserRequest) GetUsername() string {
	if x != nil && x.Username != nil {
		return *x.Username
	}
	return ""
}

type UpdateUserResponse struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	User          *User                  `protobuf:"bytes,1,opt,name=user,proto3" json:"user,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *UpdateUserResponse) Reset() {
	*x = UpdateUserResponse{}
	mi := &file_user_rpc_update_user_proto_msgTypes[1]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *UpdateUserResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UpdateUserResponse) ProtoMessage() {}

func (x *UpdateUserResponse) ProtoReflect() protoreflect.Message {
	mi := &file_user_rpc_update_user_proto_msgTypes[1]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UpdateUserResponse.ProtoReflect.Descriptor instead.
func (*UpdateUserResponse) Descriptor() ([]byte, []int) {
	return file_user_rpc_update_user_proto_rawDescGZIP(), []int{1}
}

func (x *UpdateUserResponse) GetUser() *User {
	if x != nil {
		return x.User
	}
	return nil
}

var File_user_rpc_update_user_proto protoreflect.FileDescriptor

const file_user_rpc_update_user_proto_rawDesc = "" +
	"\n" +
	"\x1auser/rpc_update_user.proto\x12\x02pb\x1a\x0fuser/user.proto\"A\n" +
	"\x11UpdateUserRequest\x12\x1f\n" +
	"\busername\x18\x01 \x01(\tH\x00R\busername\x88\x01\x01B\v\n" +
	"\t_username\"2\n" +
	"\x12UpdateUserResponse\x12\x1c\n" +
	"\x04user\x18\x01 \x01(\v2\b.pb.UserR\x04userB%Z#github.com/paweldyl/blog-backend/pbb\x06proto3"

var (
	file_user_rpc_update_user_proto_rawDescOnce sync.Once
	file_user_rpc_update_user_proto_rawDescData []byte
)

func file_user_rpc_update_user_proto_rawDescGZIP() []byte {
	file_user_rpc_update_user_proto_rawDescOnce.Do(func() {
		file_user_rpc_update_user_proto_rawDescData = protoimpl.X.CompressGZIP(unsafe.Slice(unsafe.StringData(file_user_rpc_update_user_proto_rawDesc), len(file_user_rpc_update_user_proto_rawDesc)))
	})
	return file_user_rpc_update_user_proto_rawDescData
}

var file_user_rpc_update_user_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_user_rpc_update_user_proto_goTypes = []any{
	(*UpdateUserRequest)(nil),  // 0: pb.UpdateUserRequest
	(*UpdateUserResponse)(nil), // 1: pb.UpdateUserResponse
	(*User)(nil),               // 2: pb.User
}
var file_user_rpc_update_user_proto_depIdxs = []int32{
	2, // 0: pb.UpdateUserResponse.user:type_name -> pb.User
	1, // [1:1] is the sub-list for method output_type
	1, // [1:1] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_user_rpc_update_user_proto_init() }
func file_user_rpc_update_user_proto_init() {
	if File_user_rpc_update_user_proto != nil {
		return
	}
	file_user_user_proto_init()
	file_user_rpc_update_user_proto_msgTypes[0].OneofWrappers = []any{}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: unsafe.Slice(unsafe.StringData(file_user_rpc_update_user_proto_rawDesc), len(file_user_rpc_update_user_proto_rawDesc)),
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_user_rpc_update_user_proto_goTypes,
		DependencyIndexes: file_user_rpc_update_user_proto_depIdxs,
		MessageInfos:      file_user_rpc_update_user_proto_msgTypes,
	}.Build()
	File_user_rpc_update_user_proto = out.File
	file_user_rpc_update_user_proto_goTypes = nil
	file_user_rpc_update_user_proto_depIdxs = nil
}
