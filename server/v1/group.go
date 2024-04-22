package v1

import (
	"github.com/bartoszpiechnik25/static-files-server/server/v1/handlers"
	"github.com/gin-gonic/gin"
)

func CreateAssetsV1Group(router *gin.Engine) {
	v1 := router.Group("/v1")
	v1.OPTIONS("/assets", handlers.OptionsHandler)
	v1.Use(AuthMiddleware())
	{
		v1.POST("/assets/:facility-id")
		v1.GET("/assets/:facility-id")
	}
}
