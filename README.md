# gRPC iPlayer Programme

Project for learning gRPC.

# Usage

Run server with `go run main.go`

# Client

## grpcurl

`brew install grpcurl`

```
grpcurl -plaintext -d '{"pid": "b006m86d"}' localhost:5001 programme.ProgrammeService/GetProgramme
```
