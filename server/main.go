//go:generate protoc -I ../echo --go_out=plugins=grpc:../echo ../echo/echo.proto

package main

import (
	"log"
	"net"

	pb "github.com/ajalab/echo"
	"golang.org/x/net/context"
	"google.golang.org/grpc"
)

const (
	port = ":4000"
)

type server struct{}

// Echo implements echo.EchoServer
func (s *server) Echo(ctx context.Context, req *pb.EchoRequest) (*pb.EchoResponse, error) {
	addrs, err := net.InterfaceAddrs()
	if err != nil {
		return nil, err
	}

	addr := ""

	for _, a := range addrs {
		if ipnet, ok := a.(*net.IPNet); ok && !ipnet.IP.IsLoopback() {
			if ipnet.IP.To4() != nil {
				addr = ipnet.IP.String()
			}
		}
	}
	return &pb.EchoResponse{Message: req.GetMessage(), From: addr}, nil
}

func main() {
	lis, err := net.Listen("tcp", port)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	s := grpc.NewServer()
	pb.RegisterEchoServer(s, &server{})

	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
