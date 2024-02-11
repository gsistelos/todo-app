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

	if err := userReq.Validate(); err != nil {
		return writeJSON(w, http.StatusBadRequest, apiError{Error: err.Error()})
	}

	if _, err := s.db.GetUserByEmail(userReq.Email); err == nil {
		return writeJSON(w, http.StatusConflict, apiError{Error: "Email already in use"})
	} else if !errors.Is(err, db.NotFound) {
		return writeJSON(w, http.StatusInternalServerError, apiError{Error: err.Error()})
	}

	hashedPassword, err := hashPassword(userReq.Password)
	if err != nil {
		return writeJSON(w, http.StatusInternalServerError, apiError{Error: err.Error()})
	}

	userReq.Password = hashedPassword

	user := models.NewUser(userReq.Username, userReq.Email, userReq.Password)
	id, err := s.db.CreateUser(user)
	if err != nil {
		return writeJSON(w, http.StatusInternalServerError, apiError{Error: err.Error()})
	}

	user.ID = int(id)

	return writeJSON(w, http.StatusCreated, user)
}

func (s *APIServer) handleGetUserByID(w http.ResponseWriter, r *http.Request) error {
	id := r.PathValue("id")

	user, err := s.db.GetUserByID(id)
	if err != nil {
		if errors.Is(err, db.NotFound) {
			return writeJSON(w, http.StatusNotFound, apiError{Error: "User not found"})
		} else {
			return writeJSON(w, http.StatusInternalServerError, apiError{Error: err.Error()})
		}
	}

	return writeJSON(w, http.StatusOK, user)
}

func (s *APIServer) handleGetUsers(w http.ResponseWriter, r *http.Request) error {
	users, err := s.db.GetUsers()
	if err != nil {
		if errors.Is(err, db.NotFound) {
			return writeJSON(w, http.StatusNotFound, apiError{Error: "No users found"})
		} else {
			return writeJSON(w, http.StatusInternalServerError, apiError{Error: err.Error()})
		}
	}

	return writeJSON(w, http.StatusOK, users)
}

func (s *APIServer) handleDeleteUser(w http.ResponseWriter, r *http.Request) error {
	id := r.PathValue("id")

	err := s.db.DeleteUser(id)
	if err != nil {
		if errors.Is(err, db.NotFound) {
			return writeJSON(w, http.StatusNotFound, apiError{Error: "User not found"})
		} else {
			return writeJSON(w, http.StatusInternalServerError, apiError{Error: err.Error()})
		}
	}

	return writeJSON(w, http.StatusNoContent, nil)
}

func (s *APIServer) handleUpdateUser(w http.ResponseWriter, r *http.Request) error {
	userReq := &models.UserReq{}
	if err := json.NewDecoder(r.Body).Decode(userReq); err != nil {
		if errors.Is(err, io.EOF) {
			return writeJSON(w, http.StatusBadRequest, apiError{Error: "Missing request body"})
		} else {
			return writeJSON(w, http.StatusBadRequest, apiError{Error: err.Error()})
		}
	}

	if err := userReq.Validate(); err != nil {
		return writeJSON(w, http.StatusBadRequest, apiError{Error: err.Error()})
	}

	if _, err := s.db.GetUserByEmail(userReq.Email); err == nil {
		return writeJSON(w, http.StatusConflict, apiError{Error: "Email already in use"})
	} else if !errors.Is(err, db.NotFound) {
		return writeJSON(w, http.StatusInternalServerError, apiError{Error: err.Error()})
	}

	hashedPassword, err := hashPassword(userReq.Password)
	if err != nil {
		return writeJSON(w, http.StatusInternalServerError, apiError{Error: err.Error()})
	}

	userReq.Password = hashedPassword

	id := r.PathValue("id")

	if err := s.db.UpdateUser(id, *userReq); err != nil {
		if errors.Is(err, db.NotFound) {
			return writeJSON(w, http.StatusNotFound, apiError{Error: "User not found"})
		} else if errors.Is(err, db.NotModified) {
			return writeJSON(w, http.StatusNotModified, apiError{Error: "User not modified"})
		} else {
			return writeJSON(w, http.StatusInternalServerError, apiError{Error: err.Error()})
		}
	}

	user, err := s.db.GetUserByID(id)
	if err != nil {
		return writeJSON(w, http.StatusNoContent, nil)
	}

	return writeJSON(w, http.StatusOK, user)
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

	if err := loginReq.Validate(); err != nil {
		return writeJSON(w, http.StatusBadRequest, apiError{Error: err.Error()})
	}

	user, err := s.db.GetUserByEmail(loginReq.Email)
	if err != nil {
		if errors.Is(err, db.NotFound) {
			return writeJSON(w, http.StatusUnauthorized, apiError{Error: "Unauthorized"})
		} else {
			return writeJSON(w, http.StatusInternalServerError, apiError{Error: err.Error()})
		}
	}

	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(loginReq.Password)); err != nil {
		return writeJSON(w, http.StatusUnauthorized, apiError{Error: "Unauthorized"})
	}

	token, err := newJWT(user.Email, user.Password)
	if err != nil {
		return writeJSON(w, http.StatusInternalServerError, apiError{Error: err.Error()})
	}

	return writeJSON(w, http.StatusOK, map[string]string{"token": token})
}
