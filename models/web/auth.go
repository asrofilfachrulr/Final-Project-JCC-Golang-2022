package models

type (
	LoginInput struct {
		Username string `json:"username" binding:"required"`
		Password string `json:"password" binding:"required"`
	}

	RegisterInput struct {
		Username string `json:"username" binding:"required"`
		Password string `json:"password" binding:"required"`
		FullName string `json:"full_name" binding:"required"`
		Email    string `json:"email" binding:"required"`
	}

	LoginResp struct {
		Status  string `json:"status"`
		Message string `json:"message"`
		User    string `json:"user"`
		Token   string `json:"token"`
	}
	RegisterResp struct {
		Message string `json:"message"`
		User    struct {
			Username string `json:"username"`
			Email    string `json:"email"`
		} `json:"user"`
	}
)
