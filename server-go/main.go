package main

import (
	"log"
	"net/http"
	"os"
)

type PageData struct {
	Title     string
	Items     []string
	Condition bool
}

func main() {
	http.Handle("/", http.FileServer(http.Dir("./dist")))

	log.Println("Server is running on http://localhost:" + os.Getenv("GO_PORT"))
	log.Fatal(http.ListenAndServe(":"+os.Getenv("GO_PORT"), nil))
}
