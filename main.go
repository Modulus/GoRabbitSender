package main

import (
	"log"

	"github.com/modulus/GoRabbitSender/sender"
	"github.com/modulus/goRabbitSender/quotes"
)

func main() {

	sender := sender.NewSender("config.json")

	json := quotes.GetJSON(1, 4)
	log.Printf(json)

	log.Printf("Sending message to rabbitmq")
	sender.SendMessage(json)

}
