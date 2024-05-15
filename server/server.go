package server

import (
	"log"

	v1 "github.com/bartoszpiechnik25/static-files-server/server/v1"
	"github.com/bartoszpiechnik25/static-files-server/server/v1/handlers"
	"github.com/gin-gonic/gin"
)

func CreateServer(trustedProxies []string) *gin.Engine {
	router := gin.Default()
	handler, err := handlers.NewHandler()
	if err != nil {
		log.Fatalf("could not start the server due to: %s", err.Error())
	}
	router.ForwardedByClientIP = true
	router.SetTrustedProxies(trustedProxies)
	router.Use(gin.Recovery())
	v1.CreateAssetsV1Group(router, handler)
	return router
}
