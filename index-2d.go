package main
import "fmt"

func main() {
	arr := [3][2]int{{2,4},{5,10},{8,64}}
	//fmt.Println(arr[2][1])
	fmt.Println("Len is ",len(arr))


	for i := 0; i < len(arr); i++ {

		for j := 0; j < len(arr)-1; j++ {
			fmt.Println("# value of i is # ",i," # value of  j is # ",j)

			fmt.Println(" ",arr[i][j])
		}

	}
}