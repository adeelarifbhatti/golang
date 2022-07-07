package main
import "fmt"

func main() {

	arr := [10]int{ 10,20,30,40,50,60,70,80,90,100}

	big_slice := arr[1:5]
	fmt.Println(big_slice)

	little_slice := big_slice[0:8]
	fmt.Println(little_slice)
	fmt.Println("capacity is ",cap(little_slice))
}