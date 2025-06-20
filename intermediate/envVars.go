package main

import (
	"fmt"
	"os"
	"strings"
)

func main() {
	user := os.Getenv("USER")
	home := os.Getenv("HOME")
	fmt.Println("User : ", user)
	fmt.Println("Home : ", home)

	// environment variables are key value pairs
	err := os.Setenv("FRUIT", "apple")
	if err != nil {
		fmt.Println("Error in Setting ENV Variable ", err.Error())
		return
	}
	fmt.Println("FRUIT : ", os.Getenv("FRUIT"))
	fmt.Println("Environment Variable Set Successfully ")

	err = os.Unsetenv("FRUIT")
	if err != nil {
		fmt.Println("Error in Unsetting ENV Variable ", err.Error())
		return
	}
	fmt.Println("FRUIT : ", os.Getenv("FRUIT"))

	EnvVars := os.Environ()
	fmt.Println("EnvVars : ", EnvVars)

	for _, v := range EnvVars {
		kvPair := strings.SplitN(v, "=", 2)
		fmt.Println(kvPair[0], kvPair[1])
	}

}
