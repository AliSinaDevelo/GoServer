package main

import (
	"fmt"
	"log"
	"net/http"
)

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