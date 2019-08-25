// Code generated by protoc-gen-go. DO NOT EDIT.
// source: peekaboo.proto

package pb

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

type PikabuRequest struct {
	Category             string   `protobuf:"bytes,1,opt,name=category,proto3" json:"category,omitempty"`
	Keyword              string   `protobuf:"bytes,10,opt,name=keyword,proto3" json:"keyword,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *PikabuRequest) Reset()         { *m = PikabuRequest{} }
func (m *PikabuRequest) String() string { return proto.CompactTextString(m) }
func (*PikabuRequest) ProtoMessage()    {}
func (*PikabuRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_4e9296e1ca7b7ddb, []int{0}
}

func (m *PikabuRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_PikabuRequest.Unmarshal(m, b)
}
func (m *PikabuRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_PikabuRequest.Marshal(b, m, deterministic)
}
func (m *PikabuRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_PikabuRequest.Merge(m, src)
}
func (m *PikabuRequest) XXX_Size() int {
	return xxx_messageInfo_PikabuRequest.Size(m)
}
func (m *PikabuRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_PikabuRequest.DiscardUnknown(m)
}

var xxx_messageInfo_PikabuRequest proto.InternalMessageInfo

func (m *PikabuRequest) GetCategory() string {
	if m != nil {
		return m.Category
	}
	return ""
}

func (m *PikabuRequest) GetKeyword() string {
	if m != nil {
		return m.Keyword
	}
	return ""
}

type PikabuReply struct {
	Category             string   `protobuf:"bytes,1,opt,name=category,proto3" json:"category,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *PikabuReply) Reset()         { *m = PikabuReply{} }
func (m *PikabuReply) String() string { return proto.CompactTextString(m) }
func (*PikabuReply) ProtoMessage()    {}
func (*PikabuReply) Descriptor() ([]byte, []int) {
	return fileDescriptor_4e9296e1ca7b7ddb, []int{1}
}

func (m *PikabuReply) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_PikabuReply.Unmarshal(m, b)
}
func (m *PikabuReply) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_PikabuReply.Marshal(b, m, deterministic)
}
func (m *PikabuReply) XXX_Merge(src proto.Message) {
	xxx_messageInfo_PikabuReply.Merge(m, src)
}
func (m *PikabuReply) XXX_Size() int {
	return xxx_messageInfo_PikabuReply.Size(m)
}
func (m *PikabuReply) XXX_DiscardUnknown() {
	xxx_messageInfo_PikabuReply.DiscardUnknown(m)
}

var xxx_messageInfo_PikabuReply proto.InternalMessageInfo

func (m *PikabuReply) GetCategory() string {
	if m != nil {
		return m.Category
	}
	return ""
}

func init() {
	proto.RegisterType((*PikabuRequest)(nil), "pb.PikabuRequest")
	proto.RegisterType((*PikabuReply)(nil), "pb.PikabuReply")
}

func init() { proto.RegisterFile("peekaboo.proto", fileDescriptor_4e9296e1ca7b7ddb) }

var fileDescriptor_4e9296e1ca7b7ddb = []byte{
	// 144 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0xe2, 0x2b, 0x48, 0x4d, 0xcd,
	0x4e, 0x4c, 0xca, 0xcf, 0xd7, 0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17, 0x62, 0x2a, 0x48, 0x52, 0x72,
	0xe5, 0xe2, 0x0d, 0xc8, 0xcc, 0x4e, 0x4c, 0x2a, 0x0d, 0x4a, 0x2d, 0x2c, 0x4d, 0x2d, 0x2e, 0x11,
	0x92, 0xe2, 0xe2, 0x48, 0x4e, 0x2c, 0x49, 0x4d, 0xcf, 0x2f, 0xaa, 0x94, 0x60, 0x54, 0x60, 0xd4,
	0xe0, 0x0c, 0x82, 0xf3, 0x85, 0x24, 0xb8, 0xd8, 0xb3, 0x53, 0x2b, 0xcb, 0xf3, 0x8b, 0x52, 0x24,
	0xb8, 0xc0, 0x52, 0x30, 0xae, 0x92, 0x26, 0x17, 0x37, 0xcc, 0x98, 0x82, 0x9c, 0x4a, 0x7c, 0x86,
	0x18, 0x59, 0x70, 0x71, 0x04, 0x40, 0xdd, 0x21, 0xa4, 0xc3, 0xc5, 0x06, 0xd1, 0x26, 0x24, 0xa8,
	0x57, 0x90, 0xa4, 0x87, 0xe2, 0x12, 0x29, 0x7e, 0x64, 0xa1, 0x82, 0x9c, 0xca, 0x24, 0x36, 0xb0,
	0xb3, 0x8d, 0x01, 0x01, 0x00, 0x00, 0xff, 0xff, 0x10, 0x81, 0x6c, 0xd2, 0xc8, 0x00, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// PeekabooClient is the client API for Peekaboo service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type PeekabooClient interface {
	Pikabu(ctx context.Context, in *PikabuRequest, opts ...grpc.CallOption) (*PikabuReply, error)
}

type peekabooClient struct {
	cc *grpc.ClientConn
}

func NewPeekabooClient(cc *grpc.ClientConn) PeekabooClient {
	return &peekabooClient{cc}
}

func (c *peekabooClient) Pikabu(ctx context.Context, in *PikabuRequest, opts ...grpc.CallOption) (*PikabuReply, error) {
	out := new(PikabuReply)
	err := c.cc.Invoke(ctx, "/pb.Peekaboo/Pikabu", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// PeekabooServer is the server API for Peekaboo service.
type PeekabooServer interface {
	Pikabu(context.Context, *PikabuRequest) (*PikabuReply, error)
}

// UnimplementedPeekabooServer can be embedded to have forward compatible implementations.
type UnimplementedPeekabooServer struct {
}

func (*UnimplementedPeekabooServer) Pikabu(ctx context.Context, req *PikabuRequest) (*PikabuReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Pikabu not implemented")
}

func RegisterPeekabooServer(s *grpc.Server, srv PeekabooServer) {
	s.RegisterService(&_Peekaboo_serviceDesc, srv)
}

func _Peekaboo_Pikabu_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(PikabuRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PeekabooServer).Pikabu(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pb.Peekaboo/Pikabu",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PeekabooServer).Pikabu(ctx, req.(*PikabuRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _Peekaboo_serviceDesc = grpc.ServiceDesc{
	ServiceName: "pb.Peekaboo",
	HandlerType: (*PeekabooServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Pikabu",
			Handler:    _Peekaboo_Pikabu_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "peekaboo.proto",
}