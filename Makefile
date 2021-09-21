VERSION ?= 0.0.1

create-proto:
	protoc --proto_path=. --go_out=pkg/ api/proto/v1/message/message.proto

start-server:
	go run cmd/server/main.go