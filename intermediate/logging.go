package main

import (
	"log"
	"os"
)

var infoLogger = log.New(os.Stdout, "INFO: ", log.Ldate|log.Ltime|log.Lshortfile)
var warnLogger = log.New(os.Stdout, "WARNING: ", log.Ldate|log.Ltime|log.Lshortfile)
var errorLogger = log.New(os.Stderr, "ERROR: ", log.Ldate|log.Ltime|log.Lshortfile)

func main() {
	log.Println("This is a log message")

	log.SetPrefix("Info : ")
	log.Println("this is a info message")

	log.SetFlags(log.Ldate | log.Ltime | log.Lshortfile)
	log.Println("this is a log message with date and time and file")

	infoLogger.Println("This is a info message from infoLogger")
	warnLogger.Println("This is a warning message from log")
	errorLogger.Println("This is a error message from log")

	logFile, err := os.OpenFile("app.log", os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
	if err != nil {
		log.Fatalf("Error opening file: %s", err.Error())
	}
	log.SetOutput(logFile)
	defer logFile.Close()

	DebugLogger := log.New(logFile, "DEBUG: ", log.Ldate|log.Ltime|log.Lshortfile)
	DebugLogger.Println("This is a debug message")

}
