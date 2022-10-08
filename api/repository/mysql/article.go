package mysql

import (
	"context"
	"fmt"

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
		return fmt.Errorf("db.ExecContext in mysql.InsertArticle: %w", err)
	}
	id, err := result.LastInsertId()
	if err != nil {
		return fmt.Errorf("result.LastInsertId in mysql.InsertArticle: %w", err)
	}
	a.ID = id
	return nil
}
