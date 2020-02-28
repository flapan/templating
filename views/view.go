package views

import (
	"net/http"
	"path/filepath"
	"text/template"
)

// LayoutDir has the path to Layout files
var LayoutDir string = "views/layouts"

// NewView creates and prepares a view
func NewView(layout string, files ...string) *View {
	files = append(layoutFiles(), files...)
	t, err := template.ParseFiles(files...)
	if err != nil {
		panic(err)
	}

	return &View{
		Template: t,
		Layout:   layout,
	}
}

// The View Type descibes a view
type View struct {
	Template *template.Template
	Layout   string
}

// Render executes the tmeplate
func (v *View) Render(w http.ResponseWriter, data interface{}) error {
	return v.Template.ExecuteTemplate(w, v.Layout, data)
}

func layoutFiles() []string {
	files, err := filepath.Glob(LayoutDir + "/*.html")
	if err != nil {
		panic(err)
	}
	return files
}
