package writer

import (
	"encoding/csv"
	"log"
	"os"
	"time"
)

var (
// header = []string{"timestamp", "title", "authors", "isbn13", "isbn10", "publicationYear", "publisher", "edition", "price"}
)

// WriteCsv writes new data to data CSV File
func WriteCsv(newData []string, csvFile string, header []string) error {
	if !dataCsvExists(csvFile) {
		err := initCsvFile(csvFile, header)

		if err != nil {
			return err
		}
	}

	err := updateCsvFile(newData, csvFile)

	if err != nil {
		return err
	}

	return nil
}

// dataCsvExists checks if the data csv exists
func dataCsvExists(csvFile string) bool {
	if _, err := os.Stat(csvFile); os.IsNotExist(err) {
		return false
	}
	return true
}

// initCsvFile initializes book data csv file
func initCsvFile(csvFile string, header []string) error {
	data := [][]string{header}

	file, err := os.Create(csvFile)

	if err != nil {
		return err
	}
	defer file.Close()

	writer := csv.NewWriter(file)
	defer writer.Flush()

	for _, row := range data {
		err := writer.Write(row)
		if err != nil {
			return err
		}
	}

	log.Println("Initialized Data CSV File.")
	return nil
}

// updateCsvFile appends new data to csv file
func updateCsvFile(newData []string, csvFile string) error {
	file, err := os.OpenFile(csvFile, os.O_WRONLY|os.O_APPEND, 0644)

	if err != nil {
		return err
	}
	defer file.Close()

	writer := csv.NewWriter(file)
	defer writer.Flush()

	timestamp := time.Now().Format(time.RFC3339)
	newData = append([]string{timestamp}, newData...)

	err = writer.Write(newData)
	if err != nil {
		return err
	}

	log.Println("Updated CSV File.")
	return nil
}
