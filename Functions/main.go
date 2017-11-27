package main

import (
	"fmt"
)

func main() {
	greet1("hassaan", "pasha")
	fmt.Println(greet2("Maha", "Kassim"))
	fmt.Println(greet3("Samara", "Hassaan"))
	a, b := greet4("First", "Last")
	fmt.Println(a, b)
	//calling variadic function
	fmt.Println(average(4, 5, 6, 7, 8, 9))

	//passing variadic arguments
	data := []float64{1, 2, 3, 4, 5, 6, 5, 2, 4, 5}
	fmt.Println(average(data...))

	//function expression
	greet5 := func() {
		fmt.Println("Hello Hassaan")
	}
	greet5()

	greet6 := makeGreeter()
	fmt.Println(greet6())

}

//a function that doesn't return anything (void)
func greet1(fname string, lname string) {
	fmt.Println(fname, lname)
}

//function that returns
func greet2(fname, lname string) string {
	return fmt.Sprint(fname, " ", lname)
}

//function with a names return
func greet3(fname, lname string) (s string) {
	s = fmt.Sprint(fname, " ", lname)
	return s
}

//function wit multiple returns
func greet4(fname, lname string) (string, string) {
	return fname, lname
}

//variadic function takes no or n number of arguments
func average(n ...float64) float64 {
	var total float64
	for _, v := range n {
		total += v
	}

	return total / float64(len(n))
}

//anonymous function
func makeGreeter() func() string {
	return func() string {
		return "Hello Anon Function"
	}
}
