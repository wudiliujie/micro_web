// Code generated by protoc-gen-go. DO NOT EDIT.
// source: shortmsg.proto

/*
Package yxl_micro_srv_shortmsg is a generated protocol buffer package.

It is generated from these files:
	shortmsg.proto

It has these top-level messages:
	GetRegisterCodeReq
	GetRegisterCodeRep
	SendRegisterCodeReq
	SendRegisterCodeRep
*/
package shortmsg

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

import (
	client "github.com/micro/go-micro/client"
	server "github.com/micro/go-micro/server"
	context "golang.org/x/net/context"
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

type GetRegisterCodeReq struct {
	Phone int64 `protobuf:"varint,1,opt,name=phone" json:"phone,omitempty"`
}

func (m *GetRegisterCodeReq) Reset()                    { *m = GetRegisterCodeReq{} }
func (m *GetRegisterCodeReq) String() string            { return proto.CompactTextString(m) }
func (*GetRegisterCodeReq) ProtoMessage()               {}
func (*GetRegisterCodeReq) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

func (m *GetRegisterCodeReq) GetPhone() int64 {
	if m != nil {
		return m.Phone
	}
	return 0
}

type GetRegisterCodeRep struct {
	Tag  int32 `protobuf:"varint,1,opt,name=tag" json:"tag,omitempty"`
	Code int32 `protobuf:"varint,2,opt,name=code" json:"code,omitempty"`
}

func (m *GetRegisterCodeRep) Reset()                    { *m = GetRegisterCodeRep{} }
func (m *GetRegisterCodeRep) String() string            { return proto.CompactTextString(m) }
func (*GetRegisterCodeRep) ProtoMessage()               {}
func (*GetRegisterCodeRep) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{1} }

func (m *GetRegisterCodeRep) GetTag() int32 {
	if m != nil {
		return m.Tag
	}
	return 0
}

func (m *GetRegisterCodeRep) GetCode() int32 {
	if m != nil {
		return m.Code
	}
	return 0
}

type SendRegisterCodeReq struct {
	Phone int64 `protobuf:"varint,1,opt,name=phone" json:"phone,omitempty"`
}

func (m *SendRegisterCodeReq) Reset()                    { *m = SendRegisterCodeReq{} }
func (m *SendRegisterCodeReq) String() string            { return proto.CompactTextString(m) }
func (*SendRegisterCodeReq) ProtoMessage()               {}
func (*SendRegisterCodeReq) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{2} }

func (m *SendRegisterCodeReq) GetPhone() int64 {
	if m != nil {
		return m.Phone
	}
	return 0
}

type SendRegisterCodeRep struct {
	Tag int32 `protobuf:"varint,1,opt,name=tag" json:"tag,omitempty"`
}

func (m *SendRegisterCodeRep) Reset()                    { *m = SendRegisterCodeRep{} }
func (m *SendRegisterCodeRep) String() string            { return proto.CompactTextString(m) }
func (*SendRegisterCodeRep) ProtoMessage()               {}
func (*SendRegisterCodeRep) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{3} }

func (m *SendRegisterCodeRep) GetTag() int32 {
	if m != nil {
		return m.Tag
	}
	return 0
}

func init() {
	proto.RegisterType((*GetRegisterCodeReq)(nil), "yxl.micro.srv.shortmsg.GetRegisterCodeReq")
	proto.RegisterType((*GetRegisterCodeRep)(nil), "yxl.micro.srv.shortmsg.GetRegisterCodeRep")
	proto.RegisterType((*SendRegisterCodeReq)(nil), "yxl.micro.srv.shortmsg.SendRegisterCodeReq")
	proto.RegisterType((*SendRegisterCodeRep)(nil), "yxl.micro.srv.shortmsg.SendRegisterCodeRep")
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ client.Option
var _ server.Option

// Client API for Shortmsg service

type ShortmsgClient interface {
	// 获取注册验证码
	GetRegisterCode(ctx context.Context, in *GetRegisterCodeReq, opts ...client.CallOption) (*GetRegisterCodeRep, error)
	// 发送验证码
	SendRegisterCode(ctx context.Context, in *SendRegisterCodeReq, opts ...client.CallOption) (*SendRegisterCodeRep, error)
}

type shortmsgClient struct {
	c           client.Client
	serviceName string
}

func NewShortmsgClient(serviceName string, c client.Client) ShortmsgClient {
	if c == nil {
		c = client.NewClient()
	}
	if len(serviceName) == 0 {
		serviceName = "yxl.micro.srv.shortmsg"
	}
	return &shortmsgClient{
		c:           c,
		serviceName: serviceName,
	}
}

func (c *shortmsgClient) GetRegisterCode(ctx context.Context, in *GetRegisterCodeReq, opts ...client.CallOption) (*GetRegisterCodeRep, error) {
	req := c.c.NewRequest(c.serviceName, "Shortmsg.GetRegisterCode", in)
	out := new(GetRegisterCodeRep)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *shortmsgClient) SendRegisterCode(ctx context.Context, in *SendRegisterCodeReq, opts ...client.CallOption) (*SendRegisterCodeRep, error) {
	req := c.c.NewRequest(c.serviceName, "Shortmsg.SendRegisterCode", in)
	out := new(SendRegisterCodeRep)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for Shortmsg service

type ShortmsgHandler interface {
	// 获取注册验证码
	GetRegisterCode(context.Context, *GetRegisterCodeReq, *GetRegisterCodeRep) error
	// 发送验证码
	SendRegisterCode(context.Context, *SendRegisterCodeReq, *SendRegisterCodeRep) error
}

func RegisterShortmsgHandler(s server.Server, hdlr ShortmsgHandler, opts ...server.HandlerOption) {
	s.Handle(s.NewHandler(&Shortmsg{hdlr}, opts...))
}

type Shortmsg struct {
	ShortmsgHandler
}

func (h *Shortmsg) GetRegisterCode(ctx context.Context, in *GetRegisterCodeReq, out *GetRegisterCodeRep) error {
	return h.ShortmsgHandler.GetRegisterCode(ctx, in, out)
}

func (h *Shortmsg) SendRegisterCode(ctx context.Context, in *SendRegisterCodeReq, out *SendRegisterCodeRep) error {
	return h.ShortmsgHandler.SendRegisterCode(ctx, in, out)
}

func init() { proto.RegisterFile("shortmsg.proto", fileDescriptor0) }

var fileDescriptor0 = []byte{
	// 200 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0xe2, 0x2b, 0xce, 0xc8, 0x2f,
	0x2a, 0xc9, 0x2d, 0x4e, 0xd7, 0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17, 0x12, 0xab, 0xac, 0xc8, 0xd1,
	0xcb, 0xcd, 0x4c, 0x2e, 0xca, 0xd7, 0x2b, 0x2e, 0x2a, 0xd3, 0x83, 0xc9, 0x2a, 0x69, 0x71, 0x09,
	0xb9, 0xa7, 0x96, 0x04, 0xa5, 0xa6, 0x67, 0x16, 0x97, 0xa4, 0x16, 0x39, 0xe7, 0xa7, 0xa4, 0x06,
	0xa5, 0x16, 0x0a, 0x89, 0x70, 0xb1, 0x16, 0x64, 0xe4, 0xe7, 0xa5, 0x4a, 0x30, 0x2a, 0x30, 0x6a,
	0x30, 0x07, 0x41, 0x38, 0x4a, 0x56, 0x58, 0xd4, 0x16, 0x08, 0x09, 0x70, 0x31, 0x97, 0x24, 0xa6,
	0x83, 0x55, 0xb2, 0x06, 0x81, 0x98, 0x42, 0x42, 0x5c, 0x2c, 0xc9, 0xf9, 0x29, 0xa9, 0x12, 0x4c,
	0x60, 0x21, 0x30, 0x5b, 0x49, 0x9b, 0x4b, 0x38, 0x38, 0x35, 0x2f, 0x85, 0x38, 0x8b, 0xd4, 0xb1,
	0x29, 0xc6, 0x62, 0x93, 0xd1, 0x73, 0x46, 0x2e, 0x8e, 0x60, 0xa8, 0x57, 0x84, 0xb2, 0xb9, 0xf8,
	0xd1, 0x9c, 0x27, 0xa4, 0xa5, 0x87, 0xdd, 0xdb, 0x7a, 0x98, 0x7e, 0x96, 0x22, 0x5e, 0x6d, 0x81,
	0x12, 0x83, 0x50, 0x1e, 0x97, 0x00, 0xba, 0x13, 0x85, 0xb4, 0x71, 0x99, 0x80, 0xc5, 0xe7, 0x52,
	0x24, 0x28, 0x2e, 0x50, 0x62, 0x48, 0x62, 0x03, 0x47, 0xa3, 0x31, 0x20, 0x00, 0x00, 0xff, 0xff,
	0x28, 0x94, 0xe1, 0xf9, 0xd8, 0x01, 0x00, 0x00,
}
