package main
import ("fmt"
		"log"
		"os")
func main(){
	file, err := os.OpenFile("/go/src/logging.log", os.O_APPEND|os.O_CREATE|os.O_WRONLY,0600)
	if err != nil {
		fmt.Println(err)
		return
	}

	log.SetOutput(file)
	log.Println("Writing log")
	
	log.Fatal("A Fatal Log")

}