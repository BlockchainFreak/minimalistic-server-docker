package main

import (
	"fmt"
	"net/http"
	"os"
)

func helloHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello, World!\n")
	fmt.Fprintf(w, "You accessed: %s\n", r.URL.Path)
	fmt.Fprintf(w, "Request method: %s\n", r.Method)
}

func main() {
	port := os.Getenv("PORT")
	if port == "" {
		port = "8000"
	}

	http.HandleFunc("/", helloHandler)
	fmt.Printf("Starting server at port %s\n", port)
	if err := http.ListenAndServe(":"+port, nil); err != nil {
		fmt.Println(err)
	}
}