package usecase

import (
	"context"
	"fmt"

	"github.com/Igusaya/igusaya_blog/api/domain"
	"github.com/Igusaya/igusaya_blog/api/repository/mysql"
)

type AdminUsecase struct {
	DB          mysql.Execer
	ArticleRepo ArticleRepository
}

func NewAdminUsecase(
	DB mysql.Execer,
	ArticleRepo ArticleRepository,
) *AdminUsecase {
	return &AdminUsecase{
		DB:          DB,
		ArticleRepo: ArticleRepo,
	}
}

func (u *AdminUsecase) SubmitArticle(ctx context.Context, a *domain.Article) error {
	if err := u.ArticleRepo.InsertArticle(ctx, u.DB, a); err != nil {
		return fmt.Errorf("AdminUsecase.CreateArticle to CreateArticle: %w", err)
	}
	return nil
}
