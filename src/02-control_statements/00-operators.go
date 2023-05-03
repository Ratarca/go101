package main

import (
	"fmt"
)

func main() {
	/*
			Arithmetic
		{sum:+, minus:- , multi:* , division: /, modules: % }
	*/
	fmt.Println("Arithmetic Operators")
	var a int16 = 8
	var b int16 = 4
	calculus := ((a + b) * (a - b)) / b

	fmt.Println(calculus)

	/*
			Relational
		> ; >= ; < ; <= ; == ; !=
	*/
	fmt.Println("Relational Operators")
	fmt.Println(a > b, b >= (a/2))
	fmt.Println(a < b, b <= (a/2))
	fmt.Println(a == b, b != (a/2))

	/*
			Logical
		AND {&&}; OR {||}; NOT {!}
	*/
	fmt.Println("Logical Operators")
	first_condition := (a > b) && (b > a)
	var second_condition bool = (a > b) || (b > a)

	fmt.Println(first_condition, second_condition)
}
