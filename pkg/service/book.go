package service

import (
	"encoding/json"
	"fmt"
	"log"
	"strings"

	"github.com/kirkalyn13/xyz-books-pipeline/internal/writer"
	"github.com/kirkalyn13/xyz-books-pipeline/pkg/isbn"
	"github.com/kirkalyn13/xyz-books-pipeline/pkg/model"
)

// UpdateISBNs retrieves book data and updates the CSV file
func UpdateISBNs(data []byte) error {
	var book model.Book

	err := json.Unmarshal(data, &book)
	if err != nil {
		log.Printf("Error unmarshalling json data: %v \n", err)
		return err
	}

	if book.ISBN10 == "" {
		isbn10, err := isbn.ToISBN10(book.ISBN13)

		if err != nil {
			return err
		}

		book.ISBN10 = isbn10
	}

	if book.ISBN13 == "" {
		isbn13, err := isbn.ToISBN13(book.ISBN10)

		if err != nil {
			return err
		}

		book.ISBN13 = isbn13
	}

	newData := []string{
		book.Title,
		formatAuthors(book.Authors),
		book.ISBN13,
		book.ISBN10,
		fmt.Sprint(book.PublicationYear),
		book.Publisher.Name,
		book.Edition,
		fmt.Sprint(int(book.ListPrice)),
	}

	err = writer.WriteCsv(newData)

	if err != nil {
		return err
	}

	return nil
}

// formatAuthors formats Book  authors to a single string presentation
func formatAuthors(authors []*model.Author) string {
	var authorList []string

	for _, a := range authors {
		middleName := a.MiddleName

		if middleName == "" {
			middleName = " "
		} else {
			middleName = fmt.Sprintf(" %s ", middleName)
		}

		author := a.FirstName + middleName + a.LastName
		authorList = append(authorList, author)
	}

	return strings.Join(authorList, ", ")
}
