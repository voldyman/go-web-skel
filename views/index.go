package views

import (
	"html/template"
	"net/http"
)

var (
	templateDir = "templates"
)

func Index(w http.ResponseWriter, r *http.Request) {
	templateFile, err := getTemplate("index.html")
	if err != nil {
		Error(err.Error(), w, r)
		return
	}

	templ, err := template.ParseFiles(templateFile)
	if err != nil {
		w.Write([]byte(err.Error()))
		return
	}
	templ.Execute(w, struct { Name string} {"doodle"})
	logRequest(r.Method, r.URL, "Index")
}
