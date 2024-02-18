package api

import (
	"encoding/json"
	"errors"
	"github.com/gsistelos/todo-app/db"
	"github.com/gsistelos/todo-app/models"
	"io"
	"net/http"
	"strconv"
)

func (s *APIServer) handleCreateTask(w http.ResponseWriter, r *http.Request) error {
	userID := r.PathValue("id")

	user, err := s.db.GetUserByID(userID)
	if err != nil {
		if errors.Is(err, db.NotFound) {
			return writeJSON(w, http.StatusNotFound, apiError{Error: "User not found"})
		} else {
			return err
		}
	}

	tokenString := r.Header.Get("Authorization")
	if tokenString == "" {
		return writeJSON(w, http.StatusUnauthorized, apiError{Error: "Unauthorized"})
	}

	if err := compareJWTCredentials(user.Email, user.Password, tokenString); err != nil {
		return writeJSON(w, http.StatusUnauthorized, apiError{Error: "Unauthorized"})
	}

	taskReq := &models.TaskReq{}
	if err := json.NewDecoder(r.Body).Decode(taskReq); err != nil {
		if errors.Is(err, io.EOF) {
			return writeJSON(w, http.StatusBadRequest, apiError{Error: "Missing request body"})
		} else {
			return writeJSON(w, http.StatusBadRequest, apiError{Error: err.Error()})
		}
	}

	if reqErr := taskReq.Validate(); reqErr != nil {
		return writeJSON(w, http.StatusBadRequest, reqErr)
	}

	userIDInt, _ := strconv.Atoi(userID)

	task := models.NewTask(userIDInt, taskReq.Description)

	id, err := s.db.CreateTask(task)
	if err != nil {
		return err
	}

	task.ID = int(id)

	return writeJSON(w, http.StatusCreated, task)
}

func (s *APIServer) handleGetTasks(w http.ResponseWriter, r *http.Request) error {
	userID := r.PathValue("id")

	user, err := s.db.GetUserByID(userID)
	if err != nil {
		if errors.Is(err, db.NotFound) {
			return writeJSON(w, http.StatusNotFound, apiError{Error: "User not found"})
		} else {
			return err
		}
	}

	tokenString := r.Header.Get("Authorization")
	if tokenString == "" {
		return writeJSON(w, http.StatusUnauthorized, apiError{Error: "Unauthorized"})
	}

	if err := compareJWTCredentials(user.Email, user.Password, tokenString); err != nil {
		return writeJSON(w, http.StatusUnauthorized, apiError{Error: "Unauthorized"})
	}

	tasks, err := s.db.GetTasks(userID)
	if err != nil {
		if errors.Is(err, db.NotFound) {
			return writeJSON(w, http.StatusNotFound, apiError{Error: "No tasks found"})
		} else {
			return err
		}
	}

	return writeJSON(w, http.StatusOK, tasks)
}
