package main

import (
	"log"
	"net/http"

	"bruh1/handlers"
)

func main() {
	http.HandleFunc("/user", handlers.UserHandler)

	log.Println("🚀 Server started on http://localhost:8080")
	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatal(err)
	}
}
