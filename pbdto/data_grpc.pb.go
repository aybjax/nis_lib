// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v3.6.1
// source: data.proto

package pbdto

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.32.0 or later.
const _ = grpc.SupportPackageIsVersion7

// GetStudentCoursesClient is the client API for GetStudentCourses service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type GetStudentCoursesClient interface {
	GetCourses(ctx context.Context, in *Request, opts ...grpc.CallOption) (*CoursesResponse, error)
}

type getStudentCoursesClient struct {
	cc grpc.ClientConnInterface
}

func NewGetStudentCoursesClient(cc grpc.ClientConnInterface) GetStudentCoursesClient {
	return &getStudentCoursesClient{cc}
}

func (c *getStudentCoursesClient) GetCourses(ctx context.Context, in *Request, opts ...grpc.CallOption) (*CoursesResponse, error) {
	out := new(CoursesResponse)
	err := c.cc.Invoke(ctx, "/nis_test.GetStudentCourses/GetCourses", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// GetStudentCoursesServer is the server API for GetStudentCourses service.
// All implementations should embed UnimplementedGetStudentCoursesServer
// for forward compatibility
type GetStudentCoursesServer interface {
	GetCourses(context.Context, *Request) (*CoursesResponse, error)
}

// UnimplementedGetStudentCoursesServer should be embedded to have forward compatible implementations.
type UnimplementedGetStudentCoursesServer struct {
}

func (UnimplementedGetStudentCoursesServer) GetCourses(context.Context, *Request) (*CoursesResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetCourses not implemented")
}

// UnsafeGetStudentCoursesServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to GetStudentCoursesServer will
// result in compilation errors.
type UnsafeGetStudentCoursesServer interface {
	mustEmbedUnimplementedGetStudentCoursesServer()
}

func RegisterGetStudentCoursesServer(s grpc.ServiceRegistrar, srv GetStudentCoursesServer) {
	s.RegisterService(&GetStudentCourses_ServiceDesc, srv)
}

func _GetStudentCourses_GetCourses_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Request)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GetStudentCoursesServer).GetCourses(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/nis_test.GetStudentCourses/GetCourses",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GetStudentCoursesServer).GetCourses(ctx, req.(*Request))
	}
	return interceptor(ctx, in, info, handler)
}

// GetStudentCourses_ServiceDesc is the grpc.ServiceDesc for GetStudentCourses service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var GetStudentCourses_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "nis_test.GetStudentCourses",
	HandlerType: (*GetStudentCoursesServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetCourses",
			Handler:    _GetStudentCourses_GetCourses_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "data.proto",
}

// GetCourseStudentsClient is the client API for GetCourseStudents service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type GetCourseStudentsClient interface {
	GetStudents(ctx context.Context, in *Request, opts ...grpc.CallOption) (*StudentsResponse, error)
}

type getCourseStudentsClient struct {
	cc grpc.ClientConnInterface
}

func NewGetCourseStudentsClient(cc grpc.ClientConnInterface) GetCourseStudentsClient {
	return &getCourseStudentsClient{cc}
}

func (c *getCourseStudentsClient) GetStudents(ctx context.Context, in *Request, opts ...grpc.CallOption) (*StudentsResponse, error) {
	out := new(StudentsResponse)
	err := c.cc.Invoke(ctx, "/nis_test.GetCourseStudents/GetStudents", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// GetCourseStudentsServer is the server API for GetCourseStudents service.
// All implementations should embed UnimplementedGetCourseStudentsServer
// for forward compatibility
type GetCourseStudentsServer interface {
	GetStudents(context.Context, *Request) (*StudentsResponse, error)
}

// UnimplementedGetCourseStudentsServer should be embedded to have forward compatible implementations.
type UnimplementedGetCourseStudentsServer struct {
}

func (UnimplementedGetCourseStudentsServer) GetStudents(context.Context, *Request) (*StudentsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetStudents not implemented")
}

// UnsafeGetCourseStudentsServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to GetCourseStudentsServer will
// result in compilation errors.
type UnsafeGetCourseStudentsServer interface {
	mustEmbedUnimplementedGetCourseStudentsServer()
}

func RegisterGetCourseStudentsServer(s grpc.ServiceRegistrar, srv GetCourseStudentsServer) {
	s.RegisterService(&GetCourseStudents_ServiceDesc, srv)
}

func _GetCourseStudents_GetStudents_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Request)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GetCourseStudentsServer).GetStudents(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/nis_test.GetCourseStudents/GetStudents",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GetCourseStudentsServer).GetStudents(ctx, req.(*Request))
	}
	return interceptor(ctx, in, info, handler)
}

// GetCourseStudents_ServiceDesc is the grpc.ServiceDesc for GetCourseStudents service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var GetCourseStudents_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "nis_test.GetCourseStudents",
	HandlerType: (*GetCourseStudentsServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetStudents",
			Handler:    _GetCourseStudents_GetStudents_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "data.proto",
}
