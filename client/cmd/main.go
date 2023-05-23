package main

import (
	"context"
	"encoding/json"
	"flag"
	"fmt"
	"net/http"
	"time"

	"guilhermesteves/golang-ports-service/client/pkg/client"
	"guilhermesteves/golang-ports-service/client/pkg/core"
	"guilhermesteves/golang-ports-service/client/pkg/models"
	"guilhermesteves/golang-ports-service/ps"

	"github.com/gorilla/mux"
)

type Server struct {
	client ps.PortServiceClient
}

var (
	host       = flag.String("host", "localhost", "The client host")
	port       = flag.String("port", "8080", "The client port")
	sampleFile = flag.String("file", "./ports.json", "Sample file to import")
)

func (s Server) Start() {
	s.client = client.NewPortServiceClient(fmt.Sprintf("Address: %s:%s", *host, *port))
	router := mux.NewRouter()

	router.HandleFunc("/v1/ports", s.ImportPorts).Methods(http.MethodPost)

	http.ListenAndServe(fmt.Sprintf("%s:%s", *host, *port), router)
}

func (s Server) ImportPorts(w http.ResponseWriter, r *http.Request) {
	if *sampleFile == "" {

	}
	stream := make(chan models.Port)
	//done := make(chan bool, 1)

	var err error
	var counter = 0

	go func() {
		for p := range stream {
			var ctx, cancel = context.WithTimeout(context.Background(), time.Second)
			payload, _ := json.Marshal(p)
			_, err = s.client.InsertOrUpdate(ctx, &ps.Port{ID: p.ID, Content: payload})

			if err != nil {
				close(stream)
				cancel()
				break
			}

			cancel()
			counter++
		}

		//done <- true
	}()

	err = core.StreamPorts(*sampleFile, stream)

	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Header().Set("Content-Type", "text/text")
		w.Write([]byte(fmt.Sprintf("error in %d records", counter)))
	}

	close(stream)
	//<-done
	//close(done)

	w.Write([]byte(fmt.Sprintf("%d ports were inserted or updated", counter)))
}

func main() {
	flag.Parse()

	var server = Server{}
	server.Start()
}
