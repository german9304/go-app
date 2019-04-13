package main

import (
	"fmt"
	"log"
	"net/http"
)

func helloFunc(w http.ResponseWriter, r *http.Request) {
	_, err := fmt.Fprintf(w, "hello %v ", "worlds")
	if err != nil {
		log.Fatal(err)
	}
}

func main() {

	http.HandleFunc("/", helloFunc)
	err := http.ListenAndServe(":8081", nil)
	log.Fatal(err)
}
