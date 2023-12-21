package main
import ("io"
		"os"
		"log"
		"strings"
	)
func main(){
	r := strings.NewReader("Something on the console\n")
	if _, err := io.Copy(os.Stdout,r); err != nil {
		log.Fatal(err)
	}
}