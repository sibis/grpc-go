package main

import (
	"context"
	"fmt"
	"io"
	"log"
	"math/rand"
	"time"

	numbers "github.com/sibis/grpc-go/protos"
	"google.golang.org/grpc"
)

func main() {
	conn, err := grpc.Dial("localhost:5000", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("Could not connect to server: %v", err)
	}

	c := numbers.NewNumberClient(conn)
	fmt.Println("C : ", c)
	stream, err := c.FindBiggest(context.Background())
	fmt.Println("stream : ", stream)
	if err != nil {
		log.Fatalf("Failed to call the FindBiggest method: %v", err)
	}

	// make blocking channel
	waitc := make(chan struct{})

	go func() {
		i := 0
		for {
			if i < 11 {
				s1 := rand.NewSource(time.Now().UnixNano())
				r1 := rand.New(s1)
				num := int32(r1.Intn(100))
				err = stream.Send(&numbers.NumberRequest{Number: num})
				fmt.Println("Number sent to server : ", num)
				if err != nil {
					log.Fatalf("Error while sending the number stream to server : %v", err)
				}
				time.Sleep(time.Second)
				i++
			} else {
				break
			}
		}
		err := stream.CloseSend()
		if err != nil {
			log.Fatalf("Failed to close the stream : %v", err)
		}
	}()

	go func() {
		for {
			incomingMessage, err := stream.Recv()
			if err == io.EOF {
				close(waitc)
				log.Fatalf("Error occurred while receiving EOF : %v", err)
			}
			if err != nil {
				close(waitc)
				log.Fatalf("Error occurred while receiving from the server : %v", err)
			}

			fmt.Println("--------")
			fmt.Println(incomingMessage.GetNumber())
			fmt.Println("--------")
		}
	}()
	<-waitc
}
