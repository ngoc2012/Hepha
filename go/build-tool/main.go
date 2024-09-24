package main

import component "build/Component"

type app struct {
	component.Base
	Content appContent
}

type appContent struct {
	component.Base
	Buton contentButon
}

type contentButon struct {
	component.Base
	Number int
}

func main() {
	// Create a new instance of the app
	buton := contentButon{}
	buton.Input = "src/Buton.html"
	buton.Number = 1
	buton.Render2String(buton)
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
