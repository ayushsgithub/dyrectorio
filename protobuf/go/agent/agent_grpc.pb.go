// Code generated by protoc-gen-go-grpc. DO NOT EDIT.

package agent

import (
	context "context"
	crux "gitlab.com/dyrector_io/dyrector.io/protobuf/go/crux"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.32.0 or later.
const _ = grpc.SupportPackageIsVersion7

// AgentClient is the client API for Agent service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type AgentClient interface {
	//*
	// Subscribe with pre-assigned AgentID, waiting for incoming
	// deploy requests and prefix status requests.
	// In both cases, separate, shorter-living channels are opened.
	// For deployment status reports, closed when ended.
	// For prefix status reports, should be closed by the server.
	Connect(ctx context.Context, in *AgentInfo, opts ...grpc.CallOption) (Agent_ConnectClient, error)
	DeploymentStatus(ctx context.Context, opts ...grpc.CallOption) (Agent_DeploymentStatusClient, error)
	ContainerStatus(ctx context.Context, opts ...grpc.CallOption) (Agent_ContainerStatusClient, error)
}

type agentClient struct {
	cc grpc.ClientConnInterface
}

func NewAgentClient(cc grpc.ClientConnInterface) AgentClient {
	return &agentClient{cc}
}

func (c *agentClient) Connect(ctx context.Context, in *AgentInfo, opts ...grpc.CallOption) (Agent_ConnectClient, error) {
	stream, err := c.cc.NewStream(ctx, &Agent_ServiceDesc.Streams[0], "/agent.Agent/Connect", opts...)
	if err != nil {
		return nil, err
	}
	x := &agentConnectClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type Agent_ConnectClient interface {
	Recv() (*AgentCommand, error)
	grpc.ClientStream
}

type agentConnectClient struct {
	grpc.ClientStream
}

func (x *agentConnectClient) Recv() (*AgentCommand, error) {
	m := new(AgentCommand)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *agentClient) DeploymentStatus(ctx context.Context, opts ...grpc.CallOption) (Agent_DeploymentStatusClient, error) {
	stream, err := c.cc.NewStream(ctx, &Agent_ServiceDesc.Streams[1], "/agent.Agent/DeploymentStatus", opts...)
	if err != nil {
		return nil, err
	}
	x := &agentDeploymentStatusClient{stream}
	return x, nil
}

type Agent_DeploymentStatusClient interface {
	Send(*crux.DeploymentStatusMessage) error
	CloseAndRecv() (*Empty, error)
	grpc.ClientStream
}

type agentDeploymentStatusClient struct {
	grpc.ClientStream
}

func (x *agentDeploymentStatusClient) Send(m *crux.DeploymentStatusMessage) error {
	return x.ClientStream.SendMsg(m)
}

func (x *agentDeploymentStatusClient) CloseAndRecv() (*Empty, error) {
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	m := new(Empty)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *agentClient) ContainerStatus(ctx context.Context, opts ...grpc.CallOption) (Agent_ContainerStatusClient, error) {
	stream, err := c.cc.NewStream(ctx, &Agent_ServiceDesc.Streams[2], "/agent.Agent/ContainerStatus", opts...)
	if err != nil {
		return nil, err
	}
	x := &agentContainerStatusClient{stream}
	return x, nil
}

type Agent_ContainerStatusClient interface {
	Send(*crux.ContainerStatusListMessage) error
	CloseAndRecv() (*Empty, error)
	grpc.ClientStream
}

type agentContainerStatusClient struct {
	grpc.ClientStream
}

func (x *agentContainerStatusClient) Send(m *crux.ContainerStatusListMessage) error {
	return x.ClientStream.SendMsg(m)
}

func (x *agentContainerStatusClient) CloseAndRecv() (*Empty, error) {
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	m := new(Empty)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// AgentServer is the server API for Agent service.
// All implementations must embed UnimplementedAgentServer
// for forward compatibility
type AgentServer interface {
	//*
	// Subscribe with pre-assigned AgentID, waiting for incoming
	// deploy requests and prefix status requests.
	// In both cases, separate, shorter-living channels are opened.
	// For deployment status reports, closed when ended.
	// For prefix status reports, should be closed by the server.
	Connect(*AgentInfo, Agent_ConnectServer) error
	DeploymentStatus(Agent_DeploymentStatusServer) error
	ContainerStatus(Agent_ContainerStatusServer) error
	mustEmbedUnimplementedAgentServer()
}

// UnimplementedAgentServer must be embedded to have forward compatible implementations.
type UnimplementedAgentServer struct {
}

func (UnimplementedAgentServer) Connect(*AgentInfo, Agent_ConnectServer) error {
	return status.Errorf(codes.Unimplemented, "method Connect not implemented")
}
func (UnimplementedAgentServer) DeploymentStatus(Agent_DeploymentStatusServer) error {
	return status.Errorf(codes.Unimplemented, "method DeploymentStatus not implemented")
}
func (UnimplementedAgentServer) ContainerStatus(Agent_ContainerStatusServer) error {
	return status.Errorf(codes.Unimplemented, "method ContainerStatus not implemented")
}
func (UnimplementedAgentServer) mustEmbedUnimplementedAgentServer() {}

// UnsafeAgentServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to AgentServer will
// result in compilation errors.
type UnsafeAgentServer interface {
	mustEmbedUnimplementedAgentServer()
}

func RegisterAgentServer(s grpc.ServiceRegistrar, srv AgentServer) {
	s.RegisterService(&Agent_ServiceDesc, srv)
}

func _Agent_Connect_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(AgentInfo)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(AgentServer).Connect(m, &agentConnectServer{stream})
}

type Agent_ConnectServer interface {
	Send(*AgentCommand) error
	grpc.ServerStream
}

type agentConnectServer struct {
	grpc.ServerStream
}

func (x *agentConnectServer) Send(m *AgentCommand) error {
	return x.ServerStream.SendMsg(m)
}

func _Agent_DeploymentStatus_Handler(srv interface{}, stream grpc.ServerStream) error {
	return srv.(AgentServer).DeploymentStatus(&agentDeploymentStatusServer{stream})
}

type Agent_DeploymentStatusServer interface {
	SendAndClose(*Empty) error
	Recv() (*crux.DeploymentStatusMessage, error)
	grpc.ServerStream
}

type agentDeploymentStatusServer struct {
	grpc.ServerStream
}

func (x *agentDeploymentStatusServer) SendAndClose(m *Empty) error {
	return x.ServerStream.SendMsg(m)
}

func (x *agentDeploymentStatusServer) Recv() (*crux.DeploymentStatusMessage, error) {
	m := new(crux.DeploymentStatusMessage)
	if err := x.ServerStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func _Agent_ContainerStatus_Handler(srv interface{}, stream grpc.ServerStream) error {
	return srv.(AgentServer).ContainerStatus(&agentContainerStatusServer{stream})
}

type Agent_ContainerStatusServer interface {
	SendAndClose(*Empty) error
	Recv() (*crux.ContainerStatusListMessage, error)
	grpc.ServerStream
}

type agentContainerStatusServer struct {
	grpc.ServerStream
}

func (x *agentContainerStatusServer) SendAndClose(m *Empty) error {
	return x.ServerStream.SendMsg(m)
}

func (x *agentContainerStatusServer) Recv() (*crux.ContainerStatusListMessage, error) {
	m := new(crux.ContainerStatusListMessage)
	if err := x.ServerStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// Agent_ServiceDesc is the grpc.ServiceDesc for Agent service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Agent_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "agent.Agent",
	HandlerType: (*AgentServer)(nil),
	Methods:     []grpc.MethodDesc{},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "Connect",
			Handler:       _Agent_Connect_Handler,
			ServerStreams: true,
		},
		{
			StreamName:    "DeploymentStatus",
			Handler:       _Agent_DeploymentStatus_Handler,
			ClientStreams: true,
		},
		{
			StreamName:    "ContainerStatus",
			Handler:       _Agent_ContainerStatus_Handler,
			ClientStreams: true,
		},
	},
	Metadata: "proto/agent.proto",
}
