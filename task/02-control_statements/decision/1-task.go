package main

import "fmt"

func main() {
	var age int = 10

	if age > 60 {
		fmt.Println("Pass")
	} else if age > 35 {
		fmt.Println("Pass")
	} else {
		fmt.Println("A")
	}
}
