package models

type (
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
