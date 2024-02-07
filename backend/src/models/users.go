package models

import (
	"fmt"
)

type CreateUserReq struct {
	Username string `json:"username"`
	Password string `json:"password"`
	Email    string `json:"email"`
}

type UpdateUserReq struct {
	Username string `json:"username"`
	Password string `json:"password"`
	Email    string `json:"email"`
}

type User struct {
	ID       int    `json:"id"`
	Username string `json:"username"`
	Password string `json:"password"`
	Email    string `json:"email"`
}

func (s *CreateUserReq) Validate() error {
	if s.Username == "" {
		return fmt.Errorf("username is required")
	}
	if s.Password == "" {
		return fmt.Errorf("password is required")
	}
	if s.Email == "" {
		return fmt.Errorf("email is required")
	}
	return nil
}

func (s *UpdateUserReq) Validate() error {
	if s.Username == "" {
		return fmt.Errorf("username is required")
	}
	if s.Password == "" {
		return fmt.Errorf("password is required")
	}
	if s.Email == "" {
		return fmt.Errorf("email is required")
	}
	return nil
}
