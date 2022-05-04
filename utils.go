package main

import (
	"log"
	"net/http"
)

func GetHTML(url string) *http.Response {
	response, err := http.Get(url)
	if err != nil {
		log.Fatal(err)
	}

	return response
}
