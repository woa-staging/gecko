// Code generated by protoc-gen-go. DO NOT EDIT.
// source: greadcloser.proto

package greadcloserproto

import (
	context "context"
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
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
const _ = proto.ProtoPackageIsVersion3 // please upgrade the proto package

type ReadRequest struct {
	Length               int32    `protobuf:"varint,1,opt,name=length,proto3" json:"length,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ReadRequest) Reset()         { *m = ReadRequest{} }
func (m *ReadRequest) String() string { return proto.CompactTextString(m) }
func (*ReadRequest) ProtoMessage()    {}
func (*ReadRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_21bd394eba12f5e9, []int{0}
}

func (m *ReadRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ReadRequest.Unmarshal(m, b)
}
func (m *ReadRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ReadRequest.Marshal(b, m, deterministic)
}
func (m *ReadRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ReadRequest.Merge(m, src)
}
func (m *ReadRequest) XXX_Size() int {
	return xxx_messageInfo_ReadRequest.Size(m)
}
func (m *ReadRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_ReadRequest.DiscardUnknown(m)
}

var xxx_messageInfo_ReadRequest proto.InternalMessageInfo

func (m *ReadRequest) GetLength() int32 {
	if m != nil {
		return m.Length
	}
	return 0
}

type ReadResponse struct {
	Read                 []byte   `protobuf:"bytes,1,opt,name=read,proto3" json:"read,omitempty"`
	Error                string   `protobuf:"bytes,2,opt,name=error,proto3" json:"error,omitempty"`
	Errored              bool     `protobuf:"varint,3,opt,name=errored,proto3" json:"errored,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ReadResponse) Reset()         { *m = ReadResponse{} }
func (m *ReadResponse) String() string { return proto.CompactTextString(m) }
func (*ReadResponse) ProtoMessage()    {}
func (*ReadResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_21bd394eba12f5e9, []int{1}
}

func (m *ReadResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ReadResponse.Unmarshal(m, b)
}
func (m *ReadResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ReadResponse.Marshal(b, m, deterministic)
}
func (m *ReadResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ReadResponse.Merge(m, src)
}
func (m *ReadResponse) XXX_Size() int {
	return xxx_messageInfo_ReadResponse.Size(m)
}
func (m *ReadResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_ReadResponse.DiscardUnknown(m)
}

var xxx_messageInfo_ReadResponse proto.InternalMessageInfo

func (m *ReadResponse) GetRead() []byte {
	if m != nil {
		return m.Read
	}
	return nil
}

func (m *ReadResponse) GetError() string {
	if m != nil {
		return m.Error
	}
	return ""
}

func (m *ReadResponse) GetErrored() bool {
	if m != nil {
		return m.Errored
	}
	return false
}

type CloseRequest struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *CloseRequest) Reset()         { *m = CloseRequest{} }
func (m *CloseRequest) String() string { return proto.CompactTextString(m) }
func (*CloseRequest) ProtoMessage()    {}
func (*CloseRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_21bd394eba12f5e9, []int{2}
}

func (m *CloseRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CloseRequest.Unmarshal(m, b)
}
func (m *CloseRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CloseRequest.Marshal(b, m, deterministic)
}
func (m *CloseRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CloseRequest.Merge(m, src)
}
func (m *CloseRequest) XXX_Size() int {
	return xxx_messageInfo_CloseRequest.Size(m)
}
func (m *CloseRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_CloseRequest.DiscardUnknown(m)
}

var xxx_messageInfo_CloseRequest proto.InternalMessageInfo

type CloseResponse struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *CloseResponse) Reset()         { *m = CloseResponse{} }
func (m *CloseResponse) String() string { return proto.CompactTextString(m) }
func (*CloseResponse) ProtoMessage()    {}
func (*CloseResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_21bd394eba12f5e9, []int{3}
}

func (m *CloseResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CloseResponse.Unmarshal(m, b)
}
func (m *CloseResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CloseResponse.Marshal(b, m, deterministic)
}
func (m *CloseResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CloseResponse.Merge(m, src)
}
func (m *CloseResponse) XXX_Size() int {
	return xxx_messageInfo_CloseResponse.Size(m)
}
func (m *CloseResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_CloseResponse.DiscardUnknown(m)
}

var xxx_messageInfo_CloseResponse proto.InternalMessageInfo

func init() {
	proto.RegisterType((*ReadRequest)(nil), "greadcloserproto.ReadRequest")
	proto.RegisterType((*ReadResponse)(nil), "greadcloserproto.ReadResponse")
	proto.RegisterType((*CloseRequest)(nil), "greadcloserproto.CloseRequest")
	proto.RegisterType((*CloseResponse)(nil), "greadcloserproto.CloseResponse")
}

func init() { proto.RegisterFile("greadcloser.proto", fileDescriptor_21bd394eba12f5e9) }

var fileDescriptor_21bd394eba12f5e9 = []byte{
	// 208 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x74, 0x4f, 0xcb, 0x8a, 0xc2, 0x40,
	0x10, 0x64, 0x76, 0x93, 0xec, 0x6e, 0x6f, 0xf6, 0xd5, 0x2c, 0x32, 0x04, 0xd4, 0x30, 0x20, 0xe4,
	0x94, 0x83, 0x7e, 0x82, 0x08, 0x9e, 0xe7, 0x0f, 0xa2, 0x69, 0xe2, 0x21, 0x64, 0xe2, 0x4c, 0xfc,
	0x18, 0xff, 0x56, 0xe6, 0x21, 0x04, 0x31, 0xb7, 0xaa, 0xea, 0xa2, 0xaa, 0x1a, 0xfe, 0x1a, 0x4d,
	0x55, 0x7d, 0x6c, 0x95, 0x21, 0x5d, 0xf6, 0x5a, 0x0d, 0x0a, 0x7f, 0x47, 0x92, 0x53, 0xc4, 0x0a,
	0x3e, 0x25, 0x55, 0xb5, 0xa4, 0xf3, 0x85, 0xcc, 0x80, 0x33, 0x48, 0x5a, 0xea, 0x9a, 0xe1, 0xc4,
	0x59, 0xce, 0x8a, 0x58, 0x06, 0x26, 0x24, 0xa4, 0xde, 0x66, 0x7a, 0xd5, 0x19, 0x42, 0x84, 0xc8,
	0x26, 0x39, 0x57, 0x2a, 0x1d, 0xc6, 0x7f, 0x88, 0x49, 0x6b, 0xa5, 0xf9, 0x4b, 0xce, 0x8a, 0x0f,
	0xe9, 0x09, 0x72, 0x78, 0x73, 0x80, 0x6a, 0xfe, 0x9a, 0xb3, 0xe2, 0x5d, 0xde, 0xa9, 0xf8, 0x86,
	0x74, 0x6b, 0x97, 0x84, 0x6e, 0xf1, 0x03, 0x5f, 0x81, 0xfb, 0x92, 0xf5, 0x95, 0x41, 0x62, 0x5b,
	0x49, 0xe3, 0x0e, 0x22, 0x8b, 0x70, 0x5e, 0x3e, 0x7e, 0x50, 0x8e, 0xe6, 0x67, 0x8b, 0xa9, 0x73,
	0x98, 0xbd, 0x87, 0xd8, 0x55, 0xe0, 0x13, 0xe3, 0x78, 0x4b, 0xb6, 0x9c, 0xbc, 0xfb, 0xa4, 0x43,
	0xe2, 0xc4, 0xcd, 0x2d, 0x00, 0x00, 0xff, 0xff, 0x8e, 0xaf, 0xb3, 0x35, 0x65, 0x01, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConnInterface

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion6

// ReaderClient is the client API for Reader service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type ReaderClient interface {
	Read(ctx context.Context, in *ReadRequest, opts ...grpc.CallOption) (*ReadResponse, error)
	Close(ctx context.Context, in *CloseRequest, opts ...grpc.CallOption) (*CloseResponse, error)
}

type readerClient struct {
	cc grpc.ClientConnInterface
}

func NewReaderClient(cc grpc.ClientConnInterface) ReaderClient {
	return &readerClient{cc}
}

func (c *readerClient) Read(ctx context.Context, in *ReadRequest, opts ...grpc.CallOption) (*ReadResponse, error) {
	out := new(ReadResponse)
	err := c.cc.Invoke(ctx, "/greadcloserproto.Reader/Read", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *readerClient) Close(ctx context.Context, in *CloseRequest, opts ...grpc.CallOption) (*CloseResponse, error) {
	out := new(CloseResponse)
	err := c.cc.Invoke(ctx, "/greadcloserproto.Reader/Close", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// ReaderServer is the server API for Reader service.
type ReaderServer interface {
	Read(context.Context, *ReadRequest) (*ReadResponse, error)
	Close(context.Context, *CloseRequest) (*CloseResponse, error)
}

// UnimplementedReaderServer can be embedded to have forward compatible implementations.
type UnimplementedReaderServer struct {
}

func (*UnimplementedReaderServer) Read(ctx context.Context, req *ReadRequest) (*ReadResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Read not implemented")
}
func (*UnimplementedReaderServer) Close(ctx context.Context, req *CloseRequest) (*CloseResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Close not implemented")
}

func RegisterReaderServer(s *grpc.Server, srv ReaderServer) {
	s.RegisterService(&_Reader_serviceDesc, srv)
}

func _Reader_Read_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ReadRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ReaderServer).Read(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/greadcloserproto.Reader/Read",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ReaderServer).Read(ctx, req.(*ReadRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Reader_Close_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CloseRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ReaderServer).Close(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/greadcloserproto.Reader/Close",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ReaderServer).Close(ctx, req.(*CloseRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _Reader_serviceDesc = grpc.ServiceDesc{
	ServiceName: "greadcloserproto.Reader",
	HandlerType: (*ReaderServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Read",
			Handler:    _Reader_Read_Handler,
		},
		{
			MethodName: "Close",
			Handler:    _Reader_Close_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "greadcloser.proto",
}