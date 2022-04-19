package web

type RegisterUserInput struct {
	Fullname    string `json:"fullname" binding:"required"`
	Username    string `json:"username" binding:"required"`
	Email       string `json:"email" binding:"required"`
	PhoneNumber string `json:"phone_number" binding:"required"`
	AddressLine string `json:"address_line"`
	Password    string `json:"password" binding:"required"`
}
