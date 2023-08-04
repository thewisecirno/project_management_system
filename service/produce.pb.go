// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.26.0
// 	protoc        v4.23.4
// source: produce.proto

package service

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

type Product_Request struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ProdId int32 `protobuf:"varint,1,opt,name=prod_id,json=prodId,proto3" json:"prod_id,omitempty"`
}

func (x *Product_Request) Reset() {
	*x = Product_Request{}
	if protoimpl.UnsafeEnabled {
		mi := &file_produce_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Product_Request) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Product_Request) ProtoMessage() {}

func (x *Product_Request) ProtoReflect() protoreflect.Message {
	mi := &file_produce_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Product_Request.ProtoReflect.Descriptor instead.
func (*Product_Request) Descriptor() ([]byte, []int) {
	return file_produce_proto_rawDescGZIP(), []int{0}
}

func (x *Product_Request) GetProdId() int32 {
	if x != nil {
		return x.ProdId
	}
	return 0
}

type Produce_Response struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ProdStock int32 `protobuf:"varint,2,opt,name=prod_stock,json=prodStock,proto3" json:"prod_stock,omitempty"`
}

func (x *Produce_Response) Reset() {
	*x = Produce_Response{}
	if protoimpl.UnsafeEnabled {
		mi := &file_produce_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Produce_Response) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Produce_Response) ProtoMessage() {}

func (x *Produce_Response) ProtoReflect() protoreflect.Message {
	mi := &file_produce_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Produce_Response.ProtoReflect.Descriptor instead.
func (*Produce_Response) Descriptor() ([]byte, []int) {
	return file_produce_proto_rawDescGZIP(), []int{1}
}

func (x *Produce_Response) GetProdStock() int32 {
	if x != nil {
		return x.ProdStock
	}
	return 0
}

var File_produce_proto protoreflect.FileDescriptor

var file_produce_proto_rawDesc = []byte{
	0x0a, 0x0d, 0x70, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12,
	0x07, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x22, 0x2a, 0x0a, 0x0f, 0x50, 0x72, 0x6f, 0x64,
	0x75, 0x63, 0x74, 0x5f, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x17, 0x0a, 0x07, 0x70,
	0x72, 0x6f, 0x64, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x06, 0x70, 0x72,
	0x6f, 0x64, 0x49, 0x64, 0x22, 0x31, 0x0a, 0x10, 0x50, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x65, 0x5f,
	0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x1d, 0x0a, 0x0a, 0x70, 0x72, 0x6f, 0x64,
	0x5f, 0x73, 0x74, 0x6f, 0x63, 0x6b, 0x18, 0x02, 0x20, 0x01, 0x28, 0x05, 0x52, 0x09, 0x70, 0x72,
	0x6f, 0x64, 0x53, 0x74, 0x6f, 0x63, 0x6b, 0x32, 0x52, 0x0a, 0x0b, 0x70, 0x72, 0x6f, 0x64, 0x53,
	0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x43, 0x0a, 0x0c, 0x47, 0x65, 0x74, 0x50, 0x72, 0x6f,
	0x64, 0x53, 0x74, 0x6f, 0x63, 0x6b, 0x12, 0x18, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65,
	0x2e, 0x50, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x5f, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x1a, 0x19, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x50, 0x72, 0x6f, 0x64, 0x75,
	0x63, 0x65, 0x5f, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x42, 0x0c, 0x5a, 0x0a, 0x2e,
	0x2e, 0x2f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x33,
}

var (
	file_produce_proto_rawDescOnce sync.Once
	file_produce_proto_rawDescData = file_produce_proto_rawDesc
)

func file_produce_proto_rawDescGZIP() []byte {
	file_produce_proto_rawDescOnce.Do(func() {
		file_produce_proto_rawDescData = protoimpl.X.CompressGZIP(file_produce_proto_rawDescData)
	})
	return file_produce_proto_rawDescData
}

var file_produce_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_produce_proto_goTypes = []interface{}{
	(*Product_Request)(nil),  // 0: service.Product_Request
	(*Produce_Response)(nil), // 1: service.Produce_Response
}
var file_produce_proto_depIdxs = []int32{
	0, // 0: service.prodService.GetProdStock:input_type -> service.Product_Request
	1, // 1: service.prodService.GetProdStock:output_type -> service.Produce_Response
	1, // [1:2] is the sub-list for method output_type
	0, // [0:1] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_produce_proto_init() }
func file_produce_proto_init() {
	if File_produce_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_produce_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Product_Request); i {
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
		file_produce_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Produce_Response); i {
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
			RawDescriptor: file_produce_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_produce_proto_goTypes,
		DependencyIndexes: file_produce_proto_depIdxs,
		MessageInfos:      file_produce_proto_msgTypes,
	}.Build()
	File_produce_proto = out.File
	file_produce_proto_rawDesc = nil
	file_produce_proto_goTypes = nil
	file_produce_proto_depIdxs = nil
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConnInterface

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion6

// ProdServiceClient is the client API for ProdService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type ProdServiceClient interface {
	GetProdStock(ctx context.Context, in *Product_Request, opts ...grpc.CallOption) (*Produce_Response, error)
}

type prodServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewProdServiceClient(cc grpc.ClientConnInterface) ProdServiceClient {
	return &prodServiceClient{cc}
}

func (c *prodServiceClient) GetProdStock(ctx context.Context, in *Product_Request, opts ...grpc.CallOption) (*Produce_Response, error) {
	out := new(Produce_Response)
	err := c.cc.Invoke(ctx, "/service.prodService/GetProdStock", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// ProdServiceServer is the server API for ProdService service.
type ProdServiceServer interface {
	GetProdStock(context.Context, *Product_Request) (*Produce_Response, error)
}

// UnimplementedProdServiceServer can be embedded to have forward compatible implementations.
type UnimplementedProdServiceServer struct {
}

func (*UnimplementedProdServiceServer) GetProdStock(context.Context, *Product_Request) (*Produce_Response, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetProdStock not implemented")
}

func RegisterProdServiceServer(s *grpc.Server, srv *ProduceService) {
	s.RegisterService(&_ProdService_serviceDesc, srv)
}

func _ProdService_GetProdStock_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Product_Request)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ProdServiceServer).GetProdStock(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/service.prodService/GetProdStock",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ProdServiceServer).GetProdStock(ctx, req.(*Product_Request))
	}
	return interceptor(ctx, in, info, handler)
}

var _ProdService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "service.prodService",
	HandlerType: (*ProdServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetProdStock",
			Handler:    _ProdService_GetProdStock_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "produce.proto",
}
