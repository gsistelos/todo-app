package models

type RegisterReq struct {
	Username        string `json:"username"`
	Email           string `json:"email"`
	Password        string `json:"password"`
	ConfirmPassword string `json:"confirm_password"`
}

type RegisterErr struct {
	Username        string `json:"username,omitempty"`
	Email           string `json:"email,omitempty"`
	Password        string `json:"password,omitempty"`
	ConfirmPassword string `json:"confirm_password,omitempty"`
}

func (s *RegisterReq) Validate() *RegisterErr {
	registerErr := RegisterErr{}

	if s.Username == "" {
		registerErr.Username = "Username is required"
	}

	if s.Email == "" {
		registerErr.Email = "Email is required"
	}

	if s.Password == "" {
		registerErr.Password = "Password is required"
	} else if len(s.Password) < 8 {
		registerErr.Password = "Password must be at least 8 characters"
	}

	if s.ConfirmPassword == "" {
		registerErr.ConfirmPassword = "Confirm password is required"
	} else if s.Password != s.ConfirmPassword {
		registerErr.ConfirmPassword = "Passwords do not match"
	}

	if registerErr != (RegisterErr{}) {
		return &registerErr
	}

	return nil
}
