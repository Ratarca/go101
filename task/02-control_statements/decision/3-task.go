package main

import (
	"fmt"
	"math/rand"
)

func main() {

	var random_age int
	random_age = rand.Intn(100)

	switch {
	case random_age < 10:
		fmt.Println(random_age, ": Children")
	case random_age >= 11 && random_age <= 16:
		fmt.Println(random_age, ": Teenager")
	case random_age >= 17 && random_age <= 55:
		fmt.Println(random_age, ": Adult")
	default:
		fmt.Println(random_age, ": Senior")
	}

}
