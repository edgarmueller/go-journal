package controllers

import (
	"net/http"

	"github.com/edgarmueller/go-api-journal/internal/adapters/middlewares"
	"github.com/gin-gonic/gin"
)

func NewPingController(router *gin.RouterGroup) {
	router.GET("ping", middlewares.Auth(), Ping)
}

func Ping(context *gin.Context) {
	context.JSON(http.StatusOK, gin.H{"message": "pong"})
}
