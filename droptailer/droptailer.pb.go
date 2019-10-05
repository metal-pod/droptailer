// Code generated by protoc-gen-go. DO NOT EDIT.
// source: droptailer.proto

package droptailer

import (
	context "context"
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	timestamp "github.com/golang/protobuf/ptypes/timestamp"
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

type Drop struct {
	Timestamp            *timestamp.Timestamp `protobuf:"bytes,1,opt,name=timestamp,proto3" json:"timestamp,omitempty"`
	Fields               map[string]string    `protobuf:"bytes,2,rep,name=fields,proto3" json:"fields,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
	XXX_NoUnkeyedLiteral struct{}             `json:"-"`
	XXX_unrecognized     []byte               `json:"-"`
	XXX_sizecache        int32                `json:"-"`
}

func (m *Drop) Reset()         { *m = Drop{} }
func (m *Drop) String() string { return proto.CompactTextString(m) }
func (*Drop) ProtoMessage()    {}
func (*Drop) Descriptor() ([]byte, []int) {
	return fileDescriptor_5301b1683b20a869, []int{0}
}

func (m *Drop) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Drop.Unmarshal(m, b)
}
func (m *Drop) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Drop.Marshal(b, m, deterministic)
}
func (m *Drop) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Drop.Merge(m, src)
}
func (m *Drop) XXX_Size() int {
	return xxx_messageInfo_Drop.Size(m)
}
func (m *Drop) XXX_DiscardUnknown() {
	xxx_messageInfo_Drop.DiscardUnknown(m)
}

var xxx_messageInfo_Drop proto.InternalMessageInfo

func (m *Drop) GetTimestamp() *timestamp.Timestamp {
	if m != nil {
		return m.Timestamp
	}
	return nil
}

func (m *Drop) GetFields() map[string]string {
	if m != nil {
		return m.Fields
	}
	return nil
}

type Void struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Void) Reset()         { *m = Void{} }
func (m *Void) String() string { return proto.CompactTextString(m) }
func (*Void) ProtoMessage()    {}
func (*Void) Descriptor() ([]byte, []int) {
	return fileDescriptor_5301b1683b20a869, []int{1}
}

func (m *Void) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Void.Unmarshal(m, b)
}
func (m *Void) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Void.Marshal(b, m, deterministic)
}
func (m *Void) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Void.Merge(m, src)
}
func (m *Void) XXX_Size() int {
	return xxx_messageInfo_Void.Size(m)
}
func (m *Void) XXX_DiscardUnknown() {
	xxx_messageInfo_Void.DiscardUnknown(m)
}

var xxx_messageInfo_Void proto.InternalMessageInfo

func init() {
	proto.RegisterType((*Drop)(nil), "droptailer.Drop")
	proto.RegisterMapType((map[string]string)(nil), "droptailer.Drop.FieldsEntry")
	proto.RegisterType((*Void)(nil), "droptailer.Void")
}

func init() { proto.RegisterFile("droptailer.proto", fileDescriptor_5301b1683b20a869) }

var fileDescriptor_5301b1683b20a869 = []byte{
	// 220 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0x12, 0x48, 0x29, 0xca, 0x2f,
	0x28, 0x49, 0xcc, 0xcc, 0x49, 0x2d, 0xd2, 0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17, 0xe2, 0x42, 0x88,
	0x48, 0xc9, 0xa7, 0xe7, 0xe7, 0xa7, 0xe7, 0xa4, 0xea, 0x83, 0x65, 0x92, 0x4a, 0xd3, 0xf4, 0x4b,
	0x32, 0x73, 0x53, 0x8b, 0x4b, 0x12, 0x73, 0x0b, 0x20, 0x8a, 0x95, 0x36, 0x32, 0x72, 0xb1, 0xb8,
	0x14, 0xe5, 0x17, 0x08, 0x59, 0x70, 0x71, 0xc2, 0xe5, 0x24, 0x18, 0x15, 0x18, 0x35, 0xb8, 0x8d,
	0xa4, 0xf4, 0x20, 0xba, 0xf5, 0x60, 0xba, 0xf5, 0x42, 0x60, 0x2a, 0x82, 0x10, 0x8a, 0x85, 0x4c,
	0xb8, 0xd8, 0xd2, 0x32, 0x53, 0x73, 0x52, 0x8a, 0x25, 0x98, 0x14, 0x98, 0x35, 0xb8, 0x8d, 0x64,
	0xf4, 0x90, 0x9c, 0x04, 0x32, 0x5b, 0xcf, 0x0d, 0x2c, 0xed, 0x9a, 0x57, 0x52, 0x54, 0x19, 0x04,
	0x55, 0x2b, 0x65, 0xc9, 0xc5, 0x8d, 0x24, 0x2c, 0x24, 0xc0, 0xc5, 0x9c, 0x9d, 0x5a, 0x09, 0xb6,
	0x98, 0x33, 0x08, 0xc4, 0x14, 0x12, 0xe1, 0x62, 0x2d, 0x4b, 0xcc, 0x29, 0x4d, 0x95, 0x60, 0x02,
	0x8b, 0x41, 0x38, 0x56, 0x4c, 0x16, 0x8c, 0x4a, 0x6c, 0x5c, 0x2c, 0x61, 0xf9, 0x99, 0x29, 0x46,
	0x56, 0x5c, 0x5c, 0x2e, 0x70, 0x9b, 0x84, 0x74, 0xb8, 0x58, 0x02, 0x4a, 0x8b, 0x33, 0x84, 0x04,
	0xd0, 0xad, 0x97, 0x42, 0x11, 0x01, 0xe9, 0x54, 0x62, 0x48, 0x62, 0x03, 0xfb, 0xc9, 0x18, 0x10,
	0x00, 0x00, 0xff, 0xff, 0x3e, 0x30, 0xd0, 0x30, 0x3f, 0x01, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// DroptailerClient is the client API for Droptailer service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type DroptailerClient interface {
	Push(ctx context.Context, in *Drop, opts ...grpc.CallOption) (*Void, error)
}

type droptailerClient struct {
	cc *grpc.ClientConn
}

func NewDroptailerClient(cc *grpc.ClientConn) DroptailerClient {
	return &droptailerClient{cc}
}

func (c *droptailerClient) Push(ctx context.Context, in *Drop, opts ...grpc.CallOption) (*Void, error) {
	out := new(Void)
	err := c.cc.Invoke(ctx, "/droptailer.Droptailer/Push", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// DroptailerServer is the server API for Droptailer service.
type DroptailerServer interface {
	Push(context.Context, *Drop) (*Void, error)
}

// UnimplementedDroptailerServer can be embedded to have forward compatible implementations.
type UnimplementedDroptailerServer struct {
}

func (*UnimplementedDroptailerServer) Push(ctx context.Context, req *Drop) (*Void, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Push not implemented")
}

func RegisterDroptailerServer(s *grpc.Server, srv DroptailerServer) {
	s.RegisterService(&_Droptailer_serviceDesc, srv)
}

func _Droptailer_Push_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Drop)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DroptailerServer).Push(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/droptailer.Droptailer/Push",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DroptailerServer).Push(ctx, req.(*Drop))
	}
	return interceptor(ctx, in, info, handler)
}

var _Droptailer_serviceDesc = grpc.ServiceDesc{
	ServiceName: "droptailer.Droptailer",
	HandlerType: (*DroptailerServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Push",
			Handler:    _Droptailer_Push_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "droptailer.proto",
}
