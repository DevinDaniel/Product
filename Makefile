

.PHONY: proto
proto:
	protoc --proto_path=. --micro_out=. --go_out=:. proto/product/product.proto

.PHONY: build
build: 

	SET CGO_ENABLED=0

    SET GOOS=linux

    SET GOARCH=amd64

    go build -o product-service main.go

.PHONY: test
test:
	go test -v ./... -cover

.PHONY: docker
docker:
	docker build . -t product-service:latest
