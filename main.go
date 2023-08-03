package main

import (
	"fmt"
	"html"
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/ping", func(w http.ResponseWriter, r *http.Request) {
		fmt.Println("Host:", r.Host)
		fmt.Println("Path:", r.URL.Path)
		for k, v := range r.Header {
			fmt.Printf("%s: %s\n", k, v)
		}
		fmt.Fprintf(w, "Hello, %q", html.EscapeString(r.URL.Path))
	})

	log.Fatal(http.ListenAndServe(":8080", nil))
}
