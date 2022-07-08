package main
import "fmt"

func main() {
	codes := map[string]string{
	"en" : "English",
	"fr" : "French",
	"ur" : "Urdu"}
	fmt.Println(codes["en"])
	fmt.Println(codes["ur"])
	codes["it"] = "Italian"
	value, found := codes["fr"]
	fmt.Println(found,value)
	fmt.Println(codes)

	delete(codes,"ur")
	fmt.Println(codes)

}