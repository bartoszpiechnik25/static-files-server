package v1

import (
	"github.com/bartoszpiechnik25/static-files-server/server/v1/handlers"
	"github.com/gin-gonic/gin"
)

func CreateAssetsV1Group(router *gin.Engine, handler *handlers.Handler) {
	v1 := router.Group("/api/v1")
	v1.OPTIONS("/assets", handler.OptionsHandler)
	v1.Use(AuthMiddleware())
	{
		v1.POST("/assets/:facility-id", handler.CreateAsset)
		v1.GET("/assets/:facility-id")

	}
}
