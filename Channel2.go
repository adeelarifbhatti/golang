package main
import ("fmt"
		"sync"
		)
func main(){
	var wg sync.WaitGroup
	wg.Add(2)
	// var i int
	// fmt.Print("Type a number: ")
	// fmt.Scan(&i)
	ch := make(chan int,3)
	go sell(ch, &wg)
	wg.Wait()
}
func sell(ch chan int, s *sync.WaitGroup){
	// for j := 0; j < i; j++{ 
	// 	ch <- i
	// }
	ch <- 10
	ch <- 11
	ch <- 12
	go buy(ch,s)
	s.Done()
}
func buy(ch chan int, s *sync.WaitGroup){
	fmt.Println("waiting for data")
	for i := 0; i < 3; i++{ 
	fmt.Println("Received data :", <-ch)
	}

	
	s.Done()
}