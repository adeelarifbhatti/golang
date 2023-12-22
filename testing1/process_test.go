package main
import ("testing"
		)


func TestCheckEven(t *testing.T){
	i := 10
	expected := "yes"
	res := checkEven(i)
	if expected != res {
		t.Errorf("expected : %v, got:%v ", expected, res)
	}
}