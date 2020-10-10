package main

import (
	"context"
	"fmt"
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

func (n *number) FindBiggestNum(ctx context.Context, nr *numbers.NumberRequest) (*numbers.NumberResponse, error) {
	inputNum := nr.GetNumber()
	fmt.Println("Incoming num: ", inputNum)
	if inputNum > 50 || inputNum < 0 {
		inputNum = 0
	}
	fmt.Println("Sending resp: ", inputNum)
	return &numbers.NumberResponse{Number: inputNum}, nil
}

func (*number) FindBiggestNumSS(req *numbers.NumberRequest, stream numbers.Number_FindBiggestNumSSServer) error {
	return nil
}

func (*number) FindBiggest(stream numbers.Number_FindBiggestServer) error {
	return nil
}

func (*number) FindBiggestNumCS(stream numbers.Number_FindBiggestNumCSServer) error {
	return nil
}
