package repositories

import (
	"context"
	"google-oidc/pkg/logger"
	"google-oidc/pkg/oauth2/google"
	"log/slog"

	"golang.org/x/oauth2"
)

type oauthRepository struct {
	googleAPI google.API
}

type OAuthRepository interface {
	Create(ctx context.Context, token *oauth2.Token) error
}

func NewOAuthRepository(googleAPI google.API) OAuthRepository {
	return &oauthRepository{googleAPI: googleAPI}
}

func (r *oauthRepository) Create(ctx context.Context, token *oauth2.Token) error {
	var userInfo any
	if err := r.googleAPI.GetUserInfo(ctx, token, &userInfo); err != nil {
		return err
	}
	logger.Info(ctx, "get userinfo", slog.Any("user", userInfo))
	return nil
}
