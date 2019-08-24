// Code generated by protoc-gen-go. DO NOT EDIT.
// source: grpcpb/app.proto

package grpcpb

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

type GetServerInfoRequest struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *GetServerInfoRequest) Reset()         { *m = GetServerInfoRequest{} }
func (m *GetServerInfoRequest) String() string { return proto.CompactTextString(m) }
func (*GetServerInfoRequest) ProtoMessage()    {}
func (*GetServerInfoRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_457c2b1558d77d0b, []int{0}
}

func (m *GetServerInfoRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetServerInfoRequest.Unmarshal(m, b)
}
func (m *GetServerInfoRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetServerInfoRequest.Marshal(b, m, deterministic)
}
func (m *GetServerInfoRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetServerInfoRequest.Merge(m, src)
}
func (m *GetServerInfoRequest) XXX_Size() int {
	return xxx_messageInfo_GetServerInfoRequest.Size(m)
}
func (m *GetServerInfoRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_GetServerInfoRequest.DiscardUnknown(m)
}

var xxx_messageInfo_GetServerInfoRequest proto.InternalMessageInfo

type GetServerResponse struct {
	Hostname             string   `protobuf:"bytes,1,opt,name=hostname,proto3" json:"hostname,omitempty"`
	Currentservertime    string   `protobuf:"bytes,2,opt,name=currentservertime,proto3" json:"currentservertime,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *GetServerResponse) Reset()         { *m = GetServerResponse{} }
func (m *GetServerResponse) String() string { return proto.CompactTextString(m) }
func (*GetServerResponse) ProtoMessage()    {}
func (*GetServerResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_457c2b1558d77d0b, []int{1}
}

func (m *GetServerResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetServerResponse.Unmarshal(m, b)
}
func (m *GetServerResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetServerResponse.Marshal(b, m, deterministic)
}
func (m *GetServerResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetServerResponse.Merge(m, src)
}
func (m *GetServerResponse) XXX_Size() int {
	return xxx_messageInfo_GetServerResponse.Size(m)
}
func (m *GetServerResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_GetServerResponse.DiscardUnknown(m)
}

var xxx_messageInfo_GetServerResponse proto.InternalMessageInfo

func (m *GetServerResponse) GetHostname() string {
	if m != nil {
		return m.Hostname
	}
	return ""
}

func (m *GetServerResponse) GetCurrentservertime() string {
	if m != nil {
		return m.Currentservertime
	}
	return ""
}

func init() {
	proto.RegisterType((*GetServerInfoRequest)(nil), "golang_grpc_demo.GetServerInfoRequest")
	proto.RegisterType((*GetServerResponse)(nil), "golang_grpc_demo.GetServerResponse")
}

func init() { proto.RegisterFile("grpcpb/app.proto", fileDescriptor_457c2b1558d77d0b) }

var fileDescriptor_457c2b1558d77d0b = []byte{
	// 183 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0x12, 0x48, 0x2f, 0x2a, 0x48,
	0x2e, 0x48, 0xd2, 0x4f, 0x2c, 0x28, 0xd0, 0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17, 0x12, 0x48, 0xcf,
	0xcf, 0x49, 0xcc, 0x4b, 0x8f, 0x07, 0x49, 0xc4, 0xa7, 0xa4, 0xe6, 0xe6, 0x2b, 0x89, 0x71, 0x89,
	0xb8, 0xa7, 0x96, 0x04, 0xa7, 0x16, 0x95, 0xa5, 0x16, 0x79, 0xe6, 0xa5, 0xe5, 0x07, 0xa5, 0x16,
	0x96, 0xa6, 0x16, 0x97, 0x28, 0xc5, 0x72, 0x09, 0xc2, 0xc5, 0x83, 0x52, 0x8b, 0x0b, 0xf2, 0xf3,
	0x8a, 0x53, 0x85, 0xa4, 0xb8, 0x38, 0x32, 0xf2, 0x8b, 0x4b, 0xf2, 0x12, 0x73, 0x53, 0x25, 0x18,
	0x15, 0x18, 0x35, 0x38, 0x83, 0xe0, 0x7c, 0x21, 0x1d, 0x2e, 0xc1, 0xe4, 0xd2, 0xa2, 0xa2, 0xd4,
	0xbc, 0x92, 0x62, 0xb0, 0xa6, 0x92, 0xcc, 0xdc, 0x54, 0x09, 0x26, 0xb0, 0x22, 0x4c, 0x09, 0xa3,
	0x1c, 0x2e, 0x2e, 0x84, 0x9d, 0x42, 0x71, 0x5c, 0xbc, 0xe9, 0xc8, 0x8e, 0x10, 0x52, 0xd3, 0x43,
	0x77, 0xa8, 0x1e, 0x36, 0x57, 0x4a, 0x29, 0xe3, 0x51, 0x07, 0x73, 0xb5, 0x12, 0x83, 0x13, 0x47,
	0x14, 0x1b, 0x24, 0x28, 0x92, 0xd8, 0xc0, 0xe1, 0x60, 0x0c, 0x08, 0x00, 0x00, 0xff, 0xff, 0x1a,
	0x86, 0x3a, 0x22, 0x1b, 0x01, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// ServerInfoClient is the client API for ServerInfo service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type ServerInfoClient interface {
	GetServerInfo(ctx context.Context, in *GetServerInfoRequest, opts ...grpc.CallOption) (*GetServerResponse, error)
}

type serverInfoClient struct {
	cc *grpc.ClientConn
}

func NewServerInfoClient(cc *grpc.ClientConn) ServerInfoClient {
	return &serverInfoClient{cc}
}

func (c *serverInfoClient) GetServerInfo(ctx context.Context, in *GetServerInfoRequest, opts ...grpc.CallOption) (*GetServerResponse, error) {
	out := new(GetServerResponse)
	err := c.cc.Invoke(ctx, "/golang_grpc_demo.ServerInfo/getServerInfo", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// ServerInfoServer is the server API for ServerInfo service.
type ServerInfoServer interface {
	GetServerInfo(context.Context, *GetServerInfoRequest) (*GetServerResponse, error)
}

// UnimplementedServerInfoServer can be embedded to have forward compatible implementations.
type UnimplementedServerInfoServer struct {
}

func (*UnimplementedServerInfoServer) GetServerInfo(ctx context.Context, req *GetServerInfoRequest) (*GetServerResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetServerInfo not implemented")
}

func RegisterServerInfoServer(s *grpc.Server, srv ServerInfoServer) {
	s.RegisterService(&_ServerInfo_serviceDesc, srv)
}

func _ServerInfo_GetServerInfo_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetServerInfoRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ServerInfoServer).GetServerInfo(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/golang_grpc_demo.ServerInfo/GetServerInfo",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ServerInfoServer).GetServerInfo(ctx, req.(*GetServerInfoRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _ServerInfo_serviceDesc = grpc.ServiceDesc{
	ServiceName: "golang_grpc_demo.ServerInfo",
	HandlerType: (*ServerInfoServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "getServerInfo",
			Handler:    _ServerInfo_GetServerInfo_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "grpcpb/app.proto",
}