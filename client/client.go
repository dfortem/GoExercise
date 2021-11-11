package main

import (
	"context"
	"log"
	"os"
	"time"

	"google.golang.org/grpc"
	pb "ElasticSearchGRPC"
)

const (
	address     = "localhost:50051"
	clusterURL = "http://localhost:9200"
)

func main() {
	log.SetPrefix("client-")
	// Set up a connection to the server.
	conn, err := grpc.Dial(address, grpc.WithInsecure(), grpc.WithBlock())
	if err != nil {
		log.Printf("[ERROR] Did not connect - %v \n", err)
	}
	defer conn.Close()
	c := pb.NewElasticSearchMonitoringClient(conn)

	// Contact the server and print out its response.
	url := clusterURL
	if len(os.Args) > 1 {
		url = os.Args[1]
	}
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()
	r, err := c.GetStatus(ctx, &pb.Cluster{Url: url})
	if err != nil {
		log.Printf("[ERROR] Could not get cluster's status - %v \n", err)
	} else {
	    log.Println("Cluster " + url + " has status: " + r.GetColor())	
	}
}
