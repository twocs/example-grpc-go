# example-grpc-go

This is an example gRPC in Go. Follows the general guidance from <https://grpc.io/docs/languages/go/quickstart/>.

## Installation instructions

### you need golang

### install protoc version 3

Follow guide here: <https://grpc.io/docs/protoc-installation/>

### ensure that the protoc compiler can find the plugins

```sh
export PATH="$PATH:$(go env GOPATH)/bin"
```
