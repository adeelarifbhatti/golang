package main
import ("fmt"
		"sync")

func main(){
	var wg sync.WaitGroup
	wg.Add(10)

	for i:=1;i<=10;i++{
		// concurrency, without passing i every go routine will not be running
		// because go routine doesn't start practically immediately 
		go func(i int){
			fmt.Println(i)
			wg.Done()
		}(i)
	}
	wg.Wait()
}