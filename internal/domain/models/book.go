package models

import (
	"database/sql"
	"time"

	validation "github.com/go-ozzo/ozzo-validation/v4"
)

type Book struct {
	ID          int64  `json:"id" db:"id"`
	Title       string `json:"title" db:"title"`
	Description string `json:"description" db:"description"`
	/*PDFFileURL  string `json:"pdf_file_url"`
	IMGFileURL  string `json:"img_file_url"`
	Author      Author `json:"author"`*/

	CreatedAt time.Time    `json:"created_at" db:"created_at"`
	UpdatedAt *time.Time   `json:"updated_at" db:"updated_at"`
	DeletedAt sql.NullTime `json:"deleted_at" db:"deleted_at"`
}

type Books []*Book

type CreateBookArgs struct {
	Title       string `json:"title"`
	Description string `json:"description"`
	/*PDFFileURL  string `json:"pdf_file_url"`
	IMGFileURL  string `json:"img_file_url"`
	Author      Author `json:"author"`*/
}

func (args *CreateBookArgs) Validate() error {
	err := validation.ValidateStruct(args,
		validation.Field(&args.Title, validation.Required),
		validation.Field(&args.Description, validation.Required),
		// nolint
		//validation.Field(&args.IMGFileURL, is.URL),
	)
	if err != nil {
		return err
	}

	return nil
}

type UpdateBookArgs struct {
	ID          int64  `json:"id"`
	Title       string `json:"title"`
	Description string `json:"description"`
}

func (args *UpdateBookArgs) Validate() error {
	err := validation.ValidateStruct(args,
		validation.Field(&args.ID, validation.Required),
		validation.Field(&args.Title, validation.Required),
		validation.Field(&args.Description, validation.Required),
	)
	if err != nil {
		return err
	}

	return nil
}

type BookListFilter struct {
	PageNumber uint64
	PageSize   uint64
	Title      string
}
