# gRPC-demo
My first foray into gRPC

## Installation 
`brew install protobuf`

`go get google.golang.org/grpc` 

`go get -u github.com/golang/protobuf/protoc-gen-go`

## Compile the proto file
`protoc -I $SRC/protodemo service.proto --go_out=plugins=grpc:$SRC/protodemo` 


## Setup Go and Execute
```
go mod tidy

Terminal 1:
go run server/main.go

Terminal 2:
go run client/main.go
```

Then browse to `localhost:8080/add/<num1>/<num2>` and `localhost:8080/mult/<num1>/<num2>`, then check the results.

## Performance
```
[GIN] 2020/02/18 - 12:25:07 | 200 |     553.936µs |             ::1 | GET      /add/5/3
[GIN] 2020/02/18 - 12:26:17 | 200 |    1.289908ms |             ::1 | GET      /add/5/7
[GIN] 2020/02/18 - 12:27:23 | 200 |     773.553µs |             ::1 | GET      /mult/5/7
[GIN] 2020/02/18 - 12:29:01 | 200 |    1.458086ms |             ::1 | GET      /mult/5/5
```

