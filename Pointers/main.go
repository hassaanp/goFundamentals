package main

import (
	"fmt"
)

func main() {
	a := 34

	fmt.Println(a)
	fmt.Println(&a)

	b := &a
	//var b *int = &a
	fmt.Println(b)
	//dereference a pointer to get value at memory address b or &a
	fmt.Println(*b)

	*b = 43
	fmt.Println(a)

	//everything is pass by value in go!!!
	//it means when we pass a mem addr, we are passing a value

	x := 5
	zero(x, &x) //obviously this does nothing to x because of scope
	fmt.Println(x)
}

func zero(x int, y *int) {
	//this on the other hand changes the value of the variable x in main
	*y = 123

	//this is a local copy
	//doesn't affect the variable that was passed
	x = 0
}
