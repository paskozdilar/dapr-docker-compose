// Package main implements a server for Greeter service.
package main

import (
	"flag"
	"fmt"
	"io"
	"log"
	"net"
	"os"
	"sync/atomic"

	"google.golang.org/grpc"
	pb "google.golang.org/grpc/examples/helloworld/helloworld"
)

var (
	port        = flag.Int("port", 50051, "The server port")
	healthcheck = flag.Bool("healthcheck", false, "Perform self healthcheck")

	counter = atomic.Int32{}
)

// server is used to implement helloworld.GreeterServer.
type server struct {
	pb.UnimplementedGreeterServer
}

// SayHello implements helloworld.GreeterServer
func (s *server) SayHello(stream pb.Greeter_SayHelloServer) error {
	id := counter.Add(1)
	log.Printf("Stream open: %v", id)
	defer log.Printf("Stream closed: %v", id)
	for {
		in, err := stream.Recv()
		if err == io.EOF {
			return nil
		}
		if err != nil {
			return err
		}
		log.Printf("Received: %v", in.GetName())
		err = stream.Send(&pb.HelloReply{Message: "Hello " + in.GetName()})
		if err != nil {
			return err
		}
	}
}

func main() {
	flag.Parse()

	if *healthcheck {
		conn, err := net.Dial("tcp", fmt.Sprintf(":%d", *port))
		if err != nil {
			os.Exit(1)
		}
		conn.Close()
		os.Exit(0)
	}

	lis, err := net.Listen("tcp", fmt.Sprintf(":%d", *port))
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	s := grpc.NewServer()
	pb.RegisterGreeterServer(s, &server{})
	log.Printf("server listening at %v", lis.Addr())
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
