// Wrapper for New Relic Go Agent

// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.23.0-devel
// 	protoc        v3.11.4
// source: go_agent.proto

package protos

import (
	context "context"
	proto "github.com/golang/protobuf/proto"
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

// This is a compile-time assertion that a sufficiently up-to-date version
// of the legacy proto package is being used.
const _ = proto.ProtoPackageIsVersion4

type AppConfig struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Name    string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	License string `protobuf:"bytes,2,opt,name=license,proto3" json:"license,omitempty"`
}

func (x *AppConfig) Reset() {
	*x = AppConfig{}
	if protoimpl.UnsafeEnabled {
		mi := &file_go_agent_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AppConfig) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AppConfig) ProtoMessage() {}

func (x *AppConfig) ProtoReflect() protoreflect.Message {
	mi := &file_go_agent_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AppConfig.ProtoReflect.Descriptor instead.
func (*AppConfig) Descriptor() ([]byte, []int) {
	return file_go_agent_proto_rawDescGZIP(), []int{0}
}

func (x *AppConfig) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *AppConfig) GetLicense() string {
	if x != nil {
		return x.License
	}
	return ""
}

type AppReply struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Message string `protobuf:"bytes,1,opt,name=message,proto3" json:"message,omitempty"`
}

func (x *AppReply) Reset() {
	*x = AppReply{}
	if protoimpl.UnsafeEnabled {
		mi := &file_go_agent_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AppReply) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AppReply) ProtoMessage() {}

func (x *AppReply) ProtoReflect() protoreflect.Message {
	mi := &file_go_agent_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AppReply.ProtoReflect.Descriptor instead.
func (*AppReply) Descriptor() ([]byte, []int) {
	return file_go_agent_proto_rawDescGZIP(), []int{1}
}

func (x *AppReply) GetMessage() string {
	if x != nil {
		return x.Message
	}
	return ""
}

var File_go_agent_proto protoreflect.FileDescriptor

var file_go_agent_proto_rawDesc = []byte{
	0x0a, 0x0e, 0x67, 0x6f, 0x5f, 0x61, 0x67, 0x65, 0x6e, 0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x12, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x73, 0x22, 0x39, 0x0a, 0x09, 0x41, 0x70, 0x70, 0x43,
	0x6f, 0x6e, 0x66, 0x69, 0x67, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x6c, 0x69, 0x63,
	0x65, 0x6e, 0x73, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x6c, 0x69, 0x63, 0x65,
	0x6e, 0x73, 0x65, 0x22, 0x24, 0x0a, 0x08, 0x41, 0x70, 0x70, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x12,
	0x18, 0x0a, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x32, 0x3c, 0x0a, 0x09, 0x43, 0x72, 0x65,
	0x61, 0x74, 0x65, 0x41, 0x70, 0x70, 0x12, 0x2f, 0x0a, 0x06, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65,
	0x12, 0x11, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x73, 0x2e, 0x41, 0x70, 0x70, 0x43, 0x6f, 0x6e,
	0x66, 0x69, 0x67, 0x1a, 0x10, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x73, 0x2e, 0x41, 0x70, 0x70,
	0x52, 0x65, 0x70, 0x6c, 0x79, 0x22, 0x00, 0x42, 0x0a, 0x5a, 0x08, 0x2e, 0x3b, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x73, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_go_agent_proto_rawDescOnce sync.Once
	file_go_agent_proto_rawDescData = file_go_agent_proto_rawDesc
)

func file_go_agent_proto_rawDescGZIP() []byte {
	file_go_agent_proto_rawDescOnce.Do(func() {
		file_go_agent_proto_rawDescData = protoimpl.X.CompressGZIP(file_go_agent_proto_rawDescData)
	})
	return file_go_agent_proto_rawDescData
}

var file_go_agent_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_go_agent_proto_goTypes = []interface{}{
	(*AppConfig)(nil), // 0: protos.AppConfig
	(*AppReply)(nil),  // 1: protos.AppReply
}
var file_go_agent_proto_depIdxs = []int32{
	0, // 0: protos.CreateApp.Create:input_type -> protos.AppConfig
	1, // 1: protos.CreateApp.Create:output_type -> protos.AppReply
	1, // [1:2] is the sub-list for method output_type
	0, // [0:1] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_go_agent_proto_init() }
func file_go_agent_proto_init() {
	if File_go_agent_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_go_agent_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AppConfig); i {
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
		file_go_agent_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AppReply); i {
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
			RawDescriptor: file_go_agent_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_go_agent_proto_goTypes,
		DependencyIndexes: file_go_agent_proto_depIdxs,
		MessageInfos:      file_go_agent_proto_msgTypes,
	}.Build()
	File_go_agent_proto = out.File
	file_go_agent_proto_rawDesc = nil
	file_go_agent_proto_goTypes = nil
	file_go_agent_proto_depIdxs = nil
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConnInterface

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion6

// CreateAppClient is the client API for CreateApp service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type CreateAppClient interface {
	Create(ctx context.Context, in *AppConfig, opts ...grpc.CallOption) (*AppReply, error)
}

type createAppClient struct {
	cc grpc.ClientConnInterface
}

func NewCreateAppClient(cc grpc.ClientConnInterface) CreateAppClient {
	return &createAppClient{cc}
}

func (c *createAppClient) Create(ctx context.Context, in *AppConfig, opts ...grpc.CallOption) (*AppReply, error) {
	out := new(AppReply)
	err := c.cc.Invoke(ctx, "/protos.CreateApp/Create", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// CreateAppServer is the server API for CreateApp service.
type CreateAppServer interface {
	Create(context.Context, *AppConfig) (*AppReply, error)
}

// UnimplementedCreateAppServer can be embedded to have forward compatible implementations.
type UnimplementedCreateAppServer struct {
}

func (*UnimplementedCreateAppServer) Create(context.Context, *AppConfig) (*AppReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Create not implemented")
}

func RegisterCreateAppServer(s *grpc.Server, srv CreateAppServer) {
	s.RegisterService(&_CreateApp_serviceDesc, srv)
}

func _CreateApp_Create_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(AppConfig)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CreateAppServer).Create(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/protos.CreateApp/Create",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CreateAppServer).Create(ctx, req.(*AppConfig))
	}
	return interceptor(ctx, in, info, handler)
}

var _CreateApp_serviceDesc = grpc.ServiceDesc{
	ServiceName: "protos.CreateApp",
	HandlerType: (*CreateAppServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Create",
			Handler:    _CreateApp_Create_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "go_agent.proto",
}
