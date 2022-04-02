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
// @Summary create complete address of an existing user.
// @Description Create fresh complete address of an existing user. If you want to update just use the PUT instead which provides you only update what field you like.
// @Tags User
// @Param Body body wmodels.AddressInput true "Insert new address, postal code isn't required."
// @Param Authorization header string true "Authorization. How to input in swagger : 'Bearer <insert_your_jwt_token_here>'"
// @Security BearerToken
// @Produce json
// @Success 201 {object} utils.RespWithData{data=wmodels.AddressRespData}
// @Router /user/address [post]
func PostAddress(c *gin.Context) {
	db := c.MustGet("db").(*gorm.DB)

	uid, err := token.ExtractUID(c)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}

	var input wmodels.AddressInput
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	uaddr := &models.UserAddress{
		UserID: uid,
	}

	if err := uaddr.PostAddress(db, &input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	data := wmodels.AddressRespData{
		UserID:      uaddr.UserID,
		AddressLine: uaddr.AddressLine,
		City:        uaddr.City,
		Country:     input.Country,
		PhoneNumber: uaddr.PhoneNumber,
		PostalCode:  uaddr.PostalCode,
	}

	c.JSON(http.StatusCreated, utils.RespWithData{
		Status:  "success",
		Message: "sukses menambahkan alamat",
		Data:    data,
	})
}

// Register godoc
// @Summary update address of an existing user.
// @Description Update selected address field of an existing user.
// @Tags User
// @Param Body body wmodels.AddressInputNotBinding true "Update address field you like. Remove that you won't."
// @Param Authorization header string true "Authorization. How to input in swagger : 'Bearer <insert_your_jwt_token_here>'"
// @Security BearerToken
// @Produce json
// @Success 201 {object} utils.NormalResp
// @Router /user/address [put]
func UpdateAddress(c *gin.Context) {
	db := c.MustGet("db").(*gorm.DB)

	uid, err := token.ExtractUID(c)
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{
			"error": err.Error(),
		})
		return
	}

	var input wmodels.AddressInputNotBinding
	if err := utils.ParseFromJSON(c.Request.Body, &input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	uaddr := &models.UserAddress{
		UserID: uid,
	}
	if err := uaddr.UpdateAddress(db, &input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, utils.NormalResp{
		Status:  "success",
		Message: "berhasil memperbarui info alamat user",
	})
}

// Register godoc
// @Summary change an existing user profile information.
// @Description Change an existing user profile information which consists of email, fullname, and username. Email format following RFC 5322 format. For update address info instead, please visit anya-day.herokuapp.com/user/profile/address
// @Tags User
// @Param Body body wmodels.UpdateProfileInput true "Only insert profile aspect need to be updated. Inserted value may lead to error for some reasons such updating to used username"
// @Param Authorization header string true "Authorization. How to input in swagger : 'Bearer <insert_your_jwt_token_here>'"
// @Security BearerToken
// @Produce json
// @Success 200 {object} utils.NormalResp
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

// Register godoc
// @Summary change an existing user password
// @Description  Attempt change password from an existing user
// @Tags User
// @Param Body body wmodels.ChangePwInput true "Entry existing user valid credentials and the new one."
// @Param Authorization header string true "Authorization. How to input in swagger : 'Bearer <insert_your_jwt_token_here>'"
// @Security BearerToken
// @Produce json
// @Success 200 {object} utils.NormalResp
// @Router /user/changepw [put]
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

	c.JSON(http.StatusOK, utils.NormalResp{
		Status:  "success",
		Message: "berhasil mengganti password",
	})
}
