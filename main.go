package main

import (
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/", index)

	log.Println("Start server...")
	log.Fatal(http.ListenAndServe(":8080", nil))
}

func index(req http.ResponseWriter, res *http.Request) {
	_, err := req.Write([]byte("Hello"))
	if err != nil {
		log.Println(err)
	}
}
