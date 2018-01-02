package main

import (
	"fmt"
	"net/http"

	"github.com/justdave/lenslocked.com/views"

	"github.com/gorilla/mux"
)

var homeView *views.View
var contactView *views.View

func main() {

	homeView = views.NewView("bootstrap",
		"views/home.gohtml")
	contactView = views.NewView("bootstrap",
		"views/contact.gohtml")

	var h http.Handler = http.HandlerFunc(home)
	r := mux.NewRouter()
	// This will assign the home page to the
	// NotFoundHandler
	r.NotFoundHandler = h
	r.HandleFunc("/faq", faq)
	r.HandleFunc("/", home)
	r.HandleFunc("/contact", contact)
	r.HandleFunc("/hiross", hiross)
	http.ListenAndServe(":3000", r)
}

func home(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html")
	err := homeView.Template.ExecuteTemplate(w,
		homeView.Layout, nil)
	if err != nil {
		panic(err)
	}

}

func contact(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html")
	err := contactView.Template.ExecuteTemplate(w,
		contactView.Layout, nil)
	if err != nil {
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
