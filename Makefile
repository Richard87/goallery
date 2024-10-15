MAKEFLAGS += -j2


build-frontend: generate-frontend
	cd frontend && npm run build

build-backend: generate-backend bootstrap
	cd backend && go generate ./...
	cd backend && go build .

build: build-frontend build-backend

generate-backend:
	rm -rf backend/api/api.go
	cd backend && oapi-codegen --config=api/api.server.yaml ../openapi.yaml
	cd backend && oapi-codegen --config=api/api.models.yaml ../openapi.yaml
	cd backend && oapi-codegen --config=api/api.swagger.yaml ../openapi.yaml
	cd backend && templ generate
	cd backend && go mod tidy

generate-frontend:
	openapi-generator generate -i openapi.json -g typescript-fetch -o frontend/api --additional-properties=legacyDiscriminatorBehavior=false

generate: generate-backend generate-frontend

run-backend: generate-backend
	cd backend && LOG_LEVEL=DEBUG go run .

run-frontend:
	cd frontend && npm run dev

run: run-backend run-frontend

lint-backend:
	cd backend && golangci-lint run --max-same-issues 0

lint-frontend:
	cd frontend && npm run lint

lint: lint-backend lint-frontend


HAS_SWAGGER       := $(shell command -v swagger;)
HAS_GOLANGCI_LINT := $(shell command -v golangci-lint;)
HAS_MOCKERY       := $(shell command -v mockery;)

bootstrap:
ifndef HAS_SWAGGER
	go install github.com/go-swagger/go-swagger/cmd/swagger@v0.30.5
endif
ifndef HAS_GOLANGCI_LINT
	go install github.com/golangci/golangci-lint/cmd/golangci-lint@v1.55.2
endif
ifndef HAS_MOCKERY
	go install github.com/vektra/mockery/v2@v2.40.1
endif
