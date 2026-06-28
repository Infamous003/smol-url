package app

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func SetupRoutes(router *gin.Engine) {
	v1 := router.Group("/api/v1")

	v1.GET("", func(ctx *gin.Context) {
		ctx.JSON(http.StatusOK, gin.H{"data": "welcome"})
	})
}
