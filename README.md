# gRPC
gRPC tests

## Dependencies
### Proto Compiler
```bash
sudo apt install protobuf-compiler
sudo apt install golang-goprotobuf-dev
```

### Go protobuf support
Might not be needed:
```sh
go get -u github.com/golang/protobuf/protoc-gen-go  # or
go get github.com/golang/protobuf
go get github.com/golang/protobuf/proto
```

## Compilation
Use the Makefile targets or have fun with the below.
```
protoc --go_out=plugins=grpc:chat chat.proto
```

## Resources
```
https://tutorialedge.net/golang/go-protocol-buffer-tutorial/
https://tutorialedge.net/golang/go-grpc-beginners-tutorial/
https://grpc.io/docs/languages/go/basics/
https://developers.google.com/protocol-buffers
```
