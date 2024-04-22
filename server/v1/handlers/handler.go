package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func OptionsHandler(ctx *gin.Context) {
	ctx.Header("Access-Controll-Allow-Headers", "Content-Type, Authorization")
	ctx.Header("Access-Controll-Allow-Methods", "POST, GET, PUT, OPTIONS")
	ctx.Header("Accept", "application/json,multipart/form-data")

	ctx.AbortWithStatus(http.StatusNoContent)
}
