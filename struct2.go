package main
import "fmt"

type Student struct {
	name string
	rollno int
	marks []int
	grades map[string]int
}
func main(){
	st := new(Student)
	fmt.Printf("%+v\n", st)
}