package usecase_interfaces

import (
	"context"
	"github.com/blr-coder/book_grpc/internal/domain/models"
)

type IBookUseCase interface {
	Create(ctx context.Context, createArgs *models.CreateBookArgs) (*models.Book, error)
	Get(ctx context.Context, id int64) (*models.Book, error)
	List(ctx context.Context) (models.Books, error)
}
