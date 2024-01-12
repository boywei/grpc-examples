// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v3.6.1
// source: zero/zero.proto

package zero

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

type BallRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Name string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
}

func (x *BallRequest) Reset() {
	*x = BallRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_zero_zero_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *BallRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*BallRequest) ProtoMessage() {}

func (x *BallRequest) ProtoReflect() protoreflect.Message {
	mi := &file_zero_zero_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use BallRequest.ProtoReflect.Descriptor instead.
func (*BallRequest) Descriptor() ([]byte, []int) {
	return file_zero_zero_proto_rawDescGZIP(), []int{0}
}

func (x *BallRequest) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

type BallReply struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Reply string `protobuf:"bytes,1,opt,name=reply,proto3" json:"reply,omitempty"`
}

func (x *BallReply) Reset() {
	*x = BallReply{}
	if protoimpl.UnsafeEnabled {
		mi := &file_zero_zero_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *BallReply) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*BallReply) ProtoMessage() {}

func (x *BallReply) ProtoReflect() protoreflect.Message {
	mi := &file_zero_zero_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use BallReply.ProtoReflect.Descriptor instead.
func (*BallReply) Descriptor() ([]byte, []int) {
	return file_zero_zero_proto_rawDescGZIP(), []int{1}
}

func (x *BallReply) GetReply() string {
	if x != nil {
		return x.Reply
	}
	return ""
}

var File_zero_zero_proto protoreflect.FileDescriptor

var file_zero_zero_proto_rawDesc = []byte{
	0x0a, 0x0f, 0x7a, 0x65, 0x72, 0x6f, 0x2f, 0x7a, 0x65, 0x72, 0x6f, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x22, 0x21, 0x0a, 0x0b, 0x42, 0x61, 0x6c, 0x6c, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04,
	0x6e, 0x61, 0x6d, 0x65, 0x22, 0x21, 0x0a, 0x09, 0x42, 0x61, 0x6c, 0x6c, 0x52, 0x65, 0x70, 0x6c,
	0x79, 0x12, 0x14, 0x0a, 0x05, 0x72, 0x65, 0x70, 0x6c, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x05, 0x72, 0x65, 0x70, 0x6c, 0x79, 0x32, 0x2c, 0x0a, 0x04, 0x47, 0x61, 0x6d, 0x65, 0x12,
	0x24, 0x0a, 0x08, 0x50, 0x6c, 0x61, 0x79, 0x42, 0x61, 0x6c, 0x6c, 0x12, 0x0c, 0x2e, 0x42, 0x61,
	0x6c, 0x6c, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x0a, 0x2e, 0x42, 0x61, 0x6c, 0x6c,
	0x52, 0x65, 0x70, 0x6c, 0x79, 0x42, 0x2b, 0x5a, 0x29, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e,
	0x63, 0x6f, 0x6d, 0x2f, 0x62, 0x6f, 0x79, 0x77, 0x65, 0x69, 0x2f, 0x67, 0x72, 0x70, 0x63, 0x2d,
	0x65, 0x78, 0x61, 0x6d, 0x70, 0x6c, 0x65, 0x73, 0x2f, 0x7a, 0x65, 0x72, 0x6f, 0x2f, 0x7a, 0x65,
	0x72, 0x6f, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_zero_zero_proto_rawDescOnce sync.Once
	file_zero_zero_proto_rawDescData = file_zero_zero_proto_rawDesc
)

func file_zero_zero_proto_rawDescGZIP() []byte {
	file_zero_zero_proto_rawDescOnce.Do(func() {
		file_zero_zero_proto_rawDescData = protoimpl.X.CompressGZIP(file_zero_zero_proto_rawDescData)
	})
	return file_zero_zero_proto_rawDescData
}

var file_zero_zero_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_zero_zero_proto_goTypes = []interface{}{
	(*BallRequest)(nil), // 0: BallRequest
	(*BallReply)(nil),   // 1: BallReply
}
var file_zero_zero_proto_depIdxs = []int32{
	0, // 0: Game.PlayBall:input_type -> BallRequest
	1, // 1: Game.PlayBall:output_type -> BallReply
	1, // [1:2] is the sub-list for method output_type
	0, // [0:1] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_zero_zero_proto_init() }
func file_zero_zero_proto_init() {
	if File_zero_zero_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_zero_zero_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*BallRequest); i {
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
		file_zero_zero_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*BallReply); i {
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
			RawDescriptor: file_zero_zero_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_zero_zero_proto_goTypes,
		DependencyIndexes: file_zero_zero_proto_depIdxs,
		MessageInfos:      file_zero_zero_proto_msgTypes,
	}.Build()
	File_zero_zero_proto = out.File
	file_zero_zero_proto_rawDesc = nil
	file_zero_zero_proto_goTypes = nil
	file_zero_zero_proto_depIdxs = nil
}
