package web

type RegisterUserInput struct {
	Fullname    string `json:"fullname" binding:"required"`
	Username    string `json:"username" binding:"required" validate:"min=4"`
	Email       string `json:"email" binding:"required" validate:"email"`
	PhoneNumber string `json:"phone_number" binding:"required" validate:"e164,min=11"`
	AddressLine string `json:"address_line"`
	Password    string `json:"password" binding:"required" validate:"passwd"`
}

type RegisterUserOutput struct {
	ID       uint   `json:"id"`
	Username string `json:"username"`
	Email    string `json:"email"`
}
