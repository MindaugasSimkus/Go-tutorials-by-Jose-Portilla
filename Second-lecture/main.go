package main

//import "fmt"

//func main() {
//	var speed int
//	fmt.Println(speed)
//}

//third lecture - Path separator

import (
	"fmt"
	"path"
)

func main() {
	var dir, file string
	dir, file = path.Split("css/main.css")
	fmt.Println("dir :", dir)
	fmt.Println("file :", file)

	var file2 string
	_, file2 = path.Split("css2/main2.css")
	fmt.Println("file2 :", file2)

	//short declaration

	dir3, _ := path.Split("css3/main3.css")
	fmt.Println("dir3:", dir3)

}
