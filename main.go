/*
Copyright © 2025 NAME HERE <EMAIL ADDRESS>
*/
package main

import (
	"cli-todo-app/cmd"
	"cli-todo-app/data"
	"database/sql"
	"io"
	"log"
	"os"

	"github.com/go-sql-driver/mysql"
	_ "github.com/joho/godotenv/autoload"
)

type Config struct {
	DB     *sql.DB
	Models data.Models
}

var db *sql.DB

func main() {
	logFile, err := configureLogging()

	if err != nil {
		log.Fatal(err)
	}

	defer logFile.Close()

	dbConnection := configureDatabase()

	_ = Config{
		DB:     dbConnection,
		Models: data.New(dbConnection),
	}

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

func configureDatabase() *sql.DB {
	// Capture connection properties.
	config := mysql.Config{
		User:   os.Getenv("DB_USERNAME"),
		Passwd: os.Getenv("DB_PASSWORD"),
		Net:    "tcp",
		Addr:   os.Getenv("DB_HOST") + ":" + os.Getenv("DB_PORT"),
		DBName: os.Getenv("DB_DATABASE"),
	}

	// Get a database handle.
	var err error

	db, err = sql.Open("mysql", config.FormatDSN())

	if err != nil {
		log.Fatal(err)
	}

	pingErr := db.Ping()

	if pingErr != nil {
		log.Fatal(pingErr)
	}

	return db
}
