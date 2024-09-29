package controllers

import (
	"github.com/edgarmueller/go-api-journal/internal/adapters/controllers/dto"
	"github.com/edgarmueller/go-api-journal/internal/app/usecases"
	"github.com/edgarmueller/go-api-journal/internal/domain"
	"github.com/gin-gonic/gin"
)

type UserAPIController struct {
	auth *usecases.AuthUseCases
	user *usecases.UserUseCases
}

func NewUserAPIController(router *gin.RouterGroup, auth *usecases.AuthUseCases, user *usecases.UserUseCases) {
	controller := &UserAPIController{
		auth: auth,
		user: user,
	}
	router.POST("/user/register", controller.RegisterUser)
}

func (u *UserAPIController) RegisterUser(ctx *gin.Context) {
	var registerUser domain.RegisterUser

	if err := ctx.ShouldBindJSON(&registerUser); err != nil {
		ctx.JSON(400, gin.H{"error": err.Error()})
		ctx.Abort()
		return
	}

	user, err := u.user.RegisterUser(registerUser)

	if err != nil {
		ctx.JSON(500, gin.H{"error": err.Error()})
		ctx.Abort()
		return
	}

	ctx.JSON(200, gin.H{"user": dto.ToUserResponse(user)})
}
