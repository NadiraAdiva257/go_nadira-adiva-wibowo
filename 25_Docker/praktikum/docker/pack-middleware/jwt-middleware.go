package packmiddleware

import (
	"docker/constants"

	"github.com/dgrijalva/jwt-go"
)

func CreateToken(email string, password string) (string, error) {
	claims := jwt.MapClaims{}
	claims["email"] = email
	claims["password"] = password

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	return token.SignedString([]byte(constants.SECRET_JWT))
}
