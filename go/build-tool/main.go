package main

import (
	"build/component"
	"build/file"
	"html/template"
)

type page struct {
	Html template.HTML
}

type examplesPage struct {
	Html examplesContent
}

type examplesContent struct {
	Html   template.HTML
	Button examplesButton
}

type examplesButton struct {
	Html   template.HTML
	Number int
}

func main() {
	file.Clean()

	main_index := page{
		Content: component.FileContent2Html("src/Content.html"),
	}
	component.Render2File("src/layout.html", "dist/index.html", main_index)

	button := examplesButton{Number: 1}
	button.Html = component.Render2Html("src/Button.html", button)

	content = pageContent{Button: button}
	content.Html = component.FileContent2Html("src/Content.html")

	examples_index := page{Content: content}
	component.Render2File("src/examples/layout.html", "dist/examples/index.html", examples_index)
}
