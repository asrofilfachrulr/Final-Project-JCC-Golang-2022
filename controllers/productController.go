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
// @Summary Get all avalaible products of a merchant.
// @Description Get all avalaible products of a merchant.
// @Param id path uint true "merchant id"
// @Param _ query wmodels.ProductFilter false "filter products"
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

	filterName := c.Query("name")
	filterPrice := c.Query("price")
	fitlerRating := c.Query("rating")

	filter := &wmodels.ProductFilter{
		Name:   &filterName,
		Price:  &filterPrice,
		Rating: &fitlerRating,
	}

	mproducts := &[]wmodels.ProductOutput{}
	if err := models.GetMerchantProducts(db, mproducts, &merchant, filter); err != nil {
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

// // Register godoc
// // @Summary Get detailed product.
// // @Description Get detailed product.
// // @Param id path uint true "merchant id"
// // @Tags Merchant
// // @Produce json
// // @Success 200 {object} utils.RespWithData{data=[]wmodels.ProductDetailOutput}
// // @Router /merchants/{id}/products/{productId} [GET]
// func GetProductsByProductId(c *gin.Context) {
// 	db := c.MustGet("db").(*gorm.DB)

// 	var merchantId uint = 0
// 	if v := c.Param("id"); v == "" {
// 		return
// 	} else {
// 		if n, err := utils.StringToUint(v); err != nil {
// 			c.JSON(http.StatusNotFound, gin.H{
// 				"error": "merchant id not found",
// 			})
// 			return
// 		} else {
// 			merchantId = n
// 		}
// 	}

// 	// look up merchant
// 	merchant := models.Merchant{}
// 	merchant.ID = merchantId

// 	if err := db.First(&merchant).Error; err != nil {
// 		c.JSON(http.StatusNotFound, gin.H{
// 			"error": err.Error(),
// 		})
// 		return
// 	}

// 	mproducts := &[]wmodels.ProductOutput{}
// 	if err := models.GetMerchantProducts(db, mproducts, &merchant); err != nil {
// 		c.JSON(http.StatusInternalServerError, gin.H{
// 			"error": err.Error(),
// 		})
// 		return
// 	}

// 	c.JSON(http.StatusOK, utils.RespWithData{
// 		Status:  "success",
// 		Message: "sukses mendapat daftar produk",
// 		Data:    *mproducts,
// 	})
// }
