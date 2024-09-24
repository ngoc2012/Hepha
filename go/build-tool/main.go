package main

import (
	"build/component"
	"html/template"
)

type app struct {
	Html   template.HTML
	Content appContent
}

type appContent struct {
	Html   template.HTML
	Buton contentButon
}

type contentButon struct {
	Html   template.HTML
	Number int
}

func main() {
	// Create a new instance of the app
	buton := contentButon{}
	buton.Number = 1
	buton.Html = component.Render2String("src/Buton.html", buton)
	print(buton.Html)

	content := appContent{}
	content.Buton = buton
	content.Html = component.Render2String(content)

	print(content.Html)
	index := app{}
	index.Content = content
	index.Html = component.Render2File("src/layout.html", "dist/index.html", index)
}
