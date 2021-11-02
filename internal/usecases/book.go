package usecases

import (
	"context"
	"github.com/blr-coder/book_grpc/internal/domain/models"
	"github.com/blr-coder/book_grpc/internal/domain/repository_interfaces"
)

type BookUseCase struct {
	bookRepository repository_interfaces.IBookRepository
}

func NewBookUseCase(bookRepository repository_interfaces.IBookRepository) *BookUseCase {
	return &BookUseCase{bookRepository: bookRepository}
}

func (u *BookUseCase) Create(ctx context.Context, createArgs *models.CreateBookArgs) (*models.Book, error) {
	if err := createArgs.Validate(); err != nil {
		return nil, err
	}

	book, err := u.bookRepository.Create(ctx, &models.Book{
		Title:       createArgs.Title,
		Description: createArgs.Description,
	})
	if err != nil {
		return nil, err
	}

	return book, nil
}

func (u *BookUseCase) Get(ctx context.Context, id int64) (*models.Book, error) {
	panic("GET!")
}

func (u *BookUseCase) List(ctx context.Context) (models.Books, error) {
	panic("LIST!")
}