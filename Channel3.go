package main
import ("fmt")
func main(){
	ch := make(chan int, 10)
	for j := 0; j < 10; j++{ 
	ch <- j	
	}
	close(ch)
	for j := 0; j < 15; j++{ 		
	val, ok := <- ch
	fmt.Println("Value in ch ",val, " Channel status ",ok )	
	}
}
