package models

type LoginReq struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type LoginErr struct {
	Email    string `json:"email,omitempty"`
	Password string `json:"password,omitempty"`
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
