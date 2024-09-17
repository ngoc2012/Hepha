package main

import (
	"fmt"
	"html/template"
	"os"
)

type PageData struct {
	Title     string
	Items     []string
	Condition bool
}

func main() {
	// tmpl := template.Must(template.ParseFiles("index.html", "partial.html"))
	tmpl := template.Must(template.ParseGlob("src/**/*.html"))

	data := PageData{
		Title:     "Main HTML",
		Items:     []string{"Item 1", "Item 2", "Item 3"},
		Condition: true,
	}

	// Create a file to write the output
	file, err := os.Create("dist/index.html")
	if err != nil {
		fmt.Printf("Failed to create file: %v\n", err)
	}
	defer file.Close()

	err = tmpl.Execute(file, data)
	if err != nil {
		fmt.Printf("Failed to execute template: %v\n", err)
	}

	fmt.Printf("Template output written\n")
}
