package main

import (
	"log"
	"os"

	"github.com/Yandex-Practicum/go1fl-sprint6-final/internal/server"
)

func main() {
	logger := log.New(os.Stdout, "morse-app", log.LstdFlags|log.Lshortfile)

	srv := server.Router(logger)

	err := srv.HTTP.ListenAndServe()
	if err != nil {
		logger.Fatal("Ошибка запуска сервера: ", err)
	}
}
