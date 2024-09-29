package middlewares

import (
	"github.com/edgarmueller/go-api-journal/internal/app/services"
	"github.com/gin-gonic/gin"
)

func Auth(shouldRedirectToLogin bool) gin.HandlerFunc {
	return func(context *gin.Context) {
		tokenString := context.GetHeader("Authorization")

		if tokenString == "" {
			tokenString, _ = context.Cookie("token")
		}

		if tokenString == "" {
			if shouldRedirectToLogin {
				context.Redirect(302, "/login")
				context.Abort()
				return
			}
			context.JSON(401, gin.H{"error": "request does not contain an access token"})
			context.Abort()
			return
		}

		claims, err := services.VerifyJWT(tokenString)
		if err != nil {
			context.JSON(401, gin.H{"error": err.Error()})
			context.Abort()
			return
		}

		context.Set("UserUUID", claims.Id)

		context.Next()
	}
}
