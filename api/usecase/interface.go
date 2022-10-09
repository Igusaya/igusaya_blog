package usecase

import (
	"context"

	"github.com/Igusaya/igusaya_blog/api/domain"
	"github.com/Igusaya/igusaya_blog/api/repository/mysql"
)

//go:generate mockgen -destination=./mocks/article_repository.go -package=mocks github.com/Igusaya/igusaya_blog/api/usecase ArticleRepository
type ArticleRepository interface {
	InsertArticle(ctx context.Context, db mysql.Execer, a *domain.Article) error
}
