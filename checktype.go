package main

import (
	"fmt"
	"reflect"
)

func main() {

	var name string
	var age int
	fmt.Print("Please enter name \n")
	fmt.Scanf("%s", &name)
	fmt.Print("Please enter the age \n")
	fmt.Scanf("%d", &age)
	fmt.Print("Type of name is ", reflect.TypeOf(name))
	fmt.Print("\nType of age is ", reflect.TypeOf(age), "\n")

}
