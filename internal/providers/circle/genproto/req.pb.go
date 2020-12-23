// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.21.0-devel
// 	protoc        v3.6.1
// source: req.proto

package provider

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

type Req struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Payload string `protobuf:"bytes,1,opt,name=payload,proto3" json:"payload,omitempty"`
}

func (x *Req) Reset() {
	*x = Req{}
	if protoimpl.UnsafeEnabled {
		mi := &file_req_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Req) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Req) ProtoMessage() {}

func (x *Req) ProtoReflect() protoreflect.Message {
	mi := &file_req_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Req.ProtoReflect.Descriptor instead.
func (*Req) Descriptor() ([]byte, []int) {
	return file_req_proto_rawDescGZIP(), []int{0}
}

func (x *Req) GetPayload() string {
	if x != nil {
		return x.Payload
	}
	return ""
}

type Res struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Payload string `protobuf:"bytes,1,opt,name=payload,proto3" json:"payload,omitempty"`
}

func (x *Res) Reset() {
	*x = Res{}
	if protoimpl.UnsafeEnabled {
		mi := &file_req_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Res) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Res) ProtoMessage() {}

func (x *Res) ProtoReflect() protoreflect.Message {
	mi := &file_req_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Res.ProtoReflect.Descriptor instead.
func (*Res) Descriptor() ([]byte, []int) {
	return file_req_proto_rawDescGZIP(), []int{1}
}

func (x *Res) GetPayload() string {
	if x != nil {
		return x.Payload
	}
	return ""
}

var File_req_proto protoreflect.FileDescriptor

var file_req_proto_rawDesc = []byte{
	0x0a, 0x09, 0x72, 0x65, 0x71, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x08, 0x70, 0x72, 0x6f,
	0x76, 0x69, 0x64, 0x65, 0x72, 0x22, 0x1f, 0x0a, 0x03, 0x72, 0x65, 0x71, 0x12, 0x18, 0x0a, 0x07,
	0x70, 0x61, 0x79, 0x6c, 0x6f, 0x61, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x70,
	0x61, 0x79, 0x6c, 0x6f, 0x61, 0x64, 0x22, 0x1f, 0x0a, 0x03, 0x72, 0x65, 0x73, 0x12, 0x18, 0x0a,
	0x07, 0x70, 0x61, 0x79, 0x6c, 0x6f, 0x61, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07,
	0x70, 0x61, 0x79, 0x6c, 0x6f, 0x61, 0x64, 0x32, 0x32, 0x0a, 0x04, 0x43, 0x61, 0x72, 0x64, 0x12,
	0x2a, 0x0a, 0x0a, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x43, 0x61, 0x72, 0x64, 0x12, 0x0d, 0x2e,
	0x70, 0x72, 0x6f, 0x76, 0x69, 0x64, 0x65, 0x72, 0x2e, 0x72, 0x65, 0x71, 0x1a, 0x0d, 0x2e, 0x70,
	0x72, 0x6f, 0x76, 0x69, 0x64, 0x65, 0x72, 0x2e, 0x72, 0x65, 0x73, 0x32, 0x36, 0x0a, 0x07, 0x50,
	0x61, 0x79, 0x6d, 0x65, 0x6e, 0x74, 0x12, 0x2b, 0x0a, 0x0b, 0x6d, 0x61, 0x6b, 0x65, 0x50, 0x61,
	0x79, 0x6d, 0x65, 0x6e, 0x74, 0x12, 0x0d, 0x2e, 0x70, 0x72, 0x6f, 0x76, 0x69, 0x64, 0x65, 0x72,
	0x2e, 0x72, 0x65, 0x71, 0x1a, 0x0d, 0x2e, 0x70, 0x72, 0x6f, 0x76, 0x69, 0x64, 0x65, 0x72, 0x2e,
	0x72, 0x65, 0x73, 0x42, 0x0c, 0x5a, 0x0a, 0x2e, 0x3b, 0x70, 0x72, 0x6f, 0x76, 0x69, 0x64, 0x65,
	0x72, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_req_proto_rawDescOnce sync.Once
	file_req_proto_rawDescData = file_req_proto_rawDesc
)

func file_req_proto_rawDescGZIP() []byte {
	file_req_proto_rawDescOnce.Do(func() {
		file_req_proto_rawDescData = protoimpl.X.CompressGZIP(file_req_proto_rawDescData)
	})
	return file_req_proto_rawDescData
}

var file_req_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_req_proto_goTypes = []interface{}{
	(*Req)(nil), // 0: provider.req
	(*Res)(nil), // 1: provider.res
}
var file_req_proto_depIdxs = []int32{
	0, // 0: provider.Card.createCard:input_type -> provider.req
	0, // 1: provider.Payment.makePayment:input_type -> provider.req
	1, // 2: provider.Card.createCard:output_type -> provider.res
	1, // 3: provider.Payment.makePayment:output_type -> provider.res
	2, // [2:4] is the sub-list for method output_type
	0, // [0:2] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_req_proto_init() }
func file_req_proto_init() {
	if File_req_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_req_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Req); i {
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
		file_req_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Res); i {
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
			RawDescriptor: file_req_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   2,
		},
		GoTypes:           file_req_proto_goTypes,
		DependencyIndexes: file_req_proto_depIdxs,
		MessageInfos:      file_req_proto_msgTypes,
	}.Build()
	File_req_proto = out.File
	file_req_proto_rawDesc = nil
	file_req_proto_goTypes = nil
	file_req_proto_depIdxs = nil
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConnInterface

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion6

// CardClient is the client API for Card service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type CardClient interface {
	CreateCard(ctx context.Context, in *Req, opts ...grpc.CallOption) (*Res, error)
}

type cardClient struct {
	cc grpc.ClientConnInterface
}

func NewCardClient(cc grpc.ClientConnInterface) CardClient {
	return &cardClient{cc}
}

func (c *cardClient) CreateCard(ctx context.Context, in *Req, opts ...grpc.CallOption) (*Res, error) {
	out := new(Res)
	err := c.cc.Invoke(ctx, "/provider.Card/createCard", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// CardServer is the server API for Card service.
type CardServer interface {
	CreateCard(context.Context, *Req) (*Res, error)
}

// UnimplementedCardServer can be embedded to have forward compatible implementations.
type UnimplementedCardServer struct {
}

func (*UnimplementedCardServer) CreateCard(context.Context, *Req) (*Res, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateCard not implemented")
}

func RegisterCardServer(s *grpc.Server, srv CardServer) {
	s.RegisterService(&_Card_serviceDesc, srv)
}

func _Card_CreateCard_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Req)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CardServer).CreateCard(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/provider.Card/CreateCard",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CardServer).CreateCard(ctx, req.(*Req))
	}
	return interceptor(ctx, in, info, handler)
}

var _Card_serviceDesc = grpc.ServiceDesc{
	ServiceName: "provider.Card",
	HandlerType: (*CardServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "createCard",
			Handler:    _Card_CreateCard_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "req.proto",
}

// PaymentClient is the client API for Payment service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type PaymentClient interface {
	MakePayment(ctx context.Context, in *Req, opts ...grpc.CallOption) (*Res, error)
}

type paymentClient struct {
	cc grpc.ClientConnInterface
}

func NewPaymentClient(cc grpc.ClientConnInterface) PaymentClient {
	return &paymentClient{cc}
}

func (c *paymentClient) MakePayment(ctx context.Context, in *Req, opts ...grpc.CallOption) (*Res, error) {
	out := new(Res)
	err := c.cc.Invoke(ctx, "/provider.Payment/makePayment", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// PaymentServer is the server API for Payment service.
type PaymentServer interface {
	MakePayment(context.Context, *Req) (*Res, error)
}

// UnimplementedPaymentServer can be embedded to have forward compatible implementations.
type UnimplementedPaymentServer struct {
}

func (*UnimplementedPaymentServer) MakePayment(context.Context, *Req) (*Res, error) {
	return nil, status.Errorf(codes.Unimplemented, "method MakePayment not implemented")
}

func RegisterPaymentServer(s *grpc.Server, srv PaymentServer) {
	s.RegisterService(&_Payment_serviceDesc, srv)
}

func _Payment_MakePayment_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Req)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PaymentServer).MakePayment(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/provider.Payment/MakePayment",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PaymentServer).MakePayment(ctx, req.(*Req))
	}
	return interceptor(ctx, in, info, handler)
}

var _Payment_serviceDesc = grpc.ServiceDesc{
	ServiceName: "provider.Payment",
	HandlerType: (*PaymentServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "makePayment",
			Handler:    _Payment_MakePayment_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "req.proto",
}
