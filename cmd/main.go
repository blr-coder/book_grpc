package main

import (
	v1 "github.com/blr-coder/book_grpc/api/v1"
	delivery "github.com/blr-coder/book_grpc/internal/delivery/grpc"
	"github.com/blr-coder/book_grpc/internal/repositories"
	"google.golang.org/grpc"
	"log"
	"net"
)

func main() {
	log.Println("Go GRPC")

	server := grpc.NewServer()
	bookRepository := repositories.NewBookRepository()
	bookGRPCServer := &delivery.NewBookGRPCServer(bookRepository)
	v1.RegisterBookServer(server, bookGRPCServer)

	listener, err := net.Listen("tcp", ":8040")
	if err != nil {
		log.Fatalln(err)
	}

	err = server.Serve(listener)
	if err != nil {
		log.Fatal(err)
	}
}
