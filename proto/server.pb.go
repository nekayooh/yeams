// Code generated by protoc-gen-go. DO NOT EDIT.
// source: server.proto

package proto

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

type ReturnCode int32

const (
	ReturnCode_Success ReturnCode = 0
	ReturnCode_Failure ReturnCode = 1
)

var ReturnCode_name = map[int32]string{
	0: "Success",
	1: "Failure",
}

var ReturnCode_value = map[string]int32{
	"Success": 0,
	"Failure": 1,
}

func (x ReturnCode) String() string {
	return proto.EnumName(ReturnCode_name, int32(x))
}

func (ReturnCode) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_ad098daeda4239f7, []int{0}
}

//-----------注册-----------
//注册模块
type RegMsg struct {
	Version              string     `protobuf:"bytes,1,opt,name=version,proto3" json:"version,omitempty"`
	Key                  []byte     `protobuf:"bytes,2,opt,name=key,proto3" json:"key,omitempty"`
	Name                 string     `protobuf:"bytes,3,opt,name=name,proto3" json:"name,omitempty"`
	Address              string     `protobuf:"bytes,4,opt,name=address,proto3" json:"address,omitempty"`
	Port                 int64      `protobuf:"varint,5,opt,name=port,proto3" json:"port,omitempty"`
	Thread               int64      `protobuf:"varint,6,opt,name=thread,proto3" json:"thread,omitempty"`
	Module               []*SendMsg `protobuf:"bytes,7,rep,name=module,proto3" json:"module,omitempty"`
	XXX_NoUnkeyedLiteral struct{}   `json:"-"`
	XXX_unrecognized     []byte     `json:"-"`
	XXX_sizecache        int32      `json:"-"`
}

func (m *RegMsg) Reset()         { *m = RegMsg{} }
func (m *RegMsg) String() string { return proto.CompactTextString(m) }
func (*RegMsg) ProtoMessage()    {}
func (*RegMsg) Descriptor() ([]byte, []int) {
	return fileDescriptor_ad098daeda4239f7, []int{0}
}

func (m *RegMsg) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_RegMsg.Unmarshal(m, b)
}
func (m *RegMsg) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_RegMsg.Marshal(b, m, deterministic)
}
func (m *RegMsg) XXX_Merge(src proto.Message) {
	xxx_messageInfo_RegMsg.Merge(m, src)
}
func (m *RegMsg) XXX_Size() int {
	return xxx_messageInfo_RegMsg.Size(m)
}
func (m *RegMsg) XXX_DiscardUnknown() {
	xxx_messageInfo_RegMsg.DiscardUnknown(m)
}

var xxx_messageInfo_RegMsg proto.InternalMessageInfo

func (m *RegMsg) GetVersion() string {
	if m != nil {
		return m.Version
	}
	return ""
}

func (m *RegMsg) GetKey() []byte {
	if m != nil {
		return m.Key
	}
	return nil
}

func (m *RegMsg) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *RegMsg) GetAddress() string {
	if m != nil {
		return m.Address
	}
	return ""
}

func (m *RegMsg) GetPort() int64 {
	if m != nil {
		return m.Port
	}
	return 0
}

func (m *RegMsg) GetThread() int64 {
	if m != nil {
		return m.Thread
	}
	return 0
}

func (m *RegMsg) GetModule() []*SendMsg {
	if m != nil {
		return m.Module
	}
	return nil
}

//注册成功后返回信息
type RegRtnMsg struct {
	Status               ReturnCode `protobuf:"varint,1,opt,name=status,proto3,enum=proto.ReturnCode" json:"status,omitempty"`
	Uuid                 string     `protobuf:"bytes,2,opt,name=uuid,proto3" json:"uuid,omitempty"`
	XXX_NoUnkeyedLiteral struct{}   `json:"-"`
	XXX_unrecognized     []byte     `json:"-"`
	XXX_sizecache        int32      `json:"-"`
}

func (m *RegRtnMsg) Reset()         { *m = RegRtnMsg{} }
func (m *RegRtnMsg) String() string { return proto.CompactTextString(m) }
func (*RegRtnMsg) ProtoMessage()    {}
func (*RegRtnMsg) Descriptor() ([]byte, []int) {
	return fileDescriptor_ad098daeda4239f7, []int{1}
}

func (m *RegRtnMsg) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_RegRtnMsg.Unmarshal(m, b)
}
func (m *RegRtnMsg) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_RegRtnMsg.Marshal(b, m, deterministic)
}
func (m *RegRtnMsg) XXX_Merge(src proto.Message) {
	xxx_messageInfo_RegRtnMsg.Merge(m, src)
}
func (m *RegRtnMsg) XXX_Size() int {
	return xxx_messageInfo_RegRtnMsg.Size(m)
}
func (m *RegRtnMsg) XXX_DiscardUnknown() {
	xxx_messageInfo_RegRtnMsg.DiscardUnknown(m)
}

var xxx_messageInfo_RegRtnMsg proto.InternalMessageInfo

func (m *RegRtnMsg) GetStatus() ReturnCode {
	if m != nil {
		return m.Status
	}
	return ReturnCode_Success
}

func (m *RegRtnMsg) GetUuid() string {
	if m != nil {
		return m.Uuid
	}
	return ""
}

//-----------取消注册-----------
//取消注册模块
type UnRegMsg struct {
	Uuid                 string     `protobuf:"bytes,1,opt,name=uuid,proto3" json:"uuid,omitempty"`
	Name                 string     `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	Module               []*SendMsg `protobuf:"bytes,3,rep,name=module,proto3" json:"module,omitempty"`
	XXX_NoUnkeyedLiteral struct{}   `json:"-"`
	XXX_unrecognized     []byte     `json:"-"`
	XXX_sizecache        int32      `json:"-"`
}

func (m *UnRegMsg) Reset()         { *m = UnRegMsg{} }
func (m *UnRegMsg) String() string { return proto.CompactTextString(m) }
func (*UnRegMsg) ProtoMessage()    {}
func (*UnRegMsg) Descriptor() ([]byte, []int) {
	return fileDescriptor_ad098daeda4239f7, []int{2}
}

func (m *UnRegMsg) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_UnRegMsg.Unmarshal(m, b)
}
func (m *UnRegMsg) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_UnRegMsg.Marshal(b, m, deterministic)
}
func (m *UnRegMsg) XXX_Merge(src proto.Message) {
	xxx_messageInfo_UnRegMsg.Merge(m, src)
}
func (m *UnRegMsg) XXX_Size() int {
	return xxx_messageInfo_UnRegMsg.Size(m)
}
func (m *UnRegMsg) XXX_DiscardUnknown() {
	xxx_messageInfo_UnRegMsg.DiscardUnknown(m)
}

var xxx_messageInfo_UnRegMsg proto.InternalMessageInfo

func (m *UnRegMsg) GetUuid() string {
	if m != nil {
		return m.Uuid
	}
	return ""
}

func (m *UnRegMsg) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *UnRegMsg) GetModule() []*SendMsg {
	if m != nil {
		return m.Module
	}
	return nil
}

//取消注册成功后返回信息
type UnRegRtnMsg struct {
	Status               ReturnCode `protobuf:"varint,1,opt,name=status,proto3,enum=proto.ReturnCode" json:"status,omitempty"`
	XXX_NoUnkeyedLiteral struct{}   `json:"-"`
	XXX_unrecognized     []byte     `json:"-"`
	XXX_sizecache        int32      `json:"-"`
}

func (m *UnRegRtnMsg) Reset()         { *m = UnRegRtnMsg{} }
func (m *UnRegRtnMsg) String() string { return proto.CompactTextString(m) }
func (*UnRegRtnMsg) ProtoMessage()    {}
func (*UnRegRtnMsg) Descriptor() ([]byte, []int) {
	return fileDescriptor_ad098daeda4239f7, []int{3}
}

func (m *UnRegRtnMsg) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_UnRegRtnMsg.Unmarshal(m, b)
}
func (m *UnRegRtnMsg) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_UnRegRtnMsg.Marshal(b, m, deterministic)
}
func (m *UnRegRtnMsg) XXX_Merge(src proto.Message) {
	xxx_messageInfo_UnRegRtnMsg.Merge(m, src)
}
func (m *UnRegRtnMsg) XXX_Size() int {
	return xxx_messageInfo_UnRegRtnMsg.Size(m)
}
func (m *UnRegRtnMsg) XXX_DiscardUnknown() {
	xxx_messageInfo_UnRegRtnMsg.DiscardUnknown(m)
}

var xxx_messageInfo_UnRegRtnMsg proto.InternalMessageInfo

func (m *UnRegRtnMsg) GetStatus() ReturnCode {
	if m != nil {
		return m.Status
	}
	return ReturnCode_Success
}

//-----------通信-----------
type SendMsg struct {
	Name                 string   `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	Json                 []byte   `protobuf:"bytes,2,opt,name=json,proto3" json:"json,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *SendMsg) Reset()         { *m = SendMsg{} }
func (m *SendMsg) String() string { return proto.CompactTextString(m) }
func (*SendMsg) ProtoMessage()    {}
func (*SendMsg) Descriptor() ([]byte, []int) {
	return fileDescriptor_ad098daeda4239f7, []int{4}
}

func (m *SendMsg) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_SendMsg.Unmarshal(m, b)
}
func (m *SendMsg) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_SendMsg.Marshal(b, m, deterministic)
}
func (m *SendMsg) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SendMsg.Merge(m, src)
}
func (m *SendMsg) XXX_Size() int {
	return xxx_messageInfo_SendMsg.Size(m)
}
func (m *SendMsg) XXX_DiscardUnknown() {
	xxx_messageInfo_SendMsg.DiscardUnknown(m)
}

var xxx_messageInfo_SendMsg proto.InternalMessageInfo

func (m *SendMsg) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *SendMsg) GetJson() []byte {
	if m != nil {
		return m.Json
	}
	return nil
}

type SendRtnMsg struct {
	Status               ReturnCode `protobuf:"varint,1,opt,name=status,proto3,enum=proto.ReturnCode" json:"status,omitempty"`
	Json                 []byte     `protobuf:"bytes,2,opt,name=json,proto3" json:"json,omitempty"`
	XXX_NoUnkeyedLiteral struct{}   `json:"-"`
	XXX_unrecognized     []byte     `json:"-"`
	XXX_sizecache        int32      `json:"-"`
}

func (m *SendRtnMsg) Reset()         { *m = SendRtnMsg{} }
func (m *SendRtnMsg) String() string { return proto.CompactTextString(m) }
func (*SendRtnMsg) ProtoMessage()    {}
func (*SendRtnMsg) Descriptor() ([]byte, []int) {
	return fileDescriptor_ad098daeda4239f7, []int{5}
}

func (m *SendRtnMsg) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_SendRtnMsg.Unmarshal(m, b)
}
func (m *SendRtnMsg) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_SendRtnMsg.Marshal(b, m, deterministic)
}
func (m *SendRtnMsg) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SendRtnMsg.Merge(m, src)
}
func (m *SendRtnMsg) XXX_Size() int {
	return xxx_messageInfo_SendRtnMsg.Size(m)
}
func (m *SendRtnMsg) XXX_DiscardUnknown() {
	xxx_messageInfo_SendRtnMsg.DiscardUnknown(m)
}

var xxx_messageInfo_SendRtnMsg proto.InternalMessageInfo

func (m *SendRtnMsg) GetStatus() ReturnCode {
	if m != nil {
		return m.Status
	}
	return ReturnCode_Success
}

func (m *SendRtnMsg) GetJson() []byte {
	if m != nil {
		return m.Json
	}
	return nil
}

func init() {
	proto.RegisterEnum("proto.ReturnCode", ReturnCode_name, ReturnCode_value)
	proto.RegisterType((*RegMsg)(nil), "proto.RegMsg")
	proto.RegisterType((*RegRtnMsg)(nil), "proto.RegRtnMsg")
	proto.RegisterType((*UnRegMsg)(nil), "proto.UnRegMsg")
	proto.RegisterType((*UnRegRtnMsg)(nil), "proto.UnRegRtnMsg")
	proto.RegisterType((*SendMsg)(nil), "proto.SendMsg")
	proto.RegisterType((*SendRtnMsg)(nil), "proto.SendRtnMsg")
}

func init() { proto.RegisterFile("server.proto", fileDescriptor_ad098daeda4239f7) }

var fileDescriptor_ad098daeda4239f7 = []byte{
	// 389 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x9c, 0x91, 0xbf, 0xae, 0xd3, 0x30,
	0x14, 0xc6, 0xaf, 0x6f, 0x7a, 0x93, 0x9b, 0xd3, 0x52, 0x52, 0x0f, 0xc8, 0xea, 0x14, 0x65, 0xa8,
	0x42, 0x25, 0x2a, 0x91, 0x2e, 0xec, 0x48, 0x0c, 0xa0, 0x32, 0xb8, 0x62, 0x80, 0x2d, 0x34, 0x47,
	0x21, 0xd0, 0xc6, 0x95, 0xed, 0x54, 0xe2, 0x59, 0x78, 0x0f, 0x9e, 0x0f, 0xd9, 0x71, 0xfe, 0xa8,
	0x03, 0x52, 0x3b, 0xe5, 0x7c, 0xf6, 0xf7, 0x9d, 0xfc, 0xce, 0x31, 0xcc, 0x14, 0xca, 0x0b, 0xca,
	0xcd, 0x59, 0x0a, 0x2d, 0xe8, 0x93, 0xfd, 0x24, 0x7f, 0x09, 0xf8, 0x1c, 0xcb, 0x9d, 0x2a, 0x29,
	0x83, 0xe0, 0x82, 0x52, 0x55, 0xa2, 0x66, 0x24, 0x26, 0x69, 0xc8, 0x3b, 0x49, 0x23, 0xf0, 0x7e,
	0xe1, 0x6f, 0xf6, 0x18, 0x93, 0x74, 0xc6, 0x4d, 0x49, 0x29, 0x4c, 0xea, 0xfc, 0x84, 0xcc, 0xb3,
	0x46, 0x5b, 0x9b, 0x7c, 0x5e, 0x14, 0x12, 0x95, 0x62, 0x93, 0x36, 0xef, 0xa4, 0x71, 0x9f, 0x85,
	0xd4, 0xec, 0x29, 0x26, 0xa9, 0xc7, 0x6d, 0x4d, 0x5f, 0x81, 0xaf, 0x7f, 0x48, 0xcc, 0x0b, 0xe6,
	0xdb, 0x53, 0xa7, 0xe8, 0x0a, 0xfc, 0x93, 0x28, 0x9a, 0x23, 0xb2, 0x20, 0xf6, 0xd2, 0x69, 0x36,
	0x6f, 0x79, 0x37, 0x7b, 0xac, 0x8b, 0x9d, 0x2a, 0xb9, 0xbb, 0x4d, 0x3e, 0x42, 0xc8, 0xb1, 0xe4,
	0xba, 0x36, 0xe8, 0xaf, 0xc1, 0x57, 0x3a, 0xd7, 0x8d, 0xb2, 0xe4, 0xf3, 0x6c, 0xe1, 0x42, 0x1c,
	0x75, 0x23, 0xeb, 0xf7, 0xa2, 0x40, 0xee, 0x0c, 0x86, 0xa5, 0x69, 0xaa, 0xc2, 0x0e, 0x13, 0x72,
	0x5b, 0x27, 0xdf, 0xe0, 0xf9, 0x4b, 0xed, 0xb6, 0xd0, 0xdd, 0x93, 0xe1, 0xbe, 0x9f, 0xf6, 0x71,
	0x34, 0xed, 0xc0, 0xe9, 0xfd, 0x97, 0xf3, 0x1d, 0x4c, 0x6d, 0xef, 0x9b, 0x49, 0x93, 0xb7, 0x10,
	0xb8, 0x66, 0x3d, 0x00, 0x19, 0x01, 0x50, 0x98, 0xfc, 0x54, 0xa2, 0x76, 0xaf, 0x62, 0xeb, 0xe4,
	0x13, 0x80, 0x89, 0xdc, 0xb5, 0x95, 0xeb, 0x66, 0xeb, 0x15, 0xc0, 0xe0, 0xa4, 0x53, 0x08, 0xf6,
	0xcd, 0xe1, 0x80, 0x4a, 0x45, 0x0f, 0x46, 0x7c, 0xc8, 0xab, 0x63, 0x23, 0x31, 0x22, 0xd9, 0x1f,
	0x02, 0xe1, 0x57, 0xcc, 0x77, 0x76, 0x5e, 0xfa, 0x06, 0x9e, 0x39, 0x96, 0x95, 0xd2, 0x28, 0xe9,
	0x8b, 0xfe, 0x87, 0x66, 0xb5, 0xcb, 0x68, 0x90, 0x2d, 0x61, 0xf2, 0x40, 0xb7, 0x00, 0x76, 0x3d,
	0x6d, 0xe0, 0xa5, 0x73, 0x74, 0xaf, 0xb1, 0xa4, 0xe3, 0x83, 0x3e, 0xb4, 0x06, 0xcf, 0xcc, 0x77,
	0xb5, 0xf2, 0xe5, 0x62, 0xa4, 0x3b, 0x6f, 0x76, 0xb4, 0x70, 0x9f, 0x85, 0xae, 0x0e, 0x78, 0x4b,
	0xf0, 0x2e, 0xb2, 0xef, 0xbe, 0x3d, 0xdc, 0xfe, 0x0b, 0x00, 0x00, 0xff, 0xff, 0x7f, 0xc8, 0x97,
	0x13, 0x6c, 0x03, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// YeaModuleClient is the client API for YeaModule service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type YeaModuleClient interface {
	//    注册
	Register(ctx context.Context, in *RegMsg, opts ...grpc.CallOption) (*RegRtnMsg, error)
	//    取消注册
	UnRegister(ctx context.Context, in *UnRegMsg, opts ...grpc.CallOption) (*UnRegRtnMsg, error)
	//    数据查询
	Msg(ctx context.Context, in *SendMsg, opts ...grpc.CallOption) (*SendRtnMsg, error)
}

type yeaModuleClient struct {
	cc *grpc.ClientConn
}

func NewYeaModuleClient(cc *grpc.ClientConn) YeaModuleClient {
	return &yeaModuleClient{cc}
}

func (c *yeaModuleClient) Register(ctx context.Context, in *RegMsg, opts ...grpc.CallOption) (*RegRtnMsg, error) {
	out := new(RegRtnMsg)
	err := c.cc.Invoke(ctx, "/proto.YeaModule/Register", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *yeaModuleClient) UnRegister(ctx context.Context, in *UnRegMsg, opts ...grpc.CallOption) (*UnRegRtnMsg, error) {
	out := new(UnRegRtnMsg)
	err := c.cc.Invoke(ctx, "/proto.YeaModule/UnRegister", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *yeaModuleClient) Msg(ctx context.Context, in *SendMsg, opts ...grpc.CallOption) (*SendRtnMsg, error) {
	out := new(SendRtnMsg)
	err := c.cc.Invoke(ctx, "/proto.YeaModule/Msg", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// YeaModuleServer is the server API for YeaModule service.
type YeaModuleServer interface {
	//    注册
	Register(context.Context, *RegMsg) (*RegRtnMsg, error)
	//    取消注册
	UnRegister(context.Context, *UnRegMsg) (*UnRegRtnMsg, error)
	//    数据查询
	Msg(context.Context, *SendMsg) (*SendRtnMsg, error)
}

// UnimplementedYeaModuleServer can be embedded to have forward compatible implementations.
type UnimplementedYeaModuleServer struct {
}

func (*UnimplementedYeaModuleServer) Register(ctx context.Context, req *RegMsg) (*RegRtnMsg, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Register not implemented")
}
func (*UnimplementedYeaModuleServer) UnRegister(ctx context.Context, req *UnRegMsg) (*UnRegRtnMsg, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UnRegister not implemented")
}
func (*UnimplementedYeaModuleServer) Msg(ctx context.Context, req *SendMsg) (*SendRtnMsg, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Msg not implemented")
}

func RegisterYeaModuleServer(s *grpc.Server, srv YeaModuleServer) {
	s.RegisterService(&_YeaModule_serviceDesc, srv)
}

func _YeaModule_Register_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(RegMsg)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(YeaModuleServer).Register(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/proto.YeaModule/Register",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(YeaModuleServer).Register(ctx, req.(*RegMsg))
	}
	return interceptor(ctx, in, info, handler)
}

func _YeaModule_UnRegister_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UnRegMsg)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(YeaModuleServer).UnRegister(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/proto.YeaModule/UnRegister",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(YeaModuleServer).UnRegister(ctx, req.(*UnRegMsg))
	}
	return interceptor(ctx, in, info, handler)
}

func _YeaModule_Msg_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SendMsg)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(YeaModuleServer).Msg(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/proto.YeaModule/Msg",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(YeaModuleServer).Msg(ctx, req.(*SendMsg))
	}
	return interceptor(ctx, in, info, handler)
}

var _YeaModule_serviceDesc = grpc.ServiceDesc{
	ServiceName: "proto.YeaModule",
	HandlerType: (*YeaModuleServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Register",
			Handler:    _YeaModule_Register_Handler,
		},
		{
			MethodName: "UnRegister",
			Handler:    _YeaModule_UnRegister_Handler,
		},
		{
			MethodName: "Msg",
			Handler:    _YeaModule_Msg_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "server.proto",
}

// YeaNoticeClient is the client API for YeaNotice service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type YeaNoticeClient interface {
	Msg(ctx context.Context, in *SendMsg, opts ...grpc.CallOption) (*SendRtnMsg, error)
	//    取消注册
	UnRegister(ctx context.Context, in *UnRegMsg, opts ...grpc.CallOption) (*UnRegRtnMsg, error)
}

type yeaNoticeClient struct {
	cc *grpc.ClientConn
}

func NewYeaNoticeClient(cc *grpc.ClientConn) YeaNoticeClient {
	return &yeaNoticeClient{cc}
}

func (c *yeaNoticeClient) Msg(ctx context.Context, in *SendMsg, opts ...grpc.CallOption) (*SendRtnMsg, error) {
	out := new(SendRtnMsg)
	err := c.cc.Invoke(ctx, "/proto.YeaNotice/Msg", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *yeaNoticeClient) UnRegister(ctx context.Context, in *UnRegMsg, opts ...grpc.CallOption) (*UnRegRtnMsg, error) {
	out := new(UnRegRtnMsg)
	err := c.cc.Invoke(ctx, "/proto.YeaNotice/UnRegister", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// YeaNoticeServer is the server API for YeaNotice service.
type YeaNoticeServer interface {
	Msg(context.Context, *SendMsg) (*SendRtnMsg, error)
	//    取消注册
	UnRegister(context.Context, *UnRegMsg) (*UnRegRtnMsg, error)
}

// UnimplementedYeaNoticeServer can be embedded to have forward compatible implementations.
type UnimplementedYeaNoticeServer struct {
}

func (*UnimplementedYeaNoticeServer) Msg(ctx context.Context, req *SendMsg) (*SendRtnMsg, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Msg not implemented")
}
func (*UnimplementedYeaNoticeServer) UnRegister(ctx context.Context, req *UnRegMsg) (*UnRegRtnMsg, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UnRegister not implemented")
}

func RegisterYeaNoticeServer(s *grpc.Server, srv YeaNoticeServer) {
	s.RegisterService(&_YeaNotice_serviceDesc, srv)
}

func _YeaNotice_Msg_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SendMsg)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(YeaNoticeServer).Msg(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/proto.YeaNotice/Msg",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(YeaNoticeServer).Msg(ctx, req.(*SendMsg))
	}
	return interceptor(ctx, in, info, handler)
}

func _YeaNotice_UnRegister_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UnRegMsg)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(YeaNoticeServer).UnRegister(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/proto.YeaNotice/UnRegister",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(YeaNoticeServer).UnRegister(ctx, req.(*UnRegMsg))
	}
	return interceptor(ctx, in, info, handler)
}

var _YeaNotice_serviceDesc = grpc.ServiceDesc{
	ServiceName: "proto.YeaNotice",
	HandlerType: (*YeaNoticeServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Msg",
			Handler:    _YeaNotice_Msg_Handler,
		},
		{
			MethodName: "UnRegister",
			Handler:    _YeaNotice_UnRegister_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "server.proto",
}