// Code generated by protoc-gen-go. DO NOT EDIT.
// source: src/proto/calc.proto

package proto

/*
生成的程式在 Golang 中將會屬於 `pb` 套件。
*/

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

// CalcRequest 包含了兩個數字，將會傳送至計算服務並對兩個數字進行計算。
type CalcRequest struct {
	NumberA              int32    `protobuf:"varint,1,opt,name=number_a,json=numberA,proto3" json:"number_a,omitempty"`
	NumberB              int32    `protobuf:"varint,2,opt,name=number_b,json=numberB,proto3" json:"number_b,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *CalcRequest) Reset()         { *m = CalcRequest{} }
func (m *CalcRequest) String() string { return proto.CompactTextString(m) }
func (*CalcRequest) ProtoMessage()    {}
func (*CalcRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_calc_1dba836a4d20b06a, []int{0}
}
func (m *CalcRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CalcRequest.Unmarshal(m, b)
}
func (m *CalcRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CalcRequest.Marshal(b, m, deterministic)
}
func (dst *CalcRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CalcRequest.Merge(dst, src)
}
func (m *CalcRequest) XXX_Size() int {
	return xxx_messageInfo_CalcRequest.Size(m)
}
func (m *CalcRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_CalcRequest.DiscardUnknown(m)
}

var xxx_messageInfo_CalcRequest proto.InternalMessageInfo

func (m *CalcRequest) GetNumberA() int32 {
	if m != nil {
		return m.NumberA
	}
	return 0
}

func (m *CalcRequest) GetNumberB() int32 {
	if m != nil {
		return m.NumberB
	}
	return 0
}

// CalcReply 是計算結果，將會回傳給客戶端。
type CalcReply struct {
	Result               int32    `protobuf:"varint,1,opt,name=result,proto3" json:"result,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *CalcReply) Reset()         { *m = CalcReply{} }
func (m *CalcReply) String() string { return proto.CompactTextString(m) }
func (*CalcReply) ProtoMessage()    {}
func (*CalcReply) Descriptor() ([]byte, []int) {
	return fileDescriptor_calc_1dba836a4d20b06a, []int{1}
}
func (m *CalcReply) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CalcReply.Unmarshal(m, b)
}
func (m *CalcReply) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CalcReply.Marshal(b, m, deterministic)
}
func (dst *CalcReply) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CalcReply.Merge(dst, src)
}
func (m *CalcReply) XXX_Size() int {
	return xxx_messageInfo_CalcReply.Size(m)
}
func (m *CalcReply) XXX_DiscardUnknown() {
	xxx_messageInfo_CalcReply.DiscardUnknown(m)
}

var xxx_messageInfo_CalcReply proto.InternalMessageInfo

func (m *CalcReply) GetResult() int32 {
	if m != nil {
		return m.Result
	}
	return 0
}

func init() {
	proto.RegisterType((*CalcRequest)(nil), "proto.CalcRequest")
	proto.RegisterType((*CalcReply)(nil), "proto.CalcReply")
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// CalculatorClient is the client API for Calculator service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type CalculatorClient interface {
	// Plus 會接收 CalcRequest 資料作加總，最終會回傳 CalcReply。
	Plus(ctx context.Context, in *CalcRequest, opts ...grpc.CallOption) (*CalcReply, error)
}

type calculatorClient struct {
	cc *grpc.ClientConn
}

func NewCalculatorClient(cc *grpc.ClientConn) CalculatorClient {
	return &calculatorClient{cc}
}

func (c *calculatorClient) Plus(ctx context.Context, in *CalcRequest, opts ...grpc.CallOption) (*CalcReply, error) {
	out := new(CalcReply)
	err := c.cc.Invoke(ctx, "/proto.Calculator/Plus", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// CalculatorServer is the server API for Calculator service.
type CalculatorServer interface {
	// Plus 會接收 CalcRequest 資料作加總，最終會回傳 CalcReply。
	Plus(context.Context, *CalcRequest) (*CalcReply, error)
}

func RegisterCalculatorServer(s *grpc.Server, srv CalculatorServer) {
	s.RegisterService(&_Calculator_serviceDesc, srv)
}

func _Calculator_Plus_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CalcRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CalculatorServer).Plus(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/proto.Calculator/Plus",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CalculatorServer).Plus(ctx, req.(*CalcRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _Calculator_serviceDesc = grpc.ServiceDesc{
	ServiceName: "proto.Calculator",
	HandlerType: (*CalculatorServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Plus",
			Handler:    _Calculator_Plus_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "src/proto/calc.proto",
}

func init() { proto.RegisterFile("src/proto/calc.proto", fileDescriptor_calc_1dba836a4d20b06a) }

var fileDescriptor_calc_1dba836a4d20b06a = []byte{
	// 157 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0x12, 0x29, 0x2e, 0x4a, 0xd6,
	0x2f, 0x28, 0xca, 0x2f, 0xc9, 0xd7, 0x4f, 0x4e, 0xcc, 0x49, 0xd6, 0x03, 0x33, 0x85, 0x58, 0xc1,
	0x94, 0x92, 0x33, 0x17, 0xb7, 0x73, 0x62, 0x4e, 0x72, 0x50, 0x6a, 0x61, 0x69, 0x6a, 0x71, 0x89,
	0x90, 0x24, 0x17, 0x47, 0x5e, 0x69, 0x6e, 0x52, 0x6a, 0x51, 0x7c, 0xa2, 0x04, 0xa3, 0x02, 0xa3,
	0x06, 0x6b, 0x10, 0x3b, 0x84, 0xef, 0x88, 0x24, 0x95, 0x24, 0xc1, 0x84, 0x2c, 0xe5, 0xa4, 0xa4,
	0xcc, 0xc5, 0x09, 0x31, 0xa4, 0x20, 0xa7, 0x52, 0x48, 0x8c, 0x8b, 0xad, 0x28, 0xb5, 0xb8, 0x34,
	0xa7, 0x04, 0x6a, 0x00, 0x94, 0x67, 0x64, 0xc3, 0xc5, 0x05, 0x52, 0x54, 0x9a, 0x93, 0x58, 0x92,
	0x5f, 0x24, 0xa4, 0xc7, 0xc5, 0x12, 0x90, 0x53, 0x5a, 0x2c, 0x24, 0x04, 0x71, 0x8e, 0x1e, 0x92,
	0x23, 0xa4, 0x04, 0x50, 0xc4, 0x0a, 0x72, 0x2a, 0x95, 0x18, 0x92, 0xd8, 0xc0, 0x42, 0xc6, 0x80,
	0x00, 0x00, 0x00, 0xff, 0xff, 0x34, 0xba, 0xf6, 0x53, 0xcd, 0x00, 0x00, 0x00,
}
