package main

import (
	"fmt"
	"sync"
)

func main() {
	//fmt.Println("Hello, World!")
	//go printMessage("Hi there.. 1") // this is a seperate goroutine

	//Use wait group
	var wg sync.WaitGroup

	words := []string{
		"Red",
		"Pink",
		"Green",
		"Yellow",
		"White",
		"Blue",
	}

	wg.Add(6)

	for i, x := range words {
		go printMessage(fmt.Sprintf("%d: %s", i, x), &wg)
	}

	wg.Wait()

	//time.Sleep(1 * time.Second) // sleep for 1 second - NOT a good practice. Use wait group instead

	printMessage("Hi there.. 2", &wg)
}

func printMessage(s string, wg *sync.WaitGroup) {
	defer wg.Done()
	fmt.Println(s)
}

/*
To wait for multiple goroutines to finish, we can use a wait group.
*/
