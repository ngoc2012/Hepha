package main

import "build/Component"

type app struct {
	Component.Base
	Content appContent
}

type appContent struct {
	Component.Base
	Buton contentButon
}

type contentButon struct {
	Component.Base
	Number int
}

func main() {
	// Create a new instance of the app
	buton := contentButon{}
	buton.New("src/Buton.html")
	buton.Number = 1
	buton.File2String()

	content := appContent{}
	content.New("src/Content.html")
	content.Buton = buton
	content.Render2String(content)

	index := app{}
	index.New("src/layout.html")
	index.Content = content
	index.Render2File(index, "dist/index.html")
}
