package main

import (
	"context"
	"flag"
	"fmt"
	"net"

	pb "github.com/boywei/grpc-examples/zero/zero"
	log "github.com/sirupsen/logrus"
	"google.golang.org/grpc"
)

var (
	port = flag.Int("port", 50051, "The server port")
)

type server struct {
	pb.UnimplementedGameServer
}

func (s *server) PlayBall(ctx context.Context, req *pb.BallRequest) (*pb.BallReply, error) {
	log.Printf("%s is playing ball ne", req.Name)
	return &pb.BallReply{
		Reply: req.Name + "在玩球呢",
	}, nil
}

func main() {
	flag.Parse()
	listener, err := net.Listen("tcp", fmt.Sprintf(":%d", *port))
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	s := grpc.NewServer()
	pb.RegisterGameServer(s, &server{})
	log.Printf("server listening at %v", listener.Addr())
	if err := s.Serve(listener); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
