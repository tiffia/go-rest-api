package main

import (
	"fmt"
	"net/http"
	"log"
)

func helloWorld(w http.ResponseWriter, r *http.Request){
	fmt.Fprintf(w, "Hello, World!\n")
}

func main() {
	http.HandleFunc("/", helloWorld)
	fmt.Println("Starting server on localhost:8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}