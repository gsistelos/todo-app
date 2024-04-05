package api

import (
	"encoding/json"
	"errors"
	"io"
	"net/http"
	"strconv"

	"github.com/gsistelos/todo-app/db"
	"github.com/gsistelos/todo-app/models"
)

func (s *APIServer) handleCreateTask(w http.ResponseWriter, r *http.Request) error {
	taskReq := &models.TaskReq{}
	if err := json.NewDecoder(r.Body).Decode(taskReq); err != nil {
		if errors.Is(err, io.EOF) {
			return writeJSON(w, http.StatusBadRequest, apiError{Message: "Missing request body"})
		} else {
			return writeJSON(w, http.StatusBadRequest, apiError{Message: err.Error()})
		}
	}

	if taskErr := taskReq.Validate(); taskErr != nil {
		return writeJSON(w, http.StatusBadRequest, taskErr)
	}

	userID := r.PathValue("userID")

	userIDInt, _ := strconv.Atoi(userID)

	task := models.NewTask(userIDInt, taskReq.Description, taskReq.Done, taskReq.Term)

	id, err := s.db.CreateTask(task)
	if err != nil {
		return err
	}

	task.ID = int(id)

	return writeJSON(w, http.StatusCreated, task)
}

func (s *APIServer) handleGetTasks(w http.ResponseWriter, r *http.Request) error {
	userID := r.PathValue("userID")

	tasks, err := s.db.GetTasks(userID)
	if err != nil {
		if errors.Is(err, db.ErrNotFound) {
			return writeJSON(w, http.StatusNotFound, apiError{Message: "No tasks found"})
		} else {
			return err
		}
	}

	return writeJSON(w, http.StatusOK, tasks)
}

func (s *APIServer) handleGetTaskByID(w http.ResponseWriter, r *http.Request) error {
	userID := r.PathValue("userID")
	taskID := r.PathValue("taskID")

	task, err := s.db.GetTaskByID(userID, taskID)
	if err != nil {
		if errors.Is(err, db.ErrNotFound) {
			return writeJSON(w, http.StatusNotFound, apiError{Message: "Task not found"})
		} else {
			return err
		}
	}

	return writeJSON(w, http.StatusOK, task)
}

func (s *APIServer) handleUpdateTask(w http.ResponseWriter, r *http.Request) error {
	taskReq := &models.TaskReq{}
	if err := json.NewDecoder(r.Body).Decode(taskReq); err != nil {
		if errors.Is(err, io.EOF) {
			return writeJSON(w, http.StatusBadRequest, apiError{Message: "Missing request body"})
		} else {
			return writeJSON(w, http.StatusBadRequest, apiError{Message: err.Error()})
		}
	}

	if taskErr := taskReq.Validate(); taskErr != nil {
		return writeJSON(w, http.StatusBadRequest, taskErr)
	}

	userID := r.PathValue("userID")
	taskID := r.PathValue("taskID")

	task, err := s.db.GetTaskByID(userID, taskID)
	if err != nil {
		if errors.Is(err, db.ErrNotFound) {
			return writeJSON(w, http.StatusNotFound, apiError{Message: "Task not found"})
		} else {
			return err
		}
	}

	task.Update(taskReq.Description, taskReq.Done, taskReq.Term)

	if err := s.db.UpdateTask(userID, taskID, task); err != nil {
		if errors.Is(err, db.ErrNotFound) {
			return writeJSON(w, http.StatusNotFound, apiError{Message: "Task not found"})
		} else {
			return err
		}
	}

	return writeJSON(w, http.StatusOK, task)
}

func (s *APIServer) handleDeleteTask(w http.ResponseWriter, r *http.Request) error {
	userID := r.PathValue("userID")
	taskID := r.PathValue("taskID")

	if err := s.db.DeleteTask(userID, taskID); err != nil {
		if errors.Is(err, db.ErrNotFound) {
			return writeJSON(w, http.StatusNotFound, apiError{Message: "Task not found"})
		} else {
			return err
		}
	}

	return writeJSON(w, http.StatusNoContent, nil)
}
