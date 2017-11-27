package Test

import (
	"fmt"
)

func PrintVar() {
	fmt.Println(w)
}

//This function cannot be exported because it starts as lowecase
//So this is called unexported
func printVar() {
	fmt.Println("will not be exported")
}
