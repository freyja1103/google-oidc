package main

import (
	"context"
	"flag"
	"google-oidc/internal/config"
	"google-oidc/internal/handlers"
	"google-oidc/internal/repositories"
	"google-oidc/pkg/logger/sloghandler"
	"google-oidc/pkg/oauth2/google"
	"google-oidc/pkg/requestid"
	"log"
	"log/slog"
	"net/http"
	"os"

	"github.com/coreos/go-oidc/v3/oidc"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"golang.org/x/oauth2"
	yaml "gopkg.in/yaml.v3"
)

func main() {
	conf := ReadConfig()

	logger := slog.New(sloghandler.NewHandler(os.Stdout, &slog.HandlerOptions{
		Level:     slog.LevelDebug,
		AddSource: true,
	}))
	slog.SetDefault(logger)

	googleConfig := &oauth2.Config{
		ClientID:     conf.Client.Google.ID,
		ClientSecret: conf.Client.Google.Secret,
		RedirectURL:  conf.Client.Google.RedirectURL,
		Scopes: []string{
			"openid",
			"profile",
			"email",
			"https://www.googleapis.com/auth/youtube.readonly",
		},
		Endpoint: oauth2.Endpoint{
			AuthURL:       "https://accounts.google.com/o/oauth2/v2/auth",
			TokenURL:      "https://oauth2.googleapis.com/token",
			DeviceAuthURL: "https://oauth2.googleapis.com/device/code",
			AuthStyle:     oauth2.AuthStyleInParams,
		},
	}

	provider, err := oidc.NewProvider(context.Background(), "https://accounts.google.com")
	if err != nil {
		slog.Error("failed to create oidc provider", slog.Any("error", err))
		return
	}

	googleAPI := google.NewGoogleAPI(googleConfig)

	oauth2Repo := repositories.NewOAuthRepository(googleAPI)

	oauth2Handler := handlers.NewOAuthHandler(googleConfig, oauth2Repo, provider)

	e := echo.New()
	e.Use(
		middleware.Recover(),
		middleware.RequestID(),
		requestid.SetRequestID(logger),
	)

	e.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "Hello, World!")
	})
	e.GET("/oauth2/google", oauth2Handler.OAuthGoogle)
	e.GET("/oauth2/google/callback", oauth2Handler.OAuthCallback)

	e.Logger.Fatal(e.Start(":5001"))
}

func ReadConfig() *config.Config {
	var configFile string
	flag.StringVar(&configFile, "config", "config.yaml", "set config file path")
	flag.Parse()

	data, err := os.ReadFile(configFile)
	if err != nil {
		log.Fatalf("Error reading config file %s: %s", configFile, err)
	}

	var config config.Config

	if err := yaml.Unmarshal(data, &config); err != nil {
		log.Fatalf("Error unmarshaling config: %s", err)
	}

	log.Printf("Configuration loaded successfully from %s", configFile)
	return &config
}
