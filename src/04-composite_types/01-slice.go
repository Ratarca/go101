package main

import "fmt"

// Simple slice

// Methods : append, remove, get partition
func main() {
	my_slice := []int{}
	fmt.Println(my_slice)

	my_slice = append(my_slice, 10)
	my_slice = append(my_slice, 10, 20, 30)
	fmt.Println(my_slice)

	//Danger

}
