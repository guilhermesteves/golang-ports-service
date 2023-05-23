package main

import (
	"context"
	"fmt"
	"log"
	"net"
	"os"
	"os/signal"
	"syscall"

	"guilhermesteves/golang-ports-service/ps"
	"guilhermesteves/golang-ports-service/ps/pkg/config"
	"guilhermesteves/golang-ports-service/ps/pkg/data"

	"google.golang.org/grpc"
)

type (
	db struct {
		memory data.IDatabase
	}
)

func NewPortService() ps.PortServiceServer {
	return &db{memory: data.New()}
}

func (p *db) InsertOrUpdate(ctx context.Context, req *ps.Port) (*ps.Empty, error) {
	return nil, p.memory.Create(req.ID, req.Content)
}

func main() {
	list, err := net.Listen("tcp", fmt.Sprintf("%s:%s", config.Variables.Host(), config.Variables.Port()))
	if err != nil {
		log.Fatal(err)
	}
	server := grpc.NewServer()
	srv := NewPortService()
	ps.RegisterPortServiceServer(server, srv)

	log.Printf("server listening at %v", list.Addr())
	go func() {
		stop := make(chan os.Signal, 1)
		signal.Notify(stop, os.Interrupt, syscall.SIGTERM)
		log.Printf("service is shutting down after signal %v\n", <-stop)
		server.GracefulStop()
	}()
	err = server.Serve(list)
	if err != nil {
		log.Fatal(err)
	}
}
