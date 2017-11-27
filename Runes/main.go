package main

import "fmt"

func main() {
	fmt.Println([]byte("Hello"))

	for i := 0; i < 100; i++ {
		fmt.Println(i, " - ", string(i), []byte(string(i)))
	}

	//Rune in int32 alias
	//A rune is denoted by ''
	fmt.Printf("%v %T\n", 'i', 'i')
	fmt.Println(rune('a'))
}
