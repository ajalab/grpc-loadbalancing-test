package main

import (
	"context"
	"flag"
	"fmt"
	"log"
	"time"

	pb "github.com/ajalab/grpc_loadbalancing_test/echo"
	"google.golang.org/grpc"
	"google.golang.org/grpc/balancer/roundrobin"
)

var (
	dns  = flag.String("dns", "kubernetes.default", "dns server")
	host = flag.String("host", "echoserver", "host name")
	port = flag.String("port", "4000", "port number")
)

func main() {
	flag.Parse()

	conn, err := grpc.Dial(
		fmt.Sprintf("dns://%s/%s:%s", *dns, *host, *port),
		grpc.WithInsecure(),
		grpc.WithBalancerName(roundrobin.Name),
	)

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
		time.Sleep(1 * time.Second)
	}
}
