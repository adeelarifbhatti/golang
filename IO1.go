package main
import ("fmt"
		"strings"
	)

func main(){
	r := strings.NewReader("Learning is good, currently learning IO")
	buf := make([]byte,8)
	for {
	n, err := r.Read(buf)
	fmt.Println(buf[:n]," Ascii Presentation is :",string(buf[:n]),err)
		if err != nil {
			fmt.Println("Breaking out, err == nil")
			break
		}
	}
}