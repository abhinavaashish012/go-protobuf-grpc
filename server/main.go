package main

import (
	"flag"
	"fmt"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
	"log"
	"net"
	"pcbook/pb"
	"pcbook/service"
)

func main() {

	laptopStore := service.NewInMemoryLaptopStore()

	laptopServer := service.NewLaptopServer(laptopStore)

	serverOptions := []grpc.ServerOption{}
	grpcServer := grpc.NewServer(serverOptions...)

	port := flag.Int("port", 8080, "the server port")
	flag.Parse()
	log.Printf("start server on port %d", *port)

	reflection.Register(grpcServer)
	pb.RegisterLaptopServiceServer(grpcServer, laptopServer)

	address := fmt.Sprintf("0.0.0.0:%d", *port)
	listener, err := net.Listen("tcp", address)
	if err != nil {
		log.Fatal("cannot start server: ", err)
	}

	err = grpcServer.Serve(listener)
	if err != nil {
		log.Fatal("cannot start server: ", err)
	}

}
