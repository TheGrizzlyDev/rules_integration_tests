package main

import (
	"flag"
	"fmt"
	"log"
	"net"

	pb "github.com/TheGrizzlyDev/rules_integration_tests/proto/containers"
)

var port = flag.Int("port", 12345, "")

func main() {
	flag.Parse()
	lis, err := net.Listen("tcp", fmt.Sprintf("localhost:%d", *port))
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	var opts []grpc.ServerOption
	grpcServer := grpc.NewServer(opts...)
	pb.RegisterRouteGuideServer(grpcServer, newServer())
	grpcServer.Serve(lis)
}
