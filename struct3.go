package main
import "fmt"

type Student struct {
	name string
	roleno int
}
func main(){
	st := Student{"John",35}
	fmt.Printf("%+v\n",st)
}