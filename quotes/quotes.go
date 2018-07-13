package quotes

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"time"
)

type Response struct {
	Data    []string `json:string`
	Created string   `json:string`
}

func GetJSON(paragraphs int, quotes int) string {

	response := GetJSONRaw(paragraphs, quotes)

	bytes, _ := json.Marshal(response)
	log.Printf("===================================\n")
	log.Printf("%v\n", string(bytes))
	log.Printf("===================================\n")

	return string(string(bytes))
}

func GetJSONRaw(paragraphs, quotes int) Response {
	log.Println("Starting the application...")
	response, err := http.Get(fmt.Sprintf("http://loremricksum.com/api/?paragraphs=%d&quotes=%d", paragraphs, quotes))
	if err != nil {
		log.Fatalf("The HTTP request failed with error %s\n", err)
	}
	output, _ := ioutil.ReadAll(response.Body)
	log.Println(string(output))

	currentDate := time.Now().Format(time.UnixDate)

	log.Printf("Current date %s", currentDate)

	jsonResponse := Response{}
	json.Unmarshal(output, &jsonResponse)
	jsonResponse.Created = currentDate

	return jsonResponse
}
