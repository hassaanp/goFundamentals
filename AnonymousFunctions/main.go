package main

import (
	"fmt"
)

func main() {
	func() {
		fmt.Println("I am executing myself")
		fmt.Println("Also I am anonymous")
	}()
}
