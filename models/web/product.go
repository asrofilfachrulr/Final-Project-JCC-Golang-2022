package models

type (
	ProductOutput struct {
		ID    uint   `json:"id"`
		Name  string `json:"name"`
		Price uint   `json:"price"`
	}
	ProductDetailOutput struct {
		Name  string `json:"name"`
		Price uint   `json:"price"`
		Stock uint   `json:"stock"`
		Desc  string `json:"description"`
	}
)
