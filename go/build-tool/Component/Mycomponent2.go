package Component

import (
	"bytes"
	"html/template"
	"log"
)

type Mycomponent2 struct {
	Path string
	Html string
	Data string
}

func (t Mycomponent2) Init() {
	t.Path = "build-tool/Component/Mycomponent2.html"
	t.Data = "some data 2"
	tmpl := template.Must(template.ParseFiles(t.Path))

	var buf bytes.Buffer
	err := tmpl.Execute(&buf, t)
	if err != nil {
		log.Fatalf("Error executing template: %v", err)
	}

	t.Html = buf.String()
	//log.Println(result)
}
