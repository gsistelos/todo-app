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

func (s *APIServer) handleCreateUser(w http.ResponseWriter, r *http.Request) error {
	userReq := &models.UserReq{}
	if err := json.NewDecoder(r.Body).Decode(userReq); err != nil {
		if errors.Is(err, io.EOF) {
			return writeJSON(w, http.StatusBadRequest, apiError{Error: "Missing request body"})
		} else {
			return writeJSON(w, http.StatusBadRequest, apiError{Error: err.Error()})
		}
	}

	if userErr := userReq.Validate(); userErr != nil {
		return writeJSON(w, http.StatusBadRequest, userErr)
	}

	if exists, err := s.db.UserEmailExists(userReq.Email); err != nil {
		return err
	} else if exists {
		return writeJSON(w, http.StatusConflict, models.UserErr{Email: "Email already in use"})
	}

	user, err := models.NewUser(userReq.Username, userReq.Email, userReq.Password)
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

func (s *APIServer) handleGetUserByID(w http.ResponseWriter, r *http.Request) error {
	userID := r.PathValue("userID")

	user, err := s.db.GetUserByID(userID)
	if err != nil {
		if errors.Is(err, db.ErrNotFound) {
			return writeJSON(w, http.StatusNotFound, apiError{Error: "User not found"})
		} else {
			return err
		}
	}

	return writeJSON(w, http.StatusOK, user)
}

func (s *APIServer) handleGetUsers(w http.ResponseWriter, r *http.Request) error {
	users, err := s.db.GetUsers()
	if err != nil {
		if errors.Is(err, db.ErrNotFound) {
			return writeJSON(w, http.StatusNotFound, apiError{Error: "No users found"})
		} else {
			return err
		}
	}

	return writeJSON(w, http.StatusOK, users)
}

func (s *APIServer) handleUpdateUser(w http.ResponseWriter, r *http.Request) error {
	userID := r.PathValue("userID")

	user, err := s.db.GetUserByID(userID)
	if err != nil {
		return err
	} else if user == nil {
		return writeJSON(w, http.StatusNotFound, apiError{Error: "User not found"})
	}

	userReq := &models.UserReq{}
	if err := json.NewDecoder(r.Body).Decode(userReq); err != nil {
		if errors.Is(err, io.EOF) {
			return writeJSON(w, http.StatusBadRequest, apiError{Error: "Missing request body"})
		} else {
			return writeJSON(w, http.StatusBadRequest, apiError{Error: err.Error()})
		}
	}

	// If email has changed, check if it's already in use
	if userReq.Email != user.Email {
		if exists, err := s.db.UserEmailExists(userReq.Email); err != nil {
			return err
		} else if exists {
			return writeJSON(w, http.StatusConflict, models.UserErr{Email: "Email already in use"})
		}
	}

	err = user.Update(userReq.Username, userReq.Email, userReq.Password)
	if err != nil {
		return err
	}

	if err := s.db.UpdateUser(userID, user); err != nil {
		if errors.Is(err, db.ErrNotFound) {
			return writeJSON(w, http.StatusNotFound, apiError{Error: "User not found"})
		} else if errors.Is(err, db.ErrNotModified) {
			return writeJSON(w, http.StatusNotModified, apiError{Error: "User not modified"})
		} else {
			return err
		}
	}

	return writeJSON(w, http.StatusOK, user)
}

func (s *APIServer) handleDeleteUser(w http.ResponseWriter, r *http.Request) error {
	userID := r.PathValue("userID")

	if err := s.db.DeleteUser(userID); err != nil {
		return err
	}

	return writeJSON(w, http.StatusNoContent, nil)
}

func (s *APIServer) handleLogin(w http.ResponseWriter, r *http.Request) error {
	loginReq := &models.LoginReq{}
	if err := json.NewDecoder(r.Body).Decode(loginReq); err != nil {
		if errors.Is(err, io.EOF) {
			return writeJSON(w, http.StatusBadRequest, apiError{Error: "Missing request body"})
		} else {
			return writeJSON(w, http.StatusBadRequest, apiError{Error: err.Error()})
		}
	}

	if loginErr := loginReq.Validate(); loginErr != nil {
		return writeJSON(w, http.StatusBadRequest, loginErr)
	}

	user, err := s.db.GetUserByEmail(loginReq.Email)
	if err != nil {
		if errors.Is(err, db.ErrNotFound) {
			return writeJSON(w, http.StatusUnauthorized, apiError{Error: "Unauthorized"})
		} else {
			return err
		}
	}

	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(loginReq.Password)); err != nil {
		return writeJSON(w, http.StatusUnauthorized, apiError{Error: "Unauthorized"})
	}

	tokenString, err := newJWTSignedString(user)
	if err != nil {
		return err
	}

	return writeJSON(w, http.StatusOK, Token{Token: tokenString})
}
