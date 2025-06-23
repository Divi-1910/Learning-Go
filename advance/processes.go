package main

import (
	"fmt"
	"io"
	"os/exec"
	"strings"
	"time"
)

func main() {
	cmd := exec.Command("echo", "Hello World")
	output, err := cmd.Output()
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(string(output))

	cmd = exec.Command("bash", "-c", "cd .. && ls")
	output, err = cmd.Output()
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(string(output))

	cmd = exec.Command("grep", "foo")
	cmd.Stdin = strings.NewReader("foo\nfaa\nfoo")
	output, err = cmd.Output()
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(string(output))

	cmd = exec.Command("sleep", "10")
	err = cmd.Start()
	if err != nil {
		fmt.Println(err)
		return
	}

	time.Sleep(1000 * time.Millisecond)
	err = cmd.Process.Kill()
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println("process killed")

	cmd = exec.Command("printenv", "SHELL")
	output, err = cmd.Output()
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(string(output))

	//err = cmd.Wait() // Blocking in nature
	//
	//if err != nil {
	//	fmt.Println(err)
	//	return
	//}

	pr, pw := io.Pipe()

	cmd = exec.Command("grep", "foo")
	cmd.Stdin = pr

	go func() {
		defer pw.Close()
		_, err = pw.Write([]byte("foo\nfood is love"))
		if err != nil {
			fmt.Println("Error in Writing to the pipe :", err)
			return
		}
	}()

	output, err = cmd.Output()
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(string(output))

	fmt.Println("Process is Completed ")

	cmd = exec.Command("ls", "-l")
	output, err = cmd.CombinedOutput()
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(string(output))

}

/*
Why use Process Spawning??
	Concurrency
	Isolation
	Resource Management

os/exec package
	exec.Command
	cmd.Stdin/ cmd.Stdout
	cmd.Start/ cmd.Wait
	cmd.Output

*/
