package handler

import (
	"net/http"

	"github.com/FazylovAsylkhan/encoderUrl/internal/config"
	"github.com/FazylovAsylkhan/encoderUrl/internal/logger"
	"github.com/FazylovAsylkhan/encoderUrl/internal/usecase/encoderUrl"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/cors"
)

type handler struct {
	cfg *config.Config
	encoderUrl *encoderUrl.EncoderUrlStore
}



var h handler

func Init(cfg *config.Config) *chi.Mux {
	h = handler{
		cfg: cfg,
		encoderUrl: encoderUrl.Init(),
	}
	router := chi.NewRouter()
	options := cors.Options{
		AllowedOrigins:   []string{"http://*"},
		AllowedMethods:   []string{"GET", "POST"},
		AllowedHeaders:   []string{"Content-Type"},
		ExposedHeaders:   []string{"Link"},
		AllowCredentials: false,
		MaxAge:           300,
	}
	router.Use(cors.Handler(options))

	log := logger.New()
	log.SetFormatter(&logger.HandlerFormatter{})
	router.Use(func(next http.Handler) http.Handler {
		return logger.MiddlewareHandler(log, next)
	})

	router.HandleFunc("/", h.handlerUrl)
	router.HandleFunc("/{hash}", h.handlerHash)

	return router
}
