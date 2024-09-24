package main

import (
	"build/component"
	"build/file"
	"html/template"
)

type app struct {
	Content appContent
}

type appContent struct {
	Html   template.HTML
	Button contentButton
}

type contentButton struct {
	Html   template.HTML
	Number int
}

func main() {
	file.Clean()
	button := contentButton{Number: 1}
	button.Html = component.Render2Html("src/Button.html", button)

	content := appContent{Button: button}
	content.Html = component.Render2Html("src/Content.html", content)

	index := app{Content: content}
	component.Render2File("src/layout.html", "dist/index.html", index)
}
