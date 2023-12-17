package main
import "fmt"
func modify(s []int){
	s[0] = 100
}
func main(){
	slice := []int{10,20,30}
	fmt.Println(slice)
	// slice is reference to underlying array, thats why
	// the value changes (the value in the ram changes)
	modify(slice)
	fmt.Println(slice)
}