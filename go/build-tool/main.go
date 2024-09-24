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
	buton.Html = Render2String("src/Buton.html", buton)
	print(buton.Html)

	content := appContent{}
	content.Input = "src/Content.html"
	content.Buton = buton
	content.Render2String(content)

	print(content.Html)
	index := app{}
	index.Input = "src/layout.html"
	index.Content = content
	index.Render2File(index, "dist/index.html")
}
