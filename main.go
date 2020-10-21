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
	_, err := fmt.Fprintf(res,"Hello, %s", path.Base(req.URL.Path))
	//_, err := res.Write([]byte("Hello World!"))
	if err != nil {
		log.Println(err)
	}
}
