package main

//imports are file level scopes
import (
	"fmt"
	"github.com/hassaanp/golangFundamentals/Scope/Test"
)

var x int = 23

func main() {
	//block level scope
	fmt.Println(x)
	fmt.Println(y)
	z := 11
	//order does matter in block level scope
	fmt.Println(z)
	foo()
	Test.PrintVar()
}

//order doesn't matter in package level scope
var y = 52

func foo() {
	fmt.Println(x)
}
