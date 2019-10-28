// Code generated by protoc-gen-go. DO NOT EDIT.
// source: profile.proto

package api

import (
	context "context"
	fmt "fmt"
	math "math"

	proto "github.com/golang/protobuf/proto"
	_ "google.golang.org/genproto/googleapis/api/annotations"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion3 // please upgrade the proto package

type StringMessage struct {
	Value                string   `protobuf:"bytes,1,opt,name=value,proto3" json:"value,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *StringMessage) Reset()         { *m = StringMessage{} }
func (m *StringMessage) String() string { return proto.CompactTextString(m) }
func (*StringMessage) ProtoMessage()    {}
func (*StringMessage) Descriptor() ([]byte, []int) {
	return fileDescriptor_744bf7a47b381504, []int{0}
}

func (m *StringMessage) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_StringMessage.Unmarshal(m, b)
}
func (m *StringMessage) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_StringMessage.Marshal(b, m, deterministic)
}
func (m *StringMessage) XXX_Merge(src proto.Message) {
	xxx_messageInfo_StringMessage.Merge(m, src)
}
func (m *StringMessage) XXX_Size() int {
	return xxx_messageInfo_StringMessage.Size(m)
}
func (m *StringMessage) XXX_DiscardUnknown() {
	xxx_messageInfo_StringMessage.DiscardUnknown(m)
}

var xxx_messageInfo_StringMessage proto.InternalMessageInfo

func (m *StringMessage) GetValue() string {
	if m != nil {
		return m.Value
	}
	return ""
}

type SimpleMessage struct {
	Id                   string   `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Num                  int64    `protobuf:"varint,2,opt,name=num,proto3" json:"num,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *SimpleMessage) Reset()         { *m = SimpleMessage{} }
func (m *SimpleMessage) String() string { return proto.CompactTextString(m) }
func (*SimpleMessage) ProtoMessage()    {}
func (*SimpleMessage) Descriptor() ([]byte, []int) {
	return fileDescriptor_744bf7a47b381504, []int{1}
}

func (m *SimpleMessage) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_SimpleMessage.Unmarshal(m, b)
}
func (m *SimpleMessage) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_SimpleMessage.Marshal(b, m, deterministic)
}
func (m *SimpleMessage) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SimpleMessage.Merge(m, src)
}
func (m *SimpleMessage) XXX_Size() int {
	return xxx_messageInfo_SimpleMessage.Size(m)
}
func (m *SimpleMessage) XXX_DiscardUnknown() {
	xxx_messageInfo_SimpleMessage.DiscardUnknown(m)
}

var xxx_messageInfo_SimpleMessage proto.InternalMessageInfo

func (m *SimpleMessage) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

func (m *SimpleMessage) GetNum() int64 {
	if m != nil {
		return m.Num
	}
	return 0
}

func init() {
	proto.RegisterType((*StringMessage)(nil), "api.StringMessage")
	proto.RegisterType((*SimpleMessage)(nil), "api.SimpleMessage")
}

func init() { proto.RegisterFile("profile.proto", fileDescriptor_744bf7a47b381504) }

var fileDescriptor_744bf7a47b381504 = []byte{
	// 264 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0xe2, 0x2d, 0x28, 0xca, 0x4f,
	0xcb, 0xcc, 0x49, 0xd5, 0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17, 0x62, 0x4e, 0x2c, 0xc8, 0x94, 0x92,
	0x49, 0xcf, 0xcf, 0x4f, 0xcf, 0x49, 0xd5, 0x4f, 0x2c, 0xc8, 0xd4, 0x4f, 0xcc, 0xcb, 0xcb, 0x2f,
	0x49, 0x2c, 0xc9, 0xcc, 0xcf, 0x2b, 0x86, 0x28, 0x51, 0x52, 0xe5, 0xe2, 0x0d, 0x2e, 0x29, 0xca,
	0xcc, 0x4b, 0xf7, 0x4d, 0x2d, 0x2e, 0x4e, 0x4c, 0x4f, 0x15, 0x12, 0xe1, 0x62, 0x2d, 0x4b, 0xcc,
	0x29, 0x4d, 0x95, 0x60, 0x54, 0x60, 0xd4, 0xe0, 0x0c, 0x82, 0x70, 0x94, 0x0c, 0xb9, 0x78, 0x83,
	0x33, 0x73, 0x0b, 0x72, 0x52, 0x61, 0xca, 0xf8, 0xb8, 0x98, 0x32, 0x53, 0xa0, 0x6a, 0x98, 0x32,
	0x53, 0x84, 0x04, 0xb8, 0x98, 0xf3, 0x4a, 0x73, 0x25, 0x98, 0x14, 0x18, 0x35, 0x98, 0x83, 0x40,
	0x4c, 0xa3, 0x2b, 0x8c, 0x5c, 0x7c, 0x01, 0x10, 0xe7, 0x04, 0xa7, 0x16, 0x95, 0x65, 0x26, 0xa7,
	0x0a, 0x85, 0x70, 0x71, 0xb8, 0x26, 0x67, 0xe4, 0x07, 0xe4, 0x17, 0x97, 0x08, 0x09, 0xe9, 0x25,
	0x16, 0x64, 0xea, 0xa1, 0xd8, 0x2d, 0x85, 0x45, 0x4c, 0x49, 0xa1, 0xe9, 0xf2, 0x93, 0xc9, 0x4c,
	0x52, 0x4a, 0xa2, 0xfa, 0x65, 0x86, 0xfa, 0xa9, 0x15, 0x89, 0x20, 0x47, 0xe8, 0xa7, 0x26, 0x67,
	0xe4, 0xc7, 0x17, 0xe4, 0x17, 0x97, 0x58, 0x31, 0x6a, 0x09, 0xe5, 0x70, 0xb1, 0x80, 0x4c, 0x85,
	0x99, 0x88, 0xec, 0x4c, 0x29, 0x2c, 0x62, 0x4a, 0x36, 0x60, 0x13, 0xcd, 0x30, 0x4d, 0xd4, 0xaf,
	0xce, 0x4c, 0xa9, 0x8d, 0x92, 0x15, 0x92, 0xc6, 0x2a, 0xa1, 0x5f, 0x9d, 0x57, 0x9a, 0x5b, 0xeb,
	0xc4, 0x1a, 0x05, 0x0a, 0xd5, 0x24, 0x36, 0x70, 0xf0, 0x19, 0x03, 0x02, 0x00, 0x00, 0xff, 0xff,
	0xbd, 0xfb, 0x99, 0xf3, 0x72, 0x01, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// ProfileServiceClient is the client API for ProfileService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type ProfileServiceClient interface {
	EchoPost(ctx context.Context, in *StringMessage, opts ...grpc.CallOption) (*StringMessage, error)
	Echo(ctx context.Context, in *SimpleMessage, opts ...grpc.CallOption) (*SimpleMessage, error)
}

type profileServiceClient struct {
	cc *grpc.ClientConn
}

func NewProfileServiceClient(cc *grpc.ClientConn) ProfileServiceClient {
	return &profileServiceClient{cc}
}

func (c *profileServiceClient) EchoPost(ctx context.Context, in *StringMessage, opts ...grpc.CallOption) (*StringMessage, error) {
	out := new(StringMessage)
	err := c.cc.Invoke(ctx, "/api.ProfileService/EchoPost", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *profileServiceClient) Echo(ctx context.Context, in *SimpleMessage, opts ...grpc.CallOption) (*SimpleMessage, error) {
	out := new(SimpleMessage)
	err := c.cc.Invoke(ctx, "/api.ProfileService/Echo", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// ProfileServiceServer is the server API for ProfileService service.
type ProfileServiceServer interface {
	EchoPost(context.Context, *StringMessage) (*StringMessage, error)
	Echo(context.Context, *SimpleMessage) (*SimpleMessage, error)
}

// UnimplementedProfileServiceServer can be embedded to have forward compatible implementations.
type UnimplementedProfileServiceServer struct {
}

func (*UnimplementedProfileServiceServer) EchoPost(ctx context.Context, req *StringMessage) (*StringMessage, error) {
	return nil, status.Errorf(codes.Unimplemented, "method EchoPost not implemented")
}
func (*UnimplementedProfileServiceServer) Echo(ctx context.Context, req *SimpleMessage) (*SimpleMessage, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Echo not implemented")
}

func RegisterProfileServiceServer(s *grpc.Server, srv ProfileServiceServer) {
	s.RegisterService(&_ProfileService_serviceDesc, srv)
}

func _ProfileService_EchoPost_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(StringMessage)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ProfileServiceServer).EchoPost(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/api.ProfileService/EchoPost",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ProfileServiceServer).EchoPost(ctx, req.(*StringMessage))
	}
	return interceptor(ctx, in, info, handler)
}

func _ProfileService_Echo_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SimpleMessage)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ProfileServiceServer).Echo(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/api.ProfileService/Echo",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ProfileServiceServer).Echo(ctx, req.(*SimpleMessage))
	}
	return interceptor(ctx, in, info, handler)
}

var _ProfileService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "api.ProfileService",
	HandlerType: (*ProfileServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "EchoPost",
			Handler:    _ProfileService_EchoPost_Handler,
		},
		{
			MethodName: "Echo",
			Handler:    _ProfileService_Echo_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "profile.proto",
}
