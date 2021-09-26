package main

import (
	"calhounProject/views"
	"fmt"
	"github.com/gorilla/mux"
	"log"
	"net/http"
)

var (
	contactView *views.View
	homeView *views.View
)

func Home(w http.ResponseWriter, r *http.Request) {
	fmt.Println("page visited")
	w.Header().Set("Content-Type", "text/html")
	err := homeView.Template.Execute(w, nil)
	if err != nil {
		log.Fatal(err)
	}
}

func Contact(w	http.ResponseWriter, r *http.Request)  {
	w.Header().Set("Content-Type", "text/html")
	if err := contactView.Template.Execute(w,nil); err != nil {
		log.Fatal(err)
	}
}

func notFound(w http.ResponseWriter, r *http.Request)  {
	w.Header().Set("Content-Type", "text/html")
	w.WriteHeader(http.StatusNotFound)
	fmt.Fprint(w, "<h3>sorry but an error occured.</h3>")
}

func main()  {
	homeView = views.NewView("views/home.gohtml")
	contactView = views.NewView("views/contact.gohtml")

	r := mux.NewRouter()
	r.NotFoundHandler = http.HandlerFunc(notFound)
	r.HandleFunc("/", Home)
	r.HandleFunc("/contact", Contact)
	http.ListenAndServe(":8080", r)
}