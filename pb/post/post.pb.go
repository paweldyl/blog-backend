// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.36.6
// 	protoc        v3.21.12
// source: post/post.proto

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

type Post struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Id            string                 `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Title         string                 `protobuf:"bytes,2,opt,name=title,proto3" json:"title,omitempty"`
	ShortDesc     string                 `protobuf:"bytes,3,opt,name=short_desc,json=shortDesc,proto3" json:"short_desc,omitempty"`
	Description   string                 `protobuf:"bytes,4,opt,name=description,proto3" json:"description,omitempty"`
	UserId        string                 `protobuf:"bytes,5,opt,name=user_id,json=userId,proto3" json:"user_id,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *Post) Reset() {
	*x = Post{}
	mi := &file_post_post_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *Post) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Post) ProtoMessage() {}

func (x *Post) ProtoReflect() protoreflect.Message {
	mi := &file_post_post_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Post.ProtoReflect.Descriptor instead.
func (*Post) Descriptor() ([]byte, []int) {
	return file_post_post_proto_rawDescGZIP(), []int{0}
}

func (x *Post) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *Post) GetTitle() string {
	if x != nil {
		return x.Title
	}
	return ""
}

func (x *Post) GetShortDesc() string {
	if x != nil {
		return x.ShortDesc
	}
	return ""
}

func (x *Post) GetDescription() string {
	if x != nil {
		return x.Description
	}
	return ""
}

func (x *Post) GetUserId() string {
	if x != nil {
		return x.UserId
	}
	return ""
}

var File_post_post_proto protoreflect.FileDescriptor

const file_post_post_proto_rawDesc = "" +
	"\n" +
	"\x0fpost/post.proto\x12\x02pb\"\x86\x01\n" +
	"\x04Post\x12\x0e\n" +
	"\x02id\x18\x01 \x01(\tR\x02id\x12\x14\n" +
	"\x05title\x18\x02 \x01(\tR\x05title\x12\x1d\n" +
	"\n" +
	"short_desc\x18\x03 \x01(\tR\tshortDesc\x12 \n" +
	"\vdescription\x18\x04 \x01(\tR\vdescription\x12\x17\n" +
	"\auser_id\x18\x05 \x01(\tR\x06userIdB%Z#github.com/paweldyl/blog-backend/pbb\x06proto3"

var (
	file_post_post_proto_rawDescOnce sync.Once
	file_post_post_proto_rawDescData []byte
)

func file_post_post_proto_rawDescGZIP() []byte {
	file_post_post_proto_rawDescOnce.Do(func() {
		file_post_post_proto_rawDescData = protoimpl.X.CompressGZIP(unsafe.Slice(unsafe.StringData(file_post_post_proto_rawDesc), len(file_post_post_proto_rawDesc)))
	})
	return file_post_post_proto_rawDescData
}

var file_post_post_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_post_post_proto_goTypes = []any{
	(*Post)(nil), // 0: pb.Post
}
var file_post_post_proto_depIdxs = []int32{
	0, // [0:0] is the sub-list for method output_type
	0, // [0:0] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_post_post_proto_init() }
func file_post_post_proto_init() {
	if File_post_post_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: unsafe.Slice(unsafe.StringData(file_post_post_proto_rawDesc), len(file_post_post_proto_rawDesc)),
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_post_post_proto_goTypes,
		DependencyIndexes: file_post_post_proto_depIdxs,
		MessageInfos:      file_post_post_proto_msgTypes,
	}.Build()
	File_post_post_proto = out.File
	file_post_post_proto_goTypes = nil
	file_post_post_proto_depIdxs = nil
}
