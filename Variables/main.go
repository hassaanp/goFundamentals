package main

import (
	"fmt"
)

func main() {
	//Variable Notation
	//
	//declaration with static type
	var x string
	// declaration initializes it to the zero value of the type
	fmt.Printf("%v %T \n", x, x)

	//assignment
	x = "cowboy"

	var y bool
	fmt.Printf("%v %T \n", y, y)

	y = true

	//using dynamic type
	var z = 123

	fmt.Printf("%v %T \n", x, x)
	fmt.Printf("%v %T \n", y, y)
	fmt.Printf("%v %T \n", z, z)

	//Shorthand Notation
	//dynamic types
	a := 10
	b := "golang"
	c := 4.17
	d := true

	fmt.Printf("%v %T \n", a, a)
	fmt.Printf("%v %T \n", b, b)
	fmt.Printf("%v %T \n", c, c)
	fmt.Printf("%v %T \n", d, d)
}
