package main

import "fmt"

func main() {
	fmt.Println("----- 4 TASK --------")
	customer := "basic"

	switch customer {
	case "basic":
		fmt.Println("Basic Client")
	case "premium":
		fmt.Println("Premium Client")
	default:
		fmt.Println("Another Client")
	}
}
