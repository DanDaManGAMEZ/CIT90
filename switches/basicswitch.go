package main

import "fmt"

func main() {

	var name string
	var input string

	input = "Enter your name: "

	fmt.Println(input)
	fmt.Scan(&name)

	switch name {
	case "Ruben":
		fmt.Println("Ayy lmao")
	case "Dan":
		fmt.Println("Hey...MYSELF")
	default:
		fmt.Println("Why hello there nobody.")
	}
}
