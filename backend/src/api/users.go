package api

import (
	"database/sql"
	"encoding/json"
	"github.com/gorilla/mux"
	"github.com/gsistelos/todo-app/models"
	"net/http"
	"strconv"
)

func (s *APIServer) handleCreateUser(w http.ResponseWriter, r *http.Request) error {
	userReq := &models.CreateUserReq{}
	if err := json.NewDecoder(r.Body).Decode(userReq); err != nil {
		return writeJSON(w, http.StatusBadRequest, errJSON(err.Error()))
	}

	if err := userReq.Validate(); err != nil {
		return writeJSON(w, http.StatusBadRequest, errJSON(err.Error()))
	}

	id, err := s.db.CreateUser(userReq)
	if err != nil {
		return writeJSON(w, http.StatusInternalServerError, errJSON(err.Error()))
	}

	idStr := strconv.Itoa(int(id))

	user, err := s.db.GetUser(idStr)
	if err != nil {
		return writeJSON(w, http.StatusNoContent, nil)
	}

	return writeJSON(w, http.StatusCreated, user)
}

func (s *APIServer) handleGetUser(w http.ResponseWriter, r *http.Request) error {
	vars := mux.Vars(r)
	id := vars["id"]

	user, err := s.db.GetUser(id)
	if err != nil {
		if err == sql.ErrNoRows {
			return writeJSON(w, http.StatusNotFound, errJSON("User not found"))
		}
		return writeJSON(w, http.StatusInternalServerError, errJSON(err.Error()))
	}

	return writeJSON(w, http.StatusOK, user)
}

func (s *APIServer) handleGetUsers(w http.ResponseWriter, r *http.Request) error {
	users, err := s.db.GetUsers()
	if err != nil {
		if err == sql.ErrNoRows {
			return writeJSON(w, http.StatusNotFound, errJSON("No users found"))
		}
		return writeJSON(w, http.StatusInternalServerError, errJSON(err.Error()))
	}

	return writeJSON(w, http.StatusOK, users)
}

func (s *APIServer) handleDeleteUser(w http.ResponseWriter, r *http.Request) error {
	vars := mux.Vars(r)
	id := vars["id"]

	err := s.db.DeleteUser(id)
	if err != nil {
		if err == sql.ErrNoRows {
			return writeJSON(w, http.StatusNotFound, errJSON("User not found"))
		}
		return writeJSON(w, http.StatusInternalServerError, errJSON(err.Error()))
	}

	return writeJSON(w, http.StatusNoContent, nil)
}

func (s *APIServer) handleUpdateUser(w http.ResponseWriter, r *http.Request) error {
	vars := mux.Vars(r)
	id := vars["id"]

	userReq := &models.UpdateUserReq{}
	if err := json.NewDecoder(r.Body).Decode(userReq); err != nil {
		return writeJSON(w, http.StatusBadRequest, errJSON(err.Error()))
	}

	if err := userReq.Validate(); err != nil {
		return writeJSON(w, http.StatusBadRequest, errJSON(err.Error()))
	}

	if err := s.db.UpdateUser(id, *userReq); err != nil {
		if err == sql.ErrNoRows {
			return writeJSON(w, http.StatusNotModified, errJSON("User not modified"))
		}
		return writeJSON(w, http.StatusInternalServerError, errJSON(err.Error()))
	}

	user, err := s.db.GetUser(id)
	if err != nil {
		return writeJSON(w, http.StatusNoContent, nil)
	}

	return writeJSON(w, http.StatusOK, user)
}
