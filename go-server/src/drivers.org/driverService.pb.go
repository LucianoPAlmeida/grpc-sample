// Code generated by protoc-gen-go. DO NOT EDIT.
// source: driverService.proto

package driverService

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

type DriverRequest struct {
	Number               int32    `protobuf:"varint,1,opt,name=number,proto3" json:"number,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *DriverRequest) Reset()         { *m = DriverRequest{} }
func (m *DriverRequest) String() string { return proto.CompactTextString(m) }
func (*DriverRequest) ProtoMessage()    {}
func (*DriverRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_76f40df778c75fe5, []int{0}
}

func (m *DriverRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_DriverRequest.Unmarshal(m, b)
}
func (m *DriverRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_DriverRequest.Marshal(b, m, deterministic)
}
func (m *DriverRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_DriverRequest.Merge(m, src)
}
func (m *DriverRequest) XXX_Size() int {
	return xxx_messageInfo_DriverRequest.Size(m)
}
func (m *DriverRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_DriverRequest.DiscardUnknown(m)
}

var xxx_messageInfo_DriverRequest proto.InternalMessageInfo

func (m *DriverRequest) GetNumber() int32 {
	if m != nil {
		return m.Number
	}
	return 0
}

type DriverResponse struct {
	Name                 string   `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *DriverResponse) Reset()         { *m = DriverResponse{} }
func (m *DriverResponse) String() string { return proto.CompactTextString(m) }
func (*DriverResponse) ProtoMessage()    {}
func (*DriverResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_76f40df778c75fe5, []int{1}
}

func (m *DriverResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_DriverResponse.Unmarshal(m, b)
}
func (m *DriverResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_DriverResponse.Marshal(b, m, deterministic)
}
func (m *DriverResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_DriverResponse.Merge(m, src)
}
func (m *DriverResponse) XXX_Size() int {
	return xxx_messageInfo_DriverResponse.Size(m)
}
func (m *DriverResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_DriverResponse.DiscardUnknown(m)
}

var xxx_messageInfo_DriverResponse proto.InternalMessageInfo

func (m *DriverResponse) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func init() {
	proto.RegisterType((*DriverRequest)(nil), "DriverRequest")
	proto.RegisterType((*DriverResponse)(nil), "DriverResponse")
}

func init() { proto.RegisterFile("driverService.proto", fileDescriptor_76f40df778c75fe5) }

var fileDescriptor_76f40df778c75fe5 = []byte{
	// 136 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0x12, 0x4e, 0x29, 0xca, 0x2c,
	0x4b, 0x2d, 0x0a, 0x4e, 0x2d, 0x2a, 0xcb, 0x4c, 0x4e, 0xd5, 0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x57,
	0x52, 0xe7, 0xe2, 0x75, 0x01, 0x0b, 0x07, 0xa5, 0x16, 0x96, 0xa6, 0x16, 0x97, 0x08, 0x89, 0x71,
	0xb1, 0xe5, 0x95, 0xe6, 0x26, 0xa5, 0x16, 0x49, 0x30, 0x2a, 0x30, 0x6a, 0xb0, 0x06, 0x41, 0x79,
	0x4a, 0x2a, 0x5c, 0x7c, 0x30, 0x85, 0xc5, 0x05, 0xf9, 0x79, 0xc5, 0xa9, 0x42, 0x42, 0x5c, 0x2c,
	0x79, 0x89, 0xb9, 0xa9, 0x60, 0x75, 0x9c, 0x41, 0x60, 0xb6, 0x91, 0x25, 0x17, 0x1b, 0x44, 0x95,
	0x90, 0x3e, 0x17, 0x57, 0x5a, 0x66, 0x5e, 0x0a, 0x94, 0xc7, 0xa7, 0x87, 0x62, 0x8b, 0x14, 0xbf,
	0x1e, 0xaa, 0x61, 0x4a, 0x0c, 0x49, 0x6c, 0x60, 0x07, 0x19, 0x03, 0x02, 0x00, 0x00, 0xff, 0xff,
	0x7b, 0x3a, 0x86, 0x8d, 0xa7, 0x00, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// DriverClient is the client API for Driver service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type DriverClient interface {
	FindDriver(ctx context.Context, in *DriverRequest, opts ...grpc.CallOption) (*DriverResponse, error)
}

type driverClient struct {
	cc *grpc.ClientConn
}

func NewDriverClient(cc *grpc.ClientConn) DriverClient {
	return &driverClient{cc}
}

func (c *driverClient) FindDriver(ctx context.Context, in *DriverRequest, opts ...grpc.CallOption) (*DriverResponse, error) {
	out := new(DriverResponse)
	err := c.cc.Invoke(ctx, "/Driver/findDriver", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// DriverServer is the server API for Driver service.
type DriverServer interface {
	FindDriver(context.Context, *DriverRequest) (*DriverResponse, error)
}

// UnimplementedDriverServer can be embedded to have forward compatible implementations.
type UnimplementedDriverServer struct {
}

func (*UnimplementedDriverServer) FindDriver(ctx context.Context, req *DriverRequest) (*DriverResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method FindDriver not implemented")
}

func RegisterDriverServer(s *grpc.Server, srv DriverServer) {
	s.RegisterService(&_Driver_serviceDesc, srv)
}

func _Driver_FindDriver_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DriverRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DriverServer).FindDriver(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/Driver/FindDriver",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DriverServer).FindDriver(ctx, req.(*DriverRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _Driver_serviceDesc = grpc.ServiceDesc{
	ServiceName: "Driver",
	HandlerType: (*DriverServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "findDriver",
			Handler:    _Driver_FindDriver_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "driverService.proto",
}