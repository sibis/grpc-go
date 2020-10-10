package main

import (
	"context"
	"fmt"
	"math/rand"
	"time"

	numbers "github.com/sibis/grpc-go/protos"
	"google.golang.org/grpc"
)

func main() {
	conn, err := grpc.Dial("localhost:5000", grpc.WithInsecure())
	if err != nil {
		panic(err)
	}
	defer conn.Close()

	cc := numbers.NewNumberClient(conn)

	s1 := rand.NewSource(time.Now().UnixNano())
	r1 := rand.New(s1)

	nr := &numbers.NumberRequest{
		Number: int32(r1.Intn(200)),
	}
	resp, err := cc.FindBiggestNum(context.Background(), nr)

	if err != nil {
		fmt.Println("[error] getting new number ", err)
	}
	fmt.Println("resp data : ", resp.Number)
}
