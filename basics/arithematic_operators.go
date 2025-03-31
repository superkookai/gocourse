package main

import (
	"fmt"
	"math"
)

func main() {
	var a, b int = 10, 3
	var result int

	result = a + b
	fmt.Println("a + b = ", result)

	result = a - b
	fmt.Println("a - b = ", result)

	result = a * b
	fmt.Println("a * b = ", result)

	//int divide int >> return int
	result = a / b
	fmt.Println("Integer Devision: a / b = ", result)

	//Float Division
	const p float64 = 22 / 7.0
	fmt.Println("Float Division: ", p)

	result = a % b
	fmt.Println("a % b = ", result)

	//Overflow >> more than maximum numeric can store
	//Underflow >> less than minimum numeric can store

	// Overflow with signed integers
	var maxInt int64 = 9223372036854775807 // max value that int64 can hold
	fmt.Println(maxInt)

	maxInt = maxInt + 1
	fmt.Println(maxInt)

	// Overflow with unsigned integers
	var uMaxInt uint64 = 18446744073709551615 // max value for uint64 type
	fmt.Println(uMaxInt)

	uMaxInt = uMaxInt + 1
	fmt.Println(uMaxInt)

	// Underflow with floating point numbers
	var smallFloat float64 = 1.0e-323
	fmt.Println(smallFloat)

	smallFloat = smallFloat / math.MaxFloat64
	fmt.Println(smallFloat)

}
