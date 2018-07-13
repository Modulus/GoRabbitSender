package rabbit

import (
	"fmt"
	"log"

	"github.com/streadway/amqp"
)

func buildChannel(exchangeName string) (*amqp.Channel, error) {
	conn, err := amqp.Dial("amqp://thedude:opinion@localhost:5672/")
	if err != nil {
		return nil, err
	}
	amqpChan, err := conn.Channel()
	if err != nil {
		return nil, err
	}

	err = amqpChan.ExchangeDeclare(exchangeName,
		"fanout",
		true,
		false,
		false,
		false, nil)
	if err != nil {
		return nil, err
	}

	// Clear amqp channel if connection to server is lost
	amqpErrorChan := make(chan *amqp.Error)
	amqpChan.NotifyClose(amqpErrorChan)
	go func(ec chan *amqp.Error) {
		for msg := range ec {
			log.Fatalf("Channel Cleanup %s\n", msg)
		}
	}(amqpErrorChan)

	return amqpChan, err
}

func SendMessage(json, queueName, exchangeName, rabbitmqConnectionString string) {

	log.Printf("Received message %s\n", json)
	log.Printf("Connection to rabbitmq")

	connection, err := amqp.Dial(rabbitmqConnectionString)
	failOnError(err, "Failed to connect to RabbitMQ")
	defer connection.Close()

	channel, err := buildChannel(exchangeName)
	failOnError(err, "Failed to build channel")

	queue, err := channel.QueueDeclare(
		queueName,
		false,
		false,
		false,
		false,
		nil,
	)
	failOnError(err, "Failed to declare queue")

	log.Printf("Publishing message to rabbitmq")
	err = channel.Publish(
		exchangeName, //exchange
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
