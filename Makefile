proto:
	protoc --go_out=.  --go-grpc_out=require_unimplemented_servers=false:. chat.proto --proto_path .
	protoc --go_out=. --go-grpc_out=require_unimplemented_servers=false:. person.proto --proto_path .

clean:
	rm -rf pb