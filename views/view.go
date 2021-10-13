package views

import (
	"html/template"
	"log"
	"path/filepath"
)

type View struct {
	Template *template.Template
	Layout string
}

func NewView(layout string, files ...string) *View {
	files = append(files, layoutFiles()...)
	t, err := template.ParseFiles(files...)
	if err != nil {
		log.Fatal(err)
	}

	return &View{
		Template: t,
		Layout: layout,
	}
}

func layoutFiles() []string {
	files, err := filepath.Glob("views/layouts/*.gohtml")
	if err != nil {
		panic(err)
	}
	return files
}