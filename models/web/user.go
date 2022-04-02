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
	AddressInput struct {
		AddressLine string `json:"address_line" binding:"required"`
		City        string `json:"city" binding:"required"`
		Country     string `json:"country" binding:"required"`
		PhoneNumber uint   `json:"phone_number" binding:"required"`
		PostalCode  uint   `json:"postal_code"`
	}
	AddressRespData struct {
		UserID      uint   `json:"user_id,omitempty"`
		AddressLine string `json:"address"`
		City        string `json:"city"`
		Country     string `json:"country"`
		PhoneNumber uint   `json:"number"`
		PostalCode  uint   `json:"postal_code,omitempty"`
	}
	AddressInputNotBinding struct {
		AddressLine string `json:"address_line"`
		City        string `json:"city"`
		Country     string `json:"country"`
		PhoneNumber uint   `json:"phone_number"`
		PostalCode  uint   `json:"postal_code"`
	}
	RoleDataResp struct {
		Username string `json:"username"`
		Role     string `json:"role"`
	}
	UserDataResp struct {
		Username string `json:"username"`
		Email    string `json:"email"`
		Fullname string `json:"fullname"`
	}
	UserCompleteDataResp struct {
		ID       uint            `json:"id"`
		Username string          `json:"username"`
		Email    string          `json:"email"`
		Fullname string          `json:"fullname"`
		Role     string          `json:"role"`
		Address  AddressRespData `json:"address"`
	}
)
