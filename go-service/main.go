package main

import (
	"log"
	"net"

	"google.golang.org/grpc"
	"github.com/gbengafagbola/microservice/pb"
)

type server struct{}

func (s *server) SayHello(req *pb.HelloRequest, stream pb.Greeter_SayHelloServer) error {
	log.Printf("Received: %v", req.Name)
	resp := &pb.HelloResponse{Message: "Hello, " + req.Name}
	return stream.Send(resp)
}

func main() {
	listener, err := net.Listen("tcp", ":50051")
	if err != nil {
		log.Fatalf("Failed to listen: %v", err)
	}
  
	srv := grpc.NewServer()
	pb.RegisterGreeterServer(srv, &server{})

	log.Println("Server is listening on :50051")
	if err := srv.Serve(listener); err != nil {
		log.Fatalf("Failed to serve: %v", err)
	} 
}
