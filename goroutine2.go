package main
import ("fmt" 
		"time")

func main(){
	go start()
	time.Sleep(1 * time.Second)
}
func start(){
	go stop()
	fmt.Println("in Start")
}
func stop(){
	fmt.Println("in stop")
}