// Code generated by protoc-gen-go.
// source: manager.proto
// DO NOT EDIT!

/*
Package rpc is a generated protocol buffer package.

It is generated from these files:
	manager.proto
	router.proto

It has these top-level messages:
	LoginReq
	LoginRes
	SendMsgP2PReq
	SendMsgP2PRes
*/
package rpc

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

import (
	context "golang.org/x/net/context"
	grpc "google.golang.org/grpc"
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

type LoginReq struct {
	UID   int64  `protobuf:"varint,3,opt,name=UID" json:"UID,omitempty"`
	Token string `protobuf:"bytes,4,opt,name=token" json:"token,omitempty"`
}

func (m *LoginReq) Reset()                    { *m = LoginReq{} }
func (m *LoginReq) String() string            { return proto.CompactTextString(m) }
func (*LoginReq) ProtoMessage()               {}
func (*LoginReq) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

type LoginRes struct {
	ErrCode uint32 `protobuf:"varint,3,opt,name=errCode" json:"errCode,omitempty"`
	ErrStr  string `protobuf:"bytes,4,opt,name=errStr" json:"errStr,omitempty"`
}

func (m *LoginRes) Reset()                    { *m = LoginRes{} }
func (m *LoginRes) String() string            { return proto.CompactTextString(m) }
func (*LoginRes) ProtoMessage()               {}
func (*LoginRes) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{1} }

func init() {
	proto.RegisterType((*LoginReq)(nil), "rpc.LoginReq")
	proto.RegisterType((*LoginRes)(nil), "rpc.LoginRes")
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion3

// Client API for ManagerRPC service

type ManagerRPCClient interface {
	Login(ctx context.Context, in *LoginReq, opts ...grpc.CallOption) (*LoginRes, error)
}

type managerRPCClient struct {
	cc *grpc.ClientConn
}

func NewManagerRPCClient(cc *grpc.ClientConn) ManagerRPCClient {
	return &managerRPCClient{cc}
}

func (c *managerRPCClient) Login(ctx context.Context, in *LoginReq, opts ...grpc.CallOption) (*LoginRes, error) {
	out := new(LoginRes)
	err := grpc.Invoke(ctx, "/rpc.ManagerRPC/Login", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for ManagerRPC service

type ManagerRPCServer interface {
	Login(context.Context, *LoginReq) (*LoginRes, error)
}

func RegisterManagerRPCServer(s *grpc.Server, srv ManagerRPCServer) {
	s.RegisterService(&_ManagerRPC_serviceDesc, srv)
}

func _ManagerRPC_Login_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(LoginReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ManagerRPCServer).Login(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/rpc.ManagerRPC/Login",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ManagerRPCServer).Login(ctx, req.(*LoginReq))
	}
	return interceptor(ctx, in, info, handler)
}

var _ManagerRPC_serviceDesc = grpc.ServiceDesc{
	ServiceName: "rpc.ManagerRPC",
	HandlerType: (*ManagerRPCServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Login",
			Handler:    _ManagerRPC_Login_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: fileDescriptor0,
}

func init() { proto.RegisterFile("manager.proto", fileDescriptor0) }

var fileDescriptor0 = []byte{
	// 162 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x09, 0x6e, 0x88, 0x02, 0xff, 0xe2, 0xe2, 0xcd, 0x4d, 0xcc, 0x4b,
	0x4c, 0x4f, 0x2d, 0xd2, 0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17, 0x62, 0x2e, 0x2a, 0x48, 0x56, 0x32,
	0xe2, 0xe2, 0xf0, 0xc9, 0x4f, 0xcf, 0xcc, 0x0b, 0x4a, 0x2d, 0x14, 0x12, 0xe0, 0x62, 0x0e, 0xf5,
	0x74, 0x91, 0x60, 0x56, 0x60, 0xd4, 0x60, 0x0e, 0x02, 0x31, 0x85, 0x44, 0xb8, 0x58, 0x4b, 0xf2,
	0xb3, 0x53, 0xf3, 0x24, 0x58, 0x14, 0x18, 0x35, 0x38, 0x83, 0x20, 0x1c, 0x25, 0x1b, 0xb8, 0x9e,
	0x62, 0x21, 0x09, 0x2e, 0xf6, 0xd4, 0xa2, 0x22, 0xe7, 0xfc, 0x94, 0x54, 0xb0, 0x3e, 0xde, 0x20,
	0x18, 0x57, 0x48, 0x8c, 0x8b, 0x2d, 0xb5, 0xa8, 0x28, 0xb8, 0xa4, 0x08, 0xaa, 0x19, 0xca, 0x33,
	0x32, 0xe5, 0xe2, 0xf2, 0x85, 0xb8, 0x23, 0x28, 0xc0, 0x59, 0x48, 0x9d, 0x8b, 0x15, 0x6c, 0x96,
	0x10, 0xaf, 0x5e, 0x51, 0x41, 0xb2, 0x1e, 0xcc, 0x2d, 0x52, 0x28, 0xdc, 0x62, 0x25, 0x86, 0x24,
	0x36, 0xb0, 0xa3, 0x8d, 0x01, 0x01, 0x00, 0x00, 0xff, 0xff, 0x42, 0xdb, 0xc6, 0x80, 0xc5, 0x00,
	0x00, 0x00,
}