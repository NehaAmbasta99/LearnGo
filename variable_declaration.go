package main

import "fmt"

func main() {
	//2 lines, declare and then initialise
	var num int
	num = 100


	//1 line but difficult to remember, declare and initialise 
	var floatNum float32 = 10.55
	fmt.Println(floatNum)

	//1 line, Compact one
	firstName := "Neha"

	fmt.Println(num,floatNum,firstName)
}