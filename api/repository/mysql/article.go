package mysql

import (
	"context"

	"github.com/Igusaya/igusaya_blog/api/domain"
)

func (r *Repository) InsertArticle(ctx context.Context, db Execer, a *domain.Article) error {
	sql := `INSERT INTO d_article
			(subject, body, modified, created)
			value (?, ?, ?, ?)`
	result, err := db.ExecContext(
		ctx, sql, a.Subject, a.Body, r.Clocker.Now(), r.Clocker.Now(),
	)
	if err != nil {
		return err
	}
	id, err := result.LastInsertId()
	if err != nil {
		return err
	}
	a.ID = id
	return nil
}
