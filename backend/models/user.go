package models

import (
	"golang.org/x/crypto/bcrypt"
	"time"
)

type User struct {
	ID        int       `json:"id"`
	Username  string    `json:"username"`
	Email     string    `json:"email,omitempty"`
	Password  string    `json:"password,omitempty"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

type UserReq struct {
	Username string `json:"username"`
	Email    string `json:"email"`
	Password string `json:"password"`
}

type UserErr struct {
	Username string `json:"username,omitempty"`
	Email    string `json:"email,omitempty"`
	Password string `json:"password,omitempty"`
}

type LoginReq struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type LoginErr struct {
	Email    string `json:"email,omitempty"`
	Password string `json:"password,omitempty"`
}

func NewUser(username, email, password string) (*User, error) {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return nil, err
	}

	now := time.Now().UTC()

	return &User{
		Username:  username,
		Email:     email,
		Password:  string(hashedPassword),
		CreatedAt: now,
		UpdatedAt: now,
	}, nil
}

func (s *User) Update(username, email, password string) error {
	if bcrypt.CompareHashAndPassword([]byte(s.Password), []byte(password)) != nil {
		hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
		if err != nil {
			return err
		}

		s.Password = string(hashedPassword)
	}

	s.Username = username
	s.Email = email
	s.UpdatedAt = time.Now().UTC()

	return nil
}

func (s *User) OmitSensitive() *User {
	return &User{
		ID:        s.ID,
		Username:  s.Username,
		CreatedAt: s.CreatedAt,
	}
}

func (s *UserReq) Validate() *UserErr {
	userErr := UserErr{}

	if s.Username == "" {
		userErr.Username = "Username is required"
	}

	if s.Email == "" {
		userErr.Email = "Email is required"
	}

	if s.Password == "" {
		userErr.Password = "Password is required"
	} else if len(s.Password) < 8 {
		userErr.Password = "Password must be at least 8 characters"
	}

	if userErr != (UserErr{}) {
		return &userErr
	}

	return nil
}

func (s *LoginReq) Validate() *LoginErr {
	loginErr := LoginErr{}

	if s.Email == "" {
		loginErr.Email = "Email is required"
	}

	if s.Password == "" {
		loginErr.Password = "Password is required"
	}

	if loginErr != (LoginErr{}) {
		return &loginErr
	}

	return nil
}
