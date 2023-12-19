package main
import ("fmt"
		"sync")

func main(){
	var wg sync.WaitGroup
	wg.Add(2)
	ch := make(chan string)
	go sell(ch, &wg)
	go buy(ch, &wg)
	wg.Wait()

	
}
// sending data to channel
func sell(ch chan string, wg *sync.WaitGroup){
	ch <- "Package Ready"
	fmt.Println("sending the data")
	wg.Done()
}
func buy(ch chan string, wg *sync.WaitGroup){
	fmt.Println("waiting for the data")
	val := <- ch
	fmt.Print("Data received is ", val, "\n")
	wg.Done()
}