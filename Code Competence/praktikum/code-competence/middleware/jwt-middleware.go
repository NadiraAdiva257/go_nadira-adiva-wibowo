package middleware

import (
	"github.com/golang-jwt/jwt/v5"
	"github.com/labstack/echo/v4"
)

type JwtCustomClaims struct {
	Id int `json:"id"`
	jwt.RegisteredClaims
}

func CreateToken(id int, username, password string) (string, error) {
	claims := jwt.MapClaims{}
	claims["id"] = id
	claims["username"] = username
	claims["password"] = password

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	return token.SignedString([]byte("secret"))
}

func GetClaims(c echo.Context) *JwtCustomClaims {
	user := c.Get("user").(*jwt.Token)
	claims := user.Claims.(*JwtCustomClaims)
	return claims
}
