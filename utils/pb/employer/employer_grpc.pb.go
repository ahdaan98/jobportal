// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.5.1
// - protoc             v5.27.1
// source: utils/pb/employer/employer.proto

package __

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.64.0 or later.
const _ = grpc.SupportPackageIsVersion9

const (
	Employer_CreateEmployer_FullMethodName = "/employer.Employer/CreateEmployer"
	Employer_LoginEmployer_FullMethodName  = "/employer.Employer/LoginEmployer"
)

// EmployerClient is the client API for Employer service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type EmployerClient interface {
	CreateEmployer(ctx context.Context, in *CreateEmployerReq, opts ...grpc.CallOption) (*EmployerRes, error)
	LoginEmployer(ctx context.Context, in *EmpLoginReq, opts ...grpc.CallOption) (*EmployerRes, error)
}

type employerClient struct {
	cc grpc.ClientConnInterface
}

func NewEmployerClient(cc grpc.ClientConnInterface) EmployerClient {
	return &employerClient{cc}
}

func (c *employerClient) CreateEmployer(ctx context.Context, in *CreateEmployerReq, opts ...grpc.CallOption) (*EmployerRes, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(EmployerRes)
	err := c.cc.Invoke(ctx, Employer_CreateEmployer_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *employerClient) LoginEmployer(ctx context.Context, in *EmpLoginReq, opts ...grpc.CallOption) (*EmployerRes, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(EmployerRes)
	err := c.cc.Invoke(ctx, Employer_LoginEmployer_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// EmployerServer is the server API for Employer service.
// All implementations must embed UnimplementedEmployerServer
// for forward compatibility.
type EmployerServer interface {
	CreateEmployer(context.Context, *CreateEmployerReq) (*EmployerRes, error)
	LoginEmployer(context.Context, *EmpLoginReq) (*EmployerRes, error)
	mustEmbedUnimplementedEmployerServer()
}

// UnimplementedEmployerServer must be embedded to have
// forward compatible implementations.
//
// NOTE: this should be embedded by value instead of pointer to avoid a nil
// pointer dereference when methods are called.
type UnimplementedEmployerServer struct{}

func (UnimplementedEmployerServer) CreateEmployer(context.Context, *CreateEmployerReq) (*EmployerRes, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateEmployer not implemented")
}
func (UnimplementedEmployerServer) LoginEmployer(context.Context, *EmpLoginReq) (*EmployerRes, error) {
	return nil, status.Errorf(codes.Unimplemented, "method LoginEmployer not implemented")
}
func (UnimplementedEmployerServer) mustEmbedUnimplementedEmployerServer() {}
func (UnimplementedEmployerServer) testEmbeddedByValue()                  {}

// UnsafeEmployerServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to EmployerServer will
// result in compilation errors.
type UnsafeEmployerServer interface {
	mustEmbedUnimplementedEmployerServer()
}

func RegisterEmployerServer(s grpc.ServiceRegistrar, srv EmployerServer) {
	// If the following call pancis, it indicates UnimplementedEmployerServer was
	// embedded by pointer and is nil.  This will cause panics if an
	// unimplemented method is ever invoked, so we test this at initialization
	// time to prevent it from happening at runtime later due to I/O.
	if t, ok := srv.(interface{ testEmbeddedByValue() }); ok {
		t.testEmbeddedByValue()
	}
	s.RegisterService(&Employer_ServiceDesc, srv)
}

func _Employer_CreateEmployer_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateEmployerReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(EmployerServer).CreateEmployer(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Employer_CreateEmployer_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(EmployerServer).CreateEmployer(ctx, req.(*CreateEmployerReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _Employer_LoginEmployer_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(EmpLoginReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(EmployerServer).LoginEmployer(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Employer_LoginEmployer_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(EmployerServer).LoginEmployer(ctx, req.(*EmpLoginReq))
	}
	return interceptor(ctx, in, info, handler)
}

// Employer_ServiceDesc is the grpc.ServiceDesc for Employer service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Employer_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "employer.Employer",
	HandlerType: (*EmployerServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "CreateEmployer",
			Handler:    _Employer_CreateEmployer_Handler,
		},
		{
			MethodName: "LoginEmployer",
			Handler:    _Employer_LoginEmployer_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "utils/pb/employer/employer.proto",
}
