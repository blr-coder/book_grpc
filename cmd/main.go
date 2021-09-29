package main

import (
	v1 "github.com/blr-coder/book_grpc/api/v1"
	"github.com/blr-coder/book_grpc/pkg"
	"google.golang.org/grpc"
	"log"
	"net"
)

func main() {
	log.Println("Go GRPC")

	server := grpc.NewServer()
	bookGRPCServer := &pkg.BookGRPCServer{}
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
