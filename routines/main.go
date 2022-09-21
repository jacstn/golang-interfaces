package main

import (
	"fmt"
	"sync"
)

func Printsth(s string, wg *sync.WaitGroup) {
	fmt.Println(s)
	wg.Done()
}

func main() {
	var wg sync.WaitGroup

	words := []string{"alpha", "beta", "gama"}

	for i, v := range words {
		wg.Add(1)
		go Printsth(fmt.Sprintf("%d, %s", i, v), &wg)
	}
	wg.Wait()
}
