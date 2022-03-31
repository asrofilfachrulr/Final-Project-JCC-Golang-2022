package routes

import (
	"anya-day/controllers"
	"anya-day/middleware"
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
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	r.Use(func(c *gin.Context) {
		c.Set("db", db)
	})

	// auth
	r.POST("/register", controllers.Register)
	r.POST("/login", controllers.Login)

	m := r.Use(middleware.JWTAuthMiddleware())

	m.POST("/changepw", controllers.ChangePw)

	// swagger route

	return r
}
