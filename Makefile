generate:
	protoc --go_out=plugins=grpc:. chat.proto
	cp chat.pb.go cmd/server/chat
	cp chat.pb.go cmd/client/chat
	rm chat.pb.go
