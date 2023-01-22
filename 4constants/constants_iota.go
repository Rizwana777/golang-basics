package main

import "fmt"

func main() {

	const (
		Monday    = 0
		Tuesday   = 1
		Wednesday = 2
		Thursday  = 3
		Friday    = 4
		Saturday  = 5
		Sunday    = 6
	)

	fmt.Println(Monday, Tuesday, Wednesday, Thursday, Friday, Saturday, Sunday)

	const (
		Monday1 = iota
		Tuesday1
		Wednesday1
		Thursday1
		Friday1
		Saturday1
		Sunday1
	)

	fmt.Println(Monday1, Tuesday1, Wednesday1, Thursday1, Friday1, Saturday1, Sunday1)

	// use as expression
	// -5, -7,-8
	const (
		count1 = -(5 + iota) // -(5+0) = -5
		_                    // -(5+1) = -6
		_                    // -(5+2) = -7
		count3               // -(5+3) = -8
	)

	fmt.Println(count1, count3)

}
