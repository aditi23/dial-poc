package main

import (
	"log"
	"net/http"

	"github.com/aditi23/dial-poc/handlers"
)

func main() {
	log.Println("Server started on: http://localhost:8080")
	http.HandleFunc("/tcp4", handlers.TCP4)
	http.HandleFunc("/tcp", handlers.TCP)
	http.ListenAndServe(":8080", nil)
}
