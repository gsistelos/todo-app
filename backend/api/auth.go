package api

import (
	"encoding/json"
	"errors"
	"io"
	"net/http"

	"github.com/gsistelos/todo-app/db"
	"github.com/gsistelos/todo-app/models"
	"golang.org/x/crypto/bcrypt"
)

func (s *APIServer) handleLogin(w http.ResponseWriter, r *http.Request) error {
	loginReq := &models.LoginReq{}
	if err := json.NewDecoder(r.Body).Decode(loginReq); err != nil {
		if errors.Is(err, io.EOF) {
			return writeJSON(w, http.StatusBadRequest, apiError{Message: "Missing request body"})
		} else {
			return writeJSON(w, http.StatusBadRequest, apiError{Message: err.Error()})
		}
	}

	if loginErr := loginReq.Validate(); loginErr != nil {
		return writeJSON(w, http.StatusBadRequest, loginErr)
	}

	user, err := s.db.GetUserByEmail(loginReq.Email)
	if err != nil {
		if errors.Is(err, db.ErrNotFound) {
			return writeJSON(w, http.StatusUnauthorized, apiError{Message: "Unauthorized"})
		} else {
			return err
		}
	}

	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(loginReq.Password)); err != nil {
		return writeJSON(w, http.StatusUnauthorized, apiError{Message: "Unauthorized"})
	}

	tokenString, err := newJWTSignedString(user)
	if err != nil {
		return err
	}

	return writeJSON(w, http.StatusOK, Token{Token: tokenString})
}

func (s *APIServer) handleRegister(w http.ResponseWriter, r *http.Request) error {
	registerReq := &models.RegisterReq{}
	if err := json.NewDecoder(r.Body).Decode(registerReq); err != nil {
		if errors.Is(err, io.EOF) {
			return writeJSON(w, http.StatusBadRequest, apiError{Message: "Missing request body"})
		} else {
			return writeJSON(w, http.StatusBadRequest, apiError{Message: err.Error()})
		}
	}

	if registerErr := registerReq.Validate(); registerErr != nil {
		return writeJSON(w, http.StatusBadRequest, registerErr)
	}

	if exists, err := s.db.UserEmailExists(registerReq.Email); err != nil {
		return err
	} else if exists {
		return writeJSON(w, http.StatusConflict, models.UserErr{Email: "Email already in use"})
	}

	user, err := models.NewUser(registerReq.Username, registerReq.Email, registerReq.Password)
	if err != nil {
		return err
	}

	id, err := s.db.CreateUser(user)
	if err != nil {
		return err
	}

	user.ID = int(id)

	return writeJSON(w, http.StatusCreated, user)
}
