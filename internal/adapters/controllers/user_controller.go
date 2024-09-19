package controllers

import (
	"github.com/edgarmueller/go-api-journal/internal/app/usecases"
	"github.com/edgarmueller/go-api-journal/internal/domain"
	"github.com/gin-gonic/gin"
)

type UserController struct {
	auth *usecases.AuthUseCases
}

func NewUserController(router *gin.RouterGroup, auth *usecases.AuthUseCases) {
	controller := &UserController{
		auth: auth,
	}
	router.POST("/user/register", controller.RegisterUser)
}

func (u *UserController) RegisterUser(ctx *gin.Context) {
	var registerUser domain.RegisterUser

	if err := ctx.ShouldBindJSON(&registerUser); err != nil {
		ctx.JSON(400, gin.H{"error": err.Error()})
		ctx.Abort()
		return
	}

	user, err := u.auth.RegisterUser(registerUser)

	if err != nil {
		ctx.JSON(500, gin.H{"error": err.Error()})
		ctx.Abort()
		return
	}

	ctx.JSON(200, gin.H{"user": user})
}
