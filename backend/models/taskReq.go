package models

import (
	"time"
)

type TaskReq struct {
	Description string    `json:"description"`
	Done        bool      `json:"done"`
	Term        time.Time `json:"term"`
}

type TaskErr struct {
	Description string `json:"description,omitempty"`
	Term        string `json:"term,omitempty"`
}

func (s *TaskReq) Validate() *TaskErr {
	taskErr := TaskErr{}

	if s.Description == "" {
		taskErr.Description = "Description is required"
	}

	if s.Term.IsZero() {
		taskErr.Term = "Term is required"
	} else if s.Term.Before(time.Now().UTC()) {
		taskErr.Term = "Term must be greater than current date"
	}

	if taskErr != (TaskErr{}) {
		return &taskErr
	}

	return nil
}
