// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: service/ergo.proto

package ergocon

import (
	context "context"
	fmt "fmt"
	schema "github.com/bregydoc/ergo/schema"
	proto "github.com/golang/protobuf/proto"
	grpc "google.golang.org/grpc"
	math "math"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

func init() { proto.RegisterFile("service/ergo.proto", fileDescriptor_179185ae5560f7fa) }

var fileDescriptor_179185ae5560f7fa = []byte{
	// 301 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x84, 0x91, 0xcd, 0x4e, 0xf2, 0x40,
	0x14, 0x86, 0x4b, 0xbe, 0x2f, 0x9a, 0xcc, 0x0a, 0x4e, 0xc5, 0x9f, 0x2e, 0xba, 0xe8, 0x05, 0x40,
	0xa2, 0x7b, 0xa3, 0xfc, 0xa9, 0x0b, 0xd0, 0x60, 0xdc, 0xb8, 0x1b, 0x86, 0xd7, 0xda, 0xd0, 0x76,
	0xc8, 0x99, 0x52, 0x12, 0xaf, 0xc4, 0xa5, 0x97, 0xe3, 0xd2, 0x4b, 0x30, 0x78, 0x23, 0x86, 0xc2,
	0x14, 0x1b, 0x17, 0xae, 0x66, 0xe6, 0x7d, 0x9f, 0x79, 0x72, 0x92, 0x23, 0xc8, 0x80, 0xf3, 0x48,
	0xa1, 0x0d, 0x0e, 0x75, 0x6b, 0xce, 0x3a, 0xd3, 0xb4, 0xbf, 0xbe, 0x2b, 0x9d, 0x7a, 0xae, 0x51,
	0xcf, 0x48, 0x64, 0x7b, 0x73, 0x6c, 0xda, 0xd3, 0xb7, 0x7f, 0xe2, 0x7f, 0x9f, 0x43, 0x4d, 0xe7,
	0xa2, 0x3e, 0x46, 0x18, 0x99, 0x0c, 0x3c, 0xc2, 0xb2, 0xcf, 0xac, 0x99, 0x1a, 0xad, 0x2d, 0x5b,
	0x3c, 0xef, 0x81, 0xa9, 0xd7, 0xac, 0x44, 0x37, 0xa9, 0xc9, 0x64, 0xaa, 0x10, 0x38, 0xd4, 0x15,
	0x0d, 0xfb, 0x7f, 0xb0, 0x88, 0xe3, 0x8d, 0xa0, 0xa4, 0xcb, 0xe8, 0x2f, 0x89, 0xdb, 0xd5, 0xa9,
	0x59, 0xc4, 0x59, 0xd1, 0x5c, 0x9a, 0xeb, 0x45, 0x22, 0x53, 0x3a, 0xb4, 0xfc, 0xb6, 0xdc, 0xe6,
	0x1e, 0x55, 0x3c, 0x45, 0x56, 0x48, 0x8e, 0xaa, 0x92, 0x1e, 0x72, 0xc4, 0x7a, 0x0e, 0xa6, 0x83,
	0x5f, 0xa2, 0x1e, 0x72, 0xaf, 0x5e, 0xd1, 0xf4, 0x90, 0x07, 0x0e, 0x5d, 0x09, 0x77, 0x88, 0x44,
	0x73, 0xf4, 0x82, 0x11, 0x96, 0x43, 0x18, 0x23, 0x43, 0x18, 0x3a, 0xb6, 0xe8, 0x2e, 0xbc, 0x93,
	0x2c, 0x13, 0xe3, 0x95, 0xea, 0x07, 0x03, 0xb6, 0x7c, 0xe0, 0xd0, 0x85, 0x68, 0x8e, 0xa1, 0x10,
	0xe5, 0x18, 0x00, 0xd3, 0x89, 0x54, 0xb3, 0xdb, 0xa7, 0x35, 0x42, 0xee, 0x0f, 0xd5, 0xba, 0xea,
	0x48, 0x35, 0xdb, 0x8d, 0x62, 0xe1, 0xc0, 0xe9, 0x9c, 0xbc, 0xaf, 0xfc, 0xda, 0xc7, 0xca, 0xaf,
	0x7d, 0xae, 0xfc, 0xda, 0xeb, 0x97, 0xef, 0x3c, 0xda, 0x95, 0x4e, 0xf6, 0x8a, 0x25, 0x9e, 0x7d,
	0x07, 0x00, 0x00, 0xff, 0xff, 0x74, 0x55, 0x31, 0x00, 0xf8, 0x01, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// ErgoClient is the client API for Ergo service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type ErgoClient interface {
	RegisterNewError(ctx context.Context, in *schema.ErrorSeed, opts ...grpc.CallOption) (*schema.ErrorInstance, error)
	RegisterFullError(ctx context.Context, in *schema.FullErrorSeed, opts ...grpc.CallOption) (*schema.ErrorInstance, error)
	ConsultErrorAsHuman(ctx context.Context, in *schema.ConsultAsHuman, opts ...grpc.CallOption) (*schema.ErrorHuman, error)
	ConsultErrorAsDeveloper(ctx context.Context, in *schema.ConsultAsDev, opts ...grpc.CallOption) (*schema.ErrorDev, error)
	// Save new messages
	MemorizeNewMessages(ctx context.Context, in *schema.NewMessageParams, opts ...grpc.CallOption) (*schema.UserMessages, error)
	// Save new feedback
	ReceiveFeedbackOfUser(ctx context.Context, in *schema.NewFeedBack, opts ...grpc.CallOption) (*schema.Feedback, error)
}

type ergoClient struct {
	cc *grpc.ClientConn
}

func NewErgoClient(cc *grpc.ClientConn) ErgoClient {
	return &ergoClient{cc}
}

func (c *ergoClient) RegisterNewError(ctx context.Context, in *schema.ErrorSeed, opts ...grpc.CallOption) (*schema.ErrorInstance, error) {
	out := new(schema.ErrorInstance)
	err := c.cc.Invoke(ctx, "/ergocon.Ergo/RegisterNewError", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *ergoClient) RegisterFullError(ctx context.Context, in *schema.FullErrorSeed, opts ...grpc.CallOption) (*schema.ErrorInstance, error) {
	out := new(schema.ErrorInstance)
	err := c.cc.Invoke(ctx, "/ergocon.Ergo/RegisterFullError", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *ergoClient) ConsultErrorAsHuman(ctx context.Context, in *schema.ConsultAsHuman, opts ...grpc.CallOption) (*schema.ErrorHuman, error) {
	out := new(schema.ErrorHuman)
	err := c.cc.Invoke(ctx, "/ergocon.Ergo/ConsultErrorAsHuman", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *ergoClient) ConsultErrorAsDeveloper(ctx context.Context, in *schema.ConsultAsDev, opts ...grpc.CallOption) (*schema.ErrorDev, error) {
	out := new(schema.ErrorDev)
	err := c.cc.Invoke(ctx, "/ergocon.Ergo/ConsultErrorAsDeveloper", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *ergoClient) MemorizeNewMessages(ctx context.Context, in *schema.NewMessageParams, opts ...grpc.CallOption) (*schema.UserMessages, error) {
	out := new(schema.UserMessages)
	err := c.cc.Invoke(ctx, "/ergocon.Ergo/MemorizeNewMessages", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *ergoClient) ReceiveFeedbackOfUser(ctx context.Context, in *schema.NewFeedBack, opts ...grpc.CallOption) (*schema.Feedback, error) {
	out := new(schema.Feedback)
	err := c.cc.Invoke(ctx, "/ergocon.Ergo/ReceiveFeedbackOfUser", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// ErgoServer is the server API for Ergo service.
type ErgoServer interface {
	RegisterNewError(context.Context, *schema.ErrorSeed) (*schema.ErrorInstance, error)
	RegisterFullError(context.Context, *schema.FullErrorSeed) (*schema.ErrorInstance, error)
	ConsultErrorAsHuman(context.Context, *schema.ConsultAsHuman) (*schema.ErrorHuman, error)
	ConsultErrorAsDeveloper(context.Context, *schema.ConsultAsDev) (*schema.ErrorDev, error)
	// Save new messages
	MemorizeNewMessages(context.Context, *schema.NewMessageParams) (*schema.UserMessages, error)
	// Save new feedback
	ReceiveFeedbackOfUser(context.Context, *schema.NewFeedBack) (*schema.Feedback, error)
}

func RegisterErgoServer(s *grpc.Server, srv ErgoServer) {
	s.RegisterService(&_Ergo_serviceDesc, srv)
}

func _Ergo_RegisterNewError_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(schema.ErrorSeed)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ErgoServer).RegisterNewError(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/ergocon.Ergo/RegisterNewError",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ErgoServer).RegisterNewError(ctx, req.(*schema.ErrorSeed))
	}
	return interceptor(ctx, in, info, handler)
}

func _Ergo_RegisterFullError_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(schema.FullErrorSeed)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ErgoServer).RegisterFullError(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/ergocon.Ergo/RegisterFullError",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ErgoServer).RegisterFullError(ctx, req.(*schema.FullErrorSeed))
	}
	return interceptor(ctx, in, info, handler)
}

func _Ergo_ConsultErrorAsHuman_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(schema.ConsultAsHuman)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ErgoServer).ConsultErrorAsHuman(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/ergocon.Ergo/ConsultErrorAsHuman",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ErgoServer).ConsultErrorAsHuman(ctx, req.(*schema.ConsultAsHuman))
	}
	return interceptor(ctx, in, info, handler)
}

func _Ergo_ConsultErrorAsDeveloper_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(schema.ConsultAsDev)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ErgoServer).ConsultErrorAsDeveloper(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/ergocon.Ergo/ConsultErrorAsDeveloper",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ErgoServer).ConsultErrorAsDeveloper(ctx, req.(*schema.ConsultAsDev))
	}
	return interceptor(ctx, in, info, handler)
}

func _Ergo_MemorizeNewMessages_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(schema.NewMessageParams)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ErgoServer).MemorizeNewMessages(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/ergocon.Ergo/MemorizeNewMessages",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ErgoServer).MemorizeNewMessages(ctx, req.(*schema.NewMessageParams))
	}
	return interceptor(ctx, in, info, handler)
}

func _Ergo_ReceiveFeedbackOfUser_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(schema.NewFeedBack)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ErgoServer).ReceiveFeedbackOfUser(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/ergocon.Ergo/ReceiveFeedbackOfUser",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ErgoServer).ReceiveFeedbackOfUser(ctx, req.(*schema.NewFeedBack))
	}
	return interceptor(ctx, in, info, handler)
}

var _Ergo_serviceDesc = grpc.ServiceDesc{
	ServiceName: "ergocon.Ergo",
	HandlerType: (*ErgoServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "RegisterNewError",
			Handler:    _Ergo_RegisterNewError_Handler,
		},
		{
			MethodName: "RegisterFullError",
			Handler:    _Ergo_RegisterFullError_Handler,
		},
		{
			MethodName: "ConsultErrorAsHuman",
			Handler:    _Ergo_ConsultErrorAsHuman_Handler,
		},
		{
			MethodName: "ConsultErrorAsDeveloper",
			Handler:    _Ergo_ConsultErrorAsDeveloper_Handler,
		},
		{
			MethodName: "MemorizeNewMessages",
			Handler:    _Ergo_MemorizeNewMessages_Handler,
		},
		{
			MethodName: "ReceiveFeedbackOfUser",
			Handler:    _Ergo_ReceiveFeedbackOfUser_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "service/ergo.proto",
}
