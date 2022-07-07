package main
import "fmt"

func main() {
	arr := [4]int{10,20,30,40}

	slice := arr[1:3]
	arr2 := [7]int{1,2,3,4,5,6,7}
	new_slice := arr2[0:5]
	fmt.Println(slice)
	slice = append(slice,900,23,11,33,2,2,2)

	fmt.Println(" after appending \n ",slice)
	fmt.Println("len of the slice is \n", len(slice))
	fmt.Println("capicity of the slice is \n", cap(slice))
	new_slice2 := append(new_slice, slice...)
	fmt.Println(new_slice2)
}
