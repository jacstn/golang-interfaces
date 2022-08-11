package main

import (
	"fmt"
	"io"
	"net/http"
)

type logWriter struct{}

func main() {
	resp, err := http.Get("https://google.com")

	if err != nil {
		fmt.Println("error", err)
	}
	lw := logWriter{}
	io.Copy(lw, resp.Body)
}

func (logWriter) Write(bs []byte) (n int, err error) {
	fmt.Println(string(bs))

	return len(bs), nil
}
