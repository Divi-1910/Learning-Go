package main

import "github.com/sirupsen/logrus"

func main() {
	log := logrus.New()

	// Set log Level
	log.SetLevel(logrus.InfoLevel)

	// set Log Format
	log.SetFormatter(&logrus.JSONFormatter{})

	// example
	log.Info("This is a Info message")
	log.Warn("This is a Warning message")
	log.Error("This is a Error message")

	log.WithFields(logrus.Fields{
		"username": "Guest",
		"method":   "GET",
	}).Info("User Logged in ")

}
