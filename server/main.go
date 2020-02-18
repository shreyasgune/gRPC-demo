package main

import(
	"context"
	proto "github.com/GannettDigital/gRPC-demo/protodemo"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
	"net"
)

//need to implement the functions defined in the interface on service.pb.go

type server struct{}

func main() {

	listener, err := net.Listen("tcp", ":4040")
	if err != nil {
		panic(err)
	}

	srv := grpc.NewServer() //new grpc server
	proto.RegisterAddServiceServer(srv, &server{}) //register our service on that server
	reflection.Register(srv) //for data serialization

	if e:= srv.Serve(listener); e != nil {
		panic(e)
	}

}

func (s *server) Add(ctx context.Context, request *proto.Request) (*proto.Response, error) {
	a,b := request.GetA(), request.GetB()
	result := a + b 
	return &proto.Response{Result: result}, nil 
}

func (s *server) Multiply(ctx context.Context, request *proto.Request) (*proto.Response, error) {
	a,b := request.GetA(), request.GetB()
	result := a * b 
	return &proto.Response{Result: result}, nil 
}

