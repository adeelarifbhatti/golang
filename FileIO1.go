package main
import ("fmt"
		"path/filepath"
		"os"
	)
func main() {
	path := filepath.Join("dir1","dir2","file.txt")
	directorypath := filepath.Dir(path)
	filename := filepath.Base(path)
	AbsPath := filepath.IsAbs(path)
	fileextension := filepath.Ext(path)
	fmt.Println("filepath is ",path)
	fmt.Println("Directorypath is ",directorypath)
	fmt.Println("File name is ",filename)
	fmt.Println("Abs path is ",AbsPath)
	fmt.Println("Extension path is ",fileextension)
	fileinfo, err := os.Stat("IO1.go")
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(fileinfo.Name())
	fmt.Println(fileinfo.Size())
	fmt.Println(fileinfo.Mode())
	fmt.Println(fileinfo.IsDir())
}