package main

import "fmt"

func main() {
	var doTask int = 10
	counter := 0

	// For
	fmt.Println("Task counter")
	for counter < doTask {
		fmt.Println(counter, doTask)
		counter++
	}

	// For with condition
	fmt.Println("Adult life counter")
	var adult int = 21
	for i := 0; i < adult; i++ {
		fmt.Println(i)
	}
}
