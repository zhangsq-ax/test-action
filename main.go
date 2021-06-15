package main

import (
	"fmt"
	"log"
	"net/http"
	"path"
)

func main() {
	http.HandleFunc("/", index)

	log.Println("Start server at port 8080...")
	log.Fatal(http.ListenAndServe(":8080", nil))
}


func index(res http.ResponseWriter, req *http.Request) {
	name := path.Base(req.URL.Path)
	if name == "/" {
		name = "world"
	}
	_, err := fmt.Fprintf(res,"Hello, %s!", name)
	if err != nil {
		log.Println(err)
	}
}
