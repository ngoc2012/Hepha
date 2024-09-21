package Component

import (
	"bytes"
	"html/template"
	"log"
)

type MySubComponent struct {
	Path string
	Html string
	Data string
}

func (t *MySubComponent) Init() {
	t.Path = "MySubComponent.html"
	t.Data = "another data"
	tmpl := template.Must(template.ParseFiles(t.Path))

	var buf bytes.Buffer
	err := tmpl.Execute(&buf, t)
	if err != nil {
		log.Fatalf("Error executing template: %v", err)
	}

	t.Html = buf.String()
	// log.Println(result)
}
