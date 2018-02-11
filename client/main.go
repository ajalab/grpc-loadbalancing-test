package main

import (
	"context"
	"flag"
	"log"

	pb "github.com/ajalab/echo"
	"google.golang.org/grpc"
	"google.golang.org/grpc/balancer/roundrobin"
)

var (
	address = flag.String("addr", "echoserver:4000", "$host:$port to bind for rpc")
)

func main() {
	flag.Parse()
	conn, err := grpc.Dial(*address,
		grpc.WithInsecure(),
		grpc.WithBalancerName(roundrobin.Name))
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()
	c := pb.NewEchoClient(conn)

	for i := 0; i < 1000; i++ {
		r, err := c.Echo(context.Background(), &pb.EchoRequest{Message: "hoge"})
		if err != nil {
			log.Printf("could not echo: %v", err)
		}
		log.Printf("Message '%s' from %s", r.Message, r.GetFrom())
	}
}
