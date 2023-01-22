package main

import "fmt"

func main() {

	// Print number from 1 to to 10
	for i := 1; i <= 10; i++ {
		fmt.Println(i)
	}

	fmt.Println("-------------------")

	// Print the array elements
	arr := []string{"Rizwana", "Naaz"}

	for i, v := range arr {
		fmt.Println(i+1, " ", v)
	}

}
