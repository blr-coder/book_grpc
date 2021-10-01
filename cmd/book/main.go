package main

import (
	"fmt"
	v1 "github.com/blr-coder/book_grpc/api/v1"
	"github.com/blr-coder/book_grpc/internal/db"
	delivery "github.com/blr-coder/book_grpc/internal/delivery/grpc"
	"github.com/blr-coder/book_grpc/internal/repositories"
	"google.golang.org/grpc"
	"log"
	"net"
)

func main() {
	log.Println("Go GRPC")

	grpcServer := grpc.NewServer()

	psqlDB, err := db.NewDBClient()
	if err != nil {
		fmt.Println("failed to connect to postgres")
		log.Fatalln(err)
	}

	bookRepository := repositories.NewBookRepository(psqlDB)
	bookGRPCServer := delivery.NewBookGRPCServer(*bookRepository)
	v1.RegisterBookServer(grpcServer, bookGRPCServer)

	listener, err := net.Listen("tcp", ":8040")
	if err != nil {
		log.Fatalln(err)
	}

	err = grpcServer.Serve(listener)
	if err != nil {
		log.Fatal(err)
	}
}
