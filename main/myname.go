package main

import "fmt"

func main() {
	var myname string
	var name string

	name = "Dan"
	myname = "Hello " + name + "."

	fmt.Println(myname)

	LessCodeForSameResult()
}

func LessCodeForSameResult() {
	name := "Dan"
	fmt.Print("Hello ", name, ".")
}
