package cmp

import (
	"bytes"
	"html/template"
	"log"
	"os"
	"path/filepath"
)

func Render(inputPath string, s interface{}) template.HTML {
	log.Println("Render to string: ", inputPath)
	tmpl := template.Must(template.ParseFiles(inputPath))

	var buf bytes.Buffer
	err := tmpl.Execute(&buf, s)
	if err != nil {
		log.Fatalf("Error executing template: %v", err)
		return ""
	}

	return template.HTML(buf.String())
}

func Render2File(outputPath string, inputPath string, s interface{}) {
	log.Println("Render to file: ", inputPath, outputPath)
	if err := os.MkdirAll(filepath.Dir(outputPath), 0755); err != nil {
		log.Fatalf("Failed to create output directory structure: %v\n", err)
	}

	file, err := os.Create(outputPath)
	if err != nil {
		log.Fatalf("Failed to create file: %v\n", err)
	}
	defer file.Close()

	tmpl := template.Must(template.ParseFiles(inputPath))

	err = tmpl.Execute(file, s)
	if err != nil {
		log.Fatalf("Failed to execute template: %v\n", err)
	}
}

func Read(path string) template.HTML {
	byteValue, err := os.ReadFile(path)
	if err != nil {
		log.Fatalf("Component.Base.File2Html error: %v: '%s'\n", err, path)
		return ""
	} else {
		return template.HTML(string(byteValue))
	}
}
