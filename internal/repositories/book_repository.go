package repositories

import (
	"context"
	"fmt"
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
	query := `SELECT id, title, description, created_at, updated_at FROM books WHERE id=$1`

	book := models.Book{}
	if err := r.db.GetContext(ctx, &book, query, id); err != nil {
		return nil, err
	}

	return &book, nil
}

func (r *BookRepository) List(ctx context.Context, filter *models.BookListFilter) (models.Books, error) {
	ctx, cancel := context.WithTimeout(ctx, timeout)
	defer cancel()

	query := `SELECT * FROM books`
	query, args, err := r.decodeFilter(query, filter)
	if err != nil {
		return nil, err
	}

	var books models.Books
	err = r.db.SelectContext(ctx, &books, query, args...)
	if err != nil {
		return nil, err
	}

	return books, nil
}

func (r *BookRepository) Count(ctx context.Context, filter *models.BookListFilter) (uint64, error) {
	ctx, cancel := context.WithTimeout(ctx, timeout)
	defer cancel()

	query := `SELECT count(*) FROM books`
	query, args, err := r.decodeFilter(query, filter)

	var count uint64
	err = r.db.GetContext(ctx, &count, query, args...)
	if err != nil {
		return 0, err
	}

	return count, nil
}

func (r *BookRepository) Update(ctx context.Context, book *models.Book) (*models.Book, error) {
	ctx, cancel := context.WithTimeout(ctx, timeout)
	defer cancel()

	query := `UPDATE books SET title=$1, description=$2, updated_at=$3 WHERE id=$4 RETURNING *`

	rows, err := r.db.QueryxContext(ctx, query, book.Title, book.Description, time.Now().UTC(), book.Id)
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

func (r *BookRepository) Delete(ctx context.Context, id int64) error {
	ctx, cancel := context.WithTimeout(ctx, timeout)
	defer cancel()

	query := `DELETE FROM books WHERE id=$1`
	_, err := r.db.ExecContext(ctx, query, id)
	if err != nil {
		return err
	}

	return nil
}

func (r *BookRepository) decodeFilter(query string, filter *models.BookListFilter) (string, []interface{}, error) {
	query = fmt.Sprintf("%s WHERE 1=1", query)
	var args []interface{}

	if filter.Title != "" {
		title := fmt.Sprintf("%%%s%%", filter.Title)
		query = fmt.Sprintf("%s AND title ILIKE (?)", query)
		args = append(args, title)
	}

	if filter.PageSize < 1 {
		filter.PageSize = 10
	}
	if filter.PageNumber == 0 {
		filter.PageNumber = 1
	}
	if filter.PageNumber > 1 {
		query = fmt.Sprintf("%s OFFSET %d", query, (filter.PageNumber-1)*filter.PageSize)
	}
	query = fmt.Sprintf("%s LIMIT %d", query, filter.PageSize)

	query = r.db.Rebind(query)
	return query, args, nil
}
