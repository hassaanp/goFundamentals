package main

import (
	"fmt"
)

func main() {

	// C like for loop
	for i := 0; i < 10; i++ {
		fmt.Println(i)
	}

	// C like while loop
	i := 0
	for i < 20 {
		fmt.Println(i)
		i++
	}

	// C like for(;;) loop
	i = 0
	for {
		i++
		if i%2 == 0 {
			continue
		} else {
			fmt.Println(i)
		}
		if i >= 30 {
			break
		}
	}
}
