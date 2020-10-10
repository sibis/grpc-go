package main

import (
	"context"
	"fmt"
	"log"
	"math/rand"
	"time"

	numbers "github.com/sibis/grpc-go/protos"
	"google.golang.org/grpc"
)

func main() {
	conn, err := grpc.Dial("localhost:5000", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("could not connect: %v", err)
	}
	defer conn.Close()

	c := numbers.NewNumberClient(conn)

	stream, err := c.FindBiggestNumCS(context.Background())
	if err != nil {
		log.Fatalf("request failed: %v", err)
	}

	// send 0 to 10 numbers to the stream
	for i := 0; i <= 20; i++ {
		s1 := rand.NewSource(time.Now().UnixNano())
		r1 := rand.New(s1)
		num := int32(r1.Intn(100))
		fmt.Printf("sending %v into the stream\n", num)
		stream.Send(&numbers.NumberRequest{Number: int32(num)})
		time.Sleep(100 * time.Millisecond)
	}

	// close the stream and recieve result
	result, err := stream.CloseAndRecv()
	if err != nil {
		log.Fatalf("failed to recieve response: %v", err)
	}

	fmt.Println("Max num sent to server: ", result.Number)
}
