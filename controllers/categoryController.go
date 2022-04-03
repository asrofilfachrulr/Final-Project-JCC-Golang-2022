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
// @Summary retrieve list of categories.
// @Description Get all of existing categories.
// @Tags Category
// @Produce json
// @Success 200 {object} utils.RespWithData{data=models.Category}
// @Router /categories [GET]
func GetCategories(c *gin.Context) {
	db := c.MustGet("db").(*gorm.DB)

	cat := &[]models.Category{}
	db.Find(cat)

	c.JSON(http.StatusOK, utils.RespWithData{
		Status:  "success",
		Message: "berhasil mendapat daftar kategori",
		Data:    *cat,
	})
}

// Register godoc
// @Summary [RESTRICTED] Create new category.
// @Description Create new category.
// @Tags Dev/Category
// @Param Body body wmodels.CategoryInput true "Insert new category name."
// @Param Authorization header string true "Authorization. How to input in swagger : 'Bearer <insert_your_jwt_token_here>'"
// @Security BearerToken
// @Produce json
// @Success 201 {object} utils.RespWithData{data=wmodels.Category}
// @Router /dev/categories [POST]
func DevCreateCategory(c *gin.Context) {
	db := c.MustGet("db").(*gorm.DB)

	var input wmodels.CategoryInput
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	cat := &models.Category{
		Name: input.Name,
	}
	if err := db.Create(cat).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(http.StatusCreated, utils.RespWithData{
		Status:  "success",
		Message: "sukses menambah kategori",
		Data: wmodels.Category{
			ID:   cat.ID,
			Name: cat.Name,
		},
	})
}

// Register godoc
// @Summary [RESTRICTED] Delete an existing category.
// @Description Delete an existing category.
// @Tags Dev/Category
// @Param id path uint true "category id"
// @Param Authorization header string true "Authorization. How to input in swagger : 'Bearer <insert_your_jwt_token_here>'"
// @Security BearerToken
// @Produce json
// @Success 200 {object} utils.RespWithData{data=wmodels.Category}
// @Router /dev/categories/{id} [DELETE]
func DevDeleteCategoryById(c *gin.Context) {
	db := c.MustGet("db").(*gorm.DB)

	var id uint
	if v := c.Param("id"); v == "" {
		return
	} else {
		if n, err := utils.StringToUint(v); err != nil {
			return
		} else {
			id = n
		}
	}

	cat := &models.Category{
		ID: id,
	}

	if err := db.Find(cat).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"error": "id not found",
		})
		return
	}

	if err := db.Unscoped().Delete(cat).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, utils.RespWithData{
		Status:  "success",
		Message: "sukses menghapus kategori",
		Data: wmodels.Category{
			ID:   cat.ID,
			Name: cat.Name,
		},
	})
}

// Register godoc
// @Summary [RESTRICTED] Update an existing category.
// @Description Update an existing category.
// @Tags Dev/Category
// @Param id path uint true "category id"
// @Param Body body wmodels.CategoryInput true "Insert new category name"
// @Param Authorization header string true "Authorization. How to input in swagger : 'Bearer <insert_your_jwt_token_here>'"
// @Security BearerToken
// @Produce json
// @Success 200 {object} utils.RespWithData{data=wmodels.Category}
// @Router /dev/categories/{id} [PUT]
func DevUpdateCategoryById(c *gin.Context) {
	db := c.MustGet("db").(*gorm.DB)

	var id uint
	if v := c.Param("id"); v == "" {
		return
	} else {
		if n, err := utils.StringToUint(v); err != nil {
			return
		} else {
			id = n
		}
	}

	cat := &models.Category{
		ID: id,
	}

	if err := db.Find(cat).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"error": "id not found",
		})
		return
	}

	var input wmodels.CategoryInput
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	cat.Name = input.Name
	if err := db.Save(cat).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, utils.RespWithData{
		Status:  "success",
		Message: "sukses mengubah kategori",
		Data: wmodels.Category{
			ID:   cat.ID,
			Name: cat.Name,
		},
	})
}
