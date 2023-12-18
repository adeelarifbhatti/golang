package main
import "fmt"
type Circle struct {
	x int
	y int
	radius float64
	area float64
}
func calArea(c *Circle){
	const PI float64 = 3.14
	var area float64
	area = (PI *c.radius *c.radius)
	// with reference with RAM
	(*c).area = area
}
func main() {
	c := Circle{x: 5, y: 10, radius: 6, area: 0}
	fmt.Printf("%+v\n",c)
	// Address of memory for c struct
	calArea(&c)
	fmt.Printf("%+v\n",c)
}