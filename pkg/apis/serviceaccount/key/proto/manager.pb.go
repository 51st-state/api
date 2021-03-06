// Code generated by protoc-gen-go. DO NOT EDIT.
// source: manager.proto

package serviceaccount_key

import (
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	empty "github.com/golang/protobuf/ptypes/empty"
	context "golang.org/x/net/context"
	grpc "google.golang.org/grpc"
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
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

type Identifier struct {
	GUID                 string   `protobuf:"bytes,1,opt,name=GUID,proto3" json:"GUID,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Identifier) Reset()         { *m = Identifier{} }
func (m *Identifier) String() string { return proto.CompactTextString(m) }
func (*Identifier) ProtoMessage()    {}
func (*Identifier) Descriptor() ([]byte, []int) {
	return fileDescriptor_cde9ec64f0d2c859, []int{0}
}

func (m *Identifier) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Identifier.Unmarshal(m, b)
}
func (m *Identifier) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Identifier.Marshal(b, m, deterministic)
}
func (m *Identifier) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Identifier.Merge(m, src)
}
func (m *Identifier) XXX_Size() int {
	return xxx_messageInfo_Identifier.Size(m)
}
func (m *Identifier) XXX_DiscardUnknown() {
	xxx_messageInfo_Identifier.DiscardUnknown(m)
}

var xxx_messageInfo_Identifier proto.InternalMessageInfo

func (m *Identifier) GetGUID() string {
	if m != nil {
		return m.GUID
	}
	return ""
}

type Incomplete struct {
	Name                 string   `protobuf:"bytes,1,opt,name=Name,proto3" json:"Name,omitempty"`
	Description          string   `protobuf:"bytes,2,opt,name=Description,proto3" json:"Description,omitempty"`
	PublicKey            []byte   `protobuf:"bytes,3,opt,name=PublicKey,proto3" json:"PublicKey,omitempty"`
	ServiceAccountGUID   string   `protobuf:"bytes,4,opt,name=ServiceAccountGUID,proto3" json:"ServiceAccountGUID,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Incomplete) Reset()         { *m = Incomplete{} }
func (m *Incomplete) String() string { return proto.CompactTextString(m) }
func (*Incomplete) ProtoMessage()    {}
func (*Incomplete) Descriptor() ([]byte, []int) {
	return fileDescriptor_cde9ec64f0d2c859, []int{1}
}

func (m *Incomplete) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Incomplete.Unmarshal(m, b)
}
func (m *Incomplete) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Incomplete.Marshal(b, m, deterministic)
}
func (m *Incomplete) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Incomplete.Merge(m, src)
}
func (m *Incomplete) XXX_Size() int {
	return xxx_messageInfo_Incomplete.Size(m)
}
func (m *Incomplete) XXX_DiscardUnknown() {
	xxx_messageInfo_Incomplete.DiscardUnknown(m)
}

var xxx_messageInfo_Incomplete proto.InternalMessageInfo

func (m *Incomplete) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *Incomplete) GetDescription() string {
	if m != nil {
		return m.Description
	}
	return ""
}

func (m *Incomplete) GetPublicKey() []byte {
	if m != nil {
		return m.PublicKey
	}
	return nil
}

func (m *Incomplete) GetServiceAccountGUID() string {
	if m != nil {
		return m.ServiceAccountGUID
	}
	return ""
}

type Complete struct {
	Identifier           *Identifier `protobuf:"bytes,1,opt,name=Identifier,proto3" json:"Identifier,omitempty"`
	Incomplete           *Incomplete `protobuf:"bytes,2,opt,name=Incomplete,proto3" json:"Incomplete,omitempty"`
	XXX_NoUnkeyedLiteral struct{}    `json:"-"`
	XXX_unrecognized     []byte      `json:"-"`
	XXX_sizecache        int32       `json:"-"`
}

func (m *Complete) Reset()         { *m = Complete{} }
func (m *Complete) String() string { return proto.CompactTextString(m) }
func (*Complete) ProtoMessage()    {}
func (*Complete) Descriptor() ([]byte, []int) {
	return fileDescriptor_cde9ec64f0d2c859, []int{2}
}

func (m *Complete) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Complete.Unmarshal(m, b)
}
func (m *Complete) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Complete.Marshal(b, m, deterministic)
}
func (m *Complete) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Complete.Merge(m, src)
}
func (m *Complete) XXX_Size() int {
	return xxx_messageInfo_Complete.Size(m)
}
func (m *Complete) XXX_DiscardUnknown() {
	xxx_messageInfo_Complete.DiscardUnknown(m)
}

var xxx_messageInfo_Complete proto.InternalMessageInfo

func (m *Complete) GetIdentifier() *Identifier {
	if m != nil {
		return m.Identifier
	}
	return nil
}

func (m *Complete) GetIncomplete() *Incomplete {
	if m != nil {
		return m.Incomplete
	}
	return nil
}

type ClientKey struct {
	ServiceAccountGUID   string   `protobuf:"bytes,1,opt,name=ServiceAccountGUID,proto3" json:"ServiceAccountGUID,omitempty"`
	GUID                 string   `protobuf:"bytes,2,opt,name=GUID,proto3" json:"GUID,omitempty"`
	PrivateKey           []byte   `protobuf:"bytes,3,opt,name=PrivateKey,proto3" json:"PrivateKey,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ClientKey) Reset()         { *m = ClientKey{} }
func (m *ClientKey) String() string { return proto.CompactTextString(m) }
func (*ClientKey) ProtoMessage()    {}
func (*ClientKey) Descriptor() ([]byte, []int) {
	return fileDescriptor_cde9ec64f0d2c859, []int{3}
}

func (m *ClientKey) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ClientKey.Unmarshal(m, b)
}
func (m *ClientKey) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ClientKey.Marshal(b, m, deterministic)
}
func (m *ClientKey) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ClientKey.Merge(m, src)
}
func (m *ClientKey) XXX_Size() int {
	return xxx_messageInfo_ClientKey.Size(m)
}
func (m *ClientKey) XXX_DiscardUnknown() {
	xxx_messageInfo_ClientKey.DiscardUnknown(m)
}

var xxx_messageInfo_ClientKey proto.InternalMessageInfo

func (m *ClientKey) GetServiceAccountGUID() string {
	if m != nil {
		return m.ServiceAccountGUID
	}
	return ""
}

func (m *ClientKey) GetGUID() string {
	if m != nil {
		return m.GUID
	}
	return ""
}

func (m *ClientKey) GetPrivateKey() []byte {
	if m != nil {
		return m.PrivateKey
	}
	return nil
}

func init() {
	proto.RegisterType((*Identifier)(nil), "serviceaccount_key.Identifier")
	proto.RegisterType((*Incomplete)(nil), "serviceaccount_key.Incomplete")
	proto.RegisterType((*Complete)(nil), "serviceaccount_key.Complete")
	proto.RegisterType((*ClientKey)(nil), "serviceaccount_key.ClientKey")
}

func init() { proto.RegisterFile("manager.proto", fileDescriptor_cde9ec64f0d2c859) }

var fileDescriptor_cde9ec64f0d2c859 = []byte{
	// 352 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x84, 0x92, 0x41, 0x6b, 0xea, 0x40,
	0x10, 0xc7, 0x8d, 0x4a, 0xde, 0x73, 0x7c, 0xef, 0x32, 0x87, 0x87, 0xe4, 0x59, 0x09, 0x7b, 0xf2,
	0xb4, 0x82, 0xbd, 0x4b, 0x5b, 0x15, 0x09, 0xa5, 0x45, 0x52, 0x3c, 0x97, 0x18, 0x47, 0x59, 0x9a,
	0x64, 0xc3, 0xba, 0x0a, 0x7e, 0x85, 0x9e, 0xfa, 0x15, 0xfa, 0x4d, 0x4b, 0x36, 0xc4, 0x04, 0x6a,
	0xea, 0x2d, 0xcc, 0xfc, 0x67, 0xf2, 0xff, 0xfd, 0x77, 0xe0, 0x6f, 0x1c, 0x24, 0xc1, 0x8e, 0x14,
	0x4f, 0x95, 0xd4, 0x12, 0x71, 0x4f, 0xea, 0x28, 0x42, 0x0a, 0xc2, 0x50, 0x1e, 0x12, 0xfd, 0xfa,
	0x46, 0x27, 0xe7, 0xff, 0x4e, 0xca, 0x5d, 0x44, 0x23, 0xa3, 0x58, 0x1f, 0xb6, 0x23, 0x8a, 0x53,
	0x7d, 0xca, 0x07, 0x98, 0x0b, 0xe0, 0x6d, 0x28, 0xd1, 0x62, 0x2b, 0x48, 0x21, 0x42, 0x7b, 0xb1,
	0xf2, 0x66, 0x3d, 0xcb, 0xb5, 0x86, 0x1d, 0xdf, 0x7c, 0xb3, 0x0f, 0x0b, 0xc0, 0x4b, 0x42, 0x19,
	0xa7, 0x11, 0x69, 0xca, 0x24, 0xcf, 0x41, 0x4c, 0x85, 0x24, 0xfb, 0x46, 0x17, 0xba, 0x33, 0xda,
	0x87, 0x4a, 0xa4, 0x5a, 0xc8, 0xa4, 0xd7, 0x34, 0xad, 0x6a, 0x09, 0xfb, 0xd0, 0x59, 0x1e, 0xd6,
	0x91, 0x08, 0x1f, 0xe9, 0xd4, 0x6b, 0xb9, 0xd6, 0xf0, 0x8f, 0x5f, 0x16, 0x90, 0x03, 0xbe, 0xe4,
	0xbe, 0xef, 0x73, 0xdf, 0xc6, 0x44, 0xdb, 0xac, 0xb9, 0xd0, 0x61, 0xef, 0x16, 0xfc, 0x9e, 0x16,
	0x86, 0x26, 0x55, 0x02, 0x63, 0xab, 0x3b, 0x1e, 0xf0, 0xef, 0x39, 0xf0, 0x52, 0xe5, 0x57, 0x99,
	0x27, 0x55, 0x3c, 0xe3, 0xbd, 0x6e, 0xfe, 0xac, 0xf2, 0x2b, 0x13, 0x4c, 0x42, 0x67, 0x1a, 0x09,
	0x4a, 0x74, 0x3d, 0x89, 0x55, 0x47, 0x72, 0x0e, 0xbc, 0x59, 0x06, 0x8e, 0x03, 0x80, 0xa5, 0x12,
	0xc7, 0x40, 0x53, 0x19, 0x56, 0xa5, 0x32, 0xfe, 0x6c, 0xc2, 0xaf, 0xa7, 0xfc, 0xd5, 0x71, 0x0e,
	0xad, 0x05, 0x69, 0xbc, 0xc2, 0xeb, 0xf4, 0x2f, 0xf5, 0x8b, 0x04, 0x59, 0x03, 0x3d, 0xb0, 0xa7,
	0x8a, 0x02, 0x4d, 0x78, 0x85, 0xdc, 0xb9, 0xb9, 0xb8, 0xa9, 0xe0, 0x67, 0x0d, 0xbc, 0x03, 0x7b,
	0x95, 0x6e, 0xb2, 0x55, 0x3f, 0xfe, 0xd4, 0xf9, 0xc7, 0xf3, 0xb3, 0xe4, 0xc5, 0x59, 0xf2, 0x79,
	0x76, 0x96, 0xac, 0x81, 0x0f, 0x60, 0xcf, 0xc8, 0x3c, 0xed, 0x35, 0xac, 0xda, 0x1d, 0x6b, 0xdb,
	0x54, 0x6e, 0xbf, 0x02, 0x00, 0x00, 0xff, 0xff, 0x48, 0x5c, 0x19, 0xb2, 0x1f, 0x03, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// ManagerClient is the client API for Manager service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type ManagerClient interface {
	Get(ctx context.Context, in *Identifier, opts ...grpc.CallOption) (*Complete, error)
	Create(ctx context.Context, in *Incomplete, opts ...grpc.CallOption) (*ClientKey, error)
	Update(ctx context.Context, in *Complete, opts ...grpc.CallOption) (*empty.Empty, error)
	Delete(ctx context.Context, in *Identifier, opts ...grpc.CallOption) (*empty.Empty, error)
}

type managerClient struct {
	cc *grpc.ClientConn
}

func NewManagerClient(cc *grpc.ClientConn) ManagerClient {
	return &managerClient{cc}
}

func (c *managerClient) Get(ctx context.Context, in *Identifier, opts ...grpc.CallOption) (*Complete, error) {
	out := new(Complete)
	err := c.cc.Invoke(ctx, "/serviceaccount_key.Manager/Get", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *managerClient) Create(ctx context.Context, in *Incomplete, opts ...grpc.CallOption) (*ClientKey, error) {
	out := new(ClientKey)
	err := c.cc.Invoke(ctx, "/serviceaccount_key.Manager/Create", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *managerClient) Update(ctx context.Context, in *Complete, opts ...grpc.CallOption) (*empty.Empty, error) {
	out := new(empty.Empty)
	err := c.cc.Invoke(ctx, "/serviceaccount_key.Manager/Update", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *managerClient) Delete(ctx context.Context, in *Identifier, opts ...grpc.CallOption) (*empty.Empty, error) {
	out := new(empty.Empty)
	err := c.cc.Invoke(ctx, "/serviceaccount_key.Manager/Delete", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// ManagerServer is the server API for Manager service.
type ManagerServer interface {
	Get(context.Context, *Identifier) (*Complete, error)
	Create(context.Context, *Incomplete) (*ClientKey, error)
	Update(context.Context, *Complete) (*empty.Empty, error)
	Delete(context.Context, *Identifier) (*empty.Empty, error)
}

func RegisterManagerServer(s *grpc.Server, srv ManagerServer) {
	s.RegisterService(&_Manager_serviceDesc, srv)
}

func _Manager_Get_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Identifier)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ManagerServer).Get(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/serviceaccount_key.Manager/Get",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ManagerServer).Get(ctx, req.(*Identifier))
	}
	return interceptor(ctx, in, info, handler)
}

func _Manager_Create_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Incomplete)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ManagerServer).Create(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/serviceaccount_key.Manager/Create",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ManagerServer).Create(ctx, req.(*Incomplete))
	}
	return interceptor(ctx, in, info, handler)
}

func _Manager_Update_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Complete)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ManagerServer).Update(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/serviceaccount_key.Manager/Update",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ManagerServer).Update(ctx, req.(*Complete))
	}
	return interceptor(ctx, in, info, handler)
}

func _Manager_Delete_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Identifier)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ManagerServer).Delete(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/serviceaccount_key.Manager/Delete",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ManagerServer).Delete(ctx, req.(*Identifier))
	}
	return interceptor(ctx, in, info, handler)
}

var _Manager_serviceDesc = grpc.ServiceDesc{
	ServiceName: "serviceaccount_key.Manager",
	HandlerType: (*ManagerServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Get",
			Handler:    _Manager_Get_Handler,
		},
		{
			MethodName: "Create",
			Handler:    _Manager_Create_Handler,
		},
		{
			MethodName: "Update",
			Handler:    _Manager_Update_Handler,
		},
		{
			MethodName: "Delete",
			Handler:    _Manager_Delete_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "manager.proto",
}
