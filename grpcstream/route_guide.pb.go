// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.27.1
// 	protoc        v3.17.3
// source: grpcstream/route_guide.proto

package grpcstream

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

type RouteNote struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Message string `protobuf:"bytes,2,opt,name=message,proto3" json:"message,omitempty"`
}

func (x *RouteNote) Reset() {
	*x = RouteNote{}
	if protoimpl.UnsafeEnabled {
		mi := &file_grpcstream_route_guide_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RouteNote) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RouteNote) ProtoMessage() {}

func (x *RouteNote) ProtoReflect() protoreflect.Message {
	mi := &file_grpcstream_route_guide_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RouteNote.ProtoReflect.Descriptor instead.
func (*RouteNote) Descriptor() ([]byte, []int) {
	return file_grpcstream_route_guide_proto_rawDescGZIP(), []int{0}
}

func (x *RouteNote) GetMessage() string {
	if x != nil {
		return x.Message
	}
	return ""
}

var File_grpcstream_route_guide_proto protoreflect.FileDescriptor

var file_grpcstream_route_guide_proto_rawDesc = []byte{
	0x0a, 0x1c, 0x67, 0x72, 0x70, 0x63, 0x73, 0x74, 0x72, 0x65, 0x61, 0x6d, 0x2f, 0x72, 0x6f, 0x75,
	0x74, 0x65, 0x5f, 0x67, 0x75, 0x69, 0x64, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x0a,
	0x67, 0x72, 0x70, 0x63, 0x73, 0x74, 0x72, 0x65, 0x61, 0x6d, 0x22, 0x25, 0x0a, 0x09, 0x52, 0x6f,
	0x75, 0x74, 0x65, 0x4e, 0x6f, 0x74, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61,
	0x67, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67,
	0x65, 0x32, 0x4d, 0x0a, 0x0a, 0x52, 0x6f, 0x75, 0x74, 0x65, 0x47, 0x75, 0x69, 0x64, 0x65, 0x12,
	0x3f, 0x0a, 0x09, 0x52, 0x6f, 0x75, 0x74, 0x65, 0x43, 0x68, 0x61, 0x74, 0x12, 0x15, 0x2e, 0x67,
	0x72, 0x70, 0x63, 0x73, 0x74, 0x72, 0x65, 0x61, 0x6d, 0x2e, 0x52, 0x6f, 0x75, 0x74, 0x65, 0x4e,
	0x6f, 0x74, 0x65, 0x1a, 0x15, 0x2e, 0x67, 0x72, 0x70, 0x63, 0x73, 0x74, 0x72, 0x65, 0x61, 0x6d,
	0x2e, 0x52, 0x6f, 0x75, 0x74, 0x65, 0x4e, 0x6f, 0x74, 0x65, 0x22, 0x00, 0x28, 0x01, 0x30, 0x01,
	0x42, 0x90, 0x01, 0x0a, 0x0e, 0x63, 0x6f, 0x6d, 0x2e, 0x67, 0x72, 0x70, 0x63, 0x73, 0x74, 0x72,
	0x65, 0x61, 0x6d, 0x42, 0x0f, 0x52, 0x6f, 0x75, 0x74, 0x65, 0x47, 0x75, 0x69, 0x64, 0x65, 0x50,
	0x72, 0x6f, 0x74, 0x6f, 0x50, 0x01, 0x5a, 0x25, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63,
	0x6f, 0x6d, 0x2f, 0x73, 0x6c, 0x6e, 0x74, 0x6f, 0x70, 0x70, 0x2f, 0x6e, 0x6f, 0x63, 0x6c, 0x6f,
	0x75, 0x64, 0x2f, 0x67, 0x72, 0x70, 0x63, 0x73, 0x74, 0x72, 0x65, 0x61, 0x6d, 0xa2, 0x02, 0x03,
	0x47, 0x58, 0x58, 0xaa, 0x02, 0x0a, 0x47, 0x72, 0x70, 0x63, 0x73, 0x74, 0x72, 0x65, 0x61, 0x6d,
	0xca, 0x02, 0x0a, 0x47, 0x72, 0x70, 0x63, 0x73, 0x74, 0x72, 0x65, 0x61, 0x6d, 0xe2, 0x02, 0x16,
	0x47, 0x72, 0x70, 0x63, 0x73, 0x74, 0x72, 0x65, 0x61, 0x6d, 0x5c, 0x47, 0x50, 0x42, 0x4d, 0x65,
	0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0xea, 0x02, 0x0a, 0x47, 0x72, 0x70, 0x63, 0x73, 0x74, 0x72,
	0x65, 0x61, 0x6d, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_grpcstream_route_guide_proto_rawDescOnce sync.Once
	file_grpcstream_route_guide_proto_rawDescData = file_grpcstream_route_guide_proto_rawDesc
)

func file_grpcstream_route_guide_proto_rawDescGZIP() []byte {
	file_grpcstream_route_guide_proto_rawDescOnce.Do(func() {
		file_grpcstream_route_guide_proto_rawDescData = protoimpl.X.CompressGZIP(file_grpcstream_route_guide_proto_rawDescData)
	})
	return file_grpcstream_route_guide_proto_rawDescData
}

var file_grpcstream_route_guide_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_grpcstream_route_guide_proto_goTypes = []interface{}{
	(*RouteNote)(nil), // 0: grpcstream.RouteNote
}
var file_grpcstream_route_guide_proto_depIdxs = []int32{
	0, // 0: grpcstream.RouteGuide.RouteChat:input_type -> grpcstream.RouteNote
	0, // 1: grpcstream.RouteGuide.RouteChat:output_type -> grpcstream.RouteNote
	1, // [1:2] is the sub-list for method output_type
	0, // [0:1] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_grpcstream_route_guide_proto_init() }
func file_grpcstream_route_guide_proto_init() {
	if File_grpcstream_route_guide_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_grpcstream_route_guide_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*RouteNote); i {
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
			RawDescriptor: file_grpcstream_route_guide_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_grpcstream_route_guide_proto_goTypes,
		DependencyIndexes: file_grpcstream_route_guide_proto_depIdxs,
		MessageInfos:      file_grpcstream_route_guide_proto_msgTypes,
	}.Build()
	File_grpcstream_route_guide_proto = out.File
	file_grpcstream_route_guide_proto_rawDesc = nil
	file_grpcstream_route_guide_proto_goTypes = nil
	file_grpcstream_route_guide_proto_depIdxs = nil
}
