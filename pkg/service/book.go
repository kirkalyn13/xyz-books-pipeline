package service

import (
	"bytes"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strings"

	"github.com/kirkalyn13/xyz-books-pipeline/internal/writer"
	"github.com/kirkalyn13/xyz-books-pipeline/pkg/isbn"
	"github.com/kirkalyn13/xyz-books-pipeline/pkg/model"
)

const (
	appServer = "http://localhost:8080/api/v1/books"
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
		updateBook(book)
	} else if book.ISBN13 == "" {
		isbn13, err := isbn.ToISBN13(book.ISBN10)

		if err != nil {
			return err
		}

		book.ISBN13 = isbn13
		updateBook(book)
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

// updateBook sends a PUT request to edit book with updated ISBNs
func updateBook(book model.Book) {
	jsonData, err := json.Marshal(book)
	if err != nil {
		log.Printf("Error encoding JSON: %s \n", err)
		return
	}

	url := fmt.Sprintf("%s/%v", appServer, book.ID)
	req, err := http.NewRequest(http.MethodPut, url, bytes.NewBuffer(jsonData))
	if err != nil {
		log.Printf("Error creating request: %s \n", err)
		return
	}

	req.Header.Set("Content-Type", "application/json")

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		log.Printf("Error sending request: %s \n", err)
		return
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		log.Printf("Unexpected response status: %s \n", resp.Status)
		return
	}

	log.Printf("Successfully edited ISBN for %s", book.Title)
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
