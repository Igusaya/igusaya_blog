package usecase

import (
	"context"

	"github.com/Igusaya/igusaya_blog/api/domain"
	"github.com/Igusaya/igusaya_blog/api/repository/mysql"
)

type ArticleRepository interface {
	InsertArticle(ctx context.Context, db mysql.Execer, a *domain.Article) error
}
