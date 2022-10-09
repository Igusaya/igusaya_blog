package main

import (
	"context"
	"fmt"
	"net/http"

	"github.com/Igusaya/igusaya_blog/api/clock"
	"github.com/Igusaya/igusaya_blog/api/config"
	"github.com/Igusaya/igusaya_blog/api/gen/openapi"
	"github.com/Igusaya/igusaya_blog/api/handler"
	"github.com/Igusaya/igusaya_blog/api/repository/mysql"
	"github.com/Igusaya/igusaya_blog/api/usecase"
	"github.com/go-chi/chi/v5"
)

func NewRouter(ctx context.Context, cfg *config.Config) (http.Handler, func(), error) {

	clocker := clock.RealClocker{}
	db, cleanup, err := mysql.New(ctx, cfg)
	if err != nil {
		fmt.Printf("err: %+v\n", err)
		return nil, cleanup, err
	}
	// repository
	repo := mysql.Repository{Clocker: clocker}
	// usecase
	adminUsecase := usecase.NewAdminUsecase(db, &repo)
	// hundler
	healthHandler := handler.NewHealthHandler()
	articleHandler := handler.NewArticleHandler(adminUsecase)

	routers := []openapi.Router{
		openapi.NewHealthApiController(healthHandler, openapi.WithHealthApiErrorHandler(openapi.DefaultErrorHandler)),
		openapi.NewArticleApiController(articleHandler, openapi.WithArticleApiErrorHandler(openapi.DefaultErrorHandler)),
	}
	r := chi.NewRouter()
	// middleware

	router := openapi.NewRouter(r, routers...)
	fmt.Printf("router: %#v\n", router)
	return router, cleanup, nil
}
