package models

import (
	"fmt"
	"golang.org/x/crypto/bcrypt"
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

func (s *UserReq) HashPassword() error {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(s.Password), bcrypt.DefaultCost)
	if err != nil {
		return err
	}

	s.Password = string(hashedPassword)
	return nil
}
