package api

import (
	"encoding/json"
	"errors"
	"github.com/gsistelos/todo-app/db"
	"github.com/gsistelos/todo-app/models"
	"golang.org/x/crypto/bcrypt"
	"io"
	"net/http"
	"strconv"
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

	if exists, err := s.db.UserEmailExists(userReq.Email); err != nil {
		return err
	} else if exists {
		return writeJSON(w, http.StatusConflict, apiError{Error: "Email already in use"})
	}

	if err := userReq.HashPassword(); err != nil {
		return err
	}

	id, err := s.db.CreateUser(userReq)
	if err != nil {
		return err
	}

	idStr := strconv.FormatInt(id, 10)

	user, err := s.db.GetUserByID(idStr)
	if err != nil {
		return writeJSON(w, http.StatusCreated, apiError{Error: "Failed to retrieve"})
	}

	return writeJSON(w, http.StatusCreated, user)
}

func (s *APIServer) handleGetUserByID(w http.ResponseWriter, r *http.Request) error {
	id := r.PathValue("id")

	user, err := s.db.GetUserByID(id)
	if err != nil {
		if errors.Is(err, db.NotFound) {
			return writeJSON(w, http.StatusNotFound, apiError{Error: "User not found"})
		} else {
			return err
		}
	}

	tokenString := r.Header.Get("Authorization")
	if tokenString == "" {
		return writeJSON(w, http.StatusOK, user.ToUserInfo())
	}

	authorized, err := authenticateJWT(user.Email, user.Password, tokenString)
	if err != nil {
		return err
	} else if !authorized {
		return writeJSON(w, http.StatusUnauthorized, apiError{Error: "Unauthorized"})
	}

	return writeJSON(w, http.StatusOK, user)
}

func (s *APIServer) handleGetUsers(w http.ResponseWriter, r *http.Request) error {
	users, err := s.db.GetUsers()
	if err != nil {
		if errors.Is(err, db.NotFound) {
			return writeJSON(w, http.StatusNotFound, apiError{Error: "No users found"})
		} else {
			return err
		}
	}

	return writeJSON(w, http.StatusOK, users)
}

func (s *APIServer) handleDeleteUser(w http.ResponseWriter, r *http.Request) error {
	tokenString := r.Header.Get("Authorization")
	if tokenString == "" {
		return writeJSON(w, http.StatusUnauthorized, apiError{Error: "Unauthorized"})
	}

	id := r.PathValue("id")

	user, err := s.db.GetUserByID(id)
	if err != nil {
		if errors.Is(err, db.NotFound) {
			return writeJSON(w, http.StatusNotFound, apiError{Error: "User not found"})
		} else {
			return err
		}
	}

	authorized, err := authenticateJWT(user.Email, user.Password, tokenString)
	if err != nil {
		return err
	} else if !authorized {
		return writeJSON(w, http.StatusUnauthorized, apiError{Error: "Unauthorized"})
	}

	err = s.db.DeleteUser(id)
	if err != nil {
		return err
	}

	return writeJSON(w, http.StatusNoContent, nil)
}

func (s *APIServer) handleUpdateUser(w http.ResponseWriter, r *http.Request) error {
	tokenString := r.Header.Get("Authorization")
	if tokenString == "" {
		return writeJSON(w, http.StatusUnauthorized, apiError{Error: "Unauthorized"})
	}

	id := r.PathValue("id")

	user, err := s.db.GetUserByID(id)
	if err != nil {
		return err
	} else if user == nil {
		return writeJSON(w, http.StatusNotFound, apiError{Error: "User not found"})
	}

	authorized, err := authenticateJWT(user.Email, user.Password, tokenString)
	if err != nil {
		return err
	} else if !authorized {
		return writeJSON(w, http.StatusUnauthorized, apiError{Error: "Unauthorized"})
	}

	userReq := &models.UserReq{}
	if err := json.NewDecoder(r.Body).Decode(userReq); err != nil {
		if errors.Is(err, io.EOF) {
			return writeJSON(w, http.StatusBadRequest, apiError{Error: "Missing request body"})
		} else {
			return writeJSON(w, http.StatusBadRequest, apiError{Error: err.Error()})
		}
	}

	if userReq.Email != user.Email {
		if exists, err := s.db.UserEmailExists(userReq.Email); err != nil {
			return err
		} else if exists {
			return writeJSON(w, http.StatusConflict, apiError{Error: "Email already in use"})
		}
	}

	if err := userReq.HashPassword(); err != nil {
		return err
	}

	if err := s.db.UpdateUser(id, userReq); err != nil {
		if errors.Is(err, db.NotFound) {
			return writeJSON(w, http.StatusNotFound, apiError{Error: "User not found"})
		} else if errors.Is(err, db.NotModified) {
			return writeJSON(w, http.StatusNotModified, apiError{Error: "User not modified"})
		} else {
			return err
		}
	}

	user, err = s.db.GetUserByID(id)
	if err != nil {
		return writeJSON(w, http.StatusOK, apiError{Error: "Failed to retrieve"})
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
			return err
		}
	}

	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(loginReq.Password)); err != nil {
		return writeJSON(w, http.StatusUnauthorized, apiError{Error: "Unauthorized"})
	}

	tokenString, err := newJWT(loginReq.Email, loginReq.Password)
	if err != nil {
		return err
	}

	return writeJSON(w, http.StatusOK, map[string]string{"token": tokenString})
}
