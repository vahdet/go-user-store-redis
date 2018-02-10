// Code generated by protoc-gen-go. DO NOT EDIT.
// source: user.proto

/*
Package proto is a generated protocol buffer package.

It is generated from these files:
	user.proto

It has these top-level messages:
	UserId
	User
	CreateRequest
	UpdateRequest
*/
package proto

import proto1 "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"
import google_protobuf "github.com/golang/protobuf/ptypes/timestamp"

import (
	context "golang.org/x/net/context"
	grpc "google.golang.org/grpc"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto1.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto1.ProtoPackageIsVersion2 // please upgrade the proto package

type UserId struct {
	Value int64 `protobuf:"varint,1,opt,name=value" json:"value,omitempty"`
}

func (m *UserId) Reset()                    { *m = UserId{} }
func (m *UserId) String() string            { return proto1.CompactTextString(m) }
func (*UserId) ProtoMessage()               {}
func (*UserId) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

func (m *UserId) GetValue() int64 {
	if m != nil {
		return m.Value
	}
	return 0
}

type User struct {
	Id          int64                      `protobuf:"varint,1,opt,name=id" json:"id,omitempty"`
	Name        string                     `protobuf:"bytes,2,opt,name=name" json:"name,omitempty"`
	Email       string                     `protobuf:"bytes,3,opt,name=email" json:"email,omitempty"`
	Location    string                     `protobuf:"bytes,4,opt,name=location" json:"location,omitempty"`
	Language    string                     `protobuf:"bytes,5,opt,name=language" json:"language,omitempty"`
	Created     *google_protobuf.Timestamp `protobuf:"bytes,6,opt,name=created" json:"created,omitempty"`
	LastChanged *google_protobuf.Timestamp `protobuf:"bytes,7,opt,name=lastChanged" json:"lastChanged,omitempty"`
}

func (m *User) Reset()                    { *m = User{} }
func (m *User) String() string            { return proto1.CompactTextString(m) }
func (*User) ProtoMessage()               {}
func (*User) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{1} }

func (m *User) GetId() int64 {
	if m != nil {
		return m.Id
	}
	return 0
}

func (m *User) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *User) GetEmail() string {
	if m != nil {
		return m.Email
	}
	return ""
}

func (m *User) GetLocation() string {
	if m != nil {
		return m.Location
	}
	return ""
}

func (m *User) GetLanguage() string {
	if m != nil {
		return m.Language
	}
	return ""
}

func (m *User) GetCreated() *google_protobuf.Timestamp {
	if m != nil {
		return m.Created
	}
	return nil
}

func (m *User) GetLastChanged() *google_protobuf.Timestamp {
	if m != nil {
		return m.LastChanged
	}
	return nil
}

type CreateRequest struct {
	Name     string `protobuf:"bytes,1,opt,name=name" json:"name,omitempty"`
	Email    string `protobuf:"bytes,2,opt,name=email" json:"email,omitempty"`
	Location string `protobuf:"bytes,3,opt,name=location" json:"location,omitempty"`
	Language string `protobuf:"bytes,4,opt,name=language" json:"language,omitempty"`
}

func (m *CreateRequest) Reset()                    { *m = CreateRequest{} }
func (m *CreateRequest) String() string            { return proto1.CompactTextString(m) }
func (*CreateRequest) ProtoMessage()               {}
func (*CreateRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{2} }

func (m *CreateRequest) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *CreateRequest) GetEmail() string {
	if m != nil {
		return m.Email
	}
	return ""
}

func (m *CreateRequest) GetLocation() string {
	if m != nil {
		return m.Location
	}
	return ""
}

func (m *CreateRequest) GetLanguage() string {
	if m != nil {
		return m.Language
	}
	return ""
}

type UpdateRequest struct {
	Id       int64  `protobuf:"varint,1,opt,name=id" json:"id,omitempty"`
	Name     string `protobuf:"bytes,2,opt,name=name" json:"name,omitempty"`
	Email    string `protobuf:"bytes,3,opt,name=email" json:"email,omitempty"`
	Location string `protobuf:"bytes,4,opt,name=location" json:"location,omitempty"`
	Language string `protobuf:"bytes,5,opt,name=language" json:"language,omitempty"`
}

func (m *UpdateRequest) Reset()                    { *m = UpdateRequest{} }
func (m *UpdateRequest) String() string            { return proto1.CompactTextString(m) }
func (*UpdateRequest) ProtoMessage()               {}
func (*UpdateRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{3} }

func (m *UpdateRequest) GetId() int64 {
	if m != nil {
		return m.Id
	}
	return 0
}

func (m *UpdateRequest) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *UpdateRequest) GetEmail() string {
	if m != nil {
		return m.Email
	}
	return ""
}

func (m *UpdateRequest) GetLocation() string {
	if m != nil {
		return m.Location
	}
	return ""
}

func (m *UpdateRequest) GetLanguage() string {
	if m != nil {
		return m.Language
	}
	return ""
}

func init() {
	proto1.RegisterType((*UserId)(nil), "proto.UserId")
	proto1.RegisterType((*User)(nil), "proto.User")
	proto1.RegisterType((*CreateRequest)(nil), "proto.CreateRequest")
	proto1.RegisterType((*UpdateRequest)(nil), "proto.UpdateRequest")
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// Client API for UserService service

type UserServiceClient interface {
	Get(ctx context.Context, in *UserId, opts ...grpc.CallOption) (*User, error)
	Create(ctx context.Context, in *CreateRequest, opts ...grpc.CallOption) (*UserId, error)
	Update(ctx context.Context, in *UpdateRequest, opts ...grpc.CallOption) (*UserId, error)
	Delete(ctx context.Context, in *UserId, opts ...grpc.CallOption) (*UserId, error)
}

type userServiceClient struct {
	cc *grpc.ClientConn
}

func NewUserServiceClient(cc *grpc.ClientConn) UserServiceClient {
	return &userServiceClient{cc}
}

func (c *userServiceClient) Get(ctx context.Context, in *UserId, opts ...grpc.CallOption) (*User, error) {
	out := new(User)
	err := grpc.Invoke(ctx, "/proto.UserService/Get", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userServiceClient) Create(ctx context.Context, in *CreateRequest, opts ...grpc.CallOption) (*UserId, error) {
	out := new(UserId)
	err := grpc.Invoke(ctx, "/proto.UserService/Create", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userServiceClient) Update(ctx context.Context, in *UpdateRequest, opts ...grpc.CallOption) (*UserId, error) {
	out := new(UserId)
	err := grpc.Invoke(ctx, "/proto.UserService/Update", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userServiceClient) Delete(ctx context.Context, in *UserId, opts ...grpc.CallOption) (*UserId, error) {
	out := new(UserId)
	err := grpc.Invoke(ctx, "/proto.UserService/Delete", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for UserService service

type UserServiceServer interface {
	Get(context.Context, *UserId) (*User, error)
	Create(context.Context, *CreateRequest) (*UserId, error)
	Update(context.Context, *UpdateRequest) (*UserId, error)
	Delete(context.Context, *UserId) (*UserId, error)
}

func RegisterUserServiceServer(s *grpc.Server, srv UserServiceServer) {
	s.RegisterService(&_UserService_serviceDesc, srv)
}

func _UserService_Get_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UserId)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserServiceServer).Get(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/proto.UserService/Get",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserServiceServer).Get(ctx, req.(*UserId))
	}
	return interceptor(ctx, in, info, handler)
}

func _UserService_Create_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserServiceServer).Create(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/proto.UserService/Create",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserServiceServer).Create(ctx, req.(*CreateRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _UserService_Update_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserServiceServer).Update(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/proto.UserService/Update",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserServiceServer).Update(ctx, req.(*UpdateRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _UserService_Delete_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UserId)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserServiceServer).Delete(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/proto.UserService/Delete",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserServiceServer).Delete(ctx, req.(*UserId))
	}
	return interceptor(ctx, in, info, handler)
}

var _UserService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "proto.UserService",
	HandlerType: (*UserServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Get",
			Handler:    _UserService_Get_Handler,
		},
		{
			MethodName: "Create",
			Handler:    _UserService_Create_Handler,
		},
		{
			MethodName: "Update",
			Handler:    _UserService_Update_Handler,
		},
		{
			MethodName: "Delete",
			Handler:    _UserService_Delete_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "user.proto",
}

func init() { proto1.RegisterFile("user.proto", fileDescriptor0) }

var fileDescriptor0 = []byte{
	// 332 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xbc, 0x92, 0xcf, 0x4e, 0xb3, 0x40,
	0x14, 0xc5, 0xbf, 0x01, 0x4a, 0x3f, 0x2f, 0xa9, 0x8b, 0x9b, 0x2e, 0x26, 0x2c, 0xb4, 0xc1, 0x0d,
	0x2b, 0x9a, 0x54, 0x97, 0xee, 0x6a, 0x62, 0xdc, 0xa2, 0x7d, 0x80, 0x69, 0xb9, 0x22, 0xc9, 0xc0,
	0xb4, 0x30, 0x74, 0xef, 0x8b, 0xf9, 0x44, 0x3e, 0x84, 0xe9, 0x0c, 0x6d, 0x4a, 0xfd, 0xb7, 0x73,
	0x05, 0x87, 0x73, 0xce, 0x85, 0xdf, 0x65, 0x00, 0xda, 0x86, 0xea, 0x64, 0x5d, 0x2b, 0xad, 0x70,
	0x60, 0x2e, 0xe1, 0x65, 0xae, 0x54, 0x2e, 0x69, 0x6a, 0xd4, 0xb2, 0x7d, 0x9e, 0xea, 0xa2, 0xa4,
	0x46, 0x8b, 0x72, 0x6d, 0x73, 0xd1, 0x05, 0xf8, 0x8b, 0x86, 0xea, 0x87, 0x0c, 0xc7, 0x30, 0xd8,
	0x0a, 0xd9, 0x12, 0x67, 0x13, 0x16, 0xbb, 0xa9, 0x15, 0xd1, 0x3b, 0x03, 0x6f, 0x17, 0xc0, 0x73,
	0x70, 0x8a, 0xac, 0xf3, 0x9c, 0x22, 0x43, 0x04, 0xaf, 0x12, 0x25, 0x71, 0x67, 0xc2, 0xe2, 0xb3,
	0xd4, 0xdc, 0xef, 0x46, 0x50, 0x29, 0x0a, 0xc9, 0x5d, 0xf3, 0xd0, 0x0a, 0x0c, 0xe1, 0xbf, 0x54,
	0x2b, 0xa1, 0x0b, 0x55, 0x71, 0xcf, 0x18, 0x07, 0x6d, 0x3c, 0x51, 0xe5, 0xad, 0xc8, 0x89, 0x0f,
	0x3a, 0xaf, 0xd3, 0x78, 0x03, 0xc3, 0x55, 0x4d, 0x42, 0x53, 0xc6, 0xfd, 0x09, 0x8b, 0x83, 0x59,
	0x98, 0x58, 0x9a, 0x64, 0x4f, 0x93, 0x3c, 0xed, 0x69, 0xd2, 0x7d, 0x14, 0x6f, 0x21, 0x90, 0xa2,
	0xd1, 0xf3, 0x17, 0x51, 0xe5, 0x94, 0xf1, 0xe1, 0xaf, 0xcd, 0xe3, 0x78, 0xb4, 0x81, 0xd1, 0xdc,
	0x0c, 0x4a, 0x69, 0xd3, 0x52, 0xa3, 0x0f, 0x98, 0xec, 0x2b, 0x4c, 0xe7, 0x3b, 0x4c, 0xf7, 0x07,
	0x4c, 0xaf, 0x8f, 0x19, 0xbd, 0x32, 0x18, 0x2d, 0xd6, 0xd9, 0xd1, 0x3b, 0xff, 0x7c, 0xd5, 0xb3,
	0x37, 0x06, 0xc1, 0xee, 0x2f, 0x3f, 0x52, 0xbd, 0x2d, 0x56, 0x84, 0x57, 0xe0, 0xde, 0x93, 0xc6,
	0x91, 0xdd, 0x57, 0x62, 0x4f, 0x48, 0x18, 0x1c, 0xc9, 0xe8, 0x1f, 0x4e, 0xc1, 0xb7, 0xbb, 0xc2,
	0x71, 0x67, 0xf4, 0x56, 0x17, 0xf6, 0xdb, 0xb6, 0x60, 0x41, 0x0f, 0x85, 0x1e, 0xf7, 0xe7, 0x42,
	0x0c, 0xfe, 0x1d, 0x49, 0xd2, 0x74, 0xfa, 0x25, 0xa7, 0xc9, 0xa5, 0x6f, 0xf4, 0xf5, 0x47, 0x00,
	0x00, 0x00, 0xff, 0xff, 0x91, 0xf9, 0x41, 0xc6, 0x03, 0x03, 0x00, 0x00,
}
