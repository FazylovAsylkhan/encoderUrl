package httpService

import (
	"log"
	"net/http"

	"github.com/FazylovAsylkhan/encoderUrl/config"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/cors"
)

func Start(cfg *config.Config) {
	srv := &http.Server{
		Handler: getRouter(),
		Addr:    ":" + cfg.Port,
	}

	log.Printf("Server starting on port %v", cfg.Port)
	if err := srv.ListenAndServe(); err != nil {
		log.Fatalf("server failed to start: %v", err)
	}
}

func getRouter() *chi.Mux {	
	router := chi.NewRouter()
	options := cors.Options{
		AllowedOrigins:   []string{"http://*"},
		AllowedMethods:   []string{"GET", "POST"},
		AllowedHeaders:   []string{"*"},
		ExposedHeaders:   []string{"Link"},
		AllowCredentials: false,
		MaxAge:           300,
	}
	router.Use(cors.Handler(options))

	return router
}