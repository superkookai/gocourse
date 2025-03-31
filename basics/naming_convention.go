package basics

import "fmt"

type EmployeeGoogle struct {
	FirstName string
	LastName  string
	Age       int
}

type EmployeeApple struct {
	FirstName string
	LastName  string
	Age       int
}

func main() {
	//PascalCase >> Structs, Interfaces, Enums Eg. UserInfo, NewHTTPRequest

	//snake_case >> Variables, File name Eg. user_id, first_name, http_request

	//UPPERCASE >> Constants (Immutable)

	//mixedCase >> Camel case with lowercase at first character

	const MAXRETIRED = 5

	var employeeId = 123
	fmt.Println("Employee ID: ", employeeId)
}
