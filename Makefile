


build: build-frontend
	go build .

build-frontend:
	go generate frontend/frontend.go

generate-backend:
	rm -rf backend/generated/*
	cd backend && swagger generate server --target=./generated --spec=../swagger.json --exclude-main --api-package=gaollery
	cd backend && go mod tidy

run-backend:
	cd backend && PORT=5000 go run .
