package main

import (
	"context"
	"log"
	"net"
	"time"

	"google.golang.org/grpc"
	pb "ElasticSearchGRPC"
	monitoring "monitoring"
)

const (
	port = ":50051"
)

// server is used to implement gRPC.ElasticSearchMonitoringServer.
type server struct {
	pb.UnimplementedElasticSearchMonitoringServer
}

func (s *server) GetStatus(ctx context.Context, in *pb.Cluster) (*pb.Status, error) {
	log.Printf("Received: %v \n", in.GetUrl())
	status, monitoringError := monitoring.GetStatus(in.GetUrl())
	if monitoringError != nil {
		log.Printf("[ERROR] Failed to listen: %+v \n", monitoringError)
		return nil, monitoringError
	}
	return &pb.Status{Color: status, Timestamp: time.Now().Unix()}, nil
}

func main() {
	log.SetPrefix("server-")

	lis, err := net.Listen("tcp", port)
	if err != nil {
		log.Printf("[ERROR] Failed to listen - %+v \n", err)
	}
	s := grpc.NewServer()
	pb.RegisterElasticSearchMonitoringServer(s, &server{})
	log.Println("server listening")
	if err := s.Serve(lis); err != nil {
		log.Printf("[ERROR] Failed to serve - %+v \n", err)
	}
}
