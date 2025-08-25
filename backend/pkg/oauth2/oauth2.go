package oauth2

import (
	"context"
	"net/http"

	"golang.org/x/oauth2"
)

func NewOAuthClient(ctx context.Context, oauthConfig *oauth2.Config, token *oauth2.Token) *http.Client {
	return oauthConfig.Client(ctx, token)
}
