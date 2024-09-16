package main

import (
	"html/template"
	"log"
	"net/http"
)

type PageData struct {
	Title   string
	Items   []string
	Condition bool
}

func main() {
	tmpl := template.Must(template.ParseFiles("index.html", "partial.html"))

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		data := PageData{
			Title:     "Main HTML",
			Items:     []string{"Item 1", "Item 2", "Item 3"},
			Condition: true,
		}
		err := tmpl.Execute(w, data)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}
	})

	log.Println("Server is running on http://localhost:8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}

