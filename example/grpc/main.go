package main

import (
	"grpcexp/wire"
	"log"
)

func main() {
	server, err := wire.InitializeServer("gkit_example")
	if err != nil {
		log.Fatal(err)
	}
	server.Serve()
}
