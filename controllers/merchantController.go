package controllers

import (
	models "anya-day/models/sql"
	wmodels "anya-day/models/web"
	"anya-day/utils"

	"net/http"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

// Register godoc
// @Summary Get all avalaible merchants.
// @Description Get all avalaible merchants.
// @Param _ query wmodels.MerchantFilter false "filter merchants"
// @Tags Merchant
// @Produce json
// @Success 200 {object} utils.RespWithData{data=wmodels.MerchantOutput}
// @Router /merchants [GET]
func GetMerchants(c *gin.Context) {
	db := c.MustGet("db").(*gorm.DB)

	var merchants []models.Merchant
	var merchantsOutput []wmodels.MerchantOutput

	filterName := c.Query("name")
	filterCity := c.Query("city")
	fitlerRating := c.Query("rating")

	filter := &wmodels.MerchantFilter{
		Name:   &filterName,
		City:   &filterCity,
		Rating: &fitlerRating,
	}

	if err := models.GetAll(db, filter, merchants, &merchantsOutput); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, utils.RespWithData{
		Status:  "success",
		Message: "berhasil mendapatkan daftar merchant",
		Data:    merchantsOutput,
	})
}

// Register godoc
// @Summary Get specific merchant.
// @Description Get specific merchant.
// @Param id path uint true "merchant id"
// @Tags Merchant
// @Produce json
// @Success 200 {object} utils.RespWithData{data=wmodels.MerchantDetailsOutput}
// @Router /merchants/{id} [GET]
func GetMerchantById(c *gin.Context) {
	db := c.MustGet("db").(*gorm.DB)

	var merchantId uint = 0
	if v := c.Param("id"); v == "" {
		return
	} else {
		if n, err := utils.StringToUint(v); err != nil {
			c.JSON(http.StatusNotFound, gin.H{
				"error": "merchant id not found",
			})
			return
		} else {
			merchantId = n
		}
	}

	// look up merchant
	merchant := models.Merchant{}
	merchant.ID = merchantId

	if err := db.First(&merchant).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"error": err.Error(),
		})
		return
	}

	data := &wmodels.MerchantDetailsOutput{}
	if err := merchant.GetById(db, data); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, utils.RespWithData{
		Status:  "success",
		Message: "berhasil mendapatkan merchant",
		Data:    *data,
	})
}

// Register godoc
// @Summary Delete an existing merchant.
// @Description Delete an existing merchant.
// @Param id path uint true "merchant id"
// @Tags Merchant
// @Param Authorization header string true "Authorization. How to input in swagger : 'Bearer <insert_your_jwt_token_here>'"
// @Security BearerToken
// @Produce json
// @Success 200 {object} utils.RespWithData{data=wmodels.MerchantOutput}
// @Router /merchants/{id} [DELETE]
func DeleteMerchantById(c *gin.Context) {
	db := c.MustGet("db").(*gorm.DB)

	var merchantId uint = 0
	if v := c.Param("id"); v == "" {
		return
	} else {
		if n, err := utils.StringToUint(v); err != nil {
			c.JSON(http.StatusNotFound, gin.H{
				"error": "merchant id not found",
			})
			return
		} else {
			merchantId = n
		}
	}

	// look up merchant
	merchant := models.Merchant{}
	merchant.ID = merchantId

	if err := db.First(&merchant).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"error": err.Error(),
		})
		return
	}

	if err := merchant.Delete(db); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, utils.RespWithData{
		Status:  "success",
		Message: "sukses mengapus merchant",
		Data: wmodels.MerchantOutput{
			ID:     merchant.ID,
			Name:   merchant.Name,
			Rating: merchant.Rating,
		},
	})
}
