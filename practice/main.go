package main

import (
	"fmt"
)

// lowercase, so private
// declared outside main, must be const to be used in function
const aConst int = 64

func main() {

	// var = variable (pretty basic)
	// explicitly typed
	var aString string = "This is Go!"
	fmt.Println(aString)
	// note the \n for a newline in the output
	// using %T format specifier to
	// determine the datatype of the variables
	fmt.Printf("The variable's type is %T\n", aString)

	var anInteger int = 42
	fmt.Println(anInteger)

	var defaultInt int
	fmt.Println(defaultInt)

	// implicitly typed
	var anotherString = "This is another string"
	fmt.Println(anotherString)
	fmt.Printf("The variable's type is %T\n", anotherString)

	var anotherInt = 53
	fmt.Println(anotherInt)
	fmt.Printf("The variable's type is %T\n", anotherInt)

	// implied variable, can be used inside the function without var
	myString := "This is also a string"
	fmt.Println(myString)
	fmt.Printf("The variable's type is %T\n", myString)

	fmt.Println(aConst)
	fmt.Printf("The variable's type is %T\n", aConst)

}
