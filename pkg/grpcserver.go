package pkg

import (
	"context"
	v1 "github.com/blr-coder/book_grpc/api/v1"
	"github.com/blr-coder/book_grpc/models"
)

type BookGRPCServer struct{
	v1.UnimplementedBookServer
}

func (s *BookGRPCServer) Create(ctx context.Context, request *v1.CreateBookRequest) (*v1.Book, error) {

	book := models.Book{
		Title:       request.Title,
		Description: request.Description,
	}

	return &v1.Book{
		Id:          1,
		Title:       book.Title,
		Description: book.Description,
	}, nil
}

func (s *BookGRPCServer) Get (ctx context.Context, request *v1.GetBookRequest) (*v1.Book, error) {
	return &v1.Book{
		Id:          12,
		Title:       "123",
		Description: "321",
	}, nil
}
