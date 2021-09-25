package main

import (
	"fmt"
	"net/http"
)

func main()  {

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Println("page visited")
		w.Header().Set("Content-Type", "text/html")
		fmt.Fprint(w, `<h2> testing </h2>`)
	})
	http.ListenAndServe(":8080", nil)
}