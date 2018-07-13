package main

import (
	"encoding/json"
	"fmt"
	"log"
	"os"

	"github.com/modulus/goRabbitSender/quotes"
	"github.com/modulus/goRabbitSender/rabbit"
)

type Configuration struct {
	RabbitmqConnectionString      string
	RabbitmqExchangeName          string
	RabbitmqQueueName             string
	ElasticsearchConnectionString string
}

func createConfiguration() Configuration {

	file, _ := os.Open("config.json")

	defer file.Close()

	decoder := json.NewDecoder(file)
	configuration := Configuration{}
	err := decoder.Decode(&configuration)
	if err != nil {
		fmt.Println("error:", err)
	}
	fmt.Println(configuration)

	return configuration
}

func main() {
	config := createConfiguration()
	fmt.Println(config)

	json := quotes.CreateJSON(1, 4)
	log.Printf(json)

	log.Printf("Sending message to rabbitmq")
	rabbit.SendMessage(json, config.RabbitmqQueueName, config.RabbitmqExchangeName, config.RabbitmqConnectionString)

}
