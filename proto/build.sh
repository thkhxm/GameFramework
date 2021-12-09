#protoc --go_out=../pkg/$basePath/grpc/proto --go_opt=paths=source_relative --go-grpc_out=../pkg/$basePath/grpc/proto/service --go-grpc_opt=paths=source_relative *.proto
protoc --proto_path=. --go_out=../ *Message.proto
protoc --proto_path=. --go-grpc_out=../ *Service.proto
