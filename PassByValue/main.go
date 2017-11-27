//everything in go is pass by value
package main

import (
	"fmt"
)

func changeMe(p *int) {
	fmt.Println(p)
	fmt.Println(*p)
	*p = 23
	fmt.Println(*p)
}

//slice is a reference type so no need to pass address
func changeMeString(s []string) {
	s[0] = "Pasha"
}

func main() {
	n := 45
	fmt.Println(&n)
	fmt.Println(n)
	changeMe(&n)
	fmt.Println(n)

	name := make([]string, 1, 25)
	fmt.Println(name)
	changeMeString(name)
	fmt.Println(name)
}
