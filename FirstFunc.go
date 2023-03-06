package main
import "fmt"

func returnCube(n int) int {
	return n*n*n
}

func main() {
	var i int = 5
	var j = returnCube(i)
	fmt.Print("The value is ",j,"\n")
}