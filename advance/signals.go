package main

import (
	"fmt"
	"os"
	"os/signal"
	"syscall"
	"time"
)

/*
Why use signals
	Graceful shutdowns
	Resource Cleanup
	Inter-Process Communication

Signals in Unix like OS
	SIGINT - interrupt
	SIGTERM - terminate
	SIGHUP - hang up
	SIGKILL - kill

*/

func main() {
	pid := os.Getpid()
	fmt.Printf("PID: %d\n", pid)
	sigs := make(chan os.Signal, 1)

	// Notify channel on interrupt or terminate signals
	signal.Notify(sigs, syscall.SIGINT, syscall.SIGTERM, syscall.SIGUSR1)

	go func() {
		sig := <-sigs
		switch sig {
		case syscall.SIGINT:
			fmt.Println("Received SIGINT")
		case syscall.SIGTERM:
			fmt.Println("Received SIGTERM")
		case syscall.SIGHUP:
			fmt.Println("Received SIGHUP")
		case syscall.SIGUSR1:
			fmt.Println("Received SIGUSR1")
			fmt.Println("User defined function is executed")
		default:
		}
		// fmt.Println("Recieved signal", sig)
		fmt.Println("Exiting Gracefully...")
		os.Exit(0)
	}()

	// simulate some work
	fmt.Println("Working...")
	for {
		time.Sleep(1 * time.Second)
	}

}
