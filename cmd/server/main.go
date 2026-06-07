package main

import (
	"log"
	"net/http"
	"time"

	_ "github.com/soorajsomans/url-shortener/docs"
	"github.com/soorajsomans/url-shortener/internal/generator"
	"github.com/soorajsomans/url-shortener/internal/handler"
	"github.com/soorajsomans/url-shortener/internal/repository"
	"github.com/soorajsomans/url-shortener/internal/service"
)

func main() {
	repo := repository.NewInMemoryURLRepository()

	idGenerator := generator.NewAtomicIDGenerator()

	codeGenerator := generator.NewBase62Generator()

	urlService := service.NewURLService(
		repo,
		idGenerator,
		codeGenerator,
	)

	urlHandler := handler.NewURLHandler(
		urlService,
	)

	mux := http.NewServeMux()

	urlHandler.RegisterRoutes(mux)
	handler.RegisterSwaggerRoutes(mux)

	server := &http.Server{
		Addr:              ":8080",
		Handler:           mux,
		ReadTimeout:       5 * time.Second,
		WriteTimeout:      10 * time.Second,
		IdleTimeout:       30 * time.Second,
		ReadHeaderTimeout: 2 * time.Second,
	}

	log.Println("Starting server on : 8080")

	log.Fatal(server.ListenAndServe())
}
