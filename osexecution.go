package main
import (
	"fmt"
	"os/exec"
	"runtime"
)

func execute() {
	out, err := exec.Command("/bin/bash", "-c" , "ls -l |awk '{print $9}'").Output()
	if err != nil {
		fmt.Printf("%s",err)
		fmt.Println("\nCommand failed")

	}
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