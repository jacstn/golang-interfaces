package main

import (
	"fmt"
	"net/http"
)

func main() {
	links := []string{
		"https://google.com",
		"http://faceboook.com",
		"http://golang.org",
		"https://pavoya.com",
	}
	c := make(chan string)

	for _, link := range links {
		go checkLink(link, c)
	}

	for ret := range c {
		fmt.Println(ret)
	}
}

func checkLink(link string, c chan string) {
	_, err := http.Get(link)
	if err != nil {
		c <- link + " is down"
		return
	}
	c <- link + " is up"
}
