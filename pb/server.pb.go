// Code generated by protoc-gen-go. DO NOT EDIT.
// source: pb/server.proto

package pb

import (
	context "context"
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	empty "github.com/golang/protobuf/ptypes/empty"
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

func init() {
	proto.RegisterFile("pb/server.proto", fileDescriptor_98d161fbd54d4312)
}

var fileDescriptor_98d161fbd54d4312 = []byte{
	// 447 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x7c, 0x93, 0xcd, 0x72, 0xd3, 0x30,
	0x14, 0x85, 0x5b, 0x16, 0x4c, 0xab, 0x10, 0x13, 0x5c, 0xda, 0x85, 0xd9, 0x75, 0xc5, 0xca, 0x2e,
	0x81, 0x05, 0x0b, 0xa6, 0x33, 0x9d, 0x92, 0x85, 0xf9, 0x9d, 0x71, 0xba, 0x61, 0x29, 0xbb, 0x17,
	0xc7, 0x83, 0x65, 0x09, 0x59, 0xe6, 0x27, 0x4f, 0xc2, 0xe3, 0x32, 0xf2, 0x95, 0x1c, 0x29, 0x76,
	0xb3, 0xd4, 0xa7, 0x73, 0xee, 0xf1, 0x3d, 0x4a, 0xc8, 0x53, 0x91, 0x27, 0x2d, 0xc8, 0x5f, 0x20,
	0x63, 0x21, 0xb9, 0xe2, 0xe1, 0x23, 0x91, 0x47, 0x0b, 0x91, 0x27, 0x0d, 0xa8, 0xdf, 0x5c, 0xfe,
	0x40, 0x1a, 0x05, 0x22, 0x4f, 0x2a, 0x46, 0x4b, 0x30, 0x67, 0xad, 0x60, 0xb4, 0xd8, 0x54, 0x8d,
	0x25, 0x7a, 0x50, 0xc1, 0x19, 0xe3, 0x8d, 0x01, 0x2f, 0x4a, 0xce, 0xcb, 0x1a, 0x92, 0xfe, 0x94,
	0x77, 0xdf, 0x13, 0x60, 0x42, 0xfd, 0xc5, 0xcb, 0xe5, 0xbf, 0x13, 0x72, 0xb2, 0xfa, 0x03, 0x45,
	0xa7, 0xb8, 0x0c, 0xaf, 0xc9, 0xfc, 0x56, 0x02, 0x55, 0xf0, 0x05, 0x33, 0xc3, 0xe7, 0xb1, 0xc8,
	0x63, 0x73, 0xc0, 0x9b, 0x0c, 0x7e, 0x46, 0xe7, 0x13, 0xb4, 0x15, 0x97, 0x47, 0xda, 0xff, 0x1e,
	0x6a, 0x98, 0xf6, 0xe3, 0xcd, 0xbe, 0xdf, 0xd2, 0xde, 0x7f, 0x43, 0x82, 0xb4, 0x69, 0x05, 0x14,
	0xca, 0x0e, 0x70, 0xa5, 0xe6, 0x4a, 0x4f, 0xb8, 0x98, 0xc2, 0xfd, 0x88, 0x77, 0x64, 0xf6, 0xa9,
	0x6a, 0x07, 0xff, 0x45, 0x8c, 0xcb, 0xc7, 0x76, 0xf9, 0x78, 0xa5, 0x97, 0x8f, 0xce, 0x9c, 0x01,
	0x5a, 0x6f, 0xdc, 0x6f, 0xc9, 0x2c, 0x65, 0x82, 0x4b, 0x95, 0xea, 0x8a, 0xc3, 0x50, 0xab, 0x1c,
	0xa0, 0xa3, 0xcf, 0x46, 0xcc, 0x38, 0x4f, 0xf5, 0x1c, 0xf4, 0x3d, 0x94, 0xfa, 0x4c, 0x7b, 0x07,
	0xd9, 0x2e, 0x13, 0x4b, 0x70, 0x32, 0x1d, 0x30, 0x64, 0x7a, 0xcc, 0xd6, 0x8d, 0xf5, 0x7f, 0xc6,
	0x1f, 0x00, 0xd6, 0xed, 0xa1, 0xa1, 0xee, 0x3d, 0xda, 0xfb, 0x97, 0xf6, 0xb9, 0x3c, 0xbf, 0x87,
	0xb4, 0xff, 0x54, 0xd3, 0x95, 0x94, 0x5c, 0xee, 0xfa, 0xb5, 0x8e, 0x83, 0xfd, 0x3a, 0x42, 0x93,
	0x78, 0x45, 0x9e, 0xac, 0x15, 0x95, 0x83, 0xbd, 0x97, 0xb9, 0x64, 0x94, 0x17, 0x93, 0xd9, 0xc7,
	0xaa, 0xae, 0xad, 0xa1, 0x6f, 0xc7, 0x01, 0x53, 0xfa, 0xb5, 0xe2, 0xc2, 0xd3, 0x3b, 0x60, 0xa4,
	0x5f, 0x92, 0x79, 0x06, 0x0d, 0x65, 0x7e, 0x07, 0x1e, 0x1a, 0x79, 0xde, 0x90, 0x20, 0x83, 0xd6,
	0xdd, 0xe3, 0x1c, 0x4d, 0xed, 0xa1, 0x4d, 0xae, 0xc9, 0xfc, 0x46, 0x29, 0x5a, 0x6c, 0xbc, 0x6f,
	0x43, 0xb4, 0x56, 0x12, 0x28, 0x4b, 0x1b, 0xec, 0xcd, 0x65, 0x5f, 0x3b, 0x75, 0x79, 0xf4, 0xf2,
	0xf8, 0xea, 0x38, 0x7c, 0x45, 0x16, 0x19, 0xb4, 0xd5, 0xd6, 0x7e, 0xd6, 0xdd, 0xdd, 0xb7, 0x70,
	0x61, 0x72, 0xab, 0xad, 0x3e, 0x4e, 0x3c, 0x56, 0x70, 0x4b, 0x1b, 0x1c, 0xf7, 0xa1, 0xbb, 0x2f,
	0x01, 0x0d, 0x66, 0x40, 0x7a, 0x3f, 0xfc, 0x95, 0x7c, 0x15, 0x3e, 0x56, 0xfe, 0xb8, 0x7f, 0xd3,
	0xd7, 0xff, 0x03, 0x00, 0x00, 0xff, 0xff, 0x70, 0x6f, 0x30, 0xc2, 0x9a, 0x04, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConnInterface

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion6

// ExecutorClient is the client API for Executor service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type ExecutorClient interface {
	// notwork
	CreateNetwork(ctx context.Context, in *NetworkCreateReq, opts ...grpc.CallOption) (*NetworkCreateResp, error)
	DeleteNetwork(ctx context.Context, in *NetworkDeleteReq, opts ...grpc.CallOption) (*NetworkDeleteResp, error)
	InspectNetwork(ctx context.Context, in *NetworkInspectReq, opts ...grpc.CallOption) (*NetworkInspectResp, error)
	ListNetwork(ctx context.Context, in *empty.Empty, opts ...grpc.CallOption) (*NetworkListResp, error)
	// image
	ImportImage(ctx context.Context, in *ImportImageReq, opts ...grpc.CallOption) (*ImportImageResp, error)
	ListImage(ctx context.Context, in *empty.Empty, opts ...grpc.CallOption) (*ListImageResp, error)
	DeleteImage(ctx context.Context, in *DeleteImageReq, opts ...grpc.CallOption) (*DeleteImageResp, error)
	// machine
	CreateMachine(ctx context.Context, in *CreateMachineReq, opts ...grpc.CallOption) (*CreateMachineResp, error)
	DeleteMachine(ctx context.Context, in *DeleteMachineReq, opts ...grpc.CallOption) (*Error, error)
	ListMachine(ctx context.Context, in *empty.Empty, opts ...grpc.CallOption) (*ListMachineResp, error)
	StartMachine(ctx context.Context, in *StartMachineReq, opts ...grpc.CallOption) (*Error, error)
	KillMachine(ctx context.Context, in *KillMachineReq, opts ...grpc.CallOption) (*Error, error)
	StopMachine(ctx context.Context, in *StopMachineReq, opts ...grpc.CallOption) (*Error, error)
	RenameMachine(ctx context.Context, in *RenameMachineReq, opts ...grpc.CallOption) (*Error, error)
	RestartMachine(ctx context.Context, in *RestartMachineReq, opts ...grpc.CallOption) (*Error, error)
	AttachMachine(ctx context.Context, opts ...grpc.CallOption) (Executor_AttachMachineClient, error)
	ResizeMachineTTY(ctx context.Context, in *ResizeTTYReq, opts ...grpc.CallOption) (*Error, error)
	CanAttachJudge(ctx context.Context, in *MachineIdReq, opts ...grpc.CallOption) (*CanAttachJudgeResp, error)
}

type executorClient struct {
	cc grpc.ClientConnInterface
}

func NewExecutorClient(cc grpc.ClientConnInterface) ExecutorClient {
	return &executorClient{cc}
}

func (c *executorClient) CreateNetwork(ctx context.Context, in *NetworkCreateReq, opts ...grpc.CallOption) (*NetworkCreateResp, error) {
	out := new(NetworkCreateResp)
	err := c.cc.Invoke(ctx, "/pb.Executor/CreateNetwork", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *executorClient) DeleteNetwork(ctx context.Context, in *NetworkDeleteReq, opts ...grpc.CallOption) (*NetworkDeleteResp, error) {
	out := new(NetworkDeleteResp)
	err := c.cc.Invoke(ctx, "/pb.Executor/DeleteNetwork", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *executorClient) InspectNetwork(ctx context.Context, in *NetworkInspectReq, opts ...grpc.CallOption) (*NetworkInspectResp, error) {
	out := new(NetworkInspectResp)
	err := c.cc.Invoke(ctx, "/pb.Executor/InspectNetwork", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *executorClient) ListNetwork(ctx context.Context, in *empty.Empty, opts ...grpc.CallOption) (*NetworkListResp, error) {
	out := new(NetworkListResp)
	err := c.cc.Invoke(ctx, "/pb.Executor/ListNetwork", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *executorClient) ImportImage(ctx context.Context, in *ImportImageReq, opts ...grpc.CallOption) (*ImportImageResp, error) {
	out := new(ImportImageResp)
	err := c.cc.Invoke(ctx, "/pb.Executor/ImportImage", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *executorClient) ListImage(ctx context.Context, in *empty.Empty, opts ...grpc.CallOption) (*ListImageResp, error) {
	out := new(ListImageResp)
	err := c.cc.Invoke(ctx, "/pb.Executor/ListImage", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *executorClient) DeleteImage(ctx context.Context, in *DeleteImageReq, opts ...grpc.CallOption) (*DeleteImageResp, error) {
	out := new(DeleteImageResp)
	err := c.cc.Invoke(ctx, "/pb.Executor/DeleteImage", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *executorClient) CreateMachine(ctx context.Context, in *CreateMachineReq, opts ...grpc.CallOption) (*CreateMachineResp, error) {
	out := new(CreateMachineResp)
	err := c.cc.Invoke(ctx, "/pb.Executor/CreateMachine", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *executorClient) DeleteMachine(ctx context.Context, in *DeleteMachineReq, opts ...grpc.CallOption) (*Error, error) {
	out := new(Error)
	err := c.cc.Invoke(ctx, "/pb.Executor/DeleteMachine", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *executorClient) ListMachine(ctx context.Context, in *empty.Empty, opts ...grpc.CallOption) (*ListMachineResp, error) {
	out := new(ListMachineResp)
	err := c.cc.Invoke(ctx, "/pb.Executor/ListMachine", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *executorClient) StartMachine(ctx context.Context, in *StartMachineReq, opts ...grpc.CallOption) (*Error, error) {
	out := new(Error)
	err := c.cc.Invoke(ctx, "/pb.Executor/StartMachine", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *executorClient) KillMachine(ctx context.Context, in *KillMachineReq, opts ...grpc.CallOption) (*Error, error) {
	out := new(Error)
	err := c.cc.Invoke(ctx, "/pb.Executor/KillMachine", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *executorClient) StopMachine(ctx context.Context, in *StopMachineReq, opts ...grpc.CallOption) (*Error, error) {
	out := new(Error)
	err := c.cc.Invoke(ctx, "/pb.Executor/StopMachine", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *executorClient) RenameMachine(ctx context.Context, in *RenameMachineReq, opts ...grpc.CallOption) (*Error, error) {
	out := new(Error)
	err := c.cc.Invoke(ctx, "/pb.Executor/RenameMachine", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *executorClient) RestartMachine(ctx context.Context, in *RestartMachineReq, opts ...grpc.CallOption) (*Error, error) {
	out := new(Error)
	err := c.cc.Invoke(ctx, "/pb.Executor/RestartMachine", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *executorClient) AttachMachine(ctx context.Context, opts ...grpc.CallOption) (Executor_AttachMachineClient, error) {
	stream, err := c.cc.NewStream(ctx, &_Executor_serviceDesc.Streams[0], "/pb.Executor/AttachMachine", opts...)
	if err != nil {
		return nil, err
	}
	x := &executorAttachMachineClient{stream}
	return x, nil
}

type Executor_AttachMachineClient interface {
	Send(*AttachStreamIn) error
	Recv() (*AttachStreamOut, error)
	grpc.ClientStream
}

type executorAttachMachineClient struct {
	grpc.ClientStream
}

func (x *executorAttachMachineClient) Send(m *AttachStreamIn) error {
	return x.ClientStream.SendMsg(m)
}

func (x *executorAttachMachineClient) Recv() (*AttachStreamOut, error) {
	m := new(AttachStreamOut)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *executorClient) ResizeMachineTTY(ctx context.Context, in *ResizeTTYReq, opts ...grpc.CallOption) (*Error, error) {
	out := new(Error)
	err := c.cc.Invoke(ctx, "/pb.Executor/ResizeMachineTTY", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *executorClient) CanAttachJudge(ctx context.Context, in *MachineIdReq, opts ...grpc.CallOption) (*CanAttachJudgeResp, error) {
	out := new(CanAttachJudgeResp)
	err := c.cc.Invoke(ctx, "/pb.Executor/CanAttachJudge", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// ExecutorServer is the server API for Executor service.
type ExecutorServer interface {
	// notwork
	CreateNetwork(context.Context, *NetworkCreateReq) (*NetworkCreateResp, error)
	DeleteNetwork(context.Context, *NetworkDeleteReq) (*NetworkDeleteResp, error)
	InspectNetwork(context.Context, *NetworkInspectReq) (*NetworkInspectResp, error)
	ListNetwork(context.Context, *empty.Empty) (*NetworkListResp, error)
	// image
	ImportImage(context.Context, *ImportImageReq) (*ImportImageResp, error)
	ListImage(context.Context, *empty.Empty) (*ListImageResp, error)
	DeleteImage(context.Context, *DeleteImageReq) (*DeleteImageResp, error)
	// machine
	CreateMachine(context.Context, *CreateMachineReq) (*CreateMachineResp, error)
	DeleteMachine(context.Context, *DeleteMachineReq) (*Error, error)
	ListMachine(context.Context, *empty.Empty) (*ListMachineResp, error)
	StartMachine(context.Context, *StartMachineReq) (*Error, error)
	KillMachine(context.Context, *KillMachineReq) (*Error, error)
	StopMachine(context.Context, *StopMachineReq) (*Error, error)
	RenameMachine(context.Context, *RenameMachineReq) (*Error, error)
	RestartMachine(context.Context, *RestartMachineReq) (*Error, error)
	AttachMachine(Executor_AttachMachineServer) error
	ResizeMachineTTY(context.Context, *ResizeTTYReq) (*Error, error)
	CanAttachJudge(context.Context, *MachineIdReq) (*CanAttachJudgeResp, error)
}

// UnimplementedExecutorServer can be embedded to have forward compatible implementations.
type UnimplementedExecutorServer struct {
}

func (*UnimplementedExecutorServer) CreateNetwork(ctx context.Context, req *NetworkCreateReq) (*NetworkCreateResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateNetwork not implemented")
}
func (*UnimplementedExecutorServer) DeleteNetwork(ctx context.Context, req *NetworkDeleteReq) (*NetworkDeleteResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteNetwork not implemented")
}
func (*UnimplementedExecutorServer) InspectNetwork(ctx context.Context, req *NetworkInspectReq) (*NetworkInspectResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method InspectNetwork not implemented")
}
func (*UnimplementedExecutorServer) ListNetwork(ctx context.Context, req *empty.Empty) (*NetworkListResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ListNetwork not implemented")
}
func (*UnimplementedExecutorServer) ImportImage(ctx context.Context, req *ImportImageReq) (*ImportImageResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ImportImage not implemented")
}
func (*UnimplementedExecutorServer) ListImage(ctx context.Context, req *empty.Empty) (*ListImageResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ListImage not implemented")
}
func (*UnimplementedExecutorServer) DeleteImage(ctx context.Context, req *DeleteImageReq) (*DeleteImageResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteImage not implemented")
}
func (*UnimplementedExecutorServer) CreateMachine(ctx context.Context, req *CreateMachineReq) (*CreateMachineResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateMachine not implemented")
}
func (*UnimplementedExecutorServer) DeleteMachine(ctx context.Context, req *DeleteMachineReq) (*Error, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteMachine not implemented")
}
func (*UnimplementedExecutorServer) ListMachine(ctx context.Context, req *empty.Empty) (*ListMachineResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ListMachine not implemented")
}
func (*UnimplementedExecutorServer) StartMachine(ctx context.Context, req *StartMachineReq) (*Error, error) {
	return nil, status.Errorf(codes.Unimplemented, "method StartMachine not implemented")
}
func (*UnimplementedExecutorServer) KillMachine(ctx context.Context, req *KillMachineReq) (*Error, error) {
	return nil, status.Errorf(codes.Unimplemented, "method KillMachine not implemented")
}
func (*UnimplementedExecutorServer) StopMachine(ctx context.Context, req *StopMachineReq) (*Error, error) {
	return nil, status.Errorf(codes.Unimplemented, "method StopMachine not implemented")
}
func (*UnimplementedExecutorServer) RenameMachine(ctx context.Context, req *RenameMachineReq) (*Error, error) {
	return nil, status.Errorf(codes.Unimplemented, "method RenameMachine not implemented")
}
func (*UnimplementedExecutorServer) RestartMachine(ctx context.Context, req *RestartMachineReq) (*Error, error) {
	return nil, status.Errorf(codes.Unimplemented, "method RestartMachine not implemented")
}
func (*UnimplementedExecutorServer) AttachMachine(srv Executor_AttachMachineServer) error {
	return status.Errorf(codes.Unimplemented, "method AttachMachine not implemented")
}
func (*UnimplementedExecutorServer) ResizeMachineTTY(ctx context.Context, req *ResizeTTYReq) (*Error, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ResizeMachineTTY not implemented")
}
func (*UnimplementedExecutorServer) CanAttachJudge(ctx context.Context, req *MachineIdReq) (*CanAttachJudgeResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CanAttachJudge not implemented")
}

func RegisterExecutorServer(s *grpc.Server, srv ExecutorServer) {
	s.RegisterService(&_Executor_serviceDesc, srv)
}

func _Executor_CreateNetwork_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(NetworkCreateReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ExecutorServer).CreateNetwork(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pb.Executor/CreateNetwork",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ExecutorServer).CreateNetwork(ctx, req.(*NetworkCreateReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _Executor_DeleteNetwork_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(NetworkDeleteReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ExecutorServer).DeleteNetwork(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pb.Executor/DeleteNetwork",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ExecutorServer).DeleteNetwork(ctx, req.(*NetworkDeleteReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _Executor_InspectNetwork_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(NetworkInspectReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ExecutorServer).InspectNetwork(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pb.Executor/InspectNetwork",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ExecutorServer).InspectNetwork(ctx, req.(*NetworkInspectReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _Executor_ListNetwork_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(empty.Empty)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ExecutorServer).ListNetwork(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pb.Executor/ListNetwork",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ExecutorServer).ListNetwork(ctx, req.(*empty.Empty))
	}
	return interceptor(ctx, in, info, handler)
}

func _Executor_ImportImage_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ImportImageReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ExecutorServer).ImportImage(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pb.Executor/ImportImage",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ExecutorServer).ImportImage(ctx, req.(*ImportImageReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _Executor_ListImage_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(empty.Empty)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ExecutorServer).ListImage(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pb.Executor/ListImage",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ExecutorServer).ListImage(ctx, req.(*empty.Empty))
	}
	return interceptor(ctx, in, info, handler)
}

func _Executor_DeleteImage_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeleteImageReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ExecutorServer).DeleteImage(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pb.Executor/DeleteImage",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ExecutorServer).DeleteImage(ctx, req.(*DeleteImageReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _Executor_CreateMachine_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateMachineReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ExecutorServer).CreateMachine(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pb.Executor/CreateMachine",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ExecutorServer).CreateMachine(ctx, req.(*CreateMachineReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _Executor_DeleteMachine_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeleteMachineReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ExecutorServer).DeleteMachine(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pb.Executor/DeleteMachine",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ExecutorServer).DeleteMachine(ctx, req.(*DeleteMachineReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _Executor_ListMachine_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(empty.Empty)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ExecutorServer).ListMachine(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pb.Executor/ListMachine",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ExecutorServer).ListMachine(ctx, req.(*empty.Empty))
	}
	return interceptor(ctx, in, info, handler)
}

func _Executor_StartMachine_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(StartMachineReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ExecutorServer).StartMachine(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pb.Executor/StartMachine",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ExecutorServer).StartMachine(ctx, req.(*StartMachineReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _Executor_KillMachine_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(KillMachineReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ExecutorServer).KillMachine(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pb.Executor/KillMachine",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ExecutorServer).KillMachine(ctx, req.(*KillMachineReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _Executor_StopMachine_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(StopMachineReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ExecutorServer).StopMachine(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pb.Executor/StopMachine",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ExecutorServer).StopMachine(ctx, req.(*StopMachineReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _Executor_RenameMachine_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(RenameMachineReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ExecutorServer).RenameMachine(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pb.Executor/RenameMachine",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ExecutorServer).RenameMachine(ctx, req.(*RenameMachineReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _Executor_RestartMachine_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(RestartMachineReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ExecutorServer).RestartMachine(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pb.Executor/RestartMachine",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ExecutorServer).RestartMachine(ctx, req.(*RestartMachineReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _Executor_AttachMachine_Handler(srv interface{}, stream grpc.ServerStream) error {
	return srv.(ExecutorServer).AttachMachine(&executorAttachMachineServer{stream})
}

type Executor_AttachMachineServer interface {
	Send(*AttachStreamOut) error
	Recv() (*AttachStreamIn, error)
	grpc.ServerStream
}

type executorAttachMachineServer struct {
	grpc.ServerStream
}

func (x *executorAttachMachineServer) Send(m *AttachStreamOut) error {
	return x.ServerStream.SendMsg(m)
}

func (x *executorAttachMachineServer) Recv() (*AttachStreamIn, error) {
	m := new(AttachStreamIn)
	if err := x.ServerStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func _Executor_ResizeMachineTTY_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ResizeTTYReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ExecutorServer).ResizeMachineTTY(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pb.Executor/ResizeMachineTTY",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ExecutorServer).ResizeMachineTTY(ctx, req.(*ResizeTTYReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _Executor_CanAttachJudge_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(MachineIdReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ExecutorServer).CanAttachJudge(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pb.Executor/CanAttachJudge",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ExecutorServer).CanAttachJudge(ctx, req.(*MachineIdReq))
	}
	return interceptor(ctx, in, info, handler)
}

var _Executor_serviceDesc = grpc.ServiceDesc{
	ServiceName: "pb.Executor",
	HandlerType: (*ExecutorServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "CreateNetwork",
			Handler:    _Executor_CreateNetwork_Handler,
		},
		{
			MethodName: "DeleteNetwork",
			Handler:    _Executor_DeleteNetwork_Handler,
		},
		{
			MethodName: "InspectNetwork",
			Handler:    _Executor_InspectNetwork_Handler,
		},
		{
			MethodName: "ListNetwork",
			Handler:    _Executor_ListNetwork_Handler,
		},
		{
			MethodName: "ImportImage",
			Handler:    _Executor_ImportImage_Handler,
		},
		{
			MethodName: "ListImage",
			Handler:    _Executor_ListImage_Handler,
		},
		{
			MethodName: "DeleteImage",
			Handler:    _Executor_DeleteImage_Handler,
		},
		{
			MethodName: "CreateMachine",
			Handler:    _Executor_CreateMachine_Handler,
		},
		{
			MethodName: "DeleteMachine",
			Handler:    _Executor_DeleteMachine_Handler,
		},
		{
			MethodName: "ListMachine",
			Handler:    _Executor_ListMachine_Handler,
		},
		{
			MethodName: "StartMachine",
			Handler:    _Executor_StartMachine_Handler,
		},
		{
			MethodName: "KillMachine",
			Handler:    _Executor_KillMachine_Handler,
		},
		{
			MethodName: "StopMachine",
			Handler:    _Executor_StopMachine_Handler,
		},
		{
			MethodName: "RenameMachine",
			Handler:    _Executor_RenameMachine_Handler,
		},
		{
			MethodName: "RestartMachine",
			Handler:    _Executor_RestartMachine_Handler,
		},
		{
			MethodName: "ResizeMachineTTY",
			Handler:    _Executor_ResizeMachineTTY_Handler,
		},
		{
			MethodName: "CanAttachJudge",
			Handler:    _Executor_CanAttachJudge_Handler,
		},
	},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "AttachMachine",
			Handler:       _Executor_AttachMachine_Handler,
			ServerStreams: true,
			ClientStreams: true,
		},
	},
	Metadata: "pb/server.proto",
}
