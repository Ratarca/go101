package main

import (
	"fmt"
	"math/rand"
)

func main() {
	var heinsenbergAge, darwinAge int = rand.Intn(110), rand.Intn(110)

	// if
	fmt.Println(">> Decision: Single if")
	if heinsenbergAge > darwinAge {
		fmt.Println("Heinseinberg is older than Darwin")
	} else {
		fmt.Println("Darwin is older than Heinsenberg")
	}

	if heinsenbergAge > 10 {
		fmt.Println(heinsenbergAge)
	}

	// if/else if/ else
	fmt.Println(">> Decision: else if")
	var inflation, yield int = rand.Intn(110), rand.Intn(110)
	if inflation < yield {
		fmt.Println("Put my money in Stocks")
	} else if inflation == yield {
		fmt.Println("Put my money in Bonds and Stocks")
	} else {
		fmt.Println("Put my money in Bonds")
	}

	// switch : "Study pattern matching"
	fmt.Println("Decision: switch")

	// switch (without statement, without expression)
	switch {
	case darwinAge == 100:
		fmt.Println("Magick with 100")
	case darwinAge == 50:
		fmt.Println("Magick with 50")
	default:
		fmt.Println("Yooung")
	}
}
