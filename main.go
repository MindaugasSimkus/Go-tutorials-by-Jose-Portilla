// all lectures in the same file from lecture 4
package main

import "fmt"

//lecture 4 short vs normal declarations
//if you do not know the initial value, use variable declaration

func main() {
	//	score := 0 DON't!
	// var score int
	// var ( //better to use variable declarations for other developer understanding
	// 	//related
	// 	video string

	// 	//closely related
	// 	duration int
	// 	current  int
	// )
	width, height := 100, 50 // when you know initial value, shorter code
	fmt.Println("width =", width)
	fmt.Println("height =", height)
	//to keep code concise and for redeclaration
	//width = 50
	//color := "red"

	width, color := 50, "red" // same as above
	fmt.Println("width =", width)
	fmt.Println("height =", color)

}
