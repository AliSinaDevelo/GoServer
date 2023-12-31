package main

import (
	"fmt"
	"log"
	"net/http"
)

func form_handler(w http.ResponseWriter, r *http.Request) {
	if err := r.ParseForm(); err != nil {
		fmt.Fprintf(w, "Parse form error: %v", err)
		return
	}

	fmt.Fprintf(w, "POST request succesful")
	name := r.FormValue("name")
	address := r.FormValue("address")
	fmt.Fprintf(w, "Name: %s\n", name)
	fmt.Fprintf(w, "Address: %s\n", address)
}

func hello_handler(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/hello" {
		http.Error(w, "404 NOT FOUND!", http.StatusNotFound)
		return
	}
	if r.Method != "GET" {
		http.Error(w, "method not supported!", http.StatusNotFound)
		return
	}

	fmt.Fprintf(w, "Howdy!")
}

func main() {
	// check on the static directory
	file_server := http.FileServer(http.Dir("./static"))
	http.Handle("/", file_server)
	http.HandleFunc("/form", form_handler) 
	http.HandleFunc("/hello", hello_handler)

	fmt.Printf("Server starting at port 8080\n")
	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatal(err)
	}
}