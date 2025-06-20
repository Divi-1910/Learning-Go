package main

import (
	"fmt"
	"os"
	"path/filepath"
)

/*
	Key Concepts
		1. os.Mkdir
		2. os.MkdirAll
		3. os.ReadDir
		4. os.Chdir
		5. os.Remove
		6. os.RemoveAll

	Best Practices
		1. Error Handling
		2. Permissions
		3. Cross Platform Compatibility

	Practical Applications
		1. Organizing Files
		2. File System Navigation
		3. Batch Processing

*/

func CheckError(err error) {
	if err != nil {
		fmt.Println("Fatal error ", err.Error())
		panic(err)
	}
}

func main() {
	err := os.Mkdir("subdir", 0755)
	CheckError(err)

	//defer func() {
	//	err = os.RemoveAll("subdir")
	//	CheckError(err)
	//}()

	err = os.WriteFile("subdir/file", []byte("Hello"), 0755)
	CheckError(err)

	err = os.MkdirAll("subdir/parent/child", 0755)
	CheckError(err)
	err = os.MkdirAll("subdir/parent/child1", 0755)
	CheckError(err)
	err = os.MkdirAll("subdir/parent/child2", 0755)
	CheckError(err)
	err = os.MkdirAll("subdir/parent/child3", 0755)
	CheckError(err)
	err = os.MkdirAll("subdir/parent/child4", 0755)
	CheckError(err)
	err = os.WriteFile("subdir/parent/file.txt", []byte("Hello"), 0755)
	CheckError(err)
	err = os.WriteFile("subdir/parent/child1/file.txt", []byte("Hello"), 0755)
	CheckError(err)

	directories, err := os.ReadDir("subdir/parent")
	CheckError(err)

	for _, entry := range directories {
		fmt.Println(entry.Name(), entry.IsDir())
	}

	err = os.Chdir("subdir/parent/child1")
	result, err := os.ReadDir(".")
	fmt.Println(result)
	CheckError(err)

	dir, err := os.Getwd()
	fmt.Println(dir)
	CheckError(err)

	err = os.Chdir("../../..")
	CheckError(err)

	dir, err = os.Getwd()
	fmt.Println(dir)
	CheckError(err)

	// filepath.Walk and filepath.WalkDir
	pathfile := "subdir"
	fmt.Println("Walking directory: ", pathfile)
	err = filepath.WalkDir(pathfile, func(path string, d os.DirEntry, err error) error {
		CheckError(err)
		fmt.Println(path)
		return nil
	})
	CheckError(err)

	err = os.RemoveAll("subdir") // to remove directory and the file it contains
	CheckError(err)

}
