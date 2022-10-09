package handler

import (
	"context"

	"github.com/Igusaya/igusaya_blog/api/domain"
)

//go:generate mockgen -destination=./mocks/admin_usecase.go -package=mocks github.com/Igusaya/igusaya_blog/api/handler AdminUsecase
type AdminUsecase interface {
	SubmitArticle(ctx context.Context, a *domain.Article) error
}
