package main

import "fmt"

// Simple Array : Has a fixed size and same type
var simple_array [5]int

// Array Methods

func main() {
	fmt.Println(simple_array)

	declared_array := [4]int{2, 3, 5, 7}
	fmt.Println(declared_array)

	simple_array[0] = 10
	fmt.Println(simple_array)

}
