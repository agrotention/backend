package main

import (
	"log"
	"net/http"
)

func main() {
	mux := http.NewServeMux()

	err := http.ListenAndServe(":8000", mux)
	if err != nil {
		log.Fatalln(err)
	}
}
