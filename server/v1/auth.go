package v1

import (
	"net/http"
	"slices"
	"strings"

	"github.com/gin-gonic/gin"
)

func AuthMiddleware() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		auth, ok := ctx.Request.Header["Authorization"]
		if !ok {
			ctx.JSON(http.StatusUnauthorized, ErrorResponse{
				Err: Error{
					Code:    http.StatusUnauthorized,
					Message: "Unauthorized: Missing authentication token.",
				},
			})
			ctx.Abort()
		}
		auth = strings.Split(auth[0], " ")
		if len(auth) != 2 || !slices.Contains(auth, "Bearer") {
			ctx.JSON(http.StatusUnauthorized, ErrorResponse{
				Err: Error{
					Code:    http.StatusUnauthorized,
					Message: "Unauthorized: Invalid format of authentication token.",
				},
			})
			ctx.Abort()
		}
		if auth[1] != "essasito" {
			ctx.JSON(http.StatusUnauthorized, ErrorResponse{
				Err: Error{
					Code:    http.StatusUnauthorized,
					Message: "Unauthorized: Invalid authentication token.",
				},
			})
			ctx.Abort()
		}
		ctx.Next()
	}
}
