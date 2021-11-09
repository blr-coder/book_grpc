package repositories

import (
	"context"
	"github.com/blr-coder/book_grpc/internal/domain/models"
	"github.com/jmoiron/sqlx"
	"time"
)

const (
	timeout = 5 * time.Second
)


type BookRepository struct {
	db *sqlx.DB
}

func NewBookRepository(db *sqlx.DB) *BookRepository {
	return &BookRepository{db: db}
}

func (r *BookRepository) Create(ctx context.Context, book *models.Book) (*models.Book, error) {
	ctx, cancel := context.WithTimeout(ctx, timeout)
	defer cancel()

	query := `INSERT INTO "books" ("title", description) VALUES ($1, $2) RETURNING id`

	rows, err := r.db.QueryxContext(ctx, query, book.Title, book.Description)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	for rows.Next() {
		if err = rows.StructScan(&book); err != nil {
			return nil, err
		}
	}

	return book, nil
}

func (r *BookRepository) Get(ctx context.Context, id int64) (*models.Book, error) {
	query := `SELECT id, title, description, created_at FROM books WHERE id=$1`

	book := models.Book{}
	if err := r.db.Get(&book, query, id); err != nil {
		return nil, err
	}

	return &book, nil
}

func (r *BookRepository) List(ctx context.Context) (models.Books, error) {
	ctx, cancel := context.WithTimeout(ctx, timeout)
	defer cancel()

	query := `SELECT id, title, description FROM books`

	var books models.Books
	err := r.db.SelectContext(ctx, &books, query); if err != nil {
		return nil, err
	}

	return books, nil
}
