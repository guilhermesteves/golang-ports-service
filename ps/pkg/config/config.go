package config

import (
	"log"
	"os"
)

const (
	HOST = "HOST"
	PORT = "PORT"
)

type (
	configuration struct {
		host string
		port string
	}
)

var Variables *configuration

func init() {
	var found bool
	conf := configuration{}

	conf.host, found = os.LookupEnv(HOST)
	if !found {
		log.Fatal("ERROR: env variable HOST is missing")
	}

	conf.port, found = os.LookupEnv(PORT)
	if !found {
		log.Fatal("ERROR: env variable PORT is missing")
	}

	Variables = &conf
}

func (c *configuration) Host() string {
	return c.host
}

func (c *configuration) Port() string {
	return c.port
}
