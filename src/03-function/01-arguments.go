package main

import "fmt"

// variadic FUnctions : Functions with N arguments : Pack is a
func somar(numbers ...int) {
	fmt.Println(numbers)

	total := 0
	for _, number := range numbers {
		total += number
	}

	return total
}

func main() {
	somar(1, 2, 3, 4, 5, 6, 7, 8, 10)
}
