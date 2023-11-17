package api

import (
	"crypto/rand"
	"encoding/base64"
	"time"

	"github.com/dgrijalva/jwt-go"
	"golang.org/x/crypto/bcrypt"
)

var jwtSecret = []byte(environ.AppConfig.JWTSecret)

type Claims struct {
	Email string `json:"email"`
	jwt.StandardClaims
}

func IsEmpty(data string) bool {
	return len(data) == 0
}

func HashPassword(password string) (string, error) {
	customCost := environ.AppConfig.HashCost
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), customCost)
	if err != nil {
		return "", err
	}
	return string(hashedPassword), nil
}

func VerifyPassword(hashedPassword string, password string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(password))
	return err == nil
}

func GenerateJWT(email string) (string, error) {
	claims := &Claims{
		Email: email,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().Add(time.Hour * 24 * 7).Unix(),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	signedToken, err := token.SignedString(jwtSecret)
	if err != nil {
		return "", err
	}

	return signedToken, nil
}

func GenerateRefreshToken() (string, error) {
	tokenBytes := make([]byte, environ.AppConfig.RefreshTokenLength)
	_, err := rand.Read(tokenBytes)
	if err != nil {
		return "", err
	}

	token := base64.URLEncoding.EncodeToString(tokenBytes)
	return token, nil
}
