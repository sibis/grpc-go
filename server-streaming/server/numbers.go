package main

import (
	"fmt"
	"log"
	"net"
	"time"

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

func (*number) FindBiggestNumSS(req *numbers.NumberRequest, stream numbers.Number_FindBiggestNumSSServer) error {
	t := req.GetNumber()
	for t > 0 {
		res := numbers.NumberRequest{Number: t}
		stream.Send(&res)
		t--
		time.Sleep(time.Second)
	}
	fmt.Println("Stream over from server end")
	return nil
}

func (*number) FindBiggest(stream numbers.Number_FindBiggestServer) error {
	return nil
}

func (*number) FindBiggestNumCS(stream numbers.Number_FindBiggestNumCSServer) error {
	return nil
}
