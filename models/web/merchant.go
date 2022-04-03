package models

type (
	IDTemplate struct {
		ID   uint   `json:"id"`
		Name string `json:"name"`
	}
	MerchantCreateInput struct {
		Name        string `json:"name" binding:"required"`
		Country     uint   `json:"country" binding:"required"`
		City        string `json:"city"`
		AddressLine string `json:"address_line"`
	}
	MerchantOutput struct {
		ID     uint    `json:"id"`
		Name   string  `json:"name"`
		Rating float32 `json:"rating"`
		City   string  `json:"city,omitempty"`
	}
	MerchantDetailsOutput struct {
		ID      uint               `json:"id"`
		Name    string             `json:"name"`
		Rating  float32            `json:"rating"`
		Address MerchantAddrOutput `json:"address"`
	}

	MerchantAddrOutput struct {
		City        string `json:"city"`
		AddressLine string `json:"address_line"`
		Country     string `json:"country"`
	}

	MerchantFilter struct {
		Name   *string `query:"name"`
		City   *string `query:"city" example:"Bandung"`
		Rating *string `query:"rating" example:"4.5"`
	}
)
