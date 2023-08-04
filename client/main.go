// Package main implements a client for Greeter service.
package main

import (
	"context"
	"flag"
	"io"
	"log"
	"time"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	pb "google.golang.org/grpc/examples/helloworld/helloworld"
	"google.golang.org/grpc/metadata"
)

const (
	defaultAppId = "server"
)

var (
	addr  = flag.String("addr", "localhost:50051", "the address to connect to")
	appId = flag.String("app-id", defaultAppId, "Server app ID")

	names = []string{
		"antimony", "arsenic", "aluminum", "selenium", "hydrogen", "oxygen",
		"nitrogen", "rhenium", "nickel", "neodymium", "neptunium", "germanium",
		"iron", "americium", "ruthenium", "uranium", "europium", "zirconium",
		"lutetium", "vanadium", "lanthanum", "osmium", "astatine", "radium",
		"gold", "protactinium", "indium", "gallium", "iodine", "thorium",
		"thulium", "thallium", "yttrium", "ytterbium", "actinium", "rubidium",
		"boron", "gadolinium", "niobium", "iridium", "strontium", "silicon",
		"silver", "samarium", "bismuth", "bromine", "lithium", "beryllium",
		"barium", "holmium", "helium", "hafnium", "erbium", "phosphorus",
		"francium", "fluorine", "terbium", "manganese", "mercury",
		"molybdenum", "magnesium", "dysprosium", "scium", "cerium", "cesium",
		"lead", "praseodymium", "platinum", "plutonium", "palladium",
		"promethium", "potassium", "polonium", "tantalum", "technetium",
		"titanium", "tellurium", "cadmium", "calcium", "chromium", "curium",
		"sulfur", "californium", "fermium", "berkelium", "also", "mendelevium",
		"einsteinium", "nobelium", "argon", "krypton", "neon", "radon",
		"xenon", "zinc", "rhodium", "chlorine", "carbon", "cobalt", "copper",
		"tungsten", "tin", "sodium",
	}
)

func main() {
	var (
		conn *grpc.ClientConn
		err  error
	)

	flag.Parse()

	// Set up a connection to the server.
	for {
		log.Printf("attempting to connect to: %v", *addr)
		conn, err = grpc.Dial(*addr,
			grpc.WithTransportCredentials(insecure.NewCredentials()),
			grpc.WithBlock())
		if err != nil {
			log.Printf("did not connect: %v", err)
			time.Sleep(time.Second)
		} else {
			break
		}
	}
	defer conn.Close()
	c := pb.NewGreeterClient(conn)

	// Contact the server and print out its response.
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	ctx = metadata.AppendToOutgoingContext(ctx, "dapr-app-id", *appId)
	ctx = metadata.AppendToOutgoingContext(ctx, "dapr-stream", "true")
	defer cancel()

	stream, err := c.SayHello(ctx)
	if err != nil {
		log.Fatalf("could not greet: %v", err)
	}

	// Make goroutine wait for main
	waitc := make(chan struct{})
	defer func() {
		log.Println("Sent all requests - waiting for responses...")
		<-waitc
	}()

	go func() {
		defer func() { close(waitc) }()
		for {
			in, err := stream.Recv()
			if err == io.EOF {
				return
			}
			if err != nil {
				log.Fatalf("could not recv greet: %v", err)
			}
			log.Printf("Greeting: %s", in.GetMessage())
		}
	}()

	for _, name := range names {
		err := stream.Send(&pb.HelloRequest{Name: name})
		if err != nil {
			log.Fatalf("could not greet: %v", err)
		}
	}
	stream.CloseSend()
}
