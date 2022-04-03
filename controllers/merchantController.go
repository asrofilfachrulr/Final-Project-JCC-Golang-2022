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
// @Tags Merchant
// @Produce json
// @Success 200 {object} utils.RespWithData{data=wmodels.MerchantOutput}
// @Router /merchants [GET]
func GetMerchants(c *gin.Context) {
	db := c.MustGet("db").(*gorm.DB)

	var merchants []models.Merchant
	var merchantsOutput []wmodels.MerchantOutput

	if err := models.GetMerchants(db, merchants, &merchantsOutput); err != nil {
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
