package controllers

import (
	"github.com/edgarmueller/go-api-journal/internal/app/usecases"
	"github.com/gin-gonic/gin"
)

type TokenRequest struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

func NewTokenAPIController(router *gin.RouterGroup, auth *usecases.AuthUseCases) {
	controller := &TokenAPIController{
		auth: auth,
	}
	router.POST("/token", controller.GenerateToken)
}

type TokenAPIController struct {
	auth *usecases.AuthUseCases
}

func (t *TokenAPIController) GenerateToken(ctx *gin.Context) {
	var request TokenRequest

	if err := ctx.ShouldBindJSON(&request); err != nil {
		ctx.JSON(400, gin.H{"error": err.Error()})
		ctx.Abort()
		return
	}

	token, err := t.auth.GenerateToken(request.Email, request.Password)

	if err != nil {
		ctx.JSON(401, gin.H{"error": err.Error()})
		ctx.Abort()
		return
	}

	ctx.JSON(200, gin.H{"token": token})
}
