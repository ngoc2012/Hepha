package main

import (
	"build/component"
	"build/file"
	"html/template"
)

type mainPage struct {
	Content template.HTML
}

type examplesPage struct {
	Content examplesContent
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

	main_index := page{Content: component.FileContent2Html("src/Content.html")}
	component.Render2File("src/layout.html", "dist/index.html", main_index)

	button := contentButton{Number: 1}
	button.Html = component.Render2Html("src/Button.html", button)

	content = pageContent{Button: button}
	content.Html = component.FileContent2Html("src/Content.html")

	examples_index := page{Content: content}
	component.Render2File("src/examples/layout.html", "dist/examples/index.html", examples_index)
}
