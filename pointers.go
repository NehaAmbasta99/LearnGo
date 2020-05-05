package main

import "fmt"

func main() {
	
//Pointer declared and initialised
var lastName *string = new(string)
*lastName = "Ambasta"

//Will print the memory location of lastName
fmt.Println(&lastName)

}