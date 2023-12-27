


build: build-frontend
	go build .

build-frontend:
	go generate frontend/frontend.go
