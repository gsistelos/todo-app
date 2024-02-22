package api

import (
	"net/http"
	"os"
	"time"

	"github.com/golang-jwt/jwt/v5"
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

		if err := authenticateJWT(tokenString); err != nil {
			writeJSON(w, http.StatusUnauthorized, apiError{Error: "Unauthorized"})
			return
		}

		if err := f(w, r); err != nil {
			writeJSON(w, http.StatusInternalServerError, apiError{Error: err.Error()})
		}
	}
}

func newJWTSignedString(id int, username string) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256,
		jwt.MapClaims{
			"id":        id,
			"username":  username,
			"expiresAt": time.Now().Add(time.Hour * 24).Unix(),
		})

	tokenString, err := token.SignedString(jwtSecret)
	return tokenString, err
}

func authenticateJWT(tokenString string) error {
	claims := jwt.MapClaims{}
	token, err := jwt.ParseWithClaims(tokenString, claims, func(token *jwt.Token) (interface{}, error) {
		if token.Method != jwt.SigningMethodHS256 {
			return nil, jwt.ErrInvalidKey
		}
		return jwtSecret, nil
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
