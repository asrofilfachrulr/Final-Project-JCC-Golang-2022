package middleware

import (
	models "anya-day/models/sql"
	"anya-day/token"
	"anya-day/utils"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func JWTAuthMiddleware() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		if err := token.ValidateToken(ctx); err != nil {
			ctx.JSON(http.StatusUnauthorized, gin.H{
				"error": err.Error(),
			})
			ctx.Abort()
			return
		}
		ctx.Next()
	}
}

func JWTAuthDevMiddleware() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		if err := token.ValidateToken(ctx); err != nil {
			ctx.JSON(http.StatusUnauthorized, gin.H{
				"error": err.Error(),
			})
			ctx.Abort()
			return
		}

		uid, err := token.ExtractUID(ctx)
		if err != nil {
			ctx.JSON(http.StatusUnauthorized, gin.H{
				"error": err.Error(),
			})
			ctx.Abort()
			return
		}

		_dMap := ctx.MustGet("_devMap").(map[uint]string)
		if _, ok := _dMap[uid]; !ok {
			ctx.JSON(http.StatusUnauthorized, gin.H{
				"error": "you're not allowed to continue",
			})
			ctx.Abort()
			return
		}

		ctx.Next()
	}
}

func JWTMerchantMiddleware() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		if err := token.ValidateToken(ctx); err != nil {
			ctx.JSON(http.StatusUnauthorized, gin.H{
				"error": err.Error(),
			})
			ctx.Abort()
			return
		}

		uid, err := token.ExtractUID(ctx)
		if err != nil {
			ctx.JSON(http.StatusUnauthorized, gin.H{
				"error": err.Error(),
			})
			ctx.Abort()
			return
		}

		_dMap := ctx.MustGet("_devMap").(map[uint]string)
		if _, ok := _dMap[uid]; ok {
			ctx.Next()
			return
		}

		log.Println("not a dev")

		db := ctx.MustGet("db").(*gorm.DB)

		var merchant models.Merchant
		merchant.AdminId = uid
		if err := db.Where(&merchant).First(&merchant).Error; err != nil {
			ctx.JSON(http.StatusUnauthorized, gin.H{
				"error": err.Error(),
			})
			ctx.Abort()
			return
		}

		if id, err := utils.StringToUint(ctx.Param("id")); err != nil {
			ctx.JSON(http.StatusUnauthorized, gin.H{
				"error": err.Error(),
			})
			ctx.Abort()
			return
		} else if id != merchant.ID {
			ctx.JSON(http.StatusUnauthorized, gin.H{
				"error": "you are not allowed to access",
			})
			ctx.Abort()
			return
		}

		ctx.Next()
	}
}
