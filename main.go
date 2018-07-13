package main

import (
	"log"

	"github.com/modulus/goRabbitSender/quotes"
	"github.com/modulus/goRabbitSender/rabbit"
)

func main() {

	json := quotes.CreateJson()
	log.Printf(json)

	log.Printf("Sending message to rabbitmq")
	rabbit.SendMessage(json, "messageExchange")
}
