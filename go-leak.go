package main
import ("fmt"
		"sync")
func main(){
	var wg sync.WaitGroup
	wg.Add(2)
	go leak(&wg)
	wg.Wait()
}
func leak(s *sync.WaitGroup){
	ch := make(chan int)
	go func(){
		// This will be keep on waiting for the value and no channel
		// will send the value to it
		val := <-ch
		fmt.Println("received Data at ch ",val)
		s.Done()
	}()
	fmt.Println("Existing Leak Method")
	s.Done()
}