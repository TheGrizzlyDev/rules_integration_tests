package grpc

import (
	"google.golang.org/grpc"

	"github.com/TheGrizzlyDev/rules_integration_tests/proto/containers"
)

func Setup() (*grpc.Server, error) {
	server := grpc.NewServer()

	if containerServer, err := newContainerManagerServer(); err != nil {
		return nil, err
	} else {
		containers.RegisterContainerManagerServer(server, containerServer)
	}

	return server, nil
}
