package main
import "fmt"
type shape interface {
	Area() float64
	Perimeter() float64
}
type Rectangle  struct {
	Length, Width float64
}
func (r Rectangle) Area() float64 {
	return r.Length * r.Width
}
func (r Rectangle) Perimeter() float64{
	return 2 * (r.Length * r.Width)
}
func main() {
	var s shape = Rectangle {Length: 4.0, Width: 5.0}
	fmt.Println(s.Area())
	fmt.Println(s.Perimeter())
}