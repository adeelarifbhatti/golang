package main
import ("fmt"
		"os"
		)
func main() {
	path := "/go/src/slice.go"
	data, err := os.ReadFile(path)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println("######## Data is ########\n",string(data))
	fmt.Println("######## Byte data is #######\n",data)
}