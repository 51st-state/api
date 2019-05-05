// Code generated by protoc-gen-go. DO NOT EDIT.
// source: control.proto

package rbac

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

type Rule struct {
	Rule                 string   `protobuf:"bytes,1,opt,name=Rule,proto3" json:"Rule,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Rule) Reset()         { *m = Rule{} }
func (m *Rule) String() string { return proto.CompactTextString(m) }
func (*Rule) ProtoMessage()    {}
func (*Rule) Descriptor() ([]byte, []int) {
	return fileDescriptor_0c5120591600887d, []int{0}
}

func (m *Rule) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Rule.Unmarshal(m, b)
}
func (m *Rule) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Rule.Marshal(b, m, deterministic)
}
func (m *Rule) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Rule.Merge(m, src)
}
func (m *Rule) XXX_Size() int {
	return xxx_messageInfo_Rule.Size(m)
}
func (m *Rule) XXX_DiscardUnknown() {
	xxx_messageInfo_Rule.DiscardUnknown(m)
}

var xxx_messageInfo_Rule proto.InternalMessageInfo

func (m *Rule) GetRule() string {
	if m != nil {
		return m.Rule
	}
	return ""
}

type RoleID struct {
	ID                   string   `protobuf:"bytes,1,opt,name=ID,proto3" json:"ID,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *RoleID) Reset()         { *m = RoleID{} }
func (m *RoleID) String() string { return proto.CompactTextString(m) }
func (*RoleID) ProtoMessage()    {}
func (*RoleID) Descriptor() ([]byte, []int) {
	return fileDescriptor_0c5120591600887d, []int{1}
}

func (m *RoleID) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_RoleID.Unmarshal(m, b)
}
func (m *RoleID) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_RoleID.Marshal(b, m, deterministic)
}
func (m *RoleID) XXX_Merge(src proto.Message) {
	xxx_messageInfo_RoleID.Merge(m, src)
}
func (m *RoleID) XXX_Size() int {
	return xxx_messageInfo_RoleID.Size(m)
}
func (m *RoleID) XXX_DiscardUnknown() {
	xxx_messageInfo_RoleID.DiscardUnknown(m)
}

var xxx_messageInfo_RoleID proto.InternalMessageInfo

func (m *RoleID) GetID() string {
	if m != nil {
		return m.ID
	}
	return ""
}

type RoleRules struct {
	Rules                []string `protobuf:"bytes,1,rep,name=Rules,proto3" json:"Rules,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *RoleRules) Reset()         { *m = RoleRules{} }
func (m *RoleRules) String() string { return proto.CompactTextString(m) }
func (*RoleRules) ProtoMessage()    {}
func (*RoleRules) Descriptor() ([]byte, []int) {
	return fileDescriptor_0c5120591600887d, []int{2}
}

func (m *RoleRules) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_RoleRules.Unmarshal(m, b)
}
func (m *RoleRules) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_RoleRules.Marshal(b, m, deterministic)
}
func (m *RoleRules) XXX_Merge(src proto.Message) {
	xxx_messageInfo_RoleRules.Merge(m, src)
}
func (m *RoleRules) XXX_Size() int {
	return xxx_messageInfo_RoleRules.Size(m)
}
func (m *RoleRules) XXX_DiscardUnknown() {
	xxx_messageInfo_RoleRules.DiscardUnknown(m)
}

var xxx_messageInfo_RoleRules proto.InternalMessageInfo

func (m *RoleRules) GetRules() []string {
	if m != nil {
		return m.Rules
	}
	return nil
}

type AccountID struct {
	ID                   string   `protobuf:"bytes,1,opt,name=ID,proto3" json:"ID,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *AccountID) Reset()         { *m = AccountID{} }
func (m *AccountID) String() string { return proto.CompactTextString(m) }
func (*AccountID) ProtoMessage()    {}
func (*AccountID) Descriptor() ([]byte, []int) {
	return fileDescriptor_0c5120591600887d, []int{3}
}

func (m *AccountID) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_AccountID.Unmarshal(m, b)
}
func (m *AccountID) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_AccountID.Marshal(b, m, deterministic)
}
func (m *AccountID) XXX_Merge(src proto.Message) {
	xxx_messageInfo_AccountID.Merge(m, src)
}
func (m *AccountID) XXX_Size() int {
	return xxx_messageInfo_AccountID.Size(m)
}
func (m *AccountID) XXX_DiscardUnknown() {
	xxx_messageInfo_AccountID.DiscardUnknown(m)
}

var xxx_messageInfo_AccountID proto.InternalMessageInfo

func (m *AccountID) GetID() string {
	if m != nil {
		return m.ID
	}
	return ""
}

type AccountRoles struct {
	RoleIDs              []string `protobuf:"bytes,1,rep,name=RoleIDs,proto3" json:"RoleIDs,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *AccountRoles) Reset()         { *m = AccountRoles{} }
func (m *AccountRoles) String() string { return proto.CompactTextString(m) }
func (*AccountRoles) ProtoMessage()    {}
func (*AccountRoles) Descriptor() ([]byte, []int) {
	return fileDescriptor_0c5120591600887d, []int{4}
}

func (m *AccountRoles) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_AccountRoles.Unmarshal(m, b)
}
func (m *AccountRoles) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_AccountRoles.Marshal(b, m, deterministic)
}
func (m *AccountRoles) XXX_Merge(src proto.Message) {
	xxx_messageInfo_AccountRoles.Merge(m, src)
}
func (m *AccountRoles) XXX_Size() int {
	return xxx_messageInfo_AccountRoles.Size(m)
}
func (m *AccountRoles) XXX_DiscardUnknown() {
	xxx_messageInfo_AccountRoles.DiscardUnknown(m)
}

var xxx_messageInfo_AccountRoles proto.InternalMessageInfo

func (m *AccountRoles) GetRoleIDs() []string {
	if m != nil {
		return m.RoleIDs
	}
	return nil
}

type SetRoleRulesRequest struct {
	RoleID               *RoleID    `protobuf:"bytes,1,opt,name=RoleID,proto3" json:"RoleID,omitempty"`
	RoleRules            *RoleRules `protobuf:"bytes,2,opt,name=RoleRules,proto3" json:"RoleRules,omitempty"`
	XXX_NoUnkeyedLiteral struct{}   `json:"-"`
	XXX_unrecognized     []byte     `json:"-"`
	XXX_sizecache        int32      `json:"-"`
}

func (m *SetRoleRulesRequest) Reset()         { *m = SetRoleRulesRequest{} }
func (m *SetRoleRulesRequest) String() string { return proto.CompactTextString(m) }
func (*SetRoleRulesRequest) ProtoMessage()    {}
func (*SetRoleRulesRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_0c5120591600887d, []int{5}
}

func (m *SetRoleRulesRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_SetRoleRulesRequest.Unmarshal(m, b)
}
func (m *SetRoleRulesRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_SetRoleRulesRequest.Marshal(b, m, deterministic)
}
func (m *SetRoleRulesRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SetRoleRulesRequest.Merge(m, src)
}
func (m *SetRoleRulesRequest) XXX_Size() int {
	return xxx_messageInfo_SetRoleRulesRequest.Size(m)
}
func (m *SetRoleRulesRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_SetRoleRulesRequest.DiscardUnknown(m)
}

var xxx_messageInfo_SetRoleRulesRequest proto.InternalMessageInfo

func (m *SetRoleRulesRequest) GetRoleID() *RoleID {
	if m != nil {
		return m.RoleID
	}
	return nil
}

func (m *SetRoleRulesRequest) GetRoleRules() *RoleRules {
	if m != nil {
		return m.RoleRules
	}
	return nil
}

type SetAccountRolesRequest struct {
	AccountID            *AccountID    `protobuf:"bytes,1,opt,name=AccountID,proto3" json:"AccountID,omitempty"`
	AccountRoles         *AccountRoles `protobuf:"bytes,2,opt,name=AccountRoles,proto3" json:"AccountRoles,omitempty"`
	XXX_NoUnkeyedLiteral struct{}      `json:"-"`
	XXX_unrecognized     []byte        `json:"-"`
	XXX_sizecache        int32         `json:"-"`
}

func (m *SetAccountRolesRequest) Reset()         { *m = SetAccountRolesRequest{} }
func (m *SetAccountRolesRequest) String() string { return proto.CompactTextString(m) }
func (*SetAccountRolesRequest) ProtoMessage()    {}
func (*SetAccountRolesRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_0c5120591600887d, []int{6}
}

func (m *SetAccountRolesRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_SetAccountRolesRequest.Unmarshal(m, b)
}
func (m *SetAccountRolesRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_SetAccountRolesRequest.Marshal(b, m, deterministic)
}
func (m *SetAccountRolesRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SetAccountRolesRequest.Merge(m, src)
}
func (m *SetAccountRolesRequest) XXX_Size() int {
	return xxx_messageInfo_SetAccountRolesRequest.Size(m)
}
func (m *SetAccountRolesRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_SetAccountRolesRequest.DiscardUnknown(m)
}

var xxx_messageInfo_SetAccountRolesRequest proto.InternalMessageInfo

func (m *SetAccountRolesRequest) GetAccountID() *AccountID {
	if m != nil {
		return m.AccountID
	}
	return nil
}

func (m *SetAccountRolesRequest) GetAccountRoles() *AccountRoles {
	if m != nil {
		return m.AccountRoles
	}
	return nil
}

type IsAccountAllowedRequest struct {
	AccountID            *AccountID `protobuf:"bytes,1,opt,name=AccountID,proto3" json:"AccountID,omitempty"`
	Rule                 *Rule      `protobuf:"bytes,2,opt,name=Rule,proto3" json:"Rule,omitempty"`
	XXX_NoUnkeyedLiteral struct{}   `json:"-"`
	XXX_unrecognized     []byte     `json:"-"`
	XXX_sizecache        int32      `json:"-"`
}

func (m *IsAccountAllowedRequest) Reset()         { *m = IsAccountAllowedRequest{} }
func (m *IsAccountAllowedRequest) String() string { return proto.CompactTextString(m) }
func (*IsAccountAllowedRequest) ProtoMessage()    {}
func (*IsAccountAllowedRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_0c5120591600887d, []int{7}
}

func (m *IsAccountAllowedRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_IsAccountAllowedRequest.Unmarshal(m, b)
}
func (m *IsAccountAllowedRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_IsAccountAllowedRequest.Marshal(b, m, deterministic)
}
func (m *IsAccountAllowedRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_IsAccountAllowedRequest.Merge(m, src)
}
func (m *IsAccountAllowedRequest) XXX_Size() int {
	return xxx_messageInfo_IsAccountAllowedRequest.Size(m)
}
func (m *IsAccountAllowedRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_IsAccountAllowedRequest.DiscardUnknown(m)
}

var xxx_messageInfo_IsAccountAllowedRequest proto.InternalMessageInfo

func (m *IsAccountAllowedRequest) GetAccountID() *AccountID {
	if m != nil {
		return m.AccountID
	}
	return nil
}

func (m *IsAccountAllowedRequest) GetRule() *Rule {
	if m != nil {
		return m.Rule
	}
	return nil
}

type IsAccountAllowedResponse struct {
	Allowed              bool     `protobuf:"varint,1,opt,name=Allowed,proto3" json:"Allowed,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *IsAccountAllowedResponse) Reset()         { *m = IsAccountAllowedResponse{} }
func (m *IsAccountAllowedResponse) String() string { return proto.CompactTextString(m) }
func (*IsAccountAllowedResponse) ProtoMessage()    {}
func (*IsAccountAllowedResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_0c5120591600887d, []int{8}
}

func (m *IsAccountAllowedResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_IsAccountAllowedResponse.Unmarshal(m, b)
}
func (m *IsAccountAllowedResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_IsAccountAllowedResponse.Marshal(b, m, deterministic)
}
func (m *IsAccountAllowedResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_IsAccountAllowedResponse.Merge(m, src)
}
func (m *IsAccountAllowedResponse) XXX_Size() int {
	return xxx_messageInfo_IsAccountAllowedResponse.Size(m)
}
func (m *IsAccountAllowedResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_IsAccountAllowedResponse.DiscardUnknown(m)
}

var xxx_messageInfo_IsAccountAllowedResponse proto.InternalMessageInfo

func (m *IsAccountAllowedResponse) GetAllowed() bool {
	if m != nil {
		return m.Allowed
	}
	return false
}

func init() {
	proto.RegisterType((*Rule)(nil), "rbac.Rule")
	proto.RegisterType((*RoleID)(nil), "rbac.RoleID")
	proto.RegisterType((*RoleRules)(nil), "rbac.RoleRules")
	proto.RegisterType((*AccountID)(nil), "rbac.AccountID")
	proto.RegisterType((*AccountRoles)(nil), "rbac.AccountRoles")
	proto.RegisterType((*SetRoleRulesRequest)(nil), "rbac.SetRoleRulesRequest")
	proto.RegisterType((*SetAccountRolesRequest)(nil), "rbac.SetAccountRolesRequest")
	proto.RegisterType((*IsAccountAllowedRequest)(nil), "rbac.IsAccountAllowedRequest")
	proto.RegisterType((*IsAccountAllowedResponse)(nil), "rbac.IsAccountAllowedResponse")
}

func init() { proto.RegisterFile("control.proto", fileDescriptor_0c5120591600887d) }

var fileDescriptor_0c5120591600887d = []byte{
	// 394 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x9c, 0x52, 0x4d, 0x4f, 0xea, 0x40,
	0x14, 0x2d, 0x7d, 0x3c, 0x78, 0xbd, 0xaf, 0xef, 0x61, 0xae, 0x06, 0x6b, 0x51, 0x82, 0x13, 0x17,
	0x6c, 0x28, 0x09, 0x1a, 0xe3, 0x96, 0x80, 0x21, 0xdd, 0x0e, 0xbf, 0x40, 0xea, 0x88, 0x31, 0x95,
	0x41, 0x3a, 0x8d, 0x71, 0xe5, 0xef, 0xf4, 0xdf, 0x98, 0xce, 0x07, 0xb4, 0xb5, 0xb8, 0x70, 0xd5,
	0xce, 0x9d, 0x33, 0xe7, 0x9c, 0x7b, 0xcf, 0x85, 0x7f, 0x11, 0x5f, 0x89, 0x0d, 0x8f, 0x83, 0xf5,
	0x86, 0x0b, 0x8e, 0xf5, 0xcd, 0xe2, 0x2e, 0xf2, 0x3b, 0x4b, 0xce, 0x97, 0x31, 0x1b, 0xca, 0xda,
	0x22, 0x7d, 0x18, 0xb2, 0xe7, 0xb5, 0x78, 0x53, 0x10, 0xe2, 0x43, 0x9d, 0xa6, 0x31, 0x43, 0x54,
	0x5f, 0xaf, 0xd6, 0xab, 0xf5, 0x1d, 0x2a, 0xff, 0x89, 0x07, 0x0d, 0xca, 0x63, 0x16, 0x4e, 0xf1,
	0x3f, 0xd8, 0xe1, 0x54, 0xdf, 0xd9, 0xe1, 0x94, 0x9c, 0x83, 0x93, 0xdd, 0x64, 0xa8, 0x04, 0x8f,
	0xe0, 0xb7, 0xfc, 0xf1, 0x6a, 0xbd, 0x5f, 0x7d, 0x87, 0xaa, 0x03, 0xe9, 0x80, 0x33, 0x8e, 0x22,
	0x9e, 0xae, 0x44, 0xc5, 0xfb, 0x3e, 0xb8, 0xfa, 0x32, 0xa3, 0x49, 0xd0, 0x83, 0xa6, 0x52, 0x32,
	0x24, 0xe6, 0x48, 0x9e, 0xe0, 0x70, 0xce, 0xc4, 0x56, 0x8c, 0xb2, 0x97, 0x94, 0x25, 0x02, 0x2f,
	0x8c, 0x35, 0x49, 0xfa, 0x77, 0xe4, 0x06, 0x59, 0xab, 0x81, 0xaa, 0x51, 0x63, 0x7b, 0x90, 0xb3,
	0xe9, 0xd9, 0x12, 0xd8, 0xda, 0x01, 0x15, 0xe1, 0x0e, 0x41, 0xde, 0xa1, 0x3d, 0x67, 0x22, 0x6f,
	0xcc, 0xc8, 0x0d, 0x72, 0xcd, 0x68, 0x45, 0x4d, 0xb4, 0x2d, 0xd3, 0x5c, 0xbb, 0xd7, 0xc5, 0xf6,
	0xb4, 0x34, 0x16, 0x5e, 0x28, 0xfe, 0x02, 0x8e, 0x3c, 0xc2, 0x71, 0x98, 0xe8, 0xca, 0x38, 0x8e,
	0xf9, 0x2b, 0xbb, 0xff, 0xa1, 0x83, 0xae, 0x8e, 0x53, 0x29, 0x83, 0x6e, 0x3a, 0x8d, 0x99, 0x8e,
	0xf6, 0x0a, 0xbc, 0xaf, 0x4a, 0xc9, 0x9a, 0xaf, 0x12, 0x96, 0x85, 0xa1, 0x4b, 0x52, 0xe8, 0x0f,
	0x35, 0xc7, 0xd1, 0x87, 0x0d, 0xcd, 0x89, 0xda, 0x30, 0x1c, 0x82, 0x3b, 0xcb, 0x05, 0x83, 0x85,
	0x04, 0xfc, 0xf2, 0x98, 0x89, 0x85, 0x13, 0x70, 0xf3, 0x49, 0xe2, 0x89, 0x82, 0x54, 0xa4, 0xeb,
	0xb7, 0x03, 0xb5, 0xb2, 0x81, 0x59, 0xd9, 0xe0, 0x36, 0x5b, 0x59, 0x62, 0xe1, 0x0d, 0xb4, 0x66,
	0xc5, 0x88, 0xb0, 0x3c, 0x06, 0xbf, 0x62, 0xce, 0xc4, 0xc2, 0x10, 0x5a, 0xa5, 0x70, 0xf1, 0x74,
	0xeb, 0xa0, 0x22, 0xf3, 0x6f, 0x4c, 0xcc, 0xe1, 0xa0, 0x3c, 0x3c, 0x3c, 0x53, 0x5c, 0x7b, 0xe2,
	0xf3, 0xbb, 0xfb, 0xae, 0xd5, 0xcc, 0x89, 0xb5, 0x68, 0x48, 0x99, 0xcb, 0xcf, 0x00, 0x00, 0x00,
	0xff, 0xff, 0x31, 0x65, 0x58, 0x39, 0xc3, 0x03, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// ControlClient is the client API for Control service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type ControlClient interface {
	GetRoleRules(ctx context.Context, in *RoleID, opts ...grpc.CallOption) (*RoleRules, error)
	SetRoleRules(ctx context.Context, in *SetRoleRulesRequest, opts ...grpc.CallOption) (*empty.Empty, error)
	GetAccountRoles(ctx context.Context, in *AccountID, opts ...grpc.CallOption) (*AccountRoles, error)
	SetAccountRoles(ctx context.Context, in *SetAccountRolesRequest, opts ...grpc.CallOption) (*empty.Empty, error)
	IsAccountAllowed(ctx context.Context, in *IsAccountAllowedRequest, opts ...grpc.CallOption) (*IsAccountAllowedResponse, error)
}

type controlClient struct {
	cc *grpc.ClientConn
}

func NewControlClient(cc *grpc.ClientConn) ControlClient {
	return &controlClient{cc}
}

func (c *controlClient) GetRoleRules(ctx context.Context, in *RoleID, opts ...grpc.CallOption) (*RoleRules, error) {
	out := new(RoleRules)
	err := c.cc.Invoke(ctx, "/rbac.Control/GetRoleRules", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *controlClient) SetRoleRules(ctx context.Context, in *SetRoleRulesRequest, opts ...grpc.CallOption) (*empty.Empty, error) {
	out := new(empty.Empty)
	err := c.cc.Invoke(ctx, "/rbac.Control/SetRoleRules", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *controlClient) GetAccountRoles(ctx context.Context, in *AccountID, opts ...grpc.CallOption) (*AccountRoles, error) {
	out := new(AccountRoles)
	err := c.cc.Invoke(ctx, "/rbac.Control/GetAccountRoles", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *controlClient) SetAccountRoles(ctx context.Context, in *SetAccountRolesRequest, opts ...grpc.CallOption) (*empty.Empty, error) {
	out := new(empty.Empty)
	err := c.cc.Invoke(ctx, "/rbac.Control/SetAccountRoles", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *controlClient) IsAccountAllowed(ctx context.Context, in *IsAccountAllowedRequest, opts ...grpc.CallOption) (*IsAccountAllowedResponse, error) {
	out := new(IsAccountAllowedResponse)
	err := c.cc.Invoke(ctx, "/rbac.Control/IsAccountAllowed", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// ControlServer is the server API for Control service.
type ControlServer interface {
	GetRoleRules(context.Context, *RoleID) (*RoleRules, error)
	SetRoleRules(context.Context, *SetRoleRulesRequest) (*empty.Empty, error)
	GetAccountRoles(context.Context, *AccountID) (*AccountRoles, error)
	SetAccountRoles(context.Context, *SetAccountRolesRequest) (*empty.Empty, error)
	IsAccountAllowed(context.Context, *IsAccountAllowedRequest) (*IsAccountAllowedResponse, error)
}

func RegisterControlServer(s *grpc.Server, srv ControlServer) {
	s.RegisterService(&_Control_serviceDesc, srv)
}

func _Control_GetRoleRules_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(RoleID)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ControlServer).GetRoleRules(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/rbac.Control/GetRoleRules",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ControlServer).GetRoleRules(ctx, req.(*RoleID))
	}
	return interceptor(ctx, in, info, handler)
}

func _Control_SetRoleRules_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SetRoleRulesRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ControlServer).SetRoleRules(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/rbac.Control/SetRoleRules",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ControlServer).SetRoleRules(ctx, req.(*SetRoleRulesRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Control_GetAccountRoles_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(AccountID)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ControlServer).GetAccountRoles(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/rbac.Control/GetAccountRoles",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ControlServer).GetAccountRoles(ctx, req.(*AccountID))
	}
	return interceptor(ctx, in, info, handler)
}

func _Control_SetAccountRoles_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SetAccountRolesRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ControlServer).SetAccountRoles(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/rbac.Control/SetAccountRoles",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ControlServer).SetAccountRoles(ctx, req.(*SetAccountRolesRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Control_IsAccountAllowed_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(IsAccountAllowedRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ControlServer).IsAccountAllowed(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/rbac.Control/IsAccountAllowed",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ControlServer).IsAccountAllowed(ctx, req.(*IsAccountAllowedRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _Control_serviceDesc = grpc.ServiceDesc{
	ServiceName: "rbac.Control",
	HandlerType: (*ControlServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetRoleRules",
			Handler:    _Control_GetRoleRules_Handler,
		},
		{
			MethodName: "SetRoleRules",
			Handler:    _Control_SetRoleRules_Handler,
		},
		{
			MethodName: "GetAccountRoles",
			Handler:    _Control_GetAccountRoles_Handler,
		},
		{
			MethodName: "SetAccountRoles",
			Handler:    _Control_SetAccountRoles_Handler,
		},
		{
			MethodName: "IsAccountAllowed",
			Handler:    _Control_IsAccountAllowed_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "control.proto",
}
