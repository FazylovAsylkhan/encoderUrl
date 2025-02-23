package server

import (
	"net/http"

	"github.com/FazylovAsylkhan/encoderUrl/internal/config"
	handler "github.com/FazylovAsylkhan/encoderUrl/internal/handlers/http"
	"github.com/FazylovAsylkhan/encoderUrl/internal/logger"
)


func Start(config *config.Config) {
	log := logger.New()
	log.SetFormatter(&logger.ServerFormatter{})
	srv := &http.Server{
		Handler: handler.Init(config),
		Addr:    config.Address,
	}
	logger.StartingServer(log, config.Address, config.BaseURL)
	if err := srv.ListenAndServe(); err != nil {
		log.Fatalf("Server failed to start: %v", err)
	}
}

