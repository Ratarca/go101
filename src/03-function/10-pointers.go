package main

import "fmt"

// function with pointers

func inversion_signal(number int) int {
	return number * -1
}

func inversion_signal_pointer(number *int) {
	*number = *number * -1
}

func main() {
	my_number := 20

	// first case
	fmt.Println("First case")

	number_inversion := inversion_signal(my_number)
	fmt.Println(my_number, number_inversion)

	// second case
	fmt.Println("Second case")

	inversion_signal_pointer(&my_number)
	fmt.Println(my_number)

}
