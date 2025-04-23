package main

import (
	"log"
	"net/http"
	"stocksapp/backend/internal/handler"
)

func main() {
	http.HandleFunc("/", handler.HelloHandler)
	log.Println("API corriendo en :8080")
	log.Fatal(http.ListenAndServe(":8080", nil))	
}
