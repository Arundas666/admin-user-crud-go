package jwt

import (
	"time"

	"github.com/dgrijalva/jwt-go"
)

var MysigninKey = []byte("Secret-Key")

func GenerateJWT() (string, error) {
	token := jwt.New(jwt.SigningMethodHS256)

	claims := token.Claims.(jwt.MapClaims)
	claims["authorized"] = true
	claims["user"] = "Arun"
	claims["exp"] = time.Now().Add(time.Minute * 30).Unix()

	tokenString, err := token.SignedString(MysigninKey)
	return tokenString, err
}




