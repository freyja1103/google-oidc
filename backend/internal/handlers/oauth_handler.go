package handlers

import (
	"google-oidc/internal/repositories"
	"google-oidc/pkg/logger"
	"log/slog"
	"net/http"

	"github.com/coreos/go-oidc/v3/oidc"
	"github.com/labstack/echo/v4"
	"golang.org/x/oauth2"
)

func NewOAuthHandler(oauthConf *oauth2.Config, oauthRepo repositories.OAuthRepository, provider *oidc.Provider) *OAuthHandler {
	return &OAuthHandler{
		oauthConf:    oauthConf,
		oauthRepo:    oauthRepo,
		oidcProvider: provider,
	}
}

type OAuthHandler struct {
	oauthRepo    repositories.OAuthRepository
	oauthConf    *oauth2.Config
	oidcProvider *oidc.Provider
}

func (h *OAuthHandler) OAuthGoogle(c echo.Context) error {
	return c.Redirect(http.StatusFound, h.oauthConf.AuthCodeURL("state", oauth2.AccessTypeOffline, oauth2.ApprovalForce))
}

func (h *OAuthHandler) OAuthCallback(c echo.Context) error {
	code := c.QueryParam("code")

	token, err := h.oauthConf.Exchange(c.Request().Context(), code)
	if err != nil {
		logger.Error(c.Request().Context(), "failed to exchange token", slog.Any("error", err))
		return c.NoContent(http.StatusInternalServerError)
	}

	rawIDToken, ok := token.Extra("id_token").(string)
	if !ok {
		logger.Error(c.Request().Context(), "id_token not found in token response")
		return c.NoContent(http.StatusInternalServerError)
	}

	idToken, err := h.oidcProvider.Verifier(&oidc.Config{ClientID: h.oauthConf.ClientID}).Verify(c.Request().Context(), rawIDToken)
	if err != nil {
		logger.Error(c.Request().Context(), "failed to verify id_token", slog.Any("error", err))
		return c.NoContent(http.StatusInternalServerError)
	}

	claims := new(claims)
	if err := idToken.Claims(&claims); err != nil {
		logger.Error(c.Request().Context(), "failed to parse id_token claims", slog.Any("error", err))
		return c.NoContent(http.StatusInternalServerError)
	}

	logger.Info(c.Request().Context(), "id_token claims",
		slog.String("email", claims.Email),
		slog.String("name", claims.Name),
	)

	err = h.oauthRepo.Create(c.Request().Context(), token)
	if err != nil {
		logger.Error(c.Request().Context(), "failed to create oauth2", slog.Any("error", err))
		return c.NoContent(http.StatusInternalServerError)
	}

	return c.JSON(http.StatusOK, map[string]any{
		"email": claims.Email,
		"name":  claims.Name,
	})
}

type claims struct {
	Email string `json:"email"`
	Name  string `json:"name"`
}
