package main

import (
	"context"
	"net/http"

	"github.com/Igusaya/igusaya_blog/api/gen/openapi"
	"github.com/Igusaya/igusaya_blog/api/handler"
	"github.com/go-chi/chi/v5"
)

func NewRouter(ctx context.Context) (http.Handler, error) {

	// hundler
	healthHandler := handler.NewHealthHandler()

	routers := []openapi.Router{
		openapi.NewHealthApiController(healthHandler, openapi.WithHealthApiErrorHandler(openapi.DefaultErrorHandler)),
	}
	r := chi.NewRouter()
	// middleware

	router := openapi.NewRouter(r, routers...)
	return router, nil
}
