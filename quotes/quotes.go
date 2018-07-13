package quotes

import (
	"io/ioutil"
	"log"
	"net/http"
)

type Data struct {
	data []string
}

func CreateJson() string {
	log.Println("Starting the application...")
	response, err := http.Get("http://loremricksum.com/api/?paragraphs=1&quotes=4")
	if err != nil {
		log.Fatalf("The HTTP request failed with error %s\n", err)
	}
	output, _ := ioutil.ReadAll(response.Body)
	log.Println(string(output))
	return string(output)
}
