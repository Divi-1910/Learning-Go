package main

import (
	"fmt"
	fuck "net/http"
)

func main() {
	fmt.Println("Hello from go standard library")

	response, err := fuck.Get("https://jsonplaceholder.typicode.com/posts/1")
	if error != nil {
		fmt.Println("Error :", err)
		return
	}
	defer response.Body.Close()

	fmt.Println("HTTP Response Status : ", response.Status)

}
