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
}

type TaskReq struct {
	Description string `json:"description,omitempty"`
}

func NewTask(userID int, description string) *Task {
	return &Task{
		UserID:      userID,
		Description: description,
		Done:        false,
		CreatedAt:   time.Now(),
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
