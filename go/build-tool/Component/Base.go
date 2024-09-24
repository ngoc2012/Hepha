package component

import (
	"bytes"
	"html/template"
	"log"
	"os"
	"path/filepath"
)

// type Base struct {
// 	Input string
// 	Html  template.HTML
// }

// func (t *Base) GeComponentName(s interface{}) string {
// 	// Get the reflect.Type of s
// 	typ := reflect.TypeOf(s)

// 	// If s is a pointer, get the underlying type
// 	if typ.Kind() == reflect.Ptr {
// 		typ = typ.Elem()
// 	}

// 	// Get the name of the type
// 	return typ.Name()
// }

func Render2String(s interface{}) {
	log.Println("Render to string", t.Input)
	tmpl := template.Must(template.ParseFiles(t.Input))

	var buf bytes.Buffer
	err := tmpl.Execute(&buf, s)
	if err != nil {
		log.Fatalf("Error executing template: %v", err)
	}

	t.Html = template.HTML(buf.String())
}

func Render2File(inputPath string, outputPath string, s interface{}) {
	log.Println("Render to file", inputPath, outputPath)
	// Ensure the output directory structure exists
	if err := os.MkdirAll(filepath.Dir(outputPath), 0755); err != nil {
		log.Fatalf("Failed to create output directory structure: %v\n", err)
	}

	// Create a file to write the output
	file, err := os.Create(outputPath)
	if err != nil {
		log.Fatalf("Failed to create file: %v\n", err)
	}
	defer file.Close()

	tmpl := template.Must(template.ParseFiles(inputPath))

	// Execute the template and write the output to the file
	err = tmpl.Execute(file, s)
	if err != nil {
		log.Fatalf("Failed to execute template: %v\n", err)
	}
}

func FileContent2String(path string) {
	byteValue, err := os.ReadFile(path)
	if err != nil {
		log.Fatalf("Component.Base.File2String error: %v: '%s'\n", err, path)
	} else {
		t.Html = template.HTML(string(byteValue))
	}
}
