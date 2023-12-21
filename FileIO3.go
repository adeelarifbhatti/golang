package main
import ("fmt"
		"os"
		)
func main() {
	path := "/go/src/README.md"
	file, err := os.Open(path)
	if err != nil {
		fmt.Println("err is ",err)
	}
	b := make([]byte,8)
	for {
		n, err := file.Read(b)
		if err != nil {
			fmt.Println("err from for loop is ",err)
			break
		}
		fmt.Println(string(b[0:n]))
	}
}