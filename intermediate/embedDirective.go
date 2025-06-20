package main

import (
	"embed"
	"fmt"
	"io/fs"
)

/*
	Why use Embed Directive??
		Simplicity
		Efficiency
		Security
		Default Values and Usage

	Supported Types
		Files
		Directories

	Use Cases
		Web Servers
		Configuration Files
		Testing

	Considerations
		File Size
		Update Strategy
		Compatibility

*/

//go:embed example.txt
var content string

//go:embed subdir
var basicsFolder embed.FS

func main() {
	fmt.Println("Embedded Content :", content)
	file, err := basicsFolder.ReadFile("subdir/file.go")
	if err != nil {
		fmt.Println("Error in Reading the file : ", err.Error())
		return
	}
	fmt.Println("Embedded File Content: ", string(file))

	err = fs.WalkDir(basicsFolder, "subdir", func(path string, d fs.DirEntry, err error) error {
		if err != nil {
			fmt.Println("Error in walking dir : ", err.Error())
			return err
		}
		fmt.Println("Path is : ", path)
		return nil
	})

	if err != nil {
		fmt.Println("Error in walking dir : ", err.Error())
		return
	}

}
