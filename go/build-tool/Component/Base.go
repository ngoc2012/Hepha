package Component

import (
	"bytes"
	"html/template"
	"log"
	"os"
	"path/filepath"
	"reflect"
)

type Base struct {
	Input string
	Html  string
}

func (t *Base) New(input string) Base {
	return Base{
		Input: input,
		Html:  "",
	}
}

func (t *Base) GeComponentName(s interface{}) string {
	// Get the reflect.Type of s
	typ := reflect.TypeOf(s)

	// If s is a pointer, get the underlying type
	if typ.Kind() == reflect.Ptr {
		typ = typ.Elem()
	}

	// Get the name of the type
	return typ.Name()
}

func (t *Base) Render2String(s interface{}) {
	log.Println(t.Input)
	tmpl := template.Must(template.ParseFiles(t.Input))

	var buf bytes.Buffer
	err := tmpl.Execute(&buf, s)
	if err != nil {
		log.Fatalf("Error executing template: %v", err)
	}

	t.Html = buf.String()
}

func (t *Base) Render2File(s interface{}, outputPath string) {

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

	log.Println(outputPath)
	tmpl := template.Must(template.ParseFiles(outputPath))

	// Execute the template and write the output to the file
	err = tmpl.ExecuteTemplate(file, filepath.Base(outputPath), s)
	if err != nil {
		log.Fatalf("Failed to execute template: %v\n", err)
	}
}

func (t *Base) File2String() {
	// Read the file content into a byte slice
	byteValue, err := os.ReadFile(t.Input)
	if err != nil {
		log.Fatalf("Failed to read file: %v\n", err)
	} else {
		t.Html = string(byteValue)
	}
}
