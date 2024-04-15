package models

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
