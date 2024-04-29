package writer

import (
	"encoding/csv"
	"log"
	"os"
	"time"
)

var (
	header = []string{"timestamp", "title", "author", "isbn13", "isbn10", "publicationYear", "publisher", "edition", "price"}
)

func WriteCsv() error {
	if !dataCsvExists() {
		err := initCsvFile()

		if err != nil {
			return err
		}
	} else {
		newData := []string{"American Elf", "Joel Hartse, Hannah P. Templer", "9781891830853", "1891830856", "2004", "Paste Magazine", "Book 2", "1000"}
		err := updateCsvFile(newData)

		if err != nil {
			return err
		}
	}

	return nil
}

func dataCsvExists() bool {
	if _, err := os.Stat(csvFile); os.IsNotExist(err) {
		return false
	}
	return true
}

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
