package main

import "fmt"

func main() {

	//declare and initilisation of complex number
	c := complex(10, 15)

	//accesing real part of complex number
	fmt.Println(real(c))

	//accessing imaginary part of complex number
	fmt.Println(imag(c))
}
