package main
import "fmt"

func modify(s string) {
	s = "Passing by value"
}
func main(){
	a := "hello"
	fmt.Println(a)
	modify(a)
	fmt.Println(a)
}