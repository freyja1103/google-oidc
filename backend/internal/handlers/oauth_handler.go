package handlers

import (
	"google-oidc/internal/repositories"
	"google-oidc/pkg/jwt"
	"google-oidc/pkg/logger"
	"log/slog"
	"net/http"

	"github.com/labstack/echo/v4"
	"golang.org/x/oauth2"
)

func NewOAuthHandler(oauthConf *oauth2.Config, oauthRepo repositories.OAuthRepository) *OAuthHandler {
	return &OAuthHandler{
		oauthConf: oauthConf,
		oauthRepo: oauthRepo,
	}
}

type OAuthHandler struct {
	oauthRepo repositories.OAuthRepository
	oauthConf *oauth2.Config
}

func (h *OAuthHandler) OAuthGoogle(c echo.Context) error {
	return c.Redirect(http.StatusFound, h.oauthConf.AuthCodeURL("state", oauth2.AccessTypeOffline, oauth2.ApprovalForce))
}

func (h *OAuthHandler) OAuthCallback(c echo.Context) error {
	code := c.QueryParam("code")

	token, err := h.oauthConf.Exchange(c.Request().Context(), code)
	if err != nil {
		return err
	}

	idToken, ok := token.Extra("id_token").(string)
	if !ok {
		logger.Error(c.Request().Context(), "id_token not found in token response")
		return c.NoContent(http.StatusInternalServerError)
	}

	claims, err := jwt.DecodeIDToken(idToken)
	if err != nil {
		logger.Error(c.Request().Context(), "failed to decode id_token", err)
		return c.NoContent(http.StatusInternalServerError)
	}

	logger.Info(c.Request().Context(), "id_token claims",
		slog.String("sub", claims.Sub),
		slog.String("email", claims.Email),
		slog.String("name", claims.Name),
	)

	err = h.oauthRepo.Create(c.Request().Context(), token)
	if err != nil {
		logger.Error(c.Request().Context(), "failed to create oauth2", err)
		return c.NoContent(http.StatusInternalServerError)
	}

	return c.JSON(http.StatusOK, map[string]any{
		"sub":   claims.Sub,
		"email": claims.Email,
		"name":  claims.Name,
	})
}
