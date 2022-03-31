package controllers

import (
	"anya-day/models"
	"anya-day/token"
	"net/http"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

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
	ChangePwInput struct {
		Password    string `json:"password" binding:"required"`
		NewPassword string `json:"new_password" binding:"required"`
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

// Register godoc
// @Summary Register a user
// @Description registering a new user
// @Tags Auth
// @Param Body body RegisterInput true "the body to register a user"
// @Produce json
// @Success 200 {object} RegisterResp
// @Router /register [post]
func Register(c *gin.Context) {
	db := c.MustGet("db").(*gorm.DB)
	var input RegisterInput

	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	uc := &models.UserCredential{
		Username: input.Username,
		Password: input.Password,
	}

	u := &models.User{
		Email:    input.Email,
		Username: input.Username,
		FullName: input.FullName,
	}

	uc.ConvToHash()

	err := u.SaveUser(db)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	err = uc.SaveCredential(u, db)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	user := map[string]string{
		"username": input.Username,
		"email":    input.Email,
	}

	c.JSON(http.StatusOK, gin.H{"message": "registration success", "user": user})

}

// Register godoc
// @Summary Login with an existing user
// @Description Login as existing user with valid credentials. Success attempt will return JWT token.
// @Tags Auth
// @Param Body body LoginInput true "enter valid user's credential"
// @Produce json
// @Success 200 {object} LoginResp
// @Router /login [post]
func Login(c *gin.Context) {
	db := c.MustGet("db").(*gorm.DB)
	var input LoginInput

	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	uc := &models.UserCredential{
		Username: input.Username,
		Password: input.Password,
	}

	if err := uc.AttemptLogin(db); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	t, err := token.GenerateToken(uc.UserID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "server crashed when creating your token",
		})
		return
	}

	c.JSON(http.StatusOK, LoginResp{
		Status:  "success",
		Message: "berhasil login",
		User:    input.Username,
		Token:   t,
	})
}

// Register godoc
// @Summary change an existing user password
// @Description  Attempt change password from an existing user
// @Tags Auth
// @Param Body body ChangePwInput true "Entry existing user valid credentials and the new one."
// @Param Authorization header string true "Authorization. How to input in swagger : 'Bearer <insert_your_jwt_token_here>'"
// @Security BearerToken
// @Produce json
// @Success 200 {object} map[string]interface{}
// @Router /changepw [post]
func ChangePw(c *gin.Context) {
	db := c.MustGet("db").(*gorm.DB)
	input := &ChangePwInput{}

	if err := c.ShouldBindJSON(input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	uc := &models.UserCredential{
		Password: input.Password,
	}

	if err := uc.AttemptChangePw(input.NewPassword, db, c); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"status":  "success",
		"message": "berhasil mengganti password",
	})
}
