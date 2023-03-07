package main
import (
	"fmt"
	"os/exec"
	"runtime"
)

func execute() {
	out, err := exec.Command("ls", "-lts").Output()
	if err != nil {
		fmt.Printf("%s",err)
	}
	fmt.Println("Command executed successfully")
	output := string(out[:])
	fmt.Println(output)
}
func main() {
	if runtime.GOOS == "windows" {
		fmt.Println("This node is windows")
	} else {
		execute()
	}
}