package main

import (
	"log"

	"github.com/kirkalyn13/xyz-books-pipeline/pkg/mq"
	"github.com/kirkalyn13/xyz-books-pipeline/pkg/service"
)

func main() {
	log.Println("Starting XYZ Books Pipeline")
	log.Println("Waiting for data update...")

	go service.EvaluateISBNs()

	if mq.CheckMQ(mq.Server) {
		mq.InitSubscriber("xyz-books")
	}
}
