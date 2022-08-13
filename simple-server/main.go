package main

import (
	"fmt"
	"net/http"
)

const portNumber = ":3000"

func Home(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "This is about page")
}

func About(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "This is about page")
}

func main() {
	http.HandleFunc("/", Home)
	http.HandleFunc("/", About)

	fmt.Println(fmt.Sprintf("Starting application %s", portNumber))
	http.ListenAndServe(portNumber, nil)
}
