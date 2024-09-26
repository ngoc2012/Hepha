package main

import (
	"build/component"
	"build/file"
	"html/template"
)

func examples_page() {
	button := component.Render2Html("src/examples/Button.html", map[string]interface{}{
		"Number": 1,
	})

	var buttons [10]template.HTML
	for i := 0; i < 10; i++ {
		buttons[i] = button
	}

	content := component.Render2Html("src/examples/Content.html", map[string]interface{}{
		"Buttons": buttons,
	})

	component.Render2File("src/layout.html", "dist/examples/index.html", content)
}

func main() {
	file.Clean()

	component.Render2File("dist/index.html", "src/layout.html", component.FileContent2Html("src/Content.html"))
	component.Render2File("src/layout.html", "dist/index.html", component.FileContent2Html("src/Content.html"))

	examples_page()
}
