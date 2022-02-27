package main

import (
	"fmt"
	"strconv"
)

func main() {
	var s string = "200a"
	i, err := strconv.Atoi(s)
	var n int = 200
	var m string = strconv.Itoa(n)
	fmt.Printf("%q", m)
	fmt.Print(" converted int in to string \n")
	fmt.Printf("%v %T \n", i, i)
	fmt.Printf("%v %T \n", err, err)

}
