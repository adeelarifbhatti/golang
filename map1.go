package main
import "fmt"

func main() {
	codes := map[string]string{
	"en" : "English",
	"fr" : "French",
	"ur" : "Urdu"}
	fmt.Println(codes["en"])
	fmt.Println(codes["ur"])
	// Adding a key
	codes["it"] = "Italian"
	// checking if the value exists
	value, found := codes["fr"]
	fmt.Println(found,value)
	fmt.Println(codes)
	// deleting the value
	delete(codes,"ur")
	fmt.Println(codes)
	//for loop for the key/value
	for key, value := range codes {
		fmt.Println(key, "=>" ,value)
		// deleting the map
		// delete(codes,key)
		// or another method is to reinitial the map
		// codes = make(map[string]string)
	}

}