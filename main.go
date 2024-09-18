package main

import (
	"fmt"
	"github.com/gorilla/mux"
	"log"
	"net/http"
)

func main() {
	r := mux.NewRouter()
	r.HandleFunc("/hello", helloHandler).Methods("GET")
	r.HandleFunc("/hello", helloHandler).Methods("POST")

	fmt.Println("Starting server at :8080")
	log.Fatal(http.ListenAndServe(":8080", r))
}

func helloHandler(w http.ResponseWriter, _ *http.Request) {
	fmt.Println(w, "Hello, World!")
}
