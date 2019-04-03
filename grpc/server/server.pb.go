// Code generated by protoc-gen-go. DO NOT EDIT.
// source: server.proto

package server

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

type Server struct {
	Host                 string   `protobuf:"bytes,1,opt,name=host,proto3" json:"host,omitempty"`
	Port                 string   `protobuf:"bytes,2,opt,name=port,proto3" json:"port,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Server) Reset()         { *m = Server{} }
func (m *Server) String() string { return proto.CompactTextString(m) }
func (*Server) ProtoMessage()    {}
func (*Server) Descriptor() ([]byte, []int) {
	return fileDescriptor_ad098daeda4239f7, []int{0}
}

func (m *Server) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Server.Unmarshal(m, b)
}
func (m *Server) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Server.Marshal(b, m, deterministic)
}
func (m *Server) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Server.Merge(m, src)
}
func (m *Server) XXX_Size() int {
	return xxx_messageInfo_Server.Size(m)
}
func (m *Server) XXX_DiscardUnknown() {
	xxx_messageInfo_Server.DiscardUnknown(m)
}

var xxx_messageInfo_Server proto.InternalMessageInfo

func (m *Server) GetHost() string {
	if m != nil {
		return m.Host
	}
	return ""
}

func (m *Server) GetPort() string {
	if m != nil {
		return m.Port
	}
	return ""
}

type Servers struct {
	S                    []*Server `protobuf:"bytes,3,rep,name=s,proto3" json:"s,omitempty"`
	XXX_NoUnkeyedLiteral struct{}  `json:"-"`
	XXX_unrecognized     []byte    `json:"-"`
	XXX_sizecache        int32     `json:"-"`
}

func (m *Servers) Reset()         { *m = Servers{} }
func (m *Servers) String() string { return proto.CompactTextString(m) }
func (*Servers) ProtoMessage()    {}
func (*Servers) Descriptor() ([]byte, []int) {
	return fileDescriptor_ad098daeda4239f7, []int{1}
}

func (m *Servers) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Servers.Unmarshal(m, b)
}
func (m *Servers) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Servers.Marshal(b, m, deterministic)
}
func (m *Servers) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Servers.Merge(m, src)
}
func (m *Servers) XXX_Size() int {
	return xxx_messageInfo_Servers.Size(m)
}
func (m *Servers) XXX_DiscardUnknown() {
	xxx_messageInfo_Servers.DiscardUnknown(m)
}

var xxx_messageInfo_Servers proto.InternalMessageInfo

func (m *Servers) GetS() []*Server {
	if m != nil {
		return m.S
	}
	return nil
}

type Cacheitem struct {
	Key                  string   `protobuf:"bytes,4,opt,name=key,proto3" json:"key,omitempty"`
	Value                string   `protobuf:"bytes,5,opt,name=value,proto3" json:"value,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Cacheitem) Reset()         { *m = Cacheitem{} }
func (m *Cacheitem) String() string { return proto.CompactTextString(m) }
func (*Cacheitem) ProtoMessage()    {}
func (*Cacheitem) Descriptor() ([]byte, []int) {
	return fileDescriptor_ad098daeda4239f7, []int{2}
}

func (m *Cacheitem) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Cacheitem.Unmarshal(m, b)
}
func (m *Cacheitem) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Cacheitem.Marshal(b, m, deterministic)
}
func (m *Cacheitem) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Cacheitem.Merge(m, src)
}
func (m *Cacheitem) XXX_Size() int {
	return xxx_messageInfo_Cacheitem.Size(m)
}
func (m *Cacheitem) XXX_DiscardUnknown() {
	xxx_messageInfo_Cacheitem.DiscardUnknown(m)
}

var xxx_messageInfo_Cacheitem proto.InternalMessageInfo

func (m *Cacheitem) GetKey() string {
	if m != nil {
		return m.Key
	}
	return ""
}

func (m *Cacheitem) GetValue() string {
	if m != nil {
		return m.Value
	}
	return ""
}

type Key struct {
	Key                  string   `protobuf:"bytes,6,opt,name=key,proto3" json:"key,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Key) Reset()         { *m = Key{} }
func (m *Key) String() string { return proto.CompactTextString(m) }
func (*Key) ProtoMessage()    {}
func (*Key) Descriptor() ([]byte, []int) {
	return fileDescriptor_ad098daeda4239f7, []int{3}
}

func (m *Key) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Key.Unmarshal(m, b)
}
func (m *Key) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Key.Marshal(b, m, deterministic)
}
func (m *Key) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Key.Merge(m, src)
}
func (m *Key) XXX_Size() int {
	return xxx_messageInfo_Key.Size(m)
}
func (m *Key) XXX_DiscardUnknown() {
	xxx_messageInfo_Key.DiscardUnknown(m)
}

var xxx_messageInfo_Key proto.InternalMessageInfo

func (m *Key) GetKey() string {
	if m != nil {
		return m.Key
	}
	return ""
}

type Client struct {
	Host                 string   `protobuf:"bytes,7,opt,name=host,proto3" json:"host,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Client) Reset()         { *m = Client{} }
func (m *Client) String() string { return proto.CompactTextString(m) }
func (*Client) ProtoMessage()    {}
func (*Client) Descriptor() ([]byte, []int) {
	return fileDescriptor_ad098daeda4239f7, []int{4}
}

func (m *Client) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Client.Unmarshal(m, b)
}
func (m *Client) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Client.Marshal(b, m, deterministic)
}
func (m *Client) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Client.Merge(m, src)
}
func (m *Client) XXX_Size() int {
	return xxx_messageInfo_Client.Size(m)
}
func (m *Client) XXX_DiscardUnknown() {
	xxx_messageInfo_Client.DiscardUnknown(m)
}

var xxx_messageInfo_Client proto.InternalMessageInfo

func (m *Client) GetHost() string {
	if m != nil {
		return m.Host
	}
	return ""
}

func init() {
	proto.RegisterType((*Server)(nil), "server.server")
	proto.RegisterType((*Servers)(nil), "server.servers")
	proto.RegisterType((*Cacheitem)(nil), "server.cacheitem")
	proto.RegisterType((*Key)(nil), "server.key")
	proto.RegisterType((*Client)(nil), "server.client")
}

func init() { proto.RegisterFile("server.proto", fileDescriptor_ad098daeda4239f7) }

var fileDescriptor_ad098daeda4239f7 = []byte{
	// 258 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x8c, 0x91, 0xcf, 0x6a, 0xc2, 0x40,
	0x10, 0xc6, 0xdd, 0xae, 0xae, 0x75, 0x2c, 0xfd, 0x33, 0x14, 0xba, 0x88, 0x07, 0xd9, 0x8b, 0x52,
	0xa8, 0x15, 0x7d, 0x82, 0x42, 0xc1, 0xbb, 0xa1, 0x0f, 0x90, 0x26, 0x03, 0x06, 0xd3, 0x26, 0xec,
	0x4e, 0x03, 0x3e, 0x76, 0xdf, 0xa0, 0x64, 0x37, 0x49, 0x29, 0xf5, 0xe0, 0x29, 0xdf, 0x7c, 0xf9,
	0xcd, 0xf0, 0x83, 0x85, 0x2b, 0x47, 0xb6, 0x22, 0xbb, 0x2c, 0x6d, 0xc1, 0x05, 0xaa, 0x30, 0x99,
	0x15, 0x34, 0x09, 0x11, 0xfa, 0xfb, 0xc2, 0xb1, 0x16, 0x33, 0xb1, 0x18, 0xed, 0x7c, 0xae, 0xbb,
	0xb2, 0xb0, 0xac, 0x2f, 0x42, 0x57, 0x67, 0x33, 0x87, 0x61, 0xd8, 0x70, 0x38, 0x05, 0xe1, 0xb4,
	0x9c, 0xc9, 0xc5, 0x78, 0x7d, 0xbd, 0x6c, 0xce, 0x87, 0xcf, 0x4e, 0x38, 0xb3, 0x81, 0x51, 0x12,
	0x27, 0x7b, 0xca, 0x98, 0x3e, 0xf0, 0x16, 0xe4, 0x81, 0x8e, 0xba, 0xef, 0x0f, 0xd5, 0x11, 0xef,
	0x61, 0x50, 0xc5, 0xf9, 0x17, 0xe9, 0x81, 0xef, 0xc2, 0x60, 0x1e, 0x3c, 0xd7, 0xe2, 0xaa, 0xc3,
	0xcd, 0x14, 0x54, 0x92, 0x67, 0xf4, 0xc9, 0x9d, 0xe8, 0xf0, 0x57, 0x74, 0xfd, 0x2d, 0xe0, 0x32,
	0x8a, 0xc8, 0x56, 0x59, 0x42, 0xf8, 0x0c, 0xb0, 0x25, 0x8e, 0x1a, 0xc9, 0xce, 0x2c, 0xac, 0x4f,
	0x6e, 0xfe, 0x9a, 0x3a, 0xd3, 0xc3, 0x27, 0x90, 0x2f, 0x69, 0x8a, 0x77, 0x1d, 0xd9, 0x6a, 0x4f,
	0xfe, 0x57, 0xa6, 0x87, 0x73, 0x90, 0x5b, 0x62, 0x1c, 0xb7, 0xff, 0x0e, 0x74, 0x3c, 0x0d, 0x3e,
	0x82, 0x7a, 0xa5, 0x9c, 0x98, 0xce, 0x60, 0x57, 0xa0, 0xde, 0xca, 0x34, 0x66, 0x3a, 0x57, 0xe3,
	0x5d, 0xf9, 0x97, 0xdc, 0xfc, 0x04, 0x00, 0x00, 0xff, 0xff, 0x6d, 0x28, 0xcb, 0x4e, 0xd9, 0x01,
	0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// SServiceClient is the client API for SService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type SServiceClient interface {
	GetServers(ctx context.Context, in *Client, opts ...grpc.CallOption) (*Servers, error)
	Add(ctx context.Context, in *Cacheitem, opts ...grpc.CallOption) (*Cacheitem, error)
	Get(ctx context.Context, in *Key, opts ...grpc.CallOption) (*Cacheitem, error)
	Delete(ctx context.Context, in *Key, opts ...grpc.CallOption) (*Cacheitem, error)
	Update(ctx context.Context, in *Cacheitem, opts ...grpc.CallOption) (*Cacheitem, error)
}

type sServiceClient struct {
	cc *grpc.ClientConn
}

func NewSServiceClient(cc *grpc.ClientConn) SServiceClient {
	return &sServiceClient{cc}
}

func (c *sServiceClient) GetServers(ctx context.Context, in *Client, opts ...grpc.CallOption) (*Servers, error) {
	out := new(Servers)
	err := c.cc.Invoke(ctx, "/server.SService/GetServers", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *sServiceClient) Add(ctx context.Context, in *Cacheitem, opts ...grpc.CallOption) (*Cacheitem, error) {
	out := new(Cacheitem)
	err := c.cc.Invoke(ctx, "/server.SService/Add", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *sServiceClient) Get(ctx context.Context, in *Key, opts ...grpc.CallOption) (*Cacheitem, error) {
	out := new(Cacheitem)
	err := c.cc.Invoke(ctx, "/server.SService/Get", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *sServiceClient) Delete(ctx context.Context, in *Key, opts ...grpc.CallOption) (*Cacheitem, error) {
	out := new(Cacheitem)
	err := c.cc.Invoke(ctx, "/server.SService/Delete", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *sServiceClient) Update(ctx context.Context, in *Cacheitem, opts ...grpc.CallOption) (*Cacheitem, error) {
	out := new(Cacheitem)
	err := c.cc.Invoke(ctx, "/server.SService/Update", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// SServiceServer is the server API for SService service.
type SServiceServer interface {
	GetServers(context.Context, *Client) (*Servers, error)
	Add(context.Context, *Cacheitem) (*Cacheitem, error)
	Get(context.Context, *Key) (*Cacheitem, error)
	Delete(context.Context, *Key) (*Cacheitem, error)
	Update(context.Context, *Cacheitem) (*Cacheitem, error)
}

// UnimplementedSServiceServer can be embedded to have forward compatible implementations.
type UnimplementedSServiceServer struct {
}

func (*UnimplementedSServiceServer) GetServers(ctx context.Context, req *Client) (*Servers, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetServers not implemented")
}
func (*UnimplementedSServiceServer) Add(ctx context.Context, req *Cacheitem) (*Cacheitem, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Add not implemented")
}
func (*UnimplementedSServiceServer) Get(ctx context.Context, req *Key) (*Cacheitem, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Get not implemented")
}
func (*UnimplementedSServiceServer) Delete(ctx context.Context, req *Key) (*Cacheitem, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Delete not implemented")
}
func (*UnimplementedSServiceServer) Update(ctx context.Context, req *Cacheitem) (*Cacheitem, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Update not implemented")
}

func RegisterSServiceServer(s *grpc.Server, srv SServiceServer) {
	s.RegisterService(&_SService_serviceDesc, srv)
}

func _SService_GetServers_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Client)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SServiceServer).GetServers(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/server.SService/GetServers",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SServiceServer).GetServers(ctx, req.(*Client))
	}
	return interceptor(ctx, in, info, handler)
}

func _SService_Add_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Cacheitem)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SServiceServer).Add(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/server.SService/Add",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SServiceServer).Add(ctx, req.(*Cacheitem))
	}
	return interceptor(ctx, in, info, handler)
}

func _SService_Get_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Key)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SServiceServer).Get(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/server.SService/Get",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SServiceServer).Get(ctx, req.(*Key))
	}
	return interceptor(ctx, in, info, handler)
}

func _SService_Delete_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Key)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SServiceServer).Delete(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/server.SService/Delete",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SServiceServer).Delete(ctx, req.(*Key))
	}
	return interceptor(ctx, in, info, handler)
}

func _SService_Update_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Cacheitem)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SServiceServer).Update(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/server.SService/Update",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SServiceServer).Update(ctx, req.(*Cacheitem))
	}
	return interceptor(ctx, in, info, handler)
}

var _SService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "server.SService",
	HandlerType: (*SServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetServers",
			Handler:    _SService_GetServers_Handler,
		},
		{
			MethodName: "Add",
			Handler:    _SService_Add_Handler,
		},
		{
			MethodName: "Get",
			Handler:    _SService_Get_Handler,
		},
		{
			MethodName: "Delete",
			Handler:    _SService_Delete_Handler,
		},
		{
			MethodName: "Update",
			Handler:    _SService_Update_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "server.proto",
}