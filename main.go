/*
Copyright © 2025 NAME HERE <EMAIL ADDRESS>
*/
package main

import (
	"cli-todo-app/cmd"
	"io"
	"log"
	"os"
)

func main() {
	logFile, err := configureLogging()

	if err != nil {
		log.Fatal(err)
	}

	defer logFile.Close()

	cmd.Execute()
}

func configureLogging() (*os.File, error) {
	logFile, err := os.Create("logs.txt")

	if err != nil {
		return nil, err
	}

	logger := io.MultiWriter(os.Stderr, logFile)

	log.SetOutput(logger)

	return logFile, nil
}
