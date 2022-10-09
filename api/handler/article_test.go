package handler

import (
	"context"
	"errors"
	"net/http"
	"strings"
	"testing"

	"github.com/Igusaya/igusaya_blog/api/domain"
	"github.com/Igusaya/igusaya_blog/api/gen/openapi"
	"github.com/Igusaya/igusaya_blog/api/handler/mocks"
	"github.com/golang/mock/gomock"
	"github.com/google/go-cmp/cmp"
)

func TestArticleHandler_PostArticle(t *testing.T) {
	t.Parallel()

	ctx := context.Background()
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	type args struct {
		ctx context.Context
		req openapi.CreateRequest
	}
	tests := []struct {
		name    string
		args    args
		mWant   *domain.Article
		mResult error
		want    openapi.ImplResponse
		wantErr error
	}{
		{
			"Success",
			args{ctx, openapi.CreateRequest{
				Subject: "subject",
				Body:    "body",
			}},
			&domain.Article{
				Subject: "subject",
				Body:    "body",
			},
			nil,
			openapi.Response(http.StatusOK, nil),
			nil,
		},
		{
			"Error AdminCase.SubmitArticle",
			args{ctx, openapi.CreateRequest{
				Subject: "subject",
				Body:    "body",
			}},
			&domain.Article{
				Subject: "subject",
				Body:    "body",
			},
			errors.New("dummy"),
			openapi.Response(http.StatusInternalServerError, nil),
			errors.New("ArticleHandler.CreateArticle to AdminCase.CreateArticle"),
		},
	}
	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			adminUsecase := mocks.NewMockAdminUsecase(ctrl)
			adminUsecase.EXPECT().SubmitArticle(gomock.Any(), tt.mWant).Return(tt.mResult)
			h := NewArticleHandler(adminUsecase)
			got, gotErr := h.PostArticle(tt.args.ctx, tt.args.req)
			if tt.wantErr == nil && gotErr != nil {
				t.Fatalf("unexpected err: %+v", gotErr)
			}
			if (tt.wantErr != nil && gotErr != nil) &&
				!strings.Contains(gotErr.Error(), tt.wantErr.Error()) {
				t.Fatalf("gotErr not contains wantErr\n +gotErr: %#v\n-wantErr: %#v", gotErr, tt.wantErr)
			}
			opts := []cmp.Option{}
			if gotErr == nil {
				if diff := cmp.Diff(tt.want, got, opts...); len(diff) != 0 {
					t.Fatalf("differs: (+got +want\n%s", diff)
				}
			}
		})
	}
}
