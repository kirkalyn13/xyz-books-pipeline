package mq

import (
	"log"

	"github.com/kirkalyn13/xyz-books-pipeline/pkg/service"
	"github.com/streadway/amqp"
)

func InitSubscriber(queueName string) {
	conn, err := amqp.Dial(server)

	if err != nil {
		log.Println(err)
	}
	defer conn.Close()

	ch, err := conn.Channel()

	if err != nil {
		log.Println(err)
	}
	defer ch.Close()

	msgs, err := ch.Consume(
		queueName,
		"",
		true,
		false,
		false,
		false,
		nil,
	)

	forever := make(chan bool)

	go func() {
		for d := range msgs {
			log.Printf("Received: %s\n", d.Body)
			err := service.UpdateISBNs()

			if err != nil {
				log.Printf("Error when updating ISBNs: %s \n", err)
			}
		}
	}()

	log.Println("Successfully connected to Rabbit MQ instance.")
	log.Println(" [*] - Waiting for messages...")
	<-forever
}
