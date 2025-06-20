package main

import (
	"flag"
	"fmt"
	"os"
)

func main() {
	stringFlag := flag.String("user", "Guest", "User Name")
	flag.Parse()
	fmt.Println(*stringFlag)

	subCommand1 := flag.NewFlagSet("firstSub", flag.ExitOnError)
	subCommand2 := flag.NewFlagSet("secondSub", flag.ExitOnError)

	firstFlag := subCommand1.Bool("processing", false, "Command Processing Status")
	secondFlag := subCommand1.Int("bytes", 1024, "Byte Length of Result")

	flagsc2 := subCommand2.String("language", "Go", "Enter your Language")

	if len(os.Args) < 2 {
		fmt.Println("This Program requires additional Commands")
		os.Exit(1)
	}

	switch os.Args[1] {
	case "firstSub":
		subCommand1.Parse(os.Args[2:])
		fmt.Println("SubCommand1: ", subCommand1)
		fmt.Println("Processing : ", *firstFlag)
		fmt.Println("Bytes : ", *secondFlag)
	case "secondSub":
		subCommand2.Parse(os.Args[2:])
		fmt.Println("SubCommand2: ", *subCommand2)
		fmt.Println("Language :", *flagsc2)
	default:
		fmt.Println("No Sub Command Found!!")
		os.Exit(1)
	}

}
