package main

import "fmt"

const p string = "hey dude"

//multi declaration
const (
	Pi       = 3.142
	Language = "Go"
)

const (
	A = iota
	B = iota
	C = iota
)

func main() {
	const q = 42

	fmt.Println("p - ", p)
	fmt.Println("q - ", q)
	fmt.Println("Pi - ", Pi)
	fmt.Println("Language - ", Language)

	fmt.Println(A, B, C)
}
