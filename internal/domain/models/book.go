package models

import (
	"database/sql"
	validation "github.com/go-ozzo/ozzo-validation/v4"
	"time"
)

type Book struct {
	Id          int64  `json:"id" db:"id"`
	Title       string `json:"title" db:"title"`
	Description string `json:"description" db:"description"`
	/*PDFFileURL  string `json:"pdf_file_url"`
	IMGFileURL  string `json:"img_file_url"`
	Author      Author `json:"author"`*/

	CreatedAt time.Time    `json:"created_at" db:"created_at"`
	UpdatedAt time.Time    `json:"updated_at" db:"updated_at"`
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
		//validation.Field(&args.IMGFileURL, is.URL),
	)
	if err != nil {
		return err
	}

	return nil
}
