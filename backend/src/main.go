package main

import (
	"fmt"
	"net/http"
)

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		html := "<html><body><h1>Hello, World!</h1></body></html>"
		fmt.Fprintf(w, html)
	})

	http.ListenAndServe(":8080", nil)
}
