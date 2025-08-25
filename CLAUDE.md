# CLAUDE.md

This file provides guidance to Claude Code (claude.ai/code) when working with code in this repository.

## Project Overview

This is a Google OpenID Connect (OIDC) sample application written in Go. The application demonstrates OAuth2 authentication with Google accounts, specifically targeting YouTube API access. The project is written in Japanese and serves as an educational example for implementing OIDC flows.

## Architecture

The codebase follows a standard Go application structure:

- `backend/cmd/api/main.go` - Main application entry point and HTTP server setup using Echo framework
- `backend/internal/config/` - Configuration management with YAML parsing
- `backend/internal/repositories/` - Data access layer (currently with stub OAuth repository)
- `backend/pkg/logger/` - Structured logging with custom slog handler
- `backend/pkg/requestid/` - Request ID middleware for request tracing
- `config/config.yaml` - Application configuration (OAuth credentials, cookie settings, database config)

The application uses:
- Echo v4 for HTTP routing and middleware
- Go OAuth2 library for Google authentication
- Structured logging with slog
- YAML for configuration management

## Commands

### Running the Application
```bash
just run
# or manually:
cd backend && go run ./cmd/api/main.go --config "./config/config.yaml"
```

### Development Commands
- `just` or `just --list` - Show available commands
- `go mod tidy` - Update dependencies (run from backend/ directory)
- `go build ./cmd/api` - Build the application (run from backend/ directory)

## Configuration

The application reads configuration from `config/config.yaml` which includes:
- Google OAuth2 client credentials (ID, secret, redirect URL)
- Cookie configuration for session management
- Database connection settings (currently unused)

**Security Note**: The config file may contain sensitive credentials. Ensure proper handling of OAuth secrets in production environments.

## Key Components

- **OAuth Flow**: Configured for Google OAuth2 with YouTube API scopes
- **Logging**: Custom structured logging with request ID correlation
- **Middleware**: Request recovery, request ID generation, and custom logging middleware
- **Repository Pattern**: Abstracted data access layer (stub implementation currently)

The main application currently serves a basic "Hello, World!" endpoint at `/` on port 5001, with OAuth configuration prepared but not yet fully implemented in the routing layer.