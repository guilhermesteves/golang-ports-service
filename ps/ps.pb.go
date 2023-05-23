// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.26.0
// 	protoc        v3.8.0
// source: ps.proto

package ps

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
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

type Port struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ID      string `protobuf:"bytes,1,opt,name=ID,proto3" json:"ID,omitempty"`
	Content []byte `protobuf:"bytes,2,opt,name=Content,proto3" json:"Content,omitempty"`
}

func (x *Port) Reset() {
	*x = Port{}
	if protoimpl.UnsafeEnabled {
		mi := &file_ps_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Port) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Port) ProtoMessage() {}

func (x *Port) ProtoReflect() protoreflect.Message {
	mi := &file_ps_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Port.ProtoReflect.Descriptor instead.
func (*Port) Descriptor() ([]byte, []int) {
	return file_ps_proto_rawDescGZIP(), []int{0}
}

func (x *Port) GetID() string {
	if x != nil {
		return x.ID
	}
	return ""
}

func (x *Port) GetContent() []byte {
	if x != nil {
		return x.Content
	}
	return nil
}

type Empty struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *Empty) Reset() {
	*x = Empty{}
	if protoimpl.UnsafeEnabled {
		mi := &file_ps_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Empty) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Empty) ProtoMessage() {}

func (x *Empty) ProtoReflect() protoreflect.Message {
	mi := &file_ps_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Empty.ProtoReflect.Descriptor instead.
func (*Empty) Descriptor() ([]byte, []int) {
	return file_ps_proto_rawDescGZIP(), []int{1}
}

var File_ps_proto protoreflect.FileDescriptor

var file_ps_proto_rawDesc = []byte{
	0x0a, 0x08, 0x70, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x02, 0x70, 0x73, 0x22, 0x30,
	0x0a, 0x04, 0x50, 0x6f, 0x72, 0x74, 0x12, 0x0e, 0x0a, 0x02, 0x49, 0x44, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x02, 0x49, 0x44, 0x12, 0x18, 0x0a, 0x07, 0x43, 0x6f, 0x6e, 0x74, 0x65, 0x6e,
	0x74, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x07, 0x43, 0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74,
	0x22, 0x07, 0x0a, 0x05, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x32, 0x34, 0x0a, 0x0b, 0x50, 0x6f, 0x72,
	0x74, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x25, 0x0a, 0x0e, 0x49, 0x6e, 0x73, 0x65,
	0x72, 0x74, 0x4f, 0x72, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x12, 0x08, 0x2e, 0x70, 0x73, 0x2e,
	0x50, 0x6f, 0x72, 0x74, 0x1a, 0x09, 0x2e, 0x70, 0x73, 0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x42,
	0x05, 0x5a, 0x03, 0x70, 0x73, 0x2f, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_ps_proto_rawDescOnce sync.Once
	file_ps_proto_rawDescData = file_ps_proto_rawDesc
)

func file_ps_proto_rawDescGZIP() []byte {
	file_ps_proto_rawDescOnce.Do(func() {
		file_ps_proto_rawDescData = protoimpl.X.CompressGZIP(file_ps_proto_rawDescData)
	})
	return file_ps_proto_rawDescData
}

var file_ps_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_ps_proto_goTypes = []interface{}{
	(*Port)(nil),  // 0: ps.Port
	(*Empty)(nil), // 1: ps.Empty
}
var file_ps_proto_depIdxs = []int32{
	0, // 0: ps.PortService.InsertOrUpdate:input_type -> ps.Port
	1, // 1: ps.PortService.InsertOrUpdate:output_type -> ps.Empty
	1, // [1:2] is the sub-list for method output_type
	0, // [0:1] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_ps_proto_init() }
func file_ps_proto_init() {
	if File_ps_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_ps_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Port); i {
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
		file_ps_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Empty); i {
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
			RawDescriptor: file_ps_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_ps_proto_goTypes,
		DependencyIndexes: file_ps_proto_depIdxs,
		MessageInfos:      file_ps_proto_msgTypes,
	}.Build()
	File_ps_proto = out.File
	file_ps_proto_rawDesc = nil
	file_ps_proto_goTypes = nil
	file_ps_proto_depIdxs = nil
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConnInterface

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion6

// PortServiceClient is the client API for PortService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type PortServiceClient interface {
	InsertOrUpdate(ctx context.Context, in *Port, opts ...grpc.CallOption) (*Empty, error)
}

type portServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewPortServiceClient(cc grpc.ClientConnInterface) PortServiceClient {
	return &portServiceClient{cc}
}

func (c *portServiceClient) InsertOrUpdate(ctx context.Context, in *Port, opts ...grpc.CallOption) (*Empty, error) {
	out := new(Empty)
	err := c.cc.Invoke(ctx, "/ps.PortService/InsertOrUpdate", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// PortServiceServer is the server API for PortService service.
type PortServiceServer interface {
	InsertOrUpdate(context.Context, *Port) (*Empty, error)
}

// UnimplementedPortServiceServer can be embedded to have forward compatible implementations.
type UnimplementedPortServiceServer struct {
}

func (*UnimplementedPortServiceServer) InsertOrUpdate(context.Context, *Port) (*Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method InsertOrUpdate not implemented")
}

func RegisterPortServiceServer(s *grpc.Server, srv PortServiceServer) {
	s.RegisterService(&_PortService_serviceDesc, srv)
}

func _PortService_InsertOrUpdate_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Port)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PortServiceServer).InsertOrUpdate(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/ps.PortService/InsertOrUpdate",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PortServiceServer).InsertOrUpdate(ctx, req.(*Port))
	}
	return interceptor(ctx, in, info, handler)
}

var _PortService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "ps.PortService",
	HandlerType: (*PortServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "InsertOrUpdate",
			Handler:    _PortService_InsertOrUpdate_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "ps.proto",
}