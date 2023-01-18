package main

import "sync"

var s string

func ModifyString(ns string, wg *sync.WaitGroup) {
	defer wg.Done()
	s = ns
}

func main() {
	var wg sync.WaitGroup

	wg.Add(2)
	go ModifyString("jacek", &wg)
	go ModifyString("placek", &wg)
	wg.Wait()
}
