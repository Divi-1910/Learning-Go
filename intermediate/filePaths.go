package main

import (
	"fmt"
	"path/filepath"
	"strings"
)

/*
Key Concpets :
	1. Absolute Path
	2. Relative Path

`path/filepath` package
	Functions
		1. filepath.Join
		2. filepath.Split
		3. filepath.Clean
		4. filepath.Abs
		5. filepath.Base
		6. filepath.Dir

Best Practices
	1. Platform Independence
	2. Handling Errors
	3. Security

Practical Applications
 	1. File I/O Operations
	2. Directory Navigation
	3. Path Normalization

*/

func main() {
	relativePath := "./data/file.txt"
	absolutePath := "/home/divyansh/file.txt"
	fmt.Println(relativePath)
	fmt.Println(absolutePath)

	joinedPath := filepath.Join("home", "downloads", "file.zip")
	fmt.Println("Joined Path : ", joinedPath)

	normalizedPath := filepath.Clean("./data/../data/file.txt")
	fmt.Println("Normalized Path : ", normalizedPath)

	path, file := filepath.Split("/home/user/docs/file.txt")
	fmt.Println("Path : ", path)
	fmt.Println("File : ", file)

	fmt.Println(filepath.Base("/home/user/docs/file.txt"))

	fmt.Println(filepath.IsAbs("/home/user/docs/file.txt"))

	fmt.Println(filepath.Ext(file))
	fmt.Println(strings.TrimSuffix(file, filepath.Ext(file)))

	rel, err := filepath.Rel("a/b", "a/b/t/file.txt")
	if err != nil {
		panic(err)
	}
	fmt.Println(rel)

	rel, err = filepath.Rel("a/c", "a/b/t/file.txt")
	if err != nil {
		panic(err)
	}
	fmt.Println(rel)

	abs, err := filepath.Abs(relativePath)
	if err != nil {
		panic(err)
	}
	fmt.Println(abs)

}
