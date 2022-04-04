package models

type (
	PostReview struct {
		Review string  `json:"review" binding:"required"`
		Rating float32 `json:"rating" binding:"required"`
	}
	Review struct {
		Username string  `json:"username"`
		Review   string  `json:"review"`
		Rating   float32 `json:"rating"`
	}
)
