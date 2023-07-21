test:
	go test -race ./...

lint:
	golangci-lint run ./... --config ./build/golangci-lint/config.yaml

build-app:
	go mod vendor
	CGO_ENABLED=0 go build -mod vendor -o ./dist/main cmd/main.go
	cp -r configs/ ./dist/configs
