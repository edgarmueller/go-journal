package presentation

import (
	"net/http"
	"os"

	presentation "github.com/edgarmueller/go-api-journal/internal/adapters/controllers/presentation/templates"
	"github.com/edgarmueller/go-api-journal/internal/app/usecases"
	"github.com/edgarmueller/go-api-journal/internal/domain"
	"github.com/gin-gonic/gin"
)

type AuthController struct {
	auth *usecases.AuthUseCases
	user *usecases.UserUseCases
}

func NewAuthController(router *gin.Engine, auth *usecases.AuthUseCases, user *usecases.UserUseCases) {
	controller := &AuthController{
		auth: auth,
		user: user,
	}
	router.GET("/", func(ctx *gin.Context) {
		_, exists := ctx.Get("UserUUID")

		if exists {
			ctx.Redirect(http.StatusFound, "/login")
		} else {
			ctx.Redirect(http.StatusFound, "/journal")
		}
	})
	router.GET("/login", controller.ShowLogin)
	router.GET("/register", controller.ShowRegister)
	router.POST("/login", controller.Login)
	router.POST("/register", controller.Register)
	router.POST("/logout", controller.Logout)
}

func (c *AuthController) ShowLogin(ctx *gin.Context) {
	r := New(ctx.Request.Context(), http.StatusOK, presentation.Login())
	ctx.Render(http.StatusOK, r)
}

func (c *AuthController) ShowRegister(ctx *gin.Context) {
	r := New(ctx.Request.Context(), http.StatusOK, presentation.Register())
	ctx.Render(http.StatusOK, r)
}

func (c *AuthController) Login(ctx *gin.Context) {
	username := ctx.PostForm("username")
	password := ctx.PostForm("password")

	token, err := c.auth.GenerateToken(username, password)

	if err != nil {
		r := New(ctx.Request.Context(), http.StatusOK, presentation.Error("500", err.Error()))
		ctx.Render(http.StatusOK, r)
		return
	}

	ctx.SetCookie("token", token, 3600, "/", os.Getenv("DOMAIN"), true, true)
	ctx.Redirect(http.StatusFound, "/journal")
}

func (c *AuthController) Register(ctx *gin.Context) {
	username := ctx.PostForm("username")
	password := ctx.PostForm("password")
	email := ctx.PostForm("email")

	_, err := c.user.RegisterUser(domain.RegisterUser{
		Username: username,
		Email:    email,
		Password: password,
	})

	if err != nil {
		r := New(ctx.Request.Context(), http.StatusOK, presentation.Error("500", err.Error()))
		ctx.Render(http.StatusOK, r)
		return
	}

	ctx.Redirect(http.StatusFound, "/login")
}

func (c *AuthController) Logout(ctx *gin.Context) {
	ctx.SetCookie("token", "", -1, "/", os.Getenv("DOMAIN"), true, true)
	ctx.Redirect(http.StatusFound, "/login")
}
