package routes

import (
	"anya-day/controllers"
	"anya-day/utils"

	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
	"gorm.io/gorm"
)

func InitRoute(db *gorm.DB) *gin.Engine {
	r := gin.Default()

	r.GET("/", func(ctx *gin.Context) {
		ctx.Data(200, "text/plain", []byte(utils.GetEnvWithFallback("HELLO", "hello")))
	})

	r.Use(func(c *gin.Context) {
		c.Set("db", db)
	})

	// auth
	r.POST("/register", controllers.Register)
	r.POST("/login", controllers.Login)
	r.POST("/changepw", controllers.ChangePw)

	// swagger route
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	return r
}
