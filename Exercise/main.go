package main

import (
	"fmt"
)

func computation(n int) (float64, bool) {
	// if n % 2 == 0{
	// 	return true
	// }
	return float64(n) / 2, n%2 == 0
}

func findGreatest(nums ...int) int {
	var greatest int
	for _, n := range nums {
		if n > greatest {
			greatest = n
		}
	}
	return greatest
}

func main() {
	fmt.Println(computation(43))

	computationV2 := func(n int) (float64, bool) {
		return float64(n) / 2, n%2 == 0
	}
	fmt.Println(computationV2(42))

	fmt.Println(findGreatest(1, 2, 3, 4, 5, 10, 8, 0))
}
