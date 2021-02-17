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
```
protoc --go_out=. *.proto
```

## Resources
```
https://tutorialedge.net/golang/go-protocol-buffer-tutorial/
https://grpc.io/docs/languages/go/basics/
https://developers.google.com/protocol-buffers
```