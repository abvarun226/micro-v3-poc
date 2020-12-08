
GOPATH:=$(shell go env GOPATH)
MODIFY=Mgithub.com/micro/micro/proto/api/api.proto=github.com/micro/micro/v3/proto/api

.PHONY: init
init:
	go get -u github.com/golang/protobuf/proto
	go get -u github.com/golang/protobuf/protoc-gen-go
	go get github.com/micro/micro/v3/cmd/protoc-gen-micro

.PHONY: proto
proto:    
	protoc --proto_path=. --micro_out=${MODIFY}:. --go_out=${MODIFY}:. proto/comment.proto
    

.PHONY: build
build: proto
	go build -o comments-service *.go

.PHONY: test
test:
	go test -v ./... -cover

.PHONY: docker
docker:
	docker build . -t comments-service:latest