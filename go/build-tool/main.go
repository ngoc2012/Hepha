package main

import (
	"build/component"
	"build/file"
	"html/template"
)

type page struct {
	Content pageContent
}

type pageContent struct {
	Html   template.HTML
	Button contentButton
}

type contentButton struct {
	Html   template.HTML
	Number int
}

func main() {
	file.Clean()

	index := page{Content: content}
	component.Render2File("src/layout.html", "dist/index.html", index)

	button := contentButton{Number: 1}
	button.Html = component.Render2Html("src/Button.html", button)

	content = pageContent{Button: button}
	content.Html = component.FileContent2Html("src/Content.html")

	content := pageContent{Html: button}
	examples_index := page{Content: content}
	component.Render2File("src/examples/layout.html", "dist/examples/index.html", examples_index)
}
