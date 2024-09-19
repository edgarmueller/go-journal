package services

import (
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/edgarmueller/go-api-journal/internal/domain/repositories"
)

// TODO[]
var jwtKey = []byte("supersecretkey")

type TokenService struct {
	userRepository repositories.UserRepository
	authService    AuthService
}

type JWTClaim struct {
	Username string `json:"username"`
	Email    string `json:"email"`
	jwt.StandardClaims
}

// func NewTokenService(
// 	userRepository repositories.UserRepository,
// 	authService AuthService,
// ) *TokenService {
// 	return &TokenService{userRepository: userRepository, authService: authService}
// }

func GenerateJWT(email, username string) (tokenString string, err error) {
	expirationTime := time.Now().Add(1 * time.Hour)
	claims := &JWTClaim{
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

func VerifyJWT(signedToken string) (err error) {
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
	return
}
