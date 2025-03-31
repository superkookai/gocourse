package basics

import "fmt"

// Global variable
var middleName string = "Cane"

func main() {
	var age int
	var name string = "Joe"
	var name1 = "Jane"

	//Inside a function, the := short assignment statement can be used in place of a var declaration with implicit type.
	//Outside function need to use var
	//Type inference
	count := 10
	lastname := "Smith"

	//Default Value
	//Numeric >> 0
	//Boolean >> false
	//String >> ""
	//Pointers, Slices, Maps, Functions, and Structs >> nil

	//Scope of variables >> Local variable in block scope
	fmt.Println(firstName)
	fmt.Println(middleName)
}

func printName() {
	firstName := "Micheal"
	fmt.Println(firstName)
}
