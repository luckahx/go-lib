// Code generated by protoc-gen-go. DO NOT EDIT.
// source: ping.proto

package ping

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"
import protocerr "github.com/luckahx/go-lib/cerr/proto/protocerr"

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

type PingReq struct {
	Data                 string   `protobuf:"bytes,1,opt,name=Data,proto3" json:"Data,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *PingReq) Reset()         { *m = PingReq{} }
func (m *PingReq) String() string { return proto.CompactTextString(m) }
func (*PingReq) ProtoMessage()    {}
func (*PingReq) Descriptor() ([]byte, []int) {
	return fileDescriptor_ping_422bf9a8afae311e, []int{0}
}
func (m *PingReq) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_PingReq.Unmarshal(m, b)
}
func (m *PingReq) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_PingReq.Marshal(b, m, deterministic)
}
func (dst *PingReq) XXX_Merge(src proto.Message) {
	xxx_messageInfo_PingReq.Merge(dst, src)
}
func (m *PingReq) XXX_Size() int {
	return xxx_messageInfo_PingReq.Size(m)
}
func (m *PingReq) XXX_DiscardUnknown() {
	xxx_messageInfo_PingReq.DiscardUnknown(m)
}

var xxx_messageInfo_PingReq proto.InternalMessageInfo

func (m *PingReq) GetData() string {
	if m != nil {
		return m.Data
	}
	return ""
}

type PingResponse struct {
	Err                  *protocerr.CError `protobuf:"bytes,1,opt,name=err,proto3" json:"err,omitempty"`
	Data                 string            `protobuf:"bytes,2,opt,name=Data,proto3" json:"Data,omitempty"`
	XXX_NoUnkeyedLiteral struct{}          `json:"-"`
	XXX_unrecognized     []byte            `json:"-"`
	XXX_sizecache        int32             `json:"-"`
}

func (m *PingResponse) Reset()         { *m = PingResponse{} }
func (m *PingResponse) String() string { return proto.CompactTextString(m) }
func (*PingResponse) ProtoMessage()    {}
func (*PingResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_ping_422bf9a8afae311e, []int{1}
}
func (m *PingResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_PingResponse.Unmarshal(m, b)
}
func (m *PingResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_PingResponse.Marshal(b, m, deterministic)
}
func (dst *PingResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_PingResponse.Merge(dst, src)
}
func (m *PingResponse) XXX_Size() int {
	return xxx_messageInfo_PingResponse.Size(m)
}
func (m *PingResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_PingResponse.DiscardUnknown(m)
}

var xxx_messageInfo_PingResponse proto.InternalMessageInfo

func (m *PingResponse) GetErr() *protocerr.CError {
	if m != nil {
		return m.Err
	}
	return nil
}

func (m *PingResponse) GetData() string {
	if m != nil {
		return m.Data
	}
	return ""
}

func init() {
	proto.RegisterType((*PingReq)(nil), "ping.PingReq")
	proto.RegisterType((*PingResponse)(nil), "ping.PingResponse")
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// PingClient is the client API for Ping service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type PingClient interface {
	Ping(ctx context.Context, in *PingReq, opts ...grpc.CallOption) (*PingResponse, error)
}

type pingClient struct {
	cc *grpc.ClientConn
}

func NewPingClient(cc *grpc.ClientConn) PingClient {
	return &pingClient{cc}
}

func (c *pingClient) Ping(ctx context.Context, in *PingReq, opts ...grpc.CallOption) (*PingResponse, error) {
	out := new(PingResponse)
	err := c.cc.Invoke(ctx, "/ping.Ping/Ping", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// PingServer is the server API for Ping service.
type PingServer interface {
	Ping(context.Context, *PingReq) (*PingResponse, error)
}

func RegisterPingServer(s *grpc.Server, srv PingServer) {
	s.RegisterService(&_Ping_serviceDesc, srv)
}

func _Ping_Ping_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(PingReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PingServer).Ping(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/ping.Ping/Ping",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PingServer).Ping(ctx, req.(*PingReq))
	}
	return interceptor(ctx, in, info, handler)
}

var _Ping_serviceDesc = grpc.ServiceDesc{
	ServiceName: "ping.Ping",
	HandlerType: (*PingServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Ping",
			Handler:    _Ping_Ping_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "ping.proto",
}

func init() { proto.RegisterFile("ping.proto", fileDescriptor_ping_422bf9a8afae311e) }

var fileDescriptor_ping_422bf9a8afae311e = []byte{
	// 180 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0xe2, 0x2a, 0xc8, 0xcc, 0x4b,
	0xd7, 0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17, 0x62, 0x01, 0xb1, 0xa5, 0xec, 0xd2, 0x33, 0x4b, 0x32,
	0x4a, 0x93, 0xf4, 0x92, 0xf3, 0x73, 0xf5, 0x73, 0x4a, 0x93, 0xb3, 0x13, 0x33, 0x2a, 0xf4, 0xd3,
	0xf3, 0x75, 0x73, 0x32, 0x93, 0xf4, 0x93, 0x53, 0x8b, 0x8a, 0xf4, 0xc1, 0x6a, 0x21, 0x24, 0x82,
	0x0f, 0x62, 0x41, 0x4c, 0x51, 0x92, 0xe5, 0x62, 0x0f, 0xc8, 0xcc, 0x4b, 0x0f, 0x4a, 0x2d, 0x14,
	0x12, 0xe2, 0x62, 0x71, 0x49, 0x2c, 0x49, 0x94, 0x60, 0x54, 0x60, 0xd4, 0xe0, 0x0c, 0x02, 0xb3,
	0x95, 0xdc, 0xb9, 0x78, 0x20, 0xd2, 0xc5, 0x05, 0xf9, 0x79, 0xc5, 0xa9, 0x42, 0xca, 0x5c, 0xcc,
	0xa9, 0x45, 0x45, 0x60, 0x25, 0xdc, 0x46, 0x82, 0x7a, 0x08, 0xd3, 0x9c, 0x5d, 0x8b, 0x8a, 0xf2,
	0x8b, 0x82, 0x40, 0xb2, 0x70, 0x83, 0x98, 0x10, 0x06, 0x19, 0x19, 0x73, 0xb1, 0x80, 0x0c, 0x12,
	0xd2, 0x86, 0xd2, 0xbc, 0x7a, 0x60, 0xaf, 0x40, 0xed, 0x96, 0x12, 0x42, 0xe6, 0x42, 0xec, 0x52,
	0x62, 0x48, 0x62, 0x03, 0x9b, 0x6f, 0x0c, 0x08, 0x00, 0x00, 0xff, 0xff, 0xd3, 0xf0, 0xfe, 0x3e,
	0xf7, 0x00, 0x00, 0x00,
}
