package models

import (
	"time"
)

type Task struct {
	ID          int       `json:"id"`
	UserID      int       `json:"user_id"`
	Description string    `json:"description"`
	Done        bool      `json:"done"`
	Term        time.Time `json:"term"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
}

type TaskReq struct {
	Description string    `json:"description"`
	Done        bool      `json:"done"`
	Term        time.Time `json:"term"`
}

type TaskErr struct {
	Description string `json:"description,omitempty"`
	Term        string `json:"term,omitempty"`
}

func NewTask(userID int, description string, done bool, term time.Time) *Task {
	now := time.Now().UTC()

	return &Task{
		UserID:      userID,
		Description: description,
		Done:        done,
		Term:        term,
		CreatedAt:   now,
		UpdatedAt:   now,
	}
}

func (s *Task) Update(description string, done bool, term time.Time) {
	s.Description = description
	s.Done = done
	s.Term = term
	s.UpdatedAt = time.Now().UTC()
}

func (s *TaskReq) Validate() *TaskErr {
	taskErr := TaskErr{}

	if s.Description == "" {
		taskErr.Description = "Description is required"
	}

	if s.Term.Before(time.Now().UTC()) {
		taskErr.Term = "Term must be greater than current date"
	}

	if taskErr != (TaskErr{}) {
		return &taskErr
	}

	return nil
}
