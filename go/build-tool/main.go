package main

import (
	"build/component"
	"html/template"
)

type app struct {
	Content appContent
}

type appContent struct {
	Html  template.HTML
	Buton contentButon
}

type contentButon struct {
	Html   template.HTML
	Number int
}

func main() {
	buton := contentButon{ Number: 1 }
	buton.Html = component.Render2Html("src/Buton.html", buton)

	content := appContent{ Buton: buton }
	content.Html = component.Render2Html("src/Content.html", content)

	index := app{ Content: content }
	component.Render2File("src/layout.html", "dist/index.html", index)
}
