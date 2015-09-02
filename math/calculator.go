package main

import "fmt"

func main() {
	var first int
	var second int = 10
	var input string

	input = "Enter your multiplication problem here: "

	fmt.Println(input)

	fmt.Scan(&first)
	fmt.Println(first * second)
}
