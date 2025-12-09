package backend

import (
	"fmt"
	"log"
	"net/http"
)

func helloWorld(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello, World!")
}

func Run(addr string) {
	http.HandleFunc("/", helloWorld)
	fmt.Println("Starting server on: ", addr)
	log.Fatal(http.ListenAndServe(addr, nil))
}