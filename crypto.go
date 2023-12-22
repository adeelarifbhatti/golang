package main
import ("fmt"
		"crypto/md5"
		"encoding/hex")
func encodeMD5(str string) string {
	var hashCode = md5.Sum([]byte(str))
	return hex.EncodeToString(hashCode[:])
}

func main(){
	fmt.Println("Enter the string you want to encrytp: ")
	var passwd string
	fmt.Scanln(&passwd)
	fmt.Println("Password is encrypted ", encodeMD5(passwd))
}