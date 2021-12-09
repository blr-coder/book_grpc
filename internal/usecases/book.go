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
	return u.bookRepository.Get(ctx, id)
}

func (u *BookUseCase) List(ctx context.Context, filter *models.BookListFilter) (models.Books, uint64, error) {
	books, err := u.bookRepository.List(ctx, filter)
	if err != nil {
		return nil, 0, err
	}

	count, err := u.bookRepository.Count(ctx, filter)
	if err != nil {
		return nil, 0, err
	}
	return books, count, nil
}

func (u *BookUseCase) Update(ctx context.Context, updateArgs *models.UpdateBookArgs) (*models.Book, error) {
	if err := updateArgs.Validate(); err != nil {
		return nil, err
	}

	updatedBook, err := u.bookRepository.Update(ctx, &models.Book{
		Id:          updateArgs.ID,
		Title:       updateArgs.Title,
		Description: updateArgs.Description,
	})
	if err != nil {
		return nil, err
	}

	return updatedBook, nil
}

func (u *BookUseCase) Delete(ctx context.Context, id int64) error {
	return u.bookRepository.Delete(ctx, id)
}
