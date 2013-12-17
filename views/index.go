package views

import (
	"html/template"
	"net/http"
)

var (
	templateDir = "templates"
)

func Index(w http.ResponseWriter, r *http.Request) {
	templateFile, err := getTemplate("test.html")
	if err != nil {
		Error(err.Error(), w, r)
		return
	}

	templ, err := template.ParseFiles(templateFile)
	if err != nil {
		w.Write([]byte(err.Error()))
		return
	}
	templ.Execute(w, "Index?")
	logRequest(r.Method, r.Host, "Index")
}
