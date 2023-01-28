package grpc

import (
	"context"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	pb "github.com/TheGrizzlyDev/rules_integration_tests/proto/containers"
)

type containerManagerServer struct {
	pb.UnimplementedContainerManagerServer
}

func newContainerManagerServer() (*containerManagerServer, error) {
	return &containerManagerServer{}, nil
}

func (c *containerManagerServer) Start(context.Context, *pb.StartRequest) (*pb.StartResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Start not implemented")
}

func (c *containerManagerServer) Kill(context.Context, *pb.KillRequest) (*pb.KillResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Kill not implemented")
}

func (c *containerManagerServer) Inspect(context.Context, *pb.InspectRequest) (*pb.InspectResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Inspect not implemented")
}
