package pipeline

import (
	"github.com/kirkalyn13/xyz-books-pipeline/internal/writer"
	"github.com/kirkalyn13/xyz-books-pipeline/pkg/service"
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
