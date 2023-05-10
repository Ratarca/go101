package main

import "fmt"

func main() {
	// 8 Variables <4 explicit and 4 implicit>
	var name string = "Loop"
	var hair string = "Black"
	var max_age int = 20
	var gender string = "male"

	quantity_legs := 4
	tits := true
	fly := false
	exo_skelecton := false
	var swim bool = true

	fmt.Println(name, hair, max_age,
		gender, quantity_legs, tits,
		fly, exo_skelecton, swim)
}
