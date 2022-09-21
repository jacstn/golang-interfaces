package main

import (
	"io"
	"os"
	"strings"
	"sync"
	"testing"
)

func Test_Printsth(t *testing.T) {
	stdOut := os.Stdout
	r, w, _ := os.Pipe()
	os.Stdout = w

	var wg sync.WaitGroup
	wg.Add(1)
	go Printsth("hello", &wg)

	wg.Wait()

	w.Close()

	result, _ := io.ReadAll(r)

	output := string(result)
	os.Stdout = stdOut

	if !strings.Contains(output, "hello") {

		t.Errorf("error, no hello found %s", output)
	}

}
