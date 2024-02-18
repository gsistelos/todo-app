package models

import (
	"time"
)

type Task struct {
	ID          int       `json:"id"`
	UserID      int       `json:"user_id"`
	Description string    `json:"description"`
	Done        bool      `json:"done"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
}

type TaskReq struct {
	Description string `json:"description,omitempty"`
}

func NewTask(userID int, description string) *Task {
	now := time.Now().UTC()

	return &Task{
		UserID:      userID,
		Description: description,
		Done:        false,
		CreatedAt:   now,
		UpdatedAt:   now,
	}
}

func (s *TaskReq) Validate() *TaskReq {
	taskErr := TaskReq{}

	if s.Description == "" {
		taskErr.Description = "Description is required"
	}

	if taskErr != (TaskReq{}) {
		return &taskErr
	}

	return nil
}
