package models

import (
	validation "github.com/go-ozzo/ozzo-validation/v4"
)

type Book struct {
	Id          int64  `json:"id"`
	Title       string `json:"title"`
	Description string `json:"description"`
	/*PDFFileURL  string `json:"pdf_file_url"`
	IMGFileURL  string `json:"img_file_url"`
	Author      Author `json:"author"`*/
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
