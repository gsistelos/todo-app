package api

import (
	"errors"
	"net/http"
	"os"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"github.com/gsistelos/todo-app/db"
	"golang.org/x/crypto/bcrypt"
)

var (
	jwtSecret = []byte(os.Getenv("JWT_SECRET"))
)

type Token struct {
	Token string `json:"token"`
}

func (s *APIServer) jwtHandler(f apiFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		tokenString := r.Header.Get("Authorization")
		if tokenString == "" {
			writeJSON(w, http.StatusUnauthorized, apiError{Error: "Unauthorized"})
			return
		}

		userID := r.PathValue("userID")

		user, err := s.db.GetUserByID(userID)
		if err != nil {
			if errors.Is(err, db.ErrNotFound) {
				writeJSON(w, http.StatusUnauthorized, apiError{Error: "Unauthorized"})
			} else {
				writeJSON(w, http.StatusInternalServerError, apiError{Error: err.Error()})
			}
			return
		}

		if err := compareJWTCredentials(user.Email, user.Password, tokenString); err != nil {
			writeJSON(w, http.StatusUnauthorized, apiError{Error: "Unauthorized"})
			return
		}

		if err := f(w, r); err != nil {
			writeJSON(w, http.StatusInternalServerError, apiError{Error: err.Error()})
		}
	}
}

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
