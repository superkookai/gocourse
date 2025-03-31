package main

//Name import >> foo "net/http"
import (
	"fmt"
	foo "net/http"
)

func main() {
	fmt.Println("Hello Go Standard Library")

	resp, error := foo.Get("https://jsonplaceholder.typicode.com/posts/1")
	if error != nil {
		fmt.Println("Error: ", error)
	}
	defer resp.Body.Close()
	fmt.Println("HTTP Response Status: ", resp.Status)
}
