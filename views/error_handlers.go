package views

import (
	"net/http"
)

func Error(error string, w http.ResponseWriter, r *http.Request) {
	ErrorHandler(w, r)
}

func NotFoundHandler(w http.ResponseWriter, r *http.Request) {
	http.Error(w, "Page Not Found", 404)
	logRequest(r.Method, r.Host, "NotFound Handler")
}

func ErrorHandler(w http.ResponseWriter, r *http.Request) {
	// todo: redirect(?) to error page

	http.Error(w, "Internal Server Error", 500)
	logRequest(r.Method, r.Host, "ErrorHandler")
}
