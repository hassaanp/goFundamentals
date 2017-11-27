package main

import (
	"fmt"
)

const metersToYards float64 = 1.09361

func main() {
	a := 3
	var meters float64
	fmt.Printf("%v %T\n", &a, &a)
	fmt.Print("Enter meters swam: ")
	fmt.Scan(&meters)
	yards := meters * metersToYards
	fmt.Println(meters, " meters is ", yards, " yards")
}
