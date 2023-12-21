package main
import ("fmt"
	"errors")
func process(i int) error{
	if i%2!=0{
		return errors.New("Please enter even number")
	}
	return nil
}
func checkerror(e error) {
	if e != nil {
		fmt.Println(e)
	}
	if e == nil {
		fmt.Println("The number you entered is Even !")
	}
}
func main(){
	var number int
	fmt.Print("Please enter an even number: ")
	fmt.Scanln(&number)
	err := process(number)
	checkerror(err)
}