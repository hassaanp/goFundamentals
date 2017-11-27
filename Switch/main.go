package main

import (
	"fmt"
)

func main() {
	name := "Hassaan"
	switch name {
	case "Atif":
		fmt.Println("Atif")
	case "Hassaan":
		fmt.Println(name)
		fallthrough //automatically assumes next one is true
		// no need for break
	case "Pasha":
		fmt.Println("Pasha")
		fallthrough
	case "Imtiaz":
		fmt.Println("Imtiaz")
	case "HassaanP":
		fmt.Println("Again Hassaan")

	}

	switch name {
	case "Atif":
		fmt.Println("Atif")
	case "Hassaan", "Amir": //can give multiple cases, treats as OR
		fmt.Println(name, "Amir")

	}

	a := 42

	switch {
	case a%8 == 0:
		fmt.Println("divisible by 7")
	case a > 40:
		fmt.Println("greater than 40")
	}

	var x interface{} = "21"
	// switch can be used for type assertion
	switch x.(type) {
	case string:
		fmt.Println("string")
	case int:
		fmt.Println("int")
	}

}
