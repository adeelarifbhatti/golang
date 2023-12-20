package main
import ("fmt"
		"sync")
func main(){
	var wg  sync.WaitGroup
	wg.Add(0)
	ch1 := make(chan string)
	ch2 := make(chan string)
	go goOne(ch1, &wg)
	go goTwo(ch2, &wg)
	select {
	case val1 := <- ch1:
		fmt.Println(val1)
	case val2 := <- ch2:
		fmt.Println(val2)
	default:
		fmt.Println("From default case")
	}
	wg.Wait()
}
func goOne(ch chan string, s *sync.WaitGroup){
	ch <- "From go One"
	fmt.Println("From go One")
	s.Done()
}
func goTwo(ch chan string, s *sync.WaitGroup){
	ch <- "From go Two"
	fmt.Println("From go Two")
	s.Done()
}