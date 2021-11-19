package grpc

import (
	"context"
	v1 "github.com/blr-coder/book_grpc/api/v1"
	"github.com/blr-coder/book_grpc/internal/domain/models"
	"github.com/blr-coder/book_grpc/internal/domain/usecase_interfaces"
	"github.com/golang/protobuf/ptypes/empty"
	"google.golang.org/protobuf/types/known/timestamppb"
	"time"
)

type BookGRPCServer struct {
	bookUseCase usecase_interfaces.IBookUseCase
	v1.UnimplementedBookServer
}

func NewBookGRPCServer(bookUseCase usecase_interfaces.IBookUseCase) *BookGRPCServer {
	return &BookGRPCServer{bookUseCase: bookUseCase}
}

func (s *BookGRPCServer) Create(ctx context.Context, request *v1.CreateBookRequest) (*v1.Book, error) {
	book, err := s.bookUseCase.Create(ctx, &models.CreateBookArgs{
		Title:       request.Title,
		Description: request.Description,
	})
	if err != nil {
		return nil, err
	}

	return modelsBookToGRPCBook(book), nil
}

func (s *BookGRPCServer) Get(ctx context.Context, request *v1.GetBookRequest) (*v1.Book, error) {
	book, err := s.bookUseCase.Get(ctx, request.GetId())
	if err != nil {
		return nil, err
	}

	return modelsBookToGRPCBook(book), nil
}

func (s *BookGRPCServer) List(ctx context.Context, request *v1.ListBookRequest) (*v1.Books, error) {
	books, err := s.bookUseCase.List(ctx)
	if err != nil {
		return nil, err
	}

	var grpcBooks []*v1.Book
	for _, book := range books {
		grpcBooks = append(grpcBooks, modelsBookToGRPCBook(book))
	}

	return &v1.Books{Books: grpcBooks}, nil
}

func (s *BookGRPCServer) Update(ctx context.Context, request *v1.UpdateBookRequest) (*v1.Book, error) {
	updatedBook, err := s.bookUseCase.Update(ctx, &models.UpdateBookArgs{
		ID:          request.Id,
		Title:       request.Title,
		Description: request.Description,
	})
	if err != nil {
		return nil, err
	}

	return modelsBookToGRPCBook(updatedBook), nil
}

func (s *BookGRPCServer) Delete(ctx context.Context, request *v1.GetBookRequest) (*empty.Empty, error) {
	err := s.bookUseCase.Delete(ctx, request.GetId())
	if err != nil {
		return nil, err
	}

	return &empty.Empty{}, nil
}

func modelsBookToGRPCBook(book *models.Book) *v1.Book {
	return &v1.Book{
		Id:          book.Id,
		Title:       book.Title,
		Description: book.Description,
		CreatedAt:   timestamppb.New(book.CreatedAt),
		UpdatedAt:   nilOrTimestamp(book.UpdatedAt),
	}
}

func nilOrTimestamp(time *time.Time) (timestamp *timestamppb.Timestamp) {
	if time != nil {
		timestamp = timestamppb.New(*time)
	}
	return timestamp
}
