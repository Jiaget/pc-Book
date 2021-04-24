package main

import (
	"flag"
	"fmt"
	"log"
	"net"

	"github.jiaget.com/pc-book/pb"
	"github.jiaget.com/pc-book/service"
	"google.golang.org/grpc"
)

func main() {
	// get the port from command flag
	port := flag.Int("port", 0, "server port")
	flag.Parse()
	log.Printf("start server on port %d", *port)

	laptopStore := service.NewInMemoryLaptopStore()
	imageStore := service.NewDiskImageStore("img")
	rateStore := service.NewInMemoryRateStore()
	laptopServer := service.NewLaptopServer(laptopStore, imageStore, rateStore)
	grpcServer := grpc.NewServer()
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
