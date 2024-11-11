package main

import (
	"build/cmp"
	"build/file"
	"fmt"
	"html/template"
)

type Map map[string]interface{}

var GZip = true

func react_page() {
	button := cmp.Render2Html("templates/examples/react/Button.html", Map{"Number": 1})

	var buttons [10]template.HTML
	for i := 0; i < 10; i++ {
		buttons[i] = button
	}

	content := cmp.Render2Html("templates/examples/react/Content.html", Map{"Buttons": buttons})

	cmp.Render("src/examples/react/index.html", "templates/layout.html", Map{
		"Content": content,
		"Style":   []string{"/main.css", "/fonts/inter.css"},
		"Script":  []string{"/utils.js", "button.js"},
	})
}

func remix_page() {
	cmp.Render("src/examples/remix/index.html", "templates/layout.html", Map{
		"Content": cmp.Read("templates/examples/remix/Content.html"),
		"Style":   []string{"/main.css", "/fonts/inter.css"},
		"Script":  []string{},
	})
}

func main() {
	err := file.Clean()
	if err != nil {
		fmt.Printf("failed to clean: %s\n", err)
		return
	}

	cmp.Render("src/index.html", "templates/layout.html", Map{
		"Content": cmp.Read("templates/Content.html"),
		"Style":   []string{"/main.css", "/fonts/inter.css"},
		"Script":  []string{},
	})
	cmp.Render("src/examples/index.html", "templates/layout.html", Map{
		"Content": cmp.Read("templates/examples/Content.html"),
		"Style":   []string{"/main.css", "/fonts/inter.css"},
		"Script":  []string{"/utils.js"},
	})
	cmp.Render("src/doc/index.html", "templates/doc/index.html", Map{})

	react_page()
	remix_page()
}
