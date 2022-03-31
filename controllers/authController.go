package controllers

import (
	"anya-day/models"
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

	c.JSON(http.StatusOK, LoginResp{
		Status:  "success",
		Message: "berhasil login",
		User:    input.Username,
		Token:   "",
	})
}
