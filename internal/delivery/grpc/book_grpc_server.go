package grpc

import (
	"context"
	v1 "github.com/blr-coder/book_grpc/api/v1"
	"github.com/blr-coder/book_grpc/internal/domain/models"
	"github.com/blr-coder/book_grpc/internal/repositories"
)

type BookGRPCServer struct{
	bookRepository repositories.BookRepository
	v1.UnimplementedBookServer
}

func NewBookGRPCServer(bookRepository repositories.BookRepository) *BookGRPCServer {
	return &BookGRPCServer{bookRepository: bookRepository}
}

func (s *BookGRPCServer) Create(ctx context.Context, request *v1.CreateBookRequest) (*v1.Book, error) {

	book, err := s.bookRepository.Create(ctx, &models.Book{
		Title:       request.Title,
		Description: request.Description,
	})
	if err != nil {
		return nil, err
	}

	return &v1.Book{
		Id:          book.Id,
		Title:       book.Title,
		Description: book.Description,
	}, nil
}

func (s *BookGRPCServer) Get (ctx context.Context, request *v1.GetBookRequest) (*v1.Book, error) {

	book, err := s.bookRepository.Get(ctx, request.GetId())
	if err != nil {
		return nil, err
	}

	return &v1.Book{
		Id:          book.Id,
		Title:       book.Title,
		Description: book.Description,
	}, nil
}
func (s *BookGRPCServer) List(ctx context.Context, request *v1.ListBookRequest) (*v1.Books, error) {
	books, err := s.bookRepository.List(ctx)
	if err != nil {
		return nil, err
	}

	var grpcBooks []*v1.Book
	for _, book := range books {
		grpcBooks = append(grpcBooks, &v1.Book{
			Id:          book.Id,
			Title:       book.Title,
			Description: book.Description,
		})
	}

	return &v1.Books{Books: grpcBooks}, nil
}
