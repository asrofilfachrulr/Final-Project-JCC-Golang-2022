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

		_devMap := make(map[uint]string)
		for _, u := range roles {
			_devMap[u.UserID] = u.Name
		}

		c.Set("_devMap", _devMap)
	})

	api := r.Group("/api/v1")

	// /auth
	api.POST("/register", controllers.Register)
	api.POST("/login", controllers.Login)

	// /user
	user := api.Group("/user")
	user.Use(middleware.JWTAuthMiddleware())
	user.PUT("/changepw", controllers.ChangePw)
	user.GET("/profile", controllers.GetCompleteUser)
	user.PUT("/profile", controllers.UpdateProfile)
	user.DELETE("/profile", controllers.DeleteUser)
	user.POST("/address", controllers.PostAddress)
	user.PUT("/address", controllers.UpdateAddress)
	user.PATCH("/role", controllers.ChangeUserRole)

	// /user but for dev
	devUser := api.Group("/dev/users/:id")
	devUser.Use(middleware.JWTAuthDevMiddleware())
	devUser.PUT("/changepw", controllers.DevChangePw)
	devUser.GET("/profile", controllers.DevGetCompleteUser)
	devUser.PUT("/profile", controllers.DevUpdateProfile)
	devUser.DELETE("/profile", controllers.DevDeleteUser)
	devUser.POST("/address", controllers.DevPostAddress)
	devUser.PUT("/address", controllers.DevUpdateAddress)
	devUser.PATCH("/role", controllers.DevChangeUserRole)

	// /category
	api.GET("/categories", controllers.GetCategories) // guest can access
	devCategory := api.Group("/dev/categories")
	devCategory.Use(middleware.JWTAuthDevMiddleware())
	devCategory.POST("/", controllers.DevCreateCategory)
	devCategory.PUT("/:id", controllers.DevUpdateCategoryById)
	devCategory.DELETE("/:id", controllers.DevDeleteCategoryById)

	// /merchant
	api.GET("/merchants", controllers.GetMerchants)
	api.GET("/merchants/:id", controllers.GetMerchantById)
	merchant := api.Group("/merchants", middleware.JWTMerchantMiddleware())
	merchant.DELETE("/:id", controllers.DeleteMerchantById)

	// swagger route
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	return r
}
