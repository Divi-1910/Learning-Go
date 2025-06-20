package main

import (
	"go.uber.org/zap"
	"log"
)

func main() {
	logger, err := zap.NewProduction()
	if err != nil {
		log.Fatal("Error initializing Zap logger")
	}
	defer logger.Sync()

	logger.Info("This is a Info message")

	logger.Info("User Logged In", zap.String("Name", "Divyansh"), zap.String("Her Name", "Sharvari"), zap.String("Method", "GET"))

}
