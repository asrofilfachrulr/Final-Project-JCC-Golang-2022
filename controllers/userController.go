package controllers

import (
	models "anya-day/models/sql"
	wmodels "anya-day/models/web"
	"anya-day/token"
	"anya-day/utils"
	"net/http"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

// Register godoc
// @Summary change an existing user profile information.
// @Description Change an existing user profile information which consists of email, fullname, and username. Email format following RFC 5322 format. For update address info instead, please visit anya-day.herokuapp.com/user/profile/address
// @Tags User
// @Param Body body wmodels.UpdateProfileInput true "Only insert profile aspect need to be updated. Inserted value may lead to error for some reasons such updating to used username"
// @Param Authorization header string true "Authorization. How to input in swagger : 'Bearer <insert_your_jwt_token_here>'"
// @Security BearerToken
// @Produce json
// @Success 200 {object} map[string]interface{}
// @Router /user/profile [put]
func UpdateProfile(c *gin.Context) {
	db := c.MustGet("db").(*gorm.DB)

	uid, err := token.ExtractUID(c)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}

	input := &wmodels.UpdateProfileInput{}
	if err := utils.ParseFromJSON(c.Request.Body, input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	user := &models.User{}
	db.First(user, uid)
	if err := user.UpdateUserProfile(db, input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, utils.NormalResp{
		Status:  "success",
		Message: "berhasil memperbarui profile",
	})
}
