package main

import (
	"fmt"
	"net"

	pb "github.com/khisby/go-grpc-example/grpc/User"
	userService "github.com/khisby/go-grpc-example/services"
	"google.golang.org/grpc"
)

func main() {
	listen, err := net.Listen("tcp", ":9090")
	if err != nil {
		panic(err)
	}

	s := grpc.NewServer()
	pb.RegisterUserServiceServer(s, userService.NewUser())
	fmt.Println("Server is running at port :9090")
	if err := s.Serve(listen); err != nil {
		panic(err)
	}
}
