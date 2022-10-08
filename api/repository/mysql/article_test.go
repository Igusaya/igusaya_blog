package mysql

import (
	"context"
	"errors"
	"fmt"
	"strings"
	"testing"

	"github.com/Igusaya/igusaya_blog/api/clock"
	"github.com/Igusaya/igusaya_blog/api/domain"
	"github.com/Igusaya/igusaya_blog/api/testutil"
	"github.com/Igusaya/igusaya_blog/api/testutil/fixture"
	"github.com/google/go-cmp/cmp"
	"github.com/google/go-cmp/cmp/cmpopts"
)

func TestArticle_InsertArticle(t *testing.T) {
	t.Parallel()
	ctx := context.Background()
	tx, err := testutil.OpenDBForTest(t).BeginTxx(ctx, nil)
	t.Cleanup(func() { _ = tx.Rollback() })
	if err != nil {
		t.Fatalf("BeginTxx err: %+v", err)
	}
	type args struct {
		ctx context.Context
		db  Execer
		a   *domain.Article
	}
	a := fixture.Article(t, &domain.Article{Subject: "Success"})
	want := fixture.Article(t, a)
	errA := fixture.Article(t, &domain.Article{Subject: testutil.MakeRandomStr(t, 101)})
	tests := []struct {
		name    string
		args    args
		want    *domain.Article
		wantErr error
	}{
		{
			"Success",
			args{ctx, tx, a},
			want,
			nil,
		},
		{
			"Error db.ExecContext",
			args{ctx, tx, errA},
			nil,
			errors.New("db.ExecContext in mysql.InsertArticle"),
		},
	}

	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			r := Repository{Clocker: clock.FixedClocker{}}
			gotErr := r.InsertArticle(tt.args.ctx, tt.args.db, tt.args.a)
			if tt.wantErr == nil && gotErr != nil {
				t.Fatalf("unexpected err: %+v", gotErr)
			}
			if gotErr != nil && tt.wantErr != nil {
				fmt.Printf("got: %s\nwant: %s\n", gotErr.Error(), tt.wantErr.Error())
			}
			if (tt.wantErr != nil && gotErr != nil) &&
				!strings.Contains(gotErr.Error(), tt.wantErr.Error()) {
				t.Fatalf("gotErr not contains wantErr\n +gotErr: %#v\n-wantErr: %#v", gotErr, tt.wantErr)
			}
			opts := []cmp.Option{
				cmpopts.IgnoreFields(domain.Article{}, "ID"),
			}
			if gotErr == nil {
				if diff := cmp.Diff(tt.want, tt.args.a, opts...); len(diff) != 0 {
					t.Fatalf("differs: (+got +want\n%s", diff)
				}
			}
		})
	}
}
