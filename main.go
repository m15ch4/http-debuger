package main

import (
	"fmt"
	"log"
	"net/http"
)

func handler(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Request received.")

	// Log all HTTP headers
	fmt.Println("HTTP Headers:")
	for header, values := range r.Header {
		for _, value := range values {
			fmt.Printf("%s: %s\n", header, value)
		}
	}
	fmt.Printf("Host: %s\n", r.Host)

	// Respond to the request
	fmt.Fprintf(w, "Hello, this is the HTTPS server!\n")
}

func main() {
	http.HandleFunc("/", handler)

	port := ":443" // HTTPS default port

	fmt.Printf("Server is running on https://localhost%s\n", port)

	err := http.ListenAndServeTLS(port, "certs/certificate.pem", "certs/key.pem", nil)
	if err != nil {
		log.Fatal("ListenAndServeTLS: ", err)
	}
}
