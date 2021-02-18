# gRPC
gRPC tests

## Dependencies
### Proto Compiler
```bash
sudo apt  install protobuf-compiler
```

### Go protobuf support
```
go get -u github.com/golang/protobuf/protoc-gen-go  # or
go get github.com/golang/protobuf
go get github.com/golang/protobuf/proto
```

## Compilation
Use the Makefile target or have fun with the below.
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
