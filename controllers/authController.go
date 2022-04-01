package controllers

import (
	models "anya-day/models/sql"
	wmodels "anya-day/models/web"
	"anya-day/token"
	"net/http"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

// Register godoc
// @Summary Register a user
// @Description registering a new user
// @Tags Auth
// @Param Body body wmodels.RegisterInput true "the body to register a user"
// @Produce json
// @Success 200 {object} wmodels.RegisterResp
// @Router /register [post]
func Register(c *gin.Context) {
	db := c.MustGet("db").(*gorm.DB)
	var input wmodels.RegisterInput

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
// @Param Body body wmodels.LoginInput true "enter valid user's credential"
// @Produce json
// @Success 200 {object} wmodels.LoginResp
// @Router /login [post]
func Login(c *gin.Context) {
	db := c.MustGet("db").(*gorm.DB)
	var input wmodels.LoginInput

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

	c.JSON(http.StatusOK, wmodels.LoginResp{
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
// @Param Body body wmodels.ChangePwInput true "Entry existing user valid credentials and the new one."
// @Param Authorization header string true "Authorization. How to input in swagger : 'Bearer <insert_your_jwt_token_here>'"
// @Security BearerToken
// @Produce json
// @Success 200 {object} map[string]interface{}
// @Router /changepw [post]
func ChangePw(c *gin.Context) {
	db := c.MustGet("db").(*gorm.DB)
	input := &wmodels.ChangePwInput{}

	if err := c.ShouldBindJSON(input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	uid, err := token.ExtractUID(c)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}

	uc := &models.UserCredential{
		Password: input.Password,
		UserID:   int(uid),
	}

	if err := uc.AttemptChangePw(input.NewPassword, db); err != nil {
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
