package pipeline

import (
	"github.com/kirkalyn13/xyz-books-pipeline/internal/service"
	"github.com/kirkalyn13/xyz-books-pipeline/internal/writer"
)

// UpdateISBNs retrieves book data and updates the CSV file
func UpdateISBNs() error {
	response, err := service.Get()

	if err != nil {
		return err
	}

	newData := []string{response.Title}
	err = writer.WriteCsv(newData)

	if err != nil {
		return err
	}

	return nil
}
