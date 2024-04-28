package main

import (
	"context"
	"fmt"
	"log"

	connector "github.com/bartoszpiechnik25/static-files-server/aws-connector"
	"github.com/bartoszpiechnik25/static-files-server/server"
)

func main() {
	log.Println("Creating AWS persistent agent")
	agent, err := connector.NewS3PersistentAgent(context.TODO())
	if err != nil {
		log.Fatal(err)
	}

	if ok, err := agent.BucketExists(context.TODO()); !ok {
		log.Fatal(err)
	}
	if ok, err := agent.DirExists("test-dir"); !ok {
		log.Fatal(err)
	}
	// file, err := os.Open("file.json")
	// if err != nil {
	// 	log.Fatal(err)
	// }
	// if err := agent.UploadFile([]string{"test-dir"}, "file.json", file); err != nil {
	// 	log.Fatal(err)
	// }
	// file.Close()
	fmt.Println(agent.FileExists("test-dir/file.json"))
	router := server.CreateServer()
	server := "127.0.0.1:6666"
	log.Printf("Starting server on: %s", server)
	router.Run(server)
}
