package main

import (
	"fmt"
	"html/template"
	"os"
	"path/filepath"
)

type PageData struct {
	Title     string
	Items     []string
	Condition bool
}

func main() {
	// tmpl := template.Must(template.ParseFiles("index.html", "partial.html"))
	// tmpl := template.Must(template.ParseGlob("src/**/*.html"))

	// data := PageData{
	// 	Title:     "Main HTML",
	// 	Items:     []string{"Item 1", "Item 2", "Item 3"},
	// 	Condition: true,
	// }

	// // Create a file to write the output
	// file, err := os.Create("dist/index.html")
	// if err != nil {
	// 	fmt.Printf("Failed to create file: %v\n", err)
	// }
	// defer file.Close()

	// err = tmpl.Execute(file, data)
	// if err != nil {
	// 	fmt.Printf("Failed to execute template: %v\n", err)
	// }

	// fmt.Printf("Built done\n")

	// Parse all HTML files in the src/components directory as partial templates
	partialTmpl := template.Must(template.ParseGlob("src/components/*.html"))

	inputDir := "src/app"
	outputDir := "dist"

	// Traverse the src/app directory and its subdirectories to find HTML files
	filepath.Walk(inputDir, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}

		if !info.IsDir() && filepath.Ext(path) == ".html" {

			fmt.Printf("Processing %s\n", path)

			// Parse the current HTML file and include the partial templates
			tmpl := template.Must(partialTmpl.Clone())
			tmpl = template.Must(tmpl.ParseFiles(path))

			// Create the output directory if it doesn't exist
			if err := os.MkdirAll(outputDir, 0755); err != nil {
				fmt.Printf("Failed to create output directory: %v\n", err)
				return err
			}

			// Get the relative path from src/app
			relPath, err := filepath.Rel(inputDir, path)
			if err != nil {
				fmt.Printf("Failed to get relative path: %v\n", err)
				return err
			}

			// Create the output path in the dist directory
			outputPath := filepath.Join(outputDir, relPath)

			// Ensure the output directory structure exists
			if err := os.MkdirAll(filepath.Dir(outputPath), 0755); err != nil {
				fmt.Printf("Failed to create output directory structure: %v\n", err)
				return err
			}

			// Create a file to write the output
			// fmt.Printf("Output %s\n", outputPath)
			file, err := os.Create(outputPath)
			if err != nil {
				fmt.Printf("Failed to create file: %v\n", err)
				return err
			}
			defer file.Close()

			data := PageData{
				Title:     "Main HTML",
				Items:     []string{"Item 1", "Item 2", "Item 3"},
				Condition: true,
			}

			// Execute the template and write the output to the file
			err = tmpl.ExecuteTemplate(file, filepath.Base(path), data)
			if err != nil {
				fmt.Printf("Failed to execute template: %v\n", err)
				return err
			}

			fmt.Printf("Built %s\n", outputPath)
		}
		return nil
	})
}
