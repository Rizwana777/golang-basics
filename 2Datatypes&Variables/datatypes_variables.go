package main

import "fmt"

func main() {
	var a bool
	var b int
	var c string

	fmt.Println(a, b, c)

	a = true
	b = 3
	c = "hello"

	fmt.Println(a, b, c)

	raw := `This is raw string literal
	             which is defined in single quotes`

	intr := "This is interpreted string literal which is defined in double quotes"

	fmt.Println("Raw string literal : ", raw)
	fmt.Println("Interpreted string literal : ", intr)
}
