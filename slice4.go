package main
import "fmt"

func main() {

	src_slice := []int{10,20,30,40}
	dest_slice := make([]int,3)
	num := copy(dest_slice, src_slice)
	fmt.Println(dest_slice," is the dest slice and numbers copied are ", num)
	

}