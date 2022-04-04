package models

type (
	ProductOutput struct {
		ID     uint    `json:"id"`
		Name   string  `json:"name"`
		Price  uint    `json:"price"`
		Rating float32 `json:"rating"`
	}
	ProductDetailOutput struct {
		Name   string `json:"name"`
		Price  uint   `json:"price"`
		Stock  uint   `json:"stock"`
		Rating uint   `json:"rating,omitempty"`
		Desc   string `json:"description"`
	}
	ProductFilter struct {
		Price  *string `query:"price"`
		Rating *string `query:"rating"`
		Name   *string `query:"name"`
	}
)
