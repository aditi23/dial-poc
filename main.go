package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/aditi23/dial-poc/handlers"
)

func init() {
	os.Setenv("GODEBUG", "netdns=1")
}

func main() {
	fmt.Println("GODEBUG", os.Getenv("GODEBUG"))
	log.Println("Server started on: http://localhost:8998")
	http.HandleFunc("/tcp4", handlers.TCP4)
	http.HandleFunc("/tcp", handlers.TCP)
	http.ListenAndServe(":8998", nil)
}
