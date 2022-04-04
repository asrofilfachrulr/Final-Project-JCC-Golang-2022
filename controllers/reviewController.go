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
// @Summary post a review of a product.
// @Description post a review of a product.
// @Param id path uint true "merchant id"
// @Param Body body wmodels.PostReview true "Post a review"
// @Param productId path uint true "product id"
// @Param Authorization header string true "Authorization. How to input in swagger : 'Bearer <insert_your_jwt_token_here>'"
// @Security BearerToken
// @Tags Review
// @Produce json
// @Success 201 {object} utils.NormalResp
// @Router /merchants/{id}/products/{productId}/review [POST]
func CreateProductReview(c *gin.Context) {
	db := c.MustGet("db").(*gorm.DB)

	uid, err := token.ExtractUID(c)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

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

	var productId uint = 0
	if v := c.Param("productId"); v == "" {
		return
	} else {
		if n, err := utils.StringToUint(v); err != nil {
			c.JSON(http.StatusNotFound, gin.H{
				"error": "merchant id not found",
			})
			return
		} else {
			productId = n
		}
	}

	// look up product
	product := models.Product{}
	product.ID = productId

	if err := db.First(&product).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"error": err.Error(),
		})
		return
	}

	review := models.ProductReview{
		UserID:    uid,
		ProductID: product.ID,
	}

	var input wmodels.PostReview

	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	review.Rating = input.Rating
	review.Review = input.Review

	if err := review.PostReview(db); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(http.StatusCreated, utils.NormalResp{
		Status:  "success",
		Message: "berhasil membuat review",
	})
}

// Register godoc
// @Summary get list of reviews of a product.
// @Description get list of reviews of a product.
// @Param id path uint true "merchant id"
// @Param productId path uint true "product id"
// @Tags Review
// @Produce json
// @Success 200 {object} utils.RespWithData
// @Router /merchants/{id}/products/{productId}/review [GET]
func GetReview(c *gin.Context) {
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

	var productId uint = 0
	if v := c.Param("productId"); v == "" {
		return
	} else {
		if n, err := utils.StringToUint(v); err != nil {
			c.JSON(http.StatusNotFound, gin.H{
				"error": "merchant id not found",
			})
			return
		} else {
			productId = n
		}
	}

	// look up product
	product := models.Product{}
	product.ID = productId

	if err := db.First(&product).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"error": err.Error(),
		})
		return
	}

	review := &[]wmodels.Review{}
	// lookup review
	if err := models.GetReview(db, review, product.ID); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, utils.RespWithData{
		Status:  "success",
		Message: "sukses mendapatkan daftar review",
		Data:    *review,
	})
}
