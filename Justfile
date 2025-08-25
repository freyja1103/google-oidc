default:
  @just --list

run:
  cd backend && go run ./cmd/api/main.go --config "./../config/config.yaml"