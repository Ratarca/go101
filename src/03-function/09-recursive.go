package main

import "fmt"

func find_number(position int) int {
	if position <= 1 {
		return position
	}

	return find_number(position-2) + find_number(position-1)
}

func main() {
	fmt.Println("Recursive")
	// 1 1 2 3 3 5 8 13
	position := 10
	fmt.Println(find_number(position))

}
