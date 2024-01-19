build: build-frontend build-backend

build-frontend:
	cd frontend && npm run build

build-backend:
	cd backend && go build .

generate-backend:
	rm -rf backend/generated/*
	cd backend && swagger generate server --target=./generated --spec=../swagger.json --exclude-main --api-package=gaollery
	cd backend && go mod tidy

generate-frontend:
	rm -rf frontend/src/api
	cd frontend && echo "TODO!!!" && exit 1

run-backend:
	cd backend && PORT=5000 go run .

run-frontend:
	cd frontend && npm run dev
