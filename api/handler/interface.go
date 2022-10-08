package handler

import (
	"context"

	"github.com/Igusaya/igusaya_blog/api/domain"
)

type AdminUsecase interface {
	SubmitArticle(ctx context.Context, a *domain.Article) error
}
