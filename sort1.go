package main
import (
		"fmt"
		"sort"
		)
func main(){
	vars := []int{5,1,2,5,7,7,8,3,2,2}
	vars2 := []string{"cat","Dog","spider","Apply"}
	sort.Ints(vars)
	fmt.Println(vars)
	sort.Strings(vars2)
	fmt.Println(vars2)
}