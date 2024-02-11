package api

import (
	"github.com/golang-jwt/jwt/v5"
	"golang.org/x/crypto/bcrypt"
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

func authenticateJWT(email, password, tokenString string) (bool, error) {
	claims := jwt.MapClaims{}
	token, err := jwt.ParseWithClaims(tokenString, claims, func(token *jwt.Token) (interface{}, error) {
		return jwtSecret, nil
	})
	if err != nil {
		return false, err
	}

	if !token.Valid {
		return false, nil
	}

	if claims["email"] != email {
		return false, nil
	}

	if err := bcrypt.CompareHashAndPassword([]byte(password), []byte(claims["password"].(string))); err != nil {
		return false, nil
	}

	return true, nil
}
