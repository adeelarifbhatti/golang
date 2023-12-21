package main
import ("fmt"
		"os"
		)
func main() {
	file, err := os.OpenFile("/go/src/FileWriting.txt",os.O_CREATE|os.O_APPEND|os.O_WRONLY, 0600)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer file.Close()
	_, err = file.WriteString("The file worked\n")
	if err != nil {
		fmt.Println(err)
	}

}