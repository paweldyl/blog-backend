// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.36.6
// 	protoc        v3.21.12
// source: post/rpc_delete_post.proto

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

type DeletePostRequest struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Id            string                 `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *DeletePostRequest) Reset() {
	*x = DeletePostRequest{}
	mi := &file_post_rpc_delete_post_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *DeletePostRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DeletePostRequest) ProtoMessage() {}

func (x *DeletePostRequest) ProtoReflect() protoreflect.Message {
	mi := &file_post_rpc_delete_post_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DeletePostRequest.ProtoReflect.Descriptor instead.
func (*DeletePostRequest) Descriptor() ([]byte, []int) {
	return file_post_rpc_delete_post_proto_rawDescGZIP(), []int{0}
}

func (x *DeletePostRequest) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

var File_post_rpc_delete_post_proto protoreflect.FileDescriptor

const file_post_rpc_delete_post_proto_rawDesc = "" +
	"\n" +
	"\x1apost/rpc_delete_post.proto\x12\x02pb\"#\n" +
	"\x11DeletePostRequest\x12\x0e\n" +
	"\x02id\x18\x01 \x01(\tR\x02idB%Z#github.com/paweldyl/blog-backend/pbb\x06proto3"

var (
	file_post_rpc_delete_post_proto_rawDescOnce sync.Once
	file_post_rpc_delete_post_proto_rawDescData []byte
)

func file_post_rpc_delete_post_proto_rawDescGZIP() []byte {
	file_post_rpc_delete_post_proto_rawDescOnce.Do(func() {
		file_post_rpc_delete_post_proto_rawDescData = protoimpl.X.CompressGZIP(unsafe.Slice(unsafe.StringData(file_post_rpc_delete_post_proto_rawDesc), len(file_post_rpc_delete_post_proto_rawDesc)))
	})
	return file_post_rpc_delete_post_proto_rawDescData
}

var file_post_rpc_delete_post_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_post_rpc_delete_post_proto_goTypes = []any{
	(*DeletePostRequest)(nil), // 0: pb.DeletePostRequest
}
var file_post_rpc_delete_post_proto_depIdxs = []int32{
	0, // [0:0] is the sub-list for method output_type
	0, // [0:0] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_post_rpc_delete_post_proto_init() }
func file_post_rpc_delete_post_proto_init() {
	if File_post_rpc_delete_post_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: unsafe.Slice(unsafe.StringData(file_post_rpc_delete_post_proto_rawDesc), len(file_post_rpc_delete_post_proto_rawDesc)),
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_post_rpc_delete_post_proto_goTypes,
		DependencyIndexes: file_post_rpc_delete_post_proto_depIdxs,
		MessageInfos:      file_post_rpc_delete_post_proto_msgTypes,
	}.Build()
	File_post_rpc_delete_post_proto = out.File
	file_post_rpc_delete_post_proto_goTypes = nil
	file_post_rpc_delete_post_proto_depIdxs = nil
}
