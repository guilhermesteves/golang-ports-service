package client

import (
	"log"

	"google.golang.org/grpc/credentials/insecure"
	"guilhermesteves/golang-ports-service/ps"

	"google.golang.org/grpc"
)

func NewPortServiceClient(address string) ps.PortServiceClient {
	conn, err := grpc.Dial(address, grpc.WithTransportCredentials(insecure.NewCredentials()))

	if err != nil {
		log.Println(err)
	}

	return ps.NewPortServiceClient(conn)
}
