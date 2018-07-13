package rabbit

import (
	"fmt"
	"log"

	"github.com/streadway/amqp"
)

func () buildChannel() (*amqp.Channel, error) {
	conn, err := amqp.Dial(hook.connString)
	if err != nil {
		return nil, err
	}
	amqpChan, err := conn.Channel()
	if err != nil {
		return nil, err
	}

	err = amqpChan.ExchangeDeclare(hook.exchangeName, "fanout", true, hook.AutoDeleteExchange, hook.InternalExchange, hook.NowaitExchange, nil)
	if err != nil {
		return nil, err
	}

	// Clear amqp channel if connection to server is lost
	amqpErrorChan := make(chan *amqp.Error)
	amqpChan.NotifyClose(amqpErrorChan)
	go func(h *AmqpHook, ec chan *amqp.Error) {
		for msg := range ec {
			logrus.Errorf("AmqpHook.buildChannel> Channel Cleanup %s\n", msg)
			h.amqpChan = nil
		}
	}(hook, amqpErrorChan)
	return amqpChan, err
}

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
