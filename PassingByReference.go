package main
import "fmt"

func modify(s *string) {
	*s = "Passing by reference"
}
func main(){
	a := "hello"
	fmt.Println(a)
	modify(&a)
	fmt.Println(a)
}