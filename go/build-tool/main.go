package main

import (
	"build/component"
	"build/file"
	"html/template"
)

func main() {
	file.Clean()

	component.Render2File("src/layout.html", "dist/index.html", component.FileContent2Html("src/Content.html"))

	button := component.Render2Html("src/examples/Button.html", map[string]interface{}{
		"Number": 1,
	})
	content := component.Render2Html("src/examples/Content.html", map[string]interface{}{
		"Button": template.HTML(button),
	})

	component.Render2File("src/layout.html", "dist/examples/index.html", content)
}
