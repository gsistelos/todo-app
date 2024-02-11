package models

import (
	"fmt"
)

type UserReq struct {
	Username string `json:"username"`
	Email    string `json:"email"`
	Password string `json:"password"`
}

func (s *UserReq) Validate() error {
	if s.Username == "" {
		return fmt.Errorf("username is required")
	}
	if s.Email == "" {
		return fmt.Errorf("email is required")
	}
	if s.Password == "" {
		return fmt.Errorf("password is required")
	}
	return nil
}
