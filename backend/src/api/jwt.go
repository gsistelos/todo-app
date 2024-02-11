package api

import (
	"github.com/golang-jwt/jwt/v5"
	"os"
	"time"
)

var (
	jwtSecret = []byte(os.Getenv("JWT_SECRET"))
)

func newJWT(email, password string) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256,
		jwt.MapClaims{
			"email":    email,
			"password": password,
			"exp":      time.Now().Add(time.Hour * 24).Unix(),
		})

	tokenString, err := token.SignedString(jwtSecret)
	return tokenString, err
}
