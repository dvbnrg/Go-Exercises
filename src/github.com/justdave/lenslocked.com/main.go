package main

import (
	"fmt"
	"html/template"
	"net/http"

	"github.com/gorilla/mux"
)

var homeTemplate *template.Template    //This is a frowned upon global variable, TODO: refactor dis
var contactTemplate *template.Template //This is a frowned upon global variable, TODO: refactor dis

func home(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html")
	if err := homeTemplate.Execute(w, nil); err != nil {
		panic(err)
	}
}

func contact(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html")
	if err := contactTemplate.Execute(w, nil); err != nil {
		panic(err)
	}
}

func faq(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html")
	fmt.Fprint(w, "<h1>Here are some Frequently Asked Questions:</h1>")
}

func hiross(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html")
	fmt.Fprint(w, "<h1>Hi Ross!</h1>")
}

func main() {

	var err error
	homeTemplate, err = template.ParseFiles(
		"views/home.gohtml",
		"views/layouts/footer.gohtml")
	if err != nil {
		panic(err)
	}

	contactTemplate, err = template.ParseFiles(
		"views/contact.gohtml",
		"views/layouts/footer.gohtml")
	if err != nil {
		panic(err)
	}

	var h http.Handler = http.HandlerFunc(home)
	r := mux.NewRouter()
	// This will assign the home page to the
	// NotFoundHandler
	r.NotFoundHandler = h
	r.HandleFunc("/faq", faq)
	r.HandleFunc("/contact", contact)
	r.HandleFunc("/hiross", hiross)
	http.ListenAndServe(":3000", r)
}
