package main

import (
	"log"
	"sync"

	"github.com/kirkalyn13/xyz-books-pipeline/pkg/mq"
	"github.com/kirkalyn13/xyz-books-pipeline/pkg/service"
)

func main() {
	log.Println("Starting XYZ Books Pipeline")
	log.Println("Waiting for data update...")

	var wg sync.WaitGroup

	go service.EvaluateISBNs(&wg)

	if mq.CheckMQ(mq.Server) {
		mq.InitSubscriber("xyz-books", &wg)
	}

	wg.Wait()
}
