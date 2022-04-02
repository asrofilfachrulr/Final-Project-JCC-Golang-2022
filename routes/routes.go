package routes

import (
	"anya-day/controllers"
	"anya-day/middleware"
	models "anya-day/models/sql"
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

	// attach super user id to context
	r.Use(func(c *gin.Context) {
		var roles []models.Role
		db.Where("name = ?", "dev").Find(&roles)

		superMap := make(map[uint]string)
		for _, u := range roles {
			superMap[u.UserID] = u.Name
		}

		c.Set("super_map", superMap)
	})

	// /auth
	r.POST("/register", controllers.Register)
	r.POST("/login", controllers.Login)

	// /user
	user := r.Group("/user")
	user.Use(middleware.JWTAuthMiddleware())
	user.PUT("/changepw", controllers.ChangePw)

	user.GET("/profile", controllers.GetCompleteUser)
	user.PUT("/profile", controllers.UpdateProfile)
	user.DELETE("/profile", controllers.DeleteUser)

	user.POST("/address", controllers.PostAddress)
	user.PUT("/address", controllers.UpdateAddress)

	user.PATCH("/role", controllers.ChangeUserRole)

	// swagger route
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	return r
}
