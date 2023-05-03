package main

import "fmt"

// Stuct Type
type custom_user struct {
	name   string
	age    int
	wallet float64
}

func main() {
	fmt.Println("Structure file")

	// Traditional method
	ratarca := custom_user{"ratarca", 27, 1300.2}
	fmt.Println(ratarca)

	// Parcial
	ada := custom_user{name: "ratarca"}
	fmt.Println(ada)

	// default 0
	var empty_user custom_user
	fmt.Println(empty_user)

	// fill the empty files
	empty_user.age = 19
	fmt.Println(empty_user)
}
