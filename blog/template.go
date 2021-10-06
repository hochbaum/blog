package blog

import (
	"path/filepath"
	"text/template"
)

// Very ugly approach and only here for getting the first runnable build.
// This is without question going to be refactored.

var templates = make(map[string]*template.Template)

func init() {
	registerTemplate("index.html")

	// More templates can be added here...
}

// registerTemplate parses the specified template and registers it.
func registerTemplate(path string) {
	name := filepath.Base(path)
	templates[name] = template.Must(template.ParseFiles(path))
}
