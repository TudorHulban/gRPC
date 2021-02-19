chat:
	protoc --go_out=plugins=grpc:. chat.proto
	cp chat.pb.go cmd/server/chat
	cp chat.pb.go cmd/client/chat
	rm chat.pb.go

cleanchat:
	rm cmd/server/chat/chat.pb.go || true
	rm cmd/client/chat/chat.pb.go || true


person:
	protoc --go_out=plugins=grpc:. person.proto
	cp person.pb.go cmd/server/person
	cp person.pb.go cmd/client/person
	rm person.pb.go

cleanperson:
	rm cmd/server/person/person.pb.go || true
	rm cmd/client/person/person.pb.go || true
	
clean: cleanchat cleanperson