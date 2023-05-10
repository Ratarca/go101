package main

import "fmt"

func main() {
	gdp := 1609 * 1000
	inflation := 0.05
	uneployment := 0.08
	interest_rate := 0.1
	population := 230000000
	name := "Brazil"

	fmt.Println(gdp, inflation, uneployment,
		interest_rate, population, name)
}
