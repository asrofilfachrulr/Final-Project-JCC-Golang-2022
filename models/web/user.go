package models

type (
	ChangePwInput struct {
		Password    string `json:"password" binding:"required"`
		NewPassword string `json:"new_password" binding:"required"`
	}
	UpdateProfileInput struct {
		FullName string `json:"full_name"`
		Username string `json:"username"`
		Email    string `json:"email"`
	}
)
