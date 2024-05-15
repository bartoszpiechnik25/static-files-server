package main

import (
	"log"

	"github.com/bartoszpiechnik25/static-files-server/server"
)

func main() {
	router := server.CreateServer([]string{"127.0.0.1"})
	server := "127.0.0.1:6666"
	log.Printf("Starting server on: %s", server)
	router.Run(server)
}
