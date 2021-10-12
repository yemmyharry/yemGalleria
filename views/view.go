package views

import (
	"html/template"
	"log"
)

type View struct {
	Template *template.Template
	Layout string
}

func NewView(layout string, files ...string) *View {
	files = append(files, "views/layouts/footer.gohtml", "views/layouts/bootstrap.gohtml")
	t, err := template.ParseFiles(files...)
	if err != nil {
		log.Fatal(err)
	}

	return &View{
		Template: t,
		Layout: layout,
	}
}