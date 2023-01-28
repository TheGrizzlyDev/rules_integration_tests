package grpc

import (
	pb "github.com/TheGrizzlyDev/rules_integration_tests/proto/containers"
)

type ContainerManagerServer struct {
	pb.InspectRequest
}
