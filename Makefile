build: build-frontend build-backend

build-frontend: generate-frontend
	cd frontend && npm run build

build-backend: generate-backend
	#cd backend && go generate ./... #TODO
	cd backend && go build .

generate-backend:
	jq empty swagger.json
	rm -rf backend/generated/*
	cd backend && swagger generate server --principal=models.User --template=stratoscale --target=./generated --spec=../swagger.json --exclude-main --api-package=gaollery
	cd backend && go mod tidy

generate-frontend:
	openapi-generator generate -i swagger.json -g typescript-fetch -o frontend/api --additional-properties=legacyDiscriminatorBehavior=false


run-backend:
	cd backend && PORT=5000 LOG_LEVEL=DEBUG go run .

run-frontend:
	cd frontend && npm run dev
