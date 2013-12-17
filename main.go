package main

import (
	"./views"
	"github.com/gorilla/mux"
	"log"
	"net/http"
)

var (
	ip   = ""
	port = "8080"
	addr = ip + ":" + port
)

func main() {
	router := mux.NewRouter()

	// all the view handlers
	router.HandleFunc("/", views.Index)

	// handle 404 errors
	router.NotFoundHandler = http.HandlerFunc(views.NotFoundHandler)

	
	// static dir, this should be disabled in production
	router.PathPrefix("/static/").Handler(http.FileServer(http.Dir(".")))


	http.Handle("/", router)
	
	// run the http server
	err := http.ListenAndServe(addr, nil)
	if err != nil {
		log.Fatal(err)
	}
}
