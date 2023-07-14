test:
	go test -race ./...

build-app:
	go mod vendor
	CGO_ENABLED=0 go build -mod vendor -o ./dist/main cmd/main.go
	cp -r configs/ ./dist/configs
