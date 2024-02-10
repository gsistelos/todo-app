package models

import (
	"fmt"
	"time"
)

type CreateUserReq struct {
	Username string `json:"username"`
	Email    string `json:"email"`
	Password string `json:"password"`
}

type UpdateUserReq struct {
	Username string `json:"username"`
	Email    string `json:"email"`
	Password string `json:"password"`
}

type User struct {
	ID        int       `json:"id"`
	Username  string    `json:"username"`
	Email     string    `json:"email"`
	Password  string    `json:"password"`
	CreatedAt time.Time `json:"created_at"`
}

func (s *CreateUserReq) Validate() error {
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

func (s *UpdateUserReq) Validate() error {
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

func NewUser(username, email, password string) *User {
	return &User{
		Username:  username,
		Email:     email,
		Password:  password,
		CreatedAt: time.Now().UTC(),
	}
}
