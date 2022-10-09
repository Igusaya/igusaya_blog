package usecase

import (
	"context"
	"errors"
	"strings"
	"testing"

	"github.com/Igusaya/igusaya_blog/api/domain"
	"github.com/Igusaya/igusaya_blog/api/testutil"
	"github.com/Igusaya/igusaya_blog/api/testutil/fixture"
	"github.com/Igusaya/igusaya_blog/api/usecase/mocks"
	"github.com/golang/mock/gomock"
)

func TestAdmin_SubmitArticle(t *testing.T) {
	t.Parallel()

	ctx := context.Background()
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	type args struct {
		ctx context.Context
		a   *domain.Article
	}
	a := fixture.Article(t, nil)
	tests := []struct {
		name    string
		args    args
		mWant   *domain.Article
		mResult error
		wantErr error
	}{
		{
			"Success",
			args{ctx, a},
			a,
			nil,
			nil,
		},
		{
			"Error ArticleRepo.InsertArticle",
			args{ctx, a},
			a,
			errors.New("dummy"),
			errors.New("AdminUsecase.CreateArticle to ArticleRepo.InsertArticle"),
		},
	}

	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			articleRepo := mocks.NewMockArticleRepository(ctrl)
			articleRepo.EXPECT().InsertArticle(gomock.Any(), gomock.Any(), tt.mWant).Return(tt.mResult)
			u := NewAdminUsecase(testutil.OpenDBForTest(t), articleRepo)
			gotErr := u.SubmitArticle(tt.args.ctx, tt.args.a)
			if tt.wantErr == nil && gotErr != nil {
				t.Fatalf("unexpected err: %+v", gotErr)
			}
			if (tt.wantErr != nil && gotErr != nil) &&
				!strings.Contains(gotErr.Error(), tt.wantErr.Error()) {
				t.Fatalf("gotErr not contains wantErr\n +gotErr: %#v\n-wantErr: %#v", gotErr, tt.wantErr)
			}
		})
	}
}
