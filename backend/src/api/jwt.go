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

func newJWTSignedString(email, password string) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256,
		jwt.MapClaims{
			"email":    email,
			"password": password,
			"exp":      time.Now().Add(time.Hour * 24).Unix(),
		})

	tokenString, err := token.SignedString(jwtSecret)
	return tokenString, err
}

func compareJWTCredentials(email, password, tokenString string) error {
	claims := jwt.MapClaims{}
	jwtSecret, err := jwt.ParseWithClaims(tokenString, claims, func(token *jwt.Token) (interface{}, error) {
		return jwtSecret, nil
	})
	if err != nil {
		return err
	} else if !jwtSecret.Valid {
		return jwt.ErrSignatureInvalid
	}

	if claims["email"] != email {
		return jwt.ErrInvalidKey
	}

	if err := bcrypt.CompareHashAndPassword([]byte(password), []byte(claims["password"].(string))); err != nil {
		return err
	}

	if time.Now().Unix() > int64(claims["exp"].(float64)) {
		return jwt.ErrTokenExpired
	}

	return nil
}
