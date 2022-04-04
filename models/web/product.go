package models

type (
	MerchantProductOutput struct {
		ID      uint                `json:"id"`
		Name    string              `json:"name"`
		Product ProductDetailOutput `json:"product"`
	}
	ProductOutput struct {
		ID     uint    `json:"id"`
		Name   string  `json:"name"`
		Price  uint    `json:"price"`
		Rating float32 `json:"rating"`
	}
	ProductDetailOutput struct {
		Name   string  `json:"name"`
		Price  uint    `json:"price"`
		Stock  uint    `json:"stock"`
		Rating float32 `json:"rating,omitempty"`
		Desc   string  `json:"description"`
	}
	ProductFilter struct {
		Price  *string `query:"price"`
		Rating *string `query:"rating"`
		Name   *string `query:"name"`
	}
)
