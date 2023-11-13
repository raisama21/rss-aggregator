package main

import (
	"log"
	"net/http"
	"os"

	"github.com/go-chi/chi"
	"github.com/go-chi/cors"
	"github.com/joho/godotenv"
)

func main() {
	godotenv.Load(".env")

	ports := os.Getenv("PORT")
	if ports == "" {
		log.Fatal("PORT is not foudn in environment")
	}

	router := chi.NewRouter()

	router.Use(cors.Handler(cors.Options{
		AllowedOrigins:   []string{"https://*", "http://*"},
		AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowedHeaders:   []string{"Accept", "Authorization", "Content-Type", "X-CSRF-Token"},
		ExposedHeaders:   []string{"Link"},
		AllowCredentials: false,
		MaxAge:           300,
	}))

	v1Router := chi.NewRouter()
	v1Router.Get("/healthz", handlerRediness)
	v1Router.Get("/error", handleErr)

    router.Mount("/v1", v1Router)

	server := &http.Server{
		Handler: router,
		Addr:    ":" + ports,
	}

	log.Printf("Server running on port %v", ports)
	err := server.ListenAndServe()
	if err != nil {
		log.Fatal(err)
	}
}
