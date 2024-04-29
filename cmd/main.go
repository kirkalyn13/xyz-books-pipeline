package main

import (
	"log"

	"github.com/kirkalyn13/xyz-books-pipeline/pkg/pipeline"
)

func main() {
	log.Println("Starting XYZ Books Pipeline")
	log.Println("Waiting for data update...")

	err := pipeline.UpdateISBNs()

	if err != nil {
		log.Fatal(err)
	}
}
