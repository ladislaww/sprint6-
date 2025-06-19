package server

import (
	"log"
	"net/http"
	"time"

	"github.com/Yandex-Practicum/go1fl-sprint6-final/internal/handlers"
)

type Server struct {
	HTTP *http.Server
	Logger *log.Logger
}

func Router(log *log.Logger) *Server {
	mux := http.NewServeMux()

	mux.HandleFunc("/", handlers.HtmlHandler)
	mux.HandleFunc("/upload", handlers.UploadHandler)

	httpServer := &http.Server{
		Addr: ":8080",
		Handler: mux,
		ErrorLog: log,
		ReadTimeout: 5 * time.Second,
		WriteTimeout: 10 * time.Second,
		IdleTimeout: 15 * time.Second,
	}

	return &Server{
		HTTP: httpServer,
		Logger: log,
	}

}
