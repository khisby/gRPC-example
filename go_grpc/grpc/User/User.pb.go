// Code generated by protoc-gen-go. DO NOT EDIT.
// source: User.proto

package User

import (
	context "context"
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
	emptypb "google.golang.org/protobuf/types/known/emptypb"
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

type User struct {
	UserId               string   `protobuf:"bytes,1,opt,name=userId,proto3" json:"userId,omitempty"`
	Username             string   `protobuf:"bytes,2,opt,name=username,proto3" json:"username,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *User) Reset()         { *m = User{} }
func (m *User) String() string { return proto.CompactTextString(m) }
func (*User) ProtoMessage()    {}
func (*User) Descriptor() ([]byte, []int) {
	return fileDescriptor_979821478719c248, []int{0}
}

func (m *User) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_User.Unmarshal(m, b)
}
func (m *User) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_User.Marshal(b, m, deterministic)
}
func (m *User) XXX_Merge(src proto.Message) {
	xxx_messageInfo_User.Merge(m, src)
}
func (m *User) XXX_Size() int {
	return xxx_messageInfo_User.Size(m)
}
func (m *User) XXX_DiscardUnknown() {
	xxx_messageInfo_User.DiscardUnknown(m)
}

var xxx_messageInfo_User proto.InternalMessageInfo

func (m *User) GetUserId() string {
	if m != nil {
		return m.UserId
	}
	return ""
}

func (m *User) GetUsername() string {
	if m != nil {
		return m.Username
	}
	return ""
}

type UsersResponse struct {
	Data                 []*User  `protobuf:"bytes,1,rep,name=data,proto3" json:"data,omitempty"`
	Message              string   `protobuf:"bytes,2,opt,name=message,proto3" json:"message,omitempty"`
	Status               string   `protobuf:"bytes,3,opt,name=status,proto3" json:"status,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *UsersResponse) Reset()         { *m = UsersResponse{} }
func (m *UsersResponse) String() string { return proto.CompactTextString(m) }
func (*UsersResponse) ProtoMessage()    {}
func (*UsersResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_979821478719c248, []int{1}
}

func (m *UsersResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_UsersResponse.Unmarshal(m, b)
}
func (m *UsersResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_UsersResponse.Marshal(b, m, deterministic)
}
func (m *UsersResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_UsersResponse.Merge(m, src)
}
func (m *UsersResponse) XXX_Size() int {
	return xxx_messageInfo_UsersResponse.Size(m)
}
func (m *UsersResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_UsersResponse.DiscardUnknown(m)
}

var xxx_messageInfo_UsersResponse proto.InternalMessageInfo

func (m *UsersResponse) GetData() []*User {
	if m != nil {
		return m.Data
	}
	return nil
}

func (m *UsersResponse) GetMessage() string {
	if m != nil {
		return m.Message
	}
	return ""
}

func (m *UsersResponse) GetStatus() string {
	if m != nil {
		return m.Status
	}
	return ""
}

type UserResponse struct {
	Data                 *User    `protobuf:"bytes,1,opt,name=data,proto3" json:"data,omitempty"`
	Message              string   `protobuf:"bytes,2,opt,name=message,proto3" json:"message,omitempty"`
	Status               string   `protobuf:"bytes,3,opt,name=status,proto3" json:"status,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *UserResponse) Reset()         { *m = UserResponse{} }
func (m *UserResponse) String() string { return proto.CompactTextString(m) }
func (*UserResponse) ProtoMessage()    {}
func (*UserResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_979821478719c248, []int{2}
}

func (m *UserResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_UserResponse.Unmarshal(m, b)
}
func (m *UserResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_UserResponse.Marshal(b, m, deterministic)
}
func (m *UserResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_UserResponse.Merge(m, src)
}
func (m *UserResponse) XXX_Size() int {
	return xxx_messageInfo_UserResponse.Size(m)
}
func (m *UserResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_UserResponse.DiscardUnknown(m)
}

var xxx_messageInfo_UserResponse proto.InternalMessageInfo

func (m *UserResponse) GetData() *User {
	if m != nil {
		return m.Data
	}
	return nil
}

func (m *UserResponse) GetMessage() string {
	if m != nil {
		return m.Message
	}
	return ""
}

func (m *UserResponse) GetStatus() string {
	if m != nil {
		return m.Status
	}
	return ""
}

type GetUserRequest struct {
	UserId               string   `protobuf:"bytes,1,opt,name=userId,proto3" json:"userId,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *GetUserRequest) Reset()         { *m = GetUserRequest{} }
func (m *GetUserRequest) String() string { return proto.CompactTextString(m) }
func (*GetUserRequest) ProtoMessage()    {}
func (*GetUserRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_979821478719c248, []int{3}
}

func (m *GetUserRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetUserRequest.Unmarshal(m, b)
}
func (m *GetUserRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetUserRequest.Marshal(b, m, deterministic)
}
func (m *GetUserRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetUserRequest.Merge(m, src)
}
func (m *GetUserRequest) XXX_Size() int {
	return xxx_messageInfo_GetUserRequest.Size(m)
}
func (m *GetUserRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_GetUserRequest.DiscardUnknown(m)
}

var xxx_messageInfo_GetUserRequest proto.InternalMessageInfo

func (m *GetUserRequest) GetUserId() string {
	if m != nil {
		return m.UserId
	}
	return ""
}

func init() {
	proto.RegisterType((*User)(nil), "User.User")
	proto.RegisterType((*UsersResponse)(nil), "User.UsersResponse")
	proto.RegisterType((*UserResponse)(nil), "User.UserResponse")
	proto.RegisterType((*GetUserRequest)(nil), "User.GetUserRequest")
}

func init() {
	proto.RegisterFile("User.proto", fileDescriptor_979821478719c248)
}

var fileDescriptor_979821478719c248 = []byte{
	// 269 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xac, 0x50, 0xc1, 0x4a, 0x03, 0x31,
	0x10, 0x35, 0xb6, 0xb4, 0x75, 0xaa, 0x1e, 0x46, 0x29, 0x61, 0x05, 0x29, 0x7b, 0xda, 0x53, 0x0a,
	0x15, 0x0f, 0x7a, 0x15, 0x11, 0xaf, 0x2b, 0xde, 0x4d, 0xed, 0xb8, 0x08, 0x6e, 0xb3, 0x66, 0x12,
	0xc1, 0xaf, 0xf2, 0x17, 0x25, 0xc9, 0x6e, 0x5b, 0x51, 0x6f, 0xbd, 0x2c, 0xf3, 0x66, 0xdf, 0xcb,
	0x7b, 0xf3, 0x00, 0x1e, 0x99, 0xac, 0x6a, 0xac, 0x71, 0x06, 0xfb, 0x61, 0xce, 0xce, 0x2a, 0x63,
	0xaa, 0x37, 0x9a, 0xc5, 0xdd, 0xc2, 0xbf, 0xcc, 0xa8, 0x6e, 0xdc, 0x67, 0xa2, 0xe4, 0xd7, 0x10,
	0x49, 0x38, 0x81, 0x81, 0x67, 0xb2, 0xf7, 0x4b, 0x29, 0xa6, 0xa2, 0x38, 0x28, 0x5b, 0x84, 0x19,
	0x8c, 0xc2, 0xb4, 0xd2, 0x35, 0xc9, 0xfd, 0xf8, 0x67, 0x8d, 0x73, 0x0d, 0x47, 0x41, 0xcb, 0x25,
	0x71, 0x63, 0x56, 0x4c, 0x78, 0x0e, 0xfd, 0xa5, 0x76, 0x5a, 0x8a, 0x69, 0xaf, 0x18, 0xcf, 0x41,
	0xc5, 0x28, 0xe1, 0x53, 0xc6, 0x3d, 0x4a, 0x18, 0xd6, 0xc4, 0xac, 0xab, 0xee, 0xad, 0x0e, 0x06,
	0x7b, 0x76, 0xda, 0x79, 0x96, 0xbd, 0x64, 0x9f, 0x50, 0xfe, 0x04, 0x87, 0x51, 0xff, 0xdb, 0x41,
	0xec, 0xc8, 0xa1, 0x80, 0xe3, 0x3b, 0x72, 0xc9, 0xe4, 0xdd, 0x13, 0xbb, 0xff, 0xaa, 0x98, 0x7f,
	0x09, 0x18, 0x07, 0xde, 0x03, 0xd9, 0x8f, 0xd7, 0x67, 0xc2, 0x2b, 0x18, 0xb5, 0x4a, 0xc6, 0x89,
	0x4a, 0x25, 0xab, 0xae, 0x64, 0x75, 0x1b, 0x4a, 0xce, 0x4e, 0x36, 0x09, 0xd7, 0x35, 0xe5, 0x7b,
	0x78, 0x09, 0xc3, 0x56, 0x8a, 0xa7, 0x89, 0xf1, 0x33, 0x43, 0x86, 0x5b, 0x97, 0x6d, 0x64, 0x0a,
	0xe0, 0xc6, 0x92, 0x76, 0x14, 0x95, 0x5b, 0xd7, 0xff, 0xcd, 0x5f, 0x0c, 0x62, 0x9a, 0x8b, 0xef,
	0x00, 0x00, 0x00, 0xff, 0xff, 0xde, 0x6a, 0xd2, 0xf6, 0x14, 0x02, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConnInterface

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion6

// UserServiceClient is the client API for UserService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type UserServiceClient interface {
	GetUsers(ctx context.Context, in *emptypb.Empty, opts ...grpc.CallOption) (*UsersResponse, error)
	GetUser(ctx context.Context, in *GetUserRequest, opts ...grpc.CallOption) (*UserResponse, error)
	CreateUser(ctx context.Context, in *User, opts ...grpc.CallOption) (*UserResponse, error)
}

type userServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewUserServiceClient(cc grpc.ClientConnInterface) UserServiceClient {
	return &userServiceClient{cc}
}

func (c *userServiceClient) GetUsers(ctx context.Context, in *emptypb.Empty, opts ...grpc.CallOption) (*UsersResponse, error) {
	out := new(UsersResponse)
	err := c.cc.Invoke(ctx, "/User.UserService/GetUsers", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userServiceClient) GetUser(ctx context.Context, in *GetUserRequest, opts ...grpc.CallOption) (*UserResponse, error) {
	out := new(UserResponse)
	err := c.cc.Invoke(ctx, "/User.UserService/GetUser", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userServiceClient) CreateUser(ctx context.Context, in *User, opts ...grpc.CallOption) (*UserResponse, error) {
	out := new(UserResponse)
	err := c.cc.Invoke(ctx, "/User.UserService/CreateUser", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// UserServiceServer is the server API for UserService service.
type UserServiceServer interface {
	GetUsers(context.Context, *emptypb.Empty) (*UsersResponse, error)
	GetUser(context.Context, *GetUserRequest) (*UserResponse, error)
	CreateUser(context.Context, *User) (*UserResponse, error)
}

// UnimplementedUserServiceServer can be embedded to have forward compatible implementations.
type UnimplementedUserServiceServer struct {
}

func (*UnimplementedUserServiceServer) GetUsers(ctx context.Context, req *emptypb.Empty) (*UsersResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetUsers not implemented")
}
func (*UnimplementedUserServiceServer) GetUser(ctx context.Context, req *GetUserRequest) (*UserResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetUser not implemented")
}
func (*UnimplementedUserServiceServer) CreateUser(ctx context.Context, req *User) (*UserResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateUser not implemented")
}

func RegisterUserServiceServer(s *grpc.Server, srv UserServiceServer) {
	s.RegisterService(&_UserService_serviceDesc, srv)
}

func _UserService_GetUsers_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(emptypb.Empty)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserServiceServer).GetUsers(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/User.UserService/GetUsers",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserServiceServer).GetUsers(ctx, req.(*emptypb.Empty))
	}
	return interceptor(ctx, in, info, handler)
}

func _UserService_GetUser_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetUserRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserServiceServer).GetUser(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/User.UserService/GetUser",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserServiceServer).GetUser(ctx, req.(*GetUserRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _UserService_CreateUser_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(User)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserServiceServer).CreateUser(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/User.UserService/CreateUser",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserServiceServer).CreateUser(ctx, req.(*User))
	}
	return interceptor(ctx, in, info, handler)
}

var _UserService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "User.UserService",
	HandlerType: (*UserServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetUsers",
			Handler:    _UserService_GetUsers_Handler,
		},
		{
			MethodName: "GetUser",
			Handler:    _UserService_GetUser_Handler,
		},
		{
			MethodName: "CreateUser",
			Handler:    _UserService_CreateUser_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "User.proto",
}