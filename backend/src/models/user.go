package models

import (
	"time"
)

type User struct {
	ID        int       `json:"id"`
	Username  string    `json:"username"`
	Email     string    `json:"email"`
	Password  string    `json:"password"`
	CreatedAt time.Time `json:"created_at"`
}

type UserInfo struct {
	ID        int       `json:"id"`
	Username  string    `json:"username"`
	CreatedAt time.Time `json:"created_at"`
}

func (s *User) ToUserInfo() *UserInfo {
	return &UserInfo{
		ID:        s.ID,
		Username:  s.Username,
		CreatedAt: s.CreatedAt,
	}
}
