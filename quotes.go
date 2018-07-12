package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

type Data struct {
	data []string `json:"data"`
}

func CreateQuote() string {
	restClient := &http.Client{}
	url := fmt.Sprintf("http://loremricksum.com/api/?paragraphs=%d&quotes=%d", 1, 4)
	request, err := http.NewRequest("GET", url, nil)
	if err != nil {
		// Handle error
	}

	resp, err := restClient.Do(request)

	defer resp.Body.Close()

	log.Printf("Response: %s", resp.Body)
	var data Data
	if err := json.NewDecoder(resp.Body).Decode(&data); err != nil {
		log.Println(err)
	} else {
		log.Println(data)
	}

	return data.data[0]
}
