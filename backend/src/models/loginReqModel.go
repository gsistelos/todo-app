package models

import (
	"fmt"
)

type LoginReq struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

func (s *LoginReq) Validate() error {
	if s.Email == "" {
		return fmt.Errorf("email is required")
	}
	if s.Password == "" {
		return fmt.Errorf("password is required")
	}
	return nil
}
