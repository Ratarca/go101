package main

import "fmt"

func main() {
	// Explicit declaration
	var explicit_declare string
	fmt.Println(explicit_declare)
	explicit_declare = "Some name"
	fmt.Println(explicit_declare)

	// Implicity declaration
	implicity_declare := 3.14159
	fmt.Println(implicity_declare)
}
