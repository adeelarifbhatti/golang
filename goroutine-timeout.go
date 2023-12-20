package main
import ("fmt"
		"time"
)

func main(){
	ch1 := make(chan int)
	go sendValue(ch1)
	select {
	case msg := <-ch1:
		fmt.Println("case 1 is executed ", msg)
	case  <- time.After(1* time.Second):
		fmt.Println("Case II is executed")
	}
}
func sendValue(ch chan int){
	time.Sleep(1*time.Second)
	ch <- 10
}
