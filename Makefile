BINARY_NAME=camille

install:
	go mod download

build:
	CGO_ENABLED=0 go build -o .builds/${BINARY_NAME} main.go
	chmod +x .builds/${BINARY_NAME}

start: build
	.builds/${BINARY_NAME}

clean:
	go clean
	rm -f ./.builds/*

dep:
	go mod tidy
	go mod vendor

lint:
	golangci-lint run

generate-api:
	oapi-codegen --config api/gen_config.yaml api/openapi.yaml
