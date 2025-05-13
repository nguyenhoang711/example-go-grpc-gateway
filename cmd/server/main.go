package main

import (
	"log"
	"net"

	"github.com/nguyenhoang711/example-go-grpc-gateway/internal"
	"github.com/nguyenhoang711/example-go-grpc-gateway/protogen/golang/orders"
	"google.golang.org/grpc"
)

func main() {
	const addr = "0.0.0.0:50051"
	listener, err := net.Listen("tcp", addr)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	// create a gRPC server instance
	server := grpc.NewServer()

	// create an order service instance with a reference to the DB
	db := internal.NewDB()
	orderService := internal.NewOrderService(db)

	// register the order service with the grpc server
	orders.RegisterOrdersServer(server, &orderService)

	// start listening to requests
	log.Printf("server listening at %v", listener.Addr())
	if err = server.Serve(listener); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
