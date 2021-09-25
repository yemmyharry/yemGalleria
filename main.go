package main

import (
	"fmt"
	"github.com/gorilla/mux"
	"net/http"
)

func Home(w http.ResponseWriter, r *http.Request) {
	fmt.Println("page visited")
	w.Header().Set("Content-Type", "text/html")
	fmt.Fprint(w, `<h2> testing </h2>`)
}

func notFound(w http.ResponseWriter, r *http.Request)  {
	w.Header().Set("Content-Type", "text/html")
	w.WriteHeader(http.StatusNotFound)
	fmt.Fprint(w, "<h3>sorry but an error occured.</h3>")
}

func main()  {

	r := mux.NewRouter()

	r.NotFoundHandler = http.HandlerFunc(notFound)
	r.HandleFunc("/", Home)


	http.ListenAndServe(":8080", r)
}