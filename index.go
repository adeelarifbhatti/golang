package main
import "fmt"

func main(){
	var cars [2]string = [2]string{"toyota","mehran"}
	fmt.Println(cars)

	numbers := [3]int{10,20,30}
	fmt.Println(numbers)

	names := [...]string{"John","Adeel","Rahul","Ola"}
	fmt.Println(names)
}