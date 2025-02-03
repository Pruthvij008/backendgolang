package views

import (
	"fmt"
	"html/template"
	"io/fs"
	"log"
	"net/http"
)

// Must panics if there's an error parsing a template
func Must(t Template, err error) Template {
	if err != nil {
		panic(err)
	}
	return t
}

// Template wraps an HTML template for rendering
type Template struct {
	htmlTPL *template.Template
}

// ParseFs parses templates from an embedded file system
func ParseFs(embedFS fs.FS, patterns ...string) (Template, error) {
	tpl, err := template.ParseFS(embedFS, patterns...)
	if err != nil {
		return Template{}, fmt.Errorf("parse template error: %w", err)
	}
	return Template{htmlTPL: tpl}, nil
}

// Parse parses templates from a file path
func Parse(filepath string) (Template, error) {
	tpl, err := template.ParseFiles(filepath)
	if err != nil {
		return Template{}, fmt.Errorf("parse template error: %w", err)
	}
	return Template{htmlTPL: tpl}, nil
}

// Execute renders the template with data
func (t Template) Execute(w http.ResponseWriter, data interface{}) {
	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	err := t.htmlTPL.Execute(w, data)
	if err != nil {
		log.Printf("Error executing the template: %v", err)
		http.Error(w, "Error executing template", http.StatusInternalServerError)
	}
}
