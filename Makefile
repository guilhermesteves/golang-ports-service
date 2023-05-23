.PHONY: generate

generate:
	protoc --proto_path=ps/internal/models --go_out=plugins=grpc:. ps.proto