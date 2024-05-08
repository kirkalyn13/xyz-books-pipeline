package main

import (
	"log"

	"github.com/kirkalyn13/xyz-books-pipeline/pkg/mq"
)

func main() {
	log.Println("Starting XYZ Books Pipeline")
	log.Println("Waiting for data update...")

	mq.InitSubscriber("xyz-books")
}
