package middleware

import (
	"anya-day/token"
	"net/http"

	"github.com/gin-gonic/gin"
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
			return
		}

		_dMap := ctx.MustGet("_devMap").(map[uint]string)
		if _, ok := _dMap[uid]; !ok {
			ctx.JSON(http.StatusUnauthorized, gin.H{
				"error": "you're not allowed to request",
			})
			return
		}

		ctx.Next()
	}
}
