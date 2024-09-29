package services

import (
	"fmt"
	"os"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/google/uuid"
)

var jwtKey = []byte(os.Getenv("JWT_KEY"))

type JWTClaim struct {
	Id       string `json:"id"`
	Username string `json:"username"`
	Email    string `json:"email"`
	jwt.StandardClaims
}

func GenerateJWT(userId uuid.UUID, email, username string) (tokenString string, err error) {
	expirationTime := time.Now().Add(1 * time.Hour)
	claims := &JWTClaim{
		Id:       fmt.Sprint(userId),
		Email:    email,
		Username: username,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: expirationTime.Unix(),
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err = token.SignedString([]byte(jwtKey))
	return
}

func VerifyJWT(signedToken string) (claims *JWTClaim, err error) {
	token, err := jwt.ParseWithClaims(signedToken, &JWTClaim{}, func(token *jwt.Token) (interface{}, error) {
		return []byte(jwtKey), nil
	})
	if err != nil {
		return
	}
	claims, ok := token.Claims.(*JWTClaim)
	if !ok || !token.Valid {
		err = jwt.ValidationError{
			Errors: jwt.ValidationErrorMalformed,
		}
		return
	}
	if claims.ExpiresAt < time.Now().Unix() {
		err = jwt.ValidationError{
			Errors: jwt.ValidationErrorExpired,
		}
		return
	}
	return claims, nil
}
