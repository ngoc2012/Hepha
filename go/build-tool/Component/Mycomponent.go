package Component

import (
	"bytes"
	"html/template"
	"log"
)

type Mycomponent struct {
	Path string
	Html string
	Data string
}

func (t *Mycomponent) Init() {
	t.Path = "build-tool/Component/Mycomponent.html"
	t.Data = "some data"
	tmpl := template.Must(template.ParseFiles(t.Path))

	var buf bytes.Buffer
	err := tmpl.Execute(&buf, t)
	if err != nil {
		log.Fatalf("Error executing template: %v", err)
	}

	t.Html = buf.String()
	log.Println(t.Html)
}
