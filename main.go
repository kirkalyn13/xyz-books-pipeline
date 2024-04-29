package main

import (
	"log"

	"github.com/kirkalyn13/xyz-books-pipeline/internal/writer"
)

func main() {
	log.Println("Starting XYZ Books Pipeline")
	log.Println("Waiting for data update...")

	err := writer.WriteCsv()

	if err != nil {
		log.Fatal(err)
	}
}
