package server

import (
	v1 "github.com/bartoszpiechnik25/static-files-server/server/v1"
	"github.com/gin-gonic/gin"
)

func CreateServer() *gin.Engine {
	router := gin.Default()
	router.Use(gin.Recovery())
	v1.CreateAssetsV1Group(router)
	return router
}
