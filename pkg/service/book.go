package service

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"path/filepath"
	"time"

	"github.com/kirkalyn13/xyz-books-pipeline/internal/writer"
	"github.com/kirkalyn13/xyz-books-pipeline/pkg/isbn"
	"github.com/kirkalyn13/xyz-books-pipeline/pkg/model"
	"github.com/kirkalyn13/xyz-books-pipeline/pkg/util"
)

const (
	appServer = "http://localhost:8080/api/v1/books"
)

var (
	csvFile  = filepath.Join(".", "output", "book-data.csv")
	header   = []string{"timestamp", "title", "authors", "isbn13", "isbn10", "publicationYear", "publisher", "edition", "price"}
	duration = 60 * time.Second // book evaluation duration in seconds
)

// UpdateISBNs retrieves book data and updates the CSV file
func UpdateISBNs(data []byte) error {
	var book model.Book

	err := json.Unmarshal(data, &book)
	if err != nil {
		log.Printf("Error unmarshalling json data: %v \n", err)
		return err
	}

	log.Printf("Received: ID: %v, Title: %s, ISBN13: %s, ISBN10: %s \n", book.ID, book.Title, book.ISBN13, book.ISBN10)

	err = bookToCSV(book, csvFile)

	if err != nil {
		return err
	}

	return nil
}

// EvaluateISBNs reviews existing books ISBN data if needing update, every configured duration time
func EvaluateISBNs() {
	log.Println(" [*] - Waiting for book updates...")
	url := fmt.Sprintf("%s/isbn/incomplete", appServer)

	for {
		go func() {
			books, err := getBooks(url)

			if err != nil {
				log.Printf("Error retrieving books data: %v", err)
			}

			for _, book := range books.Books {
				err := bookToCSV(book, csvFile)

				if err != nil {
					log.Printf("Error updating CSV data: %v \n", err)
					return
				}
			}
		}()

		time.Sleep(duration)
	}
}

// getBooks returns book data response
func getBooks(url string) (model.BooksResponse, error) {
	var books model.BooksResponse

	res, err := http.Get(url)
	if err != nil {
		return model.BooksResponse{}, err
	}

	defer res.Body.Close()

	if res.StatusCode != http.StatusOK {
		return model.BooksResponse{}, err
	}

	body, err := io.ReadAll(res.Body)
	if err != nil {
		return model.BooksResponse{}, err
	}

	err = json.Unmarshal(body, &books)

	if err != nil {
		return model.BooksResponse{}, err
	}

	return books, nil
}

// bookToCSV updates new book data to CSV
func bookToCSV(book model.Book, csvFile string) error {
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
		util.FormatAuthors(book.Authors),
		book.ISBN13,
		book.ISBN10,
		fmt.Sprint(book.PublicationYear),
		book.Publisher.Name,
		book.Edition,
		fmt.Sprint(int(book.ListPrice)),
	}

	err := writer.WriteCsv(newData, csvFile, header)

	if err != nil {
		return err
	}

	log.Printf("Updated ISBNs for: %s \n", book.Title)
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
