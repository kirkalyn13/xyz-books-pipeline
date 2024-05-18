package service

import (
	"os"
	"testing"

	"github.com/kirkalyn13/xyz-books-pipeline/pkg/model"
)

var (
	publisherID = uint(1)
	testBook    = model.Book{
		ID:              6,
		Title:           "Test Book",
		ISBN13:          "9780547928227",
		ISBN10:          "054792822X",
		ListPrice:       1137,
		PublicationYear: 2020,
		ImageURL:        "",
		Edition:         "First Edition",
		PublisherID:     &publisherID,
		Authors: []*model.Author{
			{ID: 1},
		},
	}
	testCsv = "test.csv"
)

func TestBookToCSV(t *testing.T) {
	err := bookToCSV(testBook, testCsv)

	if err != nil {
		t.Errorf("Test Failed: Error occurred when addding book data to CSV: %v", err)
	}

	if _, err := os.Stat(testCsv); os.IsNotExist(err) {
		t.Error("Test Failed: Expected CSV output file not found.")
	}

	os.Remove(testCsv)
}
