package handler

import (
	"context"
	"net/http"

	"github.com/Igusaya/igusaya_blog/api/gen/openapi"
)

type HealthHandler struct{}

func NewHealthHandler() *HealthHandler { return &HealthHandler{} }

func (h *HealthHandler) HealthGet(ctx context.Context) (openapi.ImplResponse, error) {
	return openapi.Response(http.StatusOK, openapi.Health{Status: "ok"}), nil
}
