package writer

import (
	"encoding/csv"
	"log"
	"os"
	"time"
)

var (
	header = []string{"timestamp", "title", "authors", "isbn13", "isbn10", "publicationYear", "publisher", "edition", "price"}
)

// WriteCsv writes new data to data CSV File
func WriteCsv(newData []string) error {
	if !dataCsvExists() {
		err := initCsvFile()

		if err != nil {
			return err
		}
	} else {
		err := updateCsvFile(newData)

		if err != nil {
			return err
		}
	}

	return nil
}

// dataCsvExists checks if the data csv exists
func dataCsvExists() bool {
	if _, err := os.Stat(csvFile); os.IsNotExist(err) {
		return false
	}
	return true
}

// initCsvFile initializes book data csv file
func initCsvFile() error {
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
func updateCsvFile(newData []string) error {
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
