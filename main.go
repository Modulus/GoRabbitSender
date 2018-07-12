package main

import (
	"context"
	"log"

	"github.com/olivere/elastic"
)

type Message struct {
	Message string `json:"message"`
}

func main() {
	message := CreateQuote()
	log.Printf("Message: %s", message)

	ctx := context.Background()
	client, err := elastic.NewClient()
	if err != nil {
		// Handle error
	}
	message := Message{Message: "Take Five"}
	_, err = client.Index().
		Index("message").
		Type("doc").
		//Id("1").
		BodyJson(message).
		Refresh("wait_for").
		Do(ctx)
	if err != nil {
		// Handle error
		panic(err)
	}

}
