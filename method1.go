package main
import "fmt"

type circle struct {
	radius float64
	area float64
}
// with * pointer is passed 
func (c *circle) calarea(){
	c.area= 3.14 * c.radius * c.radius

}

func main() {
	c := circle{radius: 5}
	c.calarea()
	fmt.Printf("%+v\n",c)
}