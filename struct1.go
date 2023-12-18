package main
import "fmt"

type student struct {
	name string
	rollno int
	marks []int
	grades map[string]int
}
func main(){
	var s student
	fmt.Printf("%+v\n",s)
}