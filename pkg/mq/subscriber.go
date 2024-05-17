package mq

import (
	"log"

	"github.com/kirkalyn13/xyz-books-pipeline/pkg/service"
	"github.com/streadway/amqp"
)

func InitSubscriber(queueName string) {
	conn, err := amqp.Dial(Server)

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
			err := service.UpdateISBNs(d.Body)

			if err != nil {
				log.Printf("Error when updating ISBNs: %s \n", err)
			}
		}
	}()

	log.Println("Successfully connected to Rabbit MQ instance.")
	log.Println(" [*] - Waiting for messages...")
	<-forever
}

// CheckMQ checks if RabbitMQ is up and running
func CheckMQ(url string) bool {
	conn, err := amqp.Dial(url)

	if err != nil {
		return false
	}
	defer conn.Close()

	return true
}
