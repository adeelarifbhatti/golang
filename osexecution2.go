package main
import (
	"fmt"
	"os/exec"
	"runtime"
)

func execute() {
	out, err := exec.Command("/bin/bash", "-c" , "ls -l |awk '{print $9}'| sed  '/^$/d'").Output()
	out2, err2 := exec.Command("/bin/bash", "-c" , "ls -l|awk '{print $9}'| wc -l").Output()

	if err != nil || err2 != nil {
		fmt.Printf("%s",err," err2 is  ", err2)
		fmt.Print("Command failed")

	}
	output := string(out[:])
	fmt.Print("value of the out2 is ",string(out2[:]))
	fmt.Print(output)
	names := [...]string{output}

	for index, value := range names {
		fmt.Print(index, " => ", value)
		
	}
}
func main() {
	if runtime.GOOS == "windows" {
		fmt.Println("This node is windows")
	} else {
		execute()
	}
}