# Variables
APP_NAME=contact-crud
MAIN_PATH=./cmd/app
GO_FILES=$(shell find . -name "*.go" -type f)

# Comandos principales
.PHONY: help build run test clean deps docker-up docker-down

help: ## Mostrar ayuda
	@echo "Comandos disponibles:"
	@grep -E '^[a-zA-Z_-]+:.*?## .*$$' $(MAKEFILE_LIST) | awk 'BEGIN {FS = ":.*?## "}; {printf "\033[36m%-15s\033[0m %s\n", $$1, $$2}'

deps: ## Instalar dependencias
	go mod tidy
	go mod download

build: ## Compilar la aplicación
	go build -o bin/$(APP_NAME) $(MAIN_PATH)

run: ## Ejecutar la aplicación
	go run $(MAIN_PATH)/main.go

test: ## Ejecutar tests
	go test -v ./...

test-coverage: ## Ejecutar tests con cobertura
	go test -v -coverprofile=coverage.out ./...
	go tool cover -html=coverage.out -o coverage.html

clean: ## Limpiar archivos generados
	rm -rf bin/
	rm -f coverage.out coverage.html

docker-up: ## Levantar base de datos MySQL con Docker
	docker-compose up -d

docker-down: ## Detener contenedores Docker
	docker-compose down

docker-logs: ## Ver logs de Docker
	docker-compose logs -f

# Comandos de desarrollo
dev: docker-up deps run ## Configurar entorno y ejecutar

lint: ## Ejecutar linter
	golangci-lint run

format: ## Formatear código
	go fmt ./...

# Por defecto mostrar ayuda
.DEFAULT_GOAL := help
