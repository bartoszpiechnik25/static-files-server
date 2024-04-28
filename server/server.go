package server

import (
	v1 "github.com/bartoszpiechnik25/static-files-server/server/v1"
	"github.com/gin-gonic/gin"
)

func CreateServer() *gin.Engine {
	router := gin.Default()
	router.ForwardedByClientIP = true
	router.SetTrustedProxies([]string{"127.0.0.1"})
	router.Use(gin.Recovery())
	v1.CreateAssetsV1Group(router)
	return router
}
