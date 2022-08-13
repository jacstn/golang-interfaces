package main

import (
	"fmt"
	"net/http"

	"github.com/jacstn/golang-playground/simple-server/pkg/handlers"
)

const portNumber = ":3000"

func main() {
	http.HandleFunc("/", handlers.Home)
	http.HandleFunc("/about", handlers.About)

	fmt.Println(fmt.Sprintf("Starting application %s", portNumber))
	http.ListenAndServe(portNumber, nil)
}
