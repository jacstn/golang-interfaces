package main

import (
	"fmt"
	"net/http"
)

const portNumber = ":3000"

func main() {
	http.HandleFunc("/", Home)
	http.HandleFunc("/about", About)

	fmt.Println(fmt.Sprintf("Starting application %s", portNumber))
	http.ListenAndServe(portNumber, nil)
}
