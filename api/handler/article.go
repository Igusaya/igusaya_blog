package handler

import (
	"context"
	"fmt"
	"net/http"

	"github.com/Igusaya/igusaya_blog/api/domain"
	"github.com/Igusaya/igusaya_blog/api/gen/openapi"
)

type ArticleHandler struct {
	AdminCase AdminUsecase
}

func NewArticleHandler(
	AdminCase AdminUsecase,
) *ArticleHandler {
	return &ArticleHandler{
		AdminCase: AdminCase,
	}
}

func (h *ArticleHandler) PostArticle(
	ctx context.Context, req openapi.CreateRequest,
) (openapi.ImplResponse, error) {
	fmt.Printf("req: %#v\n", req)
	a := &domain.Article{
		Subject: req.Subject,
		Body:    req.Body,
	}
	if err := h.AdminCase.SubmitArticle(ctx, a); err != nil {
		return openapi.Response(http.StatusInternalServerError, nil),
			fmt.Errorf("ArticleHandler.CreateArticle to AdminCase.CreateArticle: %w", err)
	}
	return openapi.Response(http.StatusOK, nil), nil
}
