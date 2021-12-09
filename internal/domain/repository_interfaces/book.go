package repository_interfaces

import (
	"context"
	"github.com/blr-coder/book_grpc/internal/domain/models"
)

type IBookRepository interface {
	Create(ctx context.Context, book *models.Book) (*models.Book, error)
	Get(ctx context.Context, id int64) (*models.Book, error)
	List(ctx context.Context, filter *models.BookListFilter) (models.Books, error)
	Count(ctx context.Context, filter *models.BookListFilter) (uint64, error)

	Update(ctx context.Context, book *models.Book) (*models.Book, error)

	Delete(ctx context.Context, id int64) error
}
