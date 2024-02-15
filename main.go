package main

import (
	"fmt"
	router "go-shadower/http"
	"go-shadower/tools"
)

func main() {
	loggerConfig := tools.LoggerConfig{
		FileName:    "log.txt",
		LogMatching: true,
	}
	logger, err := tools.NewLogger(&loggerConfig)
	if err != nil {
		panic(err)
	}

	router := router.NewRouter(&router.RouterConfig{
		Addr:         "localhost:8080",
		ReadTimeout:  5000,
		WriteTimeout: 5000,
		Target1Url:   "https://agilemanifesto.org",
		Target2Url:   "https://agilemanifesto.org",
	})
	router.Start()

	if err != nil {
		panic(err)
	}

	fmt.Print("Server started")
	logger.Log("Server started")
}
