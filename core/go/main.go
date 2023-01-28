package main

import (
	"flag"
	"fmt"
	"log"
	"net"

	"github.com/TheGrizzlyDev/rules_integration_tests/core/go/grpc"
)

var port = flag.Int("port", 12345, "")

func main() {
	flag.Parse()
	lis, err := net.Listen("tcp", fmt.Sprintf("localhost:%d", *port))
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	log.Printf("Listening on port %d", *port)

	if server, err := grpc.Setup(); err != nil {
		panic(err)
	} else {
		server.Serve(lis)
	}
}
