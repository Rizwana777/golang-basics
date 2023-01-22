package main

import "fmt"

func main() {
	num := 3

	// if statements
	if num == 3 {
		fmt.Println("num is 3")
	}

	fmt.Println("num is not 3")

	// if-else
	num = 5

	if num != 5 {
		fmt.Println("num is not 5")
	} else {
		fmt.Println("num is 5")
	}

	num = 10

	if num == 10 {
		fmt.Println("num is 10")
	} else if num == 5 {
		fmt.Println("num is 5")
	} else {
		fmt.Println("num is neither 5 nor 10")
	}

}
