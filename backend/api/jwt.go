package api

import (
	"errors"
	"net/http"
	"os"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"github.com/gsistelos/todo-app/db"
	"github.com/gsistelos/todo-app/models"
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

		if len(tokenString) < 7 || tokenString[:7] != "Bearer " {
			writeJSON(w, http.StatusUnauthorized, apiError{Error: "Unauthorized"})
			return
		}

		tokenString = tokenString[7:]

		userID := r.PathValue("userID")

		user, err := s.db.GetUserByID(userID)
		if err != nil {
			if errors.Is(err, db.ErrNotFound) {
				writeJSON(w, http.StatusNotFound, apiError{Error: "User not found"})
			} else {
				writeJSON(w, http.StatusInternalServerError, apiError{Error: err.Error()})
			}
			return
		}

		if err := authenticateJWT(user, tokenString); err != nil {
			writeJSON(w, http.StatusUnauthorized, apiError{Error: "Unauthorized"})
			return
		}

		if err := f(w, r); err != nil {
			writeJSON(w, http.StatusInternalServerError, apiError{Error: err.Error()})
		}
	}
}

func newJWTSignedString(user *models.User) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256,
		jwt.MapClaims{
			"id":        user.ID,
			"username":  user.Username,
			"expiresAt": time.Now().Add(time.Hour * 24).Unix(),
		})

	signature := append(jwtSecret, []byte(user.Email)...)
	signature = append(signature, []byte(user.Password)...)

	tokenString, err := token.SignedString(signature)
	return tokenString, err
}

func authenticateJWT(user *models.User, tokenString string) error {
	claims := jwt.MapClaims{}
	token, err := jwt.ParseWithClaims(tokenString, claims, func(token *jwt.Token) (interface{}, error) {
		if token.Method != jwt.SigningMethodHS256 {
			return nil, jwt.ErrInvalidKey
		}

		signature := append(jwtSecret, []byte(user.Email)...)
		signature = append(signature, []byte(user.Password)...)

		return signature, nil
	})
	if err != nil {
		return err
	} else if !token.Valid {
		return jwt.ErrSignatureInvalid
	}

	if time.Now().Unix() > int64(claims["expiresAt"].(float64)) {
		return jwt.ErrTokenExpired
	}

	return nil
}
