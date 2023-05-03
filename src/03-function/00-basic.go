package main

import "fmt"

// Single return
func hipotenuse(a int, b int) int {
	return a + b
}

// Multiple return
func calculus(a, b int) (int, string) {
	sum := a + b
	return sum, "ANOTHER RETURN"
}

func main() {
	//
	l3 := hipotenuse(1, 3)
	fmt.Println(l3)

	//
	var f = func(txt string) string {
		return txt
	}

	f("42")

	//
	results, _ := calculus(10, 5)
	fmt.Println(results)

}
