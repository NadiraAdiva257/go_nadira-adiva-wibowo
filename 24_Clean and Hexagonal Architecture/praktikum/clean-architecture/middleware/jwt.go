package middleware

import (
	"time"

	"github.com/golang-jwt/jwt/v4"
)

func CreateToken(userId int, email string, password string) (string, error) {
	claims := jwt.MapClaims{}
	claims["userId"] = userId
	claims["email"] = email
	claims["password"] = password
	claims["exp"] = time.Now().Add(time.Hour * 1).Unix()

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString([]byte(SECRET_JWT))
}
