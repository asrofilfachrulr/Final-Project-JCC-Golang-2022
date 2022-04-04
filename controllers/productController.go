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
// @Summary Post new products of a merchant.
// @Description Post new products of a merchant.
// @Param id path uint true "merchant id"
// @Param Body body wmodels.ProductInput true "Enter Product Details"
// @Param Authorization header string true "Authorization. How to input in swagger : 'Bearer <insert_your_jwt_token_here>'"
// @Security BearerToken
// @Tags Product
// @Produce json
// @Success 200 {object} utils.RespWithData{data=wmodels.ProductOutput}
// @Router /merchants/{id}/products [POST]
func CreateProduct(c *gin.Context) {
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

	var input wmodels.ProductInput
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	p := &models.Product{
		Name:       input.Name,
		Desc:       input.Desc,
		MerchantID: merchantId,
		Price:      input.Price,
		Stock:      input.Stock,
		CategoryID: input.CategoryID,
	}

	if err := p.PostProduct(db); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(http.StatusCreated, utils.RespWithData{
		Status:  "success",
		Message: "berhasil menambahkan produk",
		Data: wmodels.ProductOutput{
			ID:     p.ID,
			Name:   p.Name,
			Price:  p.Price,
			Rating: p.Rating,
		},
	})
}

// Register godoc
// @Summary Get all avalaible products of a merchant.
// @Description Get all avalaible products of a merchant.
// @Param id path uint true "merchant id"
// @Param _ query wmodels.ProductFilter false "filter products"
// @Tags Product
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

// Register godoc
// @Summary Get detailed product.
// @Description Get detailed product.
// @Param id path uint true "merchant id"
// @Param productId path uint true "product id"
// @Tags Product
// @Produce json
// @Success 200 {object} utils.RespWithData{data=[]wmodels.ProductDetailOutput}
// @Router /merchants/{id}/products/{productId} [GET]
func GetProductDetailById(c *gin.Context) {
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

	detailp := &wmodels.MerchantProductOutput{}
	if err := models.GetDetailedProduct(db, detailp, &merchant, &product); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, utils.RespWithData{
		Status:  "success",
		Message: "sukses mendapat detail produk",
		Data:    *detailp,
	})
}

// Register godoc
// @Summary remove a product.
// @Description remove a product.
// @Param id path uint true "merchant id"
// @Param productId path uint true "product id"
// @Param Authorization header string true "Authorization. How to input in swagger : 'Bearer <insert_your_jwt_token_here>'"
// @Security BearerToken
// @Tags Product
// @Produce json
// @Success 200 {object} utils.RespWithData
// @Router /merchants/{id}/products/{productId} [DELETE]
func DeleteProductById(c *gin.Context) {
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
				"error": "product id not found",
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

	product.MerchantID = merchantId
	if err := models.DeleteProductById(db, &product); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, utils.RespWithData{
		Status:  "success",
		Message: "sukses menghapus produk",
		Data: gin.H{
			"id":   product.ID,
			"name": product.Name,
		},
	})
}
