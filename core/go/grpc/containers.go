package grpc

import (
	pb "github.com/TheGrizzlyDev/rules_integration_tests/proto/containers"
)

type containerManagerServer struct {
	pb.UnimplementedContainerManagerServer
}

func newContainerManagerServer() (*containerManagerServer, error) {
	return &containerManagerServer{}, nil
}
