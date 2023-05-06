package main

import "fmt"

func main() {
	fmt.Println("Pointers")

	// Declare variables
	var1 := 10
	var2 := var1

	fmt.Println(var1, var2)

	// Pointer == address memory
	var value int
	var valuePtr *int

	value = 100
	valuePtr = &value

	fmt.Println(value, valuePtr, *valuePtr)
}
