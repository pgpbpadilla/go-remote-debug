package main

import (
	"log"
	"net/http"
)

func main() {
	log.Println("starting server...")
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		log.Println("responding...")
		w.Write([]byte(`Hello world`))
	})
	log.Fatal(http.ListenAndServe(":8080", nil))
}
