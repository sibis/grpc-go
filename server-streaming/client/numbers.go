package main

import (
	"context"
	"fmt"
	"io"
	"log"

	numbers "github.com/sibis/grpc-go/protos"
	"google.golang.org/grpc"
)

func main() {
	conn, err := grpc.Dial("localhost:5000", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("failed to connect: %v", err)
	}

	defer conn.Close()

	c := numbers.NewNumberClient(conn)

	// timer to start countdown for 10 sec
	timer := int32(10)

	// call Start service
	stream, err := c.FindBiggestNumSS(context.Background(), &numbers.NumberRequest{Number: timer})
	if err != nil {
		log.Fatalf("failed to start timer: %v", err)
	}

	for {
		msg, err := stream.Recv()
		if err == io.EOF {
			break
		}
		if err != nil {
			log.Fatalf("stream read failed: %v", err)
		}

		fmt.Println(msg.Number)
	}
}
