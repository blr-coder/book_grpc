package usecaseinterfaces

import (
	"context"

	"github.com/blr-coder/book_grpc/internal/domain/models"
)

type IBookUseCase interface {
	Create(ctx context.Context, createArgs *models.CreateBookArgs) (*models.Book, error)
	Get(ctx context.Context, id int64) (*models.Book, error)
	List(ctx context.Context, filter *models.BookListFilter) (models.Books, uint64, error)

	Update(ctx context.Context, updateArgs *models.UpdateBookArgs) (*models.Book, error)

	Delete(ctx context.Context, id int64) error
}
