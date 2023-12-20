package main
import (
	"fmt"
		"sync"
	)
func main() {
	var wg sync.WaitGroup
	wg.Add(2)
	ch := make(chan int)
	go sell(ch,&wg)
	go buy(ch,&wg)
	wg.Wait()
}
func sell(ch chan int, wg *sync.WaitGroup){
	for i:=0;i<100;i++{
	ch <- i
	}

	fmt.Println("Data is sent")
	close(ch)
	wg.Done()
}
func buy(ch chan int, wg *sync.WaitGroup){
	fmt.Println("Waiting for the Data")
	for val := range ch {
		fmt.Println("Receiving the data  ",val)
	}
	wg.Done()
}