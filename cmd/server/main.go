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
	router, _ := handler.Init(config)
	srv := &http.Server{
		Handler: router,
		Addr:    config.Address,
	}
	logger.StartingServer(log, config.Address, config.BaseURL)
	if err := srv.ListenAndServe(); err != nil {
		log.Fatalf("Server failed to start: %v", err)
	}
}

