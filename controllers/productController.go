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
// @Param id path uint true "merchant id"
// @Tags Merchant
// @Produce json
// @Success 200 {object} utils.RespWithData{data=[]wmodels.ProductOutput}
// @Router /merchants/{id}/products [GET]
func GetProductsByMerchantId(c *gin.Context) {
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

	mproducts := &[]wmodels.ProductOutput{}
	if err := models.GetMerchantProducts(db, mproducts, &merchant); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, utils.RespWithData{
		Status:  "success",
		Message: "sukses mendapat daftar produk",
		Data:    *mproducts,
	})
}
