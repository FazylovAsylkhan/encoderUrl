package server

import (
	"log"
	"net/http"

	"github.com/FazylovAsylkhan/encoderUrl/config"
	handler "github.com/FazylovAsylkhan/encoderUrl/handler/http"
)

func Start(config *config.Config) {
	srv := &http.Server{
		Handler: handler.Init(config),
		Addr:    ":" + config.Port,
	}

	log.Printf("Server starting on port %v", config.Port)
	if err := srv.ListenAndServe(); err != nil {
		log.Fatalf("Server failed to start: %v", err)
	}
}

