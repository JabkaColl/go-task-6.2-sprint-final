package main

import (
	"log"

	"github.com/JabkaColl/go-task-6.2-sprint-final/internal/server"
)

func main() {
	logger := log.New(log.Writer(), "morse-server: ", log.LstdFlags|log.Lshortfile)

	srv := server.NewServer(logger)
	logger.Println("Сервер запущен на порту 8080...")

	if err := srv.HTTP.ListenAndServe(); err != nil {
		logger.Fatal(err)
	}
}
