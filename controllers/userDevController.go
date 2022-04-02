package controllers

import (
	models "anya-day/models/sql"
	wmodels "anya-day/models/web"
	"anya-day/utils"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

// Register godoc
// @Summary [EXPERIMENTAL] [RESTRICTED] get complete profile of an existing user.
// @Description Get complete details of logged user.
// @Tags Dev/User
// @Param id path uint true "user id"
// @Param Authorization header string true "Authorization. How to input in swagger : 'Bearer <insert_your_jwt_token_here>'"
// @Security BearerToken
// @Produce json
// @Success 200 {object} wmodels.UserCompleteDataResp
// @Router /dev/user/{id}/profile [GET]
func DevGetCompleteUser(c *gin.Context) {
	db := c.MustGet("db").(*gorm.DB)

	var userId uint = 0
	if v := c.Param("id"); v == "" {
		return
	} else {
		if n, err := utils.StringToUint(v); err != nil {
			c.JSON(http.StatusNotFound, gin.H{
				"error": "user id not valid",
			})
			return
		} else {
			userId = n
		}
	}

	u := &models.User{}
	cu := &wmodels.UserCompleteDataResp{}
	u.ID = userId

	if err := u.GetCompleteUser(db, cu); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, *cu)
}

// Register godoc
// @Summary [EXPERIMENTAL] [RESTRICTED] create complete address of an existing user.
// @Description Create fresh complete address of an existing user. If you want to update just use the PUT instead which provides you only update what field you like.
// @Tags Dev/User
// @Param id path uint true "user id"
// @Param Body body wmodels.AddressInput true "Insert new address, postal code isn't required."
// @Param Authorization header string true "Authorization. How to input in swagger : 'Bearer <insert_your_jwt_token_here>'"
// @Security BearerToken
// @Produce json
// @Success 201 {object} utils.RespWithData{data=wmodels.AddressRespData}
// @Router /dev/user/{id}/address [post]
func DevPostAddress(c *gin.Context) {
	db := c.MustGet("db").(*gorm.DB)

	var userId uint = 0
	if v := c.Param("id"); v == "" {
		return
	} else {
		if n, err := utils.StringToUint(v); err != nil {
			c.JSON(http.StatusNotFound, gin.H{
				"error": "user id not valid",
			})
			return
		} else {
			userId = n
		}
	}

	// user_id lookup
	if err := db.First(&models.User{}, userId).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"error": "user not found!",
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
		UserID: userId,
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
// @Summary [EXPERIMENTAL] [RESTRICTED] update address of an existing user.
// @Description Update selected address field of an existing user.
// @Tags Dev/User
// @Param id path uint true "user id"
// @Param Body body wmodels.AddressInputNotBinding true "Update address field you like. Remove that you won't."
// @Param Authorization header string true "Authorization. How to input in swagger : 'Bearer <insert_your_jwt_token_here>'"
// @Security BearerToken
// @Produce json
// @Success 200 {object} utils.NormalResp
// @Router /dev/user/{id}/address [put]
func DevUpdateAddress(c *gin.Context) {
	db := c.MustGet("db").(*gorm.DB)

	var userId uint = 0
	if v := c.Param("id"); v == "" {
		return
	} else {
		if n, err := utils.StringToUint(v); err != nil {
			c.JSON(http.StatusNotFound, gin.H{
				"error": "user id not valid",
			})
			return
		} else {
			userId = n
		}
	}

	// user_id lookup
	if err := db.First(&models.User{}, userId).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"error": "user not found!",
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
		UserID: userId,
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
// @Summary [EXPERIMENTAL] [RESTRICTED] change an existing user profile information.
// @Description Change an existing user profile information which consists of email, fullname, and username. Email format following RFC 5322 format. For update address info instead, please use PUT user/profile/address instead. This endpoint has not POST method which same behaviour already handled by POST /register
// @Tags Dev/User
// @Param id path uint true "user id"
// @Param Body body wmodels.UpdateProfileInput true "Only insert profile aspect need to be updated. Inserted value may lead to error for some reasons such updating to used username"
// @Param Authorization header string true "Authorization. How to input in swagger : 'Bearer <insert_your_jwt_token_here>'"
// @Security BearerToken
// @Produce json
// @Success 200 {object} utils.NormalResp
// @Router /dev/user/{id}/profile [put]
func DevUpdateProfile(c *gin.Context) {
	db := c.MustGet("db").(*gorm.DB)

	var userId uint = 0
	if v := c.Param("id"); v == "" {
		return
	} else {
		if n, err := utils.StringToUint(v); err != nil {
			c.JSON(http.StatusNotFound, gin.H{
				"error": "user id not valid",
			})
			return
		} else {
			userId = n
		}
	}

	// user_id lookup
	if err := db.First(&models.User{}, userId).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"error": "user not found!",
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
	db.First(user, userId)
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
// @Summary [EXPERIMENTAL] [RESTRICTED] change an existing user password
// @Description  Attempt change password from an existing user
// @Tags Dev/User
// @Param id path uint true "user id"
// @Param Body body wmodels.ChangePwInput true "Entry existing user valid credentials and the new one."
// @Param Authorization header string true "Authorization. How to input in swagger : 'Bearer <insert_your_jwt_token_here>'"
// @Security BearerToken
// @Produce json
// @Success 200 {object} utils.NormalResp
// @Router /dev/user/{id}/changepw [put]
func DevChangePw(c *gin.Context) {
	db := c.MustGet("db").(*gorm.DB)
	input := &wmodels.ChangePwInput{}

	if err := c.ShouldBindJSON(input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	var userId uint = 0
	if v := c.Param("id"); v == "" {
		return
	} else {
		if n, err := utils.StringToUint(v); err != nil {
			c.JSON(http.StatusNotFound, gin.H{
				"error": "user id not valid",
			})
			return
		} else {
			userId = n
		}
	}

	// user_id lookup
	if err := db.First(&models.User{}, userId).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"error": "user not found!",
		})
		return
	}

	uc := &models.UserCredential{
		Password: input.Password,
		UserID:   int(userId),
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

// Register godoc
// @Summary [EXPERIMENTAL] [RESTRICTED] change an existing user role
// @Description  Every patch request to this endpoint lead to switch role between [customer], [merchant]. Be careful, switching from [merchant] to [customer] lead to wipe out all user's merchant data (also its products)
// @Tags Dev/User
// @Param id path uint true "user id"
// @Param Authorization header string true "Authorization. How to input in swagger : 'Bearer <insert_your_jwt_token_here>'"
// @Security BearerToken
// @Produce json
// @Success 200 {object} utils.RespWithData{data=wmodels.RoleDataResp}
// @Router /dev/user/{id}/role [patch]
func DevChangeUserRole(c *gin.Context) {
	db := c.MustGet("db").(*gorm.DB)

	var userId uint = 0
	if v := c.Param("id"); v == "" {
		return
	} else {
		if n, err := utils.StringToUint(v); err != nil {
			c.JSON(http.StatusNotFound, gin.H{
				"error": "user id not valid",
			})
			return
		} else {
			userId = n
		}
	}

	// user_id lookup
	if err := db.First(&models.User{}, userId).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"error": "user not found!",
		})
		return
	}

	r := &models.Role{
		UserID: userId,
	}
	if err := r.ChangeRole(db); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, utils.RespWithData{
		Status:  "success",
		Message: "berhasil memperbarui role user",
		Data: wmodels.RoleDataResp{
			Username: r.User.Username,
			Role:     r.Name,
		},
	})
}

// Register godoc
// @Summary [EXPERIMENTAL] [RESTRICTED] delete an existing user role
// @Description  This request lead to removing all user's related data
// @Tags Dev/User
// @Param id path uint true "user id"
// @Param Authorization header string true "Authorization. How to input in swagger : 'Bearer <insert_your_jwt_token_here>'"
// @Security BearerToken
// @Produce json
// @Success 200 {object} utils.RespWithData{data=wmodels.UserDataResp}
// @Router /dev/user/{id}/profile [delete]
func DevDeleteUser(c *gin.Context) {
	db := c.MustGet("db").(*gorm.DB)

	var userId uint = 0
	if v := c.Param("id"); v == "" {
		return
	} else {
		if n, err := utils.StringToUint(v); err != nil {
			c.JSON(http.StatusNotFound, gin.H{
				"error": "user id not valid",
			})
			log.Println(err.Error())
			return
		} else {
			userId = n
		}
	}

	// user_id lookup
	if err := db.First(&models.User{}, userId).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"error": "user not found!",
		})
		return
	}

	u := &models.User{}
	u.ID = userId

	if err := u.AttemptDeleteUser(db); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, utils.RespWithData{
		Status:  "success",
		Message: "berhasil menghapus user",
		Data: wmodels.UserDataResp{
			Username: u.Username,
			Fullname: u.FullName,
			Email:    u.Email,
		},
	})
}
