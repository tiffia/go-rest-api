package main

import (
	"fmt"
	"net/http"
	"log"

	"github.com/gorilla/mux"
)

func getRequest(w http.ResponseWriter, r *http.Request){
	fmt.Fprintf(w, "This is a GET\n")
}

func postRequest(w http.ResponseWriter, r *http.Request){
	fmt.Fprintf(w, "This is a POST\n")
}

func deleteRequest(w http.ResponseWriter, r *http.Request){
	fmt.Fprintf(w, "This is a DELETE\n")
}

func main() {

	r := mux.NewRouter()
	r.HandleFunc("/", getRequest).Methods("GET")
	r.HandleFunc("/", postRequest).Methods("POST")
	r.HandleFunc("/", deleteRequest).Methods("DELETE")

	http.Handle("/", r)
	fmt.Println("Server started at localhost: 8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}