// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.36.6
// 	protoc        v3.21.12
// source: service_blog.proto

package pb

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	reflect "reflect"
	unsafe "unsafe"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

var File_service_blog_proto protoreflect.FileDescriptor

const file_service_blog_proto_rawDesc = "" +
	"\n" +
	"\x12service_blog.proto\x12\x02pb\x1a\x15rpc_create_user.proto\x1a\x12rpc_get_user.proto\x1a\x15rpc_update_user.proto\x1a\x14rpc_login_user.proto2\xf6\x01\n" +
	"\x04Blog\x12=\n" +
	"\n" +
	"CreateUser\x12\x15.pb.CreateUserRequest\x1a\x16.pb.CreateUserResponse\"\x00\x124\n" +
	"\aGetUser\x12\x12.pb.GetUserRequest\x1a\x13.pb.GetUserResponse\"\x00\x12=\n" +
	"\n" +
	"UpdateUser\x12\x15.pb.UpdateUserRequest\x1a\x16.pb.UpdateUserResponse\"\x00\x12:\n" +
	"\tLoginUser\x12\x14.pb.LoginUserRequest\x1a\x15.pb.LoginUserResponse\"\x00B%Z#github.com/paweldyl/blog-backend/pbb\x06proto3"

var file_service_blog_proto_goTypes = []any{
	(*CreateUserRequest)(nil),  // 0: pb.CreateUserRequest
	(*GetUserRequest)(nil),     // 1: pb.GetUserRequest
	(*UpdateUserRequest)(nil),  // 2: pb.UpdateUserRequest
	(*LoginUserRequest)(nil),   // 3: pb.LoginUserRequest
	(*CreateUserResponse)(nil), // 4: pb.CreateUserResponse
	(*GetUserResponse)(nil),    // 5: pb.GetUserResponse
	(*UpdateUserResponse)(nil), // 6: pb.UpdateUserResponse
	(*LoginUserResponse)(nil),  // 7: pb.LoginUserResponse
}
var file_service_blog_proto_depIdxs = []int32{
	0, // 0: pb.Blog.CreateUser:input_type -> pb.CreateUserRequest
	1, // 1: pb.Blog.GetUser:input_type -> pb.GetUserRequest
	2, // 2: pb.Blog.UpdateUser:input_type -> pb.UpdateUserRequest
	3, // 3: pb.Blog.LoginUser:input_type -> pb.LoginUserRequest
	4, // 4: pb.Blog.CreateUser:output_type -> pb.CreateUserResponse
	5, // 5: pb.Blog.GetUser:output_type -> pb.GetUserResponse
	6, // 6: pb.Blog.UpdateUser:output_type -> pb.UpdateUserResponse
	7, // 7: pb.Blog.LoginUser:output_type -> pb.LoginUserResponse
	4, // [4:8] is the sub-list for method output_type
	0, // [0:4] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_service_blog_proto_init() }
func file_service_blog_proto_init() {
	if File_service_blog_proto != nil {
		return
	}
	file_rpc_create_user_proto_init()
	file_rpc_get_user_proto_init()
	file_rpc_update_user_proto_init()
	file_rpc_login_user_proto_init()
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: unsafe.Slice(unsafe.StringData(file_service_blog_proto_rawDesc), len(file_service_blog_proto_rawDesc)),
			NumEnums:      0,
			NumMessages:   0,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_service_blog_proto_goTypes,
		DependencyIndexes: file_service_blog_proto_depIdxs,
	}.Build()
	File_service_blog_proto = out.File
	file_service_blog_proto_goTypes = nil
	file_service_blog_proto_depIdxs = nil
}
