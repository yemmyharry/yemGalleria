package main

import (
	"fmt"
	"net/http"
)

func main()  {

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprint(w, `<h1> testing </h1>`)
	})
	http.ListenAndServe(":8080", nil)
}