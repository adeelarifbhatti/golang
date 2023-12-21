package main
import ("fmt"
	"strings")
func main(){
	str := "This is first one"
	str2 := "first one"
	str3 := "is"
	result := strings.Contains(str,str2)
	fmt.Println(result)
	result1 := strings.Count(str,str3)
	fmt.Println(result1)
	str4 := "Learning standard library  in python"
	str5 := "python"
	result2 := strings.ReplaceAll(str4, str5,"golang")
	fmt.Println(result2)

}