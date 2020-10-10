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

type number struct {
}

func main() {
	ls, err := net.Listen("tcp", ":5000")
	if err != nil {
		log.Fatalf("could not listen: %v", err)
	}

	gserver := grpc.NewServer()

	numbers.RegisterNumberServer(gserver, &number{})
	reflection.Register(gserver)

	gserver.Serve(ls)
}

func (*number) FindBiggest(stream numbers.Number_FindBiggestServer) error {
	var max int32
	var incomingNumber int32
	// read client messages
	go func() {
		for {
			msg, err := stream.Recv()
			if err == io.EOF {
				break
			}
			if err != nil {
				log.Fatalf("Could not receive the stream: %v", err)
				break
			}

			incomingNumber = msg.GetNumber()
			fmt.Println("Incoming number : ", incomingNumber)
		}
	}()
	// send a message to the client every second
	for {
		if incomingNumber > max {
			fmt.Println("Max so far: ", max)
			max = incomingNumber
			stream.Send(&numbers.NumberResponse{Number: max})
		}
	}

	return nil
}

func (*number) FindBiggestNumCS(stream numbers.Number_FindBiggestNumCSServer) error {
	return nil
}
