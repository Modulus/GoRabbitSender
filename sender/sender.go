package sender

import (
	"encoding/json"
	"fmt"
	"log"
	"os"

	"github.com/streadway/amqp"
)

type Configuration struct {
	RabbitmqConnectionString      string
	RabbitmqExchangeName          string
	RabbitmqQueueName             string
	ElasticsearchConnectionString string
}

func createConfiguration(filePath string) Configuration {
	file, err := os.Open(filePath)

	if err != nil {
		log.Printf("Failed to open file at %s", filePath)
		file, _ = os.Open("config.json")

	}

	defer file.Close()

	decoder := json.NewDecoder(file)
	configuration := Configuration{}
	err = decoder.Decode(&configuration)
	if err != nil {
		fmt.Println("error:", err)
	}
	fmt.Println(configuration)

	return configuration
}

type Sender struct {
	Config Configuration
}

func NewSender(configFilePath string) *Sender {
	sender := new(Sender)
	sender.Config = createConfiguration(configFilePath)
	return sender
}

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

func (s Sender) SendMessage(json string) {

	log.Printf("Received message %s\n", json)
	log.Printf("Connection to rabbitmq")

	connection, err := amqp.Dial(s.Config.RabbitmqConnectionString)
	failOnError(err, "Failed to connect to RabbitMQ")
	defer connection.Close()

	channel, err := buildChannel(s.Config.RabbitmqExchangeName)
	failOnError(err, "Failed to build channel")

	queue, err := channel.QueueDeclare(
		s.Config.RabbitmqQueueName,
		false,
		false,
		false,
		false,
		nil,
	)
	failOnError(err, "Failed to declare queue")

	log.Printf("Publishing message to rabbitmq")
	err = channel.Publish(
		s.Config.RabbitmqExchangeName, //exchange
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
