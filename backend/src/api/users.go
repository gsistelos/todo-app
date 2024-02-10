package api

import (
	"encoding/json"
	"errors"
	"io"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/gsistelos/todo-app/db"
	"github.com/gsistelos/todo-app/models"
)

func (s *APIServer) handleCreateUser(w http.ResponseWriter, r *http.Request) error {
	userReq := &models.CreateUserReq{}
	if err := json.NewDecoder(r.Body).Decode(userReq); err != nil {
		if errors.Is(err, io.EOF) {
			return writeJSON(w, http.StatusBadRequest, errJSON("Missing request body"))
		} else {
			return writeJSON(w, http.StatusBadRequest, errJSON(err.Error()))
		}
	}

	if err := userReq.Validate(); err != nil {
		return writeJSON(w, http.StatusBadRequest, errJSON(err.Error()))
	}

	if _, err := s.db.GetUserByEmail(userReq.Email); err == nil {
		return writeJSON(w, http.StatusConflict, errJSON("Email already in use"))
	} else if !errors.Is(err, db.NotFound) {
		return writeJSON(w, http.StatusInternalServerError, errJSON(err.Error()))
	}

	hashedPassword, err := hashPassword(userReq.Password)
	if err != nil {
		return writeJSON(w, http.StatusInternalServerError, errJSON(err.Error()))
	}

	userReq.Password = hashedPassword

	user := models.NewUser(userReq.Username, userReq.Password, userReq.Email)
	id, err := s.db.CreateUser(user)
	if err != nil {
		return writeJSON(w, http.StatusInternalServerError, errJSON(err.Error()))
	}

	user.ID = int(id)

	return writeJSON(w, http.StatusCreated, user)
}

func (s *APIServer) handleGetUser(w http.ResponseWriter, r *http.Request) error {
	vars := mux.Vars(r)
	id := vars["id"]

	user, err := s.db.GetUser(id)
	if err != nil {
		if errors.Is(err, db.NotFound) {
			return writeJSON(w, http.StatusNotFound, errJSON("User not found"))
		} else {
			return writeJSON(w, http.StatusInternalServerError, errJSON(err.Error()))
		}
	}

	return writeJSON(w, http.StatusOK, user)
}

func (s *APIServer) handleGetUsers(w http.ResponseWriter, r *http.Request) error {
	users, err := s.db.GetUsers()
	if err != nil {
		if errors.Is(err, db.NotFound) {
			return writeJSON(w, http.StatusNotFound, errJSON("No users found"))
		} else {
			return writeJSON(w, http.StatusInternalServerError, errJSON(err.Error()))
		}
	}

	return writeJSON(w, http.StatusOK, users)
}

func (s *APIServer) handleDeleteUser(w http.ResponseWriter, r *http.Request) error {
	vars := mux.Vars(r)
	id := vars["id"]

	err := s.db.DeleteUser(id)
	if err != nil {
		if errors.Is(err, db.NotFound) {
			return writeJSON(w, http.StatusNotFound, errJSON("User not found"))
		} else {
			return writeJSON(w, http.StatusInternalServerError, errJSON(err.Error()))
		}
	}

	return writeJSON(w, http.StatusNoContent, nil)
}

func (s *APIServer) handleUpdateUser(w http.ResponseWriter, r *http.Request) error {
	userReq := &models.UpdateUserReq{}
	if err := json.NewDecoder(r.Body).Decode(userReq); err != nil {
		if errors.Is(err, io.EOF) {
			return writeJSON(w, http.StatusBadRequest, errJSON("Missing request body"))
		} else {
			return writeJSON(w, http.StatusBadRequest, errJSON(err.Error()))
		}
	}

	if err := userReq.Validate(); err != nil {
		return writeJSON(w, http.StatusBadRequest, errJSON(err.Error()))
	}

	if _, err := s.db.GetUserByEmail(userReq.Email); err == nil {
		return writeJSON(w, http.StatusConflict, errJSON("Email already in use"))
	} else if !errors.Is(err, db.NotFound) {
		return writeJSON(w, http.StatusInternalServerError, errJSON(err.Error()))
	}

	hashedPassword, err := hashPassword(userReq.Password)
	if err != nil {
		return writeJSON(w, http.StatusInternalServerError, errJSON(err.Error()))
	}

	userReq.Password = hashedPassword

	vars := mux.Vars(r)
	id := vars["id"]

	if err := s.db.UpdateUser(id, *userReq); err != nil {
		if errors.Is(err, db.NotFound) {
			return writeJSON(w, http.StatusNotFound, errJSON("User not found"))
		} else if errors.Is(err, db.NotModified) {
			return writeJSON(w, http.StatusNotModified, errJSON("User not modified"))
		} else {
			return writeJSON(w, http.StatusInternalServerError, errJSON(err.Error()))
		}
	}

	user, err := s.db.GetUser(id)
	if err != nil {
		return writeJSON(w, http.StatusNoContent, nil)
	}

	return writeJSON(w, http.StatusOK, user)
}
