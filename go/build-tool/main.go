package main

import (
	"build/component"
	"build/file"
)

func main() {
	file.Clean()

	component.Render2File("src/layout.html", "dist/index.html", component.FileContent2Html("src/Content.html"))

	button := component.Render2Html("src/examples/Button.html", 1)

	content := component.Render2Html("src/examples/Content.html", map[string]interface{}{
		"Button": button,
	})

	component.Render2File("src/layout.html", "dist/examples/index.html", map[string]interface{}{
		"Content": content,
	})
}
