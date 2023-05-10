package main

import (
	"fmt"
	"math/rand"
)

func main() {
	some_number := rand.Intn(1000)

	if some_number%2 == 0 {
		fmt.Println(some_number, "is Odd")
	} else {
		fmt.Println(some_number, "isnt Odd")
	}
}
