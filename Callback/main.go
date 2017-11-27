package main

import (
	"fmt"
)

func visit(nums []int, callback func(int)) {
	for _, n := range nums {
		callback(n)
	}
}

func filter(nums []int, callback func(int) bool) []int {
	var xs []int
	for _, n := range nums {
		if callback(n) {
			xs = append(xs, n)
		}
	}
	return xs
}

func main() {
	visit([]int{1, 2, 3, 4, 5}, func(n int) {
		fmt.Println(n)
	})

	xs := filter([]int{1, 2, 3, 4, 5, 6, 7}, func(n int) bool {
		return n > 4
	})

	fmt.Println(xs)
}
