package rabbit

import (
	"fmt"
	"log"

	"github.com/streadway/amqp"
)

func SendMessage(json string) {

	log.Printf("Received message %s\n", json)
	log.Printf("Connection to rabbitmq")

	connection, err := amqp.Dial("amqp://thedude:opinion@localhost:5672/")
	failOnError(err, "Failed to connect to RabbitMQ")
	defer connection.Close()

	channel, err := connection.Channel()
	failOnError(err, "Failed to open channel")
	defer channel.Close()

	queue, err := channel.QueueDeclare(
		"messages",
		false,
		false,
		false,
		false,
		nil,
	)
	failOnError(err, "Failed to declare queue")

	log.Printf("Publishing message to rabbitmq")
	err = channel.Publish(
		"", //exchange
		queue.Name,
		false,
		false,
		amqp.Publishing{
			ContentType: "text/json",
			Body:        []byte(json),
		})
	failOnError(err, "Failed to publish a message")
}

func failOnError(err error, msg string) {
	if err != nil {
		log.Fatalf("%s: %s", msg, err)
		panic(fmt.Sprintf("%s: %s", msg, err))
	}
}
