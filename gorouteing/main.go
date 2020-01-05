package main

import (
	"fmt"
	"sync"
)

func main() {
	var wg sync.WaitGroup
	wg.Add(1)
	message := make(chan string)
	fmt.Println("start")
	go func() {
		defer wg.Done()
		message <- "hello world"
	}()
	wg.Add(1)
	go func() {
		defer wg.Done()

		fmt.Println(<-message)
	}()
	wg.Wait()
	fmt.Println("end")
}
