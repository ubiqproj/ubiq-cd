# Connect

> Connect is a family of libraries for building browser and gRPC-compatible HTTP APIs.

https://connectrpc.com/docs/go/getting-started

## Lint schema and generate code

```bash
# requires
go install github.com/bufbuild/buf/cmd/buf@latest
go install github.com/fullstorydev/grpcurl/cmd/grpcurl@latest
go install google.golang.org/protobuf/cmd/protoc-gen-go@latest
go install connectrpc.com/connectrpc/cmd/protoc-gen-connect-go@latest
```

```bash
make lint
make generate
```
