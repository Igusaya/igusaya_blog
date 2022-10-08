package fixture

import (
	"math/rand"
	"testing"
	"time"

	"github.com/Igusaya/igusaya_blog/api/domain"
	"github.com/Igusaya/igusaya_blog/api/testutil"
)

func Article(t *testing.T, a *domain.Article) *domain.Article {
	result := &domain.Article{
		ID:       rand.Int63(),
		Subject:  testutil.MakeRandomStr(t, 20),
		Body:     testutil.MakeRandomStr(t, 300),
		Tags:     nil, // TODO:　After making Tag fixture, prepare it.
		Modified: time.Now(),
		Created:  time.Now().AddDate(0, 0, -10),
	}
	if a == nil {
		return result
	}
	if a.ID != 0 {
		result.ID = a.ID
	}
	if a.Subject != "" {
		result.Subject = a.Subject
	}
	if a.Body != "" {
		result.Body = a.Body
	}
	if a.Tags != nil {
		result.Tags = a.Tags
	}
	if !a.Modified.IsZero() {
		result.Modified = a.Modified
	}
	if !a.Created.IsZero() {
		result.Created = a.Created
	}
	return result
}
