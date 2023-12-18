package main
import "fmt"

type s1 struct {
	x int
}
func main(){
	c1 := s1{x: 6}
	c2 := s1{x: 7}
	c3 := s1{x: 6}
	if c1 != c2{
		fmt.Println("C1 is not equal to C2")
	}
	if c1 == c3 {
		fmt.Println("C1 is equal to C3")
	}
}