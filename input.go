package main

import (
	"fmt"
)

func main() {
	var name string
	var age int
	fmt.Print("Please enter the name\n")
	fmt.Scanf("%s", &name)
	fmt.Print("Please enter your age\n")
	fmt.Scanf("%d", &age)
	fmt.Print("Name is ", name, " age is ", age, "\n")
}
