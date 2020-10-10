package main

import (
	"fmt"
	"io"
	"log"
	"net"

	numbers "github.com/sibis/grpc-go/protos"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

type server struct{}

func main() {
	ls, err := net.Listen("tcp", ":5000")
	if err != nil {
		log.Fatalf("could not listen: %v", err)
	}

	gserver := grpc.NewServer()

	numbers.RegisterNumberServer(gserver, &server{})
	reflection.Register(gserver)

	gserver.Serve(ls)
}

func (*server) FindBiggestNumCS(stream numbers.Number_FindBiggestNumCSServer) error {
	var max int32
	var incomingNumber int32

	for {
		msg, err := stream.Recv()
		if err == io.EOF {
			fmt.Println("Sent max num back to the client")
			return stream.SendAndClose(&numbers.NumberRequest{Number: max})
		}

		if err != nil {
			log.Fatalf("could not recieve stream: %v", err)
		}
		incomingNumber = msg.GetNumber()
		if incomingNumber > max {
			max = incomingNumber
		}
	}
}

func (*server) FindBiggest(stream numbers.Number_FindBiggestServer) error {
	return nil
}
