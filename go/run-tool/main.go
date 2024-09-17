package main

import (
	"log"
	"net/http"
)

type PageData struct {
	Title     string
	Items     []string
	Condition bool
}

func main() {
	http.Handle("/", http.FileServer(http.Dir(".")))

	log.Println("Server is running on http://localhost:8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
