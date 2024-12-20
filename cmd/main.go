package main

import (
	"FirstSprintExam/internal/app"
	"log"
	"os"
)

// Run фукнции для запуска http сервера

func setLogger() {
	log.SetOutput(os.Stdout)
}

func main() {
	setLogger()
	app.RunServer()
}
