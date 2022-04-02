package models

type (
	Category struct {
		ID   uint   `json:"id"`
		Name string `json:"name"`
	}
	CategoryInput struct {
		Name string `binding:"required"`
	}
)
