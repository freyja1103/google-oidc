package google

import (
	"context"
	"encoding/json"
	"google-oidc/pkg/oauth2"
	"google-oidc/pkg/resp"

	golangoauth2 "golang.org/x/oauth2"
)

type API interface {
	GetUserInfo(ctx context.Context, token *golangoauth2.Token, req interface{}) error
}

type api struct {
	oauthConf *golangoauth2.Config
}

func NewGoogleAPI(oauthConf *golangoauth2.Config) API {
	return &api{oauthConf: oauthConf}
}

func (a *api) GetUserInfo(ctx context.Context, token *golangoauth2.Token, req interface{}) error {
	client := oauth2.NewOAuthClient(ctx, a.oauthConf, token)
	body, err := resp.Getter(client, "https://www.googleapis.com/oauth2/v2/userinfo")
	if err != nil {
		return err
	}

	if err := json.Unmarshal(body, &req); err != nil {
		return err
	}
	return nil
}
