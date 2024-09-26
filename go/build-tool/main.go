package main

import (
	"build/cmp"
	"build/file"
	"html/template"
)

type Map map[string]interface{}

func react_page() {
	button := cmp.Render("src/examples/react/Button.html", Map{"Number": 1})

	var buttons [10]template.HTML
	for i := 0; i < 10; i++ {
		buttons[i] = button
	}

	content := cmp.Render("src/examples/react/Content.html", Map{"Buttons": buttons})

	cmp.Render2File("dist/examples/react/index.html", "src/layout.html", Map{"Content": content})
}

func remix_page() {
	button := cmp.Render("src/examples/react/Button.html", Map{"Number": 1})

	var buttons [10]template.HTML
	for i := 0; i < 10; i++ {
		buttons[i] = button
	}

	content := cmp.Render("src/examples/react/Content.html", Map{"Buttons": buttons})

	cmp.Render2File("dist/examples/react/index.html", "src/layout.html", Map{"Content": content})
}

func main() {
	file.Clean()

	cmp.Render2File("dist/index.html", "src/layout.html", Map{
		"Content": cmp.Read("src/Content.html"),
		"Style":   []string{"/main.css", "/fonts/inter.css"},
	})
	cmp.Render2File("dist/examples/index.html", "src/layout.html", Map{"Content": cmp.Read("src/examples/Content.html")})

	react_page()
	remix_page()
}
