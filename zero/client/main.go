package main

import (
	"context"
	"flag"
	"time"

	pb "github.com/boywei/grpc-examples/zero/zero"
	log "github.com/sirupsen/logrus"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

const (
	defaultName = "wodetian"
)

var (
	addr = flag.String("addr", "localhost:50051", "the address to connect to")
	name = flag.String("name", defaultName, "name to play")
)

func main() {
	flag.Parse()
	conn, err := grpc.Dial(*addr, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Printf("did not connect: %v", err)
	}
	defer conn.Close()

	c := pb.NewGameClient(conn)
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	r, err := c.PlayBall(ctx, &pb.BallRequest{Name: *name})
	if err != nil {
		log.Fatalf("could not play: %v", err)
	}
	log.Printf("Status: %s", r.GetReply())
}
