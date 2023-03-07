package main
import "fmt"

 func printDetails(student string, subjects ...string){
	//  subjects are stored in slice
	fmt.Println("Name of the student is ",student)
	// _ is blank identifier
	for _, sub :=range subjects {
		fmt.Println("subjects are ",sub)
	}
 }
 func main(){
	printDetails("Adeel","Biology","Physics","Chemistry" )
 }