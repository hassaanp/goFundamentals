package main

import (
	"fmt"
)

func main() {
	//arrays are not dynamic
	//cannot change size
	var x [20]int
	fmt.Println(len(x))
	//length is part of the type of array
	fmt.Printf("%T\n", x)

	for i := 0; i < len(x); i++ {
		x[i] = i
	}

	fmt.Println(x)
}
