package main

import (
	"fmt"
)

func main() {
	if true {
		fmt.Println("this ran")
	}

	if false {
		fmt.Println("this did not run")
	}

	food := "biryani"
	//initialization in if statement for tighter scope
	if yum := "yummy"; food == "biryani" {
		fmt.Println(yum)
	}
	// yum no longer accessible outside if
	//fmt.Println(yum)

	if false {
		fmt.Println("wont run")
	} else if true {
		fmt.Println("runs")
	} else {
		fmt.Println("wont run")
	}
}
