package server

import (
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/Yandex-Practicum/go1fl-sprint6-final/internal/handlers" // Импортируем пакет handlers
)

// Структура сервера с логгером и HTTP-сервером
type Server struct {
	logger *log.Logger
	srv    *http.Server
}

// Функция для создания нового сервера с роутером и хендлерами
func NewServer(logger *log.Logger) *Server {
	mux := http.NewServeMux()
	mux.HandleFunc("/", handlers.IndexHandler)
	mux.HandleFunc("/upload", handlers.UploadHandler)

	srv := &http.Server{
		Addr:         ":8080",
		Handler:      mux,
		ErrorLog:     logger,
		ReadTimeout:  5 * time.Second,
		WriteTimeout: 10 * time.Second,
		IdleTimeout:  15 * time.Second,
	}

	return &Server{logger: logger, srv: srv}
}

// Метод для запуска сервера
func (s *Server) Start() error {
	fmt.Println("Starting server on: 8080")
	return s.srv.ListenAndServe()
}
