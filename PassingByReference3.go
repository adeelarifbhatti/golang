package main
import "fmt"

func modify(m map[string]int){
	m["K"] = 75
}
func main(){
	map1 := make(map[string]int)
	map1["A"] = 50
	map1["B"] = 60
	map1["C"] = 70
	fmt.Println(map1)
	// maps are passed by reference, thats why
	// the value changes (the value in the ram changes)
	modify(map1)
	fmt.Println(map1)
}