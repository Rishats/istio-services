package server_info

import (
	"context"
	"fmt"
	"github.com/Rishats/istio-services/pkg/server_info"
	"net"
	"os"
	"time"

	"google.golang.org/grpc"
)

type server struct {
	server_info.UnimplementedServerInfoServiceServer
}

func (s *server) GetServerInfo(ctx context.Context, in *server_info.GetServerInfoRequest) (*server_info.GetServerInfoResponse, error) {
	hostname, err := os.Hostname()
	if err != nil {
		hostname = "unknown"
	}
	return &server_info.GetServerInfoResponse{
		Hostname:    hostname,
		Environment: "development", // Or fetch from environment variable
		Version:     "1.0",         // Or define as a constant or a variable
		UpTime:      fmt.Sprintf("%s", time.Since(startTime)),
	}, nil
}

var (
	startTime = time.Now()
)

func main() {
	lis, err := net.Listen("tcp", ":50051")
	if err != nil {
		panic(err)
	}
	s := grpc.NewServer()
	server_info.RegisterServerInfoServiceServer(s, &server{})
	fmt.Println("Server is running on port 50051")
	if err := s.Serve(lis); err != nil {
		panic(err)
	}
}
